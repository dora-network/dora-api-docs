// Small reusable client for the DORA wsplex multiplexed WebSocket protocol.
//
// Intentionally minimal: connect, send one request at a time, dispatch the
// matching response, and let callers handle notifications by path. There is
// no automatic reconnect, no subscription registry, and no retry logic.

import WebSocket from "ws";
import { uuidv7 } from "uuidv7";

export type NotificationHandler = (data: unknown) => void;

export interface PlexOptions {
  url: string; // e.g. "wss://staging.dora.co/plex"
  authHeader: string; // e.g. "ApiKey xxx" — sent verbatim
}

interface Envelope {
  id?: string;
  kind?: "response" | "notification";
  path?: string;
  data?: unknown;
  error?: string;
}

export class PlexClient {
  private readonly ws: WebSocket;
  private readonly pending = new Map<string, (env: Envelope) => void>();
  private readonly notifs = new Map<string, NotificationHandler>();
  private closed = false;

  private constructor(ws: WebSocket) {
    this.ws = ws;
    ws.on("message", (raw: WebSocket.RawData) => this.onMessage(raw));
  }

  static connect(opts: PlexOptions): Promise<PlexClient> {
    const { promise, resolve, reject } = Promise.withResolvers<PlexClient>();
    const ws = new WebSocket(opts.url, {
      headers: { Authorization: opts.authHeader },
    });
    ws.once("open", () => resolve(new PlexClient(ws)));
    ws.once("error", (err) => reject(new Error(`wsplex: dial: ${err.message}`)));
    return promise;
  }

  request<T = unknown>(
    path: string,
    data: Record<string, unknown>,
    onNotification?: NotificationHandler,
  ): Promise<T> {
    if (data === null || data === undefined) {
      return Promise.reject(new Error("wsplex: request data is required"));
    }
    const id = uuidv7();
    const payload = { id, path, data };
    const { promise, resolve, reject } = Promise.withResolvers<T>();
    this.pending.set(id, (env) => {
      if (env.error) {
        reject(new Error(`wsplex: response error: ${env.error}`));
      } else {
        resolve(env.data as T);
      }
    });
    if (onNotification) {
      this.notifs.set(path, onNotification);
    }
    this.ws.send(JSON.stringify(payload), (err) => {
      if (err) {
        this.pending.delete(id);
        reject(err);
      }
    });
    return promise;
  }

  close(): Promise<void> {
    if (this.closed) return Promise.resolve();
    this.closed = true;
    const { promise, resolve } = Promise.withResolvers<void>();
    this.ws.once("close", () => resolve());
    this.ws.close();
    return promise;
  }

  private onMessage(raw: WebSocket.RawData): void {
    let env: Envelope;
    try {
      env = JSON.parse(raw.toString()) as Envelope;
    } catch {
      return;
    }
    if (env.kind === "notification") {
      const handler = env.path ? this.notifs.get(env.path) : undefined;
      if (handler) handler(env.data);
      return;
    }
    if (!env.id) return;
    const resolver = this.pending.get(env.id);
    if (!resolver) return;
    this.pending.delete(env.id);
    resolver(env);
  }
}
