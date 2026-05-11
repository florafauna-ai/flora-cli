// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/florafauna-ai-cli/internal/mocktest"
)

func TestTechniquesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"techniques", "retrieve",
			"--technique-id", "tech_def_abc123",
		)
	})
}

func TestTechniquesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"techniques", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "1",
			"--query", "logo",
			"--workspace-id", "ws_abc123",
		)
	})
}
