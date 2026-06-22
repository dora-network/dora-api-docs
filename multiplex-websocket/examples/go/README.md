# Go example — `wsplex`

A small reusable client and a runnable demo for the [DORA multiplexed WebSocket protocol](../../README.md).

## Install

```bash
go mod download
```

Requires Go 1.26+.

## Configure

The demo reads three environment variables:

| Variable | Required | Default | Purpose |
|---|---|---|---|
| `DORA_STAGING_BASE_URL` *or* `DORA_PROD_BASE_URL` | yes | — | e.g. `https://staging.dora.co` |
| `DORA_STAGING_API_KEY` *or* `DORA_PROD_API_KEY` | yes | — | your API key |
| `DORA_DEMO_PRICE_ASSET_IDS` | no | `019c3401-9737-7106-b3d3-b7a6e6eef0e6,019c4d37-311e-7a2f-8d58-f17c39170865` | comma-separated asset ids to subscribe on `/prices` |
| `DORA_DEMO_ORDER_BOOK_ID` | no | `019c3420-5cd7-7a88-8fe6-a5a622e01ad9` | order book id to subscribe on `/trades` |

Example:

```bash
export DORA_STAGING_BASE_URL=https://staging.dora.co
export DORA_STAGING_API_KEY=your-key
export DORA_DEMO_PRICE_ASSET_IDS=019c3401-9737-7106-b3d3-b7a6e6eef0e6,019c4d37-311e-7a2f-8d58-f17c39170865
export DORA_DEMO_ORDER_BOOK_ID=019c3420-5cd7-7a88-8fe6-a5a622e01ad9
```

## Run

```bash
go run .
```

The demo:

1. Connects to `wss://<base_url>/plex`.
2. Subscribes to two asset ids on `/prices`.
3. Subscribes to one order book on `/trades`.
4. Logs responses and notifications for ~10 seconds.
5. On `Ctrl+C` or the 10-second timer: sends explicit unsubscribes, logs the responses, closes the socket.

## Verify syntax

```bash
go build ./...
```

Exits 0 if the demo and helper compile cleanly. No live server is contacted.

## Use the helper in your own code

```go
import plex "github.com/dora-network/dora-api-docs/multiplex-websocket/examples/go/plex"

func main() {
    client, _ := plex.Connect(context.Background(), plex.Options{URL: "wss://staging.dora.co/plex", AuthHeader: "ApiKey xxx"})
    defer client.Close()
    data, _ := client.Request(context.Background(), "/prices", map[string]any{"subscribe": []string{"<asset-id>"}}, nil)
    _ = data
}
```

The helper takes the request `data` as a `map[string]any` (or any JSON-marshalable type) and returns the response `data` as `json.RawMessage`. See [`client.go`](./plex/client.go) for full method signatures.
