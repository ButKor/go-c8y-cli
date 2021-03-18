// Code generated from specification version 1.0.0: DO NOT EDIT
package assigngrouptogroup

import (
	"io"
	"net/http"

	"github.com/MakeNowJust/heredoc/v2"
	"github.com/reubenmiller/go-c8y-cli/pkg/c8yfetcher"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmd/subcommand"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y-cli/pkg/cmdutil"
	"github.com/reubenmiller/go-c8y-cli/pkg/completion"
	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

// AssignGroupToGroupCmd command
type AssignGroupToGroupCmd struct {
	*subcommand.SubCommand

	factory *cmdutil.Factory
}

// NewAssignGroupToGroupCmd creates a command to Assign child group
func NewAssignGroupToGroupCmd(f *cmdutil.Factory) *AssignGroupToGroupCmd {
	ccmd := &AssignGroupToGroupCmd{
		factory: f,
	}
	cmd := &cobra.Command{
		Use:   "assignGroupToGroup",
		Short: "Assign child group",
		Long:  `Assigns a group to a group. The group will be a childAsset of the group`,
		Example: heredoc.Doc(`
$ c8y inventoryReferences assignGroupToGroup --group 12345 --newChildGroup 43234
Add a group to a group

$ c8y inventoryReferences assignGroupToGroup --group 12345 --newChildGroup 43234, 99292, 12222
Add multiple groups to a group
        `),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return f.CreateModeEnabled()
		},
		RunE: ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().StringSlice("group", []string{""}, "Group (required)")
	cmd.Flags().StringSlice("newChildGroup", []string{""}, "New child group to be added to the group as an child asset (required) (accepts pipeline)")

	completion.WithOptions(
		cmd,
	)

	flags.WithOptions(
		cmd,
		flags.WithProcessingMode(),

		flags.WithExtendedPipelineSupport("newChildGroup", "managedObject.id", true, "id"),
		flags.WithCollectionProperty("managedObject"),
	)

	// Required flags
	_ = cmd.MarkFlagRequired("group")

	ccmd.SubCommand = subcommand.NewSubCommand(cmd)

	return ccmd
}

// RunE executes the command
func (n *AssignGroupToGroupCmd) RunE(cmd *cobra.Command, args []string) error {
	cfg, err := n.factory.Config()
	if err != nil {
		return err
	}
	client, err := n.factory.Client()
	if err != nil {
		return err
	}
	inputIterators, err := flags.NewRequestInputIterators(cmd)
	if err != nil {
		return err
	}

	// query parameters
	query := flags.NewQueryTemplate()
	err = flags.WithQueryParameters(
		cmd,
		query,
		inputIterators,
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	queryValue, err := query.GetQueryUnescape(true)

	if err != nil {
		return cmderrors.NewSystemError("Invalid query parameter")
	}

	// headers
	headers := http.Header{}
	err = flags.WithHeaders(
		cmd,
		headers,
		inputIterators,
		flags.WithProcessingModeValue(),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// form data
	formData := make(map[string]io.Reader)
	err = flags.WithFormDataOptions(
		cmd,
		formData,
		inputIterators,
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// body
	body := mapbuilder.NewInitializedMapBuilder()
	err = flags.WithBody(
		cmd,
		body,
		inputIterators,
		flags.WithDataFlagValue(),
		cmdutil.WithTemplateValue(cfg),
		flags.WithTemplateVariablesValue(),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// path parameters
	path := flags.NewStringTemplate("inventory/managedObjects/{id}/childAssets")
	err = flags.WithPathParameters(
		cmd,
		path,
		inputIterators,
		c8yfetcher.WithDeviceGroupByNameFirstMatch(client, args, "group", "id"),
	)
	if err != nil {
		return err
	}

	req := c8y.RequestOptions{
		Method:       "POST",
		Path:         path.GetTemplate(),
		Query:        queryValue,
		Body:         body,
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: cfg.IgnoreAcceptHeader(),
		DryRun:       cfg.DryRun(),
	}

	return n.factory.RunWithWorkers(client, cmd, &req, inputIterators)
}
