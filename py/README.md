# Statuspage Python SDK



The Python SDK for the Statuspage API — an entity-oriented client following Pythonic conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Component()` — each
carrying a small, uniform set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`) instead of raw URL
paths and query strings. You work with named resources and verbs, which
keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to PyPI. Install it from the GitHub
release tag (`py/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/statuspage-sdk/releases)) or
from a source checkout:

```bash
pip install -e .
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```python
import os
from statuspage_sdk import StatuspageSDK

client = StatuspageSDK({
    "apikey": os.environ.get("STATUSPAGE_APIKEY"),
})
```

### 2. List component records

`list()` returns a `list` of records (each a `dict`) and raises on
error — iterate it directly.

```python
try:
    components = client.Component().list()
    for component in components:
        print(component)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a component

Component is nested under page, so provide the `page_id`.
`load()` returns the bare record (a `dict`) and raises on error.

```python
try:
    component = client.Component().load({"page_id": "example_page_id", "id": "example_id"})
    print(component)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the bare created record (a dict)
created = client.Component().create({"page_id": "example_page_id"})

# Update — the created record's id is a plain dict key
client.Component().update({"id": created["id"], "page_id": "example_page_id"})

# Remove
client.Component().remove({"id": created["id"], "page_id": "example_page_id"})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    components = client.Component().list()
    print(components)
except Exception as err:
    print(f"list failed: {err}")
```

`direct()` does **not** raise — it returns the result envelope. Branch
on `ok`; on failure `status` holds the HTTP status (for error responses)
and `err` holds a transport error, so read both defensively:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example_id"},
})

if not result["ok"]:
    print("request failed:", result.get("status"), result.get("err"))
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```python
result = client.direct({
    "path": "/api/resource/{id}",
    "method": "GET",
    "params": {"id": "example"},
})

if result["ok"]:
    print(result["status"])  # 200
    print(result["data"])    # response body
else:
    # A non-2xx response carries status + data (the error body); a
    # transport-level failure carries err instead. Only one is present, so
    # read both with .get() rather than indexing a key that may be absent.
    print(result.get("status"), result.get("err"))
```

### Prepare a request without sending it

```python
# prepare() returns the fetch definition and raises on error.
fetchdef = client.prepare({
    "path": "/api/resource/{id}",
    "method": "DELETE",
    "params": {"id": "example"},
})

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```python
client = StatuspageSDK.test()

# Entity ops return the bare record and raise on error.
component = client.Component().list()
# component contains the mock response record
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```python
def mock_fetch(url, init):
    return {
        "status": 200,
        "statusText": "OK",
        "headers": {},
        "json": lambda: {"id": "mock01"},
    }, None

client = StatuspageSDK({
    "base": "http://localhost:8080",
    "system": {
        "fetch": mock_fetch,
    },
})
```

### Run live tests

Create a `.env.local` file at the project root:

```
STATUSPAGE_TEST_LIVE=TRUE
STATUSPAGE_APIKEY=<your-key>
```

Then run:

```bash
cd py && pytest test/
```


## Reference

### StatuspageSDK

```python
from statuspage_sdk import StatuspageSDK

client = StatuspageSDK(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `str` | API key for authentication. |
| `base` | `str` | Base URL of the API server. |
| `prefix` | `str` | URL path prefix prepended to all requests. |
| `suffix` | `str` | URL path suffix appended to all requests. |
| `feature` | `dict` | Feature activation flags. |
| `extend` | `list` | Additional Feature instances to load. |
| `system` | `dict` | System overrides (e.g. custom `fetch` function). |

### test

```python
client = StatuspageSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `None`.

### StatuspageSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> dict` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> dict` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> dict` | Build and send an HTTP request. Returns a result dict (branch on `ok`). |
| `Component` | `(data) -> ComponentEntity` | Create a Component entity instance. |
| `ComponentGroupUptime` | `(data) -> ComponentGroupUptimeEntity` | Create a ComponentGroupUptime entity instance. |
| `GroupComponent` | `(data) -> GroupComponentEntity` | Create a GroupComponent entity instance. |
| `Incident` | `(data) -> IncidentEntity` | Create an Incident entity instance. |
| `IncidentPostmortem` | `(data) -> IncidentPostmortemEntity` | Create an IncidentPostmortem entity instance. |
| `IncidentSubscriber` | `(data) -> IncidentSubscriberEntity` | Create an IncidentSubscriber entity instance. |
| `IncidentTemplate` | `(data) -> IncidentTemplateEntity` | Create an IncidentTemplate entity instance. |
| `IncidentUpdate` | `(data) -> IncidentUpdateEntity` | Create an IncidentUpdate entity instance. |
| `Metric` | `(data) -> MetricEntity` | Create a Metric entity instance. |
| `MetricsProvider` | `(data) -> MetricsProviderEntity` | Create a MetricsProvider entity instance. |
| `Page` | `(data) -> PageEntity` | Create a Page entity instance. |
| `PageAccessGroup` | `(data) -> PageAccessGroupEntity` | Create a PageAccessGroup entity instance. |
| `PageAccessUser` | `(data) -> PageAccessUserEntity` | Create a PageAccessUser entity instance. |
| `Permission` | `(data) -> PermissionEntity` | Create a Permission entity instance. |
| `Postmortem` | `(data) -> PostmortemEntity` | Create a Postmortem entity instance. |
| `StatusEmbedConfig` | `(data) -> StatusEmbedConfigEntity` | Create a StatusEmbedConfig entity instance. |
| `Subscriber` | `(data) -> SubscriberEntity` | Create a Subscriber entity instance. |
| `User` | `(data) -> UserEntity` | Create an User entity instance. |

### Entity interface

All entities share the same interface.

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `(reqmatch, ctrl) -> any` | Load a single entity by match criteria. Raises on error. |
| `list` | `(reqmatch, ctrl) -> list` | List entities matching the criteria. Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> dict` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> dict` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> str` | Return the entity name. |

### Result shape

Entity operations return the bare result data (a `dict` for single-entity
ops, a `list` for `list`) and raise on error. Wrap calls in
`try`/`except` to handle failures.

The `direct()` escape hatch never raises — it returns a result `dict`
you branch on via `result["ok"]`:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `bool` | `True` if the HTTP status is 2xx. |
| `status` | `int` | HTTP status code. |
| `headers` | `dict` | Response headers. |
| `data` | `any` | Parsed JSON response body. |

On error, `ok` is `False` and `err` contains the error value.

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

Create an instance: `component = client.Component()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `automation_email` | `str` |  |
| `component` | `dict` |  |
| `created_at` | `str` |  |
| `description` | `str` |  |
| `group` | `bool` |  |
| `group_id` | `str` |  |
| `id` | `str` |  |
| `major_outage` | `int` |  |
| `name` | `str` |  |
| `only_show_if_degraded` | `bool` |  |
| `page_id` | `str` |  |
| `partial_outage` | `int` |  |
| `position` | `int` |  |
| `range_end` | `str` |  |
| `range_start` | `str` |  |
| `related_event` | `dict` |  |
| `showcase` | `bool` |  |
| `start_date` | `str` |  |
| `status` | `str` |  |
| `updated_at` | `str` |  |
| `uptime_percentage` | `float` |  |
| `warning` | `str` |  |

#### Example: Load

```python
component = client.Component().load({"id": "component_id", "page_id": "page_id"})
```

#### Example: List

```python
components = client.Component().list()
```

#### Example: Create

```python
component = client.Component().create({
    "page_id": "example_page_id",  # str
})
```


### ComponentGroupUptime

Create an instance: `component_group_uptime = client.ComponentGroupUptime()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `str` |  |
| `major_outage` | `int` |  |
| `name` | `str` |  |
| `partial_outage` | `int` |  |
| `range_end` | `str` |  |
| `range_start` | `str` |  |
| `related_event` | `dict` |  |
| `uptime_percentage` | `float` |  |
| `warning` | `str` |  |

#### Example: Load

```python
component_group_uptime = client.ComponentGroupUptime().load({"id": "component_group_uptime_id", "page_id": "page_id"})
```


### GroupComponent

Create an instance: `group_component = client.GroupComponent()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `str` |  |
| `component_group` | `dict` |  |
| `created_at` | `str` |  |
| `description` | `str` |  |
| `id` | `str` |  |
| `name` | `str` |  |
| `page_id` | `str` |  |
| `position` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
group_component = client.GroupComponent().load({"id": "group_component_id", "page_id": "page_id"})
```

#### Example: List

```python
group_components = client.GroupComponent().list()
```

#### Example: Create

```python
group_component = client.GroupComponent().create({
    "page_id": "example_page_id",  # str
})
```


### Incident

Create an instance: `incident = client.Incident()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
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
| `component` | `list` |  |
| `created_at` | `str` |  |
| `id` | `str` |  |
| `impact` | `str` |  |
| `impact_override` | `str` |  |
| `incident` | `dict` |  |
| `incident_impact` | `list` |  |
| `incident_update` | `list` |  |
| `metadata` | `dict` |  |
| `monitoring_at` | `str` |  |
| `name` | `str` |  |
| `page_id` | `str` |  |
| `postmortem_body` | `str` |  |
| `postmortem_body_last_updated_at` | `str` |  |
| `postmortem_ignored` | `bool` |  |
| `postmortem_notified_subscriber` | `bool` |  |
| `postmortem_notified_twitter` | `bool` |  |
| `postmortem_published_at` | `bool` |  |
| `reminder_interval` | `str` |  |
| `resolved_at` | `str` |  |
| `scheduled_auto_completed` | `bool` |  |
| `scheduled_auto_in_progress` | `bool` |  |
| `scheduled_for` | `str` |  |
| `scheduled_remind_prior` | `bool` |  |
| `scheduled_reminded_at` | `str` |  |
| `scheduled_until` | `str` |  |
| `shortlink` | `str` |  |
| `status` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
incident = client.Incident().load({"id": "incident_id", "page_id": "page_id"})
```

#### Example: List

```python
incidents = client.Incident().list()
```

#### Example: Create

```python
incident = client.Incident().create({
    "page_id": "example_page_id",  # str
})
```


### IncidentPostmortem

Create an instance: `incident_postmortem = client.IncidentPostmortem()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `incident_subscriber = client.IncidentSubscriber()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```python
incident_subscriber = client.IncidentSubscriber().create({
    "incident_id": "example_incident_id",  # str
    "page_id": "example_page_id",  # str
    "subscriber_id": "example_subscriber_id",  # str
})
```


### IncidentTemplate

Create an instance: `incident_template = client.IncidentTemplate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `str` |  |
| `component` | `list` |  |
| `group_id` | `str` |  |
| `id` | `str` |  |
| `name` | `str` |  |
| `should_send_notification` | `bool` |  |
| `should_tweet` | `bool` |  |
| `template` | `dict` |  |
| `title` | `str` |  |
| `update_status` | `str` |  |

#### Example: List

```python
incident_templates = client.IncidentTemplate().list()
```

#### Example: Create

```python
incident_template = client.IncidentTemplate().create({
    "page_id": "example_page_id",  # str
})
```


### IncidentUpdate

Create an instance: `incident_update = client.IncidentUpdate()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_component` | `list` |  |
| `body` | `str` |  |
| `created_at` | `str` |  |
| `custom_tweet` | `str` |  |
| `deliver_notification` | `bool` |  |
| `display_at` | `str` |  |
| `id` | `str` |  |
| `incident_id` | `str` |  |
| `incident_update` | `dict` |  |
| `status` | `str` |  |
| `tweet_id` | `str` |  |
| `twitter_updated_at` | `str` |  |
| `updated_at` | `str` |  |
| `wants_twitter_update` | `bool` |  |


### Metric

Create an instance: `metric = client.Metric()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `backfill_percentage` | `int` |  |
| `backfilled` | `bool` |  |
| `created_at` | `str` |  |
| `data` | `dict` |  |
| `decimal_place` | `int` |  |
| `display` | `bool` |  |
| `id` | `str` |  |
| `last_fetched_at` | `str` |  |
| `metric` | `dict` |  |
| `metric_identifier` | `str` |  |
| `metrics_provider_id` | `str` |  |
| `most_recent_data_at` | `str` |  |
| `name` | `str` |  |
| `reference_name` | `str` |  |
| `suffix` | `str` |  |
| `tooltip_description` | `str` |  |
| `updated_at` | `str` |  |
| `y_axis_hidden` | `bool` |  |
| `y_axis_max` | `float` |  |
| `y_axis_min` | `float` |  |

#### Example: Load

```python
metric = client.Metric().load({"id": "metric_id", "page_id": "page_id"})
```

#### Example: List

```python
metrics = client.Metric().list()
```

#### Example: Create

```python
metric = client.Metric().create({
    "metrics_provider_id": "example_metrics_provider_id",  # str
    "page_id": "example_page_id",  # str
})
```


### MetricsProvider

Create an instance: `metrics_provider = client.MetricsProvider()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` |  |
| `disabled` | `bool` |  |
| `id` | `str` |  |
| `last_revalidated_at` | `str` |  |
| `metric_base_uri` | `str` |  |
| `metrics_provider` | `dict` |  |
| `page_id` | `int` |  |
| `type` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
metrics_provider = client.MetricsProvider().load({"id": "metrics_provider_id", "page_id": "page_id"})
```

#### Example: List

```python
metrics_providers = client.MetricsProvider().list()
```

#### Example: Create

```python
metrics_provider = client.MetricsProvider().create({
    "page_id": "example_page_id",  # str
})
```


### Page

Create an instance: `page = client.Page()`

#### Operations

| Method | Description |
| --- | --- |
| `list()` | List entities, optionally matching the given criteria. |
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
| `branding` | `str` |  |
| `city` | `str` |  |
| `country` | `str` |  |
| `created_at` | `str` |  |
| `css_blue` | `str` |  |
| `css_body_background_color` | `str` |  |
| `css_border_color` | `str` |  |
| `css_font_color` | `str` |  |
| `css_graph_color` | `str` |  |
| `css_green` | `str` |  |
| `css_light_font_color` | `str` |  |
| `css_link_color` | `str` |  |
| `css_no_data` | `str` |  |
| `css_orange` | `str` |  |
| `css_red` | `str` |  |
| `css_yellow` | `str` |  |
| `domain` | `str` |  |
| `email_logo` | `str` |  |
| `favicon_logo` | `str` |  |
| `headline` | `str` |  |
| `hero_cover` | `str` |  |
| `hidden_from_search` | `bool` |  |
| `id` | `str` |  |
| `ip_restriction` | `str` |  |
| `name` | `str` |  |
| `notifications_email_footer` | `str` |  |
| `notifications_from_email` | `str` |  |
| `page` | `dict` |  |
| `page_description` | `str` |  |
| `state` | `str` |  |
| `subdomain` | `str` |  |
| `support_url` | `str` |  |
| `time_zone` | `str` |  |
| `transactional_logo` | `str` |  |
| `twitter_logo` | `str` |  |
| `twitter_username` | `str` |  |
| `updated_at` | `str` |  |
| `url` | `str` |  |
| `viewers_must_be_team_member` | `bool` |  |

#### Example: Load

```python
page = client.Page().load({"id": "page_id"})
```

#### Example: List

```python
pages = client.Page().list()
```


### PageAccessGroup

Create an instance: `page_access_group = client.PageAccessGroup()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `list` |  |
| `created_at` | `str` |  |
| `external_identifier` | `str` |  |
| `id` | `str` |  |
| `metric_id` | `list` |  |
| `name` | `str` |  |
| `page_access_group` | `dict` |  |
| `page_access_user_id` | `list` |  |
| `page_id` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
page_access_group = client.PageAccessGroup().load({"id": "page_access_group_id", "page_id": "page_id"})
```

#### Example: List

```python
page_access_groups = client.PageAccessGroup().list()
```

#### Example: Create

```python
page_access_group = client.PageAccessGroup().create({
    "id": "example_id",  # str
})
```


### PageAccessUser

Create an instance: `page_access_user = client.PageAccessUser()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `list` |  |
| `created_at` | `str` |  |
| `email` | `str` |  |
| `external_login` | `str` |  |
| `id` | `str` |  |
| `metric_id` | `list` |  |
| `page_access_group_id` | `str` |  |
| `page_access_user` | `dict` |  |
| `page_id` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
page_access_user = client.PageAccessUser().load({"id": "page_access_user_id", "page_id": "page_id"})
```

#### Example: List

```python
page_access_users = client.PageAccessUser().list()
```

#### Example: Create

```python
page_access_user = client.PageAccessUser().create({
    "id": "example_id",  # str
})
```


### Permission

Create an instance: `permission = client.Permission()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `dict` |  |
| `page` | `dict` |  |

#### Example: Load

```python
permission = client.Permission().load({"id": "permission_id", "organization_id": "organization_id"})
```


### Postmortem

Create an instance: `postmortem = client.Postmortem()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `str` |  |
| `body_draft` | `str` |  |
| `body_draft_updated_at` | `str` |  |
| `body_updated_at` | `str` |  |
| `created_at` | `str` |  |
| `custom_tweet` | `str` |  |
| `notify_subscriber` | `bool` |  |
| `notify_twitter` | `bool` |  |
| `postmortem` | `dict` |  |
| `preview_key` | `str` |  |
| `published_at` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
postmortem = client.Postmortem().load({"incident_id": "incident_id", "page_id": "page_id"})
```


### StatusEmbedConfig

Create an instance: `status_embed_config = client.StatusEmbedConfig()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `incident_background_color` | `str` |  |
| `incident_text_color` | `str` |  |
| `maintenance_background_color` | `str` |  |
| `maintenance_text_color` | `str` |  |
| `page_id` | `str` |  |
| `position` | `str` |  |
| `status_embed_config` | `dict` |  |

#### Example: Load

```python
status_embed_config = client.StatusEmbedConfig().load({"page_id": "page_id"})
```


### Subscriber

Create an instance: `subscriber = client.Subscriber()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `str` |  |
| `component_id` | `list` |  |
| `created_at` | `str` |  |
| `display_phone_number` | `str` |  |
| `email` | `str` |  |
| `endpoint` | `str` |  |
| `id` | `str` |  |
| `integration_partner` | `int` |  |
| `mode` | `str` |  |
| `obfuscated_channel_name` | `str` |  |
| `page_access_user_id` | `str` |  |
| `phone_country` | `str` |  |
| `phone_number` | `str` |  |
| `purge_at` | `str` |  |
| `quarantined_at` | `str` |  |
| `skip_confirmation_notification` | `bool` |  |
| `skip_unsubscription_notification` | `bool` |  |
| `slack` | `int` |  |
| `sms` | `int` |  |
| `state` | `str` |  |
| `subscriber` | `dict` |  |
| `team` | `int` |  |
| `type` | `str` |  |
| `webhook` | `int` |  |
| `workspace_name` | `str` |  |

#### Example: Load

```python
subscriber = client.Subscriber().load({"id": "subscriber_id", "page_id": "page_id"})
```

#### Example: List

```python
subscribers = client.Subscriber().list()
```

#### Example: Create

```python
subscriber = client.Subscriber().create({
    "page_id": "example_page_id",  # str
})
```


### User

Create an instance: `user = client.User()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list()` | List entities, optionally matching the given criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `str` |  |
| `email` | `str` |  |
| `first_name` | `str` |  |
| `id` | `str` |  |
| `last_name` | `str` |  |
| `organization_id` | `str` |  |
| `updated_at` | `str` |  |
| `user` | `dict` |  |

#### Example: List

```python
users = client.User().list()
```

#### Example: Create

```python
user = client.User().create({
    "organization_id": "example_organization_id",  # str
})
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

Features are the extension mechanism. A feature is a Python class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as dicts

The Python SDK uses plain dicts throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a dict.

### Module structure

```
py/
├── statuspage_sdk.py         -- Main SDK module
├── config.py                    -- Configuration
├── features.py                  -- Feature factory
├── core/                        -- Core types and context
├── entity/                      -- Entity implementations
├── feature/                     -- Built-in features (Base, Test, Log)
├── utility/                     -- Utility functions and struct library
└── test/                        -- Test suites
```

The main module (`statuspage_sdk`) exports the SDK class.
Import entity or utility modules directly only when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```python
component = client.Component()
component.list()

# component.data_get() now returns the component data from the last list
# component.match_get() returns the last match criteria
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
