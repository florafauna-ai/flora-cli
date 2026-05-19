import { defineCommand } from "citty";
import type { RunStartGenerationParams } from "@flora-ai/flora/resources/runs";
import {
  getClient,
  handleError,
  parseJsonArg,
  printResult,
  pollRun,
  downloadOutputs,
  pollArgs,
  type GlobalArgs,
} from "../client.js";

const startGeneration = defineCommand({
  meta: {
    name: "start-generation",
    description: "Start a model generation run (deprecated — use 'generations create')",
  },
  args: {
    type: { type: "string", description: "Generation type", required: true },
    prompt: { type: "string", description: "Prompt text", required: true },
    "workspace-id": { type: "string", description: "Workspace identifier", required: true },
    "project-id": { type: "string", description: "Project identifier", required: true },
    model: { type: "string", description: "Model identifier" },
    params: { type: "string", description: "Model parameters as JSON" },
    ...pollArgs,
  },
  async run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    try {
      const res = await client.runs.startGeneration({
        type: args.type as RunStartGenerationParams["type"],
        prompt: args.prompt,
        workspace_id: args["workspace-id"],
        project_id: args["project-id"],
        ...(args.model ? { model: args.model } : {}),
        ...(args.params
          ? { params: parseJsonArg(args.params, "params") as Record<string, unknown> }
          : {}),
      });

      printResult(res, global);

      const shouldPoll = global.poll || global.download;
      if (shouldPoll) {
        const runId = res.run_id;
        const pollUrl = res.poll_url ?? undefined;
        console.error(`\nPolling run ${runId}...`);
        const status = await pollRun(runId, global, pollUrl);
        printResult(status, global);

        if (status.status === "failed") {
          console.error(`Run failed: ${status.error_message ?? status.error_code ?? "unknown"}`);
          process.exit(1);
        }

        if (global.download) {
          await downloadOutputs(status, global);
        }
      }
    } catch (err) {
      handleError(err);
    }
  },
});

const retrieve = defineCommand({
  meta: { name: "retrieve", description: "Get run status and outputs" },
  args: {
    "run-id": {
      type: "positional",
      description: "Run identifier",
      required: true,
    },
  },
  async run({ args }) {
    const global = args as unknown as GlobalArgs;
    const baseURL = global["base-url"] ?? process.env["FLORA_BASE_URL"] ?? "https://app.flora.ai/api/v1";
    const apiKey = global["api-key"] ?? process.env["FLORA_API_KEY"] ?? "";
    const res = await fetch(`${baseURL}/runs/${encodeURIComponent(args["run-id"])}`, {
      headers: { Authorization: `Bearer ${apiKey}` },
    });
    if (!res.ok) {
      console.error(`API Error (${res.status}): ${await res.text()}`);
      process.exit(1);
    }
    const data = await res.json();
    console.log(JSON.stringify(data, null, global.format === "raw" ? 0 : 2));
  },
});

const startTechnique = defineCommand({
  meta: { name: "start-technique", description: "Start a technique run via the top-level run resource" },
  args: {
    "technique-id": { type: "string", description: "Technique identifier", required: true },
    "workspace-id": { type: "string", description: "Workspace identifier", required: true },
    inputs: { type: "string", description: "Run inputs as JSON", required: true },
    ...pollArgs,
  },
  async run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    const inputs = parseJsonArg(args.inputs, "inputs");
    try {
      const res = await client.runs.startTechnique({
        technique_id: args["technique-id"],
        workspace_id: args["workspace-id"],
        inputs,
      });

      printResult(res, global);

      const shouldPoll = global.poll || global.download;
      if (shouldPoll) {
        const runId = res.run_id;
        const pollUrl = res.poll_url ?? undefined;
        console.error(`\nPolling run ${runId}...`);
        const status = await pollRun(runId, global, pollUrl);
        printResult(status, global);

        if (status.status === "failed") {
          console.error(`Run failed: ${status.error_message ?? status.error_code ?? "unknown"}`);
          process.exit(1);
        }

        if (global.download) {
          await downloadOutputs(status, global);
        }
      }
    } catch (err) {
      handleError(err);
    }
  },
});

export default defineCommand({
  meta: { name: "runs", description: "Run management" },
  subCommands: { retrieve, "start-generation": startGeneration, "start-technique": startTechnique },
});
