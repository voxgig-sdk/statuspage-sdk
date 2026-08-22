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
  # load returns the ENTITY — call data_get for the Component record (raises on error).
  component = client.Component.load({ "page_id" => "example_page_id", "id" => "example_id" })
  puts component
rescue => err
  warn "load failed: #{err}"
end
```

### 4. Create, update, and remove

```ruby
# create returns the ENTITY — call data_get for the created Component record.
created = client.Component.create({ "page_id" => "example_page_id" })

# Update — index the record via data_get (created.data_get["id"]).
client.Component.update({ "id" => created.data_get["id"], "page_id" => "example_page_id", "automation_email" => "example_automation_email" })

# Remove
client.Component.remove({ "id" => created.data_get["id"], "page_id" => "example_page_id" })
```


## Error handling

Entity operations raise on failure, so rescue them:

```ruby
begin
  postmortem = client.Postmortem.load({ "incident_id" => "example", "page_id" => "example" })
rescue => err
  warn "load failed: #{err}"
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

Create a mock client for unit testing — no server required:

```ruby
client = StatuspageSDK.test

# Entity ops return the ENTITY (raises on error);
# call data_get for the mock record.
postmortem = client.Postmortem.load({ "incident_id" => "example", "page_id" => "example" })
puts postmortem
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
| `automation_email` | `String` | Requires a special feature flag to be enabled |
| `component` | `Hash` |  |
| `created_at` | `String` |  |
| `description` | `String` | More detailed description for component |
| `group` | `Boolean` | Is this component a group |
| `group_id` | `String` | Component Group identifier |
| `id` | `String` | Incident identifier |
| `name` | `String` | Display name for component |
| `only_show_if_degraded` | `Boolean` | Requires a special feature flag to be enabled |
| `page_id` | `String` | Page identifier |
| `position` | `Integer` | Order the component will appear on the page |
| `showcase` | `Boolean` | Should this component be showcased |
| `start_date` | `String` | The date this component started being used |
| `status` | `String` | Status of component |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Component record (raises on error).
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
| `component_id` | `String` | Component identifier |
| `incidents` | `Hash` | Related incidents |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the ComponentGroupUptime record (raises on error).
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
| `component_group` | `Hash` |  |
| `components` | `String` |  |
| `created_at` | `String` |  |
| `description` | `String` | Description of the component group. |
| `id` | `String` | Component Group Identifier |
| `name` | `String` |  |
| `page_id` | `String` |  |
| `position` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the GroupComponent record (raises on error).
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
  "component_group" => {}, # Hash
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
| `auto_transition_deliver_notifications_at_end` | `Boolean` | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `Boolean` | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `Boolean` | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `Boolean` | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `Array` | Incident components |
| `created_at` | `String` | The timestamp when the incident was created at. |
| `id` | `String` | Incident Identifier |
| `impact` | `String` | The impact of the incident. |
| `impact_override` | `String` | value to override calculated impact value |
| `incident` | `Hash` |  |
| `incident_updates` | `Array` | The incident updates for incident. |
| `metadata` | `Hash` | Metadata attached to the incident. |
| `monitoring_at` | `String` | The timestamp when incident entered monitoring state. |
| `name` | `String` | Incident Name. |
| `page_id` | `String` | Incident Page Identifier |
| `postmortem_body` | `String` | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `String` | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `Boolean` | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `Boolean` | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `Boolean` | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `Boolean` | The timestamp when the postmortem was published. |
| `reminder_intervals` | `String` | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `String` | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `Boolean` | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `Boolean` | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `String` | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `Boolean` | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `String` | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `String` | The timestamp the incident is scheduled until. |
| `shortlink` | `String` | Incident Shortlink. |
| `status` | `String` | The incident status. |
| `updated_at` | `String` | The timestamp when the incident was updated at. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Incident record (raises on error).
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
  "incident" => {}, # Hash
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
| `body` | `String` | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `Array` | Affected components |
| `group_id` | `String` | Identifier of Template Group this template belongs to |
| `id` | `String` | Incident Template Identifier |
| `name` | `String` | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `Boolean` | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `Boolean` | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `Hash` |  |
| `title` | `String` | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `String` | The status the incident or maintenance should transition to when selecting this template |

#### Example: List

```ruby
# list returns an Array of IncidentTemplate records (raises on error).
incident_templates = client.IncidentTemplate.list
```

#### Example: Create

```ruby
incident_template = client.IncidentTemplate.create({
  "page_id" => "example_page_id", # String
  "template" => {}, # Hash
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
| `affected_components` | `Array` | Affected components associated with the incident update. |
| `body` | `String` | Incident update body. |
| `created_at` | `String` | The timestamp when the incident update was created at. |
| `custom_tweet` | `String` | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `Boolean` | Controls whether to delivery notifications. |
| `display_at` | `String` | Timestamp when incident update is happened. |
| `id` | `String` | Incident Update Identifier. |
| `incident_id` | `String` | Incident Identifier. |
| `incident_update` | `Hash` |  |
| `status` | `String` | The incident status. |
| `tweet_id` | `String` | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `String` | The timestamp when twitter updated at. |
| `updated_at` | `String` | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `Boolean` | Controls whether to create twitter update. |


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
| `data` | `Hash` | Add data points to metrics |
| `decimal_places` | `Integer` |  |
| `display` | `Boolean` | Should the metric be displayed |
| `id` | `String` | Metric identifier |
| `last_fetched_at` | `String` |  |
| `metric` | `Hash` |  |
| `metric_identifier` | `String` | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `String` | Metric Provider identifier |
| `most_recent_data_at` | `String` |  |
| `name` | `String` | Name of metric |
| `reference_name` | `String` |  |
| `suffix` | `String` | Suffix to describe the units on the graph |
| `tooltip_description` | `String` |  |
| `updated_at` | `String` |  |
| `y_axis_hidden` | `Boolean` | Should the values on the y axis be hidden on render |
| `y_axis_max` | `Float` |  |
| `y_axis_min` | `Float` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Metric record (raises on error).
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
  "data" => {}, # Hash
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
| `id` | `String` | Identifier for Metrics Provider |
| `last_revalidated_at` | `String` |  |
| `metric_base_uri` | `String` |  |
| `metrics_provider` | `Hash` |  |
| `page_id` | `Integer` |  |
| `type` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the MetricsProvider record (raises on error).
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
| `allow_email_subscribers` | `Boolean` | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `Boolean` | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `Boolean` | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `Boolean` | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `Boolean` | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `Boolean` | Can your users choose to receive notifications via Webhooks |
| `branding` | `String` | The main template your statuspage will use |
| `city` | `String` |  |
| `country` | `String` |  |
| `created_at` | `String` | Timestamp the record was created |
| `css_blues` | `String` | CSS Color |
| `css_body_background_color` | `String` | CSS Color |
| `css_border_color` | `String` | CSS Color |
| `css_font_color` | `String` | CSS Color |
| `css_graph_color` | `String` | CSS Color |
| `css_greens` | `String` | CSS Color |
| `css_light_font_color` | `String` | CSS Color |
| `css_link_color` | `String` | CSS Color |
| `css_no_data` | `String` | CSS Color |
| `css_oranges` | `String` | CSS Color |
| `css_reds` | `String` | CSS Color |
| `css_yellows` | `String` | CSS Color |
| `domain` | `String` | CNAME alias for your status page |
| `email_logo` | `String` |  |
| `favicon_logo` | `String` |  |
| `headline` | `String` |  |
| `hero_cover` | `String` |  |
| `hidden_from_search` | `Boolean` | Should your page hide itself from search engines |
| `id` | `String` | Page identifier |
| `ip_restrictions` | `String` |  |
| `name` | `String` | Name of your page to be displayed |
| `notifications_email_footer` | `String` | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `String` | Allows you to customize the email address your page notifications come from |
| `page` | `Hash` |  |
| `page_description` | `String` |  |
| `state` | `String` |  |
| `subdomain` | `String` | Subdomain at which to access your status page |
| `support_url` | `String` |  |
| `time_zone` | `String` | Timezone configured for your page |
| `transactional_logo` | `String` |  |
| `twitter_logo` | `String` |  |
| `twitter_username` | `String` |  |
| `updated_at` | `String` | Timestamp the record was last updated |
| `url` | `String` | Website of your page. |
| `viewers_must_be_team_members` | `Boolean` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Page record (raises on error).
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
| `component_ids` | `Array` | List of components codes to set on the page access group |
| `created_at` | `String` |  |
| `external_identifier` | `String` | Associates group with external group. |
| `id` | `String` | Page Access Group Identifier |
| `metric_ids` | `Array` |  |
| `name` | `String` | Name for this Group. |
| `page_access_group` | `Hash` |  |
| `page_access_user_ids` | `Array` |  |
| `page_id` | `String` | Page Identifier. |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PageAccessGroup record (raises on error).
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
| `component_ids` | `Array` | List of component codes to allow access to |
| `created_at` | `String` |  |
| `email` | `String` |  |
| `external_login` | `String` | IDP login user id. |
| `id` | `String` | Page Access User Identifier |
| `metric_ids` | `Array` | List of metrics to add |
| `page_access_group_id` | `String` |  |
| `page_access_group_ids` | `String` |  |
| `page_access_user` | `Hash` |  |
| `page_id` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the PageAccessUser record (raises on error).
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
  "component_ids" => [], # Array
  "metric_ids" => [], # Array
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
| `pages` | `Hash` | Pages accessible by the user. |
| `user_id` | `String` | User identifier |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Permission record (raises on error).
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
| `body` | `String` | Postmortem body |
| `body_draft` | `String` | Body draft |
| `body_draft_updated_at` | `String` |  |
| `body_updated_at` | `String` |  |
| `created_at` | `String` |  |
| `custom_tweet` | `String` | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `Boolean` | Should email subscribers be notified. |
| `notify_twitter` | `Boolean` | Should Twitter followers be notified. |
| `postmortem` | `Hash` |  |
| `preview_key` | `String` | Preview Key |
| `published_at` | `String` |  |
| `updated_at` | `String` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Postmortem record (raises on error).
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
| `incident_background_color` | `String` | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `String` | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `String` | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `String` | Color of status embed iframe text when displaying maintenance |
| `page_id` | `String` | Page identifier |
| `position` | `String` | Corner where status embed iframe will appear on page |
| `status_embed_config` | `Hash` |  |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the StatusEmbedConfig record (raises on error).
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
| `component_ids` | `Array` | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `String` | The components for which the subscriber has elected to receive updates. |
| `created_at` | `String` |  |
| `display_phone_number` | `String` | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `String` | The email address to use to contact the subscriber. |
| `endpoint` | `String` | The URL where a webhook subscriber elects to receive updates. |
| `id` | `String` | Subscriber Identifier |
| `integration_partner` | `Integer` | The number of integration partners found by the query. |
| `mode` | `String` | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `String` | Obfuscated slack channel name |
| `page_access_user_id` | `String` | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `String` | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `String` | The phone number used to contact an SMS subscriber |
| `purge_at` | `String` | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `String` | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `Boolean` | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `Boolean` | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `Integer` | The number of Slack subscribers found by the query. |
| `sms` | `Integer` | The number of Webhook subscribers found by the query. |
| `state` | `String` | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `Hash` |  |
| `subscribers` | `String` | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `Integer` | The number of MS teams subscribers found by the query. |
| `type` | `String` | If this is present, only reactivate subscribers of this type. |
| `webhook` | `Integer` | The number of SMS subscribers found by the query. |
| `workspace_name` | `String` | The workspace name of the slack subscriber. |

#### Example: Load

```ruby
# load returns the ENTITY — call data_get for the Subscriber record (raises on error).
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
  "subscribers" => "example_subscribers", # String
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
| `email` | `String` | Email address for the team member |
| `first_name` | `String` |  |
| `id` | `String` | User identifier |
| `last_name` | `String` |  |
| `organization_id` | `String` | Organization identifier |
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
  "user" => {}, # Hash
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

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally.

```ruby
postmortem = client.Postmortem
postmortem.load({ "incident_id" => "example", "page_id" => "example" })

# postmortem.data_get now returns the postmortem data from the last load
# postmortem.match_get returns the last match criteria
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
