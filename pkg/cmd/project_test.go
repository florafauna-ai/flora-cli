// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/florafauna-ai-cli/internal/mocktest"
)

func TestProjectsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects", "create",
			"--name", "Spring Campaign",
			"--workspace-id", "ws_abc123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: Spring Campaign\n" +
			"workspace_id: ws_abc123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"projects", "create",
		)
	})
}

func TestProjectsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects", "retrieve",
			"--project-id", "prj_abc123",
		)
	})
}

func TestProjectsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects", "list",
			"--workspace-id", "ws_abc123",
			"--cursor", "cursor",
			"--limit", "1",
			"--query", "logo",
		)
	})
}

func TestProjectsListNodes(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects", "list-nodes",
			"--project-id", "prj_abc123",
			"--cursor", "cursor",
			"--limit", "1",
		)
	})
}
