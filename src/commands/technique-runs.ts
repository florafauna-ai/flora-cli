import { defineCommand } from "citty";
import type { RunCreateParams } from "@flora-ai/flora/resources/techniques/runs";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const create = defineCommand({
  meta: { name: "create", description: "Start a technique run" },
  args: {
    "technique-id": {
      type: "positional",
      description: "Technique identifier",
      required: true,
    },
    inputs: {
      type: "string",
      description: 'Run inputs as JSON array, e.g. \'[{"id":"in","type":"text","value":"hello"}]\'',
      required: true,
    },
    mode: {
      type: "string",
      description: "Run mode: async | stream",
      required: true,
    },
    "callback-url": {
      type: "string",
      description: "Optional callback URL for async notification",
    },
    "idempotency-key": {
      type: "string",
      description: "Idempotency key for retry safety",
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    let inputs: RunCreateParams["inputs"];
    try {
      inputs = JSON.parse(args.inputs) as RunCreateParams["inputs"];
    } catch {
      console.error(`Invalid JSON for inputs: ${args.inputs}`);
      process.exit(1);
    }
    return client.techniques.runs
      .create(args["technique-id"], {
        inputs,
        mode: args.mode as RunCreateParams["mode"],
        ...(args["callback-url"] ? { callback_url: args["callback-url"] } : {}),
        ...(args["idempotency-key"] ? { idempotency_key: args["idempotency-key"] } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const retrieve = defineCommand({
  meta: { name: "retrieve", description: "Get a technique run by ID" },
  args: {
    "run-id": {
      type: "positional",
      description: "Run identifier",
      required: true,
    },
    "technique-id": {
      type: "string",
      description: "Technique identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.techniques.runs
      .retrieve(args["run-id"], { techniqueId: args["technique-id"] })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "technique-runs", description: "Technique run management" },
  subCommands: { create, retrieve },
});
