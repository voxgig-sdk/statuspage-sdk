<?php
declare(strict_types=1);

// IncidentSubscriber entity test

require_once __DIR__ . '/../statuspage_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class IncidentSubscriberEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = StatuspageSDK::test(null, null);
        $ent = $testsdk->IncidentSubscriber(null);
        $this->assertNotNull($ent);
    }

    public function test_basic_flow(): void
    {
        $setup = incident_subscriber_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "incident_subscriber." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $incident_subscriber_ref01_ent = $client->IncidentSubscriber(null);
        $incident_subscriber_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.incident_subscriber"), "incident_subscriber_ref01"));
        $incident_subscriber_ref01_data["incident_id"] = $setup["idmap"]["incident01"];
        $incident_subscriber_ref01_data["page_id"] = $setup["idmap"]["page01"];
        $incident_subscriber_ref01_data["subscriber_id"] = $setup["idmap"]["subscriber01"];

        $incident_subscriber_ref01_data_result = $incident_subscriber_ref01_ent->create($incident_subscriber_ref01_data, null);
        $incident_subscriber_ref01_data = Helpers::to_map($incident_subscriber_ref01_data_result);
        $this->assertNotNull($incident_subscriber_ref01_data);

    }
}

function incident_subscriber_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/incident_subscriber/IncidentSubscriberTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = StatuspageSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["incident_subscriber01", "incident_subscriber02", "incident_subscriber03", "page01", "page02", "page03", "incident01", "incident02", "incident03", "subscriber01", "subscriber02", "subscriber03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID" => $idmap,
        "STATUSPAGE_TEST_LIVE" => "FALSE",
        "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
        "STATUSPAGE_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["STATUSPAGE_TEST_INCIDENT_SUBSCRIBER_ENTID"]);
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
