import { defineCommand } from "citty";
import { type GlobalArgs } from "../client.js";

const list = defineCommand({
  meta: { name: "list", description: "List released prebuilt Flora actions" },
  args: {},
  async run({ args }) {
    const global = args as unknown as GlobalArgs;
    const baseURL = global["base-url"] ?? process.env["FLORA_BASE_URL"] ?? "https://app.flora.ai/api/v1";
    const apiKey = global["api-key"] ?? process.env["FLORA_API_KEY"] ?? "";
    const res = await fetch(`${baseURL}/actions`, {
      headers: { Authorization: `Bearer ${apiKey}` },
    });
    if (!res.ok) {
      console.error(`API Error (${res.status}): ${await res.text()}`);
      process.exit(1);
    }
    const data = await res.json();
    console.log(JSON.stringify(data, null, global.format === "raw" ? 0 : 2));
  },
});

const get = defineCommand({
  meta: { name: "get", description: "Get a specific action by ID" },
  args: {
    "action-id": {
      type: "positional",
      description: "Action identifier or slug",
      required: true,
    },
  },
  async run({ args }) {
    const global = args as unknown as GlobalArgs;
    const baseURL = global["base-url"] ?? process.env["FLORA_BASE_URL"] ?? "https://app.flora.ai/api/v1";
    const apiKey = global["api-key"] ?? process.env["FLORA_API_KEY"] ?? "";
    const res = await fetch(`${baseURL}/actions/${encodeURIComponent(args["action-id"])}`, {
      headers: { Authorization: `Bearer ${apiKey}` },
    });
    if (!res.ok) {
      console.error(`API Error (${res.status}): ${await res.text()}`);
      process.exit(1);
    }
    const data = await res.json();
    console.log(JSON.stringify(data, null, global.format === "raw" ? 0 : 2));
  },
});

export default defineCommand({
  meta: { name: "actions", description: "Prebuilt action catalog" },
  subCommands: { list, get },
});
