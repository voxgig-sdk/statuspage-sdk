package sdktest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	sdk "github.com/voxgig-sdk/statuspage-sdk/go"
	"github.com/voxgig-sdk/statuspage-sdk/go/core"

	vs "github.com/voxgig-sdk/statuspage-sdk/go/utility/struct"
)

func TestIncidentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Incident(nil)
		if ent == nil {
			t.Fatal("expected non-nil IncidentEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := incidentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "incident." + _op, _mode); _shouldSkip {
				if _reason == "" {
					_reason = "skipped via sdk-test-control.json"
				}
				t.Skip(_reason)
				return
			}
		}
		// The basic flow consumes synthetic IDs from the fixture. In live mode
		// without an *_ENTID env override, those IDs hit the live API and 4xx.
		if setup.syntheticOnly {
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		incidentRef01Ent := client.Incident(nil)
		incidentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "incident"}, setup.data), "incident_ref01"))
		incidentRef01Data["page_id"] = setup.idmap["page01"]

		incidentRef01DataResult, err := incidentRef01Ent.Create(incidentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		incidentRef01Data = core.ToMapAny(incidentRef01DataResult)
		if incidentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if incidentRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		incidentRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		incidentRef01ListResult, err := incidentRef01Ent.List(incidentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		incidentRef01List, incidentRef01ListOk := incidentRef01ListResult.([]any)
		if !incidentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", incidentRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(incidentRef01List), map[string]any{"id": incidentRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		incidentRef01DataUp0Up := map[string]any{
			"id": incidentRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		incidentRef01MarkdefUp0Name := "created_at"
		incidentRef01MarkdefUp0Value := fmt.Sprintf("Mark01-incident_ref01_%d", setup.now)
		incidentRef01DataUp0Up[incidentRef01MarkdefUp0Name] = incidentRef01MarkdefUp0Value

		incidentRef01ResdataUp0Result, err := incidentRef01Ent.Update(incidentRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		incidentRef01ResdataUp0 := core.ToMapAny(incidentRef01ResdataUp0Result)
		if incidentRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if incidentRef01ResdataUp0["id"] != incidentRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if incidentRef01ResdataUp0[incidentRef01MarkdefUp0Name] != incidentRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", incidentRef01MarkdefUp0Name, incidentRef01ResdataUp0[incidentRef01MarkdefUp0Name])
		}

		// LOAD
		incidentRef01MatchDt0 := map[string]any{
			"id": incidentRef01Data["id"],
		}
		incidentRef01DataDt0Loaded, err := incidentRef01Ent.Load(incidentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		incidentRef01DataDt0LoadResult := core.ToMapAny(incidentRef01DataDt0Loaded)
		if incidentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if incidentRef01DataDt0LoadResult["id"] != incidentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		incidentRef01MatchRm0 := map[string]any{
			"id": incidentRef01Data["id"],
		}
		_, err = incidentRef01Ent.Remove(incidentRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		incidentRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		incidentRef01ListRt0Result, err := incidentRef01Ent.List(incidentRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		incidentRef01ListRt0, incidentRef01ListRt0Ok := incidentRef01ListRt0Result.([]any)
		if !incidentRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", incidentRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(incidentRef01ListRt0), map[string]any{"id": incidentRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func incidentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "incident", "IncidentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read incident test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse incident test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"incident01", "incident02", "incident03", "page01", "page02", "page03"},
		map[string]any{
			"`$PACK`": []any{"", map[string]any{
				"`$KEY`": "`$COPY`",
				"`$VAL`": []any{"`$FORMAT`", "upper", "`$COPY`"},
			}},
		},
	)

	// Detect ENTID env override before envOverride consumes it. When live
	// mode is on without a real override, the basic test runs against synthetic
	// IDs from the fixture and 4xx's. Surface this so the test can skip.
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_INCIDENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_INCIDENT_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_INCIDENT_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add page_id alias for update test.
	if idmapResolved["page_id"] == nil {
		idmapResolved["page_id"] = idmapResolved["page01"]
	}

	if env["STATUSPAGE_TEST_LIVE"] == "TRUE" {
		mergedOpts := vs.Merge([]any{
			map[string]any{
				"apikey": env["STATUSPAGE_APIKEY"],
			},
			extra,
		})
		client = sdk.NewStatuspageSDK(core.ToMapAny(mergedOpts))
	}

	live := env["STATUSPAGE_TEST_LIVE"] == "TRUE"
	return &entityTestSetup{
		client:        client,
		data:          entityData,
		idmap:         idmapResolved,
		env:           env,
		explain:       env["STATUSPAGE_TEST_EXPLAIN"] == "TRUE",
		live:          live,
		syntheticOnly: live && !idmapOverridden,
		now:           time.Now().UnixMilli(),
	}
}
