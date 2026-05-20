// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/florafauna-ai/flora-cli/internal/apiquery"
	"github.com/florafauna-ai/flora-go"
	"github.com/florafauna-ai/flora-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var workspacesList = cli.Command{
	Name:            "list",
	Usage:           "Returns the workspaces available to the authenticated public API key, including\neach workspace's public ID, name, creation timestamp, and caller role.",
	Suggest:         true,
	Flags:           []cli.Flag{},
	Action:          handleWorkspacesList,
	HideHelpCommand: true,
}

func handleWorkspacesList(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Workspaces.List(ctx, options...)
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
		Title:          "workspaces list",
		Transform:      transform,
	})
}
