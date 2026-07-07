# Voxgig SDK Generator — Statuspage Trial Log

Trial of `@voxgig/create-sdkgen` (npm) to build a statuspage.com SDK in
TypeScript, Python, and Go. Goal: reach the unit-testable state, review the
generated artifacts and docs, and record DX issues and bugs along the way.

## Executive summary

**Outcome: reached the fully-green unit-testable state** — TS 200/200,
Python 192/192 (mypy doc gate armed and passing), Go all ok — for an
18-entity SDK generated from the Statuspage OpenAPI spec (54 paths, 112
methods, all mapped 1:1 with zero drops). Getting there required **3
blocking fixes** (multi-target index regression, stale `guide.jsonic`
template ref, missing build step in the documented flow) and a **generator
patch series** in the local clones (op-shape requiredness policy + example
generators unified onto one source of truth) that turned 52 failing TS doc
snippets and a non-converging Go doc gate green *by construction*.

**Headline positives:** the offline test suites and doc-example
completeness gates are a genuinely differentiating idea (they caught real
generator bugs during this trial, exactly as designed); scaffold → generate
→ test is minutes-fast; project collateral (deploy tooling, gitignore
hygiene, version lockstep, disclaimers) is unusually thoughtful for
generated output.

**Headline risks (before testing against the real API):** the scaffold
default `Authorization: Bearer` contradicts Statuspage's documented `OAuth`
scheme (fixed in this project's model — generator should source it from the
spec's securityScheme); `list()` returns entity wrappers typed as plain
data in all three languages (silent wrong data); wrapped request bodies
(`{"component": {...}}`) don't flow into typed create/update; runtime point
dispatch can silently pick the wrong route; thrown errors embed the raw API
key (redaction is a commented-out stub); Python's wheel installs 8 generic
top-level packages; Go is not goroutine-safe and has no context/timeouts;
the generated CI's ts job is green while running zero tests. Full detail in
Phases 4 and 6; 60+ findings logged with file:line evidence.

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

## Phase 5 — Fix propagation and the fully-green state

Propagation flow (matches the documented `edit → add-target → generate`
loop): linked the local sdkgen into `.sdk` (`npm link ../../../voxgig/sdkgen/ts`),
re-ran `npm run add-target ts,py,go` — the **multi-target index bug fix
verified end-to-end** (all three registered, idempotent re-add, no dupes) —
then `npm run generate` (now builds first) and all three test suites.

**Result: everything green, including gates that were previously dormant.**

| Target | Before fixes | After fixes |
| --- | --- | --- |
| ts | 198/200 (35+17 snippet type errors) | **200/200** |
| py | 191 pass, 1 skip (mypy gate dormant) | **192/192, no skips** (mypy installed AND passing) |
| go | FAIL (readme snippets non-converged) | **ok** (0 repair rounds needed for op examples) |

Additional bugs found and fixed during propagation:
- 🔴 The dead-`ancestors` branch existed in **three more components** —
  `ReadmeTopQuick_ts/py/go` (root README quickstart) — and the ts/go
  variants *also* gated the nested entity's `load` on the example entity's
  ops (wrong entity). All fixed to the `relations.ancestors` +
  own-active-load + required-parent-param selection.
- 🔴 **Armed mypy gate immediately caught a fresh bug class:** the root
  README "test mode" block binds `component = client.Component().list()` —
  a singular variable rebound from `Component` to `list[Component]` across
  concatenated blocks (mypy assignment error), and misleading in any case.
  Fixed `ReadmeTopTest_ts/py` to pluralize list results and derive match
  args from the op shape. (This validates the "fail, don't skip, when mypy
  is absent" recommendation — the gate finds real things.)
- 🟡 Go constructor/feature doc blocks still declare an unused `client` and
  lean on the harness's `_ = client` auto-repair (kept deliberately — making
  constructor docs "self-consuming" would pollute them; one repair round
  handles it under the new progress-based loop).
- 🟡 Cosmetic, not fixed: the per-language "Use test mode" how-to section
  still names a `list()` result with a singular variable (compiles fine —
  each ts block checks separately; flagged for the maintainer).

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

## Phase 6 — Artifact & documentation review (6-agent workflow)

Six parallel reviewers (ts code, py code, go code, docs, API fidelity,
project scaffold) over the final regenerated artifacts. Condensed below,
severity-ordered; full evidence with file:line in the review output
(`tasks/wf0pcts1q.output` in the session scratchpad). Items already logged
in Phase 4 were excluded by the review brief.

### Cross-cutting (all three languages)

- 🔴 **Auth scheme: scaffold default `Bearer` vs Statuspage's documented
  `OAuth <api_key>`** — every live call would 401, and no generated doc
  even mentions the header format. **Fixed in this project**
  (`.sdk/model/config.jsonic` → `auth: prefix: 'OAuth'`, regenerated, all
  suites still green). Generator fix: source the prefix from the spec's
  `securityScheme` instead of a hardcoded scaffold default. *This is
  verify-step-zero when we get a real API key.*
- 🔴 **`list()` returns entity-wrapper instances typed as plain data.**
  `item.id === undefined`, `item.name` returns the internal entity name
  (`'component'`), real fields only via `item.data()` — silent wrong data
  that type-checks, in TS, py, and go alike (`load()`/`create()` return
  plain records; `list()` alone wraps). The wrapper's `name/id` shadowing
  API fields is the killer.
- 🔴 **Thrown errors embed the raw API key** — `CleanUtility`/`clean()` is
  a commented-out stub in all three runtimes; a consumer who logs a caught
  error (the READMEs' own recommended pattern) ships their key to their
  logs.
- 🟠 Missing required path params are sent as literal `{page_id}` URLs (no
  client-side validation; `reqd: true` is recorded in the model but nothing
  consumes it), and every method's match/data parameter is optional, so the
  required-field types don't actually force anything.
- 🟠 Wire pollution: match keys are duplicated into the query string
  (`DELETE …/subscribers/S?id=S&page_id=P`), `$action` leaks as
  `%24action=`, and create/update bodies carry `page_id`/`id`. Tolerated by
  the Rails backend today, but spurious.
- 🟠 `patch` is advertised in READMEs/entity tables but generated nowhere —
  every PATCH endpoint is modeled-but-unreachable (Statuspage's canonical
  update verb!).
- 🟡 Default User-Agent masquerades as a Mozilla browser string.

### TypeScript specifics

- 🔴 Entity methods call a nonexistent `utility.error` — any step-level
  failure crashes with `TypeError: error is not a function` instead of the
  intended typed error.
- 🔴 The documented `extend` middleware option is rejected by option
  validation (`Unexpected keys at field <root>: extend`), and the README's
  `{hooks: {...}}` feature shape would never fire anyway (hooks are read as
  flat methods).
- 🟠 `StatuspageError`/`Control` not exported from the package root (no
  `instanceof` contract); package.json lacks `files`/`exports`/`engines`;
  CJS-only, Node-only build.
- 🟢 Praised: strict option validation with actionable messages, zero
  runtime deps, `ctrl.explain` debug channel, offline test mode DX.

### Python specifics

- 🔴 **Packaging: the wheel installs 8 generic top-level packages/modules**
  (`config`, `core`, `entity`, `feature`, `features`, `utility`, …) —
  guaranteed site-packages collisions. Needs a single `statuspage_sdk/`
  namespace package (which would also fix PEP 561 delivery, currently
  broken for ~90% of the surface).
- 🔴 No HTTP timeout, no retries, no connection reuse (stdlib urllib, fresh
  connection per call) — and the declared `requests>=2.33` dependency is
  never imported (the SDK is stdlib-only; that's a selling point being
  wasted as a phantom dep).
- 🟠 Pipeline feature hooks are never fired (generator emitted placeholder
  comments instead of calls); the Log feature exists but is unreachable via
  config; flat exception design (no `status_code`, no subclasses).
- 🟡 Booleans serialize as `str(True)` → `'True'` in query params; SDK
  doesn't pass its own mypy; live-test env loading resolves `../../.env.local`
  outside the repo under the documented invocation.

### Go specifics

- 🔴 **Not goroutine-safe** — 14 confirmed data races on shared
  root-context maps under concurrent entity calls (can fatally panic).
- 🔴 **`ListTyped` silently returns all-nil structs** (total data loss on
  the flagship typed path — the depluralized field names mean the JSON
  never maps).
- 🟠 No `context.Context` anywhere; bare `http.DefaultClient` (no timeout,
  no injection); `Direct()` returns `nil` error with the real error buried
  in the map; `go.mod` says 1.20 but code uses `log/slog` (1.21+);
  un-idiomatic error type (no `Unwrap`, embeds full ctx incl. API key);
  67/100 files fail gofmt.

### Documentation

- 🟠 Zero Statuspage-specific onboarding: nothing explains how to get an
  API key, what `page_id` is or where to find it, or the 1 req/s rate limit
  — all of which the spec's `info.description` supplies verbatim. The spec
  content is *right there* and unused.
- 🟠 Docs advertise never-generated surfaces: "an interactive REPL" (root),
  "the CLI, and MCP server" (all three per-language READMEs).
- 🟠 Every field-description column in all six doc files is empty though
  the spec supplies descriptions and enums (component status enum appears
  nowhere); the entity table's "API path" column shows arbitrary
  sub-resource/action paths as canonical.
- 🟠 Broken/misleading links: `SECURITY.md` referenced but not emitted;
  "Upstream API" → Atlassian's contact form instead of
  developer.statuspage.io.
- 🟡 ts/README documents every entity twice (~800 duplicated lines);
  `remove()` return value documented three different ways.

### API fidelity (the "can we test the real thing" assessment)

- **4.5 of 5 core Statuspage workflows are expressible**, and the write
  path (create/resolve incident, set component status, subscribe) emits
  requests the real API would very likely accept — the user hand-builds the
  `{incident: {...}}` envelope (typed-Data gap forces `as any`, already
  logged) and the passthrough body does the right thing.
- 🔴 The most common **read** call is silently wrong:
  `Incident().list({page_id})` dispatches to `/incidents/upcoming` (the
  `select.exist` optional-params bug from Phase 4, now confirmed at
  runtime with traced URLs).
- Recommended live-test order: (0) verify OAuth prefix with a real key,
  (1) drive reads via `direct()` until point-dispatch is fixed, (2) writes
  via entity ops with hand-built envelopes.

### Project scaffold / CI

- 🔴 **The generated CI's ts job is green while running zero tests**: no
  build step before `npm test`, `dist-test/` is gitignored, and
  `node --test` with an empty glob exits 0. (go and py jobs verified to
  genuinely pass on fresh checkout.) Fix: run `make -C ts test` + a
  zero-tests guard.
- 🟠 `.sdk/apidef-warnings.txt` committed a hard `BUILD FAILED` (the
  guide.jsonic bug from Phase 2) as a "warning", full of machine-local
  absolute paths — and successful re-runs never clear it, so it's stale
  evidence of a fixed bug. Should be surfaced loudly at create time,
  gitignored, and refreshed per run.
- 🟠 Five permanently no-op CI jobs for targets never generated (go-cli,
  go-mcp, rb, php, lua), one leaking private-infra commentary
  (`GOSUMDB=off` for a private repo).
- 🟠 ts npm tarball contents are accidental (no `files` whitelist; 310
  files, 2.1 MB, dist/ included only via an npm ignore-fallback quirk).
- 🟡 LICENSE holder differs root vs per-language; root `make` lands on
  deploy help (no `make test`); `VERSION` silently falls back to 0.0.0;
  spec committed twice (root + `.sdk/def/`); `npm install` instead of
  `npm ci` in CI.
- 🟢 Praised: layered gitignore hygiene verified clean, `.env.local`
  secrets pattern safe, per-target deploy tooling with dry-run rehearsal
  and tag-last ordering, truthful "publish pending" packages table,
  unofficial-SDK disclaimers everywhere, README examples actually
  test-covered in all three languages.

---

## State of play & next steps

**Repo state:** scaffolded project + three regenerated SDKs, all test
suites green, auth prefix corrected to `OAuth`, spec + trial log committed.
**Local clone state (unpushed, for upstream PRs):** sdkgen — target-index
fix, opShape requiredness policy (+3 tests), 12 example/doc components
unified onto `opRequestShape`, Go doc-gate `-gcflags=-e` + progress-based
repair loop (77/77 tests green); create-sdkgen — `guide.aontu` template
fix, `generate`-builds-first script.

## Phase 7 — Toolchain hardening: correct-first-time generation (goal follow-up)

Goal: no per-project custom fixes — the toolchain must generate correctly
the first time, parametrically from the spec. Changes (all local clones):

### Parametric auth (replaces this project's OAuth override)
- **apidef** (`transform/top.ts` + `model/apidef.aontu`, mirrored): the top
  transform now resolves the spec's PRIMARY security scheme into model
  facts `main.kit.info.security = {scheme, type, in, name, prefix}`.
  Prefix rules: `http basic/bearer` → `Basic`/`Bearer`; `oauth2`/OIDC →
  `Bearer`; `apiKey` in an `Authorization` header → prefix extracted from
  the scheme/info prose (`Authorization: OAuth 89a2…` → `OAuth`) with a
  credential-shaped-tail regex (placeholders and bare keys don't match),
  falling back to `Bearer`; `apiKey` in any other header/query → `''`
  (raw credential). 10 new unit tests; apidef 380/380 green.
- **sdkgen** (`utility.ts`): new `resolveAuthPrefix(model)` —
  `config.auth.prefix` (user override) → `info.security.prefix`
  (spec-derived) → `'Bearer'`. All SEVEN language Config components
  (ts/js/py/go/rb/php/lua) now use it. This also fixed a latent
  inconsistency: py/go/rb/php/lua Configs defaulted to an EMPTY prefix
  when `config.auth.prefix` was absent (masked until now by the scaffold's
  hardcoded value). 5 new unit tests.
- **create-sdkgen** template `config.jsonic`: hardcoded `auth: prefix:
  'Bearer'` removed — replaced by a comment documenting the derivation and
  the override syntax. This project's `config.jsonic` reverted to match.
- ✅ **Verified end-to-end**: with zero project-local auth config,
  regeneration produces `prefix: 'OAuth'` in ts/py/go Configs, sourced
  from `api-info.jsonic`'s new `security` block. All suites green
  (ts 200/200, py 192/192, go ok).
- 📝 Release note: apidef (emit), sdkgen (consume), create-sdkgen
  (no hardcode) must ship together — the old sdkgen + new template would
  produce empty prefixes in py/go/rb/php/lua.

### First-run CLI + docs correctness
- 🔴 **Fixed (sdkgen): `voxgig-sdkgen target add ts py go` silently added
  only `ts`** — the cmd handler read `args[2]` alone, dropping extra
  positionals; create-sdkgen's own README quickstart used exactly that
  form. Both `target add` and `feature add` now parse every positional,
  comma- or space-separated, via a shared tested helper (`parseAddNames`).
- **create-sdkgen AGENTS.md** (checked out from origin, edited locally):
  step 3 / edit-loop / checklist now say `npm run generate` (which builds
  `.sdk` sources first) instead of the raw `npx voxgig-model` that fails
  on a fresh scaffold.
- **create-sdkgen README quickstart**: same generate fix + added the
  missing `npm run build` to the ts verify line.
- **create-sdkgen `--help`**: `-t/--target` documented (was parsed but
  absent from help).
- **CI template ts job**: builds before `npm test` (was permanently green
  running zero tests — `dist-test/` is gitignored and an empty test glob
  exits 0), with a comment explaining why.
- sdkgen docs (tutorial/how-tos) already used `npm run build && npm run
  generate` — still correct, untouched.

### All-language example mirrors
- js/rb/php/lua example components (20 files): same `opRequestShape`
  single-source-of-truth pattern as ts/py/go — full required match args,
  create keeps required ids, `relations.ancestors` nested branches,
  plural list variables, per-language literal helpers extended with
  placeholders. Notable extra fix: **lua create examples used `field =
  nil` placeholders, which a Lua table literal silently DROPS** — the
  executed doc example documented required fields while sending an empty
  payload; now real typed literals.
- sdkgen suite after all mirrors: **86/86 green**; create-sdkgen 12/12;
  apidef 380/380.

### Verification 1 — Ruby added to this project (first-time green)
- `npm run add-target ts py go rb` (space form — exercising the CLI fix):
  all four registered, index idempotent. `npm run generate`: clean.
- rb SDK: parametric `"prefix" => "OAuth"` ✓; full-match doc examples ✓.
- 🔴 **Found+fixed en route: the rb Makefile template's `test:` ran ONLY
  `test/exists_test.rb`** (1 test) while 40 test files exist — the same
  "green while testing nothing" class as the ts CI job. Template now runs
  the whole suite via a glob require; AGENTS.md's rb line (`ruby -Itest
  test/*_test.rb`, which executes only the first file) corrected to
  `make test`.
- **Full rb suite: 178 runs, 552 assertions, 0 failures** — including the
  readme-examples gate executing all 130 documented rb examples (3 root +
  48 README + 79 REFERENCE). ts/py/go regression: 200/200, 192/192, ok.

### Verification 2 — fresh scaffold, local toolchain end-to-end
Scaffolded a brand-new project in a scratch dir with the local
create-sdkgen, linked local sdkgen+apidef, `add-target ts rb` (space
form), `npm run generate`:
- Zero errors first-time (no guide ref failure, no missing-build failure).
- `Config.ts` / `config.rb`: `prefix: 'OAuth'` — **parametric from the
  spec with an untouched scaffold config**.
- Fresh ts suite: **200/200**; fresh rb suite: **178/178**. First time.

**For the live-API phase:** get a Statuspage API key + a test page;
step zero is verifying the `OAuth` header against a real endpoint; then
reads via `direct()` (until point dispatch is fixed) and writes via entity
ops with hand-built `{entity: {...}}` envelopes. The generator-side
blockers to fix first for a real SDK release: securityScheme-sourced auth,
list() wrapper/typing, request-body envelopes into CreateData + runtime
wrapping, point-dispatch `select.exist`, error redaction, py packaging
namespace, go concurrency/context.
