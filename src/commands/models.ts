import { defineCommand } from "citty";
import type { ModelListParams } from "@flora-ai/flora/resources/models";
import {
  getClient,
  handleError,
  printResult,
  type GlobalArgs,
} from "../client.js";

const list = defineCommand({
  meta: { name: "list", description: "List available models" },
  args: {
    type: {
      type: "string",
      description: "Filter by model type (e.g. text, image, video)",
    },
  },
  run({ args }) {
    const global = args as unknown as GlobalArgs;
    const client = getClient(global);
    return client.models
      .list(args.type ? { type: args.type as ModelListParams["type"] } : {})
      .then((res) => printResult(res, global))
      .catch(handleError);
  },
});

export default defineCommand({
  meta: { name: "models", description: "Model catalog" },
  subCommands: { list },
});
