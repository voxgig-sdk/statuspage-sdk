<?php
declare(strict_types=1);

// StatusEmbedConfig entity test

require_once __DIR__ . '/../statuspage_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class StatusEmbedConfigEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = StatuspageSDK::test(null, null);
        $ent = $testsdk->StatusEmbedConfig(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = status_embed_config_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["update", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "status_embed_config." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $status_embed_config_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.status_embed_config")));
        $status_embed_config_ref01_data = null;
        if (count($status_embed_config_ref01_data_raw) > 0) {
            $status_embed_config_ref01_data = Helpers::to_map($status_embed_config_ref01_data_raw[0][1]);
        }

        // UPDATE
        $status_embed_config_ref01_ent = $client->StatusEmbedConfig(null);
        $status_embed_config_ref01_data_up0_up = [
        ];

        $status_embed_config_ref01_markdef_up0_name = "incident_background_color";
        $status_embed_config_ref01_markdef_up0_value = "Mark01-status_embed_config_ref01_" . $setup["now"];
        $status_embed_config_ref01_data_up0_up[$status_embed_config_ref01_markdef_up0_name] = $status_embed_config_ref01_markdef_up0_value;

        $status_embed_config_ref01_resdata_up0_result = $status_embed_config_ref01_ent->update($status_embed_config_ref01_data_up0_up, null);
        $status_embed_config_ref01_resdata_up0 = Helpers::to_map($status_embed_config_ref01_resdata_up0_result);
        $this->assertNotNull($status_embed_config_ref01_resdata_up0);
        $this->assertEquals($status_embed_config_ref01_resdata_up0[$status_embed_config_ref01_markdef_up0_name], $status_embed_config_ref01_markdef_up0_value);

        // LOAD
        $status_embed_config_ref01_match_dt0 = [];
        $status_embed_config_ref01_data_dt0_loaded = $status_embed_config_ref01_ent->load($status_embed_config_ref01_match_dt0, null);
        $this->assertNotNull($status_embed_config_ref01_data_dt0_loaded);

    }
}

function status_embed_config_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/status_embed_config/StatusEmbedConfigTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = StatuspageSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["status_embed_config01", "status_embed_config02", "status_embed_config03", "page01", "page02", "page03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID" => $idmap,
        "STATUSPAGE_TEST_LIVE" => "FALSE",
        "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
        "STATUSPAGE_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }

    if ($env["STATUSPAGE_TEST_LIVE"] === "TRUE") {
        $merged_opts = Vs::merge([
            [
                "apikey" => $env["STATUSPAGE_APIKEY"],
            ],
            $extra ?? [],
        ]);
        $client = new StatuspageSDK(Helpers::to_map($merged_opts));
    }

    $live = $env["STATUSPAGE_TEST_LIVE"] === "TRUE";
    return [
        "client" => $client,
        "data" => $entity_data,
        "idmap" => $idmap_resolved,
        "env" => $env,
        "explain" => $env["STATUSPAGE_TEST_EXPLAIN"] === "TRUE",
        "live" => $live,
        "synthetic_only" => $live && !$idmap_overridden,
        "now" => (int)(microtime(true) * 1000),
    ];
}
