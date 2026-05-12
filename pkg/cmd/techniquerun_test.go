// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/florafauna-ai/flora-cli/internal/mocktest"
	"github.com/florafauna-ai/flora-cli/internal/requestflag"
)

func TestTechniquesRunsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"techniques:runs", "create",
			"--technique-id", "tech_def_abc123",
			"--input", "{id: id, type: imageUrl, value: value}",
			"--mode", "async",
			"--callback-url", "https://example.com",
			"--idempotency-key", "idempotency_key",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(techniquesRunsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"techniques:runs", "create",
			"--technique-id", "tech_def_abc123",
			"--input.id", "id",
			"--input.type", "imageUrl",
			"--input.value", "value",
			"--mode", "async",
			"--callback-url", "https://example.com",
			"--idempotency-key", "idempotency_key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"inputs:\n" +
			"  - id: id\n" +
			"    type: imageUrl\n" +
			"    value: value\n" +
			"mode: async\n" +
			"callback_url: https://example.com\n" +
			"idempotency_key: idempotency_key\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"techniques:runs", "create",
			"--technique-id", "tech_def_abc123",
		)
	})
}

func TestTechniquesRunsRetrieve(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"techniques:runs", "retrieve",
			"--technique-id", "tech_def_abc123",
			"--run-id", "run_abc123",
		)
	})
}
