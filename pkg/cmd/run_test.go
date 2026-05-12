// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/florafauna-ai/flora-cli/internal/mocktest"
)

func TestRunsStartGeneration(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"runs", "start-generation",
			"--project-id", "prj_abc123",
			"--prompt", "A cinematic product photo of a ceramic mug on a sunlit table",
			"--type", "image",
			"--workspace-id", "ws_abc123",
			"--model", "t2i-flux-2-pro",
			"--params", "{foo: bar}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"project_id: prj_abc123\n" +
			"prompt: A cinematic product photo of a ceramic mug on a sunlit table\n" +
			"type: image\n" +
			"workspace_id: ws_abc123\n" +
			"model: t2i-flux-2-pro\n" +
			"params:\n" +
			"  foo: bar\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"runs", "start-generation",
		)
	})
}

func TestRunsStartTechnique(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"runs", "start-technique",
			"--inputs", "{foo: bar}",
			"--technique-id", "tech_abcd1234",
			"--workspace-id", "ws_abc123",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inputs:\n" +
			"  foo: bar\n" +
			"technique_id: tech_abcd1234\n" +
			"workspace_id: ws_abc123\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"runs", "start-technique",
		)
	})
}
