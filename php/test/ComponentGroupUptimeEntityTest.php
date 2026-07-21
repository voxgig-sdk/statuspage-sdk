<?php
declare(strict_types=1);

// ComponentGroupUptime entity test

require_once __DIR__ . '/../statuspage_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class ComponentGroupUptimeEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = StatuspageSDK::test(null, null);
        $ent = $testsdk->ComponentGroupUptime(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = component_group_uptime_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "component_group_uptime." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $component_group_uptime_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.component_group_uptime")));
        $component_group_uptime_ref01_data = null;
        if (count($component_group_uptime_ref01_data_raw) > 0) {
            $component_group_uptime_ref01_data = Helpers::to_map($component_group_uptime_ref01_data_raw[0][1]);
        }

        // LOAD
        $component_group_uptime_ref01_ent = $client->ComponentGroupUptime(null);
        $component_group_uptime_ref01_match_dt0 = [
            "id" => $component_group_uptime_ref01_data["id"],
        ];
        $component_group_uptime_ref01_data_dt0_loaded = $component_group_uptime_ref01_ent->load($component_group_uptime_ref01_match_dt0, null);
        $component_group_uptime_ref01_data_dt0_load_result = Helpers::to_map($component_group_uptime_ref01_data_dt0_loaded);
        $this->assertNotNull($component_group_uptime_ref01_data_dt0_load_result);
        $this->assertEquals($component_group_uptime_ref01_data_dt0_load_result["id"], $component_group_uptime_ref01_data["id"]);

    }
}

function component_group_uptime_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/component_group_uptime/ComponentGroupUptimeTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = StatuspageSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["component_group_uptime01", "component_group_uptime02", "component_group_uptime03", "page01", "page02", "page03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID" => $idmap,
        "STATUSPAGE_TEST_LIVE" => "FALSE",
        "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
        "STATUSPAGE_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID"]);
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
