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

var projectsCreate = cli.Command{
	Name:    "create",
	Usage:   "Creates a new Flora project in the requested workspace. Mutating public API\nrequests support an optional Idempotency-Key header for client retries;\nduplicate keys within two hours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Project name",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier",
			Required: true,
			BodyPath: "workspace_id",
		},
	},
	Action:          handleProjectsCreate,
	HideHelpCommand: true,
}

var projectsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Returns metadata for a single project when it is accessible to the authenticated\npublic API key. Missing and inaccessible projects both return 404.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
	},
	Action:          handleProjectsRetrieve,
	HideHelpCommand: true,
}

var projectsList = cli.Command{
	Name:    "list",
	Usage:   "Returns projects in the requested workspace that are accessible to the\nauthenticated public API key, ordered by recent activity.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "workspace-id",
			Usage:     "Workspace identifier",
			Required:  true,
			QueryPath: "workspace_id",
		},
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
			Name:      "query",
			Usage:     "Search query",
			QueryPath: "query",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleProjectsList,
	HideHelpCommand: true,
}

var projectsListNodes = cli.Command{
	Name:    "list-nodes",
	Usage:   "Returns sanitized visible media nodes on a project canvas. The response omits\nraw graph documents, Liveblocks internals, raw Convex IDs, and unbounded node\ndata blobs.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleProjectsListNodes,
	HideHelpCommand: true,
}

func handleProjectsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.ProjectNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Projects.New(ctx, params, options...)
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
		Title:          "projects create",
		Transform:      transform,
	})
}

func handleProjectsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Projects.Get(ctx, cmd.Value("project-id").(string), options...)
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
		Title:          "projects retrieve",
		Transform:      transform,
	})
}

func handleProjectsList(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.ProjectListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Projects.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "projects list",
			Transform:      transform,
		})
	} else {
		iter := client.Projects.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "projects list",
			Transform:      transform,
		})
	}
}

func handleProjectsListNodes(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.ProjectListNodesParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Projects.ListNodes(
			ctx,
			cmd.Value("project-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "projects list-nodes",
			Transform:      transform,
		})
	} else {
		iter := client.Projects.ListNodesAutoPaging(
			ctx,
			cmd.Value("project-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "projects list-nodes",
			Transform:      transform,
		})
	}
}
