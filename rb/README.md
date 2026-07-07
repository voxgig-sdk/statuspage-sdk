# Statuspage Ruby SDK



The Ruby SDK for the Statuspage API — an entity-oriented client using idiomatic Ruby conventions.

The SDK exposes the API as capitalised, semantic **Entities** — for example `client.Component` — with named operations (`list`/`load`/`create`/`update`/`remove`/`patch`) instead of raw URL paths and query strings. Working with resources and verbs keeps call sites self-describing and reduces cognitive load.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
This package is not yet published to RubyGems. Install it from the
GitHub release tag (`rb/vX.Y.Z`):

- Releases: [https://github.com/voxgig-sdk/statuspage-sdk/releases](https://github.com/voxgig-sdk/statuspage-sdk/releases)


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### 1. Create a client

```ruby
require_relative "Statuspage_sdk"

client = StatuspageSDK.new({
  "apikey" => ENV["STATUSPAGE_APIKEY"],
})
```

### 2. List component records

```ruby
begin
  # list returns an Array of Component records — iterate directly.
  components = client.Component.list
  components.each do |item|
    puts "#{item["id"]} #{item["automation_email"]}"
  end
rescue => err
  warn "list failed: #{err}"
end
```

### 3. Load a component

Component is nested under page, so provide the `page_id`.

```ruby
begin
  # load returns the bare Component record (raises on error).
  component = client.Component.load({ "page_id" => "example_page_id", "id" => "example_id" })
  puts component
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the bare created Component record.
created = client.Component.create({ "page_id" => "example_page_id" })

# Update — index the bare record directly (created["id"]).
client.Component.update({ "id" => created["id"], "page_id" => "example_page_id" })

# Remove
client.Component.remove({ "id" => created["id"], "page_id" => "example_page_id" })
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  components = client.Component.list()
rescue => err
  warn "list failed: #{err}"
end
```

`direct` does **not** raise — it returns the result hash. Branch on
`ok`; on failure `status` holds the HTTP status (for error responses) and
`err` holds a transport error, so read both defensively:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example_id" },
})

warn "request failed: #{result["err"] || "HTTP #{result["status"]}"}" unless result["ok"]
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```ruby
result = client.direct({
  "path" => "/api/resource/{id}",
  "method" => "GET",
  "params" => { "id" => "example" },
})

if result["ok"]
  puts result["status"]  # 200
  puts result["data"]    # response body
else
  # On an HTTP error status there is no err (only a transport failure sets
  # it), so fall back to the status code.
  warn(result["err"] || "HTTP #{result["status"]}")
end
```

### Prepare a request without sending it

```ruby
begin
  fetchdef = client.prepare({
    "path" => "/api/resource/{id}",
    "method" => "DELETE",
    "params" => { "id" => "example" },
  })
  puts fetchdef["url"]
  puts fetchdef["method"]
  puts fetchdef["headers"]
rescue => err
  warn "prepare failed: #{err}"
end
```

### Use test mode

Create a mock client for unit testing — no server required. Seed fixture
data via the `entity` option so offline calls resolve without a live server:

```ruby
client = StatuspageSDK.test({
  "entity" => { "component" => { "test01" => { "id" => "test01" } } },
})

# Entity ops return the bare mock record (raises on error).
component = client.Component.list()
puts component
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```ruby
mock_fetch = ->(url, init) {
  return {
    "status" => 200,
    "statusText" => "OK",
    "headers" => {},
    "json" => ->() { { "id" => "mock01" } },
  }, nil
}

client = StatuspageSDK.new({
  "base" => "http://localhost:8080",
  "system" => {
    "fetch" => mock_fetch,
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
cd rb && ruby -Itest -e "Dir['test/*_test.rb'].each { |f| require_relative f }"
```


## Reference

### StatuspageSDK

```ruby
require_relative "Statuspage_sdk"
client = StatuspageSDK.new(options)
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `String` | API key for authentication. |
| `base` | `String` | Base URL of the API server. |
| `prefix` | `String` | URL path prefix prepended to all requests. |
| `suffix` | `String` | URL path suffix appended to all requests. |
| `feature` | `Hash` | Feature activation flags. |
| `extend` | `Hash` | Additional Feature instances to load. |
| `system` | `Hash` | System overrides (e.g. custom `fetch` lambda). |

### test

```ruby
client = StatuspageSDK.test(testopts, sdkopts)
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### StatuspageSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `options_map` | `() -> Hash` | Deep copy of current SDK options. |
| `get_utility` | `() -> Utility` | Copy of the SDK utility object. |
| `prepare` | `(fetchargs) -> Hash` | Build an HTTP request definition without sending. Raises on error. |
| `direct` | `(fetchargs) -> Hash` | Build and send an HTTP request. Returns a result hash (`result["ok"]`); does not raise. |
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
| `list` | `(reqmatch = nil, ctrl) -> Array` | List entities matching the criteria (call with no argument to list all). Raises on error. |
| `create` | `(reqdata, ctrl) -> any` | Create a new entity. Raises on error. |
| `update` | `(reqdata, ctrl) -> any` | Update an existing entity. Raises on error. |
| `remove` | `(reqmatch, ctrl) -> any` | Remove an entity. Raises on error. |
| `data_get` | `() -> Hash` | Get entity data. |
| `data_set` | `(data)` | Set entity data. |
| `match_get` | `() -> Hash` | Get entity match criteria. |
| `match_set` | `(match)` | Set entity match criteria. |
| `make` | `() -> Entity` | Create a new instance with the same options. |
| `get_name` | `() -> String` | Return the entity name. |

### Result shape

Entity operations return the result data directly. On failure they
raise a `StatuspageError` (a `StandardError` subclass), so wrap
calls in `begin`/`rescue` where you need to handle errors.

The `direct` escape hatch is the exception: it never raises and instead
returns a result `Hash` with these keys:

| Key | Type | Description |
| --- | --- | --- |
| `ok` | `Boolean` | `true` if the HTTP status is 2xx. |
| `status` | `Integer` | HTTP status code. |
| `headers` | `Hash` | Response headers. |
| `data` | `any` | Parsed JSON response body. |
| `err` | `Error` | Present when `ok` is `false`. |

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

Create an instance: `component = client.Component`

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
| `automation_email` | `String` |  |
| `component` | `Hash` |  |
| `created_at` | `String` |  |
| `description` | `String` |  |
| `group` | `Boolean` |  |
| `group_id` | `String` |  |
| `id` | `String` |  |
| `major_outage` | `Integer` |  |
| `name` | `String` |  |
| `only_show_if_degraded` | `Boolean` |  |
| `page_id` | `String` |  |
| `partial_outage` | `Integer` |  |
| `position` | `Integer` |  |
| `range_end` | `String` |  |
| `range_start` | `String` |  |
| `related_event` | `Hash` |  |
| `showcase` | `Boolean` |  |
| `start_date` | `String` |  |
| `status` | `String` |  |
| `updated_at` | `String` |  |
| `uptime_percentage` | `Float` |  |
| `warning` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Component record (raises on error).
component = client.Component.load({ "id" => "component_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of Component records (raises on error).
components = client.Component.list
```

#### Example: Create

```ruby
component = client.Component.create({
  "page_id" => "example_page_id", # String
})
```


### ComponentGroupUptime

Create an instance: `component_group_uptime = client.ComponentGroupUptime`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `String` |  |
| `major_outage` | `Integer` |  |
| `name` | `String` |  |
| `partial_outage` | `Integer` |  |
| `range_end` | `String` |  |
| `range_start` | `String` |  |
| `related_event` | `Hash` |  |
| `uptime_percentage` | `Float` |  |
| `warning` | `String` |  |

#### Example: Load

```ruby
# load returns the bare ComponentGroupUptime record (raises on error).
component_group_uptime = client.ComponentGroupUptime.load({ "id" => "component_group_uptime_id", "page_id" => "page_id" })
```


### GroupComponent

Create an instance: `group_component = client.GroupComponent`

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
| `component` | `String` |  |
| `component_group` | `Hash` |  |
| `created_at` | `String` |  |
| `description` | `String` |  |
| `id` | `String` |  |
| `name` | `String` |  |
| `page_id` | `String` |  |
| `position` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the bare GroupComponent record (raises on error).
group_component = client.GroupComponent.load({ "id" => "group_component_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of GroupComponent records (raises on error).
group_components = client.GroupComponent.list
```

#### Example: Create

```ruby
group_component = client.GroupComponent.create({
  "page_id" => "example_page_id", # String
})
```


### Incident

Create an instance: `incident = client.Incident`

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
| `auto_transition_deliver_notifications_at_end` | `Boolean` |  |
| `auto_transition_deliver_notifications_at_start` | `Boolean` |  |
| `auto_transition_to_maintenance_state` | `Boolean` |  |
| `auto_transition_to_operational_state` | `Boolean` |  |
| `component` | `Array` |  |
| `created_at` | `String` |  |
| `id` | `String` |  |
| `impact` | `String` |  |
| `impact_override` | `String` |  |
| `incident` | `Hash` |  |
| `incident_impact` | `Array` |  |
| `incident_update` | `Array` |  |
| `metadata` | `Hash` |  |
| `monitoring_at` | `String` |  |
| `name` | `String` |  |
| `page_id` | `String` |  |
| `postmortem_body` | `String` |  |
| `postmortem_body_last_updated_at` | `String` |  |
| `postmortem_ignored` | `Boolean` |  |
| `postmortem_notified_subscriber` | `Boolean` |  |
| `postmortem_notified_twitter` | `Boolean` |  |
| `postmortem_published_at` | `Boolean` |  |
| `reminder_interval` | `String` |  |
| `resolved_at` | `String` |  |
| `scheduled_auto_completed` | `Boolean` |  |
| `scheduled_auto_in_progress` | `Boolean` |  |
| `scheduled_for` | `String` |  |
| `scheduled_remind_prior` | `Boolean` |  |
| `scheduled_reminded_at` | `String` |  |
| `scheduled_until` | `String` |  |
| `shortlink` | `String` |  |
| `status` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Incident record (raises on error).
incident = client.Incident.load({ "id" => "incident_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of Incident records (raises on error).
incidents = client.Incident.list
```

#### Example: Create

```ruby
incident = client.Incident.create({
  "page_id" => "example_page_id", # String
})
```


### IncidentPostmortem

Create an instance: `incident_postmortem = client.IncidentPostmortem`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `incident_subscriber = client.IncidentSubscriber`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```ruby
incident_subscriber = client.IncidentSubscriber.create({
  "incident_id" => "example_incident_id", # String
  "page_id" => "example_page_id", # String
  "subscriber_id" => "example_subscriber_id", # String
})
```


### IncidentTemplate

Create an instance: `incident_template = client.IncidentTemplate`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `String` |  |
| `component` | `Array` |  |
| `group_id` | `String` |  |
| `id` | `String` |  |
| `name` | `String` |  |
| `should_send_notification` | `Boolean` |  |
| `should_tweet` | `Boolean` |  |
| `template` | `Hash` |  |
| `title` | `String` |  |
| `update_status` | `String` |  |

#### Example: List

```ruby
# list returns an Array of IncidentTemplate records (raises on error).
incident_templates = client.IncidentTemplate.list
```

#### Example: Create

```ruby
incident_template = client.IncidentTemplate.create({
  "page_id" => "example_page_id", # String
})
```


### IncidentUpdate

Create an instance: `incident_update = client.IncidentUpdate`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_component` | `Array` |  |
| `body` | `String` |  |
| `created_at` | `String` |  |
| `custom_tweet` | `String` |  |
| `deliver_notification` | `Boolean` |  |
| `display_at` | `String` |  |
| `id` | `String` |  |
| `incident_id` | `String` |  |
| `incident_update` | `Hash` |  |
| `status` | `String` |  |
| `tweet_id` | `String` |  |
| `twitter_updated_at` | `String` |  |
| `updated_at` | `String` |  |
| `wants_twitter_update` | `Boolean` |  |


### Metric

Create an instance: `metric = client.Metric`

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
| `backfill_percentage` | `Integer` |  |
| `backfilled` | `Boolean` |  |
| `created_at` | `String` |  |
| `data` | `Hash` |  |
| `decimal_place` | `Integer` |  |
| `display` | `Boolean` |  |
| `id` | `String` |  |
| `last_fetched_at` | `String` |  |
| `metric` | `Hash` |  |
| `metric_identifier` | `String` |  |
| `metrics_provider_id` | `String` |  |
| `most_recent_data_at` | `String` |  |
| `name` | `String` |  |
| `reference_name` | `String` |  |
| `suffix` | `String` |  |
| `tooltip_description` | `String` |  |
| `updated_at` | `String` |  |
| `y_axis_hidden` | `Boolean` |  |
| `y_axis_max` | `Float` |  |
| `y_axis_min` | `Float` |  |

#### Example: Load

```ruby
# load returns the bare Metric record (raises on error).
metric = client.Metric.load({ "id" => "metric_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of Metric records (raises on error).
metrics = client.Metric.list
```

#### Example: Create

```ruby
metric = client.Metric.create({
  "metrics_provider_id" => "example_metrics_provider_id", # String
  "page_id" => "example_page_id", # String
})
```


### MetricsProvider

Create an instance: `metrics_provider = client.MetricsProvider`

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
| `created_at` | `String` |  |
| `disabled` | `Boolean` |  |
| `id` | `String` |  |
| `last_revalidated_at` | `String` |  |
| `metric_base_uri` | `String` |  |
| `metrics_provider` | `Hash` |  |
| `page_id` | `Integer` |  |
| `type` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the bare MetricsProvider record (raises on error).
metrics_provider = client.MetricsProvider.load({ "id" => "metrics_provider_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of MetricsProvider records (raises on error).
metrics_providers = client.MetricsProvider.list
```

#### Example: Create

```ruby
metrics_provider = client.MetricsProvider.create({
  "page_id" => "example_page_id", # String
})
```


### Page

Create an instance: `page = client.Page`

#### Operations

| Method | Description |
| --- | --- |
| `list(match)` | List entities matching the criteria. |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activity_score` | `Float` |  |
| `allow_email_subscriber` | `Boolean` |  |
| `allow_incident_subscriber` | `Boolean` |  |
| `allow_page_subscriber` | `Boolean` |  |
| `allow_rss_atom_feed` | `Boolean` |  |
| `allow_sms_subscriber` | `Boolean` |  |
| `allow_webhook_subscriber` | `Boolean` |  |
| `branding` | `String` |  |
| `city` | `String` |  |
| `country` | `String` |  |
| `created_at` | `String` |  |
| `css_blue` | `String` |  |
| `css_body_background_color` | `String` |  |
| `css_border_color` | `String` |  |
| `css_font_color` | `String` |  |
| `css_graph_color` | `String` |  |
| `css_green` | `String` |  |
| `css_light_font_color` | `String` |  |
| `css_link_color` | `String` |  |
| `css_no_data` | `String` |  |
| `css_orange` | `String` |  |
| `css_red` | `String` |  |
| `css_yellow` | `String` |  |
| `domain` | `String` |  |
| `email_logo` | `String` |  |
| `favicon_logo` | `String` |  |
| `headline` | `String` |  |
| `hero_cover` | `String` |  |
| `hidden_from_search` | `Boolean` |  |
| `id` | `String` |  |
| `ip_restriction` | `String` |  |
| `name` | `String` |  |
| `notifications_email_footer` | `String` |  |
| `notifications_from_email` | `String` |  |
| `page` | `Hash` |  |
| `page_description` | `String` |  |
| `state` | `String` |  |
| `subdomain` | `String` |  |
| `support_url` | `String` |  |
| `time_zone` | `String` |  |
| `transactional_logo` | `String` |  |
| `twitter_logo` | `String` |  |
| `twitter_username` | `String` |  |
| `updated_at` | `String` |  |
| `url` | `String` |  |
| `viewers_must_be_team_member` | `Boolean` |  |

#### Example: Load

```ruby
# load returns the bare Page record (raises on error).
page = client.Page.load({ "id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of Page records (raises on error).
pages = client.Page.list
```


### PageAccessGroup

Create an instance: `page_access_group = client.PageAccessGroup`

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
| `component_id` | `Array` |  |
| `created_at` | `String` |  |
| `external_identifier` | `String` |  |
| `id` | `String` |  |
| `metric_id` | `Array` |  |
| `name` | `String` |  |
| `page_access_group` | `Hash` |  |
| `page_access_user_id` | `Array` |  |
| `page_id` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the bare PageAccessGroup record (raises on error).
page_access_group = client.PageAccessGroup.load({ "id" => "page_access_group_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of PageAccessGroup records (raises on error).
page_access_groups = client.PageAccessGroup.list
```

#### Example: Create

```ruby
page_access_group = client.PageAccessGroup.create({
  "id" => "example_id", # String
})
```


### PageAccessUser

Create an instance: `page_access_user = client.PageAccessUser`

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
| `component_id` | `Array` |  |
| `created_at` | `String` |  |
| `email` | `String` |  |
| `external_login` | `String` |  |
| `id` | `String` |  |
| `metric_id` | `Array` |  |
| `page_access_group_id` | `String` |  |
| `page_access_user` | `Hash` |  |
| `page_id` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the bare PageAccessUser record (raises on error).
page_access_user = client.PageAccessUser.load({ "id" => "page_access_user_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of PageAccessUser records (raises on error).
page_access_users = client.PageAccessUser.list
```

#### Example: Create

```ruby
page_access_user = client.PageAccessUser.create({
  "id" => "example_id", # String
})
```


### Permission

Create an instance: `permission = client.Permission`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `data` | `Hash` |  |
| `page` | `Hash` |  |

#### Example: Load

```ruby
# load returns the bare Permission record (raises on error).
permission = client.Permission.load({ "id" => "permission_id", "organization_id" => "organization_id" })
```


### Postmortem

Create an instance: `postmortem = client.Postmortem`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `String` |  |
| `body_draft` | `String` |  |
| `body_draft_updated_at` | `String` |  |
| `body_updated_at` | `String` |  |
| `created_at` | `String` |  |
| `custom_tweet` | `String` |  |
| `notify_subscriber` | `Boolean` |  |
| `notify_twitter` | `Boolean` |  |
| `postmortem` | `Hash` |  |
| `preview_key` | `String` |  |
| `published_at` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Postmortem record (raises on error).
postmortem = client.Postmortem.load({ "incident_id" => "incident_id", "page_id" => "page_id" })
```


### StatusEmbedConfig

Create an instance: `status_embed_config = client.StatusEmbedConfig`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `incident_background_color` | `String` |  |
| `incident_text_color` | `String` |  |
| `maintenance_background_color` | `String` |  |
| `maintenance_text_color` | `String` |  |
| `page_id` | `String` |  |
| `position` | `String` |  |
| `status_embed_config` | `Hash` |  |

#### Example: Load

```ruby
# load returns the bare StatusEmbedConfig record (raises on error).
status_embed_config = client.StatusEmbedConfig.load({ "page_id" => "page_id" })
```


### Subscriber

Create an instance: `subscriber = client.Subscriber`

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
| `component` | `String` |  |
| `component_id` | `Array` |  |
| `created_at` | `String` |  |
| `display_phone_number` | `String` |  |
| `email` | `String` |  |
| `endpoint` | `String` |  |
| `id` | `String` |  |
| `integration_partner` | `Integer` |  |
| `mode` | `String` |  |
| `obfuscated_channel_name` | `String` |  |
| `page_access_user_id` | `String` |  |
| `phone_country` | `String` |  |
| `phone_number` | `String` |  |
| `purge_at` | `String` |  |
| `quarantined_at` | `String` |  |
| `skip_confirmation_notification` | `Boolean` |  |
| `skip_unsubscription_notification` | `Boolean` |  |
| `slack` | `Integer` |  |
| `sms` | `Integer` |  |
| `state` | `String` |  |
| `subscriber` | `Hash` |  |
| `team` | `Integer` |  |
| `type` | `String` |  |
| `webhook` | `Integer` |  |
| `workspace_name` | `String` |  |

#### Example: Load

```ruby
# load returns the bare Subscriber record (raises on error).
subscriber = client.Subscriber.load({ "id" => "subscriber_id", "page_id" => "page_id" })
```

#### Example: List

```ruby
# list returns an Array of Subscriber records (raises on error).
subscribers = client.Subscriber.list
```

#### Example: Create

```ruby
subscriber = client.Subscriber.create({
  "page_id" => "example_page_id", # String
})
```


### User

Create an instance: `user = client.User`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |
| `remove(match)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `String` |  |
| `email` | `String` |  |
| `first_name` | `String` |  |
| `id` | `String` |  |
| `last_name` | `String` |  |
| `organization_id` | `String` |  |
| `updated_at` | `String` |  |
| `user` | `Hash` |  |

#### Example: List

```ruby
# list returns an Array of User records (raises on error).
users = client.User.list
```

#### Example: Create

```ruby
user = client.User.create({
  "organization_id" => "example_organization_id", # String
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

Features are the extension mechanism. A feature is a Ruby class
with hook methods named after pipeline stages (e.g. `PrePoint`,
`PreSpec`). Each method receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as hashes

The Ruby SDK uses plain Ruby hashes throughout rather than typed
objects. This mirrors the dynamic nature of the API and keeps the
SDK flexible — no code generation is needed when the API schema
changes.

Use `Helpers.to_map()` to safely validate that a value is a hash.

### Module structure

```
rb/
├── Statuspage_sdk.rb       -- Main SDK module
├── config.rb                  -- Configuration
├── features.rb                -- Feature factory
├── core/                      -- Core types and context
├── entity/                    -- Entity implementations
├── feature/                   -- Built-in features (Base, Test, Log)
├── utility/                   -- Utility functions and struct library
└── test/                      -- Test suites
```

The main module (`Statuspage_sdk`) exports the SDK class
and test helper. Import entity or utility modules directly only
when needed.

### Entity state

Entity instances are stateful. After a successful `list`, the entity
stores the returned data and match criteria internally.

```ruby
component = client.Component
component.list()

# component.data_get now returns the component data from the last list
# component.match_get returns the last match criteria
```

Call `make` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`direct` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `prepare` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
