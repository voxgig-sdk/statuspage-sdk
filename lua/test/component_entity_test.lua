-- Component entity test

local json = require("dkjson")
local vs = require("utility.struct.struct")
local sdk = require("statuspage_sdk")
local helpers = require("core.helpers")
local runner = require("test.runner")

local _test_dir = debug.getinfo(1, "S").source:match("^@(.+/)")  or "./"

describe("ComponentEntity", function()
  it("should create instance", function()
    local testsdk = sdk.test(nil, nil)
    local ent = testsdk:Component(nil)
    assert.is_not_nil(ent)
  end)

  -- Feature #4: the entity stream(action, ...) method runs the op pipeline and
  -- returns an iterator over result items. With the streaming feature active it
  -- yields the feature's incremental output; otherwise it falls back to the
  -- materialised list so stream always yields.
  it("should stream", function()
    local seed = {
      entity = {
        ["component"] = {
          s1 = { id = "s1" },
          s2 = { id = "s2" },
          s3 = { id = "s3" },
        },
      },
    }

    -- Fallback: streaming inactive -> yields the materialised list items.
    local base = sdk.test(seed, nil)
    local seen = {}
    for item in base:Component(nil):stream("list", nil, nil) do
      table.insert(seen, item)
    end
    assert.are.equal(3, #seen)

    -- Inbound: streaming active -> yields each item from the feature.
    local config = require("config")()
    if type(config.feature) == "table" and config.feature.streaming ~= nil then
      local streamsdk = sdk.test(seed, { feature = { streaming = { active = true } } })
      local got = {}
      for item in streamsdk:Component(nil):stream("list", nil, nil) do
        if vs.islist(item) then
          for _, sub in ipairs(item) do
            table.insert(got, sub)
          end
        else
          table.insert(got, item)
        end
      end
      assert.are.equal(3, #got)
    end
  end)

  it("should run basic flow", function()
    local setup = component_basic_setup(nil)
    -- Per-op sdk-test-control.json skip.
    local _live = setup.live or false
    for _, _op in ipairs({"create", "list", "update", "load", "remove"}) do
      local _should_skip, _reason = runner.is_control_skipped("entityOp", "component." .. _op, _live and "live" or "unit")
      if _should_skip then
        pending(_reason or "skipped via sdk-test-control.json")
        return
      end
    end
    -- The basic flow consumes synthetic IDs from the fixture. In live mode
    -- without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup.synthetic_only then
      pending("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_COMPONENT_ENTID JSON to run live")
      return
    end
    local client = setup.client

    -- CREATE
    local component_ref01_ent = client:Component(nil)
    local component_ref01_data = helpers.to_map(vs.getprop(
      vs.getpath(setup.data, "new.component"), "component_ref01"))
    component_ref01_data["page_access_group_id"] = setup.idmap["page_access_group01"]
    component_ref01_data["page_access_user_id"] = setup.idmap["page_access_user01"]
    component_ref01_data["page_id"] = setup.idmap["page01"]

    local component_ref01_data_result, err = component_ref01_ent:create(component_ref01_data, nil)
    assert.is_nil(err)
    component_ref01_data = helpers.to_map(component_ref01_data_result)
    assert.is_not_nil(component_ref01_data)
    assert.is_not_nil(component_ref01_data["id"])

    -- LIST
    local component_ref01_match = {
      ["page_id"] = setup.idmap["page01"],
    }

    local component_ref01_list_result, err = component_ref01_ent:list(component_ref01_match, nil)
    assert.is_nil(err)
    assert.is_table(component_ref01_list_result)

    local found_item = vs.select(
      runner.entity_list_to_data(component_ref01_list_result),
      { id = component_ref01_data["id"] })
    assert.is_false(vs.isempty(found_item))

    -- UPDATE
    local component_ref01_data_up0_up = {
      id = component_ref01_data["id"],
      ["page_id"] = setup.idmap["page_id"],
    }

    local component_ref01_markdef_up0_name = "automation_email"
    local component_ref01_markdef_up0_value = "Mark01-component_ref01_" .. tostring(setup.now)
    component_ref01_data_up0_up[component_ref01_markdef_up0_name] = component_ref01_markdef_up0_value

    local component_ref01_resdata_up0_result, err = component_ref01_ent:update(component_ref01_data_up0_up, nil)
    assert.is_nil(err)
    local component_ref01_resdata_up0 = helpers.to_map(component_ref01_resdata_up0_result)
    assert.is_not_nil(component_ref01_resdata_up0)
    assert.are.equal(component_ref01_resdata_up0["id"], component_ref01_data_up0_up["id"])
    assert.are.equal(component_ref01_resdata_up0[component_ref01_markdef_up0_name], component_ref01_markdef_up0_value)

    -- LOAD
    local component_ref01_match_dt0 = {
      id = component_ref01_data["id"],
    }
    local component_ref01_data_dt0_loaded, err = component_ref01_ent:load(component_ref01_match_dt0, nil)
    assert.is_nil(err)
    local component_ref01_data_dt0_load_result = helpers.to_map(component_ref01_data_dt0_loaded)
    assert.is_not_nil(component_ref01_data_dt0_load_result)
    assert.are.equal(component_ref01_data_dt0_load_result["id"], component_ref01_data["id"])

    -- REMOVE
    local component_ref01_match_rm0 = {
      id = component_ref01_data["id"],
    }
    local _, err = component_ref01_ent:remove(component_ref01_match_rm0, nil)
    assert.is_nil(err)

    -- LIST
    local component_ref01_match_rt0 = {
      ["page_id"] = setup.idmap["page01"],
    }

    local component_ref01_list_rt0_result, err = component_ref01_ent:list(component_ref01_match_rt0, nil)
    assert.is_nil(err)
    assert.is_table(component_ref01_list_rt0_result)

    local not_found_item = vs.select(
      runner.entity_list_to_data(component_ref01_list_rt0_result),
      { id = component_ref01_data["id"] })
    assert.is_true(vs.isempty(not_found_item))

  end)
end)

function component_basic_setup(extra)
  runner.load_env_local()

  local entity_data_file = _test_dir .. "../../.sdk/test/entity/component/ComponentTestData.json"
  local f = io.open(entity_data_file, "r")
  if f == nil then
    error("failed to read component test data: " .. entity_data_file)
  end
  local entity_data_source = f:read("*a")
  f:close()

  local entity_data = json.decode(entity_data_source)

  local options = {}
  options["entity"] = entity_data["existing"]

  local client = sdk.test(options, extra)

  -- Generate idmap via transform.
  local idmap = vs.transform(
    { "component01", "component02", "component03", "page01", "page02", "page03", "page_access_group01", "page_access_group02", "page_access_group03", "page_access_user01", "page_access_user02", "page_access_user03" },
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
  local entid_env_raw = os.getenv("STATUSPAGE_TEST_COMPONENT_ENTID")
  local idmap_overridden = entid_env_raw ~= nil and entid_env_raw:match("^%s*{") ~= nil

  local env = runner.env_override({
    ["STATUSPAGE_TEST_COMPONENT_ENTID"] = idmap,
    ["STATUSPAGE_TEST_LIVE"] = "FALSE",
    ["STATUSPAGE_TEST_EXPLAIN"] = "FALSE",
    ["STATUSPAGE_APIKEY"] = "NONE",
  })

  local idmap_resolved = helpers.to_map(
    env["STATUSPAGE_TEST_COMPONENT_ENTID"])
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
