-- Postmortem entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("statuspage_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("PostmortemEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:Postmortem(nil)
    assert.is_not_nil(ent)
  end)

  it("should run basic flow", function()
    local setup = postmortem_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"update", "load"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "postmortem." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_POSTMORTEM_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- Bootstrap entity data from existing test data.
    local postmortem_ref01_data_raw = vs.items(helpers.to_map(
      vs.getpath(setup.data, "existing.postmortem")))
    local postmortem_ref01_data = nil
    if #postmortem_ref01_data_raw > 0 then
      postmortem_ref01_data = helpers.to_map(postmortem_ref01_data_raw[1][2])
    end

    -- UPDATE
    local postmortem_ref01_ent = client:Postmortem(nil)
    local postmortem_ref01_data_up0_up = {
      ["page_id"] = setup.idmap["page_id"],
    }

    local postmortem_ref01_markdef_up0_name = "body"
    local postmortem_ref01_markdef_up0_value = "Mark01-postmortem_ref01_" .. tostring(setup.now)
    postmortem_ref01_data_up0_up[postmortem_ref01_markdef_up0_name] = postmortem_ref01_markdef_up0_value

    local postmortem_ref01_resdata_up0_result, err = postmortem_ref01_ent:update(postmortem_ref01_data_up0_up, nil)
    assert.is_nil(err)
    local postmortem_ref01_resdata_up0 = helpers.to_map(postmortem_ref01_resdata_up0_result)
    assert.is_not_nil(postmortem_ref01_resdata_up0)
    assert.are.equal(postmortem_ref01_resdata_up0[postmortem_ref01_markdef_up0_name], postmortem_ref01_markdef_up0_value)

    -- LOAD
    local postmortem_ref01_match_dt0 = {}
    local postmortem_ref01_data_dt0_loaded, err = postmortem_ref01_ent:load(postmortem_ref01_match_dt0, nil)
    assert.is_nil(err)
    assert.is_not_nil(postmortem_ref01_data_dt0_loaded)

  end)
end)

function postmortem_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/postmortem/PostmortemTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read postmortem test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "postmortem01", "postmortem02", "postmortem03", "page01", "page02", "page03", "incident01", "incident02", "incident03" },
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
  local entid_env_raw = os.getenv("STATUSPAGE_TEST_POSTMORTEM_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["STATUSPAGE_TEST_POSTMORTEM_ENTID"] = idmap,
    ["STATUSPAGE_TEST_LIVE"] = "FALSE",
    ["STATUSPAGE_TEST_EXPLAIN"] = "FALSE",
    ["STATUSPAGE_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["STATUSPAGE_TEST_POSTMORTEM_ENTID"])
  if idmap_resolved == nil then
    idmap_resolved = helpers.to_map(idmap)
  end
  if idmap_resolved["page_id"] == nil then
    idmap_resolved["page_id"] = idmap_resolved["page01"]
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
