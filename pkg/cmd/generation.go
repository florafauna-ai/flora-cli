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
	Usage:   "Starts a model generation using type, prompt, workspace_id, project_id, optional\nmodel endpoint ID, and optional model parameters. Use\ntype=image|video|audio|text and model IDs returned by GET /models or\nlist_models. Poll the returned run_id via GET /runs/{runId} for progress and\noutputs. Mutating public API requests support an optional Idempotency-Key header\nfor client retries; duplicate keys within two hours return\nidempotency_duplicate.",
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
	Action:          handleGenerationsCreate,
	HideHelpCommand: true,
}

var generationsList = cli.Command{
	Name:    "list",
	Usage:   "Lists generation history for the authenticated caller, including pending,\nrunning, completed, and failed generations. Results are newest first and can be\nfiltered by workspace_id, project_id, and status. Each item includes poll_url;\nuse it to poll pending/running generations and to fetch completed or failed run\ndetails and outputs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Opaque cursor for fetching the next page",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of results to return",
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Run status filter",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Usage:     "Workspace identifier",
			QueryPath: "workspace_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleGenerationsList,
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

func handleGenerationsList(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := flora.GenerationListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Generations.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "generations list",
			Transform:      transform,
		})
	} else {
		iter := client.Generations.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "generations list",
			Transform:      transform,
		})
	}
}
