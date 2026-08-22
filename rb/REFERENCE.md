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
| `automation_email` | `String` | No | Requires a special feature flag to be enabled |
| `component` | `Hash` | No |  |
| `created_at` | `String` | No |  |
| `description` | `String` | No | More detailed description for component |
| `group` | `Boolean` | No | Is this component a group |
| `group_id` | `String` | No | Component Group identifier |
| `id` | `String` | No | Incident identifier |
| `name` | `String` | No | Display name for component |
| `only_show_if_degraded` | `Boolean` | No | Requires a special feature flag to be enabled |
| `page_id` | `String` | No | Page identifier |
| `position` | `Integer` | No | Order the component will appear on the page |
| `showcase` | `Boolean` | No | Should this component be showcased |
| `start_date` | `String` | No | The date this component started being used |
| `status` | `String` | No | Status of component |
| `updated_at` | `String` | No |  |

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
| `component_id` | `String` | No | Component identifier |
| `incidents` | `Hash` | No | Related incidents |

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
| `component_group` | `Hash` | Yes |  |
| `components` | `String` | No |  |
| `created_at` | `String` | No |  |
| `description` | `String` | No | Description of the component group. |
| `id` | `String` | No | Component Group Identifier |
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
  "component_group" => {}, # Hash
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
| `auto_transition_deliver_notifications_at_end` | `Boolean` | No | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `Boolean` | No | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `Boolean` | No | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `Boolean` | No | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `Array` | No | Incident components |
| `created_at` | `String` | No | The timestamp when the incident was created at. |
| `id` | `String` | No | Incident Identifier |
| `impact` | `String` | No | The impact of the incident. |
| `impact_override` | `String` | No | value to override calculated impact value |
| `incident` | `Hash` | Yes |  |
| `incident_updates` | `Array` | No | The incident updates for incident. |
| `metadata` | `Hash` | No | Metadata attached to the incident. |
| `monitoring_at` | `String` | No | The timestamp when incident entered monitoring state. |
| `name` | `String` | No | Incident Name. |
| `page_id` | `String` | No | Incident Page Identifier |
| `postmortem_body` | `String` | No | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `String` | No | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `Boolean` | No | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `Boolean` | No | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `Boolean` | No | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `Boolean` | No | The timestamp when the postmortem was published. |
| `reminder_intervals` | `String` | No | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `String` | No | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `Boolean` | No | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `Boolean` | No | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `String` | No | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `Boolean` | No | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `String` | No | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `String` | No | The timestamp the incident is scheduled until. |
| `shortlink` | `String` | No | Incident Shortlink. |
| `status` | `String` | No | The incident status. |
| `updated_at` | `String` | No | The timestamp when the incident was updated at. |

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

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Incident.create({
  "page_id" => "example_page_id", # String
  "incident" => {}, # Hash
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
| `body` | `String` | No | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `Array` | No | Affected components |
| `group_id` | `String` | No | Identifier of Template Group this template belongs to |
| `id` | `String` | No | Incident Template Identifier |
| `name` | `String` | No | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `Boolean` | No | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `Boolean` | No | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `Hash` | Yes |  |
| `title` | `String` | No | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `String` | No | The status the incident or maintenance should transition to when selecting this template |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.IncidentTemplate.create({
  "page_id" => "example_page_id", # String
  "template" => {}, # Hash
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
| `affected_components` | `Array` | No | Affected components associated with the incident update. |
| `body` | `String` | No | Incident update body. |
| `created_at` | `String` | No | The timestamp when the incident update was created at. |
| `custom_tweet` | `String` | No | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `Boolean` | No | Controls whether to delivery notifications. |
| `display_at` | `String` | No | Timestamp when incident update is happened. |
| `id` | `String` | No | Incident Update Identifier. |
| `incident_id` | `String` | No | Incident Identifier. |
| `incident_update` | `Hash` | No |  |
| `status` | `String` | No | The incident status. |
| `tweet_id` | `String` | No | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `String` | No | The timestamp when twitter updated at. |
| `updated_at` | `String` | No | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `Boolean` | No | Controls whether to create twitter update. |

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
| `data` | `Hash` | Yes | Add data points to metrics |
| `decimal_places` | `Integer` | No |  |
| `display` | `Boolean` | No | Should the metric be displayed |
| `id` | `String` | No | Metric identifier |
| `last_fetched_at` | `String` | No |  |
| `metric` | `Hash` | No |  |
| `metric_identifier` | `String` | No | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `String` | No | Metric Provider identifier |
| `most_recent_data_at` | `String` | No |  |
| `name` | `String` | No | Name of metric |
| `reference_name` | `String` | No |  |
| `suffix` | `String` | No | Suffix to describe the units on the graph |
| `tooltip_description` | `String` | No |  |
| `updated_at` | `String` | No |  |
| `y_axis_hidden` | `Boolean` | No | Should the values on the y axis be hidden on render |
| `y_axis_max` | `Float` | No |  |
| `y_axis_min` | `Float` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Metric.create({
  "metrics_provider_id" => "example_metrics_provider_id", # String
  "page_id" => "example_page_id", # String
  "data" => {}, # Hash
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
| `id` | `String` | No | Identifier for Metrics Provider |
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
| `allow_email_subscribers` | `Boolean` | No | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `Boolean` | No | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `Boolean` | No | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `Boolean` | No | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `Boolean` | No | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `Boolean` | No | Can your users choose to receive notifications via Webhooks |
| `branding` | `String` | No | The main template your statuspage will use |
| `city` | `String` | No |  |
| `country` | `String` | No |  |
| `created_at` | `String` | No | Timestamp the record was created |
| `css_blues` | `String` | No | CSS Color |
| `css_body_background_color` | `String` | No | CSS Color |
| `css_border_color` | `String` | No | CSS Color |
| `css_font_color` | `String` | No | CSS Color |
| `css_graph_color` | `String` | No | CSS Color |
| `css_greens` | `String` | No | CSS Color |
| `css_light_font_color` | `String` | No | CSS Color |
| `css_link_color` | `String` | No | CSS Color |
| `css_no_data` | `String` | No | CSS Color |
| `css_oranges` | `String` | No | CSS Color |
| `css_reds` | `String` | No | CSS Color |
| `css_yellows` | `String` | No | CSS Color |
| `domain` | `String` | No | CNAME alias for your status page |
| `email_logo` | `String` | No |  |
| `favicon_logo` | `String` | No |  |
| `headline` | `String` | No |  |
| `hero_cover` | `String` | No |  |
| `hidden_from_search` | `Boolean` | No | Should your page hide itself from search engines |
| `id` | `String` | No | Page identifier |
| `ip_restrictions` | `String` | No |  |
| `name` | `String` | No | Name of your page to be displayed |
| `notifications_email_footer` | `String` | No | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `String` | No | Allows you to customize the email address your page notifications come from |
| `page` | `Hash` | No |  |
| `page_description` | `String` | No |  |
| `state` | `String` | No |  |
| `subdomain` | `String` | No | Subdomain at which to access your status page |
| `support_url` | `String` | No |  |
| `time_zone` | `String` | No | Timezone configured for your page |
| `transactional_logo` | `String` | No |  |
| `twitter_logo` | `String` | No |  |
| `twitter_username` | `String` | No |  |
| `updated_at` | `String` | No | Timestamp the record was last updated |
| `url` | `String` | No | Website of your page. |
| `viewers_must_be_team_members` | `Boolean` | No |  |

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
| `component_ids` | `Array` | No | List of components codes to set on the page access group |
| `created_at` | `String` | No |  |
| `external_identifier` | `String` | No | Associates group with external group. |
| `id` | `String` | No | Page Access Group Identifier |
| `metric_ids` | `Array` | No |  |
| `name` | `String` | No | Name for this Group. |
| `page_access_group` | `Hash` | No |  |
| `page_access_user_ids` | `Array` | No |  |
| `page_id` | `String` | No | Page Identifier. |
| `updated_at` | `String` | No |  |

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
| `component_ids` | `Array` | Yes | List of component codes to allow access to |
| `created_at` | `String` | No |  |
| `email` | `String` | No |  |
| `external_login` | `String` | No | IDP login user id. |
| `id` | `String` | No | Page Access User Identifier |
| `metric_ids` | `Array` | Yes | List of metrics to add |
| `page_access_group_id` | `String` | No |  |
| `page_access_group_ids` | `String` | No |  |
| `page_access_user` | `Hash` | No |  |
| `page_id` | `String` | No |  |
| `updated_at` | `String` | No |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.PageAccessUser.create({
  "id" => "example_id", # String
  "component_ids" => [], # Array
  "metric_ids" => [], # Array
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
| `pages` | `Hash` | No | Pages accessible by the user. |
| `user_id` | `String` | No | User identifier |

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
| `body` | `String` | No | Postmortem body |
| `body_draft` | `String` | No | Body draft |
| `body_draft_updated_at` | `String` | No |  |
| `body_updated_at` | `String` | No |  |
| `created_at` | `String` | No |  |
| `custom_tweet` | `String` | No | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `Boolean` | No | Should email subscribers be notified. |
| `notify_twitter` | `Boolean` | No | Should Twitter followers be notified. |
| `postmortem` | `Hash` | Yes |  |
| `preview_key` | `String` | No | Preview Key |
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
| `notify_subscribers` | - | - |
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
| `incident_background_color` | `String` | No | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `String` | No | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `String` | No | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `String` | No | Color of status embed iframe text when displaying maintenance |
| `page_id` | `String` | No | Page identifier |
| `position` | `String` | No | Corner where status embed iframe will appear on page |
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
| `component_ids` | `Array` | No | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `String` | No | The components for which the subscriber has elected to receive updates. |
| `created_at` | `String` | No |  |
| `display_phone_number` | `String` | No | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `String` | No | The email address to use to contact the subscriber. |
| `endpoint` | `String` | No | The URL where a webhook subscriber elects to receive updates. |
| `id` | `String` | No | Subscriber Identifier |
| `integration_partner` | `Integer` | No | The number of integration partners found by the query. |
| `mode` | `String` | No | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `String` | No | Obfuscated slack channel name |
| `page_access_user_id` | `String` | No | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `String` | No | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `String` | No | The phone number used to contact an SMS subscriber |
| `purge_at` | `String` | No | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `String` | No | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `Boolean` | No | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `Boolean` | No | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `Integer` | No | The number of Slack subscribers found by the query. |
| `sms` | `Integer` | No | The number of Webhook subscribers found by the query. |
| `state` | `String` | No | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `Hash` | No |  |
| `subscribers` | `String` | Yes | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `Integer` | No | The number of MS teams subscribers found by the query. |
| `type` | `String` | No | If this is present, only reactivate subscribers of this type. |
| `webhook` | `Integer` | No | The number of SMS subscribers found by the query. |
| `workspace_name` | `String` | No | The workspace name of the slack subscriber. |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.Subscriber.create({
  "page_id" => "example_page_id", # String
  "subscribers" => "example_subscribers", # String
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
| `email` | `String` | No | Email address for the team member |
| `first_name` | `String` | No |  |
| `id` | `String` | No | User identifier |
| `last_name` | `String` | No |  |
| `organization_id` | `String` | No | Organization identifier |
| `updated_at` | `String` | No |  |
| `user` | `Hash` | Yes |  |

### Operations

#### `create(reqdata, ctrl = nil) -> result`

Create a new entity with the given data. Raises on error.

```ruby
result = client.User.create({
  "organization_id" => "example_organization_id", # String
  "user" => {}, # Hash
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

