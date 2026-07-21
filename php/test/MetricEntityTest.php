<?php
declare(strict_types=1);

// Metric entity test

require_once __DIR__ . '/../statuspage_sdk.php';
require_once __DIR__ . '/Runner.php';

use PHPUnit\Framework\TestCase;
use Voxgig\Struct\Struct as Vs;

class MetricEntityTest extends TestCase
{
    public function test_create_instance(): void
    {
        $testsdk = StatuspageSDK::test(null, null);
        $ent = $testsdk->Metric(null);
        $this->assertNotNull($ent);
    }

    // Feature #4: the entity stream(action, ...) method runs the op pipeline
    // and yields result items. With the streaming feature active it yields the
    // feature's incremental output; otherwise it falls back to the materialised
    // list so stream always yields.
    public function test_stream(): void
    {
        $seed = [
            "entity" => [
                "metric" => [
                    "s1" => ["id" => "s1"],
                    "s2" => ["id" => "s2"],
                    "s3" => ["id" => "s3"],
                ],
            ],
        ];

        // Fallback: streaming inactive -> yields the materialised list items.
        $base = StatuspageSDK::test($seed, null);
        $seen = iterator_to_array($base->Metric(null)->stream("list", null, null), false);
        $this->assertCount(3, $seen);

        // Inbound: streaming active -> yields each item from the feature.
        $cfg = StatuspageConfig::make_config();
        if (isset($cfg["feature"]) && is_array($cfg["feature"]) && isset($cfg["feature"]["streaming"])) {
            $sdk = StatuspageSDK::test($seed, ["feature" => ["streaming" => ["active" => true]]]);
            $got = [];
            foreach ($sdk->Metric(null)->stream("list", null, null) as $item) {
                if (is_array($item) && array_is_list($item)) {
                    foreach ($item as $sub) {
                        $got[] = $sub;
                    }
                } else {
                    $got[] = $item;
                }
            }
            $this->assertCount(3, $got);
        }
    }

    public function test_basic_flow(): void
    {
        $setup = metric_basic_setup(null);
        // Per-op sdk-test-control.json skip.
        $_live = !empty($setup["live"]);
        foreach (["create", "list", "update", "load", "remove"] as $_op) {
            [$_shouldSkip, $_reason] = Runner::is_control_skipped("entityOp", "metric." . $_op, $_live ? "live" : "unit");
            if ($_shouldSkip) {
                $this->markTestSkipped($_reason ?? "skipped via sdk-test-control.json");
                return;
            }
        }
        // The basic flow consumes synthetic IDs from the fixture. In live mode
        // without an *_ENTID env override, those IDs hit the live API and 4xx.
        if (!empty($setup["synthetic_only"])) {
            $this->markTestSkipped("live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_METRIC_ENTID JSON to run live");
            return;
        }
        $client = $setup["client"];

        // CREATE
        $metric_ref01_ent = $client->Metric(null);
        $metric_ref01_data = Helpers::to_map(Vs::getprop(
            Vs::getpath($setup["data"], "new.metric"), "metric_ref01"));
        $metric_ref01_data["page_access_user_id"] = $setup["idmap"]["page_access_user01"];
        $metric_ref01_data["page_id"] = $setup["idmap"]["page01"];

        $metric_ref01_data_result = $metric_ref01_ent->create($metric_ref01_data, null);
        $metric_ref01_data = Helpers::to_map($metric_ref01_data_result);
        $this->assertNotNull($metric_ref01_data);
        $this->assertNotNull($metric_ref01_data["id"]);

        // LIST
        $metric_ref01_match = [
            "page_access_user_id" => $setup["idmap"]["page_access_user01"],
            "page_id" => $setup["idmap"]["page01"],
        ];

        $metric_ref01_list_result = $metric_ref01_ent->list($metric_ref01_match, null);
        $this->assertIsArray($metric_ref01_list_result);

        $found_item = sdk_select(
            Runner::entity_list_to_data($metric_ref01_list_result),
            ["id" => $metric_ref01_data["id"]]);
        $this->assertNotEmpty($found_item);

        // UPDATE
        $metric_ref01_data_up0_up = [
            "id" => $metric_ref01_data["id"],
            "page_id" => $setup["idmap"]["page_id"],
        ];

        $metric_ref01_markdef_up0_name = "created_at";
        $metric_ref01_markdef_up0_value = "Mark01-metric_ref01_" . $setup["now"];
        $metric_ref01_data_up0_up[$metric_ref01_markdef_up0_name] = $metric_ref01_markdef_up0_value;

        $metric_ref01_resdata_up0_result = $metric_ref01_ent->update($metric_ref01_data_up0_up, null);
        $metric_ref01_resdata_up0 = Helpers::to_map($metric_ref01_resdata_up0_result);
        $this->assertNotNull($metric_ref01_resdata_up0);
        $this->assertEquals($metric_ref01_resdata_up0["id"], $metric_ref01_data_up0_up["id"]);
        $this->assertEquals($metric_ref01_resdata_up0[$metric_ref01_markdef_up0_name], $metric_ref01_markdef_up0_value);

        // LOAD
        $metric_ref01_match_dt0 = [
            "id" => $metric_ref01_data["id"],
        ];
        $metric_ref01_data_dt0_loaded = $metric_ref01_ent->load($metric_ref01_match_dt0, null);
        $metric_ref01_data_dt0_load_result = Helpers::to_map($metric_ref01_data_dt0_loaded);
        $this->assertNotNull($metric_ref01_data_dt0_load_result);
        $this->assertEquals($metric_ref01_data_dt0_load_result["id"], $metric_ref01_data["id"]);

        // REMOVE
        $metric_ref01_match_rm0 = [
            "id" => $metric_ref01_data["id"],
        ];
        $metric_ref01_ent->remove($metric_ref01_match_rm0, null);

        // LIST
        $metric_ref01_match_rt0 = [
            "page_access_user_id" => $setup["idmap"]["page_access_user01"],
            "page_id" => $setup["idmap"]["page01"],
        ];

        $metric_ref01_list_rt0_result = $metric_ref01_ent->list($metric_ref01_match_rt0, null);
        $this->assertIsArray($metric_ref01_list_rt0_result);

        $not_found_item = sdk_select(
            Runner::entity_list_to_data($metric_ref01_list_rt0_result),
            ["id" => $metric_ref01_data["id"]]);
        $this->assertEmpty($not_found_item);

    }
}

function metric_basic_setup($extra)
{
    Runner::load_env_local();

    $entity_data_file = __DIR__ . '/../../.sdk/test/entity/metric/MetricTestData.json';
    $entity_data_source = file_get_contents($entity_data_file);
    $entity_data = json_decode($entity_data_source, true);

    $options = [];
    $options["entity"] = $entity_data["existing"];

    $client = StatuspageSDK::test($options, $extra);

    // Generate idmap.
    $idmap = [];
    foreach (["metric01", "metric02", "metric03", "page01", "page02", "page03", "metrics_provider01", "metrics_provider02", "metrics_provider03", "page_access_user01", "page_access_user02", "page_access_user03"] as $k) {
        $idmap[$k] = strtoupper($k);
    }

    // Detect ENTID env override before envOverride consumes it. When live
    // mode is on without a real override, the basic test runs against synthetic
    // IDs from the fixture and 4xx's. Surface this so the test can skip.
    $entid_env_raw = getenv("STATUSPAGE_TEST_METRIC_ENTID");
    $idmap_overridden = $entid_env_raw !== false && str_starts_with(trim($entid_env_raw), "{");

    $env = Runner::env_override([
        "STATUSPAGE_TEST_METRIC_ENTID" => $idmap,
        "STATUSPAGE_TEST_LIVE" => "FALSE",
        "STATUSPAGE_TEST_EXPLAIN" => "FALSE",
        "STATUSPAGE_APIKEY" => "NONE",
    ]);

    $idmap_resolved = Helpers::to_map(
        $env["STATUSPAGE_TEST_METRIC_ENTID"]);
    if ($idmap_resolved === null) {
        $idmap_resolved = Helpers::to_map($idmap);
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
