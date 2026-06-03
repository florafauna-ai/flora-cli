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

var techniquesRunsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Starts a run for a specific technique using the backward-compatible nested\nroute. Mutating public API requests support an optional Idempotency-Key header\nfor client retries; duplicate keys within two hours return\nidempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "technique-id",
			Usage:     "Technique identifier or slug",
			Required:  true,
			PathParam: "techniqueId",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "input",
			Usage:    "Technique inputs",
			Required: true,
			BodyPath: "inputs",
		},
		&requestflag.Flag[string]{
			Name:     "mode",
			Usage:    "Technique run execution mode",
			Required: true,
			BodyPath: "mode",
		},
		&requestflag.Flag[string]{
			Name:     "callback-url",
			Usage:    "HTTPS callback URL for asynchronous run completion notifications",
			BodyPath: "callback_url",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Idempotency key for safely retrying requests",
			BodyPath: "idempotency_key",
		},
	},
	Action:          handleTechniquesRunsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"input": {
		&requestflag.InnerFlag[string]{
			Name:       "input.id",
			Usage:      "Technique input identifier",
			InnerField: "id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.type",
			Usage:      "Technique input type",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "input.value",
			Usage:      "Technique input value",
			InnerField: "value",
		},
	},
})

var techniquesRunsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns status, progress, outputs, and error details for a technique run when it\nis accessible to the authenticated public API key.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "technique-id",
			Usage:     "Technique identifier or slug",
			Required:  true,
			PathParam: "techniqueId",
		},
		&requestflag.Flag[string]{
			Name:      "run-id",
			Usage:     "Run identifier",
			Required:  true,
			PathParam: "runId",
		},
	},
	Action:          handleTechniquesRunsRetrieve,
	HideHelpCommand: true,
}

var techniquesRunsList = cli.Command{
	Name:    "list",
	Usage:   "Lists technique run history for the authenticated caller, including pending,\nrunning, completed, and failed technique runs. Results are newest first and can\nbe filtered by workspace_id, project_id, technique_id, and status. Each item\nincludes poll_url; use it to poll pending/running technique runs and to fetch\ncompleted or failed run details and outputs.",
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
			Name:      "technique-id",
			Usage:     "Technique identifier",
			QueryPath: "technique_id",
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
	Action:          handleTechniquesRunsList,
	HideHelpCommand: true,
}

func handleTechniquesRunsCreate(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("technique-id") && len(unusedArgs) > 0 {
		cmd.Set("technique-id", unusedArgs[0])
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

	params := flora.TechniqueRunNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Techniques.Runs.New(
		ctx,
		cmd.Value("technique-id").(string),
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
		Title:          "techniques:runs create",
		Transform:      transform,
	})
}

func handleTechniquesRunsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("run-id") && len(unusedArgs) > 0 {
		cmd.Set("run-id", unusedArgs[0])
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

	params := flora.TechniqueRunGetParams{
		TechniqueID: cmd.Value("technique-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Techniques.Runs.Get(
		ctx,
		cmd.Value("run-id").(string),
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
		Title:          "techniques:runs retrieve",
		Transform:      transform,
	})
}

func handleTechniquesRunsList(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.TechniqueRunListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Techniques.Runs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "techniques:runs list",
			Transform:      transform,
		})
	} else {
		iter := client.Techniques.Runs.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "techniques:runs list",
			Transform:      transform,
		})
	}
}
