// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/stainless-sdks/florafauna-ai-cli/internal/mocktest"
)

func TestAssetsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "create",
			"--source", "signed-url",
			"--workspace-id", "ws_abc123",
			"--content-type", "image/png",
			"--file-name", "hero.png",
			"--folder", "campaign-assets",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"source: signed-url\n" +
			"workspace_id: ws_abc123\n" +
			"content_type: image/png\n" +
			"file_name: hero.png\n" +
			"folder: campaign-assets\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"assets", "create",
		)
	})
}

func TestAssetsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "retrieve",
			"--asset-id", "asset_abc123",
		)
	})
}

func TestAssetsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "list",
			"--cursor", "cursor",
			"--limit", "1",
			"--project-id", "prj_abc123",
			"--query", "logo",
			"--workspace-id", "ws_abc123",
		)
	})
}

func TestAssetsCompleteUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "complete-upload",
			"--asset-id", "asset_abc123",
		)
	})
}

func TestAssetsRetryUpload(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"assets", "retry-upload",
			"--asset-id", "asset_abc123",
		)
	})
}
