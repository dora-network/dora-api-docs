# Trade Performance Investigation

Methodology for investigating what drove a blowout or drawdown trading day.

## Core Pattern: Compare Blowout vs Baseline

When a day's P&L is an outlier (e.g. +55 vs +1 typical), compare its trade data against a nearby flat day.

### Step 1: Fetch Trades for Both Days

```bash
# Blowout day
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
  "$DORA_BASE_URL/v1/trades?user_ids=$USER_ID&order_book_ids=$OB_ID&start=2026-05-26T00:00:00Z&end=2026-05-27T00:00:00Z&page=1&limit=100"

# Baseline day
curl -L -H "Authorization: ApiKey $DORA_API_KEY" \
  "$DORA_BASE_URL/v1/trades?user_ids=$USER_ID&order_book_ids=$OB_ID&start=2026-05-25T00:00:00Z&end=2026-05-26T00:00:00Z&page=1&limit=100"
```

Pagination: iterate `page` up to 15+ (100 per page).

### Step 2: Aggregates to Compare

| Metric | What It Reveals |
|--------|----------------|
| **Avg trade quantity** | Position sizing change |
| **Max trade quantity** | Outlier trades |
| **Price range (min/max)** | Market conditions |
| **Buy/Sell balance per hour** | Delta-neutral vs directional |
| **Hourly volume trajectory** | Scaling pattern (gradual or sudden) |
| **Net taker flow** (aggressor buys vs sells) | Who's driving price |

### Step 3: Detect Regime Changes

Look for abrupt shifts in **trade size** — the most common blowout driver is a sudden position-size increase (e.g. 10x at 15:00 UTC). Cross-reference with:

- **Candle data** for the day: `resolution=1h` to see if a price catalyst aligns
- **Orderbook stats**: 24h volume, price change %, high/low
- **Hourly net position**: does the user accumulate inventory or stay flat?

### Step 4: Pitfalls

- **API P&L vs FIFO P&L**: The `/v1/pl/self` endpoint requires JWT bearer auth (not API key). FIFO simulation from raw trades will NOT match the platform's internal P&L calculation — the Dora engine accounts for spread capture, fee rebates, and settlement adjustments that trade-level FIFO can't replicate.
- **Equal buy/sell qty per hour**: Perfect balance suggests an automated market-making strategy, not directional trading. P&L comes from spread + position mark-to-market.
- **`aggressor_indicator=true` on all trades**: A taker strategy that simultaneously takes both sides of the book. Not contradictory — this is two-sided aggressive market making.

## When to Investigate

- A day's realized P&L exceeds the previous 5 days *combined*
- Win rate drops from ~70% to ~53% (edge degradation signal)
- A single orderbook accounts for 100% of trades (concentration risk)
- Unrealized P&L on open positions exceeds daily realized P&L (directional exposure growing)
