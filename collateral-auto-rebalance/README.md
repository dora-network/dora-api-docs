# Collateral Auto-Rebalance

Collateral auto-rebalance is an **opt-in** feature that automatically moves **USD cash collateral** into your isolated margin accounts when their collateral utilization gets too high. When an opted-in account rises above its configured maximum utilization threshold, DORA's auto-rebalance service pulls USD from your other opted-in isolated accounts — and, as a last resort, from your global cash account — to bring the account back down to its minimum threshold.

> **IMPORTANT:** Auto-rebalance is a basic risk-management assist. It is **not a guarantee against liquidation** and does not replace user-managed risk systems. Accounts can still be liquidated while auto-rebalance is enabled, for example when no funding source has enough transferable USD.

## Documentation in this section

| Document | Purpose |
|---|---|
| [How auto-rebalance works](./how-auto-rebalance-works.md) | Explanation of the service: utilization, thresholds, trigger policy, donor selection, outcomes, notifications, and the daily sweep |
| [Configure auto-rebalance for your accounts](./configure-auto-rebalance.md) | Step-by-step guide to enabling, tuning, and disabling auto-rebalance through the API |
| [Auto-rebalance API reference](./auto-rebalance-api.md) | Endpoint, request, response, and error reference for the `/v2/accounts/auto_rebalance` routes |

## Key facts at a glance

- **Opt-in per isolated account** — accounts that are not configured never act as rebalance targets and never donate collateral
- **USD only** — version 1 moves USD cash collateral only, between accounts of the same user
- **Two thresholds per account** — `max_utilization` (trigger line) and `min_utilization` (restore target and donor floor)
- **Deterministic funding order** — eligible isolated donors first (lowest utilization first), then your global cash account
- **Audit trail and email notifications** — every attempt is persisted and reported (completed, partial, or unsuccessful), plus a daily risk report for unresolved accounts
