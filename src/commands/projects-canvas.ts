import { defineCommand } from "citty";
import type { CanvasUpdateParams } from "@flora-ai/flora/resources/projects/canvas";
import {
  getClient,
  handleError,
  parseJsonArg,
  printResult,
  type GlobalArgs,
} from "../client.js";

const retrieve = defineCommand({
  meta: { name: "retrieve", description: "Get the canvas for a project" },
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
    return client.projects.canvas
      .retrieve(args["project-id"])
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

const update = defineCommand({
  meta: { name: "update", description: "Update the canvas for a project" },
  args: {
    "project-id": {
      type: "positional",
      description: "Project identifier",
      required: true,
    },
    diagram: {
      type: "string",
      description: "Mermaid flowchart diagram string",
      required: true,
    },
    "node-params": {
      type: "string",
      description: "Per-node parameters as JSON (keyed by Mermaid node ID)",
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    const body: CanvasUpdateParams = {
      diagram: args.diagram,
      ...(args["node-params"]
        ? {
            node_params: parseJsonArg(args["node-params"], "node-params") as CanvasUpdateParams["node_params"],
          }
        : {}),
    };
    return client.projects.canvas
      .update(args["project-id"], body)
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "projects:canvas", description: "Project canvas endpoints" },
  subCommands: { retrieve, update },
});
