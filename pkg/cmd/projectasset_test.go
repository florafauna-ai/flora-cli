// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/florafauna-ai-cli/internal/mocktest"
)

func TestProjectsAssetsAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects:assets", "attach",
			"--project-id", "prj_abc123",
			"--asset-id", "asset_abc123",
		)
	})
}
