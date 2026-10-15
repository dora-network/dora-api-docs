# How collateral auto-rebalance works

This document explains the auto-rebalance service: how utilization is measured, when rebalances trigger, where the money comes from, what outcomes are possible, and what notifications you receive. If you just want to switch the feature on, see [Configure auto-rebalance for your accounts](./configure-auto-rebalance.md) and the [API reference](./auto-rebalance-api.md).

## The problem it solves

In DORA, isolated margin accounts are collateralized independently: each account's borrowing capacity depends only on the collateral held in that account. When one account becomes highly utilized (its borrowed value grows large relative to its borrow limit), it approaches operational limits and, ultimately, liquidation risk — even when you hold plenty of spare collateral in your *other* accounts or in your global cash account.

Auto-rebalance watches the accounts you opt in and, when one crosses its danger threshold, automatically transfers USD collateral from your spare capacity into that account. It keeps highly utilized accounts away from operational limits without you having to watch utilization and move collateral manually.

It is a best-effort helper: it can only move USD that is actually transferable, and it is not a guarantee against liquidation.

## Utilization

Auto-rebalance decisions are driven by the account's **borrow-limit utilization**:

```
utilization = borrowed_value / borrow_limit
```

Both values are computed with the same valuation rules DORA uses for account health:

- `borrowed_value` — borrowed assets priced into USD, plus negative leverage accrued interest and negative pending coupon obligations
- `borrow_limit` — collateral valued in USD and weighted by its `collateral_weight`, plus positive accrued interest valued with the USD collateral weight
- USD is valued at 1; assets without a market price contribute zero

If an account's `borrow_limit` is zero or negative, utilization is undefined (`null` in the status API) and the account is skipped — it cannot be a target or a donor until its borrow limit becomes positive.

## Thresholds

Each opted-in account stores two thresholds:

| Field | Meaning |
|---|---|
| `max_utilization` | The **trigger line**. When utilization rises above this value, the account becomes a rebalance target. |
| `min_utilization` | The **restore target**. Rebalances aim to bring the account back down to (at most) this utilization. It is also the **donor floor**: an account at or above its own `min_utilization` will not donate collateral to others. |

Thresholds must satisfy `0 < min_utilization < max_utilization <= 1`.

The gap between the two thresholds is the buffer the service works with: a wider gap means each rebalance moves more USD but leaves more time between rebalances.

## When a rebalance triggers

A managed account becomes a rebalance **target** when all of the following hold:

- auto-rebalance is `enabled` for the account
- the account is an isolated margin account (not global)
- `borrow_limit > 0`
- `utilization > max_utilization`
- the account is **not** currently in an active liquidation flow

Accounts in active liquidation are excluded from target selection — rebalancing into an account that liquidation is already unwinding would interfere with the liquidation process.

Accounts that are not opted in are excluded from *both* roles: they can neither receive rebalances nor donate.

## Where the USD comes from

Funding sources are used in a fixed order:

1. **Eligible isolated donor accounts** — your other opted-in isolated accounts, used **lowest utilization first** so the most under-utilized accounts contribute first
2. **Your global cash account** — last resort, only after isolated donors are exhausted

### Isolated donor eligibility

An opted-in isolated account can donate only when:

- it is not the target account
- `borrow_limit > 0`
- its utilization is **strictly below its own `min_utilization`** — an account sitting exactly at its minimum is never drained past it

The donor's maximum removable amount is calculated so that, after donating, its utilization lands at (but does not exceed) its own `min_utilization`. In practice:

```
donor_borrow_limit_floor = donor_borrowed_value / donor_min_utilization
removable_usd            = (donor_borrow_limit - donor_borrow_limit_floor) / usd_collateral_weight
```

capped by the donor's actual available USD balance and any transfer restrictions.

### Global account restrictions

The global cash account is only a last-resort source. Having USD there is not enough — auto-rebalance may only use the **cash-reserve-aware transferable USD**, exactly as the manual balance-transfer flow does. The usable amount is:

```
max(0, available_usd - committed_usd - required_usd)
```

where the subtracted terms cover pending market-buy commitments, the required cash reserve, pending withdrawals, settlements, and other obligations. Auto-rebalance never consumes USD that is already spoken for.

## How much is moved

The service computes the USD needed to bring the target back to its `min_utilization` — never more:

```
required_borrow_limit = borrowed_value / min_utilization
required_usd          = (required_borrow_limit - borrow_limit) / usd_collateral_weight
```

Donors and the global account are consumed in order until the requirement is met or all sources are exhausted. No excess collateral is ever moved beyond what the target needs to reach its minimum threshold.

### Worked example

Account **Bond A** (opted in, `min = 0.60`, `max = 0.80`) has borrowed $92 against a borrow limit of $100:

- utilization = `92 / 100 = 0.92` → above `max = 0.80`, so Bond A becomes a target
- required borrow limit = `92 / 0.60 ≈ 153.33`, deficit = `153.33 − 100 = 53.33` → **$53.33 of USD needed**

Account **Bond B** (opted in, `min = 0.60`) has borrowed $10 against a borrow limit of $100:

- utilization = `0.10` → below its min, so Bond B is an eligible donor
- its floor is `10 / 0.60 ≈ 16.67`, so it can give up to `100 − 16.67 ≈ 83.33` of borrow limit

Bond B transfers $53.33 of USD to Bond A:

- Bond A: limit rises to `≈ 153.33`, utilization falls to `0.60` → outcome **completed**
- Bond B: limit falls to `≈ 46.67`, utilization rises to `≈ 0.21` — still safely below its own min

Had Bond B only had $20 of spare USD, the transfer would have been $20 and the outcome **partial**.

## Outcomes

Every evaluation resolves to one of three outcomes, all persisted in an audit trail:

| Status | Meaning |
|---|---|
| `completed` | Enough collateral was transferred to bring the target to `min_utilization` or better |
| `partial` | One or more transfers executed, but the target remains above `min_utilization` |
| `unsuccessful` | No transfer could be executed, or execution failed |

The most recent outcome for each account is surfaced as `last_rebalance_outcome` in the [status API](./auto-rebalance-api.md). Partial and unsuccessful attempts still trigger a notification so you know the account remains at risk.

## How the service runs

A dedicated DORA-operated daemon runs the feature. You do not run anything; configuration is entirely through the API.

### Push-driven processing (primary path)

When account state changes for a managed user — collateral balances, borrow balances, or anything else that meaningfully changes utilization — the daemon re-evaluates the affected accounts. Rebalance configuration changes (enabling, disabling, or retuning an account) also re-evaluate immediately: a newly enabled account becomes eligible without any restart, and a disabled account stops acting as both target and donor.

Bursts of updates are **coalesced**: repeated triggers for the same account while it is already queued or being processed collapse into a single re-run, and processing is serialized per target account so overlapping events never produce competing transfers for the same account.

### Daily sweep and risk report (safety net)

At least once a day, the daemon sweeps all enabled accounts that are still above their `max_utilization` and runs the same rebalance evaluation and execution path as push triggers.

After the sweep's best-effort rebalance, accounts that *still* remain above their maximum threshold are included in a **daily risk report** — a start-of-day assessment email listing the unresolved accounts (with utilization and threshold) and any transfers the sweep made. The report is a separate message from transfer notifications: it tells you which accounts are still at risk and need funds added manually.

### Double-booking protection

When multiple targets rebalance concurrently, the execution layer takes per-source locks and re-checks transferable balances inside the transaction, so the same donor or global USD can never be planned into two competing transfers. If account state changed enough to invalidate a plan before execution, the attempt fails explicitly and is recorded as `unsuccessful` rather than executing a stale transfer.

## Notifications

Auto-rebalance sends integrator-compatible email notifications for every user-visible event, through the same channel used for daily settlement reports:

| Email | When |
|---|---|
| **Collateral auto-rebalance completed** | An attempt fully restored the target to its threshold |
| **Collateral auto-rebalance partially completed** | Transfers executed but the target remains above its threshold |
| **Collateral auto-rebalance unsuccessful** | An attempt was triggered but no transfer could be executed — the account remains at risk |
| **Daily collateral risk report** | Daily summary of accounts still above their maximum after the sweep |

Transfer emails include, for every transfer in the attempt: source and destination account, asset (USD), amount, and the utilization of both accounts before and after. A single summary email covers multi-transfer attempts.

## Limitations

- **USD only** — version 1 moves USD cash collateral only; non-USD collateral is never rebalanced
- **Same user only** — collateral never moves between different users
- **Opt-in only** — accounts without configuration are never touched, in either direction
- **Transferability-bounded** — reserved, committed, or otherwise restricted USD is never used, which means a rebalance can be partial or impossible even when balances look sufficient on paper
- **No liquidation guarantee** — this is a risk-management assist, not insurance. Accounts above their thresholds can still be liquidated before, during, or after a rebalance attempt
