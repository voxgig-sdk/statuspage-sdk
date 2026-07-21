-- IncidentSubscriber entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("statuspage_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("IncidentSubscriberEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:IncidentSubscriber(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = incident_subscriber_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"create"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "incident_subscriber." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- CREATE
    local incident_subscriber_ref01_ent = client:IncidentSubscriber(nil)
    local incident_subscriber_ref01_data = helpers.to_map(vs.getprop(
      vs.getpath(setup.data, "new.incident_subscriber"), "incident_subscriber_ref01"))
    incident_subscriber_ref01_data["incident_id"] = setup.idmap["incident01"]
    incident_subscriber_ref01_data["page_id"] = setup.idmap["page01"]
    incident_subscriber_ref01_data["subscriber_id"] = setup.idmap["subscriber01"]

    local incident_subscriber_ref01_data_result, err = incident_subscriber_ref01_ent:create(incident_subscriber_ref01_data, nil)
    assert.is_nil(err)
    incident_subscriber_ref01_data = helpers.to_map(incident_subscriber_ref01_data_result)
    assert.is_not_nil(incident_subscriber_ref01_data)

  end)
end)

function incident_subscriber_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/incident_subscriber/IncidentSubscriberTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read incident_subscriber test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "incident_subscriber01", "incident_subscriber02", "incident_subscriber03", "page01", "page02", "page03", "incident01", "incident02", "incident03", "subscriber01", "subscriber02", "subscriber03" },
    {
      ["`$PACK`"] = { "", {
        ["`$KEY`"] = "`$COPY`",
        ["`$VAL`"] = { "`$FORMAT`", "upper", "`$COPY`" },
      }},
    }
  )

  -- Detect ENTID env override before envOverride consumes it. When live
  -- mode is on without a real override, the basic test runs against synthetic
  -- IDs from the fixture and 4xx's. Surface this so the test can skip.
  local entid_env_raw = os.getenv("STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"] = idmap,
    ["STATUSPAGE_TEST_LIVE"] = "FALSE",
    ["STATUSPAGE_TEST_EXPLAIN"] = "FALSE",
    ["STATUSPAGE_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"])
  if idmap_resolved == nil then
    idmap_resolved = helpers.to_map(idmap)
  end

  if env["STATUSPAGE_TEST_LIVE"] == "TRUE" then
    local merged_opts = vs.merge({
      {
        apikey = env["STATUSPAGE_APIKEY"],
      },
      extra or {},
    })
    client = sdk.new(helpers.to_map(merged_opts))
  end

  local live = env["STATUSPAGE_TEST_LIVE"] == "TRUE"
  return {
    client = client,
    data = entity_data,
    idmap = idmap_resolved,
    env = env,
    explain = env["STATUSPAGE_TEST_EXPLAIN"] == "TRUE",
    live = live,
    synthetic_only = live and not idmap_overridden,
    now = os.time() * 1000,
  }
end
