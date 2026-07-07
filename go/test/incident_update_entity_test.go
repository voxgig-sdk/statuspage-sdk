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

func TestIncidentUpdateEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.IncidentUpdate(nil)
		if ent == nil {
			t.Fatal("expected non-nil IncidentUpdateEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := incident_updateBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"update"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "incident_update." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID JSON to run live")
			return
		}
		client := setup.client

		// Bootstrap entity data from existing test data (no create step in flow).
		incidentUpdateRef01DataRaw := vs.Items(core.ToMapAny(vs.GetPath("existing.incident_update", setup.data)))
		var incidentUpdateRef01Data map[string]any
		if len(incidentUpdateRef01DataRaw) > 0 {
			incidentUpdateRef01Data = core.ToMapAny(incidentUpdateRef01DataRaw[0][1])
		}
		// Discard guards against Go's unused-var check when the flow's steps
		// happen not to consume the bootstrap data (e.g. list-only flows).
		_ = incidentUpdateRef01Data

		// UPDATE
		incidentUpdateRef01Ent := client.IncidentUpdate(nil)
		incidentUpdateRef01DataUp0Up := map[string]any{
			"id": incidentUpdateRef01Data["id"],
			"incident_id": setup.idmap["incident_id"],
			"page_id": setup.idmap["page_id"],
		}

		incidentUpdateRef01MarkdefUp0Name := "body"
		incidentUpdateRef01MarkdefUp0Value := fmt.Sprintf("Mark01-incident_update_ref01_%d", setup.now)
		incidentUpdateRef01DataUp0Up[incidentUpdateRef01MarkdefUp0Name] = incidentUpdateRef01MarkdefUp0Value

		incidentUpdateRef01ResdataUp0Result, err := incidentUpdateRef01Ent.Update(incidentUpdateRef01DataUp0Up, nil)
		if err != nil {
			t.Fatalf("update failed: %v", err)
		}
		incidentUpdateRef01ResdataUp0 := core.ToMapAny(incidentUpdateRef01ResdataUp0Result)
		if incidentUpdateRef01ResdataUp0 == nil {
			t.Fatal("expected update result to be a map")
		}
		if incidentUpdateRef01ResdataUp0["id"] != incidentUpdateRef01DataUp0Up["id"] {
			t.Fatal("expected update result id to match")
		}
		if incidentUpdateRef01ResdataUp0[incidentUpdateRef01MarkdefUp0Name] != incidentUpdateRef01MarkdefUp0Value {
			t.Fatalf("expected %s to be updated, got %v", incidentUpdateRef01MarkdefUp0Name, incidentUpdateRef01ResdataUp0[incidentUpdateRef01MarkdefUp0Name])
		}

	})
}

func incident_updateBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "incident_update", "IncidentUpdateTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read incident_update test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse incident_update test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"incident_update01", "incident_update02", "incident_update03", "page01", "page02", "page03", "incident01", "incident02", "incident03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID"])
	if idmapResolved == nil {
		idmapResolved = core.ToMapAny(idmap)
	}
	// Add incident_id alias for update test.
	if idmapResolved["incident_id"] == nil {
		idmapResolved["incident_id"] = idmapResolved["incident01"]
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
