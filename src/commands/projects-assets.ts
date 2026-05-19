import { defineCommand } from "citty";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const attach = defineCommand({
  meta: { name: "attach", description: "Attach an asset to a project canvas" },
  args: {
    "asset-id": {
      type: "positional",
      description: "Asset identifier",
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
    return client.projects.assets
      .attachAsset(args["asset-id"], { projectId: args["project-id"] })
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "projects:assets", description: "Project asset management" },
  subCommands: { attach },
});
