# User entity test

require "minitest/autorun"
require "json"
require_relative "../Statuspage_sdk"
require_relative "runner"

class UserEntityTest < Minitest::Test
  def test_create_instance
    testsdk = StatuspageSDK.test(nil, nil)
    ent = testsdk.User(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = user_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["create", "list", "remove"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "user." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_USER_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # CREATE
    user_ref01_ent = client.User(nil)
    user_ref01_data = Helpers.to_map(Vs.getprop(
      Vs.getpath(setup[:data], "new.user"), "user_ref01"))
    user_ref01_data["organization_id"] = setup[:idmap]["organization01"]

    user_ref01_data_result = user_ref01_ent.create(user_ref01_data, nil)
    user_ref01_data = Helpers.to_map(user_ref01_data_result)
    assert !user_ref01_data.nil?
    assert !user_ref01_data["id"].nil?

    # LIST
    user_ref01_match = {
      "organization_id" => setup[:idmap]["organization01"],
    }

    user_ref01_list_result = user_ref01_ent.list(user_ref01_match, nil)
    assert user_ref01_list_result.is_a?(Array)

    found_item = Vs.select(
      Runner.entity_list_to_data(user_ref01_list_result),
      { "id" => user_ref01_data["id"] })
    assert !Vs.isempty(found_item)

    # REMOVE
    user_ref01_match_rm0 = {
      "id" => user_ref01_data["id"],
    }
    user_ref01_ent.remove(user_ref01_match_rm0, nil)

    # LIST
    user_ref01_match_rt0 = {
      "organization_id" => setup[:idmap]["organization01"],
    }

    user_ref01_list_rt0_result = user_ref01_ent.list(user_ref01_match_rt0, nil)
    assert user_ref01_list_rt0_result.is_a?(Array)

    not_found_item = Vs.select(
      Runner.entity_list_to_data(user_ref01_list_rt0_result),
      { "id" => user_ref01_data["id"] })
    assert Vs.isempty(not_found_item)

  end
end

def user_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "user", "UserTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = StatuspageSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["user01", "user02", "user03", "organization01", "organization02", "organization03"],
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
  entid_env_raw = ENV["STATUSPAGE_TEST_USER_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "STATUSPAGE_TEST_USER_ENTID" => idmap,
    "STATUSPAGE_TEST_LIVE" => "FALSE",
    "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
    "STATUSPAGE_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["STATUSPAGE_TEST_USER_ENTID"])
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
