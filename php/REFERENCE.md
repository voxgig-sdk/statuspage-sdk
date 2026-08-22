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
| `automation_email` | `string` | No | Requires a special feature flag to be enabled |
| `component` | `array` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No | More detailed description for component |
| `group` | `bool` | No | Is this component a group |
| `group_id` | `string` | No | Component Group identifier |
| `id` | `string` | No | Incident identifier |
| `name` | `string` | No | Display name for component |
| `only_show_if_degraded` | `bool` | No | Requires a special feature flag to be enabled |
| `page_id` | `string` | No | Page identifier |
| `position` | `int` | No | Order the component will appear on the page |
| `showcase` | `bool` | No | Should this component be showcased |
| `start_date` | `string` | No | The date this component started being used |
| `status` | `string` | No | Status of component |
| `updated_at` | `string` | No |  |

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
| `component_id` | `string` | No | Component identifier |
| `incidents` | `array` | No | Related incidents |

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
| `component_group` | `array` | Yes |  |
| `components` | `string` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No | Description of the component group. |
| `id` | `string` | No | Component Group Identifier |
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
  "component_group" => null, // array
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
| `auto_transition_deliver_notifications_at_end` | `bool` | No | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `bool` | No | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `bool` | No | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `bool` | No | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `array` | No | Incident components |
| `created_at` | `string` | No | The timestamp when the incident was created at. |
| `id` | `string` | No | Incident Identifier |
| `impact` | `string` | No | The impact of the incident. |
| `impact_override` | `string` | No | value to override calculated impact value |
| `incident` | `array` | Yes |  |
| `incident_updates` | `array` | No | The incident updates for incident. |
| `metadata` | `array` | No | Metadata attached to the incident. |
| `monitoring_at` | `string` | No | The timestamp when incident entered monitoring state. |
| `name` | `string` | No | Incident Name. |
| `page_id` | `string` | No | Incident Page Identifier |
| `postmortem_body` | `string` | No | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `string` | No | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `bool` | No | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `bool` | No | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `bool` | No | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `bool` | No | The timestamp when the postmortem was published. |
| `reminder_intervals` | `string` | No | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `string` | No | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `bool` | No | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `bool` | No | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `string` | No | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `bool` | No | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `string` | No | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `string` | No | The timestamp the incident is scheduled until. |
| `shortlink` | `string` | No | Incident Shortlink. |
| `status` | `string` | No | The incident status. |
| `updated_at` | `string` | No | The timestamp when the incident was updated at. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | - | - | - | - | - |
| `auto_transition_deliver_notifications_at_start` | - | - | - | - | - |
| `auto_transition_to_maintenance_state` | - | - | - | - | - |
| `auto_transition_to_operational_state` | - | - | - | - | - |
| `components` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `impact` | - | - | - | - | - |
| `impact_override` | - | - | - | - | - |
| `incident` | - | - | - | Yes | - |
| `incident_updates` | - | - | - | - | - |
| `metadata` | - | - | - | - | - |
| `monitoring_at` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `page_id` | - | - | - | - | - |
| `postmortem_body` | - | - | - | - | - |
| `postmortem_body_last_updated_at` | - | - | - | - | - |
| `postmortem_ignored` | - | - | - | - | - |
| `postmortem_notified_subscribers` | - | - | - | - | - |
| `postmortem_notified_twitter` | - | - | - | - | - |
| `postmortem_published_at` | - | - | - | - | - |
| `reminder_intervals` | - | - | - | - | - |
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
  "incident" => null, // array
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
| `body` | `string` | No | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `array` | No | Affected components |
| `group_id` | `string` | No | Identifier of Template Group this template belongs to |
| `id` | `string` | No | Incident Template Identifier |
| `name` | `string` | No | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `bool` | No | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `bool` | No | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `array` | Yes |  |
| `title` | `string` | No | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `string` | No | The status the incident or maintenance should transition to when selecting this template |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->IncidentTemplate()->create([
  "page_id" => null, // string
  "template" => null, // array
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
| `affected_components` | `array` | No | Affected components associated with the incident update. |
| `body` | `string` | No | Incident update body. |
| `created_at` | `string` | No | The timestamp when the incident update was created at. |
| `custom_tweet` | `string` | No | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `bool` | No | Controls whether to delivery notifications. |
| `display_at` | `string` | No | Timestamp when incident update is happened. |
| `id` | `string` | No | Incident Update Identifier. |
| `incident_id` | `string` | No | Incident Identifier. |
| `incident_update` | `array` | No |  |
| `status` | `string` | No | The incident status. |
| `tweet_id` | `string` | No | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `string` | No | The timestamp when twitter updated at. |
| `updated_at` | `string` | No | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `bool` | No | Controls whether to create twitter update. |

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
| `data` | `array` | Yes | Add data points to metrics |
| `decimal_places` | `int` | No |  |
| `display` | `bool` | No | Should the metric be displayed |
| `id` | `string` | No | Metric identifier |
| `last_fetched_at` | `string` | No |  |
| `metric` | `array` | No |  |
| `metric_identifier` | `string` | No | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `string` | No | Metric Provider identifier |
| `most_recent_data_at` | `string` | No |  |
| `name` | `string` | No | Name of metric |
| `reference_name` | `string` | No |  |
| `suffix` | `string` | No | Suffix to describe the units on the graph |
| `tooltip_description` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `y_axis_hidden` | `bool` | No | Should the values on the y axis be hidden on render |
| `y_axis_max` | `float` | No |  |
| `y_axis_min` | `float` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Metric()->create([
  "metrics_provider_id" => null, // string
  "page_id" => null, // string
  "data" => null, // array
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
| `id` | `string` | No | Identifier for Metrics Provider |
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
| `allow_email_subscribers` | `bool` | No | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `bool` | No | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `bool` | No | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `bool` | No | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `bool` | No | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `bool` | No | Can your users choose to receive notifications via Webhooks |
| `branding` | `string` | No | The main template your statuspage will use |
| `city` | `string` | No |  |
| `country` | `string` | No |  |
| `created_at` | `string` | No | Timestamp the record was created |
| `css_blues` | `string` | No | CSS Color |
| `css_body_background_color` | `string` | No | CSS Color |
| `css_border_color` | `string` | No | CSS Color |
| `css_font_color` | `string` | No | CSS Color |
| `css_graph_color` | `string` | No | CSS Color |
| `css_greens` | `string` | No | CSS Color |
| `css_light_font_color` | `string` | No | CSS Color |
| `css_link_color` | `string` | No | CSS Color |
| `css_no_data` | `string` | No | CSS Color |
| `css_oranges` | `string` | No | CSS Color |
| `css_reds` | `string` | No | CSS Color |
| `css_yellows` | `string` | No | CSS Color |
| `domain` | `string` | No | CNAME alias for your status page |
| `email_logo` | `string` | No |  |
| `favicon_logo` | `string` | No |  |
| `headline` | `string` | No |  |
| `hero_cover` | `string` | No |  |
| `hidden_from_search` | `bool` | No | Should your page hide itself from search engines |
| `id` | `string` | No | Page identifier |
| `ip_restrictions` | `string` | No |  |
| `name` | `string` | No | Name of your page to be displayed |
| `notifications_email_footer` | `string` | No | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `string` | No | Allows you to customize the email address your page notifications come from |
| `page` | `array` | No |  |
| `page_description` | `string` | No |  |
| `state` | `string` | No |  |
| `subdomain` | `string` | No | Subdomain at which to access your status page |
| `support_url` | `string` | No |  |
| `time_zone` | `string` | No | Timezone configured for your page |
| `transactional_logo` | `string` | No |  |
| `twitter_logo` | `string` | No |  |
| `twitter_username` | `string` | No |  |
| `updated_at` | `string` | No | Timestamp the record was last updated |
| `url` | `string` | No | Website of your page. |
| `viewers_must_be_team_members` | `bool` | No |  |

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
| `component_ids` | `array` | No | List of components codes to set on the page access group |
| `created_at` | `string` | No |  |
| `external_identifier` | `string` | No | Associates group with external group. |
| `id` | `string` | No | Page Access Group Identifier |
| `metric_ids` | `array` | No |  |
| `name` | `string` | No | Name for this Group. |
| `page_access_group` | `array` | No |  |
| `page_access_user_ids` | `array` | No |  |
| `page_id` | `string` | No | Page Identifier. |
| `updated_at` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `component_ids` | - | - | Yes | - | - |
| `created_at` | - | - | - | - | - |
| `external_identifier` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `metric_ids` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `page_access_group` | - | - | - | - | - |
| `page_access_user_ids` | - | - | - | - | - |
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
| `component_ids` | `array` | Yes | List of component codes to allow access to |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `external_login` | `string` | No | IDP login user id. |
| `id` | `string` | No | Page Access User Identifier |
| `metric_ids` | `array` | Yes | List of metrics to add |
| `page_access_group_id` | `string` | No |  |
| `page_access_group_ids` | `string` | No |  |
| `page_access_user` | `array` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->PageAccessUser()->create([
  "id" => null, // string
  "component_ids" => null, // array
  "metric_ids" => null, // array
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
| `pages` | `array` | No | Pages accessible by the user. |
| `user_id` | `string` | No | User identifier |

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
| `body` | `string` | No | Postmortem body |
| `body_draft` | `string` | No | Body draft |
| `body_draft_updated_at` | `string` | No |  |
| `body_updated_at` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `bool` | No | Should email subscribers be notified. |
| `notify_twitter` | `bool` | No | Should Twitter followers be notified. |
| `postmortem` | `array` | Yes |  |
| `preview_key` | `string` | No | Preview Key |
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
| `notify_subscribers` | - | - |
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
| `incident_background_color` | `string` | No | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `string` | No | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `string` | No | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `string` | No | Color of status embed iframe text when displaying maintenance |
| `page_id` | `string` | No | Page identifier |
| `position` | `string` | No | Corner where status embed iframe will appear on page |
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
| `component_ids` | `array` | No | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `string` | No | The components for which the subscriber has elected to receive updates. |
| `created_at` | `string` | No |  |
| `display_phone_number` | `string` | No | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `string` | No | The email address to use to contact the subscriber. |
| `endpoint` | `string` | No | The URL where a webhook subscriber elects to receive updates. |
| `id` | `string` | No | Subscriber Identifier |
| `integration_partner` | `int` | No | The number of integration partners found by the query. |
| `mode` | `string` | No | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `string` | No | Obfuscated slack channel name |
| `page_access_user_id` | `string` | No | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `string` | No | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `string` | No | The phone number used to contact an SMS subscriber |
| `purge_at` | `string` | No | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `string` | No | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `bool` | No | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `bool` | No | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `int` | No | The number of Slack subscribers found by the query. |
| `sms` | `int` | No | The number of Webhook subscribers found by the query. |
| `state` | `string` | No | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `array` | No |  |
| `subscribers` | `string` | Yes | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `int` | No | The number of MS teams subscribers found by the query. |
| `type` | `string` | No | If this is present, only reactivate subscribers of this type. |
| `webhook` | `int` | No | The number of SMS subscribers found by the query. |
| `workspace_name` | `string` | No | The workspace name of the slack subscriber. |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->Subscriber()->create([
  "page_id" => null, // string
  "subscribers" => null, // string
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
| `email` | `string` | No | Email address for the team member |
| `first_name` | `string` | No |  |
| `id` | `string` | No | User identifier |
| `last_name` | `string` | No |  |
| `organization_id` | `string` | No | Organization identifier |
| `updated_at` | `string` | No |  |
| `user` | `array` | Yes |  |

### Operations

#### `create(array $reqdata, ?array $ctrl = null): mixed`

Create a new entity with the given data. Throws on error.

```php
$result = $client->User()->create([
  "organization_id" => null, // string
  "user" => null, // array
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

