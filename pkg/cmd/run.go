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

var runsStartGeneration = cli.Command{
	Name:    "start-generation",
	Usage:   "Starts a model generation run in a project canvas using type, prompt,\nworkspace_id, project_id, optional model endpoint ID, and optional model\nparameters. Use type=image|video|audio|text and model IDs returned by GET\n/models or list_models. Mutating public API requests support an optional\nIdempotency-Key header for client retries; duplicate keys within two hours\nreturn idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project identifier. Use the public API ID returned by list projects; it must start with prj_.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "prompt",
			Usage:    "Generation prompt",
			Required: true,
			BodyPath: "prompt",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    `Generation type. Use "image", "video", "audio", or "text"; do not pass model families such as "t2i" or "i2v".`,
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier. Use the public API ID returned by list workspaces; it must start with ws_.",
			Required: true,
			BodyPath: "workspace_id",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model endpoint ID, not a display name. Use list_models (or GET /models) to find accessible endpoint IDs for the requested type.",
			BodyPath: "model",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Usage:    "Model parameters",
			BodyPath: "params",
		},
	},
	Action:          handleRunsStartGeneration,
	HideHelpCommand: true,
}

var runsStartTechnique = cli.Command{
	Name:    "start-technique",
	Usage:   "Starts a technique run through the normalized top-level run resource using\ntechnique*id, workspace_id, and inputs. technique_id must use the tech* public\nAPI ID returned by list techniques. Mutating public API requests support an\noptional Idempotency-Key header for client retries; duplicate keys within two\nhours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "inputs",
			Usage:    "Technique inputs",
			Required: true,
			BodyPath: "inputs",
		},
		&requestflag.Flag[string]{
			Name:     "technique-id",
			Usage:    "Technique identifier. Use the public API ID returned by list techniques; it must start with tech_.",
			Required: true,
			BodyPath: "technique_id",
		},
		&requestflag.Flag[string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier. Use the public API ID returned by list workspaces; it must start with ws_.",
			Required: true,
			BodyPath: "workspace_id",
		},
	},
	Action:          handleRunsStartTechnique,
	HideHelpCommand: true,
}

func handleRunsStartGeneration(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := flora.RunStartGenerationParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Runs.StartGeneration(ctx, params, options...)
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
		Title:          "runs start-generation",
		Transform:      transform,
	})
}

func handleRunsStartTechnique(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := flora.RunStartTechniqueParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Runs.StartTechnique(ctx, params, options...)
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
		Title:          "runs start-technique",
		Transform:      transform,
	})
}
