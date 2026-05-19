import { defineCommand } from "citty";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const retrieve = defineCommand({
  meta: { name: "retrieve", description: "Get a technique by ID" },
  args: {
    "technique-id": {
      type: "positional",
      description: "Technique identifier or slug",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.techniques
      .retrieve(args["technique-id"])
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const list = defineCommand({
  meta: { name: "list", description: "List techniques" },
  args: {
    cursor: { type: "string", description: "Opaque cursor for pagination" },
    limit: { type: "string", description: "Maximum number of results" },
    query: { type: "string", description: "Search query" },
    "workspace-id": { type: "string", description: "Workspace identifier" },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.techniques
      .list({
        ...(args.cursor ? { cursor: args.cursor } : {}),
        ...(args.limit ? { limit: Number(args.limit) } : {}),
        ...(args.query ? { query: args.query } : {}),
        ...(args["workspace-id"] ? { workspace_id: args["workspace-id"] } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "techniques", description: "Technique catalog" },
  subCommands: { retrieve, list },
});
