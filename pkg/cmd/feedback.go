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

var feedbackRecord = cli.Command{
	Name:    "record",
	Usage:   "Records product feedback from the authenticated user, optionally linked to a\nworkspace, project, run, and attempted tools. Mutating public API requests\nsupport an optional Idempotency-Key header for client retries; duplicate keys\nwithin two hours return idempotency_duplicate.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "detail",
			Usage:    "Detailed description",
			Required: true,
			BodyPath: "detail",
		},
		&requestflag.Flag[string]{
			Name:     "kind",
			Usage:    "Feedback kind",
			Required: true,
			BodyPath: "kind",
		},
		&requestflag.Flag[string]{
			Name:     "summary",
			Usage:    "Short summary",
			Required: true,
			BodyPath: "summary",
		},
		&requestflag.Flag[[]string]{
			Name:     "attempted-tool",
			Usage:    "Tools or capabilities attempted before submitting feedback",
			BodyPath: "attempted_tools",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project identifier",
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "run-id",
			Usage:    "Run identifier",
			BodyPath: "run_id",
		},
		&requestflag.Flag[string]{
			Name:     "workspace-id",
			Usage:    "Workspace identifier",
			BodyPath: "workspace_id",
		},
	},
	Action:          handleFeedbackRecord,
	HideHelpCommand: true,
}

func handleFeedbackRecord(ctx context.Context, cmd *cli.Command) error {
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

	params := flora.FeedbackRecordParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Feedback.Record(ctx, params, options...)
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
		Title:          "feedback record",
		Transform:      transform,
	})
}
