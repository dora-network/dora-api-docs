// calculate_open_positions fetches the authenticated user's account/position
// portfolio and prints one line per (account, asset) entry that has a
// non-zero available or borrowed balance — i.e. an open position.
//
// Run from this folder:
//
//   go mod tidy
//   DORA_BASE_URL=https://staging.dora.co DORA_API_KEY=... go run .
//
// Or copy ../../example.env to .env and just `go run .`.

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// AccountV2 mirrors the V2 account/position schema. See
// openapi.json -> components.schemas.AccountV2. Numeric fields are
// decimal strings in the response; we parse to float64 for printing.
type AccountV2 struct {
	ID               string  `json:"id"`
	AssetID          string  `json:"asset_id"`
	Seq              int     `json:"seq"`
	IsGlobal         bool    `json:"is_global"`
	Available        string  `json:"available"`
	Locked           string  `json:"locked"`
	Supplied         string  `json:"supplied"`
	Borrowed         string  `json:"borrowed"`
	ImpendingBorrows string  `json:"impending_borrows"`
	AccountName      string  `json:"account_name"`
}

type Portfolio struct {
	Accounts map[string]map[string]AccountV2 `json:"accounts"`
}

type Response struct {
	Data struct {
		Portfolio Portfolio `json:"portfolio"`
	} `json:"data"`
}

func main() {
	_ = godotenv.Load() // .env in cwd, optional
	base := strings.TrimRight(os.Getenv("DORA_BASE_URL"), "/")
	key := os.Getenv("DORA_API_KEY")
	if base == "" || key == "" {
		fmt.Fprintln(os.Stderr, "set DORA_BASE_URL and DORA_API_KEY (or copy ../example.env to .env)")
		os.Exit(2)
	}

	req, _ := http.NewRequest(http.MethodGet, base+"/v2/ledger/accounts/self", nil)
	req.Header.Set("Authorization", "ApiKey "+key)
	req.Header.Set("User-Agent", "DoraAPIClient/1.0 (calculating_open_positions)")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Fprintln(os.Stderr, "request failed:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		fmt.Fprintf(os.Stderr, "HTTP %d: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	var out Response
	if err := json.Unmarshal(body, &out); err != nil {
		fmt.Fprintln(os.Stderr, "decode failed:", err)
		os.Exit(1)
	}

	fmt.Printf("%-36s  %-12s  %-10s  %14s  %14s\n",
		"account_id", "asset_id", "is_global", "available", "borrowed")
	fmt.Println(strings.Repeat("-", 96))

	rowCount := 0
	for _, byAsset := range out.Data.Portfolio.Accounts {
		for _, a := range byAsset {
			avail := parseFloat(a.Available)
			borrowed := parseFloat(a.Borrowed)
			if avail == 0 && borrowed == 0 {
				continue
			}
			rowCount++
			fmt.Printf("%-36s  %-12s  %-10v  %14s  %14s\n",
				a.ID, a.AssetID[:8]+"…", a.IsGlobal, a.Available, a.Borrowed)
		}
	}
	if rowCount == 0 {
		fmt.Println("No open positions.")
	}
}

// parseFloat returns 0 for empty/invalid strings (DORA omits zero fields
// in some responses, and "0" parses to 0 anyway — both are fine for our
// "is anything non-zero" check).
func parseFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}
