import { defineCommand } from "citty";
import type { GenerationCreateParams } from "@flora-ai/flora/resources/generations";
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

const create = defineCommand({
  meta: { name: "create", description: "Start a generation" },
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
      const res = await client.generations.create({
        type: args.type as GenerationCreateParams["type"],
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

export default defineCommand({
  meta: { name: "generations", description: "Generation endpoints" },
  subCommands: { create },
});
