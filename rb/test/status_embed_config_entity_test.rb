# StatusEmbedConfig entity test

require "minitest/autorun"
require "json"
require_relative "../Statuspage_sdk"
require_relative "runner"

class StatusEmbedConfigEntityTest < Minitest::Test
  def test_create_instance
    testsdk = StatuspageSDK.test(nil, nil)
    ent = testsdk.StatusEmbedConfig(nil)
    assert !ent.nil?
  end

  def test_basic_flow
    setup = status_embed_config_basic_setup(nil)
    # Per-op sdk-test-control.json skip.
    _live = setup[:live] || false
    ["update", "load"].each do |_op|
      _should_skip, _reason = Runner.is_control_skipped("entityOp", "status_embed_config." + _op, _live ? "live" : "unit")
      if _should_skip
        skip(_reason || "skipped via sdk-test-control.json")
        return
      end
    end
    # The basic flow consumes synthetic IDs from the fixture. In live mode
    # without an *_ENTID env override, those IDs hit the live API and 4xx.
    if setup[:synthetic_only]
      skip "live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID JSON to run live"
      return
    end
    client = setup[:client]

    # Bootstrap entity data from existing test data.
    status_embed_config_ref01_data_raw = Vs.items(Helpers.to_map(
      Vs.getpath(setup[:data], "existing.status_embed_config")))
    status_embed_config_ref01_data = nil
    if status_embed_config_ref01_data_raw.length > 0
      status_embed_config_ref01_data = Helpers.to_map(status_embed_config_ref01_data_raw[0][1])
    end

    # UPDATE
    status_embed_config_ref01_ent = client.StatusEmbedConfig(nil)
    status_embed_config_ref01_data_up0_up = {
    }

    status_embed_config_ref01_markdef_up0_name = "incident_background_color"
    status_embed_config_ref01_markdef_up0_value = "Mark01-status_embed_config_ref01_#{setup[:now]}"
    status_embed_config_ref01_data_up0_up[status_embed_config_ref01_markdef_up0_name] = status_embed_config_ref01_markdef_up0_value

    status_embed_config_ref01_resdata_up0_result = status_embed_config_ref01_ent.update(status_embed_config_ref01_data_up0_up, nil)
    status_embed_config_ref01_resdata_up0 = Helpers.to_map(status_embed_config_ref01_resdata_up0_result)
    assert !status_embed_config_ref01_resdata_up0.nil?
    assert_equal status_embed_config_ref01_resdata_up0[status_embed_config_ref01_markdef_up0_name], status_embed_config_ref01_markdef_up0_value

    # LOAD
    status_embed_config_ref01_match_dt0 = {}
    status_embed_config_ref01_data_dt0_loaded = status_embed_config_ref01_ent.load(status_embed_config_ref01_match_dt0, nil)
    assert !status_embed_config_ref01_data_dt0_loaded.nil?

  end
end

def status_embed_config_basic_setup(extra)
  Runner.load_env_local

  entity_data_file = File.join(__dir__, "..", "..", ".sdk", "test", "entity", "status_embed_config", "StatusEmbedConfigTestData.json")
  entity_data_source = File.read(entity_data_file)
  entity_data = JSON.parse(entity_data_source)

  options = {}
  options["entity"] = entity_data["existing"]

  client = StatuspageSDK.test(options, extra)

  # Generate idmap via transform.
  idmap = Vs.transform(
    ["status_embed_config01", "status_embed_config02", "status_embed_config03", "page01", "page02", "page03"],
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
  entid_env_raw = ENV["STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID"]
  idmap_overridden = !entid_env_raw.nil? && entid_env_raw.strip.start_with?("{")

  env = Runner.env_override({
    "STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID" => idmap,
    "STATUSPAGE_TEST_LIVE" => "FALSE",
    "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
    "STATUSPAGE_APIKEY" => "NONE",
  })

  idmap_resolved = Helpers.to_map(
    env["STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID"])
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
