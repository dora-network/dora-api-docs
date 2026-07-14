"""Fetch the authenticated user's open positions and print one line per
(account, asset) entry with a non-zero available or borrowed balance.

Run from this folder:

    DORA_BASE_URL=https://staging.dora.co DORA_API_KEY=... python3 calculate_positions.py

Or copy ../example.env to .env and just `python3 calculate_positions.py`.
A .env reader is included below so no third-party packages are required.
"""

from __future__ import annotations

import json
import os
import sys
import urllib.error
import urllib.request
from pathlib import Path


def load_dotenv(path: Path) -> None:
    """Minimal .env loader. Lines like `KEY=value` populate os.environ.
    Quotes and `${VAR}` expansion are not supported; this is just enough
    for DORA_BASE_URL / DORA_API_KEY in this example. To get the full
    behaviour, `pip install python-dotenv` and replace this with
    `from dotenv import load_dotenv; load_dotenv()`.
    """
    if not path.exists():
        return
    for raw in path.read_text().splitlines():
        line = raw.strip()
        if not line or line.startswith("#") or "=" not in line:
            continue
        key, _, value = line.partition("=")
        os.environ.setdefault(key.strip(), value.strip())


def main() -> int:
    load_dotenv(Path(__file__).parent / ".env")
    base = os.environ.get("DORA_BASE_URL", "").rstrip("/")
    key = os.environ.get("DORA_API_KEY", "")
    if not base or not key:
        print(
            "set DORA_BASE_URL and DORA_API_KEY (or copy ../example.env to .env)",
            file=sys.stderr,
        )
        return 2

    req = urllib.request.Request(
        f"{base}/v2/ledger/accounts/self",
        headers={
            "Authorization": f"ApiKey {key}",
            "User-Agent": "DoraAPIClient/1.0 (calculating_open_positions)",
        },
    )
    try:
        with urllib.request.urlopen(req) as resp:
            body = resp.read()
            if resp.status != 200:
                print(f"HTTP {resp.status}: {body.decode()}", file=sys.stderr)
                return 1
    except urllib.error.HTTPError as e:
        print(f"HTTP {e.code}: {e.read().decode()}", file=sys.stderr)
        return 1

    data = json.loads(body)
    accounts = data["data"]["portfolio"]["accounts"]

    print(f"{'account_id':38}  {'asset_id':14}  {'is_global':10}  {'available':>16}  {'borrowed':>16}")
    print("-" * 100)
    row_count = 0
    for _acct_id, by_asset in accounts.items():
        for _asset_id, a in by_asset.items():
            try:
                avail = float(a["available"])
            except (TypeError, ValueError):
                avail = 0.0
            try:
                borrowed = float(a["borrowed"])
            except (TypeError, ValueError):
                borrowed = 0.0
            if avail == 0 and borrowed == 0:
                continue
            row_count += 1
            short_asset = a["asset_id"][:8] + "…"
            print(
                f"{a['id']:38}  {short_asset:14}  {str(a['is_global']):10}  "
                f"{a['available']:>16}  {a['borrowed']:>16}"
            )
    if row_count == 0:
        print("No open positions.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
