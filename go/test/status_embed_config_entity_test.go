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

func TestStatusEmbedConfigEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.StatusEmbedConfig(nil)
		if ent == nil {
			t.Fatal("expected non-nil StatusEmbedConfigEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := status_embed_configBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update", "load"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "status_embed_config." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		statusEmbedConfigRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.status_embed_config", setup.data)))
		var statusEmbedConfigRef01Data map[string]any
		if len(statusEmbedConfigRef01DataRaw) > 0 {
			statusEmbedConfigRef01Data = core.ToMapAny(statusEmbedConfigRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = statusEmbedConfigRef01Data

		// UPDATE
		statusEmbedConfigRef01Ent := client.StatusEmbedConfig(nil)
		statusEmbedConfigRef01DataUp0Up := map[string]any{
		}

		statusEmbedConfigRef01MarkdefUp0Name := "incident_background_color"
		statusEmbedConfigRef01MarkdefUp0Value := fmt.Sprintf("Mark01-status_embed_config_ref01_%d", setup.now)
		statusEmbedConfigRef01DataUp0Up[statusEmbedConfigRef01MarkdefUp0Name] = statusEmbedConfigRef01MarkdefUp0Value

		statusEmbedConfigRef01ResdataUp0Result, err := statusEmbedConfigRef01Ent.Update(statusEmbedConfigRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		statusEmbedConfigRef01ResdataUp0 := core.ToMapAny(statusEmbedConfigRef01ResdataUp0Result)
		if statusEmbedConfigRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if statusEmbedConfigRef01ResdataUp0[statusEmbedConfigRef01MarkdefUp0Name] != statusEmbedConfigRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", statusEmbedConfigRef01MarkdefUp0Name, statusEmbedConfigRef01ResdataUp0[statusEmbedConfigRef01MarkdefUp0Name])
		}

		// LOAD
		statusEmbedConfigRef01MatchDt0 := map[string]any{}
		statusEmbedConfigRef01DataDt0Loaded, err := statusEmbedConfigRef01Ent.Load(statusEmbedConfigRef01MatchDt0, nil)
		if err != nil {
			t.Fatalf("load failed: %v", err)
		}
		if statusEmbedConfigRef01DataDt0Loaded == nil {
			t.Fatal("expected load result to be non-nil")
		}

	})
}

func status_embed_configBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "status_embed_config", "StatusEmbedConfigTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read status_embed_config test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse status_embed_config test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"status_embed_config01", "status_embed_config02", "status_embed_config03", "page01", "page02", "page03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
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
