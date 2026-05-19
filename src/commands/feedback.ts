import { defineCommand } from "citty";
import type { FeedbackRecordParams } from "@flora-ai/flora/resources/feedback";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const record = defineCommand({
  meta: { name: "record", description: "Record product feedback" },
  args: {
    kind: { type: "string", description: "Feedback kind", required: true },
    summary: { type: "string", description: "Short summary", required: true },
    detail: { type: "string", description: "Detailed description", required: true },
    "workspace-id": { type: "string", description: "Workspace identifier" },
    "project-id": { type: "string", description: "Project identifier" },
    "run-id": { type: "string", description: "Run identifier" },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.feedback
      .record({
        kind: args.kind as FeedbackRecordParams["kind"],
        summary: args.summary,
        detail: args.detail,
        ...(args["workspace-id"] ? { workspace_id: args["workspace-id"] } : {}),
        ...(args["project-id"] ? { project_id: args["project-id"] } : {}),
        ...(args["run-id"] ? { run_id: args["run-id"] } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "feedback", description: "Product feedback" },
  subCommands: { record },
});
