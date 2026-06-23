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
        self._lock = asyncio.Lock()

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
        fut: asyncio.Future[dict[str, Any]] = asyncio.get_running_loop().create_future()
        async with self._lock:
            if self._closed:
                raise RuntimeError("wsplex: client is closed")
            self._pending[request_id] = fut
        if on_notification is not None:
            self._notifs[path] = on_notification
        try:
            await self._conn.send(json.dumps(payload))
        except Exception:
            # If the send fails, surface that to the caller instead of hanging.
            async with self._lock:
                self._pending.pop(request_id, None)
            raise
        try:
            response = await fut
        finally:
            async with self._lock:
                self._pending.pop(request_id, None)
        if "error" in response:
            raise RuntimeError(f"wsplex: response error: {response['error']}")
        return response.get("data", {})

    async def close(self) -> None:
        async with self._lock:
            if self._closed:
                return
            self._closed = True
        try:
            await self._conn.close()
        finally:
            if self._read_task is not None:
                self._read_task.cancel()
            # Wake any callers still awaiting a response.
            async with self._lock:
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
                async with self._lock:
                    fut = self._pending.get(rid)
                    if fut is not None and not fut.done():
                        fut.set_result(msg)
                        # Drop from the map so close()/except don't double-finalize.
                        self._pending.pop(rid, None)
        except Exception as exc:
            # If the read loop dies unexpectedly, fail any callers still awaiting a
            # response so they don't hang forever. `close()` cancels them itself
            # first, so a CancelledError reaching this point is harmless.
            async with self._lock:
                if not self._closed:
                    err = ConnectionError(f"wsplex: connection lost: {exc!r}")
                    for fut in self._pending.values():
                        if not fut.done():
                            fut.set_exception(err)
                    self._pending.clear()
