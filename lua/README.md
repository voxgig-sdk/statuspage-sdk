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
client:Component():update({ id = created:data_get()["id"], page_id = "example_page_id", automation_email = "example_automation_email" })

-- Remove
client:Component():remove({ id = created:data_get()["id"], page_id = "example_page_id" })
```


## Error handling

Entity operations return `(value, err)`. Check `err` before using
the value:

```lua
local postmortem, err = client:Postmortem():load({ incident_id = "example", page_id = "example" })
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

local result, err = client:Postmortem():load({ incident_id = "example", page_id = "example" })
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
| `automation_email` | `string` | Requires a special feature flag to be enabled |
| `component` | `table` |  |
| `created_at` | `string` |  |
| `description` | `string` | More detailed description for component |
| `group` | `boolean` | Is this component a group |
| `group_id` | `string` | Component Group identifier |
| `id` | `string` | Incident identifier |
| `name` | `string` | Display name for component |
| `only_show_if_degraded` | `boolean` | Requires a special feature flag to be enabled |
| `page_id` | `string` | Page identifier |
| `position` | `number` | Order the component will appear on the page |
| `showcase` | `boolean` | Should this component be showcased |
| `start_date` | `string` | The date this component started being used |
| `status` | `string` | Status of component |
| `updated_at` | `string` |  |

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
| `component_id` | `string` | Component identifier |
| `incidents` | `table` | Related incidents |

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
| `component_group` | `table` |  |
| `components` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` | Description of the component group. |
| `id` | `string` | Component Group Identifier |
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
  component_group = {}, -- table
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
| `auto_transition_deliver_notifications_at_end` | `boolean` | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `boolean` | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `boolean` | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `boolean` | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `table` | Incident components |
| `created_at` | `string` | The timestamp when the incident was created at. |
| `id` | `string` | Incident Identifier |
| `impact` | `string` | The impact of the incident. |
| `impact_override` | `string` | value to override calculated impact value |
| `incident` | `table` |  |
| `incident_updates` | `table` | The incident updates for incident. |
| `metadata` | `table` | Metadata attached to the incident. |
| `monitoring_at` | `string` | The timestamp when incident entered monitoring state. |
| `name` | `string` | Incident Name. |
| `page_id` | `string` | Incident Page Identifier |
| `postmortem_body` | `string` | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `string` | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `boolean` | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `boolean` | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `boolean` | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `boolean` | The timestamp when the postmortem was published. |
| `reminder_intervals` | `string` | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `string` | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `boolean` | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `boolean` | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `string` | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `boolean` | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `string` | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `string` | The timestamp the incident is scheduled until. |
| `shortlink` | `string` | Incident Shortlink. |
| `status` | `string` | The incident status. |
| `updated_at` | `string` | The timestamp when the incident was updated at. |

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
  incident = {}, -- table
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
| `body` | `string` | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `table` | Affected components |
| `group_id` | `string` | Identifier of Template Group this template belongs to |
| `id` | `string` | Incident Template Identifier |
| `name` | `string` | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `boolean` | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `boolean` | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `table` |  |
| `title` | `string` | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `string` | The status the incident or maintenance should transition to when selecting this template |

#### Example: List

```lua
local incident_templates, err = client:IncidentTemplate():list()
```

#### Example: Create

```lua
local incident_template, err = client:IncidentTemplate():create({
  page_id = "example_page_id", -- string
  template = {}, -- table
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
| `affected_components` | `table` | Affected components associated with the incident update. |
| `body` | `string` | Incident update body. |
| `created_at` | `string` | The timestamp when the incident update was created at. |
| `custom_tweet` | `string` | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `boolean` | Controls whether to delivery notifications. |
| `display_at` | `string` | Timestamp when incident update is happened. |
| `id` | `string` | Incident Update Identifier. |
| `incident_id` | `string` | Incident Identifier. |
| `incident_update` | `table` |  |
| `status` | `string` | The incident status. |
| `tweet_id` | `string` | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `string` | The timestamp when twitter updated at. |
| `updated_at` | `string` | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `boolean` | Controls whether to create twitter update. |


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
| `data` | `table` | Add data points to metrics |
| `decimal_places` | `number` |  |
| `display` | `boolean` | Should the metric be displayed |
| `id` | `string` | Metric identifier |
| `last_fetched_at` | `string` |  |
| `metric` | `table` |  |
| `metric_identifier` | `string` | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `string` | Metric Provider identifier |
| `most_recent_data_at` | `string` |  |
| `name` | `string` | Name of metric |
| `reference_name` | `string` |  |
| `suffix` | `string` | Suffix to describe the units on the graph |
| `tooltip_description` | `string` |  |
| `updated_at` | `string` |  |
| `y_axis_hidden` | `boolean` | Should the values on the y axis be hidden on render |
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
  data = {}, -- table
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
| `id` | `string` | Identifier for Metrics Provider |
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
| `allow_email_subscribers` | `boolean` | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `boolean` | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `boolean` | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `boolean` | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `boolean` | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `boolean` | Can your users choose to receive notifications via Webhooks |
| `branding` | `string` | The main template your statuspage will use |
| `city` | `string` |  |
| `country` | `string` |  |
| `created_at` | `string` | Timestamp the record was created |
| `css_blues` | `string` | CSS Color |
| `css_body_background_color` | `string` | CSS Color |
| `css_border_color` | `string` | CSS Color |
| `css_font_color` | `string` | CSS Color |
| `css_graph_color` | `string` | CSS Color |
| `css_greens` | `string` | CSS Color |
| `css_light_font_color` | `string` | CSS Color |
| `css_link_color` | `string` | CSS Color |
| `css_no_data` | `string` | CSS Color |
| `css_oranges` | `string` | CSS Color |
| `css_reds` | `string` | CSS Color |
| `css_yellows` | `string` | CSS Color |
| `domain` | `string` | CNAME alias for your status page |
| `email_logo` | `string` |  |
| `favicon_logo` | `string` |  |
| `headline` | `string` |  |
| `hero_cover` | `string` |  |
| `hidden_from_search` | `boolean` | Should your page hide itself from search engines |
| `id` | `string` | Page identifier |
| `ip_restrictions` | `string` |  |
| `name` | `string` | Name of your page to be displayed |
| `notifications_email_footer` | `string` | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `string` | Allows you to customize the email address your page notifications come from |
| `page` | `table` |  |
| `page_description` | `string` |  |
| `state` | `string` |  |
| `subdomain` | `string` | Subdomain at which to access your status page |
| `support_url` | `string` |  |
| `time_zone` | `string` | Timezone configured for your page |
| `transactional_logo` | `string` |  |
| `twitter_logo` | `string` |  |
| `twitter_username` | `string` |  |
| `updated_at` | `string` | Timestamp the record was last updated |
| `url` | `string` | Website of your page. |
| `viewers_must_be_team_members` | `boolean` |  |

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
| `component_ids` | `table` | List of components codes to set on the page access group |
| `created_at` | `string` |  |
| `external_identifier` | `string` | Associates group with external group. |
| `id` | `string` | Page Access Group Identifier |
| `metric_ids` | `table` |  |
| `name` | `string` | Name for this Group. |
| `page_access_group` | `table` |  |
| `page_access_user_ids` | `table` |  |
| `page_id` | `string` | Page Identifier. |
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
| `component_ids` | `table` | List of component codes to allow access to |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `external_login` | `string` | IDP login user id. |
| `id` | `string` | Page Access User Identifier |
| `metric_ids` | `table` | List of metrics to add |
| `page_access_group_id` | `string` |  |
| `page_access_group_ids` | `string` |  |
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
  component_ids = {}, -- table
  metric_ids = {}, -- table
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
| `pages` | `table` | Pages accessible by the user. |
| `user_id` | `string` | User identifier |

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
| `body` | `string` | Postmortem body |
| `body_draft` | `string` | Body draft |
| `body_draft_updated_at` | `string` |  |
| `body_updated_at` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `boolean` | Should email subscribers be notified. |
| `notify_twitter` | `boolean` | Should Twitter followers be notified. |
| `postmortem` | `table` |  |
| `preview_key` | `string` | Preview Key |
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
| `incident_background_color` | `string` | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `string` | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `string` | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `string` | Color of status embed iframe text when displaying maintenance |
| `page_id` | `string` | Page identifier |
| `position` | `string` | Corner where status embed iframe will appear on page |
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
| `component_ids` | `table` | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `string` | The components for which the subscriber has elected to receive updates. |
| `created_at` | `string` |  |
| `display_phone_number` | `string` | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `string` | The email address to use to contact the subscriber. |
| `endpoint` | `string` | The URL where a webhook subscriber elects to receive updates. |
| `id` | `string` | Subscriber Identifier |
| `integration_partner` | `number` | The number of integration partners found by the query. |
| `mode` | `string` | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `string` | Obfuscated slack channel name |
| `page_access_user_id` | `string` | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `string` | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `string` | The phone number used to contact an SMS subscriber |
| `purge_at` | `string` | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `string` | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `boolean` | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `boolean` | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `number` | The number of Slack subscribers found by the query. |
| `sms` | `number` | The number of Webhook subscribers found by the query. |
| `state` | `string` | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `table` |  |
| `subscribers` | `string` | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `number` | The number of MS teams subscribers found by the query. |
| `type` | `string` | If this is present, only reactivate subscribers of this type. |
| `webhook` | `number` | The number of SMS subscribers found by the query. |
| `workspace_name` | `string` | The workspace name of the slack subscriber. |

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
  subscribers = "example_subscribers", -- string
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
| `email` | `string` | Email address for the team member |
| `first_name` | `string` |  |
| `id` | `string` | User identifier |
| `last_name` | `string` |  |
| `organization_id` | `string` | Organization identifier |
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
  user = {}, -- table
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

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```lua
local postmortem = client:Postmortem()
postmortem:load({ incident_id = "example", page_id = "example" })

-- postmortem:data_get() now returns the postmortem data from the last load
-- postmortem:match_get() returns the last match criteria
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
