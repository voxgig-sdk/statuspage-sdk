package sdktest

import (
	"encoding/json"
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

func TestIncidentSubscriberEntity(t *testing.T) {
	t.Run("instance", func(t *testing.T) {
		testsdk := sdk.TestSDK(nil, nil)
		ent := testsdk.IncidentSubscriber(nil)
		if ent == nil {
			t.Fatal("expected non-nil IncidentSubscriberEntity")
		}
	})

	t.Run("basic", func(t *testing.T) {
		setup := incident_subscriberBasicSetup(nil)
		// Per-op sdk-test-control.json skip — basic test exercises a flow
		// with multiple ops; skipping any op skips the whole flow.
		_mode := "unit"
		if setup.live {
			_mode = "live"
		}
		for _, _op := range []string{"create"} {
			if _shouldSkip, _reason := isControlSkipped("entityOp", "incident_subscriber." + _op, _mode); _shouldSkip {
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
			t.Skip("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID JSON to run live")
			return
		}
		client := setup.client

		// CREATE
		incidentSubscriberRef01Ent := client.IncidentSubscriber(nil)
		incidentSubscriberRef01Data := core.ToMapAny(vs.GetProp(
			vs.GetPath([]any{"new", "incident_subscriber"}, setup.data), "incident_subscriber_ref01"))
		incidentSubscriberRef01Data["incident_id"] = setup.idmap["incident01"]
		incidentSubscriberRef01Data["page_id"] = setup.idmap["page01"]
		incidentSubscriberRef01Data["subscriber_id"] = setup.idmap["subscriber01"]

		incidentSubscriberRef01DataResult, err := incidentSubscriberRef01Ent.Create(incidentSubscriberRef01Data, nil)
		if err != nil {
			t.Fatalf("create failed: %v", err)
		}
		incidentSubscriberRef01Data = core.ToMapAny(incidentSubscriberRef01DataResult)
		if incidentSubscriberRef01Data == nil {
			t.Fatal("expected create result to be a map")
		}

	})
}

func incident_subscriberBasicSetup(extra map[string]any) *entityTestSetup {
	loadEnvLocal()

	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	entityDataFile := filepath.Join(dir, "..", "..", ".sdk", "test", "entity", "incident_subscriber", "IncidentSubscriberTestData.json")

	entityDataSource, err := os.ReadFile(entityDataFile)
	if err != nil {
		panic("failed to read incident_subscriber test data: " + err.Error())
	}

	var entityData map[string]any
	if err := json.Unmarshal(entityDataSource, &entityData); err != nil {
		panic("failed to parse incident_subscriber test data: " + err.Error())
	}

	options := map[string]any{}
	options["entity"] = entityData["existing"]

	client := sdk.TestSDK(options, extra)

	// Generate idmap via transform, matching TS pattern.
	idmap := vs.Transform(
		[]any{"incident_subscriber01", "incident_subscriber02", "incident_subscriber03", "page01", "page02", "page03", "incident01", "incident02", "incident03", "subscriber01", "subscriber02", "subscriber03"},
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
	entidEnvRaw := os.Getenv("STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID")
	idmapOverridden := entidEnvRaw != "" && strings.HasPrefix(strings.TrimSpace(entidEnvRaw), "{")

	env := envOverride(map[string]any{
		"STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID": idmap,
		"STATUSPAGE_TEST_LIVE":      "FALSE",
		"STATUSPAGE_TEST_EXPLAIN":   "FALSE",
		"STATUSPAGE_APIKEY":         "NONE",
	})

	idmapResolved := core.ToMapAny(env["STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"])
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
