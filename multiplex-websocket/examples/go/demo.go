package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	plex "github.com/dora-network/dora-api-docs/multiplex-websocket/examples/go/plex"
)

func main() {
	baseURL, apiKey := pickBaseURLAndKey()
	if baseURL == "" || apiKey == "" {
		fmt.Fprintln(os.Stderr, "wsplex demo: set DORA_STAGING_BASE_URL+DORA_STAGING_API_KEY (or _PROD_ equivalents)")
		os.Exit(2)
	}
	priceIDs := splitNonEmpty(os.Getenv("DORA_DEMO_PRICE_ASSET_IDS"),
		"019c3401-9737-7106-b3d3-b7a6e6eef0e6", "019c4d37-311e-7a2f-8d58-f17c39170865")
	orderBookID := firstNonEmpty(os.Getenv("DORA_DEMO_ORDER_BOOK_ID"),
		"019c3420-5cd7-7a88-8fe6-a5a622e01ad9")
	userID := firstNonEmpty(os.Getenv("DORA_DEMO_USER_ID"),
		"019c4d37-311e-7a2f-8d58-f17c39170865")
	assetID := firstNonEmpty(os.Getenv("DORA_DEMO_ASSET_ID"),
		"019c3401-9737-7106-b3d3-b7a6e6eef0e6")

	url := strings.TrimRight(baseURL, "/")
	url = strings.Replace(url, "https://", "wss://", 1)
	url = strings.Replace(url, "http://", "ws://", 1)
	url = url + "/plex"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	client, err := plex.Connect(ctx, plex.Options{URL: url, AuthHeader: "ApiKey " + apiKey})
	if err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: connect:", err)
		os.Exit(1)
	}
	defer client.Close()

	notif := func(path string) func(json.RawMessage) {
		return func(data json.RawMessage) {
			fmt.Printf("[notif %s] %s\n", path, string(data))
		}
	}

	// --- / (list routes) ---
	if resp, err := client.Request(ctx, "/", map[string]any{}, nil); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: / routes:", err)
	} else {
		fmt.Printf("[resp / routes] %s\n", string(resp))
	}

	// --- /prices (subscribe to specific asset ids) ---
	priceSub := map[string]any{"subscribe": priceIDs}
	if resp, err := client.Request(ctx, "/prices", priceSub, notif("/prices")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /prices subscribe:", err)
		os.Exit(1)
	} else {
		fmt.Printf("[resp /prices subscribe] %s\n", string(resp))
	}

	// --- /trades (subscribe to one order book, all users) ---
	tradeSub := map[string]any{"subscribe": []map[string]any{{"order_book_ids": []string{orderBookID}}}}
	if resp, err := client.Request(ctx, "/trades", tradeSub, notif("/trades")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /trades subscribe:", err)
		os.Exit(1)
	} else {
		fmt.Printf("[resp /trades subscribe] %s\n", string(resp))
	}

	// --- /assets (subscribe to full asset updates) ---
	if resp, err := client.Request(ctx, "/assets", map[string]any{"subscribe": true}, notif("/assets")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /assets subscribe:", err)
	} else {
		fmt.Printf("[resp /assets subscribe] %s\n", string(resp))
	}

	// --- /orderbook/stats (all order books) ---
	if resp, err := client.Request(ctx, "/orderbook/stats", map[string]any{"subscribe_all": true}, notif("/orderbook/stats")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /orderbook/stats subscribe:", err)
	} else {
		fmt.Printf("[resp /orderbook/stats subscribe] %s\n", string(resp))
	}

	// --- /charts/candles (1-minute candles for one order book) ---
	const oneMinuteNs = 60_000_000_000
	candleSub := map[string]any{"subscribe": map[string]any{"orderbook_ids": []string{orderBookID}, "resolution": oneMinuteNs}}
	if resp, err := client.Request(ctx, "/charts/candles", candleSub, notif("/charts/candles")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /charts/candles subscribe:", err)
	} else {
		fmt.Printf("[resp /charts/candles subscribe] %s\n", string(resp))
	}

	// --- /accounts/balance (auth required; subscribe to one user) ---
	if resp, err := client.Request(ctx, "/accounts/balance", map[string]any{"subscribe": []string{userID}}, notif("/accounts/balance")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /accounts/balance subscribe:", err)
	} else {
		fmt.Printf("[resp /accounts/balance subscribe] %s\n", string(resp))
	}

	// --- /pools/balance (all pools) ---
	if resp, err := client.Request(ctx, "/pools/balance", map[string]any{"subscribe_all": true}, notif("/pools/balance")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /pools/balance subscribe:", err)
	} else {
		fmt.Printf("[resp /pools/balance subscribe] %s\n", string(resp))
	}

	// --- /orders/byuser (auth required; one user, all order books) ---
	ordersSub := map[string]any{"user_id": userID, "subscribe_all_orderbooks": true}
	if resp, err := client.Request(ctx, "/orders/byuser", ordersSub, notif("/orders/byuser")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /orders/byuser subscribe:", err)
	} else {
		fmt.Printf("[resp /orders/byuser subscribe] %s\n", string(resp))
	}

	// --- /debug/notify (echo a ping after 100ms) ---
	debugReq := map[string]any{"delay": 100_000_000, "data": map[string]any{"ping": "pong", "asset_id": assetID}}
	if resp, err := client.Request(ctx, "/debug/notify", debugReq, notif("/debug/notify")); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /debug/notify:", err)
	} else {
		fmt.Printf("[resp /debug/notify] %s\n", string(resp))
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	timer := time.NewTimer(10 * time.Second)
	select {
	case <-stop:
	case <-timer.C:
	}

	fmt.Println("wsplex demo: shutting down")

	// Clean up subscriptions we opened.
	unsub := func(path string, data map[string]any, label string) {
		if resp, err := client.Request(ctx, path, data, nil); err != nil {
			fmt.Fprintf(os.Stderr, "wsplex demo: %s unsubscribe: %s\n", label, err)
		} else {
			fmt.Printf("[resp %s unsubscribe] %s\n", label, string(resp))
		}
	}

	unsub("/prices", map[string]any{"unsubscribe": priceIDs}, "/prices")
	unsub("/trades", map[string]any{"unsubscribe": []map[string]any{{"order_book_ids": []string{orderBookID}}}}, "/trades")
	unsub("/assets", map[string]any{"subscribe": false}, "/assets")
	unsub("/orderbook/stats", map[string]any{"unsubscribe_all": true}, "/orderbook/stats")
	unsub("/charts/candles", map[string]any{"unsubscribe": map[string]any{"all": true}}, "/charts/candles")
	unsub("/accounts/balance", map[string]any{"unsubscribe": []string{userID}}, "/accounts/balance")
	unsub("/pools/balance", map[string]any{"unsubscribe_all": true}, "/pools/balance")
	unsub("/orders/byuser", map[string]any{"user_id": userID, "unsubscribe_all_orderbooks": true}, "/orders/byuser")
}

func pickBaseURLAndKey() (string, string) {
	if v := os.Getenv("DORA_STAGING_BASE_URL"); v != "" {
		return v, os.Getenv("DORA_STAGING_API_KEY")
	}
	if v := os.Getenv("DORA_PROD_BASE_URL"); v != "" {
		return v, os.Getenv("DORA_PROD_API_KEY")
	}
	return "", ""
}

func splitNonEmpty(s string, fallback ...string) []string {
	if s == "" {
		return fallback
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	if len(out) == 0 {
		return fallback
	}
	return out
}

func firstNonEmpty(s string, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}
