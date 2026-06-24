# TypeScript example — `wsplex`

A small reusable client and a runnable demo for the [DORA multiplexed WebSocket protocol](../../README.md).

## Install

```bash
npm install
```

Requires Node 22+.

## Configure

| Variable | Required | Default | Purpose |
|---|---|---|---|
| `DORA_STAGING_BASE_URL` *or* `DORA_PROD_BASE_URL` | yes | — | e.g. `https://staging.dora.co` |
| `DORA_STAGING_API_KEY` *or* `DORA_PROD_API_KEY` | yes | — | your API key |
| `DORA_DEMO_PRICE_ASSET_IDS` | no | `019c3401-9737-7106-b3d3-b7a6e6eef0e6,019c4d37-311e-7a2f-8d58-f17c39170865` | comma-separated asset ids |
| `DORA_DEMO_ORDER_BOOK_ID` | no | `019c3420-5cd7-7a88-8fe6-a5a622e01ad9` | order book id |

```bash
export DORA_STAGING_BASE_URL=https://staging.dora.co
export DORA_STAGING_API_KEY=your-key
```

## Run

```bash
npm run demo
```

Connects, subscribes to `/prices` and `/trades`, logs responses and notifications for ~10 seconds, then sends explicit unsubscribes and closes.

## Verify syntax

```bash
npm run typecheck
```

Exits 0 if the demo and helper compile cleanly under `strict: true`. No live server is contacted.

## Use the helper in your own code

```ts
import { PlexClient, PlexOptions } from "./client.js";

const client = await PlexClient.connect({
  url: "wss://staging.dora.co/plex",
  authHeader: "ApiKey xxx",
  userAgent: "MyClient/1.0",
} satisfies PlexOptions);
const data = await client.request("/prices", { subscribe: ["<asset-id>"] });
await client.close();
```
