# Statuspage Lua SDK



The Lua SDK for the Statuspage API — an entity-oriented client using Lua conventions.

It exposes the API as capitalised, semantic **Entities** — e.g. `client:Component()` — each with the same small set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to LuaRocks. Install it from the
GitHub release tag (`lua/vX.Y.Z`, see [Releases](https://github.com/voxgig-sdk/statuspage-sdk/releases)),
or add the source directory to your `LUA_PATH`:

```bash
export LUA_PATH="path/to/lua/?.lua;path/to/lua/?/init.lua;;"
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```lua
local sdk = require("statuspage_sdk")

local client = sdk.new({
  apikey = os.getenv("STATUSPAGE_APIKEY"),
})
```

### 2. List component records

Entity operations return `(value, err)`. For `list`, `value` is the
array of records itself — iterate it directly (there is no wrapper).

```lua
local components, err = client:Component():list()
if err then error(err) end

for _, item in ipairs(components) do
  print(item["id"], item["automation_email"])
end
```

### 3. Load a component

Component is nested under page, so provide the `page_id`.

```lua
local component, err = client:Component():load({ page_id = "example_page_id", id = "example_id" })
if err then error(err) end
print(component)
```

### 4. Create, update, and remove

```lua
-- Create
local created, err = client:Component():create({ page_id = "example_page_id" })
if err then error(err) end

-- Update
client:Component():update({ id = created["id"], page_id = "example_page_id" })

-- Remove
client:Component():remove({ id = created["id"], page_id = "example_page_id" })
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local components, err = client:Component():list()
if err then error(err) end
```

`direct` follows the same `(value, err)` convention:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example_id" },
})
if err then error(err) end
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```lua
local result, err = client:direct({
  path = "/api/resource/{id}",
  method = "GET",
  params = { id = "example" },
})
if err then error(err) end

if result["ok"] then
  print(result["status"])  -- 200
  print(result["data"])    -- response body
end
```

### Prepare a request without sending it

```lua
local fetchdef, err = client:prepare({
  path = "/api/resource/{id}",
  method = "DELETE",
  params = { id = "example" },
})
if err then error(err) end

print(fetchdef["url"])
print(fetchdef["method"])
print(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```lua
local client = sdk.test()

local result, err = client:Component():list()
-- result is the returned data; err is set on failure
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```lua
local function mock_fetch(url, init)
  return {
    status = 200,
    statusText = "OK",
    headers = {},
    json = function()
      return { id = "mock01" }
    end,
  }, nil
end

local client = sdk.new({
  base = "http://localhost:8080",
  system = {
    fetch = mock_fetch,
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
cd lua && busted test/
```


## Reference

### StatuspageSDK

```lua
local sdk = require("statuspage_sdk")
local client = sdk.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `table` | Feature activation flags. |
| `extend` | `table` | Additional Feature instances to load. |
| `system` | `table` | System overrides (e.g. custom `fetch` function). |

### test

```lua
local client = sdk.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### StatuspageSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> table` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> table, err` | Build an HTTP request definition without sending. |
| `direct` | `(fetchargs) -> table, err` | Build and send an HTTP request. |
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
| `load` | `(reqmatch, ctrl) -> any, err` | Load a single entity by match criteria. |
| `list` | `(reqmatch, ctrl) -> any, err` | List entities matching the criteria. |
| `create` | `(reqdata, ctrl) -> any, err` | Create a new entity. |
| `update` | `(reqdata, ctrl) -> any, err` | Update an existing entity. |
| `remove` | `(reqmatch, ctrl) -> any, err` | Remove an entity. |
| `data_get` | `() -> table` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> table` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> string` | Return the entity name. |

### Result shape

Entity operations return `(value, err)`. The `value` is the operation's
data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `load` / `create` / `update` / `remove` | the entity record (a `table`) |
| `list` | an array (`table`) of entity records |

Check `err` first (it is non-`nil` on failure), then use `value`:

    local component, err = client:Component():load({ id = "example_id" })
    if err then error(err) end
    -- component is the loaded record

Only `direct()` returns a response envelope — a `table` with `ok`,
`status`, `headers`, and `data` keys.

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

Create an instance: `local component = client:Component(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `automation_email` | `string` |  |
| `component` | `table` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `group` | `boolean` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `major_outage` | `number` |  |
| `name` | `string` |  |
| `only_show_if_degraded` | `boolean` |  |
| `page_id` | `string` |  |
| `partial_outage` | `number` |  |
| `position` | `number` |  |
| `range_end` | `string` |  |
| `range_start` | `string` |  |
| `related_event` | `table` |  |
| `showcase` | `boolean` |  |
| `start_date` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |
| `uptime_percentage` | `number` |  |
| `warning` | `string` |  |

#### Example: Load

```lua
local component, err = client:Component():load({ id = "component_id", page_id = "page_id" })
```

#### Example: List

```lua
local components, err = client:Component():list()
```

#### Example: Create

```lua
local component, err = client:Component():create({
  page_id = "example_page_id", -- string
})
```


### ComponentGroupUptime

Create an instance: `local component_group_uptime = client:ComponentGroupUptime(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `major_outage` | `number` |  |
| `name` | `string` |  |
| `partial_outage` | `number` |  |
| `range_end` | `string` |  |
| `range_start` | `string` |  |
| `related_event` | `table` |  |
| `uptime_percentage` | `number` |  |
| `warning` | `string` |  |

#### Example: Load

```lua
local component_group_uptime, err = client:ComponentGroupUptime():load({ id = "component_group_uptime_id", page_id = "page_id" })
```


### GroupComponent

Create an instance: `local group_component = client:GroupComponent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `string` |  |
| `component_group` | `table` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local group_component, err = client:GroupComponent():load({ id = "group_component_id", page_id = "page_id" })
```

#### Example: List

```lua
local group_components, err = client:GroupComponent():list()
```

#### Example: Create

```lua
local group_component, err = client:GroupComponent():create({
  page_id = "example_page_id", -- string
})
```


### Incident

Create an instance: `local incident = client:Incident(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `boolean` |  |
| `auto_transition_deliver_notifications_at_start` | `boolean` |  |
| `auto_transition_to_maintenance_state` | `boolean` |  |
| `auto_transition_to_operational_state` | `boolean` |  |
| `component` | `table` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `impact` | `string` |  |
| `impact_override` | `string` |  |
| `incident` | `table` |  |
| `incident_impact` | `table` |  |
| `incident_update` | `table` |  |
| `metadata` | `table` |  |
| `monitoring_at` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `postmortem_body` | `string` |  |
| `postmortem_body_last_updated_at` | `string` |  |
| `postmortem_ignored` | `boolean` |  |
| `postmortem_notified_subscriber` | `boolean` |  |
| `postmortem_notified_twitter` | `boolean` |  |
| `postmortem_published_at` | `boolean` |  |
| `reminder_interval` | `string` |  |
| `resolved_at` | `string` |  |
| `scheduled_auto_completed` | `boolean` |  |
| `scheduled_auto_in_progress` | `boolean` |  |
| `scheduled_for` | `string` |  |
| `scheduled_remind_prior` | `boolean` |  |
| `scheduled_reminded_at` | `string` |  |
| `scheduled_until` | `string` |  |
| `shortlink` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local incident, err = client:Incident():load({ id = "incident_id", page_id = "page_id" })
```

#### Example: List

```lua
local incidents, err = client:Incident():list()
```

#### Example: Create

```lua
local incident, err = client:Incident():create({
  page_id = "example_page_id", -- string
})
```


### IncidentPostmortem

Create an instance: `local incident_postmortem = client:IncidentPostmortem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `local incident_subscriber = client:IncidentSubscriber(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```lua
local incident_subscriber, err = client:IncidentSubscriber():create({
  incident_id = "example_incident_id", -- string
  page_id = "example_page_id", -- string
  subscriber_id = "example_subscriber_id", -- string
})
```


### IncidentTemplate

Create an instance: `local incident_template = client:IncidentTemplate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` |  |
| `component` | `table` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `should_send_notification` | `boolean` |  |
| `should_tweet` | `boolean` |  |
| `template` | `table` |  |
| `title` | `string` |  |
| `update_status` | `string` |  |

#### Example: List

```lua
local incident_templates, err = client:IncidentTemplate():list()
```

#### Example: Create

```lua
local incident_template, err = client:IncidentTemplate():create({
  page_id = "example_page_id", -- string
})
```


### IncidentUpdate

Create an instance: `local incident_update = client:IncidentUpdate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_component` | `table` |  |
| `body` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `deliver_notification` | `boolean` |  |
| `display_at` | `string` |  |
| `id` | `string` |  |
| `incident_id` | `string` |  |
| `incident_update` | `table` |  |
| `status` | `string` |  |
| `tweet_id` | `string` |  |
| `twitter_updated_at` | `string` |  |
| `updated_at` | `string` |  |
| `wants_twitter_update` | `boolean` |  |


### Metric

Create an instance: `local metric = client:Metric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `backfill_percentage` | `number` |  |
| `backfilled` | `boolean` |  |
| `created_at` | `string` |  |
| `data` | `table` |  |
| `decimal_place` | `number` |  |
| `display` | `boolean` |  |
| `id` | `string` |  |
| `last_fetched_at` | `string` |  |
| `metric` | `table` |  |
| `metric_identifier` | `string` |  |
| `metrics_provider_id` | `string` |  |
| `most_recent_data_at` | `string` |  |
| `name` | `string` |  |
| `reference_name` | `string` |  |
| `suffix` | `string` |  |
| `tooltip_description` | `string` |  |
| `updated_at` | `string` |  |
| `y_axis_hidden` | `boolean` |  |
| `y_axis_max` | `number` |  |
| `y_axis_min` | `number` |  |

#### Example: Load

```lua
local metric, err = client:Metric():load({ id = "metric_id", page_id = "page_id" })
```

#### Example: List

```lua
local metrics, err = client:Metric():list()
```

#### Example: Create

```lua
local metric, err = client:Metric():create({
  metrics_provider_id = "example_metrics_provider_id", -- string
  page_id = "example_page_id", -- string
})
```


### MetricsProvider

Create an instance: `local metrics_provider = client:MetricsProvider(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `disabled` | `boolean` |  |
| `id` | `string` |  |
| `last_revalidated_at` | `string` |  |
| `metric_base_uri` | `string` |  |
| `metrics_provider` | `table` |  |
| `page_id` | `number` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local metrics_provider, err = client:MetricsProvider():load({ id = "metrics_provider_id", page_id = "page_id" })
```

#### Example: List

```lua
local metrics_providers, err = client:MetricsProvider():list()
```

#### Example: Create

```lua
local metrics_provider, err = client:MetricsProvider():create({
  page_id = "example_page_id", -- string
})
```


### Page

Create an instance: `local page = client:Page(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activity_score` | `number` |  |
| `allow_email_subscriber` | `boolean` |  |
| `allow_incident_subscriber` | `boolean` |  |
| `allow_page_subscriber` | `boolean` |  |
| `allow_rss_atom_feed` | `boolean` |  |
| `allow_sms_subscriber` | `boolean` |  |
| `allow_webhook_subscriber` | `boolean` |  |
| `branding` | `string` |  |
| `city` | `string` |  |
| `country` | `string` |  |
| `created_at` | `string` |  |
| `css_blue` | `string` |  |
| `css_body_background_color` | `string` |  |
| `css_border_color` | `string` |  |
| `css_font_color` | `string` |  |
| `css_graph_color` | `string` |  |
| `css_green` | `string` |  |
| `css_light_font_color` | `string` |  |
| `css_link_color` | `string` |  |
| `css_no_data` | `string` |  |
| `css_orange` | `string` |  |
| `css_red` | `string` |  |
| `css_yellow` | `string` |  |
| `domain` | `string` |  |
| `email_logo` | `string` |  |
| `favicon_logo` | `string` |  |
| `headline` | `string` |  |
| `hero_cover` | `string` |  |
| `hidden_from_search` | `boolean` |  |
| `id` | `string` |  |
| `ip_restriction` | `string` |  |
| `name` | `string` |  |
| `notifications_email_footer` | `string` |  |
| `notifications_from_email` | `string` |  |
| `page` | `table` |  |
| `page_description` | `string` |  |
| `state` | `string` |  |
| `subdomain` | `string` |  |
| `support_url` | `string` |  |
| `time_zone` | `string` |  |
| `transactional_logo` | `string` |  |
| `twitter_logo` | `string` |  |
| `twitter_username` | `string` |  |
| `updated_at` | `string` |  |
| `url` | `string` |  |
| `viewers_must_be_team_member` | `boolean` |  |

#### Example: Load

```lua
local page, err = client:Page():load({ id = "page_id" })
```

#### Example: List

```lua
local pages, err = client:Page():list()
```


### PageAccessGroup

Create an instance: `local page_access_group = client:PageAccessGroup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `table` |  |
| `created_at` | `string` |  |
| `external_identifier` | `string` |  |
| `id` | `string` |  |
| `metric_id` | `table` |  |
| `name` | `string` |  |
| `page_access_group` | `table` |  |
| `page_access_user_id` | `table` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local page_access_group, err = client:PageAccessGroup():load({ id = "page_access_group_id", page_id = "page_id" })
```

#### Example: List

```lua
local page_access_groups, err = client:PageAccessGroup():list()
```

#### Example: Create

```lua
local page_access_group, err = client:PageAccessGroup():create({
  id = "example_id", -- string
})
```


### PageAccessUser

Create an instance: `local page_access_user = client:PageAccessUser(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `table` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `external_login` | `string` |  |
| `id` | `string` |  |
| `metric_id` | `table` |  |
| `page_access_group_id` | `string` |  |
| `page_access_user` | `table` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local page_access_user, err = client:PageAccessUser():load({ id = "page_access_user_id", page_id = "page_id" })
```

#### Example: List

```lua
local page_access_users, err = client:PageAccessUser():list()
```

#### Example: Create

```lua
local page_access_user, err = client:PageAccessUser():create({
  id = "example_id", -- string
})
```


### Permission

Create an instance: `local permission = client:Permission(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `table` |  |
| `page` | `table` |  |

#### Example: Load

```lua
local permission, err = client:Permission():load({ id = "permission_id", organization_id = "organization_id" })
```


### Postmortem

Create an instance: `local postmortem = client:Postmortem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` |  |
| `body_draft` | `string` |  |
| `body_draft_updated_at` | `string` |  |
| `body_updated_at` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `notify_subscriber` | `boolean` |  |
| `notify_twitter` | `boolean` |  |
| `postmortem` | `table` |  |
| `preview_key` | `string` |  |
| `published_at` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```lua
local postmortem, err = client:Postmortem():load({ incident_id = "incident_id", page_id = "page_id" })
```


### StatusEmbedConfig

Create an instance: `local status_embed_config = client:StatusEmbedConfig(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `incident_background_color` | `string` |  |
| `incident_text_color` | `string` |  |
| `maintenance_background_color` | `string` |  |
| `maintenance_text_color` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `status_embed_config` | `table` |  |

#### Example: Load

```lua
local status_embed_config, err = client:StatusEmbedConfig():load({ page_id = "page_id" })
```


### Subscriber

Create an instance: `local subscriber = client:Subscriber(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `remove(match)` | Remove the matching entity. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `string` |  |
| `component_id` | `table` |  |
| `created_at` | `string` |  |
| `display_phone_number` | `string` |  |
| `email` | `string` |  |
| `endpoint` | `string` |  |
| `id` | `string` |  |
| `integration_partner` | `number` |  |
| `mode` | `string` |  |
| `obfuscated_channel_name` | `string` |  |
| `page_access_user_id` | `string` |  |
| `phone_country` | `string` |  |
| `phone_number` | `string` |  |
| `purge_at` | `string` |  |
| `quarantined_at` | `string` |  |
| `skip_confirmation_notification` | `boolean` |  |
| `skip_unsubscription_notification` | `boolean` |  |
| `slack` | `number` |  |
| `sms` | `number` |  |
| `state` | `string` |  |
| `subscriber` | `table` |  |
| `team` | `number` |  |
| `type` | `string` |  |
| `webhook` | `number` |  |
| `workspace_name` | `string` |  |

#### Example: Load

```lua
local subscriber, err = client:Subscriber():load({ id = "subscriber_id", page_id = "page_id" })
```

#### Example: List

```lua
local subscribers, err = client:Subscriber():list()
```

#### Example: Create

```lua
local subscriber, err = client:Subscriber():create({
  page_id = "example_page_id", -- string
})
```


### User

Create an instance: `local user = client:User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `first_name` | `string` |  |
| `id` | `string` |  |
| `last_name` | `string` |  |
| `organization_id` | `string` |  |
| `updated_at` | `string` |  |
| `user` | `table` |  |

#### Example: List

```lua
local users, err = client:User():list()
```

#### Example: Create

```lua
local user, err = client:User():create({
  organization_id = "example_organization_id", -- string
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

Features are the extension mechanism. A feature is a Lua table
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as tables

The Lua SDK uses plain Lua tables throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `helpers.to_map()` to safely validate that a value is a table.

### Module structure

```
lua/
├── statuspage_sdk.lua    -- Main SDK module
├── config.lua               -- Configuration
├── features.lua             -- Feature factory
├── core/                    -- Core types and context
├── entity/                  -- Entity implementations
├── feature/                 -- Built-in features (Base, Test, Log)
├── utility/                 -- Utility functions and struct library
└── test/                    -- Test suites
```

The main module (`statuspage_sdk`) exports the SDK constructor
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```lua
local component = client:Component()
component:list()

-- component:data_get() now returns the component data from the last list
-- component:match_get() returns the last match criteria
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
