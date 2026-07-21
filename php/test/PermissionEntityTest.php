<?php
declare(strict_types=1);

// Permission entity test

require_once __DIR__ . '/../statuspage_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class PermissionEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = StatuspageSDK::test(null, null);
        $ent = $testsdk->Permission(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = permission_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["update", "load"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "permission." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_PERMISSION_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $permission_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.permission")));
        $permission_ref01_data = null;
        if (count($permission_ref01_data_raw) > 0) {
            $permission_ref01_data = Helpers::to_map($permission_ref01_data_raw[0][1]);
        }

        // UPDATE
        $permission_ref01_ent = $client->Permission(null);
        $permission_ref01_data_up0_up = [
            "organization_id" => $setup["idmap"]["organization_id"],
        ];

        $permission_ref01_resdata_up0_result = $permission_ref01_ent->update($permission_ref01_data_up0_up, null);
        $permission_ref01_resdata_up0 = Helpers::to_map($permission_ref01_resdata_up0_result);
        $this->assertNotNull($permission_ref01_resdata_up0);

        // LOAD
        $permission_ref01_match_dt0 = [];
        $permission_ref01_data_dt0_loaded = $permission_ref01_ent->load($permission_ref01_match_dt0, null);
        $this->assertNotNull($permission_ref01_data_dt0_loaded);

    }
}

function permission_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/permission/PermissionTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = StatuspageSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["permission01", "permission02", "permission03", "organization01", "organization02", "organization03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("STATUSPAGE_TEST_PERMISSION_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "STATUSPAGE_TEST_PERMISSION_ENTID" => $idmap,
        "STATUSPAGE_TEST_LIVE" => "FALSE",
        "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
        "STATUSPAGE_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["STATUSPAGE_TEST_PERMISSION_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }
    if (!isset($idmap_resolved["organization_id"])) {
        $idmap_resolved["organization_id"] = $idmap_resolved["organization01"];
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
