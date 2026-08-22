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
    components = client.Component().list({"page_id": "example"})
    for component in components:
        print(component)
except Exception as err:
    print(f"list failed: {err}")
```

### 3. Load a component

Component is nested under page, so provide the `page_id`.
`load()` returns the ENTITY — call data_get() for the record — and raises on error.

```python
try:
    component = client.Component().load({"page_id": "example_page_id", "id": "example_id"})
    print(component)
except Exception as err:
    print(f"load failed: {err}")
```

### 4. Create, update, and remove

```python
# Create — returns the ENTITY (call data_get() for the record)
created = client.Component().create({"page_id": "example_page_id"})

# Update — the created record's id is a plain dict key
client.Component().update({"id": created.data_get()["id"], "page_id": "example_page_id", "automation_email": "example_automation_email"})

# Remove
client.Component().remove({"id": created.data_get()["id"], "page_id": "example_page_id"})
```


## Error handling

Entity operations raise on failure, so wrap them in `try` / `except`:

```python
try:
    postmortem = client.Postmortem().load({"incident_id": "example", "page_id": "example"})
    print(postmortem)
except Exception as err:
    print(f"load failed: {err}")
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

# Entity ops return the ENTITY and raises on error;
# call data_get() for the record.
postmortem = client.Postmortem().load({"incident_id": "example", "page_id": "example"})
# postmortem contains the mock response record
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

Entity operations return the ENTITY (call data_get() for the record) (a `dict` for single-entity
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
| `automation_email` | Requires a special feature flag to be enabled |
| `component` |  |
| `created_at` |  |
| `description` | More detailed description for component |
| `group` | Is this component a group |
| `group_id` | Component Group identifier |
| `id` | Incident identifier |
| `name` | Display name for component |
| `only_show_if_degraded` | Requires a special feature flag to be enabled |
| `page_id` | Page identifier |
| `position` | Order the component will appear on the page |
| `showcase` | Should this component be showcased |
| `start_date` | The date this component started being used |
| `status` | Status of component |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/components/{component_id}/page_access_groups`

#### ComponentGroupUptime

| Field | Description |
| --- | --- |
| `component_id` | Component identifier |
| `incidents` | Related incidents |

Operations: Load.

API path: `/pages/{page_id}/component-groups/{id}/uptime`

#### GroupComponent

| Field | Description |
| --- | --- |
| `component_group` |  |
| `components` |  |
| `created_at` |  |
| `description` | Description of the component group. |
| `id` | Component Group Identifier |
| `name` |  |
| `page_id` |  |
| `position` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/component-groups`

#### Incident

| Field | Description |
| --- | --- |
| `auto_transition_deliver_notifications_at_end` | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | Incident components |
| `created_at` | The timestamp when the incident was created at. |
| `id` | Incident Identifier |
| `impact` | The impact of the incident. |
| `impact_override` | value to override calculated impact value |
| `incident` |  |
| `incident_updates` | The incident updates for incident. |
| `metadata` | Metadata attached to the incident. |
| `monitoring_at` | The timestamp when incident entered monitoring state. |
| `name` | Incident Name. |
| `page_id` | Incident Page Identifier |
| `postmortem_body` | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | The timestamp when the postmortem was published. |
| `reminder_intervals` | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | The timestamp the incident is scheduled until. |
| `shortlink` | Incident Shortlink. |
| `status` | The incident status. |
| `updated_at` | The timestamp when the incident was updated at. |

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
| `body` | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | Affected components |
| `group_id` | Identifier of Template Group this template belongs to |
| `id` | Incident Template Identifier |
| `name` | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` |  |
| `title` | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | The status the incident or maintenance should transition to when selecting this template |

Operations: Create, List.

API path: `/pages/{page_id}/incident_templates`

#### IncidentUpdate

| Field | Description |
| --- | --- |
| `affected_components` | Affected components associated with the incident update. |
| `body` | Incident update body. |
| `created_at` | The timestamp when the incident update was created at. |
| `custom_tweet` | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | Controls whether to delivery notifications. |
| `display_at` | Timestamp when incident update is happened. |
| `id` | Incident Update Identifier. |
| `incident_id` | Incident Identifier. |
| `incident_update` |  |
| `status` | The incident status. |
| `tweet_id` | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | The timestamp when twitter updated at. |
| `updated_at` | The timestamp when the incident update is updated. |
| `wants_twitter_update` | Controls whether to create twitter update. |

Operations: Patch, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}`

#### Metric

| Field | Description |
| --- | --- |
| `backfill_percentage` |  |
| `backfilled` |  |
| `created_at` |  |
| `data` | Add data points to metrics |
| `decimal_places` |  |
| `display` | Should the metric be displayed |
| `id` | Metric identifier |
| `last_fetched_at` |  |
| `metric` |  |
| `metric_identifier` | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | Metric Provider identifier |
| `most_recent_data_at` |  |
| `name` | Name of metric |
| `reference_name` |  |
| `suffix` | Suffix to describe the units on the graph |
| `tooltip_description` |  |
| `updated_at` |  |
| `y_axis_hidden` | Should the values on the y axis be hidden on render |
| `y_axis_max` |  |
| `y_axis_min` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/metrics/{metric_id}/data`

#### MetricsProvider

| Field | Description |
| --- | --- |
| `created_at` |  |
| `disabled` |  |
| `id` | Identifier for Metrics Provider |
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
| `allow_email_subscribers` | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | Can your users choose to receive notifications via Webhooks |
| `branding` | The main template your statuspage will use |
| `city` |  |
| `country` |  |
| `created_at` | Timestamp the record was created |
| `css_blues` | CSS Color |
| `css_body_background_color` | CSS Color |
| `css_border_color` | CSS Color |
| `css_font_color` | CSS Color |
| `css_graph_color` | CSS Color |
| `css_greens` | CSS Color |
| `css_light_font_color` | CSS Color |
| `css_link_color` | CSS Color |
| `css_no_data` | CSS Color |
| `css_oranges` | CSS Color |
| `css_reds` | CSS Color |
| `css_yellows` | CSS Color |
| `domain` | CNAME alias for your status page |
| `email_logo` |  |
| `favicon_logo` |  |
| `headline` |  |
| `hero_cover` |  |
| `hidden_from_search` | Should your page hide itself from search engines |
| `id` | Page identifier |
| `ip_restrictions` |  |
| `name` | Name of your page to be displayed |
| `notifications_email_footer` | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | Allows you to customize the email address your page notifications come from |
| `page` |  |
| `page_description` |  |
| `state` |  |
| `subdomain` | Subdomain at which to access your status page |
| `support_url` |  |
| `time_zone` | Timezone configured for your page |
| `transactional_logo` |  |
| `twitter_logo` |  |
| `twitter_username` |  |
| `updated_at` | Timestamp the record was last updated |
| `url` | Website of your page. |
| `viewers_must_be_team_members` |  |

Operations: List, Load, Patch, Update.

API path: `/pages`

#### PageAccessGroup

| Field | Description |
| --- | --- |
| `component_ids` | List of components codes to set on the page access group |
| `created_at` |  |
| `external_identifier` | Associates group with external group. |
| `id` | Page Access Group Identifier |
| `metric_ids` |  |
| `name` | Name for this Group. |
| `page_access_group` |  |
| `page_access_user_ids` |  |
| `page_id` | Page Identifier. |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_groups/{page_access_group_id}/components`

#### PageAccessUser

| Field | Description |
| --- | --- |
| `component_ids` | List of component codes to allow access to |
| `created_at` |  |
| `email` |  |
| `external_login` | IDP login user id. |
| `id` | Page Access User Identifier |
| `metric_ids` | List of metrics to add |
| `page_access_group_id` |  |
| `page_access_group_ids` |  |
| `page_access_user` |  |
| `page_id` |  |
| `updated_at` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_users/{page_access_user_id}/components`

#### Permission

| Field | Description |
| --- | --- |
| `pages` | Pages accessible by the user. |
| `user_id` | User identifier |

Operations: Load, Update.

API path: `/organizations/{organization_id}/permissions/{user_id}`

#### Postmortem

| Field | Description |
| --- | --- |
| `body` | Postmortem body |
| `body_draft` | Body draft |
| `body_draft_updated_at` |  |
| `body_updated_at` |  |
| `created_at` |  |
| `custom_tweet` | Custom tweet for Incident Postmortem |
| `notify_subscribers` | Should email subscribers be notified. |
| `notify_twitter` | Should Twitter followers be notified. |
| `postmortem` |  |
| `preview_key` | Preview Key |
| `published_at` |  |
| `updated_at` |  |

Operations: Load, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/postmortem`

#### StatusEmbedConfig

| Field | Description |
| --- | --- |
| `incident_background_color` | Color of status embed iframe background when displaying incident |
| `incident_text_color` | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | Color of status embed iframe text when displaying maintenance |
| `page_id` | Page identifier |
| `position` | Corner where status embed iframe will appear on page |
| `status_embed_config` |  |

Operations: Load, Patch, Update.

API path: `/pages/{page_id}/status_embed_config`

#### Subscriber

| Field | Description |
| --- | --- |
| `component_ids` | A list of component ids for which the subscriber should recieve updates for. |
| `components` | The components for which the subscriber has elected to receive updates. |
| `created_at` |  |
| `display_phone_number` | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | The email address to use to contact the subscriber. |
| `endpoint` | The URL where a webhook subscriber elects to receive updates. |
| `id` | Subscriber Identifier |
| `integration_partner` | The number of integration partners found by the query. |
| `mode` | The communication mode of the subscriber. |
| `obfuscated_channel_name` | Obfuscated slack channel name |
| `page_access_user_id` | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | The phone number used to contact an SMS subscriber |
| `purge_at` | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | The number of Slack subscribers found by the query. |
| `sms` | The number of Webhook subscribers found by the query. |
| `state` | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` |  |
| `subscribers` | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | The number of MS teams subscribers found by the query. |
| `type` | If this is present, only reactivate subscribers of this type. |
| `webhook` | The number of SMS subscribers found by the query. |
| `workspace_name` | The workspace name of the slack subscriber. |

Operations: Create, List, Load, Remove, Update.

API path: `/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation`

#### User

| Field | Description |
| --- | --- |
| `created_at` |  |
| `email` | Email address for the team member |
| `first_name` |  |
| `id` | User identifier |
| `last_name` |  |
| `organization_id` | Organization identifier |
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
| `automation_email` | `str` | Requires a special feature flag to be enabled |
| `component` | `dict` |  |
| `created_at` | `str` |  |
| `description` | `str` | More detailed description for component |
| `group` | `bool` | Is this component a group |
| `group_id` | `str` | Component Group identifier |
| `id` | `str` | Incident identifier |
| `name` | `str` | Display name for component |
| `only_show_if_degraded` | `bool` | Requires a special feature flag to be enabled |
| `page_id` | `str` | Page identifier |
| `position` | `int` | Order the component will appear on the page |
| `showcase` | `bool` | Should this component be showcased |
| `start_date` | `str` | The date this component started being used |
| `status` | `str` | Status of component |
| `updated_at` | `str` |  |

#### Example: Load

```python
component = client.Component().load({"id": "component_id", "page_id": "page_id"})
```

#### Example: List

```python
components = client.Component().list({"page_id": "example"})
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
| `component_id` | `str` | Component identifier |
| `incidents` | `dict` | Related incidents |

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
| `component_group` | `dict` |  |
| `components` | `str` |  |
| `created_at` | `str` |  |
| `description` | `str` | Description of the component group. |
| `id` | `str` | Component Group Identifier |
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
group_components = client.GroupComponent().list({"page_id": "example"})
```

#### Example: Create

```python
group_component = client.GroupComponent().create({
    "page_id": "example_page_id",  # str
    "component_group": {},  # dict
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
| `auto_transition_deliver_notifications_at_end` | `bool` | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `bool` | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `bool` | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `bool` | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `list` | Incident components |
| `created_at` | `str` | The timestamp when the incident was created at. |
| `id` | `str` | Incident Identifier |
| `impact` | `str` | The impact of the incident. |
| `impact_override` | `str` | value to override calculated impact value |
| `incident` | `dict` |  |
| `incident_updates` | `list` | The incident updates for incident. |
| `metadata` | `dict` | Metadata attached to the incident. |
| `monitoring_at` | `str` | The timestamp when incident entered monitoring state. |
| `name` | `str` | Incident Name. |
| `page_id` | `str` | Incident Page Identifier |
| `postmortem_body` | `str` | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `str` | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `bool` | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `bool` | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `bool` | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `bool` | The timestamp when the postmortem was published. |
| `reminder_intervals` | `str` | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `str` | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `bool` | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `bool` | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `str` | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `bool` | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `str` | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `str` | The timestamp the incident is scheduled until. |
| `shortlink` | `str` | Incident Shortlink. |
| `status` | `str` | The incident status. |
| `updated_at` | `str` | The timestamp when the incident was updated at. |

#### Example: Load

```python
incident = client.Incident().load({"id": "incident_id", "page_id": "page_id"})
```

#### Example: List

```python
incidents = client.Incident().list({"page_id": "example"})
```

#### Example: Create

```python
incident = client.Incident().create({
    "page_id": "example_page_id",  # str
    "incident": {},  # dict
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
| `body` | `str` | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `list` | Affected components |
| `group_id` | `str` | Identifier of Template Group this template belongs to |
| `id` | `str` | Incident Template Identifier |
| `name` | `str` | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `bool` | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `bool` | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `dict` |  |
| `title` | `str` | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `str` | The status the incident or maintenance should transition to when selecting this template |

#### Example: List

```python
incident_templates = client.IncidentTemplate().list({"page_id": "example"})
```

#### Example: Create

```python
incident_template = client.IncidentTemplate().create({
    "page_id": "example_page_id",  # str
    "template": {},  # dict
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
| `affected_components` | `list` | Affected components associated with the incident update. |
| `body` | `str` | Incident update body. |
| `created_at` | `str` | The timestamp when the incident update was created at. |
| `custom_tweet` | `str` | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `bool` | Controls whether to delivery notifications. |
| `display_at` | `str` | Timestamp when incident update is happened. |
| `id` | `str` | Incident Update Identifier. |
| `incident_id` | `str` | Incident Identifier. |
| `incident_update` | `dict` |  |
| `status` | `str` | The incident status. |
| `tweet_id` | `str` | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `str` | The timestamp when twitter updated at. |
| `updated_at` | `str` | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `bool` | Controls whether to create twitter update. |


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
| `data` | `dict` | Add data points to metrics |
| `decimal_places` | `int` |  |
| `display` | `bool` | Should the metric be displayed |
| `id` | `str` | Metric identifier |
| `last_fetched_at` | `str` |  |
| `metric` | `dict` |  |
| `metric_identifier` | `str` | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `str` | Metric Provider identifier |
| `most_recent_data_at` | `str` |  |
| `name` | `str` | Name of metric |
| `reference_name` | `str` |  |
| `suffix` | `str` | Suffix to describe the units on the graph |
| `tooltip_description` | `str` |  |
| `updated_at` | `str` |  |
| `y_axis_hidden` | `bool` | Should the values on the y axis be hidden on render |
| `y_axis_max` | `float` |  |
| `y_axis_min` | `float` |  |

#### Example: Load

```python
metric = client.Metric().load({"id": "metric_id", "page_id": "page_id"})
```

#### Example: List

```python
metrics = client.Metric().list({"page_access_user_id": "example", "page_id": "example"})
```

#### Example: Create

```python
metric = client.Metric().create({
    "metrics_provider_id": "example_metrics_provider_id",  # str
    "page_id": "example_page_id",  # str
    "data": {},  # dict
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
| `id` | `str` | Identifier for Metrics Provider |
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
metrics_providers = client.MetricsProvider().list({"page_id": "example"})
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
| `allow_email_subscribers` | `bool` | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `bool` | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `bool` | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `bool` | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `bool` | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `bool` | Can your users choose to receive notifications via Webhooks |
| `branding` | `str` | The main template your statuspage will use |
| `city` | `str` |  |
| `country` | `str` |  |
| `created_at` | `str` | Timestamp the record was created |
| `css_blues` | `str` | CSS Color |
| `css_body_background_color` | `str` | CSS Color |
| `css_border_color` | `str` | CSS Color |
| `css_font_color` | `str` | CSS Color |
| `css_graph_color` | `str` | CSS Color |
| `css_greens` | `str` | CSS Color |
| `css_light_font_color` | `str` | CSS Color |
| `css_link_color` | `str` | CSS Color |
| `css_no_data` | `str` | CSS Color |
| `css_oranges` | `str` | CSS Color |
| `css_reds` | `str` | CSS Color |
| `css_yellows` | `str` | CSS Color |
| `domain` | `str` | CNAME alias for your status page |
| `email_logo` | `str` |  |
| `favicon_logo` | `str` |  |
| `headline` | `str` |  |
| `hero_cover` | `str` |  |
| `hidden_from_search` | `bool` | Should your page hide itself from search engines |
| `id` | `str` | Page identifier |
| `ip_restrictions` | `str` |  |
| `name` | `str` | Name of your page to be displayed |
| `notifications_email_footer` | `str` | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `str` | Allows you to customize the email address your page notifications come from |
| `page` | `dict` |  |
| `page_description` | `str` |  |
| `state` | `str` |  |
| `subdomain` | `str` | Subdomain at which to access your status page |
| `support_url` | `str` |  |
| `time_zone` | `str` | Timezone configured for your page |
| `transactional_logo` | `str` |  |
| `twitter_logo` | `str` |  |
| `twitter_username` | `str` |  |
| `updated_at` | `str` | Timestamp the record was last updated |
| `url` | `str` | Website of your page. |
| `viewers_must_be_team_members` | `bool` |  |

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
| `component_ids` | `list` | List of components codes to set on the page access group |
| `created_at` | `str` |  |
| `external_identifier` | `str` | Associates group with external group. |
| `id` | `str` | Page Access Group Identifier |
| `metric_ids` | `list` |  |
| `name` | `str` | Name for this Group. |
| `page_access_group` | `dict` |  |
| `page_access_user_ids` | `list` |  |
| `page_id` | `str` | Page Identifier. |
| `updated_at` | `str` |  |

#### Example: Load

```python
page_access_group = client.PageAccessGroup().load({"id": "page_access_group_id", "page_id": "page_id"})
```

#### Example: List

```python
page_access_groups = client.PageAccessGroup().list({"id": "example_id"})
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
| `component_ids` | `list` | List of component codes to allow access to |
| `created_at` | `str` |  |
| `email` | `str` |  |
| `external_login` | `str` | IDP login user id. |
| `id` | `str` | Page Access User Identifier |
| `metric_ids` | `list` | List of metrics to add |
| `page_access_group_id` | `str` |  |
| `page_access_group_ids` | `str` |  |
| `page_access_user` | `dict` |  |
| `page_id` | `str` |  |
| `updated_at` | `str` |  |

#### Example: Load

```python
page_access_user = client.PageAccessUser().load({"id": "page_access_user_id", "page_id": "page_id"})
```

#### Example: List

```python
page_access_users = client.PageAccessUser().list({"id": "example_id"})
```

#### Example: Create

```python
page_access_user = client.PageAccessUser().create({
    "id": "example_id",  # str
    "component_ids": [],  # list
    "metric_ids": [],  # list
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
| `pages` | `dict` | Pages accessible by the user. |
| `user_id` | `str` | User identifier |

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
| `body` | `str` | Postmortem body |
| `body_draft` | `str` | Body draft |
| `body_draft_updated_at` | `str` |  |
| `body_updated_at` | `str` |  |
| `created_at` | `str` |  |
| `custom_tweet` | `str` | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `bool` | Should email subscribers be notified. |
| `notify_twitter` | `bool` | Should Twitter followers be notified. |
| `postmortem` | `dict` |  |
| `preview_key` | `str` | Preview Key |
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
| `incident_background_color` | `str` | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `str` | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `str` | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `str` | Color of status embed iframe text when displaying maintenance |
| `page_id` | `str` | Page identifier |
| `position` | `str` | Corner where status embed iframe will appear on page |
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
| `component_ids` | `list` | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `str` | The components for which the subscriber has elected to receive updates. |
| `created_at` | `str` |  |
| `display_phone_number` | `str` | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `str` | The email address to use to contact the subscriber. |
| `endpoint` | `str` | The URL where a webhook subscriber elects to receive updates. |
| `id` | `str` | Subscriber Identifier |
| `integration_partner` | `int` | The number of integration partners found by the query. |
| `mode` | `str` | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `str` | Obfuscated slack channel name |
| `page_access_user_id` | `str` | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `str` | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `str` | The phone number used to contact an SMS subscriber |
| `purge_at` | `str` | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `str` | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `bool` | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `bool` | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `int` | The number of Slack subscribers found by the query. |
| `sms` | `int` | The number of Webhook subscribers found by the query. |
| `state` | `str` | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `dict` |  |
| `subscribers` | `str` | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `int` | The number of MS teams subscribers found by the query. |
| `type` | `str` | If this is present, only reactivate subscribers of this type. |
| `webhook` | `int` | The number of SMS subscribers found by the query. |
| `workspace_name` | `str` | The workspace name of the slack subscriber. |

#### Example: Load

```python
subscriber = client.Subscriber().load({"id": "subscriber_id", "page_id": "page_id"})
```

#### Example: List

```python
subscribers = client.Subscriber().list({"page_id": "example"})
```

#### Example: Create

```python
subscriber = client.Subscriber().create({
    "page_id": "example_page_id",  # str
    "subscribers": "example_subscribers",  # str
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
| `email` | `str` | Email address for the team member |
| `first_name` | `str` |  |
| `id` | `str` | User identifier |
| `last_name` | `str` |  |
| `organization_id` | `str` | Organization identifier |
| `updated_at` | `str` |  |
| `user` | `dict` |  |

#### Example: List

```python
users = client.User().list({"organization_id": "example"})
```

#### Example: Create

```python
user = client.User().create({
    "organization_id": "example_organization_id",  # str
    "user": {},  # dict
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

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```python
postmortem = client.Postmortem()
postmortem.load({"incident_id": "example", "page_id": "example"})

# postmortem.data_get() now returns the postmortem data from the last load
# postmortem.match_get() returns the last match criteria
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
