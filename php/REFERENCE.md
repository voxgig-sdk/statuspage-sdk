# Statuspage PHP SDK Reference

Complete API reference for the Statuspage PHP SDK.


## StatuspageSDK

### Constructor

```php
require_once __DIR__ . '/statuspage_sdk.php';

$client = new StatuspageSDK($options);
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$options` | `array` | SDK configuration options. |
| `$options["apikey"]` | `string` | API key for authentication. |
| `$options["base"]` | `string` | Base URL for API requests. |
| `$options["prefix"]` | `string` | URL prefix appended after base. |
| `$options["suffix"]` | `string` | URL suffix appended after path. |
| `$options["headers"]` | `array` | Custom headers for all requests. |
| `$options["feature"]` | `array` | Feature configuration. |
| `$options["system"]` | `array` | System overrides (e.g. custom fetch). |


### Static Methods

#### `StatuspageSDK::test($testopts = null, $sdkopts = null)`

Create a test client with mock features active. Both arguments may be `null`.

```php
$client = StatuspageSDK::test();
```


### Instance Methods

#### `Component($data = null)`

Create a new `ComponentEntity` instance. Pass `null` for no initial data.

#### `ComponentGroupUptime($data = null)`

Create a new `ComponentGroupUptimeEntity` instance. Pass `null` for no initial data.

#### `GroupComponent($data = null)`

Create a new `GroupComponentEntity` instance. Pass `null` for no initial data.

#### `Incident($data = null)`

Create a new `IncidentEntity` instance. Pass `null` for no initial data.

#### `IncidentPostmortem($data = null)`

Create a new `IncidentPostmortemEntity` instance. Pass `null` for no initial data.

#### `IncidentSubscriber($data = null)`

Create a new `IncidentSubscriberEntity` instance. Pass `null` for no initial data.

#### `IncidentTemplate($data = null)`

Create a new `IncidentTemplateEntity` instance. Pass `null` for no initial data.

#### `IncidentUpdate($data = null)`

Create a new `IncidentUpdateEntity` instance. Pass `null` for no initial data.

#### `Metric($data = null)`

Create a new `MetricEntity` instance. Pass `null` for no initial data.

#### `MetricsProvider($data = null)`

Create a new `MetricsProviderEntity` instance. Pass `null` for no initial data.

#### `Page($data = null)`

Create a new `PageEntity` instance. Pass `null` for no initial data.

#### `PageAccessGroup($data = null)`

Create a new `PageAccessGroupEntity` instance. Pass `null` for no initial data.

#### `PageAccessUser($data = null)`

Create a new `PageAccessUserEntity` instance. Pass `null` for no initial data.

#### `Permission($data = null)`

Create a new `PermissionEntity` instance. Pass `null` for no initial data.

#### `Postmortem($data = null)`

Create a new `PostmortemEntity` instance. Pass `null` for no initial data.

#### `StatusEmbedConfig($data = null)`

Create a new `StatusEmbedConfigEntity` instance. Pass `null` for no initial data.

#### `Subscriber($data = null)`

Create a new `SubscriberEntity` instance. Pass `null` for no initial data.

#### `User($data = null)`

Create a new `UserEntity` instance. Pass `null` for no initial data.

#### `options_map(): array`

Return a deep copy of the current SDK options.

#### `get_utility(): StatuspageUtility`

Return a copy of the SDK utility object.

#### `direct(array $fetchargs = []): array`

Make a direct HTTP request to any API endpoint. This is the raw-HTTP escape
hatch: it does **not** throw. It returns a result array
`["ok" => bool, "status" => int, "headers" => array, "data" => mixed]`, or
`["ok" => false, "err" => \Exception]` on failure. Branch on `$result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `$fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `$fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `$fetchargs["params"]` | `array` | Path parameter values for `{param}` substitution. |
| `$fetchargs["query"]` | `array` | Query string parameters. |
| `$fetchargs["headers"]` | `array` | Request headers (merged with defaults). |
| `$fetchargs["body"]` | `mixed` | Request body (arrays are JSON-serialized). |
| `$fetchargs["ctrl"]` | `array` | Control options. |

**Returns:** `array` — the result dict (see above); never throws.

#### `prepare(array $fetchargs = []): mixed`

Prepare a fetch definition without sending the request. Returns the
`$fetchdef` array. Throws on error.


---

## ComponentEntity

```php
$component = $client->Component();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `automation_email` | `string` | No |  |
| `component` | `array` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No |  |
| `group` | `bool` | No |  |
| `group_id` | `string` | No |  |
| `id` | `string` | No |  |
| `major_outage` | `int` | No |  |
| `name` | `string` | No |  |
| `only_show_if_degraded` | `bool` | No |  |
| `page_id` | `string` | No |  |
| `partial_outage` | `int` | No |  |
| `position` | `int` | No |  |
| `range_end` | `string` | No |  |
| `range_start` | `string` | No |  |
| `related_event` | `array` | No |  |
| `showcase` | `bool` | No |  |
| `start_date` | `string` | No |  |
| `status` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `uptime_percentage` | `float` | No |  |
| `warning` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Component()->create([
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Component()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Component()->load(["id" => "component_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Component()->remove(["id" => "component_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Component()->update([
  "id" => "component_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ComponentEntity`

Create a new `ComponentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## ComponentGroupUptimeEntity

```php
$component_group_uptime = $client->ComponentGroupUptime();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |
| `major_outage` | `int` | No |  |
| `name` | `string` | No |  |
| `partial_outage` | `int` | No |  |
| `range_end` | `string` | No |  |
| `range_start` | `string` | No |  |
| `related_event` | `array` | No |  |
| `uptime_percentage` | `float` | No |  |
| `warning` | `string` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->ComponentGroupUptime()->load(["id" => "component_group_uptime_id", "page_id" => "page_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): ComponentGroupUptimeEntity`

Create a new `ComponentGroupUptimeEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## GroupComponentEntity

```php
$group_component = $client->GroupComponent();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `string` | No |  |
| `component_group` | `array` | Yes |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `name` | `string` | No |  |
| `page_id` | `string` | No |  |
| `position` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->GroupComponent()->create([
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->GroupComponent()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->GroupComponent()->load(["id" => "group_component_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->GroupComponent()->remove(["id" => "group_component_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->GroupComponent()->update([
  "id" => "group_component_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): GroupComponentEntity`

Create a new `GroupComponentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IncidentEntity

```php
$incident = $client->Incident();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `bool` | No |  |
| `auto_transition_deliver_notifications_at_start` | `bool` | No |  |
| `auto_transition_to_maintenance_state` | `bool` | No |  |
| `auto_transition_to_operational_state` | `bool` | No |  |
| `component` | `array` | No |  |
| `created_at` | `string` | No |  |
| `id` | `string` | No |  |
| `impact` | `string` | No |  |
| `impact_override` | `string` | No |  |
| `incident` | `array` | Yes |  |
| `incident_impact` | `array` | No |  |
| `incident_update` | `array` | No |  |
| `metadata` | `array` | No |  |
| `monitoring_at` | `string` | No |  |
| `name` | `string` | No |  |
| `page_id` | `string` | No |  |
| `postmortem_body` | `string` | No |  |
| `postmortem_body_last_updated_at` | `string` | No |  |
| `postmortem_ignored` | `bool` | No |  |
| `postmortem_notified_subscriber` | `bool` | No |  |
| `postmortem_notified_twitter` | `bool` | No |  |
| `postmortem_published_at` | `bool` | No |  |
| `reminder_interval` | `string` | No |  |
| `resolved_at` | `string` | No |  |
| `scheduled_auto_completed` | `bool` | No |  |
| `scheduled_auto_in_progress` | `bool` | No |  |
| `scheduled_for` | `string` | No |  |
| `scheduled_remind_prior` | `bool` | No |  |
| `scheduled_reminded_at` | `string` | No |  |
| `scheduled_until` | `string` | No |  |
| `shortlink` | `string` | No |  |
| `status` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | - | - | - | - | - |
| `auto_transition_deliver_notifications_at_start` | - | - | - | - | - |
| `auto_transition_to_maintenance_state` | - | - | - | - | - |
| `auto_transition_to_operational_state` | - | - | - | - | - |
| `component` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `impact` | - | - | - | - | - |
| `impact_override` | - | - | - | - | - |
| `incident` | - | - | - | Yes | - |
| `incident_impact` | - | - | - | - | - |
| `incident_update` | - | - | - | - | - |
| `metadata` | - | - | - | - | - |
| `monitoring_at` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `page_id` | - | - | - | - | - |
| `postmortem_body` | - | - | - | - | - |
| `postmortem_body_last_updated_at` | - | - | - | - | - |
| `postmortem_ignored` | - | - | - | - | - |
| `postmortem_notified_subscriber` | - | - | - | - | - |
| `postmortem_notified_twitter` | - | - | - | - | - |
| `postmortem_published_at` | - | - | - | - | - |
| `reminder_interval` | - | - | - | - | - |
| `resolved_at` | - | - | - | - | - |
| `scheduled_auto_completed` | - | - | - | - | - |
| `scheduled_auto_in_progress` | - | - | - | - | - |
| `scheduled_for` | - | - | - | - | - |
| `scheduled_remind_prior` | - | - | - | - | - |
| `scheduled_reminded_at` | - | - | - | - | - |
| `scheduled_until` | - | - | - | - | - |
| `shortlink` | - | - | - | - | - |
| `status` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Incident()->create([
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Incident()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Incident()->load(["id" => "incident_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Incident()->remove(["id" => "incident_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Incident()->update([
  "id" => "incident_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IncidentEntity`

Create a new `IncidentEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IncidentPostmortemEntity

```php
$incident_postmortem = $client->IncidentPostmortem();
```

### Operations

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->IncidentPostmortem()->remove(["id" => "id", "page_id" => "page_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IncidentPostmortemEntity`

Create a new `IncidentPostmortemEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IncidentSubscriberEntity

```php
$incident_subscriber = $client->IncidentSubscriber();
```

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IncidentSubscriber()->create([
  "incident_id" => null, // string
  "page_id" => null, // string
  "subscriber_id" => null, // string
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IncidentSubscriberEntity`

Create a new `IncidentSubscriberEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IncidentTemplateEntity

```php
$incident_template = $client->IncidentTemplate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `string` | No |  |
| `component` | `array` | No |  |
| `group_id` | `string` | No |  |
| `id` | `string` | No |  |
| `name` | `string` | No |  |
| `should_send_notification` | `bool` | No |  |
| `should_tweet` | `bool` | No |  |
| `template` | `array` | Yes |  |
| `title` | `string` | No |  |
| `update_status` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IncidentTemplate()->create([
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->IncidentTemplate()->list();
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IncidentTemplateEntity`

Create a new `IncidentTemplateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## IncidentUpdateEntity

```php
$incident_update = $client->IncidentUpdate();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `affected_component` | `array` | No |  |
| `body` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No |  |
| `deliver_notification` | `bool` | No |  |
| `display_at` | `string` | No |  |
| `id` | `string` | No |  |
| `incident_id` | `string` | No |  |
| `incident_update` | `array` | No |  |
| `status` | `string` | No |  |
| `tweet_id` | `string` | No |  |
| `twitter_updated_at` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `wants_twitter_update` | `bool` | No |  |

### Operations

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->IncidentUpdate()->update([
  "id" => "id",
  "incident_id" => "incident_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): IncidentUpdateEntity`

Create a new `IncidentUpdateEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MetricEntity

```php
$metric = $client->Metric();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `backfill_percentage` | `int` | No |  |
| `backfilled` | `bool` | No |  |
| `created_at` | `string` | No |  |
| `data` | `array` | Yes |  |
| `decimal_place` | `int` | No |  |
| `display` | `bool` | No |  |
| `id` | `string` | No |  |
| `last_fetched_at` | `string` | No |  |
| `metric` | `array` | No |  |
| `metric_identifier` | `string` | No |  |
| `metrics_provider_id` | `string` | No |  |
| `most_recent_data_at` | `string` | No |  |
| `name` | `string` | No |  |
| `reference_name` | `string` | No |  |
| `suffix` | `string` | No |  |
| `tooltip_description` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `y_axis_hidden` | `bool` | No |  |
| `y_axis_max` | `float` | No |  |
| `y_axis_min` | `float` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Metric()->create([
  "metrics_provider_id" => null, // string
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Metric()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Metric()->load(["id" => "metric_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Metric()->remove(["id" => "metric_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Metric()->update([
  "id" => "metric_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MetricEntity`

Create a new `MetricEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## MetricsProviderEntity

```php
$metrics_provider = $client->MetricsProvider();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `disabled` | `bool` | No |  |
| `id` | `string` | No |  |
| `last_revalidated_at` | `string` | No |  |
| `metric_base_uri` | `string` | No |  |
| `metrics_provider` | `array` | No |  |
| `page_id` | `int` | No |  |
| `type` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->MetricsProvider()->create([
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->MetricsProvider()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->MetricsProvider()->load(["id" => "metrics_provider_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->MetricsProvider()->remove(["id" => "metrics_provider_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->MetricsProvider()->update([
  "id" => "metrics_provider_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): MetricsProviderEntity`

Create a new `MetricsProviderEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PageEntity

```php
$page = $client->Page();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activity_score` | `float` | No |  |
| `allow_email_subscriber` | `bool` | No |  |
| `allow_incident_subscriber` | `bool` | No |  |
| `allow_page_subscriber` | `bool` | No |  |
| `allow_rss_atom_feed` | `bool` | No |  |
| `allow_sms_subscriber` | `bool` | No |  |
| `allow_webhook_subscriber` | `bool` | No |  |
| `branding` | `string` | No |  |
| `city` | `string` | No |  |
| `country` | `string` | No |  |
| `created_at` | `string` | No |  |
| `css_blue` | `string` | No |  |
| `css_body_background_color` | `string` | No |  |
| `css_border_color` | `string` | No |  |
| `css_font_color` | `string` | No |  |
| `css_graph_color` | `string` | No |  |
| `css_green` | `string` | No |  |
| `css_light_font_color` | `string` | No |  |
| `css_link_color` | `string` | No |  |
| `css_no_data` | `string` | No |  |
| `css_orange` | `string` | No |  |
| `css_red` | `string` | No |  |
| `css_yellow` | `string` | No |  |
| `domain` | `string` | No |  |
| `email_logo` | `string` | No |  |
| `favicon_logo` | `string` | No |  |
| `headline` | `string` | No |  |
| `hero_cover` | `string` | No |  |
| `hidden_from_search` | `bool` | No |  |
| `id` | `string` | No |  |
| `ip_restriction` | `string` | No |  |
| `name` | `string` | No |  |
| `notifications_email_footer` | `string` | No |  |
| `notifications_from_email` | `string` | No |  |
| `page` | `array` | No |  |
| `page_description` | `string` | No |  |
| `state` | `string` | No |  |
| `subdomain` | `string` | No |  |
| `support_url` | `string` | No |  |
| `time_zone` | `string` | No |  |
| `transactional_logo` | `string` | No |  |
| `twitter_logo` | `string` | No |  |
| `twitter_username` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `url` | `string` | No |  |
| `viewers_must_be_team_member` | `bool` | No |  |

### Operations

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Page()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Page()->load(["id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Page()->update([
  "id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PageEntity`

Create a new `PageEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PageAccessGroupEntity

```php
$page_access_group = $client->PageAccessGroup();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `array` | No |  |
| `created_at` | `string` | No |  |
| `external_identifier` | `string` | No |  |
| `id` | `string` | No |  |
| `metric_id` | `array` | No |  |
| `name` | `string` | No |  |
| `page_access_group` | `array` | No |  |
| `page_access_user_id` | `array` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `component_id` | - | - | Yes | - | - |
| `created_at` | - | - | - | - | - |
| `external_identifier` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `metric_id` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `page_access_group` | - | - | - | - | - |
| `page_access_user_id` | - | - | - | - | - |
| `page_id` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PageAccessGroup()->create([
  "id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PageAccessGroup()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PageAccessGroup()->load(["id" => "page_access_group_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PageAccessGroup()->remove(["id" => "page_access_group_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->PageAccessGroup()->update([
  "id" => "page_access_group_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PageAccessGroupEntity`

Create a new `PageAccessGroupEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PageAccessUserEntity

```php
$page_access_user = $client->PageAccessUser();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `array` | Yes |  |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `external_login` | `string` | No |  |
| `id` | `string` | No |  |
| `metric_id` | `array` | Yes |  |
| `page_access_group_id` | `string` | No |  |
| `page_access_user` | `array` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PageAccessUser()->create([
  "id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->PageAccessUser()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->PageAccessUser()->load(["id" => "page_access_user_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->PageAccessUser()->remove(["id" => "page_access_user_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->PageAccessUser()->update([
  "id" => "page_access_user_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PageAccessUserEntity`

Create a new `PageAccessUserEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PermissionEntity

```php
$permission = $client->Permission();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `array` | No |  |
| `page` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Permission()->load(["id" => "permission_id", "organization_id" => "organization_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Permission()->update([
  "id" => "permission_id",
  "organization_id" => "organization_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PermissionEntity`

Create a new `PermissionEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## PostmortemEntity

```php
$postmortem = $client->Postmortem();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `string` | No |  |
| `body_draft` | `string` | No |  |
| `body_draft_updated_at` | `string` | No |  |
| `body_updated_at` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No |  |
| `notify_subscriber` | `bool` | No |  |
| `notify_twitter` | `bool` | No |  |
| `postmortem` | `array` | Yes |  |
| `preview_key` | `string` | No |  |
| `published_at` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Field Usage by Operation

| Field | load | update |
| --- | --- | --- |
| `body` | - | - |
| `body_draft` | - | - |
| `body_draft_updated_at` | - | - |
| `body_updated_at` | - | - |
| `created_at` | - | - |
| `custom_tweet` | - | - |
| `notify_subscriber` | - | - |
| `notify_twitter` | - | - |
| `postmortem` | - | Yes |
| `preview_key` | - | - |
| `published_at` | - | - |
| `updated_at` | - | - |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Postmortem()->load(["incident_id" => "incident_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Postmortem()->update([
  "incident_id" => "incident_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): PostmortemEntity`

Create a new `PostmortemEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## StatusEmbedConfigEntity

```php
$status_embed_config = $client->StatusEmbedConfig();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `incident_background_color` | `string` | No |  |
| `incident_text_color` | `string` | No |  |
| `maintenance_background_color` | `string` | No |  |
| `maintenance_text_color` | `string` | No |  |
| `page_id` | `string` | No |  |
| `position` | `string` | No |  |
| `status_embed_config` | `array` | No |  |

### Operations

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->StatusEmbedConfig()->load(["page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->StatusEmbedConfig()->update([
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): StatusEmbedConfigEntity`

Create a new `StatusEmbedConfigEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## SubscriberEntity

```php
$subscriber = $client->Subscriber();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `string` | No |  |
| `component_id` | `array` | No |  |
| `created_at` | `string` | No |  |
| `display_phone_number` | `string` | No |  |
| `email` | `string` | No |  |
| `endpoint` | `string` | No |  |
| `id` | `string` | No |  |
| `integration_partner` | `int` | No |  |
| `mode` | `string` | No |  |
| `obfuscated_channel_name` | `string` | No |  |
| `page_access_user_id` | `string` | No |  |
| `phone_country` | `string` | No |  |
| `phone_number` | `string` | No |  |
| `purge_at` | `string` | No |  |
| `quarantined_at` | `string` | No |  |
| `skip_confirmation_notification` | `bool` | No |  |
| `skip_unsubscription_notification` | `bool` | No |  |
| `slack` | `int` | No |  |
| `sms` | `int` | No |  |
| `state` | `string` | No |  |
| `subscriber` | `array` | No |  |
| `team` | `int` | No |  |
| `type` | `string` | No |  |
| `webhook` | `int` | No |  |
| `workspace_name` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `component` | - | - | - | - | - |
| `component_id` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `display_phone_number` | - | - | - | - | - |
| `email` | - | - | - | - | - |
| `endpoint` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `integration_partner` | - | - | - | - | - |
| `mode` | - | - | - | - | - |
| `obfuscated_channel_name` | - | - | - | - | - |
| `page_access_user_id` | - | - | - | - | - |
| `phone_country` | - | - | - | - | - |
| `phone_number` | - | - | - | - | - |
| `purge_at` | - | - | - | - | - |
| `quarantined_at` | - | - | - | - | - |
| `skip_confirmation_notification` | - | - | - | - | - |
| `skip_unsubscription_notification` | - | - | - | - | - |
| `slack` | - | - | - | - | - |
| `sms` | - | - | - | - | - |
| `state` | - | - | - | - | - |
| `subscriber` | - | - | Yes | - | - |
| `team` | - | - | - | - | - |
| `type` | - | - | - | - | - |
| `webhook` | - | - | - | - | - |
| `workspace_name` | - | - | - | - | - |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Subscriber()->create([
  "page_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->Subscriber()->list();
```

#### `load(array $reqmatch, ?array $ctrl = null): mixed`

Load a single entity matching the given criteria. Throws on error.

```php
$result = $client->Subscriber()->load(["id" => "subscriber_id", "page_id" => "page_id"]);
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->Subscriber()->remove(["id" => "subscriber_id", "page_id" => "page_id"]);
```

#### `update(array $reqdata, ?array $ctrl = null): mixed`

Update an existing entity. The data must include the entity `id`. Throws on error.

```php
$result = $client->Subscriber()->update([
  "id" => "subscriber_id",
  "page_id" => "page_id",
  // Fields to update
]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): SubscriberEntity`

Create a new `SubscriberEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## UserEntity

```php
$user = $client->User();
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `first_name` | `string` | No |  |
| `id` | `string` | No |  |
| `last_name` | `string` | No |  |
| `organization_id` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `user` | `array` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->User()->create([
  "organization_id" => null, // string
]);
```

#### `list(?array $reqmatch = null, ?array $ctrl = null): mixed`

List entities matching the given criteria (call with no argument to list all). Returns an array. Throws on error.

```php
$results = $client->User()->list();
```

#### `remove(array $reqmatch, ?array $ctrl = null): mixed`

Remove the entity matching the given criteria. Throws on error.

```php
$result = $client->User()->remove(["id" => "id", "organization_id" => "organization_id"]);
```

### Common Methods

#### `data_get(): array`

Get the entity data. Returns a copy of the current data.

#### `data_set($data): void`

Set the entity data.

#### `match_get(): array`

Get the entity match criteria.

#### `match_set($match): void`

Set the entity match criteria.

#### `make(): UserEntity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name(): string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```php
$client = new StatuspageSDK([
  "feature" => [
    "test" => ["active" => true],
  ],
]);
```

