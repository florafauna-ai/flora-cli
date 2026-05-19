import { defineCommand } from "citty";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const create = defineCommand({
  meta: { name: "create", description: "Upload a new asset" },
  args: {
    source: {
      type: "string",
      description: "Asset source URL",
      required: true,
    },
    "workspace-id": {
      type: "string",
      description: "Workspace identifier",
      required: true,
    },
    folder: {
      type: "string",
      description: "Optional folder path",
    },
    "content-type": {
      type: "string",
      description: "MIME content type",
    },
    "file-name": {
      type: "string",
      description: "File name",
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.assets
      .create({
        source: args.source,
        workspace_id: args["workspace-id"],
        ...(args.folder ? { folder: args.folder } : {}),
        ...(args["content-type"] ? { content_type: args["content-type"] } : {}),
        ...(args["file-name"] ? { file_name: args["file-name"] } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const retrieve = defineCommand({
  meta: { name: "retrieve", description: "Get an asset by ID" },
  args: {
    "asset-id": {
      type: "positional",
      description: "Asset identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.assets
      .retrieve(args["asset-id"])
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const list = defineCommand({
  meta: { name: "list", description: "List assets" },
  args: {
    cursor: { type: "string", description: "Opaque cursor for pagination" },
    limit: { type: "string", description: "Maximum number of results" },
    query: { type: "string", description: "Search query" },
    "workspace-id": { type: "string", description: "Workspace identifier" },
    "project-id": { type: "string", description: "Project identifier" },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.assets
      .list({
        ...(args.cursor ? { cursor: args.cursor } : {}),
        ...(args.limit ? { limit: Number(args.limit) } : {}),
        ...(args.query ? { query: args.query } : {}),
        ...(args["workspace-id"] ? { workspace_id: args["workspace-id"] } : {}),
        ...(args["project-id"] ? { project_id: args["project-id"] } : {}),
      })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const complete = defineCommand({
  meta: { name: "complete", description: "Mark an asset upload as complete" },
  args: {
    "asset-id": {
      type: "positional",
      description: "Asset identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.assets
      .complete(args["asset-id"])
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const retry = defineCommand({
  meta: { name: "retry", description: "Retry a failed asset upload" },
  args: {
    "asset-id": {
      type: "positional",
      description: "Asset identifier",
      required: true,
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.assets
      .retry(args["asset-id"])
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "assets", description: "Asset upload and retrieval" },
  subCommands: { create, retrieve, list, complete, retry },
});
