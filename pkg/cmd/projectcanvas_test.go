// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/florafauna-ai/flora-cli/internal/mocktest"
)

func TestProjectsCanvasRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects:canvas", "retrieve",
			"--project-id", "prj_abc123",
		)
	})
}

func TestProjectsCanvasUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"projects:canvas", "update",
			"--project-id", "prj_abc123",
			"--diagram", "graph LR\n  source[\"Product photo (Image)\"]\n  output[\"Editorial campaign image (Image)\"]\n  source --> output",
			"--node-params", "{foo: {aspect_ratio: aspect_ratio, content_url: https://example.com, model: model, model_parameters: {foo: bar}, prompt: prompt, resolution: resolution}}",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"diagram: |-\n" +
			"  graph LR\n" +
			"    source[\"Product photo (Image)\"]\n" +
			"    output[\"Editorial campaign image (Image)\"]\n" +
			"    source --> output\n" +
			"node_params:\n" +
			"  foo:\n" +
			"    aspect_ratio: aspect_ratio\n" +
			"    content_url: https://example.com\n" +
			"    model: model\n" +
			"    model_parameters:\n" +
			"      foo: bar\n" +
			"    prompt: prompt\n" +
			"    resolution: resolution\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"projects:canvas", "update",
			"--project-id", "prj_abc123",
		)
	})
}
