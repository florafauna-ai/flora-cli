import { defineCommand, runMain } from "citty";

const main = defineCommand({
  meta: {
    name: "flora",
    version: "0.1.0",
    description: "CLI for the Flora.ai API",
  },
  args: {
    "api-key": {
      type: "string",
      description: "Flora API key (or set FLORA_API_KEY)",
    },
    "base-url": {
      type: "string",
      description: "Override the base URL for API requests",
    },
    debug: {
      type: "boolean",
      description: "Enable debug logging",
    },
    format: {
      type: "string",
      description: "Output format: pretty, raw",
      default: "pretty",
    },
  },
  subCommands: {
    workspaces: () => import("./commands/workspaces.js").then((m) => m.default),
    projects: () => import("./commands/projects.js").then((m) => m.default),
    "projects:canvas": () => import("./commands/projects-canvas.js").then((m) => m.default),
    "projects:actions": () => import("./commands/projects-actions.js").then((m) => m.default),
    "projects:assets": () => import("./commands/projects-assets.js").then((m) => m.default),
    techniques: () => import("./commands/techniques.js").then((m) => m.default),
    "technique-runs": () => import("./commands/technique-runs.js").then((m) => m.default),
    assets: () => import("./commands/assets.js").then((m) => m.default),
    models: () => import("./commands/models.js").then((m) => m.default),
    runs: () => import("./commands/runs.js").then((m) => m.default),
    generations: () => import("./commands/generations.js").then((m) => m.default),
    feedback: () => import("./commands/feedback.js").then((m) => m.default),
    actions: () => import("./commands/actions.js").then((m) => m.default),
  },
});

runMain(main);
