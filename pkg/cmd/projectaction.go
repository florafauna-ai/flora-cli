// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/florafauna-ai/flora-cli/internal/apiquery"
	"github.com/florafauna-ai/flora-cli/internal/requestflag"
	"github.com/florafauna-ai/flora-go"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var projectsActionsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a prebuilt action node on a project canvas using a raw action slug.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
		&requestflag.Flag[string]{
			Name:     "action-id",
			Usage:    "Action identifier",
			Required: true,
			BodyPath: "action_id",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Usage:    "Action parameters",
			BodyPath: "params",
		},
	},
	Action:          handleProjectsActionsCreate,
	HideHelpCommand: true,
}

var projectsActionsRun = cli.Command{
	Name:    "run",
	Usage:   "Runs an existing canvas action node through the action execution workflow.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
		&requestflag.Flag[string]{
			Name:      "node-id",
			Usage:     "Canvas action node identifier",
			Required:  true,
			PathParam: "nodeId",
		},
	},
	Action:          handleProjectsActionsRun,
	HideHelpCommand: true,
}

func handleProjectsActionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("project-id") && len(unusedArgs) > 0 {
		cmd.Set("project-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := flora.ProjectActionNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Projects.Actions.New(
		ctx,
		cmd.Value("project-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "projects:actions create",
		Transform:      transform,
	})
}

func handleProjectsActionsRun(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("node-id") && len(unusedArgs) > 0 {
		cmd.Set("node-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	params := flora.ProjectActionRunParams{
		ProjectID: cmd.Value("project-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Projects.Actions.Run(
		ctx,
		cmd.Value("node-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "projects:actions run",
		Transform:      transform,
	})
}
