import FLORA from "@flora-ai/flora";

export interface GlobalArgs {
  "api-key"?: string;
  "base-url"?: string;
  debug?: boolean;
  format?: string;
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
