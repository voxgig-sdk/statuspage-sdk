<?php
declare(strict_types=1);

// Statuspage SDK

require_once __DIR__ . '/utility/struct/Struct.php';
require_once __DIR__ . '/core/UtilityType.php';
require_once __DIR__ . '/core/Spec.php';
require_once __DIR__ . '/core/Helpers.php';

// Load utility registration
require_once __DIR__ . '/utility/Register.php';

// Load config and features
require_once __DIR__ . '/config.php';
require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/features.php';

use Voxgig\Struct\Struct;

// Features record diagnostic state on the client as dynamic properties
// (_retry, _cache, _metrics, ...); allow them explicitly (PHP 8.2+
// deprecates implicit dynamic properties).
#[\AllowDynamicProperties]
class StatuspageSDK
{
    public string $mode;
    public array $features;
    public ?array $options;

    private $_utility;
    private $_rootctx;

    public function __construct(array $options = [])
    {
        $this->mode = "live";
        $this->features = [];
        $this->options = null;

        $utility = new StatuspageUtility();
        $this->_utility = $utility;

        $config = StatuspageConfig::make_config();

        $this->_rootctx = ($utility->make_context)([
            "client" => $this,
            "utility" => $utility,
            "config" => $config,
            "options" => $options ?? [],
            "shared" => [],
        ], null);

        $this->options = ($utility->make_options)($this->_rootctx);

        if (Struct::getpath($this->options, "feature.test.active") === true) {
            $this->mode = "test";
        }

        $this->_rootctx->options = $this->options;

        // Add features in the resolved order (make_options puts an explicit
        // list order first, else defaults to test-first). Ordering matters: the
        // `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is
        // current, so `test` must be added before them to sit at the base.
        $feature_opts = StatuspageHelpers::to_map(Struct::getprop($this->options, "feature"));
        if ($feature_opts) {
            $featureorder = Struct::getpath($this->options, "__derived__.featureorder");
            if (is_array($featureorder)) {
                foreach ($featureorder as $fname) {
                    $fopts = StatuspageHelpers::to_map($feature_opts[$fname] ?? null);
                    if ($fopts && isset($fopts["active"]) && $fopts["active"] === true) {
                        ($utility->feature_add)($this->_rootctx, StatuspageFeatures::make_feature($fname));
                    }
                }
            }
        }

        // Add extension features.
        $extend_val = Struct::getprop($this->options, "extend");
        if (is_array($extend_val)) {
            foreach ($extend_val as $f) {
                if (is_object($f) && method_exists($f, 'get_name')) {
                    ($utility->feature_add)($this->_rootctx, $f);
                }
            }
        }

        // Initialize features.
        foreach ($this->features as $f) {
            ($utility->feature_init)($this->_rootctx, $f);
        }

        ($utility->feature_hook)($this->_rootctx, "PostConstruct");
    }

    public function options_map(): array
    {
        $out = Struct::clone($this->options);
        return is_array($out) ? $out : [];
    }

    public function get_utility()
    {
        return StatuspageUtility::copy($this->_utility);
    }

    public function get_root_ctx()
    {
        return $this->_rootctx;
    }

    public function prepare(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;
        $fetchargs = $fetchargs ?? [];

        $ctrl = StatuspageHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "prepare",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $opts = $this->options;
        $path = Struct::getprop($fetchargs, "path") ?? "";
        $path = is_string($path) ? $path : "";
        $method_val = Struct::getprop($fetchargs, "method") ?? "GET";
        $method_val = is_string($method_val) ? $method_val : "GET";
        $params = StatuspageHelpers::to_map(Struct::getprop($fetchargs, "params")) ?? [];
        $query = StatuspageHelpers::to_map(Struct::getprop($fetchargs, "query")) ?? [];
        $headers = ($utility->prepare_headers)($ctx);

        $base = Struct::getprop($opts, "base") ?? "";
        $base = is_string($base) ? $base : "";
        $prefix = Struct::getprop($opts, "prefix") ?? "";
        $prefix = is_string($prefix) ? $prefix : "";
        $suffix = Struct::getprop($opts, "suffix") ?? "";
        $suffix = is_string($suffix) ? $suffix : "";

        $ctx->spec = new StatuspageSpec([
            "base" => $base, "prefix" => $prefix, "suffix" => $suffix,
            "path" => $path, "method" => $method_val,
            "params" => $params, "query" => $query, "headers" => $headers,
            "body" => Struct::getprop($fetchargs, "body"),
            "step" => "start",
        ]);

        // Merge user-provided headers.
        $uh = Struct::getprop($fetchargs, "headers");
        if (is_array($uh)) {
            foreach ($uh as $k => $v) {
                $ctx->spec->headers[$k] = $v;
            }
        }

        [$_, $err] = ($utility->prepare_auth)($ctx);
        if ($err) {
            return ($utility->make_error)($ctx, $err);
        }

        [$fetchdef, $fd_err] = ($utility->make_fetch_def)($ctx);
        if ($fd_err) {
            return ($utility->make_error)($ctx, $fd_err);
        }
        return $fetchdef;
    }

    public function direct(array $fetchargs = []): mixed
    {
        $utility = $this->_utility;

        // direct() is the raw-HTTP escape hatch: it never throws, it returns
        // an {ok, err, ...} dict. prepare() now raises on error, so catch it
        // and surface the failure through the dict instead.
        try {
            $fetchdef = $this->prepare($fetchargs);
        } catch (\Throwable $err) {
            return ["ok" => false, "err" => $err];
        }

        $fetchargs = $fetchargs ?? [];
        $ctrl = StatuspageHelpers::to_map(Struct::getprop($fetchargs, "ctrl")) ?? [];

        $ctx = ($utility->make_context)([
            "opname" => "direct",
            "ctrl" => $ctrl,
        ], $this->_rootctx);

        $url = $fetchdef["url"] ?? "";
        [$fetched, $fetch_err] = ($utility->fetcher)($ctx, $url, $fetchdef);

        if ($fetch_err) {
            return ["ok" => false, "err" => $fetch_err];
        }

        if ($fetched === null) {
            return [
                "ok" => false,
                "err" => $ctx->make_error("direct_no_response", "response: undefined"),
            ];
        }

        if (is_array($fetched)) {
            $status = StatuspageHelpers::to_int(Struct::getprop($fetched, "status"));
            $headers = Struct::getprop($fetched, "headers") ?? [];

            // No-body responses (204, 304) and explicit zero content-length
            // must skip JSON parsing — calling json() on an empty body errors.
            $content_length = is_array($headers) ? ($headers["content-length"] ?? null) : null;
            $no_body = $status === 204 || $status === 304 || (string)$content_length === "0";

            $json_data = null;
            if (!$no_body) {
                $jf = Struct::getprop($fetched, "json");
                if (is_callable($jf)) {
                    try {
                        $json_data = $jf();
                    } catch (\Throwable $e) {
                        // Non-JSON body — leave data null but keep status/ok.
                        $json_data = null;
                    }
                }
            }

            return [
                "ok" => $status >= 200 && $status < 300,
                "status" => $status,
                "headers" => Struct::getprop($fetched, "headers"),
                "data" => $json_data,
            ];
        }

        return [
            "ok" => false,
            "err" => $ctx->make_error("direct_invalid", "invalid response type"),
        ];
    }


    private $_component = null;

    // Canonical facade: $client->Component()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->component()
    // resolves here too.
    public function Component($data = null)
    {
        require_once __DIR__ . '/entity/component_entity.php';
        if ($data === null) {
            if ($this->_component === null) {
                $this->_component = new ComponentEntity($this, null);
            }
            return $this->_component;
        }
        return new ComponentEntity($this, $data);
    }


    private $_component_group_uptime = null;

    // Canonical facade: $client->ComponentGroupUptime()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->component_group_uptime()
    // resolves here too.
    public function ComponentGroupUptime($data = null)
    {
        require_once __DIR__ . '/entity/component_group_uptime_entity.php';
        if ($data === null) {
            if ($this->_component_group_uptime === null) {
                $this->_component_group_uptime = new ComponentGroupUptimeEntity($this, null);
            }
            return $this->_component_group_uptime;
        }
        return new ComponentGroupUptimeEntity($this, $data);
    }


    private $_group_component = null;

    // Canonical facade: $client->GroupComponent()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->group_component()
    // resolves here too.
    public function GroupComponent($data = null)
    {
        require_once __DIR__ . '/entity/group_component_entity.php';
        if ($data === null) {
            if ($this->_group_component === null) {
                $this->_group_component = new GroupComponentEntity($this, null);
            }
            return $this->_group_component;
        }
        return new GroupComponentEntity($this, $data);
    }


    private $_incident = null;

    // Canonical facade: $client->Incident()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->incident()
    // resolves here too.
    public function Incident($data = null)
    {
        require_once __DIR__ . '/entity/incident_entity.php';
        if ($data === null) {
            if ($this->_incident === null) {
                $this->_incident = new IncidentEntity($this, null);
            }
            return $this->_incident;
        }
        return new IncidentEntity($this, $data);
    }


    private $_incident_postmortem = null;

    // Canonical facade: $client->IncidentPostmortem()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->incident_postmortem()
    // resolves here too.
    public function IncidentPostmortem($data = null)
    {
        require_once __DIR__ . '/entity/incident_postmortem_entity.php';
        if ($data === null) {
            if ($this->_incident_postmortem === null) {
                $this->_incident_postmortem = new IncidentPostmortemEntity($this, null);
            }
            return $this->_incident_postmortem;
        }
        return new IncidentPostmortemEntity($this, $data);
    }


    private $_incident_subscriber = null;

    // Canonical facade: $client->IncidentSubscriber()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->incident_subscriber()
    // resolves here too.
    public function IncidentSubscriber($data = null)
    {
        require_once __DIR__ . '/entity/incident_subscriber_entity.php';
        if ($data === null) {
            if ($this->_incident_subscriber === null) {
                $this->_incident_subscriber = new IncidentSubscriberEntity($this, null);
            }
            return $this->_incident_subscriber;
        }
        return new IncidentSubscriberEntity($this, $data);
    }


    private $_incident_template = null;

    // Canonical facade: $client->IncidentTemplate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->incident_template()
    // resolves here too.
    public function IncidentTemplate($data = null)
    {
        require_once __DIR__ . '/entity/incident_template_entity.php';
        if ($data === null) {
            if ($this->_incident_template === null) {
                $this->_incident_template = new IncidentTemplateEntity($this, null);
            }
            return $this->_incident_template;
        }
        return new IncidentTemplateEntity($this, $data);
    }


    private $_incident_update = null;

    // Canonical facade: $client->IncidentUpdate()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->incident_update()
    // resolves here too.
    public function IncidentUpdate($data = null)
    {
        require_once __DIR__ . '/entity/incident_update_entity.php';
        if ($data === null) {
            if ($this->_incident_update === null) {
                $this->_incident_update = new IncidentUpdateEntity($this, null);
            }
            return $this->_incident_update;
        }
        return new IncidentUpdateEntity($this, $data);
    }


    private $_metric = null;

    // Canonical facade: $client->Metric()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->metric()
    // resolves here too.
    public function Metric($data = null)
    {
        require_once __DIR__ . '/entity/metric_entity.php';
        if ($data === null) {
            if ($this->_metric === null) {
                $this->_metric = new MetricEntity($this, null);
            }
            return $this->_metric;
        }
        return new MetricEntity($this, $data);
    }


    private $_metrics_provider = null;

    // Canonical facade: $client->MetricsProvider()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->metrics_provider()
    // resolves here too.
    public function MetricsProvider($data = null)
    {
        require_once __DIR__ . '/entity/metrics_provider_entity.php';
        if ($data === null) {
            if ($this->_metrics_provider === null) {
                $this->_metrics_provider = new MetricsProviderEntity($this, null);
            }
            return $this->_metrics_provider;
        }
        return new MetricsProviderEntity($this, $data);
    }


    private $_page = null;

    // Canonical facade: $client->Page()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->page()
    // resolves here too.
    public function Page($data = null)
    {
        require_once __DIR__ . '/entity/page_entity.php';
        if ($data === null) {
            if ($this->_page === null) {
                $this->_page = new PageEntity($this, null);
            }
            return $this->_page;
        }
        return new PageEntity($this, $data);
    }


    private $_page_access_group = null;

    // Canonical facade: $client->PageAccessGroup()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->page_access_group()
    // resolves here too.
    public function PageAccessGroup($data = null)
    {
        require_once __DIR__ . '/entity/page_access_group_entity.php';
        if ($data === null) {
            if ($this->_page_access_group === null) {
                $this->_page_access_group = new PageAccessGroupEntity($this, null);
            }
            return $this->_page_access_group;
        }
        return new PageAccessGroupEntity($this, $data);
    }


    private $_page_access_user = null;

    // Canonical facade: $client->PageAccessUser()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->page_access_user()
    // resolves here too.
    public function PageAccessUser($data = null)
    {
        require_once __DIR__ . '/entity/page_access_user_entity.php';
        if ($data === null) {
            if ($this->_page_access_user === null) {
                $this->_page_access_user = new PageAccessUserEntity($this, null);
            }
            return $this->_page_access_user;
        }
        return new PageAccessUserEntity($this, $data);
    }


    private $_permission = null;

    // Canonical facade: $client->Permission()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->permission()
    // resolves here too.
    public function Permission($data = null)
    {
        require_once __DIR__ . '/entity/permission_entity.php';
        if ($data === null) {
            if ($this->_permission === null) {
                $this->_permission = new PermissionEntity($this, null);
            }
            return $this->_permission;
        }
        return new PermissionEntity($this, $data);
    }


    private $_postmortem = null;

    // Canonical facade: $client->Postmortem()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->postmortem()
    // resolves here too.
    public function Postmortem($data = null)
    {
        require_once __DIR__ . '/entity/postmortem_entity.php';
        if ($data === null) {
            if ($this->_postmortem === null) {
                $this->_postmortem = new PostmortemEntity($this, null);
            }
            return $this->_postmortem;
        }
        return new PostmortemEntity($this, $data);
    }


    private $_status_embed_config = null;

    // Canonical facade: $client->StatusEmbedConfig()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->status_embed_config()
    // resolves here too.
    public function StatusEmbedConfig($data = null)
    {
        require_once __DIR__ . '/entity/status_embed_config_entity.php';
        if ($data === null) {
            if ($this->_status_embed_config === null) {
                $this->_status_embed_config = new StatusEmbedConfigEntity($this, null);
            }
            return $this->_status_embed_config;
        }
        return new StatusEmbedConfigEntity($this, $data);
    }


    private $_subscriber = null;

    // Canonical facade: $client->Subscriber()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->subscriber()
    // resolves here too.
    public function Subscriber($data = null)
    {
        require_once __DIR__ . '/entity/subscriber_entity.php';
        if ($data === null) {
            if ($this->_subscriber === null) {
                $this->_subscriber = new SubscriberEntity($this, null);
            }
            return $this->_subscriber;
        }
        return new SubscriberEntity($this, $data);
    }


    private $_user = null;

    // Canonical facade: $client->User()->list() / ->load(["id" => ...]).
    // PHP method names are case-insensitive, so lowercase $client->user()
    // resolves here too.
    public function User($data = null)
    {
        require_once __DIR__ . '/entity/user_entity.php';
        if ($data === null) {
            if ($this->_user === null) {
                $this->_user = new UserEntity($this, null);
            }
            return $this->_user;
        }
        return new UserEntity($this, $data);
    }



    public static function test(?array $testopts = null, ?array $sdkopts = null): self
    {
        $sdkopts = $sdkopts ?? [];
        $sdkopts = Struct::clone($sdkopts);
        $sdkopts = is_array($sdkopts) ? $sdkopts : [];

        $testopts = $testopts ?? [];
        $testopts = Struct::clone($testopts);
        $testopts = is_array($testopts) ? $testopts : [];
        $testopts["active"] = true;

        if (!isset($sdkopts["feature"])) {
            $sdkopts["feature"] = [];
        }
        $sdkopts["feature"]["test"] = $testopts;

        $sdk = new StatuspageSDK($sdkopts);
        $sdk->mode = "test";
        return $sdk;
    }
}
