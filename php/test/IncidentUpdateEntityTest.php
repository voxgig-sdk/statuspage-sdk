<?php
declare(strict_types=1);

// IncidentUpdate entity test

require_once __DIR__ . '/../statuspage_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class IncidentUpdateEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = StatuspageSDK::test(null, null);
        $ent = $testsdk->IncidentUpdate(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = incident_update_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["update"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "incident_update." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // Bootstrap entity data from existing test data.
        $incident_update_ref01_data_raw = Vs::items(Helpers::to_map(
            Vs::getpath($setup["data"], "existing.incident_update")));
        $incident_update_ref01_data = null;
        if (count($incident_update_ref01_data_raw) > 0) {
            $incident_update_ref01_data = Helpers::to_map($incident_update_ref01_data_raw[0][1]);
        }

        // UPDATE
        $incident_update_ref01_ent = $client->IncidentUpdate(null);
        $incident_update_ref01_data_up0_up = [
            "id" => $incident_update_ref01_data["id"],
            "incident_id" => $setup["idmap"]["incident_id"],
            "page_id" => $setup["idmap"]["page_id"],
        ];

        $incident_update_ref01_markdef_up0_name = "body";
        $incident_update_ref01_markdef_up0_value = "Mark01-incident_update_ref01_" . $setup["now"];
        $incident_update_ref01_data_up0_up[$incident_update_ref01_markdef_up0_name] = $incident_update_ref01_markdef_up0_value;

        $incident_update_ref01_resdata_up0_result = $incident_update_ref01_ent->update($incident_update_ref01_data_up0_up, null);
        $incident_update_ref01_resdata_up0 = Helpers::to_map($incident_update_ref01_resdata_up0_result);
        $this->assertNotNull($incident_update_ref01_resdata_up0);
        $this->assertEquals($incident_update_ref01_resdata_up0["id"], $incident_update_ref01_data_up0_up["id"]);
        $this->assertEquals($incident_update_ref01_resdata_up0[$incident_update_ref01_markdef_up0_name], $incident_update_ref01_markdef_up0_value);

    }
}

function incident_update_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/incident_update/IncidentUpdateTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = StatuspageSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["incident_update01", "incident_update02", "incident_update03", "page01", "page02", "page03", "incident01", "incident02", "incident03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID" => $idmap,
        "STATUSPAGE_TEST_LIVE" => "FALSE",
        "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
        "STATUSPAGE_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
    }
    if (!isset($idmap_resolved["incident_id"])) {
        $idmap_resolved["incident_id"] = $idmap_resolved["incident01"];
    }
    if (!isset($idmap_resolved["page_id"])) {
        $idmap_resolved["page_id"] = $idmap_resolved["page01"];
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
