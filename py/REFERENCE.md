# Statuspage Python SDK Reference

Complete API reference for the Statuspage Python SDK.


## StatuspageSDK

### Constructor

```python
from statuspage_sdk import StatuspageSDK

client = StatuspageSDK(options)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `dict` | SDK configuration options. |
| `options["apikey"]` | `str` | API key for authentication. |
| `options["base"]` | `str` | Base URL for API requests. |
| `options["prefix"]` | `str` | URL prefix appended after base. |
| `options["suffix"]` | `str` | URL suffix appended after path. |
| `options["headers"]` | `dict` | Custom headers for all requests. |
| `options["feature"]` | `dict` | Feature configuration. |
| `options["system"]` | `dict` | System overrides (e.g. custom fetch). |


### Static Methods

#### `StatuspageSDK.test(testopts=None, sdkopts=None)`

Create a test client with mock features active. Both arguments may be `None`.

```python
client = StatuspageSDK.test()
```


### Instance Methods

#### `Component(data=None)`

Create a new `ComponentEntity` instance. Pass `None` for no initial data.

#### `ComponentGroupUptime(data=None)`

Create a new `ComponentGroupUptimeEntity` instance. Pass `None` for no initial data.

#### `GroupComponent(data=None)`

Create a new `GroupComponentEntity` instance. Pass `None` for no initial data.

#### `Incident(data=None)`

Create a new `IncidentEntity` instance. Pass `None` for no initial data.

#### `IncidentPostmortem(data=None)`

Create a new `IncidentPostmortemEntity` instance. Pass `None` for no initial data.

#### `IncidentSubscriber(data=None)`

Create a new `IncidentSubscriberEntity` instance. Pass `None` for no initial data.

#### `IncidentTemplate(data=None)`

Create a new `IncidentTemplateEntity` instance. Pass `None` for no initial data.

#### `IncidentUpdate(data=None)`

Create a new `IncidentUpdateEntity` instance. Pass `None` for no initial data.

#### `Metric(data=None)`

Create a new `MetricEntity` instance. Pass `None` for no initial data.

#### `MetricsProvider(data=None)`

Create a new `MetricsProviderEntity` instance. Pass `None` for no initial data.

#### `Page(data=None)`

Create a new `PageEntity` instance. Pass `None` for no initial data.

#### `PageAccessGroup(data=None)`

Create a new `PageAccessGroupEntity` instance. Pass `None` for no initial data.

#### `PageAccessUser(data=None)`

Create a new `PageAccessUserEntity` instance. Pass `None` for no initial data.

#### `Permission(data=None)`

Create a new `PermissionEntity` instance. Pass `None` for no initial data.

#### `Postmortem(data=None)`

Create a new `PostmortemEntity` instance. Pass `None` for no initial data.

#### `StatusEmbedConfig(data=None)`

Create a new `StatusEmbedConfigEntity` instance. Pass `None` for no initial data.

#### `Subscriber(data=None)`

Create a new `SubscriberEntity` instance. Pass `None` for no initial data.

#### `User(data=None)`

Create a new `UserEntity` instance. Pass `None` for no initial data.

#### `options_map() -> dict`

Return a deep copy of the current SDK options.

#### `get_utility() -> Utility`

Return a copy of the SDK utility object.

#### `direct(fetchargs=None) -> dict`

Make a direct HTTP request to any API endpoint. Returns a result `dict` with `ok`, `status`, `headers`, and `data` (or `err` on failure). This escape hatch never raises — branch on `result["ok"]`.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs["path"]` | `str` | URL path with optional `{param}` placeholders. |
| `fetchargs["method"]` | `str` | HTTP method (default: `"GET"`). |
| `fetchargs["params"]` | `dict` | Path parameter values. |
| `fetchargs["query"]` | `dict` | Query string parameters. |
| `fetchargs["headers"]` | `dict` | Request headers (merged with defaults). |
| `fetchargs["body"]` | `any` | Request body (dicts are JSON-serialized). |

**Returns:** `result_dict`

#### `prepare(fetchargs=None) -> dict`

Prepare a fetch definition without sending. Returns the `fetchdef` and raises on error.


---

## ComponentEntity

```python
component = client.Component()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `automation_email` | `str` | No | Requires a special feature flag to be enabled |
| `component` | `dict` | No |  |
| `created_at` | `str` | No |  |
| `description` | `str` | No | More detailed description for component |
| `group` | `bool` | No | Is this component a group |
| `group_id` | `str` | No | Component Group identifier |
| `id` | `str` | No | Incident identifier |
| `name` | `str` | No | Display name for component |
| `only_show_if_degraded` | `bool` | No | Requires a special feature flag to be enabled |
| `page_id` | `str` | No | Page identifier |
| `position` | `int` | No | Order the component will appear on the page |
| `showcase` | `bool` | No | Should this component be showcased |
| `start_date` | `str` | No | The date this component started being used |
| `status` | `str` | No | Status of component |
| `updated_at` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Component().create({
    "page_id": "example_page_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Component().list({"page_id": "example"})
for component in results:
    print(component)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Component().load({"id": "component_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Component().remove({"id": "component_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Component().update({
    "id": "component_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ComponentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## ComponentGroupUptimeEntity

```python
component_group_uptime = client.ComponentGroupUptime()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `str` | No | Component identifier |
| `incidents` | `dict` | No | Related incidents |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ComponentGroupUptime().load({"id": "component_group_uptime_id", "page_id": "page_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `ComponentGroupUptimeEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## GroupComponentEntity

```python
group_component = client.GroupComponent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_group` | `dict` | Yes |  |
| `components` | `str` | No |  |
| `created_at` | `str` | No |  |
| `description` | `str` | No | Description of the component group. |
| `id` | `str` | No | Component Group Identifier |
| `name` | `str` | No |  |
| `page_id` | `str` | No |  |
| `position` | `str` | No |  |
| `updated_at` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GroupComponent().create({
    "page_id": "example_page_id",  # str
    "component_group": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.GroupComponent().list({"page_id": "example"})
for group_component in results:
    print(group_component)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.GroupComponent().load({"id": "group_component_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.GroupComponent().remove({"id": "group_component_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.GroupComponent().update({
    "id": "group_component_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `GroupComponentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IncidentEntity

```python
incident = client.Incident()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `bool` | No | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `bool` | No | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `bool` | No | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `bool` | No | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `list` | No | Incident components |
| `created_at` | `str` | No | The timestamp when the incident was created at. |
| `id` | `str` | No | Incident Identifier |
| `impact` | `str` | No | The impact of the incident. |
| `impact_override` | `str` | No | value to override calculated impact value |
| `incident` | `dict` | Yes |  |
| `incident_updates` | `list` | No | The incident updates for incident. |
| `metadata` | `dict` | No | Metadata attached to the incident. |
| `monitoring_at` | `str` | No | The timestamp when incident entered monitoring state. |
| `name` | `str` | No | Incident Name. |
| `page_id` | `str` | No | Incident Page Identifier |
| `postmortem_body` | `str` | No | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `str` | No | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `bool` | No | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `bool` | No | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `bool` | No | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `bool` | No | The timestamp when the postmortem was published. |
| `reminder_intervals` | `str` | No | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `str` | No | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `bool` | No | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `bool` | No | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `str` | No | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `bool` | No | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `str` | No | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `str` | No | The timestamp the incident is scheduled until. |
| `shortlink` | `str` | No | Incident Shortlink. |
| `status` | `str` | No | The incident status. |
| `updated_at` | `str` | No | The timestamp when the incident was updated at. |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Incident().create({
    "page_id": "example_page_id",  # str
    "incident": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Incident().list({"page_id": "example"})
for incident in results:
    print(incident)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Incident().load({"id": "incident_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Incident().remove({"id": "incident_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Incident().update({
    "id": "incident_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IncidentEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IncidentPostmortemEntity

```python
incident_postmortem = client.IncidentPostmortem()
```

### Operations

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.IncidentPostmortem().remove({"id": "id", "page_id": "page_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IncidentPostmortemEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IncidentSubscriberEntity

```python
incident_subscriber = client.IncidentSubscriber()
```

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.IncidentSubscriber().create({
    "incident_id": "example_incident_id",  # str
    "page_id": "example_page_id",  # str
    "subscriber_id": "example_subscriber_id",  # str
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IncidentSubscriberEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IncidentTemplateEntity

```python
incident_template = client.IncidentTemplate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `str` | No | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `list` | No | Affected components |
| `group_id` | `str` | No | Identifier of Template Group this template belongs to |
| `id` | `str` | No | Incident Template Identifier |
| `name` | `str` | No | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `bool` | No | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `bool` | No | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `dict` | Yes |  |
| `title` | `str` | No | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `str` | No | The status the incident or maintenance should transition to when selecting this template |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.IncidentTemplate().create({
    "page_id": "example_page_id",  # str
    "template": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.IncidentTemplate().list({"page_id": "example"})
for incident_template in results:
    print(incident_template)
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IncidentTemplateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## IncidentUpdateEntity

```python
incident_update = client.IncidentUpdate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `affected_components` | `list` | No | Affected components associated with the incident update. |
| `body` | `str` | No | Incident update body. |
| `created_at` | `str` | No | The timestamp when the incident update was created at. |
| `custom_tweet` | `str` | No | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `bool` | No | Controls whether to delivery notifications. |
| `display_at` | `str` | No | Timestamp when incident update is happened. |
| `id` | `str` | No | Incident Update Identifier. |
| `incident_id` | `str` | No | Incident Identifier. |
| `incident_update` | `dict` | No |  |
| `status` | `str` | No | The incident status. |
| `tweet_id` | `str` | No | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `str` | No | The timestamp when twitter updated at. |
| `updated_at` | `str` | No | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `bool` | No | Controls whether to create twitter update. |

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.IncidentUpdate().update({
    "id": "id",
    "incident_id": "incident_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `IncidentUpdateEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MetricEntity

```python
metric = client.Metric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `backfill_percentage` | `int` | No |  |
| `backfilled` | `bool` | No |  |
| `created_at` | `str` | No |  |
| `data` | `dict` | Yes | Add data points to metrics |
| `decimal_places` | `int` | No |  |
| `display` | `bool` | No | Should the metric be displayed |
| `id` | `str` | No | Metric identifier |
| `last_fetched_at` | `str` | No |  |
| `metric` | `dict` | No |  |
| `metric_identifier` | `str` | No | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `str` | No | Metric Provider identifier |
| `most_recent_data_at` | `str` | No |  |
| `name` | `str` | No | Name of metric |
| `reference_name` | `str` | No |  |
| `suffix` | `str` | No | Suffix to describe the units on the graph |
| `tooltip_description` | `str` | No |  |
| `updated_at` | `str` | No |  |
| `y_axis_hidden` | `bool` | No | Should the values on the y axis be hidden on render |
| `y_axis_max` | `float` | No |  |
| `y_axis_min` | `float` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Metric().create({
    "metrics_provider_id": "example_metrics_provider_id",  # str
    "page_id": "example_page_id",  # str
    "data": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Metric().list({"page_access_user_id": "example", "page_id": "example"})
for metric in results:
    print(metric)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Metric().load({"id": "metric_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Metric().remove({"id": "metric_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Metric().update({
    "id": "metric_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetricEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## MetricsProviderEntity

```python
metrics_provider = client.MetricsProvider()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No |  |
| `disabled` | `bool` | No |  |
| `id` | `str` | No | Identifier for Metrics Provider |
| `last_revalidated_at` | `str` | No |  |
| `metric_base_uri` | `str` | No |  |
| `metrics_provider` | `dict` | No |  |
| `page_id` | `int` | No |  |
| `type` | `str` | No |  |
| `updated_at` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.MetricsProvider().create({
    "page_id": "example_page_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MetricsProvider().list({"page_id": "example"})
for metrics_provider in results:
    print(metrics_provider)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.MetricsProvider().load({"id": "metrics_provider_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.MetricsProvider().remove({"id": "metrics_provider_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.MetricsProvider().update({
    "id": "metrics_provider_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `MetricsProviderEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PageEntity

```python
page = client.Page()
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
| `branding` | `str` | No | The main template your statuspage will use |
| `city` | `str` | No |  |
| `country` | `str` | No |  |
| `created_at` | `str` | No | Timestamp the record was created |
| `css_blues` | `str` | No | CSS Color |
| `css_body_background_color` | `str` | No | CSS Color |
| `css_border_color` | `str` | No | CSS Color |
| `css_font_color` | `str` | No | CSS Color |
| `css_graph_color` | `str` | No | CSS Color |
| `css_greens` | `str` | No | CSS Color |
| `css_light_font_color` | `str` | No | CSS Color |
| `css_link_color` | `str` | No | CSS Color |
| `css_no_data` | `str` | No | CSS Color |
| `css_oranges` | `str` | No | CSS Color |
| `css_reds` | `str` | No | CSS Color |
| `css_yellows` | `str` | No | CSS Color |
| `domain` | `str` | No | CNAME alias for your status page |
| `email_logo` | `str` | No |  |
| `favicon_logo` | `str` | No |  |
| `headline` | `str` | No |  |
| `hero_cover` | `str` | No |  |
| `hidden_from_search` | `bool` | No | Should your page hide itself from search engines |
| `id` | `str` | No | Page identifier |
| `ip_restrictions` | `str` | No |  |
| `name` | `str` | No | Name of your page to be displayed |
| `notifications_email_footer` | `str` | No | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `str` | No | Allows you to customize the email address your page notifications come from |
| `page` | `dict` | No |  |
| `page_description` | `str` | No |  |
| `state` | `str` | No |  |
| `subdomain` | `str` | No | Subdomain at which to access your status page |
| `support_url` | `str` | No |  |
| `time_zone` | `str` | No | Timezone configured for your page |
| `transactional_logo` | `str` | No |  |
| `twitter_logo` | `str` | No |  |
| `twitter_username` | `str` | No |  |
| `updated_at` | `str` | No | Timestamp the record was last updated |
| `url` | `str` | No | Website of your page. |
| `viewers_must_be_team_members` | `bool` | No |  |

### Operations

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Page().list()
for page in results:
    print(page)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Page().load({"id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Page().update({
    "id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PageEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PageAccessGroupEntity

```python
page_access_group = client.PageAccessGroup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_ids` | `list` | No | List of components codes to set on the page access group |
| `created_at` | `str` | No |  |
| `external_identifier` | `str` | No | Associates group with external group. |
| `id` | `str` | No | Page Access Group Identifier |
| `metric_ids` | `list` | No |  |
| `name` | `str` | No | Name for this Group. |
| `page_access_group` | `dict` | No |  |
| `page_access_user_ids` | `list` | No |  |
| `page_id` | `str` | No | Page Identifier. |
| `updated_at` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PageAccessGroup().create({
    "id": "example_id",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PageAccessGroup().list({"id": "example_id"})
for page_access_group in results:
    print(page_access_group)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PageAccessGroup().load({"id": "page_access_group_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PageAccessGroup().remove({"id": "page_access_group_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PageAccessGroup().update({
    "id": "page_access_group_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PageAccessGroupEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PageAccessUserEntity

```python
page_access_user = client.PageAccessUser()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_ids` | `list` | Yes | List of component codes to allow access to |
| `created_at` | `str` | No |  |
| `email` | `str` | No |  |
| `external_login` | `str` | No | IDP login user id. |
| `id` | `str` | No | Page Access User Identifier |
| `metric_ids` | `list` | Yes | List of metrics to add |
| `page_access_group_id` | `str` | No |  |
| `page_access_group_ids` | `str` | No |  |
| `page_access_user` | `dict` | No |  |
| `page_id` | `str` | No |  |
| `updated_at` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PageAccessUser().create({
    "id": "example_id",  # str
    "component_ids": [],  # list
    "metric_ids": [],  # list
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PageAccessUser().list({"id": "example_id"})
for page_access_user in results:
    print(page_access_user)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PageAccessUser().load({"id": "page_access_user_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PageAccessUser().remove({"id": "page_access_user_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PageAccessUser().update({
    "id": "page_access_user_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PageAccessUserEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PermissionEntity

```python
permission = client.Permission()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `pages` | `dict` | No | Pages accessible by the user. |
| `user_id` | `str` | No | User identifier |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Permission().load({"id": "permission_id", "organization_id": "organization_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Permission().update({
    "id": "permission_id",
    "organization_id": "organization_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PermissionEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## PostmortemEntity

```python
postmortem = client.Postmortem()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `str` | No | Postmortem body |
| `body_draft` | `str` | No | Body draft |
| `body_draft_updated_at` | `str` | No |  |
| `body_updated_at` | `str` | No |  |
| `created_at` | `str` | No |  |
| `custom_tweet` | `str` | No | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `bool` | No | Should email subscribers be notified. |
| `notify_twitter` | `bool` | No | Should Twitter followers be notified. |
| `postmortem` | `dict` | Yes |  |
| `preview_key` | `str` | No | Preview Key |
| `published_at` | `str` | No |  |
| `updated_at` | `str` | No |  |

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

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Postmortem().load({"incident_id": "incident_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Postmortem().update({
    "incident_id": "incident_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `PostmortemEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## StatusEmbedConfigEntity

```python
status_embed_config = client.StatusEmbedConfig()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `incident_background_color` | `str` | No | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `str` | No | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `str` | No | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `str` | No | Color of status embed iframe text when displaying maintenance |
| `page_id` | `str` | No | Page identifier |
| `position` | `str` | No | Corner where status embed iframe will appear on page |
| `status_embed_config` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.StatusEmbedConfig().load({"page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.StatusEmbedConfig().update({
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `StatusEmbedConfigEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## SubscriberEntity

```python
subscriber = client.Subscriber()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_ids` | `list` | No | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `str` | No | The components for which the subscriber has elected to receive updates. |
| `created_at` | `str` | No |  |
| `display_phone_number` | `str` | No | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `str` | No | The email address to use to contact the subscriber. |
| `endpoint` | `str` | No | The URL where a webhook subscriber elects to receive updates. |
| `id` | `str` | No | Subscriber Identifier |
| `integration_partner` | `int` | No | The number of integration partners found by the query. |
| `mode` | `str` | No | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `str` | No | Obfuscated slack channel name |
| `page_access_user_id` | `str` | No | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `str` | No | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `str` | No | The phone number used to contact an SMS subscriber |
| `purge_at` | `str` | No | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `str` | No | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `bool` | No | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `bool` | No | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `int` | No | The number of Slack subscribers found by the query. |
| `sms` | `int` | No | The number of Webhook subscribers found by the query. |
| `state` | `str` | No | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `dict` | No |  |
| `subscribers` | `str` | Yes | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `int` | No | The number of MS teams subscribers found by the query. |
| `type` | `str` | No | If this is present, only reactivate subscribers of this type. |
| `webhook` | `int` | No | The number of SMS subscribers found by the query. |
| `workspace_name` | `str` | No | The workspace name of the slack subscriber. |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Subscriber().create({
    "page_id": "example_page_id",  # str
    "subscribers": "example_subscribers",  # str
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Subscriber().list({"page_id": "example"})
for subscriber in results:
    print(subscriber)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Subscriber().load({"id": "subscriber_id", "page_id": "page_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Subscriber().remove({"id": "subscriber_id", "page_id": "page_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Subscriber().update({
    "id": "subscriber_id",
    "page_id": "page_id",
    # Fields to update
})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `SubscriberEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## UserEntity

```python
user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `str` | No |  |
| `email` | `str` | No | Email address for the team member |
| `first_name` | `str` | No |  |
| `id` | `str` | No | User identifier |
| `last_name` | `str` | No |  |
| `organization_id` | `str` | No | Organization identifier |
| `updated_at` | `str` | No |  |
| `user` | `dict` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.User().create({
    "organization_id": "example_organization_id",  # str
    "user": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.User().list({"organization_id": "example"})
for user in results:
    print(user)
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.User().remove({"id": "id", "organization_id": "organization_id"})
```

### Common Methods

#### `data_get() -> dict`

Get the entity data.

#### `data_set(data)`

Set the entity data.

#### `match_get() -> dict`

Get the entity match criteria.

#### `match_set(match)`

Set the entity match criteria.

#### `make() -> Entity`

Create a new `UserEntity` instance with the same options.

#### `get_name() -> str`

Return the entity name.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```python
client = StatuspageSDK({
    "feature": {
        "test": {"active": True},
    },
})
```

