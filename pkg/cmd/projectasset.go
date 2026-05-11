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

var projectsAssetsAttach = cli.Command{
	Name:    "attach",
	Usage:   "Attaches an existing ready asset to a project canvas as a static media node.\nMutating public API requests support an optional Idempotency-Key header for\nclient retries; duplicate keys within two hours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project identifier",
			Required:  true,
			PathParam: "projectId",
		},
		&requestflag.Flag[string]{
			Name:      "asset-id",
			Usage:     "Asset identifier",
			Required:  true,
			PathParam: "assetId",
		},
	},
	Action:          handleProjectsAssetsAttach,
	HideHelpCommand: true,
}

func handleProjectsAssetsAttach(ctx context.Context, cmd *cli.Command) error {
	client := florafaunaai.NewClient(getDefaultRequestOptions(cmd)...)
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

	params := florafaunaai.ProjectAssetAttachParams{
		ProjectID: cmd.Value("project-id").(string),
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Projects.Assets.Attach(
		ctx,
		cmd.Value("asset-id").(string),
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
		Title:          "projects:assets attach",
		Transform:      transform,
	})
}
