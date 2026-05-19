import { defineCommand } from "citty";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const create = defineCommand({
  meta: { name: "create", description: "Create a new project" },
  args: {
    name: {
      type: "string",
      description: "Project name",
      required: true,
    },
    "workspace-id": {
      type: "string",
      description: "Workspace identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.projects
      .create({
        name: args.name,
        workspace_id: args["workspace-id"],
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const retrieve = defineCommand({
  meta: { name: "retrieve", description: "Get a project by ID" },
  args: {
    "project-id": {
      type: "positional",
      description: "Project identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.projects
      .retrieve(args["project-id"])
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const list = defineCommand({
  meta: { name: "list", description: "List projects" },
  args: {
    cursor: { type: "string", description: "Opaque cursor for pagination" },
    limit: { type: "string", description: "Maximum number of results" },
    query: { type: "string", description: "Search query" },
    "workspace-id": { type: "string", description: "Workspace identifier", required: true },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.projects
      .list({
        workspace_id: args["workspace-id"],
        ...(args.cursor ? { cursor: args.cursor } : {}),
        ...(args.limit ? { limit: Number(args.limit) } : {}),
        ...(args.query ? { query: args.query } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const listNodes = defineCommand({
  meta: { name: "list-nodes", description: "List canvas nodes for a project" },
  args: {
    "project-id": {
      type: "positional",
      description: "Project identifier",
      required: true,
    },
    cursor: { type: "string", description: "Opaque cursor for pagination" },
    limit: { type: "string", description: "Maximum number of results" },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.projects
      .listNodes(args["project-id"], {
        ...(args.cursor ? { cursor: args.cursor } : {}),
        ...(args.limit ? { limit: Number(args.limit) } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "projects", description: "Project management" },
  subCommands: { create, retrieve, list, "list-nodes": listNodes },
});
