// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/florafauna-ai-cli/internal/mocktest"
)

func TestFeedbackRecord(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"feedback", "record",
			"--detail", "I want to export all generated campaign images at once.",
			"--kind", "feature_request",
			"--summary", "Need batch export support",
			"--attempted-tool", "generate_image",
			"--project-id", "prj_abc123",
			"--run-id", "run_abc123",
			"--workspace-id", "ws_abc123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"detail: I want to export all generated campaign images at once.\n" +
			"kind: feature_request\n" +
			"summary: Need batch export support\n" +
			"attempted_tools:\n" +
			"  - generate_image\n" +
			"project_id: prj_abc123\n" +
			"run_id: run_abc123\n" +
			"workspace_id: ws_abc123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"feedback", "record",
		)
	})
}
