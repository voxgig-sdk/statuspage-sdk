# Incident entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from statuspage_sdk import StatuspageSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestIncidentEntity:

    def test_should_create_instance(self):
        testsdk = StatuspageSDK.test(None, None)
        ent = testsdk.Incident(None)
        assert ent is not None

    def test_should_stream(self):
        # Feature #4: the entity stream(action, ...) method runs the op
        # pipeline and yields result items. With the streaming feature active
        # it yields the feature's incremental output; otherwise it falls back
        # to the materialised list so stream always yields.
        seed = {
            "entity": {
                "incident": {
                    "s1": {"id": "s1"},
                    "s2": {"id": "s2"},
                    "s3": {"id": "s3"},
                }
            }
        }

        # Fallback: streaming inactive -> yields the materialised list items.
        base = StatuspageSDK.test(seed, None)
        seen = list(base.Incident(None).stream("list", None, None))
        assert len(seen) == 3

        # Inbound: streaming active -> yields each item from the feature.
        from config import make_config
        cfg = make_config()
        if isinstance(cfg.get("feature"), dict) and "streaming" in cfg["feature"]:
            sdk = StatuspageSDK.test(
                seed, {"feature": {"streaming": {"active": True}}})
            got = []
            for item in sdk.Incident(None).stream("list", None, None):
                if isinstance(item, list):
                    got.extend(item)
                else:
                    got.append(item)
            assert len(got) == 3

    def test_should_run_basic_flow(self):
        setup = _incident_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "incident." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set STATUSPAGE_TEST_INCIDENT_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        incident_ref01_ent = client.Incident(None)
        incident_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.incident"), "incident_ref01"))
        incident_ref01_data["page_id"] = setup["idmap"]["page01"]

        incident_ref01_data = helpers.to_map(incident_ref01_ent.create(incident_ref01_data, None))
        assert incident_ref01_data is not None
        assert incident_ref01_data["id"] is not None

        # LIST
        incident_ref01_match = {
            "page_id": setup["idmap"]["page01"],
        }

        incident_ref01_list_result = incident_ref01_ent.list(incident_ref01_match, None)
        assert isinstance(incident_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(incident_ref01_list_result),
            {"id": incident_ref01_data["id"]})
        assert not vs.isempty(found_item)

        # UPDATE
        incident_ref01_data_up0_up = {
            "id": incident_ref01_data["id"],
            "page_id": setup["idmap"]["page_id"],
        }

        incident_ref01_markdef_up0_name = "created_at"
        incident_ref01_markdef_up0_value = "Mark01-incident_ref01_" + str(setup["now"])
        incident_ref01_data_up0_up[incident_ref01_markdef_up0_name] = incident_ref01_markdef_up0_value

        incident_ref01_resdata_up0 = helpers.to_map(incident_ref01_ent.update(incident_ref01_data_up0_up, None))
        assert incident_ref01_resdata_up0 is not None
        assert incident_ref01_resdata_up0["id"] == incident_ref01_data_up0_up["id"]
        assert incident_ref01_resdata_up0[incident_ref01_markdef_up0_name] == incident_ref01_markdef_up0_value

        # LOAD
        incident_ref01_match_dt0 = {
            "id": incident_ref01_data["id"],
        }
        incident_ref01_data_dt0_loaded = incident_ref01_ent.load(incident_ref01_match_dt0, None)
        incident_ref01_data_dt0_load_result = helpers.to_map(incident_ref01_data_dt0_loaded)
        assert incident_ref01_data_dt0_load_result is not None
        assert incident_ref01_data_dt0_load_result["id"] == incident_ref01_data["id"]

        # REMOVE
        incident_ref01_match_rm0 = {
            "id": incident_ref01_data["id"],
        }
        incident_ref01_ent.remove(incident_ref01_match_rm0, None)

        # LIST
        incident_ref01_match_rt0 = {
            "page_id": setup["idmap"]["page01"],
        }

        incident_ref01_list_rt0_result = incident_ref01_ent.list(incident_ref01_match_rt0, None)
        assert isinstance(incident_ref01_list_rt0_result, list)

        not_found_item = vs.select(
            runner.entity_list_to_data(incident_ref01_list_rt0_result),
            {"id": incident_ref01_data["id"]})
        assert vs.isempty(not_found_item)



def _incident_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/incident/IncidentTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = StatuspageSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["incident01", "incident02", "incident03", "page01", "page02", "page03"],
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
        "STATUSPAGE_TEST_INCIDENT_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "STATUSPAGE_TEST_INCIDENT_ENTID": idmap,
        "STATUSPAGE_TEST_LIVE": "FALSE",
        "STATUSPAGE_TEST_EXPLAIN": "FALSE",
        "STATUSPAGE_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("STATUSPAGE_TEST_INCIDENT_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)
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
