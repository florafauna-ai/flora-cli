import { defineCommand } from "citty";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const list = defineCommand({
  meta: { name: "list", description: "List workspaces available to the authenticated API key" },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.workspaces
      .list()
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "workspaces", description: "Workspace discovery" },
  subCommands: { list },
});
