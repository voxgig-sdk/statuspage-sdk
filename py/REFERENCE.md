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
| `automation_email` | `str` | No |  |
| `component` | `dict` | No |  |
| `created_at` | `str` | No |  |
| `description` | `str` | No |  |
| `group` | `bool` | No |  |
| `group_id` | `str` | No |  |
| `id` | `str` | No |  |
| `major_outage` | `int` | No |  |
| `name` | `str` | No |  |
| `only_show_if_degraded` | `bool` | No |  |
| `page_id` | `str` | No |  |
| `partial_outage` | `int` | No |  |
| `position` | `int` | No |  |
| `range_end` | `str` | No |  |
| `range_start` | `str` | No |  |
| `related_event` | `dict` | No |  |
| `showcase` | `bool` | No |  |
| `start_date` | `str` | No |  |
| `status` | `str` | No |  |
| `updated_at` | `str` | No |  |
| `uptime_percentage` | `float` | No |  |
| `warning` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Component().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Component().list()
for component in results:
    print(component)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Component().load({"id": "component_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Component().remove({"id": "component_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Component().update({
    "id": "component_id",
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
| `id` | `str` | No |  |
| `major_outage` | `int` | No |  |
| `name` | `str` | No |  |
| `partial_outage` | `int` | No |  |
| `range_end` | `str` | No |  |
| `range_start` | `str` | No |  |
| `related_event` | `dict` | No |  |
| `uptime_percentage` | `float` | No |  |
| `warning` | `str` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.ComponentGroupUptime().load({"id": "component_group_uptime_id"})
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
| `component` | `str` | No |  |
| `component_group` | `dict` | Yes |  |
| `created_at` | `str` | No |  |
| `description` | `str` | No |  |
| `id` | `str` | No |  |
| `name` | `str` | No |  |
| `page_id` | `str` | No |  |
| `position` | `str` | No |  |
| `updated_at` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.GroupComponent().create({
    "component_group": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.GroupComponent().list()
for group_component in results:
    print(group_component)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.GroupComponent().load({"id": "group_component_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.GroupComponent().remove({"id": "group_component_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.GroupComponent().update({
    "id": "group_component_id",
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
| `auto_transition_deliver_notifications_at_end` | `bool` | No |  |
| `auto_transition_deliver_notifications_at_start` | `bool` | No |  |
| `auto_transition_to_maintenance_state` | `bool` | No |  |
| `auto_transition_to_operational_state` | `bool` | No |  |
| `component` | `list` | No |  |
| `created_at` | `str` | No |  |
| `id` | `str` | No |  |
| `impact` | `str` | No |  |
| `impact_override` | `str` | No |  |
| `incident` | `dict` | Yes |  |
| `incident_impact` | `list` | No |  |
| `incident_update` | `list` | No |  |
| `metadata` | `dict` | No |  |
| `monitoring_at` | `str` | No |  |
| `name` | `str` | No |  |
| `page_id` | `str` | No |  |
| `postmortem_body` | `str` | No |  |
| `postmortem_body_last_updated_at` | `str` | No |  |
| `postmortem_ignored` | `bool` | No |  |
| `postmortem_notified_subscriber` | `bool` | No |  |
| `postmortem_notified_twitter` | `bool` | No |  |
| `postmortem_published_at` | `bool` | No |  |
| `reminder_interval` | `str` | No |  |
| `resolved_at` | `str` | No |  |
| `scheduled_auto_completed` | `bool` | No |  |
| `scheduled_auto_in_progress` | `bool` | No |  |
| `scheduled_for` | `str` | No |  |
| `scheduled_remind_prior` | `bool` | No |  |
| `scheduled_reminded_at` | `str` | No |  |
| `scheduled_until` | `str` | No |  |
| `shortlink` | `str` | No |  |
| `status` | `str` | No |  |
| `updated_at` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Incident().create({
    "incident": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Incident().list()
for incident in results:
    print(incident)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Incident().load({"id": "incident_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Incident().remove({"id": "incident_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Incident().update({
    "id": "incident_id",
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
result = client.IncidentPostmortem().remove()
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
| `body` | `str` | No |  |
| `component` | `list` | No |  |
| `group_id` | `str` | No |  |
| `id` | `str` | No |  |
| `name` | `str` | No |  |
| `should_send_notification` | `bool` | No |  |
| `should_tweet` | `bool` | No |  |
| `template` | `dict` | Yes |  |
| `title` | `str` | No |  |
| `update_status` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.IncidentTemplate().create({
    "template": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.IncidentTemplate().list()
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
| `affected_component` | `list` | No |  |
| `body` | `str` | No |  |
| `created_at` | `str` | No |  |
| `custom_tweet` | `str` | No |  |
| `deliver_notification` | `bool` | No |  |
| `display_at` | `str` | No |  |
| `id` | `str` | No |  |
| `incident_id` | `str` | No |  |
| `incident_update` | `dict` | No |  |
| `status` | `str` | No |  |
| `tweet_id` | `str` | No |  |
| `twitter_updated_at` | `str` | No |  |
| `updated_at` | `str` | No |  |
| `wants_twitter_update` | `bool` | No |  |

### Operations

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.IncidentUpdate().update({
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
| `data` | `dict` | Yes |  |
| `decimal_place` | `int` | No |  |
| `display` | `bool` | No |  |
| `id` | `str` | No |  |
| `last_fetched_at` | `str` | No |  |
| `metric` | `dict` | No |  |
| `metric_identifier` | `str` | No |  |
| `metrics_provider_id` | `str` | No |  |
| `most_recent_data_at` | `str` | No |  |
| `name` | `str` | No |  |
| `reference_name` | `str` | No |  |
| `suffix` | `str` | No |  |
| `tooltip_description` | `str` | No |  |
| `updated_at` | `str` | No |  |
| `y_axis_hidden` | `bool` | No |  |
| `y_axis_max` | `float` | No |  |
| `y_axis_min` | `float` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Metric().create({
    "data": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Metric().list()
for metric in results:
    print(metric)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Metric().load({"id": "metric_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Metric().remove({"id": "metric_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Metric().update({
    "id": "metric_id",
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
| `id` | `str` | No |  |
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
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.MetricsProvider().list()
for metrics_provider in results:
    print(metrics_provider)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.MetricsProvider().load({"id": "metrics_provider_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.MetricsProvider().remove({"id": "metrics_provider_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.MetricsProvider().update({
    "id": "metrics_provider_id",
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
| `allow_email_subscriber` | `bool` | No |  |
| `allow_incident_subscriber` | `bool` | No |  |
| `allow_page_subscriber` | `bool` | No |  |
| `allow_rss_atom_feed` | `bool` | No |  |
| `allow_sms_subscriber` | `bool` | No |  |
| `allow_webhook_subscriber` | `bool` | No |  |
| `branding` | `str` | No |  |
| `city` | `str` | No |  |
| `country` | `str` | No |  |
| `created_at` | `str` | No |  |
| `css_blue` | `str` | No |  |
| `css_body_background_color` | `str` | No |  |
| `css_border_color` | `str` | No |  |
| `css_font_color` | `str` | No |  |
| `css_graph_color` | `str` | No |  |
| `css_green` | `str` | No |  |
| `css_light_font_color` | `str` | No |  |
| `css_link_color` | `str` | No |  |
| `css_no_data` | `str` | No |  |
| `css_orange` | `str` | No |  |
| `css_red` | `str` | No |  |
| `css_yellow` | `str` | No |  |
| `domain` | `str` | No |  |
| `email_logo` | `str` | No |  |
| `favicon_logo` | `str` | No |  |
| `headline` | `str` | No |  |
| `hero_cover` | `str` | No |  |
| `hidden_from_search` | `bool` | No |  |
| `id` | `str` | No |  |
| `ip_restriction` | `str` | No |  |
| `name` | `str` | No |  |
| `notifications_email_footer` | `str` | No |  |
| `notifications_from_email` | `str` | No |  |
| `page` | `dict` | No |  |
| `page_description` | `str` | No |  |
| `state` | `str` | No |  |
| `subdomain` | `str` | No |  |
| `support_url` | `str` | No |  |
| `time_zone` | `str` | No |  |
| `transactional_logo` | `str` | No |  |
| `twitter_logo` | `str` | No |  |
| `twitter_username` | `str` | No |  |
| `updated_at` | `str` | No |  |
| `url` | `str` | No |  |
| `viewers_must_be_team_member` | `bool` | No |  |

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
| `component_id` | `list` | No |  |
| `created_at` | `str` | No |  |
| `external_identifier` | `str` | No |  |
| `id` | `str` | No |  |
| `metric_id` | `list` | No |  |
| `name` | `str` | No |  |
| `page_access_group` | `dict` | No |  |
| `page_access_user_id` | `list` | No |  |
| `page_id` | `str` | No |  |
| `updated_at` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PageAccessGroup().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PageAccessGroup().list()
for page_access_group in results:
    print(page_access_group)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PageAccessGroup().load({"id": "page_access_group_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PageAccessGroup().remove({"id": "page_access_group_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PageAccessGroup().update({
    "id": "page_access_group_id",
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
| `component_id` | `list` | Yes |  |
| `created_at` | `str` | No |  |
| `email` | `str` | No |  |
| `external_login` | `str` | No |  |
| `id` | `str` | No |  |
| `metric_id` | `list` | Yes |  |
| `page_access_group_id` | `str` | No |  |
| `page_access_user` | `dict` | No |  |
| `page_id` | `str` | No |  |
| `updated_at` | `str` | No |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.PageAccessUser().create({
    "component_id": [],  # list
    "metric_id": [],  # list
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.PageAccessUser().list()
for page_access_user in results:
    print(page_access_user)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.PageAccessUser().load({"id": "page_access_user_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.PageAccessUser().remove({"id": "page_access_user_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.PageAccessUser().update({
    "id": "page_access_user_id",
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
| `data` | `dict` | No |  |
| `page` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Permission().load({"id": "permission_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Permission().update({
    "id": "permission_id",
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
| `body` | `str` | No |  |
| `body_draft` | `str` | No |  |
| `body_draft_updated_at` | `str` | No |  |
| `body_updated_at` | `str` | No |  |
| `created_at` | `str` | No |  |
| `custom_tweet` | `str` | No |  |
| `notify_subscriber` | `bool` | No |  |
| `notify_twitter` | `bool` | No |  |
| `postmortem` | `dict` | Yes |  |
| `preview_key` | `str` | No |  |
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
| `notify_subscriber` | - | - |
| `notify_twitter` | - | - |
| `postmortem` | - | Yes |
| `preview_key` | - | - |
| `published_at` | - | - |
| `updated_at` | - | - |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Postmortem().load()
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Postmortem().update({
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
| `incident_background_color` | `str` | No |  |
| `incident_text_color` | `str` | No |  |
| `maintenance_background_color` | `str` | No |  |
| `maintenance_text_color` | `str` | No |  |
| `page_id` | `str` | No |  |
| `position` | `str` | No |  |
| `status_embed_config` | `dict` | No |  |

### Operations

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.StatusEmbedConfig().load()
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.StatusEmbedConfig().update({
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
| `component` | `str` | No |  |
| `component_id` | `list` | No |  |
| `created_at` | `str` | No |  |
| `display_phone_number` | `str` | No |  |
| `email` | `str` | No |  |
| `endpoint` | `str` | No |  |
| `id` | `str` | No |  |
| `integration_partner` | `int` | No |  |
| `mode` | `str` | No |  |
| `obfuscated_channel_name` | `str` | No |  |
| `page_access_user_id` | `str` | No |  |
| `phone_country` | `str` | No |  |
| `phone_number` | `str` | No |  |
| `purge_at` | `str` | No |  |
| `quarantined_at` | `str` | No |  |
| `skip_confirmation_notification` | `bool` | No |  |
| `skip_unsubscription_notification` | `bool` | No |  |
| `slack` | `int` | No |  |
| `sms` | `int` | No |  |
| `state` | `str` | No |  |
| `subscriber` | `dict` | No |  |
| `team` | `int` | No |  |
| `type` | `str` | No |  |
| `webhook` | `int` | No |  |
| `workspace_name` | `str` | No |  |

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

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.Subscriber().create({
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.Subscriber().list()
for subscriber in results:
    print(subscriber)
```

#### `load(reqmatch, ctrl=None) -> dict`

Load a single entity matching the given criteria. Returns the entity data and raises on error.

```python
result = client.Subscriber().load({"id": "subscriber_id"})
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.Subscriber().remove({"id": "subscriber_id"})
```

#### `update(reqdata, ctrl=None) -> dict`

Update an existing entity. The data must include the entity `id`. Returns the updated entity data and raises on error.

```python
result = client.Subscriber().update({
    "id": "subscriber_id",
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
| `email` | `str` | No |  |
| `first_name` | `str` | No |  |
| `id` | `str` | No |  |
| `last_name` | `str` | No |  |
| `organization_id` | `str` | No |  |
| `updated_at` | `str` | No |  |
| `user` | `dict` | Yes |  |

### Operations

#### `create(reqdata, ctrl=None) -> dict`

Create a new entity with the given data. Returns the created entity data and raises on error.

```python
result = client.User().create({
    "user": {},  # dict
})
```

#### `list(reqmatch=None, ctrl=None) -> list`

List entities matching the given criteria. The match is optional — call `list()` with no argument to list all records. Returns a list and raises on error.

```python
results = client.User().list()
for user in results:
    print(user)
```

#### `remove(reqmatch, ctrl=None) -> dict`

Remove the entity matching the given criteria. Raises on error.

```python
result = client.User().remove()
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

