"""Runnable end-to-end demo for the DORA wsplex multiplexed WebSocket."""

from __future__ import annotations

import asyncio
import os
import signal
from typing import Any

from client import PlexClient, PlexOptions

DEFAULT_PRICE_IDS = (
    "019c3401-9737-7106-b3d3-b7a6e6eef0e6",
    "019c4d37-311e-7a2f-8d58-f17c39170865",
)
DEFAULT_ORDER_BOOK_ID = "019c3420-5cd7-7a88-8fe6-a5a622e01ad9"


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


def to_plex_url(base_url: str) -> str:
    url = base_url.rstrip("/")
    if url.startswith("https://"):
        url = "wss://" + url[len("https://"):]
    elif url.startswith("http://"):
        url = "ws://" + url[len("http://"):]
    return url + "/plex"


async def run() -> int:
    base_url, api_key = pick_base_url_and_key()
    if not base_url or not api_key:
        print("wsplex demo: set DORA_STAGING_BASE_URL+DORA_STAGING_API_KEY (or _PROD_ equivalents)", flush=True)
        return 2

    price_ids = split_non_empty(os.environ.get("DORA_DEMO_PRICE_ASSET_IDS", ""), DEFAULT_PRICE_IDS)
    order_book_id = os.environ.get("DORA_DEMO_ORDER_BOOK_ID") or DEFAULT_ORDER_BOOK_ID

    client = await PlexClient.connect(PlexOptions(
        url=to_plex_url(base_url),
        auth_header=f"ApiKey {api_key}",
    ))

    def on_prices(data: dict[str, Any]) -> None:
        print(f"[notif /prices] {data}", flush=True)

    def on_trades(data: dict[str, Any]) -> None:
        print(f"[notif /trades] {data}", flush=True)

    try:
        resp = await client.request("/prices", {"subscribe": price_ids}, on_notification=on_prices)
        print(f"[resp /prices subscribe] {resp}", flush=True)

        resp = await client.request(
            "/trades",
            {"subscribe": [{"order_book_ids": [order_book_id]}]},
            on_notification=on_trades,
        )
        print(f"[resp /trades subscribe] {resp}", flush=True)

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

        resp = await client.request("/prices", {"unsubscribe": price_ids})
        print(f"[resp /prices unsubscribe] {resp}", flush=True)

        resp = await client.request(
            "/trades",
            {"unsubscribe": [{"order_book_ids": [order_book_id]}]},
        )
        print(f"[resp /trades unsubscribe] {resp}", flush=True)
    finally:
        await client.close()

    return 0


if __name__ == "__main__":
    raise SystemExit(asyncio.run(run()))
