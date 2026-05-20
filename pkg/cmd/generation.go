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

var generationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Starts a model generation using a prompt, workspace, project, optional model,\nand optional model parameters. Poll the returned run_id via GET /runs/{runId}\nfor progress and outputs. Mutating public API requests support an optional\nIdempotency-Key header for client retries; duplicate keys within two hours\nreturn idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project identifier",
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
			Usage:    "Generation type",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier",
			Required: true,
			BodyPath: "workspace_id",
		},
		&requestflag.Flag[string]{
			Name:     "model",
			Usage:    "Model endpoint ID",
			BodyPath: "model",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "params",
			Usage:    "Model parameters",
			BodyPath: "params",
		},
	},
	Action:          handleGenerationsCreate,
	HideHelpCommand: true,
}

func handleGenerationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.GenerationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Generations.New(ctx, params, options...)
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
		Title:          "generations create",
		Transform:      transform,
	})
}
