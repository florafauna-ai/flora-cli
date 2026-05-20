// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/florafauna-ai/flora-cli/internal/mocktest"
)

func TestProjectsActionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects:actions", "create",
			"--project-id", "prj_abc123",
			"--action-id", "split-text",
			"--params", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"action_id: split-text\n" +
			"params:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"projects:actions", "create",
			"--project-id", "prj_abc123",
		)
	})
}

func TestProjectsActionsRun(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects:actions", "run",
			"--project-id", "prj_abc123",
			"--node-id", "nodeId",
		)
	})
}
