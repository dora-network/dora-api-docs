# DORA REST API Auto-generated documentation

This repository contains the auto-generated documentation for the DORA (U.S. Patent 12,596,578) REST API.
The documentation is generated from the OpenAPI specification file `openapi.json`.
It uses the [Swagger Codegen CLI tool](https://swagger.io/tools/swagger-codegen/) to generate the documentation.

## Pre-requisites

To generate the documentation, you need to have the following tools installed and in your PATH:

- java
- make

To check if you have these tools installed, run the following commands:

```bash
java --version
make --version
```

## Generating the documentation

To generate the documentation, run the following command in the root directory of the repository:

```bash
make gen
```

## Getting started

Follow the getting started guide in the [Getting started guide](getting-started.md).

## AI Agent Integration

This repository ships with a [Dora Network API Skill](../skills/dora-api/) designed for AI agent integration. The skill enables integrators and their users to interact with the Dora REST API using environment-specific base URLs and API keys.

### What it provides

- **Environment-aware API access** — Configurable base URLs and API keys for `staging` and `production` environments via `DORA_*_BASE_URL` and `DORA_*_API_KEY` environment variables.
- **API Key authentication** — Uses `Authorization: ApiKey <key>` on all authenticated endpoints.
- **Ready-to-use curl examples** — Common operations (list assets, fetch trades, get orderbook data, WebSocket streams) are pre-built and copy-paste ready.

### How to use it

1. Install the skill into your AI agent's skill directory.
2. Set your environment variables for the target environment:
   ```
   DORA_STAGING_BASE_URL=https://staging.dora.co/
   DORA_STAGING_API_KEY=your_dora_staging_key
   ```
   If `*_BASE_URL` variables are not set, the skill falls back to base URLs defined in the OpenAPI spec.
3. Reference the skill in your agent session. The agent will auto-load the curl commands and authentication patterns.

### Quick start

```bash
export DORA_BASE_URL=$DORA_STAGING_BASE_URL
export DORA_API_KEY=$DORA_STAGING_API_KEY
curl -L -H "Authorization: ApiKey $DORA_API_KEY" "$DORA_BASE_URL/v1/assets?limit=10"
```

### Platform support

The skill works on Linux, macOS, and Windows.
