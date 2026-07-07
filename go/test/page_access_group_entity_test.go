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

func TestPageAccessGroupEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PageAccessGroup(nil)
		if ent == nil {
			t.Fatal("expected non-nil PageAccessGroupEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := page_access_groupBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "page_access_group." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		pageAccessGroupRef01Ent := client.PageAccessGroup(nil)
		pageAccessGroupRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "page_access_group"}, setup.data), "page_access_group_ref01"))
		pageAccessGroupRef01Data["page_id"] = setup.idmap["page01"]

		pageAccessGroupRef01DataResult, err := pageAccessGroupRef01Ent.Create(pageAccessGroupRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		pageAccessGroupRef01Data = core.ToMapAny(pageAccessGroupRef01DataResult)
		if pageAccessGroupRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if pageAccessGroupRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		pageAccessGroupRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		pageAccessGroupRef01ListResult, err := pageAccessGroupRef01Ent.List(pageAccessGroupRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pageAccessGroupRef01List, pageAccessGroupRef01ListOk := pageAccessGroupRef01ListResult.([]any)
		if !pageAccessGroupRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", pageAccessGroupRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(pageAccessGroupRef01List), map[string]any{"id": pageAccessGroupRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		pageAccessGroupRef01DataUp0Up := map[string]any{
			"id": pageAccessGroupRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		pageAccessGroupRef01MarkdefUp0Name := "created_at"
		pageAccessGroupRef01MarkdefUp0Value := fmt.Sprintf("Mark01-page_access_group_ref01_%d", setup.now)
		pageAccessGroupRef01DataUp0Up[pageAccessGroupRef01MarkdefUp0Name] = pageAccessGroupRef01MarkdefUp0Value

		pageAccessGroupRef01ResdataUp0Result, err := pageAccessGroupRef01Ent.Update(pageAccessGroupRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		pageAccessGroupRef01ResdataUp0 := core.ToMapAny(pageAccessGroupRef01ResdataUp0Result)
		if pageAccessGroupRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if pageAccessGroupRef01ResdataUp0["id"] != pageAccessGroupRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if pageAccessGroupRef01ResdataUp0[pageAccessGroupRef01MarkdefUp0Name] != pageAccessGroupRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", pageAccessGroupRef01MarkdefUp0Name, pageAccessGroupRef01ResdataUp0[pageAccessGroupRef01MarkdefUp0Name])
		}

		// LOAD
		pageAccessGroupRef01MatchDt0 := map[string]any{
			"id": pageAccessGroupRef01Data["id"],
		}
		pageAccessGroupRef01DataDt0Loaded, err := pageAccessGroupRef01Ent.Load(pageAccessGroupRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		pageAccessGroupRef01DataDt0LoadResult := core.ToMapAny(pageAccessGroupRef01DataDt0Loaded)
		if pageAccessGroupRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if pageAccessGroupRef01DataDt0LoadResult["id"] != pageAccessGroupRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		pageAccessGroupRef01MatchRm0 := map[string]any{
			"id": pageAccessGroupRef01Data["id"],
		}
		_, err = pageAccessGroupRef01Ent.Remove(pageAccessGroupRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		pageAccessGroupRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		pageAccessGroupRef01ListRt0Result, err := pageAccessGroupRef01Ent.List(pageAccessGroupRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pageAccessGroupRef01ListRt0, pageAccessGroupRef01ListRt0Ok := pageAccessGroupRef01ListRt0Result.([]any)
		if !pageAccessGroupRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", pageAccessGroupRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(pageAccessGroupRef01ListRt0), map[string]any{"id": pageAccessGroupRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func page_access_groupBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "page_access_group", "PageAccessGroupTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read page_access_group test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse page_access_group test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"page_access_group01", "page_access_group02", "page_access_group03", "page01", "page02", "page03", "component01", "component02", "component03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID"])
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
