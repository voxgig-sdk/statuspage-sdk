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

## Phase 1 — Scaffold (`npm create @voxgig/sdkgen@latest`)

Command (from the parent dir, targeting the pre-existing repo folder):

```
npm create @voxgig/sdkgen@latest -- statuspage \
  -d statuspage-sdk/statuspage-openapi3.json -o statuspage-sdk \
  -t ts -t py -t go -f test
```

- 🟢 Installed `@voxgig/create-sdkgen@0.16.0` from npm; scaffold + `.sdk`
  npm install + target/feature add completed in ~6s with clean, readable
  structured logging. The spec file was copied to `.sdk/def/` as promised.
  Merging into an existing non-empty git repo (README, LICENSE, .gitignore
  already present) worked without complaint.
- 🔴 **BUG 1 (sdkgen 1.3.1): `target add ts,py,go` registers only the LAST
  target in `target-index.jsonic`.** All three target model files
  (`ts.jsonic`, `py.jsonic`, `go.jsonic`) are written, but the index ends up
  containing only `@"go.jsonic"` — a subsequent generate silently produces a
  **Go-only SDK**. Root cause: in `sdkgen/ts/src/action/target.ts`,
  `TargetRoot` renders `File('target-index.jsonic')` once **per target**
  inside the `each(targets)` loop, passing `names: [tname]` against the
  index content loaded once up-front — so each render overwrites the
  previous and the last one wins. Regression introduced in commit
  `1d1af3c external-targets` (changed `names: targets` → `names: [tname]`;
  the old line is still there, commented out). `feature add` does not have
  this bug (it passes the full `names: features` list each render).
  - **Workaround:** re-run adds one at a time — `npm run add-target ts`,
    then `npm run add-target py` — each single add appends correctly.
  - **Fix applied** to local clone `~/Projects/voxgig/sdkgen`
    (accumulate resolved `tname`s; pass the accumulated list to
    `UpdateIndex`, preserving alias resolution). `npm run build && npm test`:
    74/74 pass.
  - 🟡 Severity note: this failure is **silent** at scaffold time. Nothing
    warns that requested targets weren't registered; you only notice when
    `ts/`/`py/` never appear after generate. A post-add sanity check
    ("N targets requested, M registered") would catch this class.
- 📝 Scaffold-time `-f test` ran `feature add test` twice (once implicitly
  during `target add` — "the test feature is required by every target" —
  and once for the explicit flag). Idempotent, harmless, but slightly
  confusing in the logs.

## Phase 2 — Generate (`cd .sdk && npm run build && npm run generate`)

- 🔴 **BUG 2 (create-sdkgen 0.16.0): generation fails out-of-the-box —
  scaffolded `model/guide/guide.jsonic` references
  `@"@voxgig/apidef/model/guide.jsonic"`, but apidef ≥6.x ships
  `model/guide.aontu`.** First `npm run generate` dies with
  `aontu/multisource_not_found`. The template file
  `project/standard/.sdk/model/guide/guide.jsonic` in the published package
  carries the stale `.jsonic` extension. Notably the same rename was already
  fixed for `sdk.fragment.jsonic` (`apidef.aontu` / `sdkgen.aontu` refs) —
  in an **uncommitted local change** in `~/Projects/voxgig/create-sdkgen` —
  but the guide template was missed. Every fresh scaffold hits this.
  - **Fix applied:** one-char-class edit (`.jsonic` → `.aontu`) in the
    project's `.sdk/model/guide/guide.jsonic` **and** in the local
    create-sdkgen clone's template.
  - 🟡 The error report itself is decent (shows the source line and all 30
    search paths), but 30 search paths for a wrong-extension miss is noise;
    a "did you mean `guide.aontu`? (found in same folder)" hint would make
    this self-diagnosing.
- 🔴 **BUG 3 (docs): `create-sdkgen/AGENTS.md` step 3 omits the required
  build step.** It says generation is just `npx voxgig-model
  model/sdk.jsonic`, but on a fresh scaffold that fails with
  `Cannot find module '.sdk/dist/Root.js'` — the `.sdk/src/*.ts` build
  components must be compiled first (`npm run build` in `.sdk/`). sdkgen's
  README Quick start has it right (`npm run build && npm run generate`).
  Either AGENTS.md needs the build step, or (better DX) the scaffold's
  install phase / the generate script should build automatically
  (`"generate": "tsc --build src && voxgig-model model/sdk.jsonic"` or a
  prepare hook).
- 🟢 With those three fixed: **generation succeeds** for all 3 targets.
  apidef heuristics on the statuspage spec: `entities=18 paths=54
  methods=112 tags=19` (epr=0.333, emr=0.161). 18 entities generated per
  target: component, component_group_uptime, group_component, incident,
  incident_postmortem, incident_subscriber, incident_template,
  incident_update, metric, metrics_provider, page, page_access_group,
  page_access_user, permission, postmortem, status_embed_config,
  subscriber, user.
- 📝 Output per target: full runtime core (`src/` / `core/`), one entity
  module per entity, feature system, offline test suite, README.md,
  REFERENCE.md, Makefile, package metadata (`package.json` /
  `pyproject.toml` / `go.mod`). Root README.md and Makefile also generated.

---
