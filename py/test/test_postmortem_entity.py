# Postmortem entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from statuspage_sdk import StatuspageSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestPostmortemEntity:

    def test_should_create_instance(self):
        testsdk = StatuspageSDK.test(None, None)
        ent = testsdk.Postmortem(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _postmortem_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["update", "load"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "postmortem." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set STATUSPAGE_TEST_POSTMORTEM_ENTID JSON to run live")
        client = setup["client"]

        # Bootstrap entity data from existing test data.
        postmortem_ref01_data_raw = vs.items(helpers.to_map(
            vs.getpath(setup["data"], "existing.postmortem")))
        postmortem_ref01_data = None
        if len(postmortem_ref01_data_raw) > 0:
            postmortem_ref01_data = helpers.to_map(postmortem_ref01_data_raw[0][1])

        # UPDATE
        postmortem_ref01_ent = client.Postmortem(None)
        postmortem_ref01_data_up0_up = {
            "page_id": setup["idmap"]["page_id"],
        }

        postmortem_ref01_markdef_up0_name = "body"
        postmortem_ref01_markdef_up0_value = "Mark01-postmortem_ref01_" + str(setup["now"])
        postmortem_ref01_data_up0_up[postmortem_ref01_markdef_up0_name] = postmortem_ref01_markdef_up0_value

        postmortem_ref01_resdata_up0 = helpers.to_map(postmortem_ref01_ent.update(postmortem_ref01_data_up0_up, None))
        assert postmortem_ref01_resdata_up0 is not None
        assert postmortem_ref01_resdata_up0[postmortem_ref01_markdef_up0_name] == postmortem_ref01_markdef_up0_value

        # LOAD
        postmortem_ref01_match_dt0 = {}
        postmortem_ref01_data_dt0_loaded = postmortem_ref01_ent.load(postmortem_ref01_match_dt0, None)
        assert postmortem_ref01_data_dt0_loaded is not None



def _postmortem_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/postmortem/PostmortemTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = StatuspageSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["postmortem01", "postmortem02", "postmortem03", "page01", "page02", "page03", "incident01", "incident02", "incident03"],
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
        "STATUSPAGE_TEST_POSTMORTEM_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "STATUSPAGE_TEST_POSTMORTEM_ENTID": idmap,
        "STATUSPAGE_TEST_LIVE": "FALSE",
        "STATUSPAGE_TEST_EXPLAIN": "FALSE",
        "STATUSPAGE_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("STATUSPAGE_TEST_POSTMORTEM_ENTID"))
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
