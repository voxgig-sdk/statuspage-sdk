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
| `automation_email` | `string` | No | Requires a special feature flag to be enabled |
| `component` | `map[string]any` | No |  |
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
| `component_id` | `string` | No | Component identifier |
| `incidents` | `map[string]any` | No | Related incidents |

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
| `component_group` | `map[string]any` | Yes |  |
| `components` | `string` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No | Description of the component group. |
| `id` | `string` | No | Component Group Identifier |
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
    "component_group": map[string]any{},
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
| `auto_transition_deliver_notifications_at_end` | `bool` | No | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `bool` | No | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `bool` | No | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `bool` | No | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `[]any` | No | Incident components |
| `created_at` | `string` | No | The timestamp when the incident was created at. |
| `id` | `string` | No | Incident Identifier |
| `impact` | `string` | No | The impact of the incident. |
| `impact_override` | `string` | No | value to override calculated impact value |
| `incident` | `map[string]any` | Yes |  |
| `incident_updates` | `[]any` | No | The incident updates for incident. |
| `metadata` | `map[string]any` | No | Metadata attached to the incident. |
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
    "incident": map[string]any{},
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
| `body` | `string` | No | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `[]any` | No | Affected components |
| `group_id` | `string` | No | Identifier of Template Group this template belongs to |
| `id` | `string` | No | Incident Template Identifier |
| `name` | `string` | No | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `bool` | No | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `bool` | No | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `map[string]any` | Yes |  |
| `title` | `string` | No | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `string` | No | The status the incident or maintenance should transition to when selecting this template |

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
    "template": map[string]any{},
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
| `affected_components` | `[]any` | No | Affected components associated with the incident update. |
| `body` | `string` | No | Incident update body. |
| `created_at` | `string` | No | The timestamp when the incident update was created at. |
| `custom_tweet` | `string` | No | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `bool` | No | Controls whether to delivery notifications. |
| `display_at` | `string` | No | Timestamp when incident update is happened. |
| `id` | `string` | No | Incident Update Identifier. |
| `incident_id` | `string` | No | Incident Identifier. |
| `incident_update` | `map[string]any` | No |  |
| `status` | `string` | No | The incident status. |
| `tweet_id` | `string` | No | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `string` | No | The timestamp when twitter updated at. |
| `updated_at` | `string` | No | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `bool` | No | Controls whether to create twitter update. |

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
| `data` | `map[string]any` | Yes | Add data points to metrics |
| `decimal_places` | `int` | No |  |
| `display` | `bool` | No | Should the metric be displayed |
| `id` | `string` | No | Metric identifier |
| `last_fetched_at` | `string` | No |  |
| `metric` | `map[string]any` | No |  |
| `metric_identifier` | `string` | No | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `string` | No | Metric Provider identifier |
| `most_recent_data_at` | `string` | No |  |
| `name` | `string` | No | Name of metric |
| `reference_name` | `string` | No |  |
| `suffix` | `string` | No | Suffix to describe the units on the graph |
| `tooltip_description` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `y_axis_hidden` | `bool` | No | Should the values on the y axis be hidden on render |
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
    "data": map[string]any{},
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
| `id` | `string` | No | Identifier for Metrics Provider |
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
| `page` | `map[string]any` | No |  |
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
| `component_ids` | `[]any` | No | List of components codes to set on the page access group |
| `created_at` | `string` | No |  |
| `external_identifier` | `string` | No | Associates group with external group. |
| `id` | `string` | No | Page Access Group Identifier |
| `metric_ids` | `[]any` | No |  |
| `name` | `string` | No | Name for this Group. |
| `page_access_group` | `map[string]any` | No |  |
| `page_access_user_ids` | `[]any` | No |  |
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
| `component_ids` | `[]any` | Yes | List of component codes to allow access to |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `external_login` | `string` | No | IDP login user id. |
| `id` | `string` | No | Page Access User Identifier |
| `metric_ids` | `[]any` | Yes | List of metrics to add |
| `page_access_group_id` | `string` | No |  |
| `page_access_group_ids` | `string` | No |  |
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
    "component_ids": []any{},
    "metric_ids": []any{},
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
| `pages` | `map[string]any` | No | Pages accessible by the user. |
| `user_id` | `string` | No | User identifier |

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
| `body` | `string` | No | Postmortem body |
| `body_draft` | `string` | No | Body draft |
| `body_draft_updated_at` | `string` | No |  |
| `body_updated_at` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `bool` | No | Should email subscribers be notified. |
| `notify_twitter` | `bool` | No | Should Twitter followers be notified. |
| `postmortem` | `map[string]any` | Yes |  |
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
| `incident_background_color` | `string` | No | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `string` | No | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `string` | No | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `string` | No | Color of status embed iframe text when displaying maintenance |
| `page_id` | `string` | No | Page identifier |
| `position` | `string` | No | Corner where status embed iframe will appear on page |
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
| `component_ids` | `[]any` | No | A list of component ids for which the subscriber should recieve updates for. |
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
| `subscriber` | `map[string]any` | No |  |
| `subscribers` | `string` | Yes | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `int` | No | The number of MS teams subscribers found by the query. |
| `type` | `string` | No | If this is present, only reactivate subscribers of this type. |
| `webhook` | `int` | No | The number of SMS subscribers found by the query. |
| `workspace_name` | `string` | No | The workspace name of the slack subscriber. |

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
    "subscribers": "example_subscribers",
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
| `email` | `string` | No | Email address for the team member |
| `first_name` | `string` | No |  |
| `id` | `string` | No | User identifier |
| `last_name` | `string` | No |  |
| `organization_id` | `string` | No | Organization identifier |
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
    "user": map[string]any{},
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

