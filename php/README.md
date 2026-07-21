# Statuspage PHP SDK



The PHP SDK for the Statuspage API — an entity-oriented client using PHP conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `$client->Component()` — with named operations (`list`/`load`/`create`/`update`/`remove`/`patch`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to Packagist. Install it from the
GitHub release tag (`php/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/statuspage-sdk/releases](https://github.com/voxgig-sdk/statuspage-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```php
<?php
require_once 'statuspage_sdk.php';

$client = new StatuspageSDK([
    "apikey" => getenv("STATUSPAGE_APIKEY"),
]);
```

### 2. List component records

```php
try {
    // list() returns an array of Component records — iterate directly.
    $components = $client->Component()->list();
    foreach ($components as $item) {
        echo $item["id"] . " " . $item["automation_email"] . "\n";
    }
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 3. Load a component

Component is nested under page, so provide the `page_id`.

```php
try {
    // load() returns the bare Component record (throws on error).
    $component = $client->Component()->load(["page_id" => "example_page_id", "id" => "example_id"]);
    print_r($component);
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

### 4. Create, update, and remove

```php
// create() returns the bare created Component record.
$created = $client->Component()->create(["page_id" => "example_page_id"]);

// Update — index the bare record directly ($created["id"]).
$client->Component()->update(["id" => $created["id"], "page_id" => "example_page_id"]);

// Remove
$client->Component()->remove(["id" => $created["id"], "page_id" => "example_page_id"]);
```


## Error handling

Entity operations throw a `\Throwable` on failure, so wrap them in
`try` / `catch`:

```php
try {
    $components = $client->Component()->list();
} catch (\Throwable $err) {
    echo "Error: " . $err->getMessage();
}
```

`direct()` does **not** throw — it returns the result array. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```php
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example_id"],
]);

if (! $result["ok"]) {
    $err = $result["err"] ?? null;
    echo "request failed: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```php
// direct() is the raw-HTTP escape hatch: it returns a result array
// (it does not throw). Branch on $result["ok"].
$result = $client->direct([
    "path" => "/api/resource/{id}",
    "method" => "GET",
    "params" => ["id" => "example"],
]);

if ($result["ok"]) {
    echo $result["status"];  // 200
    print_r($result["data"]);  // response body
} else {
    // On an HTTP error status there is no err (only a transport failure sets
    // it), so fall back to the status code.
    $err = $result["err"] ?? null;
    echo "Error: " . ($err ? $err->getMessage() : "HTTP " . $result["status"]);
}
```

### Prepare a request without sending it

```php
// prepare() throws on error and returns the fetch definition.
$fetchdef = $client->prepare([
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => ["id" => "example"],
]);

echo $fetchdef["url"];
echo $fetchdef["method"];
print_r($fetchdef["headers"]);
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```php
$client = StatuspageSDK::test([
    "entity" => ["component" => ["test01" => ["id" => "test01"]]],
]);

// Entity ops return the bare mock record (throws on error).
$component = $client->Component()->list();
print_r($component);
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```php
$mock_fetch = function ($url, $init) {
    return [
        [
            "status" => 200,
            "statusText" => "OK",
            "headers" => [],
            "json" => function () { return ["id" => "mock01"]; },
        ],
        null,
    ];
};

$client = new StatuspageSDK([
    "base" => "http://localhost:8080",
    "system" => [
        "fetch" => $mock_fetch,
    ],
]);
```

### Run live tests

Create a `.env.local` file at the project root:

```
STATUSPAGE_TEST_LIVE=TRUE
STATUSPAGE_APIKEY=<your-key>
```

Then run:

```bash
cd php && ./vendor/bin/phpunit test/
```


## Reference

### StatuspageSDK

```php
require_once 'statuspage_sdk.php';
$client = new StatuspageSDK($options);
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `array` | Feature activation flags. |
| `extend` | `array` | Additional Feature instances to load. |
| `system` | `array` | System overrides (e.g. custom `fetch` callable). |

### test

```php
$client = StatuspageSDK::test($testopts, $sdkopts);
```

Creates a test-mode client with mock transport. Both arguments may be `null`.

### StatuspageSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `(): array` | Deep copy of current SDK options. |
| `get_utility` | `(): Utility` | Copy of the SDK utility object. |
| `prepare` | `(array $fetchargs): array` | Build an HTTP request definition without sending. |
| `direct` | `(array $fetchargs): array` | Build and send an HTTP request. |
| `Component` | `($data): ComponentEntity` | Create a Component entity instance. |
| `ComponentGroupUptime` | `($data): ComponentGroupUptimeEntity` | Create a ComponentGroupUptime entity instance. |
| `GroupComponent` | `($data): GroupComponentEntity` | Create a GroupComponent entity instance. |
| `Incident` | `($data): IncidentEntity` | Create an Incident entity instance. |
| `IncidentPostmortem` | `($data): IncidentPostmortemEntity` | Create an IncidentPostmortem entity instance. |
| `IncidentSubscriber` | `($data): IncidentSubscriberEntity` | Create an IncidentSubscriber entity instance. |
| `IncidentTemplate` | `($data): IncidentTemplateEntity` | Create an IncidentTemplate entity instance. |
| `IncidentUpdate` | `($data): IncidentUpdateEntity` | Create an IncidentUpdate entity instance. |
| `Metric` | `($data): MetricEntity` | Create a Metric entity instance. |
| `MetricsProvider` | `($data): MetricsProviderEntity` | Create a MetricsProvider entity instance. |
| `Page` | `($data): PageEntity` | Create a Page entity instance. |
| `PageAccessGroup` | `($data): PageAccessGroupEntity` | Create a PageAccessGroup entity instance. |
| `PageAccessUser` | `($data): PageAccessUserEntity` | Create a PageAccessUser entity instance. |
| `Permission` | `($data): PermissionEntity` | Create a Permission entity instance. |
| `Postmortem` | `($data): PostmortemEntity` | Create a Postmortem entity instance. |
| `StatusEmbedConfig` | `($data): StatusEmbedConfigEntity` | Create a StatusEmbedConfig entity instance. |
| `Subscriber` | `($data): SubscriberEntity` | Create a Subscriber entity instance. |
| `User` | `($data): UserEntity` | Create an User entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `($reqmatch, $ctrl): array` | Load a single entity by match criteria. |
| `list` | `(?array $reqmatch = null, $ctrl): array` | List entities matching the criteria (call with no argument to list all). |
| `create` | `($reqdata, $ctrl): array` | Create a new entity. |
| `update` | `($reqdata, $ctrl): array` | Update an existing entity. |
| `remove` | `($reqmatch, $ctrl): array` | Remove an entity. |
| `data_get` | `(): array` | Get entity data. |
| `data_set` | `($data): void` | Set entity data. |
| `match_get` | `(): array` | Get entity match criteria. |
| `match_set` | `($match): void` | Set entity match criteria. |
| `make` | `(): Entity` | Create a new instance with the same options. |
| `get_name` | `(): string` | Return the entity name. |

### Result shape

Entity operations return the bare result data (an `array` for single-entity
ops, a `list` for `list`) and throw on error. Wrap calls in
`try`/`catch` to handle failures.

The `direct()` escape hatch never throws — it returns a result `array`
you branch on via `$result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `true` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `array` | Response headers. |
| `data` | `mixed` | Parsed JSON response body. |

On error, `ok` is `false` and `$err` contains the error value.

### Entities

#### Component

| Field | Description |
| --- | --- |
| `automation_email` |  |
| `component` |  |
| `created_at` |  |
| `description` |  |
| `group` |  |
| `group_id` |  |
| `id` |  |
| `major_outage` |  |
| `name` |  |
| `only_show_if_degraded` |  |
| `page_id` |  |
| `partial_outage` |  |
| `position` |  |
| `range_end` |  |
| `range_start` |  |
| `related_event` |  |
| `showcase` |  |
| `start_date` |  |
| `status` |  |
| `updated_at` |  |
| `uptime_percentage` |  |
| `warning` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/components/{component_id}/page_access_groups`

#### ComponentGroupUptime

| Field | Description |
| --- | --- |
| `id` |  |
| `major_outage` |  |
| `name` |  |
| `partial_outage` |  |
| `range_end` |  |
| `range_start` |  |
| `related_event` |  |
| `uptime_percentage` |  |
| `warning` |  |

Operations: Load.

API path: `/pages/{page_id}/component-groups/{id}/uptime`

#### GroupComponent

| Field | Description |
| --- | --- |
| `component` |  |
| `component_group` |  |
| `created_at` |  |
| `description` |  |
| `id` |  |
| `name` |  |
| `page_id` |  |
| `position` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/component-groups`

#### Incident

| Field | Description |
| --- | --- |
| `auto_transition_deliver_notifications_at_end` |  |
| `auto_transition_deliver_notifications_at_start` |  |
| `auto_transition_to_maintenance_state` |  |
| `auto_transition_to_operational_state` |  |
| `component` |  |
| `created_at` |  |
| `id` |  |
| `impact` |  |
| `impact_override` |  |
| `incident` |  |
| `incident_impact` |  |
| `incident_update` |  |
| `metadata` |  |
| `monitoring_at` |  |
| `name` |  |
| `page_id` |  |
| `postmortem_body` |  |
| `postmortem_body_last_updated_at` |  |
| `postmortem_ignored` |  |
| `postmortem_notified_subscriber` |  |
| `postmortem_notified_twitter` |  |
| `postmortem_published_at` |  |
| `reminder_interval` |  |
| `resolved_at` |  |
| `scheduled_auto_completed` |  |
| `scheduled_auto_in_progress` |  |
| `scheduled_for` |  |
| `scheduled_remind_prior` |  |
| `scheduled_reminded_at` |  |
| `scheduled_until` |  |
| `shortlink` |  |
| `status` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/incidents`

#### IncidentPostmortem

| Field | Description |
| --- | --- |

Operations: Remove.

API path: `/pages/{page_id}/incidents/{incident_id}/postmortem`

#### IncidentSubscriber

| Field | Description |
| --- | --- |

Operations: Create.

API path: `/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation`

#### IncidentTemplate

| Field | Description |
| --- | --- |
| `body` |  |
| `component` |  |
| `group_id` |  |
| `id` |  |
| `name` |  |
| `should_send_notification` |  |
| `should_tweet` |  |
| `template` |  |
| `title` |  |
| `update_status` |  |

Operations: Create, List.

API path: `/pages/{page_id}/incident_templates`

#### IncidentUpdate

| Field | Description |
| --- | --- |
| `affected_component` |  |
| `body` |  |
| `created_at` |  |
| `custom_tweet` |  |
| `deliver_notification` |  |
| `display_at` |  |
| `id` |  |
| `incident_id` |  |
| `incident_update` |  |
| `status` |  |
| `tweet_id` |  |
| `twitter_updated_at` |  |
| `updated_at` |  |
| `wants_twitter_update` |  |

Operations: Patch, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}`

#### Metric

| Field | Description |
| --- | --- |
| `backfill_percentage` |  |
| `backfilled` |  |
| `created_at` |  |
| `data` |  |
| `decimal_place` |  |
| `display` |  |
| `id` |  |
| `last_fetched_at` |  |
| `metric` |  |
| `metric_identifier` |  |
| `metrics_provider_id` |  |
| `most_recent_data_at` |  |
| `name` |  |
| `reference_name` |  |
| `suffix` |  |
| `tooltip_description` |  |
| `updated_at` |  |
| `y_axis_hidden` |  |
| `y_axis_max` |  |
| `y_axis_min` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/metrics/{metric_id}/data`

#### MetricsProvider

| Field | Description |
| --- | --- |
| `created_at` |  |
| `disabled` |  |
| `id` |  |
| `last_revalidated_at` |  |
| `metric_base_uri` |  |
| `metrics_provider` |  |
| `page_id` |  |
| `type` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/metrics_providers`

#### Page

| Field | Description |
| --- | --- |
| `activity_score` |  |
| `allow_email_subscriber` |  |
| `allow_incident_subscriber` |  |
| `allow_page_subscriber` |  |
| `allow_rss_atom_feed` |  |
| `allow_sms_subscriber` |  |
| `allow_webhook_subscriber` |  |
| `branding` |  |
| `city` |  |
| `country` |  |
| `created_at` |  |
| `css_blue` |  |
| `css_body_background_color` |  |
| `css_border_color` |  |
| `css_font_color` |  |
| `css_graph_color` |  |
| `css_green` |  |
| `css_light_font_color` |  |
| `css_link_color` |  |
| `css_no_data` |  |
| `css_orange` |  |
| `css_red` |  |
| `css_yellow` |  |
| `domain` |  |
| `email_logo` |  |
| `favicon_logo` |  |
| `headline` |  |
| `hero_cover` |  |
| `hidden_from_search` |  |
| `id` |  |
| `ip_restriction` |  |
| `name` |  |
| `notifications_email_footer` |  |
| `notifications_from_email` |  |
| `page` |  |
| `page_description` |  |
| `state` |  |
| `subdomain` |  |
| `support_url` |  |
| `time_zone` |  |
| `transactional_logo` |  |
| `twitter_logo` |  |
| `twitter_username` |  |
| `updated_at` |  |
| `url` |  |
| `viewers_must_be_team_member` |  |

Operations: List, Load, Patch, Update.

API path: `/pages`

#### PageAccessGroup

| Field | Description |
| --- | --- |
| `component_id` |  |
| `created_at` |  |
| `external_identifier` |  |
| `id` |  |
| `metric_id` |  |
| `name` |  |
| `page_access_group` |  |
| `page_access_user_id` |  |
| `page_id` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_groups/{page_access_group_id}/components`

#### PageAccessUser

| Field | Description |
| --- | --- |
| `component_id` |  |
| `created_at` |  |
| `email` |  |
| `external_login` |  |
| `id` |  |
| `metric_id` |  |
| `page_access_group_id` |  |
| `page_access_user` |  |
| `page_id` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_users/{page_access_user_id}/components`

#### Permission

| Field | Description |
| --- | --- |
| `data` |  |
| `page` |  |

Operations: Load, Update.

API path: `/organizations/{organization_id}/permissions/{user_id}`

#### Postmortem

| Field | Description |
| --- | --- |
| `body` |  |
| `body_draft` |  |
| `body_draft_updated_at` |  |
| `body_updated_at` |  |
| `created_at` |  |
| `custom_tweet` |  |
| `notify_subscriber` |  |
| `notify_twitter` |  |
| `postmortem` |  |
| `preview_key` |  |
| `published_at` |  |
| `updated_at` |  |

Operations: Load, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/postmortem`

#### StatusEmbedConfig

| Field | Description |
| --- | --- |
| `incident_background_color` |  |
| `incident_text_color` |  |
| `maintenance_background_color` |  |
| `maintenance_text_color` |  |
| `page_id` |  |
| `position` |  |
| `status_embed_config` |  |

Operations: Load, Patch, Update.

API path: `/pages/{page_id}/status_embed_config`

#### Subscriber

| Field | Description |
| --- | --- |
| `component` |  |
| `component_id` |  |
| `created_at` |  |
| `display_phone_number` |  |
| `email` |  |
| `endpoint` |  |
| `id` |  |
| `integration_partner` |  |
| `mode` |  |
| `obfuscated_channel_name` |  |
| `page_access_user_id` |  |
| `phone_country` |  |
| `phone_number` |  |
| `purge_at` |  |
| `quarantined_at` |  |
| `skip_confirmation_notification` |  |
| `skip_unsubscription_notification` |  |
| `slack` |  |
| `sms` |  |
| `state` |  |
| `subscriber` |  |
| `team` |  |
| `type` |  |
| `webhook` |  |
| `workspace_name` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation`

#### User

| Field | Description |
| --- | --- |
| `created_at` |  |
| `email` |  |
| `first_name` |  |
| `id` |  |
| `last_name` |  |
| `organization_id` |  |
| `updated_at` |  |
| `user` |  |

Operations: Create, List, Remove.

API path: `/organizations/{organization_id}/users`



## Entities


### Component

Create an instance: `$component = $client->Component();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `automation_email` | `string` |  |
| `component` | `array` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `group` | `bool` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `major_outage` | `int` |  |
| `name` | `string` |  |
| `only_show_if_degraded` | `bool` |  |
| `page_id` | `string` |  |
| `partial_outage` | `int` |  |
| `position` | `int` |  |
| `range_end` | `string` |  |
| `range_start` | `string` |  |
| `related_event` | `array` |  |
| `showcase` | `bool` |  |
| `start_date` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |
| `uptime_percentage` | `float` |  |
| `warning` | `string` |  |

#### Example: Load

```php
// load() returns the bare Component record (throws on error).
$component = $client->Component()->load(["id" => "component_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of Component records (throws on error).
$components = $client->Component()->list();
```

#### Example: Create

```php
$component = $client->Component()->create([
    "page_id" => null, // string
]);
```


### ComponentGroupUptime

Create an instance: `$component_group_uptime = $client->ComponentGroupUptime();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `major_outage` | `int` |  |
| `name` | `string` |  |
| `partial_outage` | `int` |  |
| `range_end` | `string` |  |
| `range_start` | `string` |  |
| `related_event` | `array` |  |
| `uptime_percentage` | `float` |  |
| `warning` | `string` |  |

#### Example: Load

```php
// load() returns the bare ComponentGroupUptime record (throws on error).
$component_group_uptime = $client->ComponentGroupUptime()->load(["id" => "component_group_uptime_id", "page_id" => "page_id"]);
```


### GroupComponent

Create an instance: `$group_component = $client->GroupComponent();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `string` |  |
| `component_group` | `array` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the bare GroupComponent record (throws on error).
$group_component = $client->GroupComponent()->load(["id" => "group_component_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of GroupComponent records (throws on error).
$group_components = $client->GroupComponent()->list();
```

#### Example: Create

```php
$group_component = $client->GroupComponent()->create([
    "page_id" => null, // string
]);
```


### Incident

Create an instance: `$incident = $client->Incident();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `bool` |  |
| `auto_transition_deliver_notifications_at_start` | `bool` |  |
| `auto_transition_to_maintenance_state` | `bool` |  |
| `auto_transition_to_operational_state` | `bool` |  |
| `component` | `array` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `impact` | `string` |  |
| `impact_override` | `string` |  |
| `incident` | `array` |  |
| `incident_impact` | `array` |  |
| `incident_update` | `array` |  |
| `metadata` | `array` |  |
| `monitoring_at` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `postmortem_body` | `string` |  |
| `postmortem_body_last_updated_at` | `string` |  |
| `postmortem_ignored` | `bool` |  |
| `postmortem_notified_subscriber` | `bool` |  |
| `postmortem_notified_twitter` | `bool` |  |
| `postmortem_published_at` | `bool` |  |
| `reminder_interval` | `string` |  |
| `resolved_at` | `string` |  |
| `scheduled_auto_completed` | `bool` |  |
| `scheduled_auto_in_progress` | `bool` |  |
| `scheduled_for` | `string` |  |
| `scheduled_remind_prior` | `bool` |  |
| `scheduled_reminded_at` | `string` |  |
| `scheduled_until` | `string` |  |
| `shortlink` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the bare Incident record (throws on error).
$incident = $client->Incident()->load(["id" => "incident_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of Incident records (throws on error).
$incidents = $client->Incident()->list();
```

#### Example: Create

```php
$incident = $client->Incident()->create([
    "page_id" => null, // string
]);
```


### IncidentPostmortem

Create an instance: `$incident_postmortem = $client->IncidentPostmortem();`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `$incident_subscriber = $client->IncidentSubscriber();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```php
$incident_subscriber = $client->IncidentSubscriber()->create([
    "incident_id" => null, // string
    "page_id" => null, // string
    "subscriber_id" => null, // string
]);
```


### IncidentTemplate

Create an instance: `$incident_template = $client->IncidentTemplate();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` |  |
| `component` | `array` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `should_send_notification` | `bool` |  |
| `should_tweet` | `bool` |  |
| `template` | `array` |  |
| `title` | `string` |  |
| `update_status` | `string` |  |

#### Example: List

```php
// list() returns an array of IncidentTemplate records (throws on error).
$incident_templates = $client->IncidentTemplate()->list();
```

#### Example: Create

```php
$incident_template = $client->IncidentTemplate()->create([
    "page_id" => null, // string
]);
```


### IncidentUpdate

Create an instance: `$incident_update = $client->IncidentUpdate();`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_component` | `array` |  |
| `body` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `deliver_notification` | `bool` |  |
| `display_at` | `string` |  |
| `id` | `string` |  |
| `incident_id` | `string` |  |
| `incident_update` | `array` |  |
| `status` | `string` |  |
| `tweet_id` | `string` |  |
| `twitter_updated_at` | `string` |  |
| `updated_at` | `string` |  |
| `wants_twitter_update` | `bool` |  |


### Metric

Create an instance: `$metric = $client->Metric();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `backfill_percentage` | `int` |  |
| `backfilled` | `bool` |  |
| `created_at` | `string` |  |
| `data` | `array` |  |
| `decimal_place` | `int` |  |
| `display` | `bool` |  |
| `id` | `string` |  |
| `last_fetched_at` | `string` |  |
| `metric` | `array` |  |
| `metric_identifier` | `string` |  |
| `metrics_provider_id` | `string` |  |
| `most_recent_data_at` | `string` |  |
| `name` | `string` |  |
| `reference_name` | `string` |  |
| `suffix` | `string` |  |
| `tooltip_description` | `string` |  |
| `updated_at` | `string` |  |
| `y_axis_hidden` | `bool` |  |
| `y_axis_max` | `float` |  |
| `y_axis_min` | `float` |  |

#### Example: Load

```php
// load() returns the bare Metric record (throws on error).
$metric = $client->Metric()->load(["id" => "metric_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of Metric records (throws on error).
$metrics = $client->Metric()->list();
```

#### Example: Create

```php
$metric = $client->Metric()->create([
    "metrics_provider_id" => null, // string
    "page_id" => null, // string
]);
```


### MetricsProvider

Create an instance: `$metrics_provider = $client->MetricsProvider();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `disabled` | `bool` |  |
| `id` | `string` |  |
| `last_revalidated_at` | `string` |  |
| `metric_base_uri` | `string` |  |
| `metrics_provider` | `array` |  |
| `page_id` | `int` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the bare MetricsProvider record (throws on error).
$metrics_provider = $client->MetricsProvider()->load(["id" => "metrics_provider_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of MetricsProvider records (throws on error).
$metrics_providers = $client->MetricsProvider()->list();
```

#### Example: Create

```php
$metrics_provider = $client->MetricsProvider()->create([
    "page_id" => null, // string
]);
```


### Page

Create an instance: `$page = $client->Page();`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activity_score` | `float` |  |
| `allow_email_subscriber` | `bool` |  |
| `allow_incident_subscriber` | `bool` |  |
| `allow_page_subscriber` | `bool` |  |
| `allow_rss_atom_feed` | `bool` |  |
| `allow_sms_subscriber` | `bool` |  |
| `allow_webhook_subscriber` | `bool` |  |
| `branding` | `string` |  |
| `city` | `string` |  |
| `country` | `string` |  |
| `created_at` | `string` |  |
| `css_blue` | `string` |  |
| `css_body_background_color` | `string` |  |
| `css_border_color` | `string` |  |
| `css_font_color` | `string` |  |
| `css_graph_color` | `string` |  |
| `css_green` | `string` |  |
| `css_light_font_color` | `string` |  |
| `css_link_color` | `string` |  |
| `css_no_data` | `string` |  |
| `css_orange` | `string` |  |
| `css_red` | `string` |  |
| `css_yellow` | `string` |  |
| `domain` | `string` |  |
| `email_logo` | `string` |  |
| `favicon_logo` | `string` |  |
| `headline` | `string` |  |
| `hero_cover` | `string` |  |
| `hidden_from_search` | `bool` |  |
| `id` | `string` |  |
| `ip_restriction` | `string` |  |
| `name` | `string` |  |
| `notifications_email_footer` | `string` |  |
| `notifications_from_email` | `string` |  |
| `page` | `array` |  |
| `page_description` | `string` |  |
| `state` | `string` |  |
| `subdomain` | `string` |  |
| `support_url` | `string` |  |
| `time_zone` | `string` |  |
| `transactional_logo` | `string` |  |
| `twitter_logo` | `string` |  |
| `twitter_username` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |
| `viewers_must_be_team_member` | `bool` |  |

#### Example: Load

```php
// load() returns the bare Page record (throws on error).
$page = $client->Page()->load(["id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of Page records (throws on error).
$pages = $client->Page()->list();
```


### PageAccessGroup

Create an instance: `$page_access_group = $client->PageAccessGroup();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `array` |  |
| `created_at` | `string` |  |
| `external_identifier` | `string` |  |
| `id` | `string` |  |
| `metric_id` | `array` |  |
| `name` | `string` |  |
| `page_access_group` | `array` |  |
| `page_access_user_id` | `array` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the bare PageAccessGroup record (throws on error).
$page_access_group = $client->PageAccessGroup()->load(["id" => "page_access_group_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of PageAccessGroup records (throws on error).
$page_access_groups = $client->PageAccessGroup()->list();
```

#### Example: Create

```php
$page_access_group = $client->PageAccessGroup()->create([
    "id" => null, // string
]);
```


### PageAccessUser

Create an instance: `$page_access_user = $client->PageAccessUser();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `array` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `external_login` | `string` |  |
| `id` | `string` |  |
| `metric_id` | `array` |  |
| `page_access_group_id` | `string` |  |
| `page_access_user` | `array` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the bare PageAccessUser record (throws on error).
$page_access_user = $client->PageAccessUser()->load(["id" => "page_access_user_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of PageAccessUser records (throws on error).
$page_access_users = $client->PageAccessUser()->list();
```

#### Example: Create

```php
$page_access_user = $client->PageAccessUser()->create([
    "id" => null, // string
]);
```


### Permission

Create an instance: `$permission = $client->Permission();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `array` |  |
| `page` | `array` |  |

#### Example: Load

```php
// load() returns the bare Permission record (throws on error).
$permission = $client->Permission()->load(["id" => "permission_id", "organization_id" => "organization_id"]);
```


### Postmortem

Create an instance: `$postmortem = $client->Postmortem();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` |  |
| `body_draft` | `string` |  |
| `body_draft_updated_at` | `string` |  |
| `body_updated_at` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `notify_subscriber` | `bool` |  |
| `notify_twitter` | `bool` |  |
| `postmortem` | `array` |  |
| `preview_key` | `string` |  |
| `published_at` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```php
// load() returns the bare Postmortem record (throws on error).
$postmortem = $client->Postmortem()->load(["incident_id" => "incident_id", "page_id" => "page_id"]);
```


### StatusEmbedConfig

Create an instance: `$status_embed_config = $client->StatusEmbedConfig();`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `incident_background_color` | `string` |  |
| `incident_text_color` | `string` |  |
| `maintenance_background_color` | `string` |  |
| `maintenance_text_color` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `status_embed_config` | `array` |  |

#### Example: Load

```php
// load() returns the bare StatusEmbedConfig record (throws on error).
$status_embed_config = $client->StatusEmbedConfig()->load(["page_id" => "page_id"]);
```


### Subscriber

Create an instance: `$subscriber = $client->Subscriber();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `string` |  |
| `component_id` | `array` |  |
| `created_at` | `string` |  |
| `display_phone_number` | `string` |  |
| `email` | `string` |  |
| `endpoint` | `string` |  |
| `id` | `string` |  |
| `integration_partner` | `int` |  |
| `mode` | `string` |  |
| `obfuscated_channel_name` | `string` |  |
| `page_access_user_id` | `string` |  |
| `phone_country` | `string` |  |
| `phone_number` | `string` |  |
| `purge_at` | `string` |  |
| `quarantined_at` | `string` |  |
| `skip_confirmation_notification` | `bool` |  |
| `skip_unsubscription_notification` | `bool` |  |
| `slack` | `int` |  |
| `sms` | `int` |  |
| `state` | `string` |  |
| `subscriber` | `array` |  |
| `team` | `int` |  |
| `type` | `string` |  |
| `webhook` | `int` |  |
| `workspace_name` | `string` |  |

#### Example: Load

```php
// load() returns the bare Subscriber record (throws on error).
$subscriber = $client->Subscriber()->load(["id" => "subscriber_id", "page_id" => "page_id"]);
```

#### Example: List

```php
// list() returns an array of Subscriber records (throws on error).
$subscribers = $client->Subscriber()->list();
```

#### Example: Create

```php
$subscriber = $client->Subscriber()->create([
    "page_id" => null, // string
]);
```


### User

Create an instance: `$user = $client->User();`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `organization_id` | `string` |  |
| `updated_at` | `string` |  |
| `user` | `array` |  |

#### Example: List

```php
// list() returns an array of User records (throws on error).
$users = $client->User()->list();
```

#### Example: Create

```php
$user = $client->User()->create([
    "organization_id" => null, // string
]);
```


## Advanced

> The sections above cover everyday use. The material below explains the
> SDK's internals — useful when extending it with custom features, but not
> needed for normal use.

### The operation pipeline

Every entity operation follows a six-stage pipeline. Each stage fires a
feature hook before executing:

```
PrePoint → PreSpec → PreRequest → PreResponse → PreResult → PreDone
```

- **PrePoint**: Resolves which API endpoint to call based on the
  operation name and entity configuration.
- **PreSpec**: Builds the HTTP spec — URL, method, headers, body —
  from the resolved point and the caller's parameters.
- **PreRequest**: Sends the HTTP request. Features can intercept here
  to replace the transport (as TestFeature does with mocks).
- **PreResponse**: Parses the raw HTTP response.
- **PreResult**: Extracts the business data from the parsed response.
- **PreDone**: Final stage before returning to the caller. Entity
  state (match, data) is updated here.

If any stage errors, the pipeline short-circuits and the error surfaces
to the caller — see [Error handling](#error-handling) for how that looks
in this language.

### Features and hooks

Features are the extension mechanism. A feature is a PHP class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as arrays

The PHP SDK uses plain PHP associative arrays throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers::to_map()` to safely validate that a value is an array.

### Directory structure

```
php/
├── statuspage_sdk.php          -- Main SDK class
├── config.php                     -- Configuration
├── features.php                   -- Feature factory
├── core/                          -- Core types and context
├── entity/                        -- Entity implementations
├── feature/                       -- Built-in features (Base, Test, Log)
├── utility/                       -- Utility functions and struct library
└── test/                          -- Test suites
```

The main class (`statuspage_sdk.php`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```php
$component = $client->Component();
$component->list();

// $component->data_get() now returns the component data from the last list
// $component->match_get() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
