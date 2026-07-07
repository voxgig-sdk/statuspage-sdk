# IncidentSubscriber entity test

require "minitest/autorun"
require "json"
require_relative "../Statuspage_sdk"
require_relative "runner"

class IncidentSubscriberEntityTest < Minitest::Test
  def test_create_instance
    testsdk = StatuspageSDK.test(nil, nil)
    ent = testsdk.IncidentSubscriber(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = incident_subscriber_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "incident_subscriber." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    incident_subscriber_ref01_ent = client.IncidentSubscriber(nil)
    incident_subscriber_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.incident_subscriber"), "incident_subscriber_ref01"))
    incident_subscriber_ref01_data["incident_id"] = setup[:idmap]["incident01"]
    incident_subscriber_ref01_data["page_id"] = setup[:idmap]["page01"]
    incident_subscriber_ref01_data["subscriber_id"] = setup[:idmap]["subscriber01"]

    incident_subscriber_ref01_data_result = incident_subscriber_ref01_ent.create(incident_subscriber_ref01_data, nil)
    incident_subscriber_ref01_data = Helpers.to_map(incident_subscriber_ref01_data_result)
    assert !incident_subscriber_ref01_data.nil?

  end
end

def incident_subscriber_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "incident_subscriber", "IncidentSubscriberTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = StatuspageSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["incident_subscriber01", "incident_subscriber02", "incident_subscriber03", "page01", "page02", "page03", "incident01", "incident02", "incident03", "subscriber01", "subscriber02", "subscriber03"],
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
  entid_env_raw = ENV["STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID" => idmap,
    "STATUSPAGE_TEST_LIVE" => "FALSE",
    "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
    "STATUSPAGE_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"])
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
