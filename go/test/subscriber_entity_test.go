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

func TestSubscriberEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.Subscriber(nil)
		if ent == nil {
			t.Fatal("expected non-nil SubscriberEntity")
		}
	})

	// Feature #4: the entity Stream(action, ...) method runs the op pipeline and
	// returns a channel over result items. With the streaming feature active it
	// yields the feature's incremental output; otherwise it falls back to the
	// materialised list so Stream always yields.
	t.Run("stream", func(t *testing.T) {
		seed := map[string]any{
			"entity": map[string]any{
				"subscriber": map[string]any{
					"s1": map[string]any{"id": "s1"},
					"s2": map[string]any{"id": "s2"},
					"s3": map[string]any{"id": "s3"},
				},
			},
		}

		// Fallback: streaming inactive -> yields the materialised list items.
		base := sdk.TestSDK(seed, nil)
		var seen []any
		for item := range base.Subscriber(nil).Stream("list", nil, nil) {
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
			for item := range streamSdk.Subscriber(nil).Stream("list", nil, nil) {
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
		setup := subscriberBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "subscriber." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_SUBSCRIBER_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		subscriberRef01Ent := client.Subscriber(nil)
		subscriberRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "subscriber"}, setup.data), "subscriber_ref01"))
		subscriberRef01Data["incident_id"] = setup.idmap["incident01"]
		subscriberRef01Data["page_id"] = setup.idmap["page01"]

		subscriberRef01DataResult, err := subscriberRef01Ent.Create(subscriberRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		subscriberRef01Data = core.ToMapAny(subscriberRef01DataResult)
		if subscriberRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if subscriberRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		subscriberRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		subscriberRef01ListResult, err := subscriberRef01Ent.List(subscriberRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		subscriberRef01List, subscriberRef01ListOk := subscriberRef01ListResult.([]any)
		if !subscriberRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", subscriberRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(subscriberRef01List), map[string]any{"id": subscriberRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		subscriberRef01DataUp0Up := map[string]any{
			"id": subscriberRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		subscriberRef01MarkdefUp0Name := "component"
		subscriberRef01MarkdefUp0Value := fmt.Sprintf("Mark01-subscriber_ref01_%d", setup.now)
		subscriberRef01DataUp0Up[subscriberRef01MarkdefUp0Name] = subscriberRef01MarkdefUp0Value

		subscriberRef01ResdataUp0Result, err := subscriberRef01Ent.Update(subscriberRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		subscriberRef01ResdataUp0 := core.ToMapAny(subscriberRef01ResdataUp0Result)
		if subscriberRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if subscriberRef01ResdataUp0["id"] != subscriberRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if subscriberRef01ResdataUp0[subscriberRef01MarkdefUp0Name] != subscriberRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", subscriberRef01MarkdefUp0Name, subscriberRef01ResdataUp0[subscriberRef01MarkdefUp0Name])
		}

		// LOAD
		subscriberRef01MatchDt0 := map[string]any{
			"id": subscriberRef01Data["id"],
		}
		subscriberRef01DataDt0Loaded, err := subscriberRef01Ent.Load(subscriberRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		subscriberRef01DataDt0LoadResult := core.ToMapAny(subscriberRef01DataDt0Loaded)
		if subscriberRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if subscriberRef01DataDt0LoadResult["id"] != subscriberRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		subscriberRef01MatchRm0 := map[string]any{
			"id": subscriberRef01Data["id"],
		}
		_, err = subscriberRef01Ent.Remove(subscriberRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		subscriberRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		subscriberRef01ListRt0Result, err := subscriberRef01Ent.List(subscriberRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		subscriberRef01ListRt0, subscriberRef01ListRt0Ok := subscriberRef01ListRt0Result.([]any)
		if !subscriberRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", subscriberRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(subscriberRef01ListRt0), map[string]any{"id": subscriberRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func subscriberBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "subscriber", "SubscriberTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read subscriber test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse subscriber test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"subscriber01", "subscriber02", "subscriber03", "page01", "page02", "page03", "incident01", "incident02", "incident03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_SUBSCRIBER_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_SUBSCRIBER_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_SUBSCRIBER_ENTID"])
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
