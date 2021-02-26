package cmd

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"

	"github.com/reubenmiller/go-c8y-cli/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/pkg/iterator"
	"github.com/reubenmiller/go-c8y-cli/pkg/progressbar"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

type BatchOptions struct {
	StartIndex        int
	NumJobs           int
	TotalWorkers      int
	Delay             int
	AbortOnErrorCount int

	InputData []string

	inputIndex int
}

func (b *BatchOptions) GetItem() (string, error) {
	defer func() {
		b.inputIndex++
	}()

	if b.useInputData() {
		if b.inputIndex < len(b.InputData) {

			return b.InputData[b.inputIndex], nil
		}
		return "", fmt.Errorf("end of input data")
	}

	if b.inputIndex >= b.NumJobs {
		return "", fmt.Errorf("end of input data")
	}
	return fmt.Sprintf("%d", b.inputIndex+b.StartIndex), nil

}

func (b *BatchOptions) useInputData() bool {
	return b.InputData != nil && len(b.InputData) > 0
}

func getBatchOptions(cmd *cobra.Command) (*BatchOptions, error) {
	options := &BatchOptions{
		AbortOnErrorCount: 10,
		TotalWorkers:      1,
		Delay:             1000,
	}

	if v, err := cmd.Flags().GetInt("count"); err == nil {
		options.NumJobs = v
	}

	if v, err := cmd.Flags().GetInt("startIndex"); err == nil {
		options.StartIndex = v
	}

	if v, err := cmd.Flags().GetInt("delay"); err == nil {
		options.Delay = v
	}

	if v, err := cmd.Flags().GetInt("workers"); err == nil {
		if v > globalFlagBatchMaxWorkers {
			return nil, fmt.Errorf("number of workers exceeds the maximum workers limit of %d", globalFlagBatchMaxWorkers)
		}
		options.TotalWorkers = v
	}

	if v, err := cmd.Flags().GetInt("abortOnErrors"); err == nil {
		options.AbortOnErrorCount = v
	}

	return options, nil
}

type batchArgument struct {
	id            int64
	request       c8y.RequestOptions
	commonOptions CommonCommandOptions
	batchOptions  BatchOptions
}

func processRequestAndResponseWithWorkers(cmd *cobra.Command, r *c8y.RequestOptions, inputIterators *flags.RequestInputIterators) error {
	var err error
	var pathIter iterator.Iterator

	if inputIterators != nil && inputIterators.Total > 0 {
		if inputIterators.Path != nil {
			pathIter = inputIterators.Path
		} else {
			// use continuous path repeater so that it does not stop the other interators
			pathIter = iterator.NewRepeatIterator(r.Path, 0)
		}
		if inputIterators.Body != nil {
			r.Body = inputIterators.Body
		}
	}
	if pathIter == nil {
		pathIter = iterator.NewRepeatIterator(r.Path, 1)
	}
	// Note: Body accepts iterator types, so no need for special handling here
	requestIter := NewRequestIterator(*r, pathIter, inputIterators.Query, r.Body)

	// get common options and batch settings
	commonOptions, err := getCommonOptions(cmd)
	if err != nil {
		return cmderrors.NewUserError(fmt.Sprintf("Failed to get common options. err=%s", err))
	}

	batchOptions, err := getBatchOptions(cmd)
	if err != nil {
		return err
	}

	return runBatched(requestIter, commonOptions, *batchOptions)
}

func runBatched(requestIterator *RequestIterator, commonOptions CommonCommandOptions, batchOptions BatchOptions) error {
	// Two channels - to send them work and to collect their results.
	// buffer size does not really matter, it just needs to be high
	// enough not to block the workers

	// TODO: how to detect when request iterator is finished when using the body iterator (total number of requests?)
	if batchOptions.TotalWorkers < 1 {
		batchOptions.TotalWorkers = 1
	}
	jobs := make(chan batchArgument, batchOptions.TotalWorkers-1)
	results := make(chan error, batchOptions.TotalWorkers-1)
	workers := sync.WaitGroup{}

	progbar := progressbar.NewMultiProgressBar(rootCmd.ErrOrStderr(), 1, batchOptions.TotalWorkers, "requests", globalFlagProgressBar)
	progbar.Start(float64(batchOptions.Delay * 2 / 1000))

	for w := 1; w <= batchOptions.TotalWorkers; w++ {
		Logger.Infof("starting worker: %d", w)
		workers.Add(1)
		go batchWorker(w, jobs, results, progbar, &workers)
	}

	jobID := int64(0)

	// add jobs async
	go func() {
		defer close(jobs)
		jobInputErrors := int64(0)
		for {
			jobID++
			Logger.Infof("checking job iterator: %d", jobID)

			if jobID > globalFlagBatchMaxJobs {
				Logger.Warningf("maximum jobs reached: limit=%d", globalFlagBatchMaxJobs)
				break
			}

			request, err := requestIterator.GetNext()

			if err != nil {
				if errors.Is(err, io.EOF) {
					// no more requests
					break
				}
				jobInputErrors++

				rootCauseErr := err
				if errors.Is(err, ErrNoMatchesFound) {
					rootCauseErr = err
				} else if parentErr := errors.Unwrap(err); parentErr != nil {
					rootCauseErr = parentErr
				}
				Logger.Warningf("skipping job: %d. %s", jobID, rootCauseErr)
				results <- err

				// Note: stop adding jobs if total errors are exceeded
				// This is necessary as the worker still needs time to process
				// the current job, so there can be a delay before the results are read.
				if jobInputErrors >= int64(batchOptions.AbortOnErrorCount) {
					break
				}

				// move to next job
				continue
			}
			Logger.Infof("adding job: %d", jobID)

			if jobID == 1 {
				if commonOptions.Filters != nil && len(commonOptions.Filters.Pluck) > 0 {
					// fmt.Printf("%s\n", strings.Join(commonOptions.Filters.Pluck, ","))
				}
			}

			jobs <- batchArgument{
				id:            jobID,
				batchOptions:  batchOptions,
				request:       *request,
				commonOptions: commonOptions,
			}
		}

		Logger.Info("finished adding jobs")
	}()

	// collect all the results of the work.
	totalErrors := make([]error, 0)

	// close the results when the works are finished, but don't block reading the results
	wasCancelled := int32(0)
	go func() {
		workers.Wait()
		time.Sleep(200 * time.Microsecond)

		// prevent closing channel twice
		if atomic.AddInt32(&wasCancelled, 1) == 1 {
			close(results)
		}
	}()

	for err := range results {
		Logger.Infof("reading job result: %s", err)
		if err != nil && err != io.EOF {
			totalErrors = append(totalErrors, err)
		}
		// exit early
		if batchOptions.AbortOnErrorCount != 0 && len(totalErrors) >= batchOptions.AbortOnErrorCount {
			if atomic.AddInt32(&wasCancelled, 1) == 1 {
				close(results)
			}
			return cmderrors.NewUserErrorWithExitCode(103, fmt.Sprintf("aborted batch as error count has been exceeded. totalErrors=%d", batchOptions.AbortOnErrorCount))
		}
	}
	if progbar.IsEnabled() && jobID > 1 {
		// wait for progress bar to update last increment
		time.Sleep(progbar.RefreshRate())
	}

	maxJobsReached := jobID > globalFlagBatchMaxJobs
	if total := len(totalErrors); total > 0 {
		if total == 1 {
			// return only error
			return totalErrors[0]
		}
		// aggregate error
		return cmderrors.NewUserErrorWithExitCode(104, fmt.Sprintf("jobs completed with %d errors. job limit exceeded=%v", total, maxJobsReached))
	}
	if maxJobsReached {
		return cmderrors.NewUserErrorWithExitCode(105, fmt.Sprintf("max job limit exceeded. limit=%d", globalFlagBatchMaxJobs))
	}
	return nil
}

// These workers will receive work on the `jobs` channel and send the corresponding
// results on `results`
func batchWorker(id int, jobs <-chan batchArgument, results chan<- error, prog *progressbar.ProgressBar, wg *sync.WaitGroup) {
	var err error
	onStartup := true

	var total int64

	defer wg.Done()
	for job := range jobs {
		total++
		workerStart := prog.StartJob(id, total)
		if !onStartup {
			if !errors.Is(err, io.EOF) && job.batchOptions.Delay > 0 {
				Logger.Infof("worker %d: sleeping %dms before fetching next job", id, job.batchOptions.Delay)
				time.Sleep(time.Duration(job.batchOptions.Delay) * time.Millisecond)
			}
		} else {
			jitter := rand.Int31n(50)
			time.Sleep(time.Duration(jitter) * time.Millisecond)
		}
		onStartup = false

		Logger.Infof("worker %d: started job %d", id, job.id)
		startTime := time.Now().UnixNano()

		err = processRequestAndResponse([]c8y.RequestOptions{job.request}, job.commonOptions)
		elapsedMS := (time.Now().UnixNano() - startTime) / 1000.0 / 1000.0

		Logger.Infof("worker %d: finished job %d in %dms", id, job.id, elapsedMS)
		prog.FinishedJob(id, workerStart)

		// return result before delay, so errors can be handled before the sleep
		results <- err
	}
	prog.WorkerCompleted(id)
}
