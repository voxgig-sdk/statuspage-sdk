# MetricsProvider entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from statuspage_sdk import StatuspageSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestMetricsProviderEntity:

    def test_should_create_instance(self):
        testsdk = StatuspageSDK.test(None, None)
        ent = testsdk.MetricsProvider(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _metrics_provider_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create", "list", "update", "load", "remove"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "metrics_provider." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set STATUSPAGE_TEST_METRICS_PROVIDER_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        metrics_provider_ref01_ent = client.MetricsProvider(None)
        metrics_provider_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.metrics_provider"), "metrics_provider_ref01"))
        metrics_provider_ref01_data["page_id"] = setup["idmap"]["page01"]

        metrics_provider_ref01_data = helpers.to_map(metrics_provider_ref01_ent.create(metrics_provider_ref01_data, None))
        assert metrics_provider_ref01_data is not None
        assert metrics_provider_ref01_data["id"] is not None

        # LIST
        metrics_provider_ref01_match = {
            "page_id": setup["idmap"]["page01"],
        }

        metrics_provider_ref01_list_result = metrics_provider_ref01_ent.list(metrics_provider_ref01_match, None)
        assert isinstance(metrics_provider_ref01_list_result, list)

        found_item = vs.select(
            runner.entity_list_to_data(metrics_provider_ref01_list_result),
            {"id": metrics_provider_ref01_data["id"]})
        assert not vs.isempty(found_item)

        # UPDATE
        metrics_provider_ref01_data_up0_up = {
            "id": metrics_provider_ref01_data["id"],
            "page_id": setup["idmap"]["page_id"],
        }

        metrics_provider_ref01_markdef_up0_name = "created_at"
        metrics_provider_ref01_markdef_up0_value = "Mark01-metrics_provider_ref01_" + str(setup["now"])
        metrics_provider_ref01_data_up0_up[metrics_provider_ref01_markdef_up0_name] = metrics_provider_ref01_markdef_up0_value

        metrics_provider_ref01_resdata_up0 = helpers.to_map(metrics_provider_ref01_ent.update(metrics_provider_ref01_data_up0_up, None))
        assert metrics_provider_ref01_resdata_up0 is not None
        assert metrics_provider_ref01_resdata_up0["id"] == metrics_provider_ref01_data_up0_up["id"]
        assert metrics_provider_ref01_resdata_up0[metrics_provider_ref01_markdef_up0_name] == metrics_provider_ref01_markdef_up0_value

        # LOAD
        metrics_provider_ref01_match_dt0 = {
            "id": metrics_provider_ref01_data["id"],
        }
        metrics_provider_ref01_data_dt0_loaded = metrics_provider_ref01_ent.load(metrics_provider_ref01_match_dt0, None)
        metrics_provider_ref01_data_dt0_load_result = helpers.to_map(metrics_provider_ref01_data_dt0_loaded)
        assert metrics_provider_ref01_data_dt0_load_result is not None
        assert metrics_provider_ref01_data_dt0_load_result["id"] == metrics_provider_ref01_data["id"]

        # REMOVE
        metrics_provider_ref01_match_rm0 = {
            "id": metrics_provider_ref01_data["id"],
        }
        metrics_provider_ref01_ent.remove(metrics_provider_ref01_match_rm0, None)

        # LIST
        metrics_provider_ref01_match_rt0 = {
            "page_id": setup["idmap"]["page01"],
        }

        metrics_provider_ref01_list_rt0_result = metrics_provider_ref01_ent.list(metrics_provider_ref01_match_rt0, None)
        assert isinstance(metrics_provider_ref01_list_rt0_result, list)

        not_found_item = vs.select(
            runner.entity_list_to_data(metrics_provider_ref01_list_rt0_result),
            {"id": metrics_provider_ref01_data["id"]})
        assert vs.isempty(not_found_item)



def _metrics_provider_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/metrics_provider/MetricsProviderTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = StatuspageSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["metrics_provider01", "metrics_provider02", "metrics_provider03", "page01", "page02", "page03"],
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
        "STATUSPAGE_TEST_METRICS_PROVIDER_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "STATUSPAGE_TEST_METRICS_PROVIDER_ENTID": idmap,
        "STATUSPAGE_TEST_LIVE": "FALSE",
        "STATUSPAGE_TEST_EXPLAIN": "FALSE",
        "STATUSPAGE_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("STATUSPAGE_TEST_METRICS_PROVIDER_ENTID"))
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
