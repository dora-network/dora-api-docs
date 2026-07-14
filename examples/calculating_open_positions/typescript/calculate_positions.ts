// Fetch the authenticated user's open positions and print one line per
// (account, asset) entry with a non-zero available or borrowed balance.
//
// Run from this folder:
//
//   npm install
//   DORA_BASE_URL=https://staging.dora.co DORA_API_KEY=... npm start
//
// Or copy ../example.env to .env and just `npm start`.
// A .env reader is included so no extra dependencies are required at runtime.

import { readFileSync, existsSync } from "node:fs";
import { join } from "node:path";

function loadDotenv(path: string): void {
  if (!existsSync(path)) return;
  // Tiny .env loader — lines like `KEY=value` set process.env. Quotes
  // and ${VAR} expansion are not supported; this is just enough for
  // DORA_BASE_URL / DORA_API_KEY in this example. For full behaviour,
  // run `npm install dotenv` and replace this with
  // `import 'dotenv/config';`.
  for (const raw of readFileSync(path, "utf8").split("\n")) {
    const line = raw.trim();
    if (!line || line.startsWith("#") || !line.includes("=")) continue;
    const [k, ...rest] = line.split("=");
    if (!process.env[k.trim()]) process.env[k.trim()] = rest.join("=").trim();
  }
}

interface AccountV2 {
  id: string;
  asset_id: string;
  is_global: boolean;
  available: string;
  borrowed: string;
}

interface Portfolio {
  accounts: Record<string, Record<string, AccountV2>>;
}

interface Response {
  data: { portfolio: Portfolio };
}

async function main(): Promise<number> {
  loadDotenv(join(import.meta.dirname ?? ".", ".env"));
  const base = (process.env.DORA_BASE_URL ?? "").replace(/\/$/, "");
  const key = process.env.DORA_API_KEY ?? "";
  if (!base || !key) {
    console.error("set DORA_BASE_URL and DORA_API_KEY (or copy ../example.env to .env)");
    return 2;
  }

  const res = await fetch(`${base}/v2/ledger/accounts/self`, {
    headers: {
      Authorization: `ApiKey ${key}`,
      "User-Agent": "DoraAPIClient/1.0 (calculating_open_positions)",
    },
  });
  if (!res.ok) {
    console.error(`HTTP ${res.status}: ${await res.text()}`);
    return 1;
  }
  const data = (await res.json()) as Response;
  const accounts = data.data.portfolio.accounts;

  console.log(
    `${"account_id".padEnd(38)}  ${"asset_id".padEnd(14)}  ${"is_global".padEnd(10)}  ${"available".padStart(16)}  ${"borrowed".padStart(16)}`,
  );
  console.log("-".repeat(100));

  let rows = 0;
  for (const byAsset of Object.values(accounts)) {
    for (const a of Object.values(byAsset)) {
      const avail = Number(a.available);
      const borrowed = Number(a.borrowed);
      if (avail === 0 && borrowed === 0) continue;
      rows++;
      console.log(
        `${a.id.padEnd(38)}  ${(a.asset_id.slice(0, 8) + "…").padEnd(14)}  ${String(a.is_global).padEnd(10)}  ${a.available.padStart(16)}  ${a.borrowed.padStart(16)}`,
      );
    }
  }
  if (rows === 0) console.log("No open positions.");
  return 0;
}

main().then((code) => process.exit(code));
