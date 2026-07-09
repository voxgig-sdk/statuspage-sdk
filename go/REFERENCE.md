# Statuspage Golang SDK Reference

Complete API reference for the Statuspage Golang SDK.


## StatuspageSDK

### Constructor

```go
func NewStatuspageSDK(options map[string]any) *StatuspageSDK
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `map[string]any` | SDK configuration options. |
| `options["apikey"]` | `string` | API key for authentication. |
| `options["base"]` | `string` | Base URL for API requests. |
| `options["prefix"]` | `string` | URL prefix appended after base. |
| `options["suffix"]` | `string` | URL suffix appended after path. |
| `options["headers"]` | `map[string]any` | Custom headers for all requests. |
| `options["feature"]` | `map[string]any` | Feature configuration. |
| `options["system"]` | `map[string]any` | System overrides (e.g. custom fetch). |


### Static Methods

#### `Test() *StatuspageSDK`

No-arg convenience constructor for the common no-options test case.

```go
client := sdk.Test()
```

#### `TestSDK(testopts, sdkopts map[string]any) *StatuspageSDK`

Test client with options. Both arguments may be `nil`.

```go
client := sdk.TestSDK(testopts, sdkopts)
```


### Instance Methods

#### `Component(data map[string]any) StatuspageEntity`

Create a new `Component` entity instance. Pass `nil` for no initial data.

#### `ComponentGroupUptime(data map[string]any) StatuspageEntity`

Create a new `ComponentGroupUptime` entity instance. Pass `nil` for no initial data.

#### `GroupComponent(data map[string]any) StatuspageEntity`

Create a new `GroupComponent` entity instance. Pass `nil` for no initial data.

#### `Incident(data map[string]any) StatuspageEntity`

Create a new `Incident` entity instance. Pass `nil` for no initial data.

#### `IncidentPostmortem(data map[string]any) StatuspageEntity`

Create a new `IncidentPostmortem` entity instance. Pass `nil` for no initial data.

#### `IncidentSubscriber(data map[string]any) StatuspageEntity`

Create a new `IncidentSubscriber` entity instance. Pass `nil` for no initial data.

#### `IncidentTemplate(data map[string]any) StatuspageEntity`

Create a new `IncidentTemplate` entity instance. Pass `nil` for no initial data.

#### `IncidentUpdate(data map[string]any) StatuspageEntity`

Create a new `IncidentUpdate` entity instance. Pass `nil` for no initial data.

#### `Metric(data map[string]any) StatuspageEntity`

Create a new `Metric` entity instance. Pass `nil` for no initial data.

#### `MetricsProvider(data map[string]any) StatuspageEntity`

Create a new `MetricsProvider` entity instance. Pass `nil` for no initial data.

#### `Page(data map[string]any) StatuspageEntity`

Create a new `Page` entity instance. Pass `nil` for no initial data.

#### `PageAccessGroup(data map[string]any) StatuspageEntity`

Create a new `PageAccessGroup` entity instance. Pass `nil` for no initial data.

#### `PageAccessUser(data map[string]any) StatuspageEntity`

Create a new `PageAccessUser` entity instance. Pass `nil` for no initial data.

#### `Permission(data map[string]any) StatuspageEntity`

Create a new `Permission` entity instance. Pass `nil` for no initial data.

#### `Postmortem(data map[string]any) StatuspageEntity`

Create a new `Postmortem` entity instance. Pass `nil` for no initial data.

#### `StatusEmbedConfig(data map[string]any) StatuspageEntity`

Create a new `StatusEmbedConfig` entity instance. Pass `nil` for no initial data.

#### `Subscriber(data map[string]any) StatuspageEntity`

Create a new `Subscriber` entity instance. Pass `nil` for no initial data.

#### `User(data map[string]any) StatuspageEntity`

Create a new `User` entity instance. Pass `nil` for no initial data.

#### `OptionsMap() map[string]any`

Return a deep copy of the current SDK options.

#### `GetUtility() *Utility`

Return a copy of the SDK utility object.

#### `Direct(fetchargs map[string]any) (map[string]any, error)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `string` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `map[string]any` | Path parameter values for `{param}` substitution. |
| `fetchargs["query"]` | `map[string]any` | Query string parameters. |
| `fetchargs["headers"]` | `map[string]any` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (maps are JSON-serialized). |
| `fetchargs["ctrl"]` | `map[string]any` | Control options (e.g. `map[string]any{"explain": true}`). |

**Returns:** `(map[string]any, error)`

#### `Prepare(fetchargs map[string]any) (map[string]any, error)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `Direct()`.

**Returns:** `(map[string]any, error)`


---

## ComponentEntity

```go
component := client.Component(nil)
fmt.Println(component.GetName()) // "component"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `automation_email` | `string` | No |  |
| `component` | `map[string]any` | No |  |
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
| `related_event` | `map[string]any` | No |  |
| `showcase` | `bool` | No |  |
| `start_date` | `string` | No |  |
| `status` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `uptime_percentage` | `float64` | No |  |
| `warning` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Component(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Component(nil).Load(map[string]any{"id": "component_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Component(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Component(nil).Update(map[string]any{
    "id": "component_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Component(nil).Remove(map[string]any{"id": "component_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ComponentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## ComponentGroupUptimeEntity

```go
componentGroupUptime := client.ComponentGroupUptime(nil)
fmt.Println(componentGroupUptime.GetName()) // "component_group_uptime"
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
| `related_event` | `map[string]any` | No |  |
| `uptime_percentage` | `float64` | No |  |
| `warning` | `string` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.ComponentGroupUptime(nil).Load(map[string]any{"id": "component_group_uptime_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `ComponentGroupUptimeEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## GroupComponentEntity

```go
groupComponent := client.GroupComponent(nil)
fmt.Println(groupComponent.GetName()) // "group_component"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `string` | No |  |
| `component_group` | `map[string]any` | Yes |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `name` | `string` | No |  |
| `page_id` | `string` | No |  |
| `position` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.GroupComponent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.GroupComponent(nil).Load(map[string]any{"id": "group_component_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.GroupComponent(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.GroupComponent(nil).Update(map[string]any{
    "id": "group_component_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.GroupComponent(nil).Remove(map[string]any{"id": "group_component_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `GroupComponentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IncidentEntity

```go
incident := client.Incident(nil)
fmt.Println(incident.GetName()) // "incident"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `bool` | No |  |
| `auto_transition_deliver_notifications_at_start` | `bool` | No |  |
| `auto_transition_to_maintenance_state` | `bool` | No |  |
| `auto_transition_to_operational_state` | `bool` | No |  |
| `component` | `[]any` | No |  |
| `created_at` | `string` | No |  |
| `id` | `string` | No |  |
| `impact` | `string` | No |  |
| `impact_override` | `string` | No |  |
| `incident` | `map[string]any` | Yes |  |
| `incident_impact` | `[]any` | No |  |
| `incident_update` | `[]any` | No |  |
| `metadata` | `map[string]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Incident(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Incident(nil).Load(map[string]any{"id": "incident_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Incident(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Incident(nil).Update(map[string]any{
    "id": "incident_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Incident(nil).Remove(map[string]any{"id": "incident_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IncidentEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IncidentPostmortemEntity

```go
incidentPostmortem := client.IncidentPostmortem(nil)
fmt.Println(incidentPostmortem.GetName()) // "incident_postmortem"
```

### Operations

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.IncidentPostmortem(nil).Remove(map[string]any{"id": "id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IncidentPostmortemEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IncidentSubscriberEntity

```go
incidentSubscriber := client.IncidentSubscriber(nil)
fmt.Println(incidentSubscriber.GetName()) // "incident_subscriber"
```

### Operations

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.IncidentSubscriber(nil).Create(map[string]any{
    "incident_id": "example_incident_id",
    "page_id": "example_page_id",
    "subscriber_id": "example_subscriber_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IncidentSubscriberEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IncidentTemplateEntity

```go
incidentTemplate := client.IncidentTemplate(nil)
fmt.Println(incidentTemplate.GetName()) // "incident_template"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `string` | No |  |
| `component` | `[]any` | No |  |
| `group_id` | `string` | No |  |
| `id` | `string` | No |  |
| `name` | `string` | No |  |
| `should_send_notification` | `bool` | No |  |
| `should_tweet` | `bool` | No |  |
| `template` | `map[string]any` | Yes |  |
| `title` | `string` | No |  |
| `update_status` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.IncidentTemplate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.IncidentTemplate(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IncidentTemplateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## IncidentUpdateEntity

```go
incidentUpdate := client.IncidentUpdate(nil)
fmt.Println(incidentUpdate.GetName()) // "incident_update"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `affected_component` | `[]any` | No |  |
| `body` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No |  |
| `deliver_notification` | `bool` | No |  |
| `display_at` | `string` | No |  |
| `id` | `string` | No |  |
| `incident_id` | `string` | No |  |
| `incident_update` | `map[string]any` | No |  |
| `status` | `string` | No |  |
| `tweet_id` | `string` | No |  |
| `twitter_updated_at` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `wants_twitter_update` | `bool` | No |  |

### Operations

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.IncidentUpdate(nil).Update(map[string]any{
    "id": "id",
    "incident_id": "incident_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `IncidentUpdateEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MetricEntity

```go
metric := client.Metric(nil)
fmt.Println(metric.GetName()) // "metric"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `backfill_percentage` | `int` | No |  |
| `backfilled` | `bool` | No |  |
| `created_at` | `string` | No |  |
| `data` | `map[string]any` | Yes |  |
| `decimal_place` | `int` | No |  |
| `display` | `bool` | No |  |
| `id` | `string` | No |  |
| `last_fetched_at` | `string` | No |  |
| `metric` | `map[string]any` | No |  |
| `metric_identifier` | `string` | No |  |
| `metrics_provider_id` | `string` | No |  |
| `most_recent_data_at` | `string` | No |  |
| `name` | `string` | No |  |
| `reference_name` | `string` | No |  |
| `suffix` | `string` | No |  |
| `tooltip_description` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `y_axis_hidden` | `bool` | No |  |
| `y_axis_max` | `float64` | No |  |
| `y_axis_min` | `float64` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Metric(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Metric(nil).Load(map[string]any{"id": "metric_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Metric(nil).Create(map[string]any{
    "metrics_provider_id": "example_metrics_provider_id",
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Metric(nil).Update(map[string]any{
    "id": "metric_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Metric(nil).Remove(map[string]any{"id": "metric_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MetricEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## MetricsProviderEntity

```go
metricsProvider := client.MetricsProvider(nil)
fmt.Println(metricsProvider.GetName()) // "metrics_provider"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `disabled` | `bool` | No |  |
| `id` | `string` | No |  |
| `last_revalidated_at` | `string` | No |  |
| `metric_base_uri` | `string` | No |  |
| `metrics_provider` | `map[string]any` | No |  |
| `page_id` | `int` | No |  |
| `type` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.MetricsProvider(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.MetricsProvider(nil).Load(map[string]any{"id": "metrics_provider_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.MetricsProvider(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.MetricsProvider(nil).Update(map[string]any{
    "id": "metrics_provider_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.MetricsProvider(nil).Remove(map[string]any{"id": "metrics_provider_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `MetricsProviderEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PageEntity

```go
page := client.Page(nil)
fmt.Println(page.GetName()) // "page"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activity_score` | `float64` | No |  |
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
| `page` | `map[string]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Page(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Page(nil).Load(map[string]any{"id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Page(nil).Update(map[string]any{
    "id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PageEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PageAccessGroupEntity

```go
pageAccessGroup := client.PageAccessGroup(nil)
fmt.Println(pageAccessGroup.GetName()) // "page_access_group"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `[]any` | No |  |
| `created_at` | `string` | No |  |
| `external_identifier` | `string` | No |  |
| `id` | `string` | No |  |
| `metric_id` | `[]any` | No |  |
| `name` | `string` | No |  |
| `page_access_group` | `map[string]any` | No |  |
| `page_access_user_id` | `[]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PageAccessGroup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PageAccessGroup(nil).Load(map[string]any{"id": "page_access_group_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PageAccessGroup(nil).Create(map[string]any{
    "id": "example_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.PageAccessGroup(nil).Update(map[string]any{
    "id": "page_access_group_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PageAccessGroup(nil).Remove(map[string]any{"id": "page_access_group_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PageAccessGroupEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PageAccessUserEntity

```go
pageAccessUser := client.PageAccessUser(nil)
fmt.Println(pageAccessUser.GetName()) // "page_access_user"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `[]any` | Yes |  |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `external_login` | `string` | No |  |
| `id` | `string` | No |  |
| `metric_id` | `[]any` | Yes |  |
| `page_access_group_id` | `string` | No |  |
| `page_access_user` | `map[string]any` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.PageAccessUser(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.PageAccessUser(nil).Load(map[string]any{"id": "page_access_user_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.PageAccessUser(nil).Create(map[string]any{
    "id": "example_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.PageAccessUser(nil).Update(map[string]any{
    "id": "page_access_user_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.PageAccessUser(nil).Remove(map[string]any{"id": "page_access_user_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PageAccessUserEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PermissionEntity

```go
permission := client.Permission(nil)
fmt.Println(permission.GetName()) // "permission"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `map[string]any` | No |  |
| `page` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Permission(nil).Load(map[string]any{"id": "permission_id", "organization_id": "organization_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Permission(nil).Update(map[string]any{
    "id": "permission_id",
    "organization_id": "organization_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PermissionEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## PostmortemEntity

```go
postmortem := client.Postmortem(nil)
fmt.Println(postmortem.GetName()) // "postmortem"
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
| `postmortem` | `map[string]any` | Yes |  |
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

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Postmortem(nil).Load(map[string]any{"incident_id": "incident_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Postmortem(nil).Update(map[string]any{
    "incident_id": "incident_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `PostmortemEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## StatusEmbedConfigEntity

```go
statusEmbedConfig := client.StatusEmbedConfig(nil)
fmt.Println(statusEmbedConfig.GetName()) // "status_embed_config"
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
| `status_embed_config` | `map[string]any` | No |  |

### Operations

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.StatusEmbedConfig(nil).Load(map[string]any{"page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.StatusEmbedConfig(nil).Update(map[string]any{
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `StatusEmbedConfigEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## SubscriberEntity

```go
subscriber := client.Subscriber(nil)
fmt.Println(subscriber.GetName()) // "subscriber"
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `string` | No |  |
| `component_id` | `[]any` | No |  |
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
| `subscriber` | `map[string]any` | No |  |
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

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.Subscriber(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Load(reqmatch, ctrl map[string]any) (any, error)`

Load a single entity matching the given criteria.

```go
result, err := client.Subscriber(nil).Load(map[string]any{"id": "subscriber_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.Subscriber(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Update(reqdata, ctrl map[string]any) (any, error)`

Update an existing entity. The data must include the entity `id`.

```go
result, err := client.Subscriber(nil).Update(map[string]any{
    "id": "subscriber_id",
    "page_id": "page_id",
    // Fields to update
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.Subscriber(nil).Remove(map[string]any{"id": "subscriber_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `SubscriberEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## UserEntity

```go
user := client.User(nil)
fmt.Println(user.GetName()) // "user"
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
| `user` | `map[string]any` | Yes |  |

### Operations

#### `List(reqmatch, ctrl map[string]any) (any, error)`

List entities matching the given criteria. Returns an array.

```go
results, err := client.User(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(results)
```

#### `Create(reqdata, ctrl map[string]any) (any, error)`

Create a new entity with the given data.

```go
result, err := client.User(nil).Create(map[string]any{
    "organization_id": "example_organization_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

#### `Remove(reqmatch, ctrl map[string]any) (any, error)`

Remove the entity matching the given criteria.

```go
result, err := client.User(nil).Remove(map[string]any{"id": "id", "organization_id": "organization_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```

### Common Methods

#### `Data(args ...any) any`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `Match(args ...any) any`

Get or set the entity match criteria. Works the same as `Data()`.

#### `Make() Entity`

Create a new `UserEntity` instance with the same client and
options.

#### `GetName() string`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```go
client := sdk.NewStatuspageSDK(map[string]any{
    "feature": map[string]any{
        "test": map[string]any{"active": true},
    },
})
```

