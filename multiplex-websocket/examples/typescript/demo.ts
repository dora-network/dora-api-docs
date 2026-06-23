// Runnable end-to-end demo for the DORA wsplex multiplexed WebSocket.

import { PlexClient } from "./client.js";

const DEFAULT_PRICE_IDS = [
  "019c3401-9737-7106-b3d3-b7a6e6eef0e6",
  "019c4d37-311e-7a2f-8d58-f17c39170865",
];
const DEFAULT_ORDER_BOOK_ID = "019c3420-5cd7-7a88-8fe6-a5a622e01ad9";

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

async function main(): Promise<number> {
  const picked = pickBaseUrlAndKey();
  if (!picked || !picked.apiKey) {
    console.error("wsplex demo: set DORA_STAGING_BASE_URL+DORA_STAGING_API_KEY (or _PROD_ equivalents)");
    return 2;
  }

  const priceIds = splitNonEmpty(process.env.DORA_DEMO_PRICE_ASSET_IDS, DEFAULT_PRICE_IDS);
  const orderBookId = process.env.DORA_DEMO_ORDER_BOOK_ID || DEFAULT_ORDER_BOOK_ID;

  const client = await PlexClient.connect({
    url: toPlexUrl(picked.url),
    authHeader: `ApiKey ${picked.apiKey}`,
  });

  const onPrices = (data: unknown): void => console.log(`[notif /prices] ${JSON.stringify(data)}`);
  const onTrades = (data: unknown): void => console.log(`[notif /trades] ${JSON.stringify(data)}`);

  try {
    const r1 = await client.request("/prices", { subscribe: priceIds }, onPrices);
    console.log(`[resp /prices subscribe] ${JSON.stringify(r1)}`);

    const r2 = await client.request(
      "/trades",
      { subscribe: [{ order_book_ids: [orderBookId] }] },
      onTrades,
    );
    console.log(`[resp /trades subscribe] ${JSON.stringify(r2)}`);

    const stop = new Promise<void>((resolve) => {
      const onSignal = (): void => {
        resolve();
      };
      process.once("SIGINT", onSignal);
      process.once("SIGTERM", onSignal);
      setTimeout(resolve, 10_000);
    });
    await stop;

    console.log("wsplex demo: shutting down");

    const r3 = await client.request("/prices", { unsubscribe: priceIds });
    console.log(`[resp /prices unsubscribe] ${JSON.stringify(r3)}`);

    const r4 = await client.request("/trades", {
      unsubscribe: [{ order_book_ids: [orderBookId] }],
    });
    console.log(`[resp /trades unsubscribe] ${JSON.stringify(r4)}`);
  } finally {
    await client.close();
  }

  return 0;
}

main().then((code) => process.exit(code)).catch((err) => {
  console.error(err);
  process.exit(1);
});
