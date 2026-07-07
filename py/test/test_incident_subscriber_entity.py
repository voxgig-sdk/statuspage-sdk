# IncidentSubscriber entity test

import json
import os
import time

import pytest

from utility.voxgig_struct import voxgig_struct as vs
from statuspage_sdk import StatuspageSDK
from core import helpers

_TEST_DIR = os.path.dirname(os.path.abspath(__file__))
from test import runner


class TestIncidentSubscriberEntity:

    def test_should_create_instance(self):
        testsdk = StatuspageSDK.test(None, None)
        ent = testsdk.IncidentSubscriber(None)
        assert ent is not None

    def test_should_run_basic_flow(self):
        setup = _incident_subscriber_basic_setup(None)
        # Per-op sdk-test-control.json skip — basic test exercises a flow with
        # multiple ops; skipping any one skips the whole flow (steps depend
        # on each other).
        _live = setup.get("live", False)
        for _op in ["create"]:
            _skip, _reason = runner.is_control_skipped("entityOp", "incident_subscriber." + _op, "live" if _live else "unit")
            if _skip:
                pytest.skip(_reason or "skipped via sdk-test-control.json")
                return
        # The basic flow consumes synthetic IDs from the fixture. In live mode
        # without an *_ENTID env override, those IDs hit the live API and 4xx.
        if setup.get("synthetic_only"):
            pytest.skip("live entity test uses synthetic IDs from fixture — "
                        "set STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID JSON to run live")
        client = setup["client"]

        # CREATE
        incident_subscriber_ref01_ent = client.IncidentSubscriber(None)
        incident_subscriber_ref01_data = helpers.to_map(vs.getprop(
            vs.getpath(setup["data"], "new.incident_subscriber"), "incident_subscriber_ref01"))
        incident_subscriber_ref01_data["incident_id"] = setup["idmap"]["incident01"]
        incident_subscriber_ref01_data["page_id"] = setup["idmap"]["page01"]
        incident_subscriber_ref01_data["subscriber_id"] = setup["idmap"]["subscriber01"]

        incident_subscriber_ref01_data = helpers.to_map(incident_subscriber_ref01_ent.create(incident_subscriber_ref01_data, None))
        assert incident_subscriber_ref01_data is not None



def _incident_subscriber_basic_setup(extra):
    runner.load_env_local()

    entity_data_file = os.path.join(_TEST_DIR, "../../.sdk/test/entity/incident_subscriber/IncidentSubscriberTestData.json")
    with open(entity_data_file, "r") as f:
        entity_data_source = f.read()

    entity_data = json.loads(entity_data_source)

    options = {}
    options["entity"] = entity_data.get("existing")

    client = StatuspageSDK.test(options, extra)

    # Generate idmap via transform.
    idmap = vs.transform(
        ["incident_subscriber01", "incident_subscriber02", "incident_subscriber03", "page01", "page02", "page03", "incident01", "incident02", "incident03", "subscriber01", "subscriber02", "subscriber03"],
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
        "STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID")
    _idmap_overridden = _entid_env_raw is not None and _entid_env_raw.strip().startswith("{")

    env = runner.env_override({
        "STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID": idmap,
        "STATUSPAGE_TEST_LIVE": "FALSE",
        "STATUSPAGE_TEST_EXPLAIN": "FALSE",
        "STATUSPAGE_APIKEY": "NONE",
    })

    idmap_resolved = helpers.to_map(
        env.get("STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"))
    if idmap_resolved is None:
        idmap_resolved = helpers.to_map(idmap)

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
