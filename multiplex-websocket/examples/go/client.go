// Package plex provides a small reusable client for the DORA wsplex
// multiplexed WebSocket protocol, plus a runnable demo (see demo.go).
//
// The Client is intentionally minimal: connect, send one request at a time,
// dispatch the matching response, and let callers handle notifications by
// path. There is no automatic reconnect, no subscription registry, and no
// retry logic. Integrators compose higher-level behavior on top.
package plex
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/coder/websocket"
	"github.com/google/uuid"
)

// Options configures a Client.
type Options struct {
	URL        string // e.g. "wss://staging.dora.co/plex"
	AuthHeader string // e.g. "ApiKey xxx" — sent verbatim as the Authorization header
}

// Client is a multiplexed WebSocket client for wsplex.
type Client struct {
	conn *websocket.Conn

	mu      sync.Mutex
	pending map[string]chan envelope          // request id -> response channel
	notifs  map[string]func(json.RawMessage)  // path -> notification handler
	closed  bool
	closeCh chan struct{}
	readErr error
}

// Connect dials the multiplexed WebSocket and starts the read loop.
// Caller must call Close when done.
func Connect(ctx context.Context, opts Options) (*Client, error) {
	hdr := http.Header{}
	hdr.Set("Authorization", opts.AuthHeader)
	conn, _, err := websocket.Dial(ctx, opts.URL, &websocket.DialOptions{HTTPHeader: hdr})
	if err != nil {
		return nil, fmt.Errorf("wsplex: dial: %w", err)
	}
	c := &Client{
		conn:    conn,
		pending: make(map[string]chan envelope),
		notifs:  make(map[string]func(json.RawMessage)),
		closeCh: make(chan struct{}),
	}
	go c.readLoop()
	return c, nil
}

// Request sends one multiplexed request and blocks until the matching response arrives.
// `data` is required. `onNotification` (optional) becomes the active handler for
// further notifications on this path until the next Request to the same path replaces it.
func (c *Client) Request(ctx context.Context, path string, data any, onNotification func(json.RawMessage)) (json.RawMessage, error) {
	if data == nil {
		return nil, errors.New("wsplex: request data is required")
	}
	id := uuid.Must(uuid.NewV7()).String()
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("wsplex: marshal data: %w", err)
	}
	req := struct {
		ID   string          `json:"id"`
		Path string          `json:"path"`
		Data json.RawMessage `json:"data"`
	}{ID: id, Path: path, Data: payload}

	respCh := make(chan envelope, 1)
	c.mu.Lock()
	c.pending[id] = respCh
	if onNotification != nil {
		c.notifs[path] = onNotification
	}
	c.mu.Unlock()

	if err := c.conn.Write(ctx, websocket.MessageText, mustJSON(req)); err != nil {
		c.cleanup(id)
		return nil, fmt.Errorf("wsplex: write: %w", err)
	}

	select {
	case env := <-respCh:
		c.cleanup(id)
		if env.Error != "" {
			return nil, fmt.Errorf("wsplex: response error: %s", env.Error)
		}
		return env.Data, nil
	case <-ctx.Done():
		c.cleanup(id)
		return nil, ctx.Err()
	case <-c.closeCh:
		return nil, fmt.Errorf("wsplex: connection closed: %w", c.readErr)
	}
}

// Close closes the underlying WebSocket. Idempotent and safe after errors.
func (c *Client) Close() error {
	c.mu.Lock()
	if c.closed {
		c.mu.Unlock()
		return nil
	}
	c.closed = true
	close(c.closeCh)
	c.mu.Unlock()
	return c.conn.Close(websocket.StatusNormalClosure, "client closing")
}

func (c *Client) cleanup(id string) {
	c.mu.Lock()
	delete(c.pending, id)
	c.mu.Unlock()
}

type envelope struct {
	ID    string          `json:"id"`
	Kind  string          `json:"kind"`
	Path  string          `json:"path"`
	Data  json.RawMessage `json:"data,omitempty"`
	Error string          `json:"error,omitempty"`
}

func (c *Client) readLoop() {
	for {
		_, raw, err := c.conn.Read(context.Background())
		if err != nil {
			c.mu.Lock()
			c.readErr = err
			c.mu.Unlock()
			close(c.closeCh)
			return
		}
		var env envelope
		if err := json.Unmarshal(raw, &env); err != nil {
			continue
		}
		if env.Kind == "notification" {
			c.mu.Lock()
			h := c.notifs[env.Path]
			c.mu.Unlock()
			if h != nil {
				h(env.Data)
			}
			continue
		}
		c.mu.Lock()
		ch, ok := c.pending[env.ID]
		c.mu.Unlock()
		if ok {
			ch <- env
		}
	}
}

func mustJSON(v any) []byte {
	b, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return b
}
