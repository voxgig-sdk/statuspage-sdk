# IncidentUpdate entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from statuspage_sdk import StatuspageSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestIncidentUpdateEntity:

    def test_should_create_instance(self):
        testsdk = StatuspageSDK.test(None, None)
        ent = testsdk.IncidentUpdate(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _incident_update_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["update"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "incident_update." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        incident_update_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.incident_update")))
        incident_update_ref01_data = None
        if len(incident_update_ref01_data_raw) > 0:
            incident_update_ref01_data = helpers.to_map(incident_update_ref01_data_raw[0][1])

        # UPDATE
        incident_update_ref01_ent = client.IncidentUpdate(None)
        incident_update_ref01_data_up0_up = {
            "id": incident_update_ref01_data["id"],
            "incident_id": setup["idmap"]["incident_id"],
            "page_id": setup["idmap"]["page_id"],
        }

        incident_update_ref01_markdef_up0_name = "body"
        incident_update_ref01_markdef_up0_value = "Mark01-incident_update_ref01_" + str(setup["now"])
        incident_update_ref01_data_up0_up[incident_update_ref01_markdef_up0_name] = incident_update_ref01_markdef_up0_value

        incident_update_ref01_resdata_up0 = helpers.to_map(incident_update_ref01_ent.update(incident_update_ref01_data_up0_up, None))
        assert incident_update_ref01_resdata_up0 is not None
        assert incident_update_ref01_resdata_up0["id"] == incident_update_ref01_data_up0_up["id"]
        assert incident_update_ref01_resdata_up0[incident_update_ref01_markdef_up0_name] == incident_update_ref01_markdef_up0_value



def _incident_update_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/incident_update/IncidentUpdateTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = StatuspageSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["incident_update01", "incident_update02", "incident_update03", "page01", "page02", "page03", "incident01", "incident02", "incident03"],
        {
            "`$PACK`": ["", {
                "`$KEY`": "`$COPY`",
                "`$VAL`": ["`$FORMAT`", "upper", "`$COPY`"],
            }],
        }
    )

    # Detect ENTID env override before envOverride consumes it. When live
    # mode is on without a real override, the basic test runs against synthetic
    # IDs from the fixture and 4xx's. We surface this so the test can skip.
    _entid_env_raw = os.environ.get(
        "STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID": idmap,
        "STATUSPAGE_TEST_LIVE": "FALSE",
        "STATUSPAGE_TEST_EXPLAIN": "FALSE",
        "STATUSPAGE_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)
    if idmap_resolved.get("incident_id") is None:
        idmap_resolved["incident_id"] = idmap_resolved.get("incident01")
    if idmap_resolved.get("page_id") is None:
        idmap_resolved["page_id"] = idmap_resolved.get("page01")

    if env.get("STATUSPAGE_TEST_LIVE") == "TRUE":
        merged_opts = vs.merge([
            {
                "apikey": env.get("STATUSPAGE_APIKEY"),
            },
            extra or {},
        ])
        client = StatuspageSDK(helpers.to_map(merged_opts))

    _live = env.get("STATUSPAGE_TEST_LIVE") == "TRUE"
    return {
        "client": client,
        "data": entity_data,
        "idmap": idmap_resolved,
        "env": env,
        "explain": env.get("STATUSPAGE_TEST_EXPLAIN") == "TRUE",
        "live": _live,
        "synthetic_only": _live and not _idmap_overridden,
        "now": int(time.time() * 1000),
    }
