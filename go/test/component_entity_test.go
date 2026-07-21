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

func TestComponentEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Component(nil)
		if ent == nil {
			t.Fatal("expected non-nil ComponentEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"component": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.Component(nil).Stream("list", nil, nil) {
			seen = append(seen, item)
		}
		if len(seen) != 3 {
			t.Fatalf("expected 3 streamed items, got %d", len(seen))
		}

		// Inbound: streaming active -> yields each item from the feature iterator.
		hasStreaming := false
		if fm, ok := core.MakeConfig()["feature"].(map[string]any); ok {
			_, hasStreaming = fm["streaming"]
		}
		if hasStreaming {
			streamSdk := sdk.TestSDK(seed, map[string]any{
				"feature": map[string]any{"streaming": map[string]any{"active": true}},
			})
			var got []any
			for item := range streamSdk.Component(nil).Stream("list", nil, nil) {
				if sub, ok := item.([]any); ok {
					got = append(got, sub...)
				} else {
					got = append(got, item)
				}
			}
			if len(got) != 3 {
				t.Fatalf("expected 3 items via streaming feature, got %d", len(got))
			}
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := componentBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "component." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_COMPONENT_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		componentRef01Ent := client.Component(nil)
		componentRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "component"}, setup.data), "component_ref01"))
		componentRef01Data["page_access_group_id"] = setup.idmap["page_access_group01"]
		componentRef01Data["page_access_user_id"] = setup.idmap["page_access_user01"]
		componentRef01Data["page_id"] = setup.idmap["page01"]

		componentRef01DataResult, err := componentRef01Ent.Create(componentRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		componentRef01Data = core.ToMapAny(componentRef01DataResult)
		if componentRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if componentRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		componentRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		componentRef01ListResult, err := componentRef01Ent.List(componentRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		componentRef01List, componentRef01ListOk := componentRef01ListResult.([]any)
		if !componentRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", componentRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(componentRef01List), map[string]any{"id": componentRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		componentRef01DataUp0Up := map[string]any{
			"id": componentRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		componentRef01MarkdefUp0Name := "automation_email"
		componentRef01MarkdefUp0Value := fmt.Sprintf("Mark01-component_ref01_%d", setup.now)
		componentRef01DataUp0Up[componentRef01MarkdefUp0Name] = componentRef01MarkdefUp0Value

		componentRef01ResdataUp0Result, err := componentRef01Ent.Update(componentRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		componentRef01ResdataUp0 := core.ToMapAny(componentRef01ResdataUp0Result)
		if componentRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if componentRef01ResdataUp0["id"] != componentRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if componentRef01ResdataUp0[componentRef01MarkdefUp0Name] != componentRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", componentRef01MarkdefUp0Name, componentRef01ResdataUp0[componentRef01MarkdefUp0Name])
		}

		// LOAD
		componentRef01MatchDt0 := map[string]any{
			"id": componentRef01Data["id"],
		}
		componentRef01DataDt0Loaded, err := componentRef01Ent.Load(componentRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		componentRef01DataDt0LoadResult := core.ToMapAny(componentRef01DataDt0Loaded)
		if componentRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if componentRef01DataDt0LoadResult["id"] != componentRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		componentRef01MatchRm0 := map[string]any{
			"id": componentRef01Data["id"],
		}
		_, err = componentRef01Ent.Remove(componentRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		componentRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		componentRef01ListRt0Result, err := componentRef01Ent.List(componentRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		componentRef01ListRt0, componentRef01ListRt0Ok := componentRef01ListRt0Result.([]any)
		if !componentRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", componentRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(componentRef01ListRt0), map[string]any{"id": componentRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func componentBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "component", "ComponentTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read component test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse component test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"component01", "component02", "component03", "page01", "page02", "page03", "page_access_group01", "page_access_group02", "page_access_group03", "page_access_user01", "page_access_user02", "page_access_user03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_COMPONENT_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_COMPONENT_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_COMPONENT_ENTID"])
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
