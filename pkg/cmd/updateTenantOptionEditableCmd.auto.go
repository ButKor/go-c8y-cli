// Code generated from specification version 1.0.0: DO NOT EDIT
package cmd

import (
	"io"
	"net/http"

	"github.com/reubenmiller/go-c8y-cli/pkg/cmderrors"
	"github.com/reubenmiller/go-c8y-cli/pkg/completion"
	"github.com/reubenmiller/go-c8y-cli/pkg/flags"
	"github.com/reubenmiller/go-c8y-cli/pkg/mapbuilder"
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/spf13/cobra"
)

// UpdateTenantOptionEditableCmd command
type UpdateTenantOptionEditableCmd struct {
	*baseCmd
}

// NewUpdateTenantOptionEditableCmd creates a command to Update tenant option edit setting
func NewUpdateTenantOptionEditableCmd() *UpdateTenantOptionEditableCmd {
	ccmd := &UpdateTenantOptionEditableCmd{}
	cmd := &cobra.Command{
		Use:   "updateEdit",
		Short: "Update tenant option edit setting",
		Long: `Update read-only setting of an existing tenant option Required role:: ROLE_OPTION_MANAGEMENT_ADMIN, Required tenant management Example Request:: Update access.control.allow.origin option.
`,
		Example: `
$ c8y tenantOptions updateEdit --category "c8y_cli_tests" --key "option8" --editable "true"
Update editable property for an existing tenant option
        `,
		PreRunE: validateUpdateMode,
		RunE:    ccmd.RunE,
	}

	cmd.SilenceUsage = true

	cmd.Flags().String("category", "", "Tenant Option category (required)")
	cmd.Flags().String("key", "", "Tenant Option key (required) (accepts pipeline)")
	cmd.Flags().String("editable", "", "Whether the tenant option should be editable or not (required)")
	addProcessingModeFlag(cmd)

	completion.WithOptions(
		cmd,
		completion.WithValidateSet("editable", "true", "false"),
	)

	flags.WithOptions(
		cmd,
		flags.WithExtendedPipelineSupport("key", "key", true, "id"),
	)

	// Required flags
	cmd.MarkFlagRequired("category")
	cmd.MarkFlagRequired("editable")

	ccmd.baseCmd = newBaseCmd(cmd)

	return ccmd
}

// RunE executes the command
func (n *UpdateTenantOptionEditableCmd) RunE(cmd *cobra.Command, args []string) error {
	var err error
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
		WithDataValue(),
		flags.WithStringValue("editable", "editable"),
		WithTemplateValue(),
		WithTemplateVariablesValue(),
	)
	if err != nil {
		return cmderrors.NewUserError(err)
	}

	// path parameters
	path := flags.NewStringTemplate("/tenant/options/{category}/{key}/editable")
	err = flags.WithPathParameters(
		cmd,
		path,
		inputIterators,
		flags.WithStringValue("category", "category"),
		flags.WithStringValue("key", "key"),
	)
	if err != nil {
		return err
	}

	req := c8y.RequestOptions{
		Method:       "PUT",
		Path:         path.GetTemplate(),
		Query:        queryValue,
		Body:         body,
		FormData:     formData,
		Header:       headers,
		IgnoreAccept: globalFlagIgnoreAccept,
		DryRun:       globalFlagDryRun,
	}

	return processRequestAndResponseWithWorkers(cmd, &req, inputIterators)
}
