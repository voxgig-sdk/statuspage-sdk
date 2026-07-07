# Incident entity test

require "minitest/autorun"
require "json"
require_relative "../Statuspage_sdk"
require_relative "runner"

class IncidentEntityTest < Minitest::Test
  def test_create_instance
    testsdk = StatuspageSDK.test(nil, nil)
    ent = testsdk.Incident(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = incident_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "update", "load", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "incident." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    incident_ref01_ent = client.Incident(nil)
    incident_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.incident"), "incident_ref01"))
    incident_ref01_data["page_id"] = setup[:idmap]["page01"]

    incident_ref01_data_result = incident_ref01_ent.create(incident_ref01_data, nil)
    incident_ref01_data = Helpers.to_map(incident_ref01_data_result)
    assert !incident_ref01_data.nil?
    assert !incident_ref01_data["id"].nil?

    # LIST
    incident_ref01_match = {
      "page_id" => setup[:idmap]["page01"],
    }

    incident_ref01_list_result = incident_ref01_ent.list(incident_ref01_match, nil)
    assert incident_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(incident_ref01_list_result),
      { "id" => incident_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # UPDATE
    incident_ref01_data_up0_up = {
      "id" => incident_ref01_data["id"],
      "page_id" => setup[:idmap]["page_id"],
    }

    incident_ref01_markdef_up0_name = "created_at"
    incident_ref01_markdef_up0_value = "Mark01-incident_ref01_#{setup[:now]}"
    incident_ref01_data_up0_up[incident_ref01_markdef_up0_name] = incident_ref01_markdef_up0_value

    incident_ref01_resdata_up0_result = incident_ref01_ent.update(incident_ref01_data_up0_up, nil)
    incident_ref01_resdata_up0 = Helpers.to_map(incident_ref01_resdata_up0_result)
    assert !incident_ref01_resdata_up0.nil?
    assert_equal incident_ref01_resdata_up0["id"], incident_ref01_data_up0_up["id"]
    assert_equal incident_ref01_resdata_up0[incident_ref01_markdef_up0_name], incident_ref01_markdef_up0_value

    # LOAD
    incident_ref01_match_dt0 = {
      "id" => incident_ref01_data["id"],
    }
    incident_ref01_data_dt0_loaded = incident_ref01_ent.load(incident_ref01_match_dt0, nil)
    incident_ref01_data_dt0_load_result = Helpers.to_map(incident_ref01_data_dt0_loaded)
    assert !incident_ref01_data_dt0_load_result.nil?
    assert_equal incident_ref01_data_dt0_load_result["id"], incident_ref01_data["id"]

    # REMOVE
    incident_ref01_match_rm0 = {
      "id" => incident_ref01_data["id"],
    }
    incident_ref01_ent.remove(incident_ref01_match_rm0, nil)

    # LIST
    incident_ref01_match_rt0 = {
      "page_id" => setup[:idmap]["page01"],
    }

    incident_ref01_list_rt0_result = incident_ref01_ent.list(incident_ref01_match_rt0, nil)
    assert incident_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(incident_ref01_list_rt0_result),
      { "id" => incident_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def incident_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "incident", "IncidentTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = StatuspageSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["incident01", "incident02", "incident03", "page01", "page02", "page03"],
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
  entid_env_raw = ENV["STATUSPAGE_TEST_INCIDENT_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "STATUSPAGE_TEST_INCIDENT_ENTID" => idmap,
    "STATUSPAGE_TEST_LIVE" => "FALSE",
    "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
    "STATUSPAGE_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["STATUSPAGE_TEST_INCIDENT_ENTID"])
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
