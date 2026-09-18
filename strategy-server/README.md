# Strategy Server — User Guide

The Dora Strategy Server is an HTTP API for trading bond strategies on the
Dora Network. This guide covers everything an authenticated API user needs:
backtesting, running prebuilt strategies live, and using the built-in AI agent
to create, backtest, and deploy custom strategies.

## Contents

- [Features](#features)
- [Base URLs](#base-urls)
- [Authentication](#authentication)
- [Conventions](#conventions)
- [Prebuilt strategies](#prebuilt-strategies)
- [Discovery endpoints](#discovery-endpoints)
- [Backtesting a prebuilt strategy](#backtesting-a-prebuilt-strategy)
- [Running a prebuilt strategy live](#running-a-prebuilt-strategy-live)
- [Building your own strategy with the agent](#building-your-own-strategy-with-the-agent)
- [Endpoint reference](#endpoint-reference)

## Features

| Capability | Description |
|---|---|
| **Backtesting** | Simulate any prebuilt signal strategy or agent-built strategy against historical Dora market data and retrieve P&L, drawdown, Sharpe ratio, and per-trade results. |
| **Live strategy runs** | Run prebuilt strategies against real Dora order books. Place real orders, pause/resume/stop at any time, and monitor every trading decision. |
| **Prebuilt signal strategies** | Four maintained, parameterised strategies: mean reversion, copy trading, breakout (volatility compression), and momentum (trend following). Backtestable and runnable live. |
| **Execution strategies** | TWAP and VWAP order execution: work a larger order into time- or volume-shaped slices so it completes without moving the market price as much as a one-shot large order would. Live runs only. |
| **AI agent strategy builder** | Describe a strategy in natural language; the agent generates Go code, compiles it to a sandboxed WebAssembly module, validates it, and makes it backtestable and deployable — all through the API. |
| **Real-time notifications** | A WebSocket stream of lifecycle events (run started, order filled, backtest completed, …) with replay support. |
| **Trading decision log** | Every decision a live strategy makes (with its inputs and the resulting order) is recorded and queryable. |

## Base URLs

| Environment | Base URL |
|---|---|
| **Production** | `https://strategy-prod.dora.co` |
| **Staging** | `https://strategy-staging.dora.co` |

Staging is connected to Dora's staging environment. Test integrations there
first; production places real orders.

All examples below assume:

```sh
BASE_URL="https://strategy-prod.dora.co"   # or the staging URL
DORA_API_KEY="your-dora-api-key"
TENANT_ID="your-dora-tenant-id"
```

## Authentication

Every endpoint except `/healthz` and `/v1/openapi` requires:

1. An `Authorization` header with your Dora API key:

   ```
   Authorization: ApiKey <dora-api-key>
   ```

2. A `tenant-id` header with your Dora tenant identifier:

   ```
   tenant-id: <tenant-id>
   ```

Both are required on every authenticated request. Your API key and tenant ID
are available from your Dora account settings.

> **Note:** the agent endpoints (`/v1/agent/*`) accept **only** the literal
> `ApiKey` prefix — `Bearer <token>` is rejected there. Use `ApiKey`
> everywhere for simplicity.

Verify your credentials before anything else:

```sh
curl -s "$BASE_URL/v1/dora/user" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID"
```

A `200` with your user ID means authentication is working.

## Conventions

- **Timestamps** are RFC3339 (`2026-08-01T00:00:00Z`). Query filters also
  accept bare dates (`2026-08-01`).
- **Pagination**: list endpoints take `page` (1-based) and `limit` query
  parameters and return `{ "items": [...], "page": n, "limit": m }`.
- **Async jobs**: creating a backtest or run returns immediately with a job
  record. Poll the corresponding `GET` endpoint until `status` is terminal
  (`completed`, `failed`, `cancelled`, `stopped`), or subscribe to the
  [notifications WebSocket](#real-time-notifications).
- **Errors** are JSON: `{ "error": "<message>" }` with an appropriate HTTP
  status. `401` = missing/invalid credentials, `400` = invalid request body,
  `404` = unknown ID, `409` = conflicting state, `429` = rate limited.
- **Rate limits**: the agent endpoints are rate-limited per user. If you hit
  `429`, slow down and retry.
- **Browser clients**: CORS is restricted to Dora-approved origins. Server-side
  clients and `curl` are unaffected.
- **Full API schema**: the complete OpenAPI 3.1 spec is served (no auth) at
  `GET /v1/openapi`.

## Prebuilt strategies

Six prebuilt strategies are available. Each is identified by a `strategy_type`
and configured through a `config` object. The live field list (with types,
defaults, and validation constraints) is always available from
`GET /v1/strategies`:

```sh
curl -s "$BASE_URL/v1/strategies" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" | jq
```

### Mean Reversion (`mean_reversion`)

Statistical arbitrage on a bond's spread (price minus benchmark yield). Buys
when the spread z-score drops unusually low, sells when high, closes when the
z-score reverts. Optional z-score stop-loss.

Key fields: `order_book_id`, `tenor` (benchmark Treasury tenor, e.g. `10Y`),
`lookback_window`, `entry_z_score` / `exit_z_score`, `max_position_size`,
`initial_balance` (required > 0 for backtests), `leverage`.

### Copy Trading (`copytrading`)

Mirrors trades of a followed trader across all open Dora order books, sizing
each copied order as a percentage of available balance with min/max clamps.

Key fields: `followed_trader` (UUID, required — see
[discovery endpoints](#discovery-endpoints)), `percentage_of_available`,
`min_order_size` / `max_order_size`, `disallowed_bonds`, `leverage`.

### Breakout — Volatility Compression (`breakout`)

Enters when short-window volatility compresses relative to the long window,
then price breaks out of an ATR-scaled trigger band for `confirmation_bars`
consecutive closes. Optional on-balance-volume confirmation, ATR stop-loss
and take-profit.

Key fields: `order_book_id`, `short_vol_window` / `long_vol_window`,
`compression_threshold`, `breakout_atr_multiple`, `confirmation_bars`,
`stop_loss_atr` / `take_profit_atr`, `obv_window`, `initial_balance`,
`leverage`.

### Momentum — Trend Following (`momentum`)

Rides sustained trends using a fast/slow moving-average crossover on a
configurable series — the bond's `price`, its `ytm`, or its `spread` versus a
FRED benchmark yield. Exits via ATR-anchored stop-loss, take-profit, or MA
reversal.

Key fields: `order_book_id`, `signal_source` (`price` | `ytm` | `spread`;
`spread` also requires `tenor`), `fast_window` / `slow_window`,
`stop_loss_atr` / `take_profit_atr`, `max_position_size`, `initial_balance`,
`leverage`.

### TWAP — Time-Weighted Average Price (`twap`)

An **execution strategy** for working a larger order into the market in equal
slices on a fixed clock, so the order completes over time without impacting
the market price as much as a one-shot large order would. Live runs only —
execution strategies cannot be backtested.

Splits `total_amount` evenly across the `start_time` → `end_time` window and
places a child order every `interval_seconds`. Progress is checkpointed, so a
server restart resumes the schedule rather than restarting it.

Key fields: `order_book_id`, `total_amount` (> 0), `side` (`buy` |
`sell`), `start_time` / `end_time` (ISO 8601; `end_time` must be in the
future and strictly after `start_time`), `interval_seconds` (default 300).

### VWAP — Volume-Weighted Average Price (`vwap`)

An **execution strategy** that follows the market's own volume curve instead
of a fixed clock: it sizes each child order in proportion to the historical
trade volume in that time-of-day bucket, executing more when the market is
usually liquid and less when it is quiet. Live runs only.

The bucket schedule is derived from `window_days` of historical trades on the
order book, at `bucket_minutes` granularity.

Key fields: `order_book_id`, `total_amount` (> 0), `side` (`buy` | `sell`),
`start_time` / `end_time` (ISO 8601; same rules as TWAP), `window_days`
(default 30), `bucket_minutes` (default 5).

## Discovery endpoints

Use these to find the IDs your strategy configs need:

```sh
# Order books you can trade (order_book_id values)
curl -s "$BASE_URL/v1/dora/orderbooks" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Traders available to follow (followed_trader values for copytrading)
curl -s "$BASE_URL/v1/copy-traders" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Supported benchmark Treasury tenors (tenor values)
curl -s "$BASE_URL/v1/tenors" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

## Backtesting a prebuilt strategy

Backtests are asynchronous: submit, poll until complete, then fetch results.
Backtesting is available for the four signal strategies — the execution
strategies (`twap`, `vwap`) are live-only and cannot be backtested.

**1. Submit the backtest** — `POST /v1/backtests` with the strategy type, its
config, and the historical window:

```sh
curl -s -X POST "$BASE_URL/v1/backtests" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "strategy_type": "mean_reversion",
    "config": {
      "order_book_id": "019c3420-5cd7-7a88-8fe6-a5a622e01ad9",
      "tenor": "10Y",
      "lookback_window": 20,
      "entry_z_score": 2.0,
      "exit_z_score": 0.5,
      "max_position_size": 0.5,
      "initial_balance": 100000,
      "leverage": 1.0
    },
    "start": "2026-07-01T00:00:00Z",
    "end": "2026-08-01T00:00:00Z"
  }' | jq
```

The response is `202 Accepted` with a backtest record containing its `id`.

**2. Poll status** — the light-weight metadata endpoint:

```sh
curl -s "$BASE_URL/v1/backtests/$BACKTEST_ID/metadata" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq '.status'
```

`status` moves through `running` to a terminal `completed`, `failed`, or
`cancelled`.

**3. Fetch the result summary**:

```sh
curl -s "$BASE_URL/v1/backtests/$BACKTEST_ID" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

Returns the flattened result: `total_pnl`, `win_count`, `loss_count`,
`max_drawdown`, `sharpe_ratio`, `asset_name`, `config`, and final `status`.

**4. Fetch per-trade records** (paginated):

```sh
# All trade records
curl -s "$BASE_URL/v1/backtests/$BACKTEST_ID/trades?page=1&limit=100" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Closed (round-trip) trades with per-trade P&L
curl -s "$BASE_URL/v1/backtests/$BACKTEST_ID/closed-trades?page=1&limit=100" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

**List and cancel**:

```sh
# Your backtests (filters: status, from, to; paginated)
curl -s "$BASE_URL/v1/backtests?status=completed&page=1&limit=20" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Cancel a running backtest
curl -s -X DELETE "$BASE_URL/v1/backtests/$BACKTEST_ID" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

## Running a prebuilt strategy live

A live run places **real orders** on the configured Dora order book using
your account balance. Config shape is the same as for backtests, minus the
date window.

**1. Start the run** — `POST /v1/runs`:

```sh
curl -s -X POST "$BASE_URL/v1/runs" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "strategy_type": "breakout",
    "config": {
      "order_book_id": "019c3420-5cd7-7a88-8fe6-a5a622e01ad9",
      "compression_threshold": 0.3,
      "confirmation_bars": 5,
      "stop_loss_atr": 20,
      "leverage": 1.0
    }
  }' | jq
```

Returns `201` with the run `id`. A `409` means a strategy is already active
for that order book — stop it first.

**Execution strategies** (`twap`, `vwap`) are run-only and use the same
`POST /v1/runs` endpoint with their own config shape — for example, working
500,000 units into the market in 5-minute TWAP slices over an hour:

```sh
curl -s -X POST "$BASE_URL/v1/runs" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "strategy_type": "twap",
    "config": {
      "order_book_id": "019c3420-5cd7-7a88-8fe6-a5a622e01ad9",
      "total_amount": 500000,
      "side": "buy",
      "start_time": "2026-09-18T10:00:00Z",
      "end_time": "2026-09-18T11:00:00Z",
      "interval_seconds": 300
    }
  }' | jq
```

Every child order the execution strategy places shows up in the run's
trading decision log (step 3 below), and order-fill updates flow
over the [notifications WebSocket](#real-time-notifications).

**2. Manage the run**:

```sh
curl -s -X POST "$BASE_URL/v1/runs/$RUN_ID/pause" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

curl -s -X POST "$BASE_URL/v1/runs/$RUN_ID/resume" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Stop and delete the run
curl -s -X DELETE "$BASE_URL/v1/runs/$RUN_ID" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# List your runs (run status: running | paused | stopped)
curl -s "$BASE_URL/v1/runs" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

**3. Audit trading decisions** — every decision the strategy made, newest
first:

```sh
curl -s "$BASE_URL/v1/trading-decisions/$RUN_ID" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

### Real-time notifications

`GET /v1/notifications/ws` upgrades to a WebSocket streaming JSON-encoded
lifecycle events (run status changes, order fills, backtest completions) as
text frames. Authenticate with the standard `Authorization` header; clients
that cannot set headers may pass `?x-api-key=<key>` instead.

Optional query parameters:

- `Last-Event-ID=<uuidv7>` — replay missed events after a reconnect.
- `types=<comma-separated>` — filter to specific event types.

The server sends WebSocket-level pings every 30 seconds; respond with pongs
to keep the connection alive.

## Building your own strategy with the agent

The agent (`/v1/agent/*`) turns a natural-language description into a working,
sandboxed strategy. The workflow:

1. [Configure an LLM provider](#1-configure-an-llm-provider) (once per provider)
2. [Chat: describe your strategy](#2-chat-describe-your-strategy)
3. [Save the captured strategy](#3-save-the-captured-strategy)
4. [Backtest it](#4-backtest-an-agent-strategy)
5. [Deploy it live](#5-deploy-an-agent-strategy-live)

The agent generates a Go module implementing the `dorastrategy.Strategy`
interface, compiles it to WebAssembly (TinyGo), validates it against the
framework surface in a sandboxed runtime, and records it as a versioned
strategy you own. Strategy code is validated (no network access at init,
allowlisted imports only) and runs sandboxed with memory caps both in
backtests and live deployments.

### 1. Configure an LLM provider

The agent needs credentials for the LLM that will author your strategy code.
This is separate from your Dora API key. Supported providers: `openai`,
`anthropic`, and `openrouter` (unified access to many models).

```sh
curl -s -X POST "$BASE_URL/v1/agent/provider-config" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "provider": "openai",
    "api_key": "sk-...",
    "default_model": "gpt-5"
  }' | jq
```

Fields:

| Field | Required | Notes |
|---|---|---|
| `provider` | yes | `openai`, `anthropic`, or `openrouter` |
| `api_key` | yes | Your provider API key. Stored encrypted (AES-256-GCM); only ever returned masked. |
| `default_model` | yes | Default model for new sessions. |
| `base_url` | no | Override for OpenRouter-style relays or local proxies. Must be an `https://` URL. |

**Model naming**: for `openai` and `anthropic` use the bare model family name
(`gpt-5`, `claude-opus-4`); for `openrouter` use the full `vendor/model`
address (`openai/gpt-5`), optionally with `:free` for free-tier inference.

Manage your configs:

```sh
# List configured providers (keys shown masked)
curl -s "$BASE_URL/v1/agent/provider-config" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Delete one provider's config
curl -s -X DELETE "$BASE_URL/v1/agent/provider-config/openai" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq
```

### 2. Chat: describe your strategy

Create a session bound to a provider and model:

```sh
SID=$(curl -s -X POST "$BASE_URL/v1/agent/sessions" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{"provider":"openai","model":"gpt-5","title":"My duration strategy"}' \
  | jq -r '.session_id')
```

Send your strategy description. The response is a **Server-Sent Events**
stream — assistant text streams as `delta` events, and a `strategy_saved`
event is emitted when the agent has produced and validated a strategy
version. The stream ends with a `done` event.

```sh
curl -N -X POST "$BASE_URL/v1/agent/sessions/$SID/messages" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -H "Accept: text/event-stream" \
  -d '{"prompt":"Build a mean-reversion strategy on the 10Y Treasury spread that buys when the spread z-score over a 240-tick window drops below 2, with a 20-ATR stop-loss. Order book: 019c3420-5cd7-7a88-8fe6-a5a622e01ad9."}'
```

Iterate in follow-up messages (same endpoint, same session) until you are
happy with the generated strategy — the agent refines the code and re-validates
it each turn. Session history is persisted; list sessions with
`GET /v1/agent/sessions`, replay one with `GET /v1/agent/sessions/{id}`, and
delete with `DELETE /v1/agent/sessions/{id}`.

### 3. Save the captured strategy

Each validated strategy the agent produces is stashed as a pending capture.
Confirm it to record it under your account:

```sh
curl -s -X POST "$BASE_URL/v1/agent/sessions/$SID/save" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" | jq
```

Returns a `Version` record with the `strategy_id` and its `revision`. From
here on the strategy is a first-class object:

```sh
# Your strategies
curl -s "$BASE_URL/v1/agent/strategies" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# One strategy's versions, newest first
curl -s "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/versions" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Full source code of one version
curl -s "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/versions/$REVISION" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Point the strategy's head back at a known-good revision
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/rollback" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{"revision":"'$REVISION'"}' | jq
```

### 4. Backtest an agent strategy

Backtest a specific version over a historical window of candle data:

```sh
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/versions/$REVISION/backtest" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "order_book_id": "019c3420-5cd7-7a88-8fe6-a5a622e01ad9",
    "resolution": "1h",
    "start": "2026-07-01T00:00:00Z",
    "end": "2026-08-01T00:00:00Z"
  }' | jq
```

Request fields:

| Field | Required | Notes |
|---|---|---|
| `order_book_id` | yes | Order book to trade and fetch historic candles for. |
| `resolution` | yes | Candle resolution: `1m`, `5m`, `15m`, `1h`, `4h`, `1d`. |
| `start` / `end` | yes | Backtest window (RFC3339). |
| `params` | no | Runtime parameters passed to the strategy (string map). |
| `warmup_candles` | no | Extra candles fetched before `start` so the strategy starts with history. |

The response contains a `backtest_id`. Poll it:

```sh
curl -s "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/backtests/$BACKTEST_ID" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# All backtest jobs for the strategy
curl -s "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/backtests" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Cancel a running one
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/backtests/$BACKTEST_ID/cancel" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID"
```

### 5. Deploy an agent strategy live

Deploying a version starts it in live mode: it receives real-time candles
from the order book and submits real orders on your behalf, with automatic
crash-restart within a bounded budget.

```sh
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/versions/$REVISION/deploy" \
  -H "Authorization: ApiKey $DORA_API_KEY" \
  -H "tenant-id: $TENANT_ID" \
  -H "Content-Type: application/json" \
  -d '{
    "order_book_id": "019c3420-5cd7-7a88-8fe6-a5a622e01ad9",
    "resolution": "1m"
  }' | jq
```

Request fields: `order_book_id` (required), `resolution` (required, e.g.
`1m` / `5m` / `1h`), optional `params` (string map) and `warmup_candles`.

**Manage deployments**:

```sh
# List deployments for the strategy (status: running | stopped | crashed | halted)
curl -s "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/deployments" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# One deployment's status
curl -s "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/deployments/$DEPLOYMENT_ID" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID" | jq

# Stop a running deployment
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/deployments/$DEPLOYMENT_ID/stop" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID"
```

**Safety controls**:

```sh
# Halt: kill switch for ALL your strategies (immediately stops live instances)
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/halt" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID"

# Resume: clear the kill switch
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/resume" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID"

# Restart: clear a crashed deployment's restart budget so it can restart again
curl -s -X POST "$BASE_URL/v1/agent/strategies/$STRATEGY_ID/restart" \
  -H "Authorization: ApiKey $DORA_API_KEY" -H "tenant-id: $TENANT_ID"
```

If a deployment exhausts its restart budget it stays `crashed` until you
explicitly clear the budget with `restart`.

## Endpoint reference

Full OpenAPI 3.1 spec: `GET /v1/openapi` (no auth).

| Method | Path | Purpose |
|---|---|---|
| GET | `/healthz` | Health check (no auth) |
| GET | `/v1/dora/orderbooks` | List tradeable order books |
| GET | `/v1/dora/user` | Current Dora user |
| GET | `/v1/copy-traders` | List followable traders |
| GET | `/v1/tenors` | List benchmark tenors |
| GET | `/v1/strategies` | List prebuilt strategies + config fields |
| GET / POST | `/v1/backtests` | List / create prebuilt backtests |
| GET / DELETE | `/v1/backtests/{id}` | Result summary / cancel |
| GET | `/v1/backtests/{id}/metadata` | Status + timestamps (light) |
| GET | `/v1/backtests/{id}/trades` | Paginated trade records |
| GET | `/v1/backtests/{id}/closed-trades` | Paginated closed trades |
| GET / POST | `/v1/runs` | List / start live runs |
| GET / DELETE | `/v1/runs/{id}` | Get / stop a run |
| POST | `/v1/runs/{id}/pause` | Pause |
| POST | `/v1/runs/{id}/resume` | Resume |
| GET | `/v1/trading-decisions/{run_id}` | Decision log for a run |
| GET | `/v1/notifications/ws` | WebSocket lifecycle notifications |
| GET / POST | `/v1/agent/provider-config` | List / save LLM provider configs |
| DELETE | `/v1/agent/provider-config/{provider}` | Delete a provider config |
| GET / POST | `/v1/agent/sessions` | List / create chat sessions |
| GET / DELETE | `/v1/agent/sessions/{id}` | Get (with messages) / delete |
| POST | `/v1/agent/sessions/{id}/messages` | Send prompt; SSE response |
| POST | `/v1/agent/sessions/{id}/save` | Save pending strategy capture |
| GET | `/v1/agent/strategies` | List your agent strategies |
| GET | `/v1/agent/strategies/{id}` | Get a strategy |
| GET | `/v1/agent/strategies/{id}/versions` | List versions |
| GET | `/v1/agent/strategies/{id}/versions/{rev}` | Version with source |
| POST | `/v1/agent/strategies/{id}/rollback` | Roll head to a revision |
| POST | `/v1/agent/strategies/{id}/versions/{rev}/backtest` | Backtest a version |
| GET | `/v1/agent/strategies/{id}/backtests` | List backtest jobs |
| GET | `/v1/agent/strategies/{id}/backtests/{bid}` | Backtest status/result |
| POST | `/v1/agent/strategies/{id}/backtests/{bid}/cancel` | Cancel a backtest |
| POST | `/v1/agent/strategies/{id}/versions/{rev}/deploy` | Deploy a version live |
| GET | `/v1/agent/strategies/{id}/deployments` | List deployments |
| GET | `/v1/agent/strategies/{id}/deployments/{did}` | Deployment status |
| POST | `/v1/agent/strategies/{id}/deployments/{did}/stop` | Stop a deployment |
| POST | `/v1/agent/strategies/{id}/halt` | Kill switch on |
| POST | `/v1/agent/strategies/{id}/resume` | Kill switch off |
| POST | `/v1/agent/strategies/{id}/restart` | Clear restart budget |
