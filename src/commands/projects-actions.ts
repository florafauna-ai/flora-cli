import { defineCommand } from "citty";
import type { ActionCreateParams } from "@flora-ai/flora/resources/projects/actions";
import {
  getClient,
  handleError,
  parseJsonArg,
  printResult,
  type GlobalArgs,
} from "../client.js";

const create = defineCommand({
  meta: { name: "create", description: "Create a prebuilt action node on a project canvas" },
  args: {
    "project-id": {
      type: "string",
      description: "Project identifier",
      required: true,
    },
    "action-id": {
      type: "string",
      description: "Action identifier / slug (e.g. rotate-image)",
      required: true,
    },
    params: {
      type: "string",
      description: "Optional action params as JSON",
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.projects.actions
      .create(args["project-id"], {
        action_id: args["action-id"] as ActionCreateParams["action_id"],
        ...(args.params ? { params: parseJsonArg(args.params, "params") } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const run = defineCommand({
  meta: { name: "run", description: "Run a canvas action node" },
  args: {
    "node-id": {
      type: "positional",
      description: "Canvas action node identifier",
      required: true,
    },
    "project-id": {
      type: "string",
      description: "Project identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.projects.actions
      .run(args["node-id"], { projectId: args["project-id"] })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "projects:actions", description: "Canvas action management" },
  subCommands: { create, run },
});
