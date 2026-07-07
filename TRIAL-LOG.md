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

## Phase 3 — Unit tests (the "unit testable state")

| Target | Command | Result |
| --- | --- | --- |
| **py** | `python3 -m venv .venv && pip install pytest requests; pytest test/` | 🟢 **191 passed, 1 skipped** (mypy gate skips when mypy absent) |
| **ts** | `npm install && npm run build && npm test` | 🔶 **198/200 pass; 2 fail** — both doc-example completeness gates (35 snippet type errors) |
| **go** | `go test ./...` | 🔶 all pass except `TestReadmeGoSnippets` (2 snippets don't compile) |

The failures are all in the **doc-example gates** — the generator's own
"every documented example is tested" guarantee catching real generator
defects. The runtime/unit layers (entity ops, pipeline, struct utilities,
mock transport) are green across all three languages: ~190 tests each.

- 🔴 **BUG 4 (sdkgen ts examples): generated examples don't satisfy the
  generated types for nested entities.** Statuspage nests everything under
  `/pages/{page_id}/…`, so e.g. `ComponentLoadMatch = {id: string, page_id:
  string}` (both required). But the generated README shows
  `client.Component().load({ id: 'example_id' })` — no `page_id` → TS2345.
  Inconsistently, the **update** example DOES include both
  (`update({id, page_id})`). load/remove/create examples disagree with the
  type generator; ~22 of the 35 errors are this class, across essentially
  every nested entity (component, incident, subscriber, metric,
  group_component, page_access_*, …).
- 🔴 **BUG 5 (sdkgen ts examples): `create({})` / partial-create examples
  vs required CreateData fields** — e.g. `create({ page_id })` fails because
  `ComponentCreateData` also requires `id`. (Requiring `id` in a *create*
  payload is itself a modeling smell — see Phase 4.)
- 🔴 **BUG 6 (sdkgen ts docs/test-harness): illustration blocks type-checked
  as code.** ~13 errors (TS1005/TS1109/TS2693) come from field-type tables /
  `{...}`-placeholder blocks fenced as ` ```ts ` (e.g. a block using bare
  `boolean` / `number` / `any` as values). The readme test has a "known
  illustration" exemption, but these blocks don't qualify for it — either
  the generator should fence them as non-`ts` or emit them in the
  illustration-recognizable shape.
- 🔴 **BUG 7 (sdkgen go examples): two snippets fail `go vet`-level checks**
  — `declared and not used: err` (snip98) and unused variable
  `status_embed_config` (snip99) in README/REFERENCE snippets.
- 🟡 **DX: cross-language test rigor is uneven, and it matters.** Python's
  doc gate *executes* snippets (nice design: seeded offline mode, programming
  errors fail, only 404-class tolerated) but its static gate silently skips
  when mypy isn't installed — so py/README.md contains the **identical**
  `load({"id"})`-without-`page_id` flaw as TS and still reports 191-green.
  The TS suite catches statically what py can't. Consider shipping mypy in
  the py dev extras (or failing, not skipping, when it's absent).
- 📝 Diagnosis of bugs 4–7 root causes fanned out to a 4-agent workflow
  (ts snippets, go snippets, sdkgen example-generation internals, apidef
  modeling audit) — results in Phase 4.

---

## Phase 4 — Root-cause diagnosis (4-agent workflow) and fixes

Corrected counts: **52** failing TS snippets (17 README + 35 REFERENCE), not
35; all TS2345 argument mismatches. The Go failure is not "2 bad snippets":
**142** unused-variable errors across ~70 REFERENCE blocks are silently
auto-repaired by the test's fix-loop, which falls exactly one round short
(Go caps at 10 reported errors/round; the 15-attempt budget < 142/10 + final
rebuild). The two names that surface (snip98/99) are just the
lexicographically-last fragments.

### Root causes (sdkgen)

- 🔴 **BUG 4 root cause — two generators, two sources of truth.** The
  Match/Data *types* are generated from `opRequestShape()` (op params across
  points), but the README/REFERENCE *examples* are generated from only the
  entity's id field (`entityIdField`). Every entity nested under
  `/pages/{page_id}/…` (14 of 18 here) gets examples that can't type-check.
  The update example agreed with its type **only by accident** (its "sample
  patch fields" happened to include `page_id`).
- 🔴 **BUG 4a — dead code.** `ReadmeQuick_ts.ts` has a nested-entity branch
  that would have added the parent param — it checks `e.ancestors`, but the
  model stores `relations.ancestors`, so it can never fire. (The repo's own
  `buildIdNames` helper reads the correct path.)
- 🔴 **BUG 5 root cause — union-required op shapes.** `opShape.ts#opParams`
  merges params across ALL points of an op, first-seen `reqd` wins — so
  params required by *any* route become required for the whole op, and
  `$action` sub-resource routes (POST `…/components/{id}/page_access_groups`
  folded into `create`) inject a required `id` into `CreateData`. Worse than
  cosmetic: the union types steer callers into passing key sets that flip
  the runtime point-dispatch onto the wrong route.
- 🔴 **BUG 7 root cause — Go examples aren't idiomatic.** REFERENCE op
  examples emit `result, err := …` and never consume either (and accessor
  examples bind snake_case variables like `status_embed_config`). The
  auto-repair gate hides 140 of 142 of these.

### Fixes applied (local clones, all sdkgen tests green: 77/77)

| Fix | Where |
| --- | --- |
| `opParams`: exclude `$action` points from canonical shapes; requiredness = intersection across alternative points (+3 new unit tests pinning this) | `sdkgen/ts/src/helpers/opShape.ts`, `test/opshape.test.ts` |
| Example args derived from `opRequestShape` — the SAME source as the types (load/remove/update/create; id-first ordering; required-id no longer dropped from create) | `sdkgen/…/cmp/ts/ReadmeQuick_ts.ts`, `ReadmeRef_ts.ts`, `ReadmeEntity_ts.ts` |
| Dead nested-entity branch fixed (`relations.ancestors`), selection requires an active load + required parent param; prose names the actual parent | `ReadmeQuick_ts.ts` |
| Same policy mirrored to Python + Go components; Go examples made idiomatic (`if err != nil { panic(err) }` + consume values, camelCase vars); Go repair loop: `-gcflags=-e` + progress-based bound | `cmp/py/*`, `cmp/go/*` (agents) |
| `generate` script builds `.sdk` sources first (`tsc --build src && voxgig-model …`) so the AGENTS.md sequence works as documented | `create-sdkgen` template + this project |

After the `opParams` fix the statuspage shapes become sane:
`ComponentCreateData {page_id}` (no phantom `id`), `ComponentListMatch`
`{page_id}` required + group/user ids optional, `ComponentLoadMatch`
`{id, page_id}` — and examples agree with them by construction.

### Deeper findings logged for upstream (NOT fixed in this trial)

**apidef modeling (affects correctness of any nested/wrapped API):**
1. 🔴 **Wrapped request bodies never reach create/update.** Statuspage wraps
   every payload (`{"component": {...}}`). apidef records only path params on
   op points; the body schema's inner fields don't flow into `CreateData`,
   and the runtime sends `reqdata` verbatim — so a type-correct `create()`
   cannot produce a valid Statuspage API call. This is the single biggest
   blocker for using the generated SDK against the real API.
2. 🔴 **Point dispatch can silently pick the wrong route**: `select.exist`
   includes optional query params (`page`, `per_page`), so
   `list({page_id, page_access_group_id})` matches nothing and falls through
   to the *last* point (`GET /pages/{p}/components`) — wrong data, no error.
3. 🟠 `required:[...]` on a *nested* schema is misread as "this property is
   required" (`!!fielddef.required`), e.g. `Incident.incident` wrongly
   required.
4. 🟠 Field names are depluralized away from the wire format with an empty
   alias map (`components` → `component?: any[]`) — generated types
   mis-describe actual response JSON; no `orig` retained to map back.
5. 🟠 `$action` response schemas pollute entity fields (Component gains
   uptime-report fields; Subscriber gains count-by-type integers).
6. 🟠 Guide heuristic renames `page_id → id` on `page_access_*` collection
   routes via a `startsWith(parent + '_')` string coincidence — `id` means
   *page* id in list/create but *group/user* id in load/remove.
7. 🟠 One resource split across two entities: `postmortem` (GET/PUT) vs
   `incident_postmortem` (DELETE-only, empty interface) because DELETE→204
   has no response schema to key entity resolution on.
8. 🟡 `POST /subscribers/unsubscribe` (bulk delete) is modeled as
   `subscriber.create($action)` — a create() that deletes; deep-nested
   `resend_confirmation` mints a fake field-less `incident_subscriber`
   entity.
9. 🟡 `GET /pages/{p}/metrics` says "Get a list" but declares a non-array
   response, so it's classified `load` not `list` (spec quirk; heuristic
   could use plural-segment/paging-param counter-signals).
10. 🟡 Census: all 54 paths/112 methods map 1:1 (nothing dropped — good);
    `Permission` keeps an un-unwrapped `data` envelope; model `patch` ops
    generate no method at all (modeled but unreachable).
11. 🟡 gofmt non-compliance across 66 generated Go files (`go vet` is clean);
    a `format.Source` post-pass would fix it.

**Test-harness observations:**
- The TS illustration allowlist (`?:`, `/*`, type-table shape) is
  content-sniffed; it correctly exempted all 16 pseudo-code blocks but also
  *masks* 6 create examples that carry the real missing-param bug (they ride
  the `/* type */` placeholders). An explicit marker would be more honest.
- The Go auto-repair loop (`_ = name` insertion) actively hides unidiomatic
  examples — 140 of 142 defects invisible until the budget ran out.
- The Python execute-gate tolerance for 404-class errors is what lets
  missing-required-param examples pass; with mypy present it would have
  caught what TS caught.

---
