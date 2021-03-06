package cmd

import (
	"context"
	"regexp"

	"github.com/pkg/errors"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
)

type applicationFetcher struct {
	client *c8y.Client
}

func newApplicationFetcher(client *c8y.Client) *applicationFetcher {
	return &applicationFetcher{
		client: client,
	}
}

func (f *applicationFetcher) getByID(id string) ([]fetcherResultSet, error) {
	app, resp, err := client.Application.GetApplication(
		context.Background(),
		id,
	)

	if err != nil {
		return nil, errors.Wrap(err, "Could not fetch by id")
	}

	results := make([]fetcherResultSet, 1)
	results[0] = fetcherResultSet{
		ID:    app.ID,
		Name:  app.Name,
		Value: *resp.JSON,
	}
	return results, nil
}

// getByName returns applications matching a given using regular expression
func (f *applicationFetcher) getByName(name string) ([]fetcherResultSet, error) {
	col, _, err := client.Application.GetApplications(
		context.Background(),
		&c8y.ApplicationOptions{
			PaginationOptions: *c8y.NewPaginationOptions(2000),
		},
	)
	if err != nil {
		return nil, err
	}

	pattern, err := regexp.Compile(name)

	if err != nil {
		return nil, errors.Wrap(err, "invalid regex")
	}

	if err != nil {
		return nil, errors.Wrap(err, "Could not fetch by id")
	}

	results := make([]fetcherResultSet, 0)

	for i, app := range col.Applications {
		if pattern.MatchString(app.Name) {
			results = append(results, fetcherResultSet{
				ID:    app.ID,
				Name:  app.Name,
				Value: col.Items[i],
			})
		}

	}

	return results, nil
}
