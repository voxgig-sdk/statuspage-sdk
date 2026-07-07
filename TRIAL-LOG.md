# Voxgig SDK Generator — Statuspage Trial Log

Trial of `@voxgig/create-sdkgen` (npm) to build a statuspage.com SDK in
TypeScript, Python, and Go. Goal: reach the unit-testable state, review the
generated artifacts and docs, and record DX issues and bugs along the way.

- **Date:** 2026-07-07
- **Environment:** macOS (darwin 25.2.0), Node v24.16.0, npm 11.13.0, Python 3.14.6, go 1.26.3 (arm64)
- **Working repo:** `~/Projects/voxgig-sdk/statuspage-sdk` (fresh repo, initial commit only)
- **Local tool clones:** `~/Projects/voxgig/{create-sdkgen, sdkgen, apidef, apidef-validate, model, struct, util}` (available for fixes/linking)

Legend: 🟢 progress · 🟡 DX friction · 🔴 bug/defect · 📝 observation

---

## Phase 0 — Discovery & setup

### 🟢 Finding the usage instructions
- `create-sdkgen` README on GitHub/npm is a single line ("Create a Voxgig SDK
  Generator project") — no usage at all. The real entry points found by digging:
  - `create-sdkgen --help` text (in `bin/create-sdkgen`).
  - `sdkgen` README "Quick start".
  - **`create-sdkgen/AGENTS.md`** (GitHub `main`) — the actual end-to-end
    spec → scaffold → generate → test guide. This is good and dense.
- 🟡 **DX: README is a dead end.** A first-time visitor to
  npmjs.com/package/@voxgig/create-sdkgen or the GitHub repo gets no usage,
  no link to AGENTS.md, no example. The one-line README should at minimum
  carry the Quick start block from sdkgen's README and a link to AGENTS.md.
- 🟡 **DX: docs entry points are scattered.** Usage lives in sdkgen/README
  ("Quick start"), sdkgen/docs/tutorial.md, and create-sdkgen/AGENTS.md, with
  slightly different command sequences (see next item).
- 🟡 **DX: inconsistent generate command across docs.** sdkgen README says
  `npm run build && npm run generate` from `.sdk/`; AGENTS.md says
  `npx voxgig-model model/sdk.jsonic`. The template `.sdk/package.json` maps
  `generate` → `voxgig-model model/sdk.jsonic`, so they converge, but the
  README's extra `npm run build` step (and whether it's needed) is unexplained.
- 🔴 **Bug (minor): `--target`/`-t` missing from `--help`.** The bin parses
  `-t/--target` (repeatable; extra positionals are also treated as targets)
  and AGENTS.md documents it, but the help text omits it entirely.
- 📝 The local `~/Projects/voxgig/create-sdkgen` clone has no `AGENTS.md`
  though GitHub `main` has one — local clone slightly out of sync with
  origin. Used the GitHub version.
- 📝 AGENTS.md says `-t ts,py,go` (comma-separated); the bin actually parses
  repeated `-t` flags / positionals. Not a bug — scaffold joins them and
  `voxgig-sdkgen target add` splits on commas — but the two syntaxes are only
  discoverable by reading source.

### 🟢 Obtaining the Statuspage OpenAPI definition
- statuspage.com (Atlassian) does not publish a raw spec URL.
  `https://developer.statuspage.io/swagger.json` → S3 404. The docs site is a
  pre-rendered ReDoc page with the whole spec embedded in a
  `__redoc_state` JS blob.
- Extracted the spec with a balanced-brace scanner over the HTML →
  `statuspage-openapi3.json`: **OpenAPI 3.0.0, "Statuspage API" v1.0.0,
  54 paths, 74 schemas.** (Not a voxgig issue — Atlassian's choice — but it
  is the input provenance for this trial.)

---
