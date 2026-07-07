# Page entity test

require "minitest/autorun"
require "json"
require_relative "../Statuspage_sdk"
require_relative "runner"

class PageEntityTest < Minitest::Test
  def test_create_instance
    testsdk = StatuspageSDK.test(nil, nil)
    ent = testsdk.Page(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = page_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["list", "update", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "page." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_PAGE_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    page_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.page")))
    page_ref01_data = nil
    if page_ref01_data_raw.length > 0
      page_ref01_data = Helpers.to_map(page_ref01_data_raw[0][1])
    end

    # LIST
    page_ref01_ent = client.Page(nil)
    page_ref01_match = {}

    page_ref01_list_result = page_ref01_ent.list(page_ref01_match, nil)
    assert page_ref01_list_result.is_a?(Array)

    # UPDATE
    page_ref01_data_up0_up = {
      "id" => page_ref01_data["id"],
    }

    page_ref01_markdef_up0_name = "branding"
    page_ref01_markdef_up0_value = "Mark01-page_ref01_#{setup[:now]}"
    page_ref01_data_up0_up[page_ref01_markdef_up0_name] = page_ref01_markdef_up0_value

    page_ref01_resdata_up0_result = page_ref01_ent.update(page_ref01_data_up0_up, nil)
    page_ref01_resdata_up0 = Helpers.to_map(page_ref01_resdata_up0_result)
    assert !page_ref01_resdata_up0.nil?
    assert_equal page_ref01_resdata_up0["id"], page_ref01_data_up0_up["id"]
    assert_equal page_ref01_resdata_up0[page_ref01_markdef_up0_name], page_ref01_markdef_up0_value

    # LOAD
    page_ref01_match_dt0 = {
      "id" => page_ref01_data["id"],
    }
    page_ref01_data_dt0_loaded = page_ref01_ent.load(page_ref01_match_dt0, nil)
    page_ref01_data_dt0_load_result = Helpers.to_map(page_ref01_data_dt0_loaded)
    assert !page_ref01_data_dt0_load_result.nil?
    assert_equal page_ref01_data_dt0_load_result["id"], page_ref01_data["id"]

  end
end

def page_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "page", "PageTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = StatuspageSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["page01", "page02", "page03"],
    {
      "`$PACK`" => ["", {
        "`$KEY`" => "`$COPY`",
        "`$VAL`" => ["`$FORMAT`", "upper", "`$COPY`"],
      }],
    }
  )

  # Detect ENTID env override before envOverride consumes it. When live
  # mode is on without a real override, the basic test runs against synthetic
  # IDs from the fixture and 4xx's. Surface this so the test can skip.
  entid_env_raw = ENV["STATUSPAGE_TEST_PAGE_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "STATUSPAGE_TEST_PAGE_ENTID" => idmap,
    "STATUSPAGE_TEST_LIVE" => "FALSE",
    "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
    "STATUSPAGE_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["STATUSPAGE_TEST_PAGE_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end

  if env["STATUSPAGE_TEST_LIVE"] == "TRUE"
    merged_opts = Vs.merge([
      {
        "apikey" => env["STATUSPAGE_APIKEY"],
      },
      extra || {},
    ])
    client = StatuspageSDK.new(Helpers.to_map(merged_opts))
  end

  live = env["STATUSPAGE_TEST_LIVE"] == "TRUE"
  {
    client: client,
    data: entity_data,
    idmap: idmap_resolved,
    env: env,
    explain: env["STATUSPAGE_TEST_EXPLAIN"] == "TRUE",
    live: live,
    synthetic_only: live && !idmap_overridden,
    now: (Time.now.to_f * 1000).to_i,
  }
end
