# Flora CLI

The official CLI for the Flora REST API, built with [citty](https://github.com/unjs/citty) and the [@flora-ai/flora](https://www.npmjs.com/package/@flora-ai/flora) TypeScript SDK.

## Installation

Requires Node.js 18+.

```sh
npm install -g @flora-ai/cli
```

### Running Locally

```sh
git clone https://github.com/florafauna-ai/flora-cli.git
cd flora-cli
pnpm install
pnpm dev -- --help
```

## Usage

The CLI follows a resource-based command structure:

```sh
flora [OPTIONS] <resource> <command> [flags...]
```

```sh
flora workspaces list --api-key 'My API Key'
```

For details about specific commands, use the `--help` flag.

### Environment variables

| Environment variable | Required | Default value |
| -------------------- | -------- | ------------- |
| `FLORA_API_KEY`      | no       | `null`        |

### Global flags

- `--api-key` — Flora API key (or set `FLORA_API_KEY`)
- `--base-url` — Override the base URL for API requests
- `--debug` — Enable debug logging
- `--format` — Output format: `pretty` (default), `raw`
- `--help` — Show command usage
- `--version` — Show the CLI version

### Resources

| Command | Subcommands |
| --- | --- |
| `workspaces` | `list` |
| `projects` | `create`, `retrieve`, `list`, `list-nodes` |
| `projects:canvas` | `retrieve`, `update` |
| `projects:actions` | `create`, `run` |
| `projects:assets` | `attach` |
| `techniques` | `retrieve`, `list` |
| `technique-runs` | `create`, `retrieve` |
| `assets` | `create`, `retrieve`, `list`, `complete`, `retry` |
| `models` | `list` |
| `runs` | `start-generation`, `start-technique` |
| `generations` | `create` |
| `feedback` | `record` |
| `actions` | `list`, `get` |

### Examples

```sh
# List workspaces
flora workspaces list

# Create a project
flora projects create --workspace-id ws_abc123 --name "My Project"

# List techniques with search
flora techniques list --workspace-id ws_abc123 --query "upscale"

# Start a generation
flora generations create \
  --type image \
  --prompt "A sunset over mountains" \
  --workspace-id ws_abc123 \
  --project-id prj_abc123

# Get canvas for a project
flora projects:canvas retrieve prj_abc123
```

## Development

```sh
pnpm install       # install dependencies
pnpm dev           # run via tsx (no build needed)
pnpm build         # compile to dist/
pnpm typecheck     # type-check without emitting
```
