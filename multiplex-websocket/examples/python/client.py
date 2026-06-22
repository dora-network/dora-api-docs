"""Small reusable client for the DORA wsplex multiplexed WebSocket protocol.

Intentionally minimal: connect, send one request at a time, dispatch the
matching response, and let callers handle notifications by path. There is
no automatic reconnect, no subscription registry, and no retry logic.
"""

from __future__ import annotations

import asyncio
import json
from dataclasses import dataclass
from typing import Any, Awaitable, Callable, Optional
from uuid import uuid7

import websockets
from websockets.asyncio.client import ClientConnection


NotificationHandler = Callable[[dict[str, Any]], None]


@dataclass
class PlexOptions:
    url: str
    auth_header: str  # e.g. "ApiKey xxx"


class PlexClient:
    def __init__(self, conn: ClientConnection) -> None:
        self._conn = conn
        self._pending: dict[str, asyncio.Future[dict[str, Any]]] = {}
        self._notifs: dict[str, NotificationHandler] = {}
        self._read_task: Optional[asyncio.Task[None]] = None
        self._closed = False

    @classmethod
    async def connect(cls, opts: PlexOptions) -> "PlexClient":
        headers = {"Authorization": opts.auth_header}
        conn = await websockets.connect(opts.url, additional_headers=headers)
        client = cls(conn)
        client._read_task = asyncio.create_task(client._read_loop())
        return client

    async def request(
        self,
        path: str,
        data: dict[str, Any],
        on_notification: Optional[NotificationHandler] = None,
    ) -> dict[str, Any]:
        if data is None:
            raise ValueError("wsplex: request data is required")
        request_id = str(uuid7())
        payload = {"id": request_id, "path": path, "data": data}
        fut: asyncio.Future[dict[str, Any]] = asyncio.get_event_loop().create_future()
        self._pending[request_id] = fut
        if on_notification is not None:
            self._notifs[path] = on_notification
        await self._conn.send(json.dumps(payload))
        response = await fut
        self._pending.pop(request_id, None)
        if "error" in response:
            raise RuntimeError(f"wsplex: response error: {response['error']}")
        return response.get("data", {})

    async def close(self) -> None:
        if self._closed:
            return
        self._closed = True
        try:
            await self._conn.close()
        finally:
            if self._read_task is not None:
                self._read_task.cancel()
            # Wake any callers still awaiting a response.
            for fut in self._pending.values():
                if not fut.done():
                    fut.cancel()
            self._pending.clear()

    async def _read_loop(self) -> None:
        try:
            async for raw in self._conn:
                msg = json.loads(raw)
                kind = msg.get("kind")
                if kind == "notification":
                    handler = self._notifs.get(msg.get("path", ""))
                    if handler is not None:
                        handler(msg.get("data", {}))
                    continue
                rid = msg.get("id")
                if rid is None:
                    continue
                fut = self._pending.get(rid)
                if fut is not None and not fut.done():
                    fut.set_result(msg)
        except Exception:
            pass
