import { defineCommand } from "citty";
import type { RunStartGenerationParams } from "@flora-ai/flora/resources/runs";
import {
  getClient,
  handleError,
  parseJsonArg,
  printResult,
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
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.runs
      .startGeneration({
        type: args.type as RunStartGenerationParams["type"],
        prompt: args.prompt,
        workspace_id: args["workspace-id"],
        project_id: args["project-id"],
        ...(args.model ? { model: args.model } : {}),
        ...(args.params
          ? { params: parseJsonArg(args.params, "params") as Record<string, unknown> }
          : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const startTechnique = defineCommand({
  meta: { name: "start-technique", description: "Start a technique run via the top-level run resource" },
  args: {
    "technique-id": { type: "string", description: "Technique identifier", required: true },
    "workspace-id": { type: "string", description: "Workspace identifier", required: true },
    inputs: { type: "string", description: "Run inputs as JSON", required: true },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    const inputs = parseJsonArg(args.inputs, "inputs");
    return client.runs
      .startTechnique({
        technique_id: args["technique-id"],
        workspace_id: args["workspace-id"],
        inputs,
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "runs", description: "Top-level run creation" },
  subCommands: { "start-generation": startGeneration, "start-technique": startTechnique },
});
