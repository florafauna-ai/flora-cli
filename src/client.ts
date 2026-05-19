import { writeFile } from "node:fs/promises";
import { basename, join } from "node:path";
import FLORA from "@flora-ai/flora";

export interface GlobalArgs {
  "api-key"?: string;
  "base-url"?: string;
  debug?: boolean;
  format?: string;
  poll?: boolean;
  download?: boolean;
  "output-dir"?: string;
}

export interface RunStatus {
  run_id: string;
  status: "pending" | "running" | "completed" | "failed";
  progress?: number;
  error_code?: string;
  error_message?: string;
  outputs?: Array<{ output_id: string; type: string; url: string }>;
  poll_url?: string;
  [key: string]: unknown;
}

export function getClient(args: GlobalArgs): FLORA {
  return new FLORA({
    apiKey: args["api-key"] ?? process.env["FLORA_API_KEY"] ?? undefined,
    baseURL: args["base-url"] ?? undefined,
  });
}

export function printResult(data: unknown, args: GlobalArgs): void {
  const fmt = args.format ?? "pretty";
  if (fmt === "raw") {
    console.log(JSON.stringify(data));
  } else {
    console.log(JSON.stringify(data, null, 2));
  }
}

export function handleError(err: unknown): never {
  if (err instanceof FLORA.APIError) {
    console.error(`API Error (${err.status}): ${err.message}`);
    process.exit(1);
  }
  if (err instanceof Error) {
    console.error(`Error: ${err.message}`);
    process.exit(1);
  }
  console.error("Unknown error:", err);
  process.exit(1);
}

export function parseJsonArg(value: string | undefined, name: string): Record<string, unknown> {
  if (!value) {
    console.error(`Missing required JSON argument: ${name}`);
    process.exit(1);
  }
  try {
    return JSON.parse(value) as Record<string, unknown>;
  } catch {
    console.error(`Invalid JSON for ${name}: ${value}`);
    process.exit(1);
  }
}

function resolveBaseURL(args: GlobalArgs): string {
  return args["base-url"] ?? process.env["FLORA_BASE_URL"] ?? "https://app.flora.ai/api/v1";
}

function resolveApiKey(args: GlobalArgs): string {
  return args["api-key"] ?? process.env["FLORA_API_KEY"] ?? "";
}

async function fetchRunStatus(pollUrl: string, apiKey: string): Promise<RunStatus> {
  const res = await fetch(pollUrl, {
    headers: { Authorization: `Bearer ${apiKey}` },
  });
  if (!res.ok) {
    const text = await res.text();
    throw new Error(`Poll failed (${res.status}): ${text}`);
  }
  return (await res.json()) as RunStatus;
}

export async function pollRun(
  runId: string,
  args: GlobalArgs,
  pollUrlOverride?: string,
): Promise<RunStatus> {
  const baseURL = resolveBaseURL(args);
  const apiKey = resolveApiKey(args);
  const pollUrl = pollUrlOverride ?? `${baseURL}/runs/${encodeURIComponent(runId)}`;
  const intervalMs = 2000;
  const maxAttempts = 300; // 10 minutes max

  for (let i = 0; i < maxAttempts; i++) {
    const status = await fetchRunStatus(pollUrl, apiKey);

    if (args.debug) {
      console.error(`[poll ${i + 1}] status=${status.status} progress=${status.progress ?? "?"}`);
    } else {
      const pct = status.progress != null ? `${Math.round(status.progress * 100)}%` : "...";
      process.stderr.write(`\r  ${status.status} ${pct}  `);
    }

    if (status.status === "completed" || status.status === "failed") {
      if (!args.debug) process.stderr.write("\n");
      return status;
    }

    await new Promise((r) => setTimeout(r, intervalMs));
  }

  if (!args.debug) process.stderr.write("\n");
  throw new Error("Polling timed out after 10 minutes");
}

export async function downloadOutputs(
  status: RunStatus,
  args: GlobalArgs,
): Promise<string[]> {
  const outputs = status.outputs ?? [];
  if (outputs.length === 0) {
    console.error("No outputs to download.");
    return [];
  }

  const outDir = args["output-dir"] ?? ".";
  const downloaded: string[] = [];

  for (const output of outputs) {
    if (!output.url) continue;
    const url = new URL(output.url);
    const filename = basename(url.pathname) || `${output.output_id}.bin`;
    const dest = join(outDir, filename);

    console.error(`Downloading ${output.type}: ${filename}`);
    const res = await fetch(output.url);
    if (!res.ok) {
      console.error(`  Failed (${res.status}): ${output.url}`);
      continue;
    }
    const buf = Buffer.from(await res.arrayBuffer());
    await writeFile(dest, buf);
    console.error(`  Saved to ${dest} (${buf.length} bytes)`);
    downloaded.push(dest);
  }

  return downloaded;
}

export const pollArgs = {
  poll: {
    type: "boolean" as const,
    description: "Poll until the run completes or fails",
  },
  download: {
    type: "boolean" as const,
    description: "Download output files when the run completes (implies --poll)",
  },
  "output-dir": {
    type: "string" as const,
    description: "Directory to save downloaded outputs (default: current dir)",
  },
};
