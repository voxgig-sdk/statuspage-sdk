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

func TestPageAccessUserEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.PageAccessUser(nil)
		if ent == nil {
			t.Fatal("expected non-nil PageAccessUserEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := page_access_userBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "page_access_user." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_PAGE_ACCESS_USER_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		pageAccessUserRef01Ent := client.PageAccessUser(nil)
		pageAccessUserRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "page_access_user"}, setup.data), "page_access_user_ref01"))
		pageAccessUserRef01Data["page_id"] = setup.idmap["page01"]

		pageAccessUserRef01DataResult, err := pageAccessUserRef01Ent.Create(pageAccessUserRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		pageAccessUserRef01Data = core.ToMapAny(pageAccessUserRef01DataResult)
		if pageAccessUserRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if pageAccessUserRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		pageAccessUserRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		pageAccessUserRef01ListResult, err := pageAccessUserRef01Ent.List(pageAccessUserRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pageAccessUserRef01List, pageAccessUserRef01ListOk := pageAccessUserRef01ListResult.([]any)
		if !pageAccessUserRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", pageAccessUserRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(pageAccessUserRef01List), map[string]any{"id": pageAccessUserRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		pageAccessUserRef01DataUp0Up := map[string]any{
			"id": pageAccessUserRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		pageAccessUserRef01MarkdefUp0Name := "created_at"
		pageAccessUserRef01MarkdefUp0Value := fmt.Sprintf("Mark01-page_access_user_ref01_%d", setup.now)
		pageAccessUserRef01DataUp0Up[pageAccessUserRef01MarkdefUp0Name] = pageAccessUserRef01MarkdefUp0Value

		pageAccessUserRef01ResdataUp0Result, err := pageAccessUserRef01Ent.Update(pageAccessUserRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		pageAccessUserRef01ResdataUp0 := core.ToMapAny(pageAccessUserRef01ResdataUp0Result)
		if pageAccessUserRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if pageAccessUserRef01ResdataUp0["id"] != pageAccessUserRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if pageAccessUserRef01ResdataUp0[pageAccessUserRef01MarkdefUp0Name] != pageAccessUserRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", pageAccessUserRef01MarkdefUp0Name, pageAccessUserRef01ResdataUp0[pageAccessUserRef01MarkdefUp0Name])
		}

		// LOAD
		pageAccessUserRef01MatchDt0 := map[string]any{
			"id": pageAccessUserRef01Data["id"],
		}
		pageAccessUserRef01DataDt0Loaded, err := pageAccessUserRef01Ent.Load(pageAccessUserRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		pageAccessUserRef01DataDt0LoadResult := core.ToMapAny(pageAccessUserRef01DataDt0Loaded)
		if pageAccessUserRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if pageAccessUserRef01DataDt0LoadResult["id"] != pageAccessUserRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		pageAccessUserRef01MatchRm0 := map[string]any{
			"id": pageAccessUserRef01Data["id"],
		}
		_, err = pageAccessUserRef01Ent.Remove(pageAccessUserRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		pageAccessUserRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		pageAccessUserRef01ListRt0Result, err := pageAccessUserRef01Ent.List(pageAccessUserRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		pageAccessUserRef01ListRt0, pageAccessUserRef01ListRt0Ok := pageAccessUserRef01ListRt0Result.([]any)
		if !pageAccessUserRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", pageAccessUserRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(pageAccessUserRef01ListRt0), map[string]any{"id": pageAccessUserRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func page_access_userBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "page_access_user", "PageAccessUserTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read page_access_user test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse page_access_user test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"page_access_user01", "page_access_user02", "page_access_user03", "page01", "page02", "page03", "component01", "component02", "component03", "metric01", "metric02", "metric03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_PAGE_ACCESS_USER_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_PAGE_ACCESS_USER_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_PAGE_ACCESS_USER_ENTID"])
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
