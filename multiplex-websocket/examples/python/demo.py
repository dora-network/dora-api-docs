"""Runnable end-to-end demo for the DORA wsplex multiplexed WebSocket.

Demonstrates every documented path: /, /prices, /trades, /assets,
/orderbook/stats, /charts/candles, /accounts/balance, /pools/balance,
/orders/byuser, and /debug/notify. The /prices notification payload is a
map keyed by asset id (not an array) — the handler prints it as-is.
"""

from __future__ import annotations

import asyncio
import json
import os
import signal
from typing import Any

from client import PlexClient, PlexOptions

DEFAULT_PRICE_IDS = (
    "019c3401-9737-7106-b3d3-b7a6e6eef0e6",
    "019c4d37-311e-7a2f-8d58-f17c39170865",
)
DEFAULT_ORDER_BOOK_ID = "019c3420-5cd7-7a88-8fe6-a5a622e01ad9"
DEFAULT_USER_ID = "019c4d37-311e-7a2f-8d58-f17c39170865"
DEFAULT_ASSET_ID = "019c3401-9737-7106-b3d3-b7a6e6eef0e6"

ONE_MINUTE_NS = 60_000_000_000


def pick_base_url_and_key() -> tuple[str | None, str | None]:
    if os.environ.get("DORA_STAGING_BASE_URL"):
        return os.environ["DORA_STAGING_BASE_URL"], os.environ.get("DORA_STAGING_API_KEY")
    if os.environ.get("DORA_PROD_BASE_URL"):
        return os.environ["DORA_PROD_BASE_URL"], os.environ.get("DORA_PROD_API_KEY")
    return None, None


def split_non_empty(s: str, fallback: tuple[str, ...]) -> list[str]:
    if not s:
        return list(fallback)
    parts = [p.strip() for p in s.split(",") if p.strip()]
    return parts or list(fallback)


def first_non_empty(s: str | None, fallback: str) -> str:
    return s if s else fallback


def to_plex_url(base_url: str) -> str:
    url = base_url.rstrip("/")
    if url.startswith("https://"):
        url = "wss://" + url[len("https://"):]
    elif url.startswith("http://"):
        url = "ws://" + url[len("http://"):]
    return url + "/plex"


def make_handler(path: str):
    def handler(data: dict[str, Any]) -> None:
        print(f"[notif {path}] {json.dumps(data)}", flush=True)
    return handler


async def run() -> int:
    base_url, api_key = pick_base_url_and_key()
    if not base_url or not api_key:
        print("wsplex demo: set DORA_STAGING_BASE_URL+DORA_STAGING_API_KEY (or _PROD_ equivalents)", flush=True)
        return 2

    price_ids = split_non_empty(os.environ.get("DORA_DEMO_PRICE_ASSET_IDS", ""), DEFAULT_PRICE_IDS)
    order_book_id = first_non_empty(os.environ.get("DORA_DEMO_ORDER_BOOK_ID"), DEFAULT_ORDER_BOOK_ID)
    user_id = first_non_empty(os.environ.get("DORA_DEMO_USER_ID"), DEFAULT_USER_ID)
    asset_id = first_non_empty(os.environ.get("DORA_DEMO_ASSET_ID"), DEFAULT_ASSET_ID)

    client = await PlexClient.connect(PlexOptions(
        url=to_plex_url(base_url),
        auth_header=f"ApiKey {api_key}",
    ))

    async def req(path: str, data: dict[str, Any], label: str, notif: Any = None) -> None:
        resp = await client.request(path, data, on_notification=notif)
        print(f"[resp {label}] {json.dumps(resp)}", flush=True)

    async def unsub(path: str, data: dict[str, Any], label: str) -> None:
        try:
            resp = await client.request(path, data)
            print(f"[resp {label} unsubscribe] {json.dumps(resp)}", flush=True)
        except RuntimeError as exc:
            print(f"wsplex demo: {label} unsubscribe: {exc}", file=__import__("sys").stderr, flush=True)

    try:
        # --- / (list routes) ---
        await req("/", {}, "/ routes")

        # --- /prices (subscribe to specific asset ids; notifications are a map keyed by asset id) ---
        await req("/prices", {"subscribe": price_ids}, "/prices subscribe",
                  notif=make_handler("/prices"))

        # --- /trades (one order book, all users) ---
        await req("/trades", {"subscribe": [{"order_book_ids": [order_book_id]}]},
                  "/trades subscribe", notif=make_handler("/trades"))

        # --- /assets (full asset updates) ---
        await req("/assets", {"subscribe": True}, "/assets subscribe",
                  notif=make_handler("/assets"))

        # --- /orderbook/stats (all order books) ---
        await req("/orderbook/stats", {"subscribe_all": True}, "/orderbook/stats subscribe",
                  notif=make_handler("/orderbook/stats"))

        # --- /charts/candles (1-minute candles for one order book) ---
        await req("/charts/candles",
                  {"subscribe": {"orderbook_ids": [order_book_id], "resolution": ONE_MINUTE_NS}},
                  "/charts/candles subscribe", notif=make_handler("/charts/candles"))

        # --- /accounts/balance (auth required; one user) ---
        await req("/accounts/balance", {"subscribe": [user_id]}, "/accounts/balance subscribe",
                  notif=make_handler("/accounts/balance"))

        # --- /pools/balance (all pools) ---
        await req("/pools/balance", {"subscribe_all": True}, "/pools/balance subscribe",
                  notif=make_handler("/pools/balance"))

        # --- /orders/byuser (auth required; one user, all order books) ---
        await req("/orders/byuser", {"user_id": user_id, "subscribe_all_orderbooks": True},
                  "/orders/byuser subscribe", notif=make_handler("/orders/byuser"))

        # --- /debug/notify (echo a ping after 100ms) ---
        await req("/debug/notify", {"delay": 100_000_000, "data": {"ping": "pong", "asset_id": asset_id}},
                  "/debug/notify", notif=make_handler("/debug/notify"))

        stop = asyncio.Event()
        loop = asyncio.get_running_loop()
        for sig in (signal.SIGINT, signal.SIGTERM):
            try:
                loop.add_signal_handler(sig, stop.set)
            except NotImplementedError:
                pass

        try:
            await asyncio.wait_for(stop.wait(), timeout=10.0)
        except asyncio.TimeoutError:
            pass

        print("wsplex demo: shutting down", flush=True)

        # Clean up subscriptions.
        await unsub("/prices", {"unsubscribe": price_ids}, "/prices")
        await unsub("/trades", {"unsubscribe": [{"order_book_ids": [order_book_id]}]}, "/trades")
        await unsub("/assets", {"subscribe": False}, "/assets")
        await unsub("/orderbook/stats", {"unsubscribe_all": True}, "/orderbook/stats")
        await unsub("/charts/candles", {"unsubscribe": {"all": True}}, "/charts/candles")
        await unsub("/accounts/balance", {"unsubscribe": [user_id]}, "/accounts/balance")
        await unsub("/pools/balance", {"unsubscribe_all": True}, "/pools/balance")
        await unsub("/orders/byuser", {"user_id": user_id, "unsubscribe_all_orderbooks": True}, "/orders/byuser")
    finally:
        await client.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(asyncio.run(run()))
