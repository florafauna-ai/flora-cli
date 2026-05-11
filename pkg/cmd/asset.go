// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/stainless-sdks/florafauna-ai-cli/internal/apiquery"
	"github.com/stainless-sdks/florafauna-ai-cli/internal/requestflag"
	"github.com/stainless-sdks/florafauna-ai-go"
	"github.com/stainless-sdks/florafauna-ai-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var assetsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates an asset from an allowlisted source URL or reserves a signed upload URL.\nMutating public API requests support an optional Idempotency-Key header for\nclient retries; duplicate keys within two hours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "source",
			Usage:    "Asset source URL or signed-url upload mode",
			Required: true,
			BodyPath: "source",
		},
		&requestflag.Flag[string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier",
			Required: true,
			BodyPath: "workspace_id",
		},
		&requestflag.Flag[string]{
			Name:     "content-type",
			Usage:    "Asset content type",
			BodyPath: "content_type",
		},
		&requestflag.Flag[string]{
			Name:     "file-name",
			Usage:    "Asset file name",
			BodyPath: "file_name",
		},
		&requestflag.Flag[string]{
			Name:     "folder",
			Usage:    "Destination folder",
			BodyPath: "folder",
		},
	},
	Action:          handleAssetsCreate,
	HideHelpCommand: true,
}

var assetsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns metadata for one asset when it is accessible to the authenticated public\nAPI key. Missing and inaccessible assets both return 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "asset-id",
			Usage:     "Asset identifier",
			Required:  true,
			PathParam: "assetId",
		},
	},
	Action:          handleAssetsRetrieve,
	HideHelpCommand: true,
}

var assetsList = cli.Command{
	Name:    "list",
	Usage:   "Returns assets visible to the authenticated public API key. Filter by workspace,\nproject canvas, search query, cursor, and limit without exposing raw file bytes\nor internal graph data.",
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
			Name:      "query",
			Usage:     "Search query",
			QueryPath: "query",
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
	Action:          handleAssetsList,
	HideHelpCommand: true,
}

var assetsComplete = cli.Command{
	Name:    "complete",
	Usage:   "Marks a signed asset upload as complete after the file has been uploaded.\nMutating public API requests support an optional Idempotency-Key header for\nclient retries; duplicate keys within two hours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "asset-id",
			Usage:     "Asset identifier",
			Required:  true,
			PathParam: "assetId",
		},
	},
	Action:          handleAssetsComplete,
	HideHelpCommand: true,
}

var assetsRetry = cli.Command{
	Name:    "retry",
	Usage:   "Creates a fresh signed upload reservation for a failed or expired asset upload.\nMutating public API requests support an optional Idempotency-Key header for\nclient retries; duplicate keys within two hours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "asset-id",
			Usage:     "Asset identifier",
			Required:  true,
			PathParam: "assetId",
		},
	},
	Action:          handleAssetsRetry,
	HideHelpCommand: true,
}

func handleAssetsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.AssetNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Assets.New(ctx, params, options...)
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
		Title:          "assets create",
		Transform:      transform,
	})
}

func handleAssetsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("asset-id") && len(unusedArgs) > 0 {
		cmd.Set("asset-id", unusedArgs[0])
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
	_, err = client.Assets.Get(ctx, cmd.Value("asset-id").(string), options...)
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
		Title:          "assets retrieve",
		Transform:      transform,
	})
}

func handleAssetsList(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.AssetListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Assets.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "assets list",
			Transform:      transform,
		})
	} else {
		iter := client.Assets.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "assets list",
			Transform:      transform,
		})
	}
}

func handleAssetsComplete(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("asset-id") && len(unusedArgs) > 0 {
		cmd.Set("asset-id", unusedArgs[0])
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
	_, err = client.Assets.Complete(ctx, cmd.Value("asset-id").(string), options...)
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
		Title:          "assets complete",
		Transform:      transform,
	})
}

func handleAssetsRetry(ctx context.Context, cmd *cli.Command) error {
	client := flora.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("asset-id") && len(unusedArgs) > 0 {
		cmd.Set("asset-id", unusedArgs[0])
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
	_, err = client.Assets.Retry(ctx, cmd.Value("asset-id").(string), options...)
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
		Title:          "assets retry",
		Transform:      transform,
	})
}
