// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/florafauna-ai/flora-cli/internal/mocktest"
)

func TestTechniquesRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"techniques", "retrieve",
			"--technique-id", "art-directors-critique",
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
			"--cursor", "eyJvZmZzZXQiOjIwfQ",
			"--limit", "1",
			"--query", "logo",
			"--workspace-id", "ws_abc123",
		)
	})
}
