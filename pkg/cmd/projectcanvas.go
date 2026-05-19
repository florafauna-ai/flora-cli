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

var projectsCanvasRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns the current project canvas topology as a Mermaid flowchart using the\nsame serializer as the Fauna agent.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
	},
	Action:          handleProjectsCanvasRetrieve,
	HideHelpCommand: true,
}

var projectsCanvasUpdate = cli.Command{
	Name:    "update",
	Usage:   "Applies a Mermaid flowchart patch to the project canvas using the same\ncreate_workflow path as the Fauna agent. The diagram may add nodes, connect\nnodes, and reference existing canvas nodes by their Mermaid short IDs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
		&requestflag.Flag[string]{
			Name:     "diagram",
			Usage:    "Mermaid flowchart diagram to apply",
			Required: true,
			BodyPath: "diagram",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "node-params",
			Usage:    "Optional per-node parameters keyed by Mermaid node ID.",
			BodyPath: "node_params",
		},
	},
	Action:          handleProjectsCanvasUpdate,
	HideHelpCommand: true,
}

func handleProjectsCanvasRetrieve(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Projects.Canvas.Get(ctx, cmd.Value("project-id").(string), options...)
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
		Title:          "projects:canvas retrieve",
		Transform:      transform,
	})
}

func handleProjectsCanvasUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.ProjectCanvasUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Projects.Canvas.Update(
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
		Title:          "projects:canvas update",
		Transform:      transform,
	})
}
