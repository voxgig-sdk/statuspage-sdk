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

func TestMetricsProviderEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.MetricsProvider(nil)
		if ent == nil {
			t.Fatal("expected non-nil MetricsProviderEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := metrics_providerBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create", "list", "update", "load", "remove"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "metrics_provider." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_METRICS_PROVIDER_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		metricsProviderRef01Ent := client.MetricsProvider(nil)
		metricsProviderRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "metrics_provider"}, setup.data), "metrics_provider_ref01"))
		metricsProviderRef01Data["page_id"] = setup.idmap["page01"]

		metricsProviderRef01DataResult, err := metricsProviderRef01Ent.Create(metricsProviderRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		metricsProviderRef01Data = core.ToMapAny(metricsProviderRef01DataResult)
		if metricsProviderRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}
		if metricsProviderRef01Data["id"] == nil {
			t.Fatal("expected created entity to have an id")
		}

		// LIST
		metricsProviderRef01Match := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		metricsProviderRef01ListResult, err := metricsProviderRef01Ent.List(metricsProviderRef01Match, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		metricsProviderRef01List, metricsProviderRef01ListOk := metricsProviderRef01ListResult.([]any)
		if !metricsProviderRef01ListOk {
			t.Fatalf("expected list result to be an array, got %T", metricsProviderRef01ListResult)
		}

		foundItem := vs.Select(entityListToData(metricsProviderRef01List), map[string]any{"id": metricsProviderRef01Data["id"]})
		if vs.IsEmpty(foundItem) {
			t.Fatal("expected to find created entity in list")
		}

		// UPDATE
		metricsProviderRef01DataUp0Up := map[string]any{
			"id": metricsProviderRef01Data["id"],
			"page_id": setup.idmap["page_id"],
		}

		metricsProviderRef01MarkdefUp0Name := "created_at"
		metricsProviderRef01MarkdefUp0Value := fmt.Sprintf("Mark01-metrics_provider_ref01_%d", setup.now)
		metricsProviderRef01DataUp0Up[metricsProviderRef01MarkdefUp0Name] = metricsProviderRef01MarkdefUp0Value

		metricsProviderRef01ResdataUp0Result, err := metricsProviderRef01Ent.Update(metricsProviderRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		metricsProviderRef01ResdataUp0 := core.ToMapAny(metricsProviderRef01ResdataUp0Result)
		if metricsProviderRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if metricsProviderRef01ResdataUp0["id"] != metricsProviderRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if metricsProviderRef01ResdataUp0[metricsProviderRef01MarkdefUp0Name] != metricsProviderRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", metricsProviderRef01MarkdefUp0Name, metricsProviderRef01ResdataUp0[metricsProviderRef01MarkdefUp0Name])
		}

		// LOAD
		metricsProviderRef01MatchDt0 := map[string]any{
			"id": metricsProviderRef01Data["id"],
		}
		metricsProviderRef01DataDt0Loaded, err := metricsProviderRef01Ent.Load(metricsProviderRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		metricsProviderRef01DataDt0LoadResult := core.ToMapAny(metricsProviderRef01DataDt0Loaded)
		if metricsProviderRef01DataDt0LoadResult == nil {
			t.Fatal("expected load result to be a map")
		}
		if metricsProviderRef01DataDt0LoadResult["id"] != metricsProviderRef01Data["id"] {
			t.Fatal("expected load result id to match")
		}

		// REMOVE
		metricsProviderRef01MatchRm0 := map[string]any{
			"id": metricsProviderRef01Data["id"],
		}
		_, err = metricsProviderRef01Ent.Remove(metricsProviderRef01MatchRm0, nil)
		if err != nil {
			t.Fatalf("remove failed: %v", err)
		}

		// LIST
		metricsProviderRef01MatchRt0 := map[string]any{
			"page_id": setup.idmap["page01"],
		}

		metricsProviderRef01ListRt0Result, err := metricsProviderRef01Ent.List(metricsProviderRef01MatchRt0, nil)
		if err != nil {
			t.Fatalf("list failed: %v", err)
		}
		metricsProviderRef01ListRt0, metricsProviderRef01ListRt0Ok := metricsProviderRef01ListRt0Result.([]any)
		if !metricsProviderRef01ListRt0Ok {
			t.Fatalf("expected list result to be an array, got %T", metricsProviderRef01ListRt0Result)
		}

		notFoundItem := vs.Select(entityListToData(metricsProviderRef01ListRt0), map[string]any{"id": metricsProviderRef01Data["id"]})
		if !vs.IsEmpty(notFoundItem) {
			t.Fatal("expected removed entity to not be in list")
		}

	})
}

func metrics_providerBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "metrics_provider", "MetricsProviderTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read metrics_provider test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse metrics_provider test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"metrics_provider01", "metrics_provider02", "metrics_provider03", "page01", "page02", "page03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_METRICS_PROVIDER_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_METRICS_PROVIDER_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_METRICS_PROVIDER_ENTID"])
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
