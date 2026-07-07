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

func TestGroupComponentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.GroupComponent(nil)
		if ent == nil {
			t.Fatal("expected non-nil GroupComponentEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := group_componentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "group_component." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_GROUP_COMPONENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		groupComponentRef01Ent := client.GroupComponent(nil)
		groupComponentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "group_component"}, setup.data), "group_component_ref01"))
		groupComponentRef01Data["page_id"] = setup.idmap["page01"]

		groupComponentRef01DataResult, err := groupComponentRef01Ent.Create(groupComponentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		groupComponentRef01Data = core.ToMapAny(groupComponentRef01DataResult)
		if groupComponentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if groupComponentRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		groupComponentRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		groupComponentRef01ListResult, err := groupComponentRef01Ent.List(groupComponentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		groupComponentRef01List, groupComponentRef01ListOk := groupComponentRef01ListResult.([]any)
		if !groupComponentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", groupComponentRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(groupComponentRef01List), map[string]any{"id": groupComponentRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		groupComponentRef01DataUp0Up := map[string]any{
			"id": groupComponentRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		groupComponentRef01MarkdefUp0Name := "component"
		groupComponentRef01MarkdefUp0Value := fmt.Sprintf("Mark01-group_component_ref01_%d", setup.now)
		groupComponentRef01DataUp0Up[groupComponentRef01MarkdefUp0Name] = groupComponentRef01MarkdefUp0Value

		groupComponentRef01ResdataUp0Result, err := groupComponentRef01Ent.Update(groupComponentRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		groupComponentRef01ResdataUp0 := core.ToMapAny(groupComponentRef01ResdataUp0Result)
		if groupComponentRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if groupComponentRef01ResdataUp0["id"] != groupComponentRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if groupComponentRef01ResdataUp0[groupComponentRef01MarkdefUp0Name] != groupComponentRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", groupComponentRef01MarkdefUp0Name, groupComponentRef01ResdataUp0[groupComponentRef01MarkdefUp0Name])
		}

		// LOAD
		groupComponentRef01MatchDt0 := map[string]any{
			"id": groupComponentRef01Data["id"],
		}
		groupComponentRef01DataDt0Loaded, err := groupComponentRef01Ent.Load(groupComponentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		groupComponentRef01DataDt0LoadResult := core.ToMapAny(groupComponentRef01DataDt0Loaded)
		if groupComponentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if groupComponentRef01DataDt0LoadResult["id"] != groupComponentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		groupComponentRef01MatchRm0 := map[string]any{
			"id": groupComponentRef01Data["id"],
		}
		_, err = groupComponentRef01Ent.Remove(groupComponentRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		groupComponentRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		groupComponentRef01ListRt0Result, err := groupComponentRef01Ent.List(groupComponentRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		groupComponentRef01ListRt0, groupComponentRef01ListRt0Ok := groupComponentRef01ListRt0Result.([]any)
		if !groupComponentRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", groupComponentRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(groupComponentRef01ListRt0), map[string]any{"id": groupComponentRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func group_componentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "group_component", "GroupComponentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read group_component test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse group_component test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"group_component01", "group_component02", "group_component03", "page01", "page02", "page03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_GROUP_COMPONENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_GROUP_COMPONENT_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_GROUP_COMPONENT_ENTID"])
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
