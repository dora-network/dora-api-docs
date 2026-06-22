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

	notifPrices := func(data json.RawMessage) {
		fmt.Printf("[notif /prices] %s\n", string(data))
	}
	notifTrades := func(data json.RawMessage) {
		fmt.Printf("[notif /trades] %s\n", string(data))
	}

	priceSub := map[string]any{"subscribe": priceIDs}
	if resp, err := client.Request(ctx, "/prices", priceSub, notifPrices); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /prices subscribe:", err)
		os.Exit(1)
	} else {
		fmt.Printf("[resp /prices subscribe] %s\n", string(resp))
	}

	tradeSub := map[string]any{"subscribe": []map[string]any{{"order_book_ids": []string{orderBookID}}}}
	if resp, err := client.Request(ctx, "/trades", tradeSub, notifTrades); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /trades subscribe:", err)
		os.Exit(1)
	} else {
		fmt.Printf("[resp /trades subscribe] %s\n", string(resp))
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	timer := time.NewTimer(10 * time.Second)
	select {
	case <-stop:
	case <-timer.C:
	}

	fmt.Println("wsplex demo: shutting down")

	if resp, err := client.Request(ctx, "/prices", map[string]any{"unsubscribe": priceIDs}, nil); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /prices unsubscribe:", err)
	} else {
		fmt.Printf("[resp /prices unsubscribe] %s\n", string(resp))
	}

	if resp, err := client.Request(ctx, "/trades", map[string]any{"unsubscribe": []map[string]any{{"order_book_ids": []string{orderBookID}}}}, nil); err != nil {
		fmt.Fprintln(os.Stderr, "wsplex demo: /trades unsubscribe:", err)
	} else {
		fmt.Printf("[resp /trades unsubscribe] %s\n", string(resp))
	}

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
