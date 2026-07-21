# PageAccessGroup entity test

require "minitest/autorun"
require "json"
require_relative "../Statuspage_sdk"
require_relative "runner"

class PageAccessGroupEntityTest < Minitest::Test
  def test_create_instance
    testsdk = StatuspageSDK.test(nil, nil)
    ent = testsdk.PageAccessGroup(nil)
    assert !ent.nil?
  end

  # Feature #4: the entity stream(action, ...) method runs the op pipeline and
  # returns an Enumerator over result items. With the streaming feature active
  # it yields the feature's incremental output; otherwise it falls back to the
  # materialised list so stream always yields.
  def test_stream
    seed = {
      "entity" => {
        "page_access_group" => {
          "s1" => { "id" => "s1" },
          "s2" => { "id" => "s2" },
          "s3" => { "id" => "s3" },
        },
      },
    }

    # Fallback: streaming inactive -> yields the materialised list items.
    base = StatuspageSDK.test(seed, nil)
    seen = base.PageAccessGroup(nil).stream("list", nil, nil).to_a
    assert_equal 3, seen.length

    # Inbound: streaming active -> yields each item from the feature.
    cfg = StatuspageConfig.make_config
    if cfg["feature"].is_a?(Hash) && cfg["feature"].key?("streaming")
      sdk = StatuspageSDK.test(seed, { "feature" => { "streaming" => { "active" => true } } })
      got = []
      sdk.PageAccessGroup(nil).stream("list", nil, nil).each do |item|
        if item.is_a?(Array)
          got.concat(item)
        else
          got << item
        end
      end
      assert_equal 3, got.length
    end
  end

  def test_basic_flow
    setup = page_access_group_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "page_access_group." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    page_access_group_ref01_ent = client.PageAccessGroup(nil)
    page_access_group_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.page_access_group"), "page_access_group_ref01"))
    page_access_group_ref01_data["page_id"] = setup[:idmap]["page01"]

    page_access_group_ref01_data_result = page_access_group_ref01_ent.create(page_access_group_ref01_data, nil)
    page_access_group_ref01_data = Helpers.to_map(page_access_group_ref01_data_result)
    assert !page_access_group_ref01_data.nil?
    assert !page_access_group_ref01_data["id"].nil?

    # LIST
    page_access_group_ref01_match = {
      "page_id" => setup[:idmap]["page01"],
    }

    page_access_group_ref01_list_result = page_access_group_ref01_ent.list(page_access_group_ref01_match, nil)
    assert page_access_group_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(page_access_group_ref01_list_result),
      { "id" => page_access_group_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # UPDATE
    page_access_group_ref01_data_up0_up = {
      "id" => page_access_group_ref01_data["id"],
      "page_id" => setup[:idmap]["page_id"],
    }

    page_access_group_ref01_markdef_up0_name = "created_at"
    page_access_group_ref01_markdef_up0_value = "Mark01-page_access_group_ref01_#{setup[:now]}"
    page_access_group_ref01_data_up0_up[page_access_group_ref01_markdef_up0_name] = page_access_group_ref01_markdef_up0_value

    page_access_group_ref01_resdata_up0_result = page_access_group_ref01_ent.update(page_access_group_ref01_data_up0_up, nil)
    page_access_group_ref01_resdata_up0 = Helpers.to_map(page_access_group_ref01_resdata_up0_result)
    assert !page_access_group_ref01_resdata_up0.nil?
    assert_equal page_access_group_ref01_resdata_up0["id"], page_access_group_ref01_data_up0_up["id"]
    assert_equal page_access_group_ref01_resdata_up0[page_access_group_ref01_markdef_up0_name], page_access_group_ref01_markdef_up0_value

    # LOAD
    page_access_group_ref01_match_dt0 = {
      "id" => page_access_group_ref01_data["id"],
    }
    page_access_group_ref01_data_dt0_loaded = page_access_group_ref01_ent.load(page_access_group_ref01_match_dt0, nil)
    page_access_group_ref01_data_dt0_load_result = Helpers.to_map(page_access_group_ref01_data_dt0_loaded)
    assert !page_access_group_ref01_data_dt0_load_result.nil?
    assert_equal page_access_group_ref01_data_dt0_load_result["id"], page_access_group_ref01_data["id"]

    # REMOVE
    page_access_group_ref01_match_rm0 = {
      "id" => page_access_group_ref01_data["id"],
    }
    page_access_group_ref01_ent.remove(page_access_group_ref01_match_rm0, nil)

    # LIST
    page_access_group_ref01_match_rt0 = {
      "page_id" => setup[:idmap]["page01"],
    }

    page_access_group_ref01_list_rt0_result = page_access_group_ref01_ent.list(page_access_group_ref01_match_rt0, nil)
    assert page_access_group_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(page_access_group_ref01_list_rt0_result),
      { "id" => page_access_group_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def page_access_group_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "page_access_group", "PageAccessGroupTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = StatuspageSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["page_access_group01", "page_access_group02", "page_access_group03", "page01", "page02", "page03", "component01", "component02", "component03"],
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
  entid_env_raw = ENV["STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID" => idmap,
    "STATUSPAGE_TEST_LIVE" => "FALSE",
    "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
    "STATUSPAGE_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID"])
  if idmap_resolved.nil?
    idmap_resolved = Helpers.to_map(idmap)
  end
  if idmap_resolved["page_id"].nil?
    idmap_resolved["page_id"] = idmap_resolved["page01"]
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
