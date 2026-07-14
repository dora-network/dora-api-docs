// Runnable end-to-end demo for the DORA wsplex multiplexed WebSocket.
//
// Demonstrates every documented path: /, /prices, /trades, /assets,
// /orderbook/stats, /charts/candles, /accounts/balance, /pools/balance,
// /orders/byuser, and /debug/notify. The /prices notification payload is a
// map keyed by asset id (not an array) — the handler prints it as-is.

import { PlexClient } from "./client.js";

const DEFAULT_PRICE_IDS = [
  "019c3401-9737-7106-b3d3-b7a6e6eef0e6",
  "019c4d37-311e-7a2f-8d58-f17c39170865",
];
const DEFAULT_ORDER_BOOK_ID = "019c3420-5cd7-7a88-8fe6-a5a622e01ad9";
const DEFAULT_USER_ID = "019c4d37-311e-7a2f-8d58-f17c39170865";
const DEFAULT_ASSET_ID = "019c3401-9737-7106-b3d3-b7a6e6eef0e6";


function pickBaseUrlAndKey(): { url: string; apiKey: string } | null {
  const staging = process.env.DORA_STAGING_BASE_URL;
  if (staging) return { url: staging, apiKey: process.env.DORA_STAGING_API_KEY ?? "" };
  const prod = process.env.DORA_PROD_BASE_URL;
  if (prod) return { url: prod, apiKey: process.env.DORA_PROD_API_KEY ?? "" };
  return null;
}

function toPlexUrl(baseUrl: string): string {
  let url = baseUrl.replace(/\/$/, "");
  if (url.startsWith("https://")) url = "wss://" + url.slice("https://".length);
  else if (url.startsWith("http://")) url = "ws://" + url.slice("http://".length);
  return url + "/plex";
}

function splitNonEmpty(value: string | undefined, fallback: readonly string[]): string[] {
  if (!value) return [...fallback];
  const parts = value.split(",").map((s) => s.trim()).filter(Boolean);
  return parts.length > 0 ? parts : [...fallback];
}

function firstNonEmpty(value: string | undefined, fallback: string): string {
  return value || fallback;
}

function makeHandler(path: string): (data: unknown) => void {
  return (data: unknown): void => console.log(`[notif ${path}] ${JSON.stringify(data)}`);
}

async function main(): Promise<number> {
  const picked = pickBaseUrlAndKey();
  if (!picked || !picked.apiKey) {
    console.error("wsplex demo: set DORA_STAGING_BASE_URL+DORA_STAGING_API_KEY (or _PROD_ equivalents)");
    return 2;
  }

  const priceIds = splitNonEmpty(process.env.DORA_DEMO_PRICE_ASSET_IDS, DEFAULT_PRICE_IDS);
  const orderBookId = firstNonEmpty(process.env.DORA_DEMO_ORDER_BOOK_ID, DEFAULT_ORDER_BOOK_ID);
  const userId = firstNonEmpty(process.env.DORA_DEMO_USER_ID, DEFAULT_USER_ID);
  const assetId = firstNonEmpty(process.env.DORA_DEMO_ASSET_ID, DEFAULT_ASSET_ID);

  const client = await PlexClient.connect({
    url: toPlexUrl(picked.url),
    authHeader: `ApiKey ${picked.apiKey}`,
  });

  async function req(
    path: string,
    data: Record<string, unknown>,
    label: string,
    onNotif?: (data: unknown) => void,
  ): Promise<void> {
    const resp = await client.request(path, data, onNotif);
    console.log(`[resp ${label}] ${JSON.stringify(resp)}`);
  }

  async function unsub(path: string, data: Record<string, unknown>, label: string): Promise<void> {
    try {
      const resp = await client.request(path, data);
      console.log(`[resp ${label} unsubscribe] ${JSON.stringify(resp)}`);
    } catch (err) {
      console.error(`wsplex demo: ${label} unsubscribe: ${err instanceof Error ? err.message : err}`);
    }
  }

  try {
    // --- / (list routes) ---
    await req("/", {}, "/ routes");

    // --- /prices (subscribe to specific asset ids; notifications are a map keyed by asset id) ---
    await req("/prices", { subscribe: priceIds }, "/prices subscribe", makeHandler("/prices"));

    // --- /trades (one order book, all users) ---
    await req(
      "/trades",
      { subscribe: [{ order_book_ids: [orderBookId] }] },
      "/trades subscribe",
      makeHandler("/trades"),
    );

    // --- /assets (full asset updates) ---
    await req("/assets", { subscribe: true }, "/assets subscribe", makeHandler("/assets"));

    // --- /orderbook/stats (all order books) ---
    await req(
      "/orderbook/stats",
      { subscribe_all: true },
      "/orderbook/stats subscribe",
      makeHandler("/orderbook/stats"),
    );

    // --- /charts/candles (1-minute candles for one order book) ---
    await req(
      "/charts/candles",
      { subscribe: { orderbook_ids: [orderBookId], resolution: "1m" } },
      "/charts/candles subscribe",
      makeHandler("/charts/candles"),
    );

    // --- /accounts/balance (auth required; one user) ---
    await req(
      "/accounts/balance",
      { subscribe: [userId] },
      "/accounts/balance subscribe",
      makeHandler("/accounts/balance"),
    );

    // --- /pools/balance (all pools) ---
    await req(
      "/pools/balance",
      { subscribe_all: true },
      "/pools/balance subscribe",
      makeHandler("/pools/balance"),
    );

    // --- /orders/byuser (auth required; one user, all order books) ---
    await req(
      "/orders/byuser",
      { user_id: userId, subscribe_all_orderbooks: true },
      "/orders/byuser subscribe",
      makeHandler("/orders/byuser"),
    );

    // --- /debug/notify (echo a ping after 100ms) ---
    await req(
      "/debug/notify",
      { delay: 100_000_000, data: { ping: "pong", asset_id: assetId } },
      "/debug/notify",
      makeHandler("/debug/notify"),
    );

    const stop = new Promise<void>((resolve) => {
      const onSignal = (): void => resolve();
      process.once("SIGINT", onSignal);
      process.once("SIGTERM", onSignal);
      setTimeout(resolve, 10_000);
    });
    await stop;

    console.log("wsplex demo: shutting down");

    // Clean up subscriptions.
    await unsub("/prices", { unsubscribe: priceIds }, "/prices");
    await unsub("/trades", { unsubscribe: [{ order_book_ids: [orderBookId] }] }, "/trades");
    await unsub("/assets", { subscribe: false }, "/assets");
    await unsub("/orderbook/stats", { unsubscribe_all: true }, "/orderbook/stats");
    await unsub("/charts/candles", { unsubscribe: { all: true } }, "/charts/candles");
    await unsub("/accounts/balance", { unsubscribe: [userId] }, "/accounts/balance");
    await unsub("/pools/balance", { unsubscribe_all: true }, "/pools/balance");
    await unsub("/orders/byuser", { user_id: userId, unsubscribe_all_orderbooks: true }, "/orders/byuser");
  } finally {
    await client.close();
  }

  return 0;
}

main().then((code) => process.exit(code)).catch((err) => {
  console.error(err);
  process.exit(1);
});
