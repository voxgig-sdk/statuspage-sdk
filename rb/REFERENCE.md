# Statuspage Ruby SDK Reference

Complete API reference for the Statuspage Ruby SDK.


## StatuspageSDK

### Constructor

```ruby
require_relative 'Statuspage_sdk'

client = StatuspageSDK.new(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `Hash` | SDK configuration options. |
| `options["apikey"]` | `String` | API key for authentication. |
| `options["base"]` | `String` | Base URL for API requests. |
| `options["prefix"]` | `String` | URL prefix appended after base. |
| `options["suffix"]` | `String` | URL suffix appended after path. |
| `options["headers"]` | `Hash` | Custom headers for all requests. |
| `options["feature"]` | `Hash` | Feature configuration. |
| `options["system"]` | `Hash` | System overrides (e.g. custom fetch). |


### Static Methods

#### `StatuspageSDK.test(testopts = nil, sdkopts = nil)`

Create a test client with mock features active. Both arguments may be `nil`.

```ruby
client = StatuspageSDK.test
```


### Instance Methods

#### `Component(data = nil)`

Create a new `Component` entity instance. Pass `nil` for no initial data.

#### `ComponentGroupUptime(data = nil)`

Create a new `ComponentGroupUptime` entity instance. Pass `nil` for no initial data.

#### `GroupComponent(data = nil)`

Create a new `GroupComponent` entity instance. Pass `nil` for no initial data.

#### `Incident(data = nil)`

Create a new `Incident` entity instance. Pass `nil` for no initial data.

#### `IncidentPostmortem(data = nil)`

Create a new `IncidentPostmortem` entity instance. Pass `nil` for no initial data.

#### `IncidentSubscriber(data = nil)`

Create a new `IncidentSubscriber` entity instance. Pass `nil` for no initial data.

#### `IncidentTemplate(data = nil)`

Create a new `IncidentTemplate` entity instance. Pass `nil` for no initial data.

#### `IncidentUpdate(data = nil)`

Create a new `IncidentUpdate` entity instance. Pass `nil` for no initial data.

#### `Metric(data = nil)`

Create a new `Metric` entity instance. Pass `nil` for no initial data.

#### `MetricsProvider(data = nil)`

Create a new `MetricsProvider` entity instance. Pass `nil` for no initial data.

#### `Page(data = nil)`

Create a new `Page` entity instance. Pass `nil` for no initial data.

#### `PageAccessGroup(data = nil)`

Create a new `PageAccessGroup` entity instance. Pass `nil` for no initial data.

#### `PageAccessUser(data = nil)`

Create a new `PageAccessUser` entity instance. Pass `nil` for no initial data.

#### `Permission(data = nil)`

Create a new `Permission` entity instance. Pass `nil` for no initial data.

#### `Postmortem(data = nil)`

Create a new `Postmortem` entity instance. Pass `nil` for no initial data.

#### `StatusEmbedConfig(data = nil)`

Create a new `StatusEmbedConfig` entity instance. Pass `nil` for no initial data.

#### `Subscriber(data = nil)`

Create a new `Subscriber` entity instance. Pass `nil` for no initial data.

#### `User(data = nil)`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `options_map -> Hash`

Return a deep copy of the current SDK options.

#### `get_utility -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs = {}) -> Hash`

Make a direct HTTP request to any API endpoint. Returns a result hash
(`{ "ok" => ..., "status" => ..., "data" => ..., "err" => ... }`); it
does not raise — inspect `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `String` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `String` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `Hash` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `Hash` | Query string parameters. |
| `fetchargs["headers"]` | `Hash` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (hashes are JSON-serialized). |
| `fetchargs["ctrl"]` | `Hash` | Control options (e.g. `{ "explain" => true }`). |

**Returns:** `Hash`

#### `prepare(fetchargs = {}) -> Hash`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`. Raises on error.

**Returns:** `Hash` (the fetch definition; raises on error)


---

## ComponentEntity

```ruby
component = client.Component
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `automation_email` | `String` | No |  |
| `component` | `Hash` | No |  |
| `created_at` | `String` | No |  |
| `description` | `String` | No |  |
| `group` | `Boolean` | No |  |
| `group_id` | `String` | No |  |
| `id` | `String` | No |  |
| `major_outage` | `Integer` | No |  |
| `name` | `String` | No |  |
| `only_show_if_degraded` | `Boolean` | No |  |
| `page_id` | `String` | No |  |
| `partial_outage` | `Integer` | No |  |
| `position` | `Integer` | No |  |
| `range_end` | `String` | No |  |
| `range_start` | `String` | No |  |
| `related_event` | `Hash` | No |  |
| `showcase` | `Boolean` | No |  |
| `start_date` | `String` | No |  |
| `status` | `String` | No |  |
| `updated_at` | `String` | No |  |
| `uptime_percentage` | `Float` | No |  |
| `warning` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Component.create({
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Component.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Component.load({ "id" => "component_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Component.remove({ "id" => "component_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Component.update({
  "id" => "component_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ComponentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## ComponentGroupUptimeEntity

```ruby
component_group_uptime = client.ComponentGroupUptime
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `String` | No |  |
| `major_outage` | `Integer` | No |  |
| `name` | `String` | No |  |
| `partial_outage` | `Integer` | No |  |
| `range_end` | `String` | No |  |
| `range_start` | `String` | No |  |
| `related_event` | `Hash` | No |  |
| `uptime_percentage` | `Float` | No |  |
| `warning` | `String` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.ComponentGroupUptime.load({ "id" => "component_group_uptime_id", "page_id" => "page_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `ComponentGroupUptimeEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## GroupComponentEntity

```ruby
group_component = client.GroupComponent
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `String` | No |  |
| `component_group` | `Hash` | Yes |  |
| `created_at` | `String` | No |  |
| `description` | `String` | No |  |
| `id` | `String` | No |  |
| `name` | `String` | No |  |
| `page_id` | `String` | No |  |
| `position` | `String` | No |  |
| `updated_at` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.GroupComponent.create({
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.GroupComponent.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.GroupComponent.load({ "id" => "group_component_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.GroupComponent.remove({ "id" => "group_component_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.GroupComponent.update({
  "id" => "group_component_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `GroupComponentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IncidentEntity

```ruby
incident = client.Incident
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `Boolean` | No |  |
| `auto_transition_deliver_notifications_at_start` | `Boolean` | No |  |
| `auto_transition_to_maintenance_state` | `Boolean` | No |  |
| `auto_transition_to_operational_state` | `Boolean` | No |  |
| `component` | `Array` | No |  |
| `created_at` | `String` | No |  |
| `id` | `String` | No |  |
| `impact` | `String` | No |  |
| `impact_override` | `String` | No |  |
| `incident` | `Hash` | Yes |  |
| `incident_impact` | `Array` | No |  |
| `incident_update` | `Array` | No |  |
| `metadata` | `Hash` | No |  |
| `monitoring_at` | `String` | No |  |
| `name` | `String` | No |  |
| `page_id` | `String` | No |  |
| `postmortem_body` | `String` | No |  |
| `postmortem_body_last_updated_at` | `String` | No |  |
| `postmortem_ignored` | `Boolean` | No |  |
| `postmortem_notified_subscriber` | `Boolean` | No |  |
| `postmortem_notified_twitter` | `Boolean` | No |  |
| `postmortem_published_at` | `Boolean` | No |  |
| `reminder_interval` | `String` | No |  |
| `resolved_at` | `String` | No |  |
| `scheduled_auto_completed` | `Boolean` | No |  |
| `scheduled_auto_in_progress` | `Boolean` | No |  |
| `scheduled_for` | `String` | No |  |
| `scheduled_remind_prior` | `Boolean` | No |  |
| `scheduled_reminded_at` | `String` | No |  |
| `scheduled_until` | `String` | No |  |
| `shortlink` | `String` | No |  |
| `status` | `String` | No |  |
| `updated_at` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Incident.create({
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Incident.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Incident.load({ "id" => "incident_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Incident.remove({ "id" => "incident_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Incident.update({
  "id" => "incident_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IncidentEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IncidentPostmortemEntity

```ruby
incident_postmortem = client.IncidentPostmortem
```

### Operations

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.IncidentPostmortem.remove({ "id" => "id", "page_id" => "page_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IncidentPostmortemEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IncidentSubscriberEntity

```ruby
incident_subscriber = client.IncidentSubscriber
```

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.IncidentSubscriber.create({
  "incident_id" => "example_incident_id", # String
  "page_id" => "example_page_id", # String
  "subscriber_id" => "example_subscriber_id", # String
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IncidentSubscriberEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IncidentTemplateEntity

```ruby
incident_template = client.IncidentTemplate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `String` | No |  |
| `component` | `Array` | No |  |
| `group_id` | `String` | No |  |
| `id` | `String` | No |  |
| `name` | `String` | No |  |
| `should_send_notification` | `Boolean` | No |  |
| `should_tweet` | `Boolean` | No |  |
| `template` | `Hash` | Yes |  |
| `title` | `String` | No |  |
| `update_status` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.IncidentTemplate.create({
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.IncidentTemplate.list
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IncidentTemplateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## IncidentUpdateEntity

```ruby
incident_update = client.IncidentUpdate
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `affected_component` | `Array` | No |  |
| `body` | `String` | No |  |
| `created_at` | `String` | No |  |
| `custom_tweet` | `String` | No |  |
| `deliver_notification` | `Boolean` | No |  |
| `display_at` | `String` | No |  |
| `id` | `String` | No |  |
| `incident_id` | `String` | No |  |
| `incident_update` | `Hash` | No |  |
| `status` | `String` | No |  |
| `tweet_id` | `String` | No |  |
| `twitter_updated_at` | `String` | No |  |
| `updated_at` | `String` | No |  |
| `wants_twitter_update` | `Boolean` | No |  |

### Operations

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.IncidentUpdate.update({
  "id" => "id",
  "incident_id" => "incident_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `IncidentUpdateEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MetricEntity

```ruby
metric = client.Metric
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `backfill_percentage` | `Integer` | No |  |
| `backfilled` | `Boolean` | No |  |
| `created_at` | `String` | No |  |
| `data` | `Hash` | Yes |  |
| `decimal_place` | `Integer` | No |  |
| `display` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `last_fetched_at` | `String` | No |  |
| `metric` | `Hash` | No |  |
| `metric_identifier` | `String` | No |  |
| `metrics_provider_id` | `String` | No |  |
| `most_recent_data_at` | `String` | No |  |
| `name` | `String` | No |  |
| `reference_name` | `String` | No |  |
| `suffix` | `String` | No |  |
| `tooltip_description` | `String` | No |  |
| `updated_at` | `String` | No |  |
| `y_axis_hidden` | `Boolean` | No |  |
| `y_axis_max` | `Float` | No |  |
| `y_axis_min` | `Float` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Metric.create({
  "metrics_provider_id" => "example_metrics_provider_id", # String
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Metric.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Metric.load({ "id" => "metric_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Metric.remove({ "id" => "metric_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Metric.update({
  "id" => "metric_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MetricEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## MetricsProviderEntity

```ruby
metrics_provider = client.MetricsProvider
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `String` | No |  |
| `disabled` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `last_revalidated_at` | `String` | No |  |
| `metric_base_uri` | `String` | No |  |
| `metrics_provider` | `Hash` | No |  |
| `page_id` | `Integer` | No |  |
| `type` | `String` | No |  |
| `updated_at` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.MetricsProvider.create({
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.MetricsProvider.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.MetricsProvider.load({ "id" => "metrics_provider_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.MetricsProvider.remove({ "id" => "metrics_provider_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.MetricsProvider.update({
  "id" => "metrics_provider_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `MetricsProviderEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PageEntity

```ruby
page = client.Page
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activity_score` | `Float` | No |  |
| `allow_email_subscriber` | `Boolean` | No |  |
| `allow_incident_subscriber` | `Boolean` | No |  |
| `allow_page_subscriber` | `Boolean` | No |  |
| `allow_rss_atom_feed` | `Boolean` | No |  |
| `allow_sms_subscriber` | `Boolean` | No |  |
| `allow_webhook_subscriber` | `Boolean` | No |  |
| `branding` | `String` | No |  |
| `city` | `String` | No |  |
| `country` | `String` | No |  |
| `created_at` | `String` | No |  |
| `css_blue` | `String` | No |  |
| `css_body_background_color` | `String` | No |  |
| `css_border_color` | `String` | No |  |
| `css_font_color` | `String` | No |  |
| `css_graph_color` | `String` | No |  |
| `css_green` | `String` | No |  |
| `css_light_font_color` | `String` | No |  |
| `css_link_color` | `String` | No |  |
| `css_no_data` | `String` | No |  |
| `css_orange` | `String` | No |  |
| `css_red` | `String` | No |  |
| `css_yellow` | `String` | No |  |
| `domain` | `String` | No |  |
| `email_logo` | `String` | No |  |
| `favicon_logo` | `String` | No |  |
| `headline` | `String` | No |  |
| `hero_cover` | `String` | No |  |
| `hidden_from_search` | `Boolean` | No |  |
| `id` | `String` | No |  |
| `ip_restriction` | `String` | No |  |
| `name` | `String` | No |  |
| `notifications_email_footer` | `String` | No |  |
| `notifications_from_email` | `String` | No |  |
| `page` | `Hash` | No |  |
| `page_description` | `String` | No |  |
| `state` | `String` | No |  |
| `subdomain` | `String` | No |  |
| `support_url` | `String` | No |  |
| `time_zone` | `String` | No |  |
| `transactional_logo` | `String` | No |  |
| `twitter_logo` | `String` | No |  |
| `twitter_username` | `String` | No |  |
| `updated_at` | `String` | No |  |
| `url` | `String` | No |  |
| `viewers_must_be_team_member` | `Boolean` | No |  |

### Operations

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Page.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Page.load({ "id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Page.update({
  "id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PageEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PageAccessGroupEntity

```ruby
page_access_group = client.PageAccessGroup
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `Array` | No |  |
| `created_at` | `String` | No |  |
| `external_identifier` | `String` | No |  |
| `id` | `String` | No |  |
| `metric_id` | `Array` | No |  |
| `name` | `String` | No |  |
| `page_access_group` | `Hash` | No |  |
| `page_access_user_id` | `Array` | No |  |
| `page_id` | `String` | No |  |
| `updated_at` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PageAccessGroup.create({
  "id" => "example_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PageAccessGroup.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PageAccessGroup.load({ "id" => "page_access_group_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PageAccessGroup.remove({ "id" => "page_access_group_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.PageAccessGroup.update({
  "id" => "page_access_group_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PageAccessGroupEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PageAccessUserEntity

```ruby
page_access_user = client.PageAccessUser
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `Array` | Yes |  |
| `created_at` | `String` | No |  |
| `email` | `String` | No |  |
| `external_login` | `String` | No |  |
| `id` | `String` | No |  |
| `metric_id` | `Array` | Yes |  |
| `page_access_group_id` | `String` | No |  |
| `page_access_user` | `Hash` | No |  |
| `page_id` | `String` | No |  |
| `updated_at` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PageAccessUser.create({
  "id" => "example_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.PageAccessUser.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.PageAccessUser.load({ "id" => "page_access_user_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.PageAccessUser.remove({ "id" => "page_access_user_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.PageAccessUser.update({
  "id" => "page_access_user_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PageAccessUserEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PermissionEntity

```ruby
permission = client.Permission
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `Hash` | No |  |
| `page` | `Hash` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Permission.load({ "id" => "permission_id", "organization_id" => "organization_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Permission.update({
  "id" => "permission_id",
  "organization_id" => "organization_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PermissionEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## PostmortemEntity

```ruby
postmortem = client.Postmortem
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `String` | No |  |
| `body_draft` | `String` | No |  |
| `body_draft_updated_at` | `String` | No |  |
| `body_updated_at` | `String` | No |  |
| `created_at` | `String` | No |  |
| `custom_tweet` | `String` | No |  |
| `notify_subscriber` | `Boolean` | No |  |
| `notify_twitter` | `Boolean` | No |  |
| `postmortem` | `Hash` | Yes |  |
| `preview_key` | `String` | No |  |
| `published_at` | `String` | No |  |
| `updated_at` | `String` | No |  |

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

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Postmortem.load({ "incident_id" => "incident_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Postmortem.update({
  "incident_id" => "incident_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `PostmortemEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## StatusEmbedConfigEntity

```ruby
status_embed_config = client.StatusEmbedConfig
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `incident_background_color` | `String` | No |  |
| `incident_text_color` | `String` | No |  |
| `maintenance_background_color` | `String` | No |  |
| `maintenance_text_color` | `String` | No |  |
| `page_id` | `String` | No |  |
| `position` | `String` | No |  |
| `status_embed_config` | `Hash` | No |  |

### Operations

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.StatusEmbedConfig.load({ "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.StatusEmbedConfig.update({
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `StatusEmbedConfigEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## SubscriberEntity

```ruby
subscriber = client.Subscriber
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `String` | No |  |
| `component_id` | `Array` | No |  |
| `created_at` | `String` | No |  |
| `display_phone_number` | `String` | No |  |
| `email` | `String` | No |  |
| `endpoint` | `String` | No |  |
| `id` | `String` | No |  |
| `integration_partner` | `Integer` | No |  |
| `mode` | `String` | No |  |
| `obfuscated_channel_name` | `String` | No |  |
| `page_access_user_id` | `String` | No |  |
| `phone_country` | `String` | No |  |
| `phone_number` | `String` | No |  |
| `purge_at` | `String` | No |  |
| `quarantined_at` | `String` | No |  |
| `skip_confirmation_notification` | `Boolean` | No |  |
| `skip_unsubscription_notification` | `Boolean` | No |  |
| `slack` | `Integer` | No |  |
| `sms` | `Integer` | No |  |
| `state` | `String` | No |  |
| `subscriber` | `Hash` | No |  |
| `team` | `Integer` | No |  |
| `type` | `String` | No |  |
| `webhook` | `Integer` | No |  |
| `workspace_name` | `String` | No |  |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Subscriber.create({
  "page_id" => "example_page_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.Subscriber.list
```

#### `load(reqmatch, ctrl = nil) -> result`

Load a single entity matching the given criteria. Raises on error.

```ruby
result = client.Subscriber.load({ "id" => "subscriber_id", "page_id" => "page_id" })
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.Subscriber.remove({ "id" => "subscriber_id", "page_id" => "page_id" })
```

#### `update(reqdata, ctrl = nil) -> result`

Update an existing entity. The data must include the entity `id`. Raises on error.

```ruby
result = client.Subscriber.update({
  "id" => "subscriber_id",
  "page_id" => "page_id",
  # Fields to update
})
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `SubscriberEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## UserEntity

```ruby
user = client.User
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `String` | No |  |
| `email` | `String` | No |  |
| `first_name` | `String` | No |  |
| `id` | `String` | No |  |
| `last_name` | `String` | No |  |
| `organization_id` | `String` | No |  |
| `updated_at` | `String` | No |  |
| `user` | `Hash` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.User.create({
  "organization_id" => "example_organization_id", # String
})
```

#### `list(reqmatch = nil, ctrl = nil) -> Array`

List entities matching the given criteria (call with no argument to list all). Returns an array. Raises on error.

```ruby
results = client.User.list
```

#### `remove(reqmatch, ctrl = nil) -> result`

Remove the entity matching the given criteria. Raises on error.

```ruby
result = client.User.remove({ "id" => "id", "organization_id" => "organization_id" })
```

### Common Methods

#### `data_get -> Hash`

Get the entity data. Returns a copy of the current data.

#### `data_set(data)`

Set the entity data.

#### `match_get -> Hash`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make -> Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `get_name -> String`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ruby
client = StatuspageSDK.new({
  "feature" => {
    "test" => { "active" => true },
  },
})
```

