# Statuspage Golang SDK



The Golang SDK for the Statuspage API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Component(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`, `Patch`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Also generated from this model: `go-cli`, `go-mcp`, `lua`, `php`, `py`, `rb`, `ts` — see
> the [top-level README](../README.md).


## Install
```bash
go get github.com/voxgig-sdk/statuspage-sdk/go@latest
```

The Go module proxy resolves the version from the `go/vX.Y.Z` GitHub
release tag — see [Releases](https://github.com/voxgig-sdk/statuspage-sdk/releases) for the available versions.

To vendor from a local checkout instead, clone this repo alongside your
project and add a `replace` directive pointing at the checked-out
`go/` directory:

```bash
go mod edit -replace github.com/voxgig-sdk/statuspage-sdk/go=../statuspage-sdk/go
```


## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.

### Quickstart

A complete program: create a client, then call the entity operations.
Each operation returns `(value, error)` — the value is the data itself
(there is no `{ok, data}` wrapper), so check `err` and use the value
directly.

```go
package main

import (
    "fmt"
    "os"
    sdk "github.com/voxgig-sdk/statuspage-sdk/go"
)

func main() {
    client := sdk.NewStatuspageSDK(map[string]any{
        "apikey": os.Getenv("STATUSPAGE_APIKEY"),
    })

    // List component records — the value is the array of records itself.
    components, err := client.Component(nil).List(nil, nil)
    if err != nil {
        panic(err)
    }
    for _, item := range components.([]any) {
        fmt.Println(item)
    }

    // Load a single component — the value is the loaded record.
    component, err := client.Component(nil).Load(map[string]any{"id": "example_id", "page_id": "example_page_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(component)

    // Create a component.
    created, err := client.Component(nil).Create(map[string]any{"page_id": "example_page_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Update a component.
    updated, err := client.Component(nil).Update(map[string]any{"id": "example_id", "page_id": "example_page_id", "automation_email": "example_automation_email"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(updated)

    // Remove a component.
    removed, err := client.Component(nil).Remove(map[string]any{"id": "example_id", "page_id": "example_page_id"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(removed)
}
```


## Error handling

Every entity operation returns `(value, error)`. Check `err` before
using the value — there is no exception to catch:

```go
postmortem, err := client.Postmortem(nil).Load(map[string]any{"incident_id": "example", "page_id": "example"}, nil)
if err != nil {
    // handle err
    return
}
_ = postmortem
```

`Direct` follows the same `(value, error)` convention:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example_id"},
})
if err != nil {
    // handle err
}
_ = result
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```go
result, err := client.Direct(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "GET",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

if result["ok"] == true {
    fmt.Println(result["status"]) // 200
    fmt.Println(result["data"])   // response body
}
```

### Prepare a request without sending it

```go
fetchdef, err := client.Prepare(map[string]any{
    "path":   "/api/resource/{id}",
    "method": "DELETE",
    "params": map[string]any{"id": "example"},
})
if err != nil {
    panic(err)
}

fmt.Println(fetchdef["url"])
fmt.Println(fetchdef["method"])
fmt.Println(fetchdef["headers"])
```

### Use test mode

Create a mock client for unit testing — no server required:

```go
client := sdk.Test()

postmortem, err := client.Postmortem(nil).Load(
    map[string]any{"incident_id": "example", "page_id": "example"}, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(postmortem) // the returned mock data
```

### Use a custom fetch function

Replace the HTTP transport with your own function:

```go
mockFetch := func(url string, init map[string]any) (map[string]any, error) {
    return map[string]any{
        "status":     200,
        "statusText": "OK",
        "headers":    map[string]any{},
        "json": (func() any)(func() any {
            return map[string]any{"id": "mock01"}
        }),
    }, nil
}

client := sdk.NewStatuspageSDK(map[string]any{
    "base": "http://localhost:8080",
    "system": map[string]any{
        "fetch": (func(string, map[string]any) (map[string]any, error))(mockFetch),
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
cd go && go test ./test/...
```


## Reference

### NewStatuspageSDK

```go
func NewStatuspageSDK(options map[string]any) *StatuspageSDK
```

Creates a new SDK client.

| Option | Type | Description |
| --- | --- | --- |
| `"apikey"` | `string` | API key for authentication. |
| `"base"` | `string` | Base URL of the API server. |
| `"prefix"` | `string` | URL path prefix prepended to all requests. |
| `"suffix"` | `string` | URL path suffix appended to all requests. |
| `"feature"` | `map[string]any` | Feature activation flags. |
| `"extend"` | `[]any` | Additional Feature instances to load. |
| `"system"` | `map[string]any` | System overrides (e.g. custom `"fetch"` function). |

### TestSDK

```go
func TestSDK(testopts map[string]any, sdkopts map[string]any) *StatuspageSDK
```

Creates a test-mode client with mock transport. Both arguments may be `nil`.

### StatuspageSDK methods

| Method | Signature | Description |
| --- | --- | --- |
| `OptionsMap` | `() map[string]any` | Deep copy of current SDK options. |
| `GetUtility` | `() *Utility` | Copy of the SDK utility object. |
| `Prepare` | `(fetchargs map[string]any) (map[string]any, error)` | Build an HTTP request definition without sending. |
| `Direct` | `(fetchargs map[string]any) (map[string]any, error)` | Build and send an HTTP request. |
| `Component` | `(data map[string]any) StatuspageEntity` | Create a Component entity instance. |
| `ComponentGroupUptime` | `(data map[string]any) StatuspageEntity` | Create a ComponentGroupUptime entity instance. |
| `GroupComponent` | `(data map[string]any) StatuspageEntity` | Create a GroupComponent entity instance. |
| `Incident` | `(data map[string]any) StatuspageEntity` | Create an Incident entity instance. |
| `IncidentPostmortem` | `(data map[string]any) StatuspageEntity` | Create an IncidentPostmortem entity instance. |
| `IncidentSubscriber` | `(data map[string]any) StatuspageEntity` | Create an IncidentSubscriber entity instance. |
| `IncidentTemplate` | `(data map[string]any) StatuspageEntity` | Create an IncidentTemplate entity instance. |
| `IncidentUpdate` | `(data map[string]any) StatuspageEntity` | Create an IncidentUpdate entity instance. |
| `Metric` | `(data map[string]any) StatuspageEntity` | Create a Metric entity instance. |
| `MetricsProvider` | `(data map[string]any) StatuspageEntity` | Create a MetricsProvider entity instance. |
| `Page` | `(data map[string]any) StatuspageEntity` | Create a Page entity instance. |
| `PageAccessGroup` | `(data map[string]any) StatuspageEntity` | Create a PageAccessGroup entity instance. |
| `PageAccessUser` | `(data map[string]any) StatuspageEntity` | Create a PageAccessUser entity instance. |
| `Permission` | `(data map[string]any) StatuspageEntity` | Create a Permission entity instance. |
| `Postmortem` | `(data map[string]any) StatuspageEntity` | Create a Postmortem entity instance. |
| `StatusEmbedConfig` | `(data map[string]any) StatuspageEntity` | Create a StatusEmbedConfig entity instance. |
| `Subscriber` | `(data map[string]any) StatuspageEntity` | Create a Subscriber entity instance. |
| `User` | `(data map[string]any) StatuspageEntity` | Create an User entity instance. |

### Entity interface (StatuspageEntity)

All entities implement the `StatuspageEntity` interface.

| Method | Signature | Description |
| --- | --- | --- |
| `Load` | `(reqmatch, ctrl map[string]any) (any, error)` | Load a single entity by match criteria. |
| `List` | `(reqmatch, ctrl map[string]any) (any, error)` | List entities matching the criteria. |
| `Create` | `(reqdata, ctrl map[string]any) (any, error)` | Create a new entity. |
| `Update` | `(reqdata, ctrl map[string]any) (any, error)` | Update an existing entity. |
| `Remove` | `(reqmatch, ctrl map[string]any) (any, error)` | Remove an entity. |
| `Data` | `(args ...any) any` | Get or set entity data. |
| `Match` | `(args ...any) any` | Get or set entity match criteria. |
| `Make` | `() Entity` | Create a new instance with the same options. |
| `GetName` | `() string` | Return the entity name. |

### Result shape

Entity operations return `(value, error)`. The `value` is the
operation's data **directly** — there is no wrapper:

| Operation | `value` |
| --- | --- |
| `Load` / `Create` / `Update` / `Remove` | the entity record (`map[string]any`) |
| `List` | a `[]any` of entity records |

Check `err` first, then use the value directly (or the typed
`...Typed` variants, which return the entity's model struct and a typed
slice):

    component, err := client.Component(nil).List(map[string]any{/* fields */}, nil)
    if err != nil { /* handle */ }
    // component is the returned record

Only `Direct()` returns a response envelope — a `map[string]any` with
`"ok"`, `"status"`, `"headers"`, and `"data"` keys.

### Entities

#### Component

| Field | Description |
| --- | --- |
| `"automation_email"` | Requires a special feature flag to be enabled |
| `"component"` |  |
| `"created_at"` |  |
| `"description"` | More detailed description for component |
| `"group"` | Is this component a group |
| `"group_id"` | Component Group identifier |
| `"id"` | Incident identifier |
| `"name"` | Display name for component |
| `"only_show_if_degraded"` | Requires a special feature flag to be enabled |
| `"page_id"` | Page identifier |
| `"position"` | Order the component will appear on the page |
| `"showcase"` | Should this component be showcased |
| `"start_date"` | The date this component started being used |
| `"status"` | Status of component |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/components/{component_id}/page_access_groups`

#### ComponentGroupUptime

| Field | Description |
| --- | --- |
| `"component_id"` | Component identifier |
| `"incidents"` | Related incidents |

Operations: Load.

API path: `/pages/{page_id}/component-groups/{id}/uptime`

#### GroupComponent

| Field | Description |
| --- | --- |
| `"component_group"` |  |
| `"components"` |  |
| `"created_at"` |  |
| `"description"` | Description of the component group. |
| `"id"` | Component Group Identifier |
| `"name"` |  |
| `"page_id"` |  |
| `"position"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/component-groups`

#### Incident

| Field | Description |
| --- | --- |
| `"auto_transition_deliver_notifications_at_end"` | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `"auto_transition_deliver_notifications_at_start"` | Controls whether send notification when scheduled maintenances auto transition to started. |
| `"auto_transition_to_maintenance_state"` | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `"auto_transition_to_operational_state"` | Controls whether change components status to operational once scheduled maintenance completes. |
| `"components"` | Incident components |
| `"created_at"` | The timestamp when the incident was created at. |
| `"id"` | Incident Identifier |
| `"impact"` | The impact of the incident. |
| `"impact_override"` | value to override calculated impact value |
| `"incident"` |  |
| `"incident_updates"` | The incident updates for incident. |
| `"metadata"` | Metadata attached to the incident. |
| `"monitoring_at"` | The timestamp when incident entered monitoring state. |
| `"name"` | Incident Name. |
| `"page_id"` | Incident Page Identifier |
| `"postmortem_body"` | Body of the Postmortem. |
| `"postmortem_body_last_updated_at"` | The timestamp when the incident postmortem body was last updated at. |
| `"postmortem_ignored"` | Controls whether the incident will have postmortem. |
| `"postmortem_notified_subscribers"` | Indicates whether subscribers are already notificed about postmortem. |
| `"postmortem_notified_twitter"` | Controls whether to decide if notify postmortem on twitter. |
| `"postmortem_published_at"` | The timestamp when the postmortem was published. |
| `"reminder_intervals"` | Custom reminder intervals for unresolved/open incidents. |
| `"resolved_at"` | The timestamp when incident was resolved. |
| `"scheduled_auto_completed"` | Controls whether the incident is scheduled to automatically change to complete. |
| `"scheduled_auto_in_progress"` | Controls whether the incident is scheduled to automatically change to in progress. |
| `"scheduled_for"` | The timestamp the incident is scheduled for. |
| `"scheduled_remind_prior"` | Controls whether to remind subscribers prior to scheduled incidents. |
| `"scheduled_reminded_at"` | The timestamp when the scheduled incident reminder was sent at. |
| `"scheduled_until"` | The timestamp the incident is scheduled until. |
| `"shortlink"` | Incident Shortlink. |
| `"status"` | The incident status. |
| `"updated_at"` | The timestamp when the incident was updated at. |

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
| `"body"` | Body of the incident or maintenance update to be applied when selecting this template |
| `"components"` | Affected components |
| `"group_id"` | Identifier of Template Group this template belongs to |
| `"id"` | Incident Template Identifier |
| `"name"` | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `"should_send_notifications"` | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `"should_tweet"` | Whether the "tweet update" checkbox should be selected when selecting this template |
| `"template"` |  |
| `"title"` | Title to be applied to the incident or maintenance when selecting this template |
| `"update_status"` | The status the incident or maintenance should transition to when selecting this template |

Operations: Create, List.

API path: `/pages/{page_id}/incident_templates`

#### IncidentUpdate

| Field | Description |
| --- | --- |
| `"affected_components"` | Affected components associated with the incident update. |
| `"body"` | Incident update body. |
| `"created_at"` | The timestamp when the incident update was created at. |
| `"custom_tweet"` | An optional customized tweet message for incident postmortem. |
| `"deliver_notifications"` | Controls whether to delivery notifications. |
| `"display_at"` | Timestamp when incident update is happened. |
| `"id"` | Incident Update Identifier. |
| `"incident_id"` | Incident Identifier. |
| `"incident_update"` |  |
| `"status"` | The incident status. |
| `"tweet_id"` | Tweet identifier associated to this incident update. |
| `"twitter_updated_at"` | The timestamp when twitter updated at. |
| `"updated_at"` | The timestamp when the incident update is updated. |
| `"wants_twitter_update"` | Controls whether to create twitter update. |

Operations: Patch, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}`

#### Metric

| Field | Description |
| --- | --- |
| `"backfill_percentage"` |  |
| `"backfilled"` |  |
| `"created_at"` |  |
| `"data"` | Add data points to metrics |
| `"decimal_places"` |  |
| `"display"` | Should the metric be displayed |
| `"id"` | Metric identifier |
| `"last_fetched_at"` |  |
| `"metric"` |  |
| `"metric_identifier"` | Metric Display identifier used to look up the metric data from the provider |
| `"metrics_provider_id"` | Metric Provider identifier |
| `"most_recent_data_at"` |  |
| `"name"` | Name of metric |
| `"reference_name"` |  |
| `"suffix"` | Suffix to describe the units on the graph |
| `"tooltip_description"` |  |
| `"updated_at"` |  |
| `"y_axis_hidden"` | Should the values on the y axis be hidden on render |
| `"y_axis_max"` |  |
| `"y_axis_min"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/metrics/{metric_id}/data`

#### MetricsProvider

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"disabled"` |  |
| `"id"` | Identifier for Metrics Provider |
| `"last_revalidated_at"` |  |
| `"metric_base_uri"` |  |
| `"metrics_provider"` |  |
| `"page_id"` |  |
| `"type"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/metrics_providers`

#### Page

| Field | Description |
| --- | --- |
| `"activity_score"` |  |
| `"allow_email_subscribers"` | Can your users choose to receive notifications via email |
| `"allow_incident_subscribers"` | Can your users subscribe to notifications for a single incident |
| `"allow_page_subscribers"` | Can your users subscribe to all notifications on the page |
| `"allow_rss_atom_feeds"` | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `"allow_sms_subscribers"` | Can your users choose to receive notifications via SMS |
| `"allow_webhook_subscribers"` | Can your users choose to receive notifications via Webhooks |
| `"branding"` | The main template your statuspage will use |
| `"city"` |  |
| `"country"` |  |
| `"created_at"` | Timestamp the record was created |
| `"css_blues"` | CSS Color |
| `"css_body_background_color"` | CSS Color |
| `"css_border_color"` | CSS Color |
| `"css_font_color"` | CSS Color |
| `"css_graph_color"` | CSS Color |
| `"css_greens"` | CSS Color |
| `"css_light_font_color"` | CSS Color |
| `"css_link_color"` | CSS Color |
| `"css_no_data"` | CSS Color |
| `"css_oranges"` | CSS Color |
| `"css_reds"` | CSS Color |
| `"css_yellows"` | CSS Color |
| `"domain"` | CNAME alias for your status page |
| `"email_logo"` |  |
| `"favicon_logo"` |  |
| `"headline"` |  |
| `"hero_cover"` |  |
| `"hidden_from_search"` | Should your page hide itself from search engines |
| `"id"` | Page identifier |
| `"ip_restrictions"` |  |
| `"name"` | Name of your page to be displayed |
| `"notifications_email_footer"` | Allows you to customize the footer appearing on your notification emails. |
| `"notifications_from_email"` | Allows you to customize the email address your page notifications come from |
| `"page"` |  |
| `"page_description"` |  |
| `"state"` |  |
| `"subdomain"` | Subdomain at which to access your status page |
| `"support_url"` |  |
| `"time_zone"` | Timezone configured for your page |
| `"transactional_logo"` |  |
| `"twitter_logo"` |  |
| `"twitter_username"` |  |
| `"updated_at"` | Timestamp the record was last updated |
| `"url"` | Website of your page. |
| `"viewers_must_be_team_members"` |  |

Operations: List, Load, Patch, Update.

API path: `/pages`

#### PageAccessGroup

| Field | Description |
| --- | --- |
| `"component_ids"` | List of components codes to set on the page access group |
| `"created_at"` |  |
| `"external_identifier"` | Associates group with external group. |
| `"id"` | Page Access Group Identifier |
| `"metric_ids"` |  |
| `"name"` | Name for this Group. |
| `"page_access_group"` |  |
| `"page_access_user_ids"` |  |
| `"page_id"` | Page Identifier. |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_groups/{page_access_group_id}/components`

#### PageAccessUser

| Field | Description |
| --- | --- |
| `"component_ids"` | List of component codes to allow access to |
| `"created_at"` |  |
| `"email"` |  |
| `"external_login"` | IDP login user id. |
| `"id"` | Page Access User Identifier |
| `"metric_ids"` | List of metrics to add |
| `"page_access_group_id"` |  |
| `"page_access_group_ids"` |  |
| `"page_access_user"` |  |
| `"page_id"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_users/{page_access_user_id}/components`

#### Permission

| Field | Description |
| --- | --- |
| `"pages"` | Pages accessible by the user. |
| `"user_id"` | User identifier |

Operations: Load, Update.

API path: `/organizations/{organization_id}/permissions/{user_id}`

#### Postmortem

| Field | Description |
| --- | --- |
| `"body"` | Postmortem body |
| `"body_draft"` | Body draft |
| `"body_draft_updated_at"` |  |
| `"body_updated_at"` |  |
| `"created_at"` |  |
| `"custom_tweet"` | Custom tweet for Incident Postmortem |
| `"notify_subscribers"` | Should email subscribers be notified. |
| `"notify_twitter"` | Should Twitter followers be notified. |
| `"postmortem"` |  |
| `"preview_key"` | Preview Key |
| `"published_at"` |  |
| `"updated_at"` |  |

Operations: Load, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/postmortem`

#### StatusEmbedConfig

| Field | Description |
| --- | --- |
| `"incident_background_color"` | Color of status embed iframe background when displaying incident |
| `"incident_text_color"` | Color of status embed iframe text when displaying incident |
| `"maintenance_background_color"` | Color of status embed iframe background when displaying maintenance |
| `"maintenance_text_color"` | Color of status embed iframe text when displaying maintenance |
| `"page_id"` | Page identifier |
| `"position"` | Corner where status embed iframe will appear on page |
| `"status_embed_config"` |  |

Operations: Load, Patch, Update.

API path: `/pages/{page_id}/status_embed_config`

#### Subscriber

| Field | Description |
| --- | --- |
| `"component_ids"` | A list of component ids for which the subscriber should recieve updates for. |
| `"components"` | The components for which the subscriber has elected to receive updates. |
| `"created_at"` |  |
| `"display_phone_number"` | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `"email"` | The email address to use to contact the subscriber. |
| `"endpoint"` | The URL where a webhook subscriber elects to receive updates. |
| `"id"` | Subscriber Identifier |
| `"integration_partner"` | The number of integration partners found by the query. |
| `"mode"` | The communication mode of the subscriber. |
| `"obfuscated_channel_name"` | Obfuscated slack channel name |
| `"page_access_user_id"` | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `"phone_country"` | The two-character country code representing the country of which the phone_number is a part. |
| `"phone_number"` | The phone number used to contact an SMS subscriber |
| `"purge_at"` | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `"quarantined_at"` | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `"skip_confirmation_notification"` | If this is true, do not notify the user with changes to their subscription. |
| `"skip_unsubscription_notification"` | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `"slack"` | The number of Slack subscribers found by the query. |
| `"sms"` | The number of Webhook subscribers found by the query. |
| `"state"` | If this is present, only unsubscribe subscribers in this state. |
| `"subscriber"` |  |
| `"subscribers"` | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `"teams"` | The number of MS teams subscribers found by the query. |
| `"type"` | If this is present, only reactivate subscribers of this type. |
| `"webhook"` | The number of SMS subscribers found by the query. |
| `"workspace_name"` | The workspace name of the slack subscriber. |

Operations: Create, List, Load, Remove, Update.

API path: `/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation`

#### User

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"email"` | Email address for the team member |
| `"first_name"` |  |
| `"id"` | User identifier |
| `"last_name"` |  |
| `"organization_id"` | Organization identifier |
| `"updated_at"` |  |
| `"user"` |  |

Operations: Create, List, Remove.

API path: `/organizations/{organization_id}/users`



## Entities


### Component

Create an instance: `component := client.Component(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `automation_email` | `string` | Requires a special feature flag to be enabled |
| `component` | `map[string]any` |  |
| `created_at` | `string` |  |
| `description` | `string` | More detailed description for component |
| `group` | `bool` | Is this component a group |
| `group_id` | `string` | Component Group identifier |
| `id` | `string` | Incident identifier |
| `name` | `string` | Display name for component |
| `only_show_if_degraded` | `bool` | Requires a special feature flag to be enabled |
| `page_id` | `string` | Page identifier |
| `position` | `int` | Order the component will appear on the page |
| `showcase` | `bool` | Should this component be showcased |
| `start_date` | `string` | The date this component started being used |
| `status` | `string` | Status of component |
| `updated_at` | `string` |  |

#### Example: Load

```go
component, err := client.Component(nil).Load(map[string]any{"id": "component_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(component) // the loaded record
```

#### Example: List

```go
components, err := client.Component(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(components) // the array of records
```

#### Example: Create

```go
result, err := client.Component(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### ComponentGroupUptime

Create an instance: `componentGroupUptime := client.ComponentGroupUptime(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `string` | Component identifier |
| `incidents` | `map[string]any` | Related incidents |

#### Example: Load

```go
componentGroupUptime, err := client.ComponentGroupUptime(nil).Load(map[string]any{"id": "component_group_uptime_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(componentGroupUptime) // the loaded record
```


### GroupComponent

Create an instance: `groupComponent := client.GroupComponent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_group` | `map[string]any` |  |
| `components` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` | Description of the component group. |
| `id` | `string` | Component Group Identifier |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
groupComponent, err := client.GroupComponent(nil).Load(map[string]any{"id": "group_component_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(groupComponent) // the loaded record
```

#### Example: List

```go
groupComponents, err := client.GroupComponent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(groupComponents) // the array of records
```

#### Example: Create

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


### Incident

Create an instance: `incident := client.Incident(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `bool` | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `bool` | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `bool` | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `bool` | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `[]any` | Incident components |
| `created_at` | `string` | The timestamp when the incident was created at. |
| `id` | `string` | Incident Identifier |
| `impact` | `string` | The impact of the incident. |
| `impact_override` | `string` | value to override calculated impact value |
| `incident` | `map[string]any` |  |
| `incident_updates` | `[]any` | The incident updates for incident. |
| `metadata` | `map[string]any` | Metadata attached to the incident. |
| `monitoring_at` | `string` | The timestamp when incident entered monitoring state. |
| `name` | `string` | Incident Name. |
| `page_id` | `string` | Incident Page Identifier |
| `postmortem_body` | `string` | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `string` | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `bool` | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `bool` | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `bool` | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `bool` | The timestamp when the postmortem was published. |
| `reminder_intervals` | `string` | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `string` | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `bool` | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `bool` | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `string` | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `bool` | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `string` | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `string` | The timestamp the incident is scheduled until. |
| `shortlink` | `string` | Incident Shortlink. |
| `status` | `string` | The incident status. |
| `updated_at` | `string` | The timestamp when the incident was updated at. |

#### Example: Load

```go
incident, err := client.Incident(nil).Load(map[string]any{"id": "incident_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(incident) // the loaded record
```

#### Example: List

```go
incidents, err := client.Incident(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(incidents) // the array of records
```

#### Example: Create

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


### IncidentPostmortem

Create an instance: `incidentPostmortem := client.IncidentPostmortem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `incidentSubscriber := client.IncidentSubscriber(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Example: Create

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


### IncidentTemplate

Create an instance: `incidentTemplate := client.IncidentTemplate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `[]any` | Affected components |
| `group_id` | `string` | Identifier of Template Group this template belongs to |
| `id` | `string` | Incident Template Identifier |
| `name` | `string` | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `bool` | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `bool` | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `map[string]any` |  |
| `title` | `string` | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `string` | The status the incident or maintenance should transition to when selecting this template |

#### Example: List

```go
incidentTemplates, err := client.IncidentTemplate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(incidentTemplates) // the array of records
```

#### Example: Create

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


### IncidentUpdate

Create an instance: `incidentUpdate := client.IncidentUpdate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_components` | `[]any` | Affected components associated with the incident update. |
| `body` | `string` | Incident update body. |
| `created_at` | `string` | The timestamp when the incident update was created at. |
| `custom_tweet` | `string` | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `bool` | Controls whether to delivery notifications. |
| `display_at` | `string` | Timestamp when incident update is happened. |
| `id` | `string` | Incident Update Identifier. |
| `incident_id` | `string` | Incident Identifier. |
| `incident_update` | `map[string]any` |  |
| `status` | `string` | The incident status. |
| `tweet_id` | `string` | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `string` | The timestamp when twitter updated at. |
| `updated_at` | `string` | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `bool` | Controls whether to create twitter update. |


### Metric

Create an instance: `metric := client.Metric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `backfill_percentage` | `int` |  |
| `backfilled` | `bool` |  |
| `created_at` | `string` |  |
| `data` | `map[string]any` | Add data points to metrics |
| `decimal_places` | `int` |  |
| `display` | `bool` | Should the metric be displayed |
| `id` | `string` | Metric identifier |
| `last_fetched_at` | `string` |  |
| `metric` | `map[string]any` |  |
| `metric_identifier` | `string` | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `string` | Metric Provider identifier |
| `most_recent_data_at` | `string` |  |
| `name` | `string` | Name of metric |
| `reference_name` | `string` |  |
| `suffix` | `string` | Suffix to describe the units on the graph |
| `tooltip_description` | `string` |  |
| `updated_at` | `string` |  |
| `y_axis_hidden` | `bool` | Should the values on the y axis be hidden on render |
| `y_axis_max` | `float64` |  |
| `y_axis_min` | `float64` |  |

#### Example: Load

```go
metric, err := client.Metric(nil).Load(map[string]any{"id": "metric_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(metric) // the loaded record
```

#### Example: List

```go
metrics, err := client.Metric(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(metrics) // the array of records
```

#### Example: Create

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


### MetricsProvider

Create an instance: `metricsProvider := client.MetricsProvider(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `disabled` | `bool` |  |
| `id` | `string` | Identifier for Metrics Provider |
| `last_revalidated_at` | `string` |  |
| `metric_base_uri` | `string` |  |
| `metrics_provider` | `map[string]any` |  |
| `page_id` | `int` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
metricsProvider, err := client.MetricsProvider(nil).Load(map[string]any{"id": "metrics_provider_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(metricsProvider) // the loaded record
```

#### Example: List

```go
metricsProviders, err := client.MetricsProvider(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(metricsProviders) // the array of records
```

#### Example: Create

```go
result, err := client.MetricsProvider(nil).Create(map[string]any{
    "page_id": "example_page_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### Page

Create an instance: `page := client.Page(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `activity_score` | `float64` |  |
| `allow_email_subscribers` | `bool` | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `bool` | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `bool` | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `bool` | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `bool` | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `bool` | Can your users choose to receive notifications via Webhooks |
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
| `hidden_from_search` | `bool` | Should your page hide itself from search engines |
| `id` | `string` | Page identifier |
| `ip_restrictions` | `string` |  |
| `name` | `string` | Name of your page to be displayed |
| `notifications_email_footer` | `string` | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `string` | Allows you to customize the email address your page notifications come from |
| `page` | `map[string]any` |  |
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
| `viewers_must_be_team_members` | `bool` |  |

#### Example: Load

```go
page, err := client.Page(nil).Load(map[string]any{"id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(page) // the loaded record
```

#### Example: List

```go
pages, err := client.Page(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pages) // the array of records
```


### PageAccessGroup

Create an instance: `pageAccessGroup := client.PageAccessGroup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_ids` | `[]any` | List of components codes to set on the page access group |
| `created_at` | `string` |  |
| `external_identifier` | `string` | Associates group with external group. |
| `id` | `string` | Page Access Group Identifier |
| `metric_ids` | `[]any` |  |
| `name` | `string` | Name for this Group. |
| `page_access_group` | `map[string]any` |  |
| `page_access_user_ids` | `[]any` |  |
| `page_id` | `string` | Page Identifier. |
| `updated_at` | `string` |  |

#### Example: Load

```go
pageAccessGroup, err := client.PageAccessGroup(nil).Load(map[string]any{"id": "page_access_group_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(pageAccessGroup) // the loaded record
```

#### Example: List

```go
pageAccessGroups, err := client.PageAccessGroup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pageAccessGroups) // the array of records
```

#### Example: Create

```go
result, err := client.PageAccessGroup(nil).Create(map[string]any{
    "id": "example_id",
}, nil)
if err != nil {
    panic(err)
}
fmt.Println(result)
```


### PageAccessUser

Create an instance: `pageAccessUser := client.PageAccessUser(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_ids` | `[]any` | List of component codes to allow access to |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `external_login` | `string` | IDP login user id. |
| `id` | `string` | Page Access User Identifier |
| `metric_ids` | `[]any` | List of metrics to add |
| `page_access_group_id` | `string` |  |
| `page_access_group_ids` | `string` |  |
| `page_access_user` | `map[string]any` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
pageAccessUser, err := client.PageAccessUser(nil).Load(map[string]any{"id": "page_access_user_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(pageAccessUser) // the loaded record
```

#### Example: List

```go
pageAccessUsers, err := client.PageAccessUser(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(pageAccessUsers) // the array of records
```

#### Example: Create

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


### Permission

Create an instance: `permission := client.Permission(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `pages` | `map[string]any` | Pages accessible by the user. |
| `user_id` | `string` | User identifier |

#### Example: Load

```go
permission, err := client.Permission(nil).Load(map[string]any{"id": "permission_id", "organization_id": "organization_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(permission) // the loaded record
```


### Postmortem

Create an instance: `postmortem := client.Postmortem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` | Postmortem body |
| `body_draft` | `string` | Body draft |
| `body_draft_updated_at` | `string` |  |
| `body_updated_at` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `bool` | Should email subscribers be notified. |
| `notify_twitter` | `bool` | Should Twitter followers be notified. |
| `postmortem` | `map[string]any` |  |
| `preview_key` | `string` | Preview Key |
| `published_at` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
postmortem, err := client.Postmortem(nil).Load(map[string]any{"incident_id": "incident_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(postmortem) // the loaded record
```


### StatusEmbedConfig

Create an instance: `statusEmbedConfig := client.StatusEmbedConfig(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `incident_background_color` | `string` | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `string` | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `string` | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `string` | Color of status embed iframe text when displaying maintenance |
| `page_id` | `string` | Page identifier |
| `position` | `string` | Corner where status embed iframe will appear on page |
| `status_embed_config` | `map[string]any` |  |

#### Example: Load

```go
statusEmbedConfig, err := client.StatusEmbedConfig(nil).Load(map[string]any{"page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(statusEmbedConfig) // the loaded record
```


### Subscriber

Create an instance: `subscriber := client.Subscriber(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Update(data, ctrl)` | Update an existing entity. |
| `Remove(match, ctrl)` | Remove the matching entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_ids` | `[]any` | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `string` | The components for which the subscriber has elected to receive updates. |
| `created_at` | `string` |  |
| `display_phone_number` | `string` | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `string` | The email address to use to contact the subscriber. |
| `endpoint` | `string` | The URL where a webhook subscriber elects to receive updates. |
| `id` | `string` | Subscriber Identifier |
| `integration_partner` | `int` | The number of integration partners found by the query. |
| `mode` | `string` | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `string` | Obfuscated slack channel name |
| `page_access_user_id` | `string` | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `string` | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `string` | The phone number used to contact an SMS subscriber |
| `purge_at` | `string` | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `string` | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `bool` | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `bool` | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `int` | The number of Slack subscribers found by the query. |
| `sms` | `int` | The number of Webhook subscribers found by the query. |
| `state` | `string` | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `map[string]any` |  |
| `subscribers` | `string` | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `int` | The number of MS teams subscribers found by the query. |
| `type` | `string` | If this is present, only reactivate subscribers of this type. |
| `webhook` | `int` | The number of SMS subscribers found by the query. |
| `workspace_name` | `string` | The workspace name of the slack subscriber. |

#### Example: Load

```go
subscriber, err := client.Subscriber(nil).Load(map[string]any{"id": "subscriber_id", "page_id": "page_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(subscriber) // the loaded record
```

#### Example: List

```go
subscribers, err := client.Subscriber(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(subscribers) // the array of records
```

#### Example: Create

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


### User

Create an instance: `user := client.User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `Remove(match, ctrl)` | Remove the matching entity. |

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
| `user` | `map[string]any` |  |

#### Example: List

```go
users, err := client.User(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(users) // the array of records
```

#### Example: Create

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

Features are the extension mechanism. A feature implements the
`Feature` interface and provides hooks — functions keyed by pipeline
stage names.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Data as maps

The Go SDK uses `map[string]any` throughout rather than typed structs.
This mirrors the dynamic nature of the API and keeps the SDK
flexible — no code generation is needed when the API schema changes.

Use `core.ToMapAny()` to safely cast results and nested data.

### Package structure

```
github.com/voxgig-sdk/statuspage-sdk/go/
├── statuspage.go        # Root package — type aliases and constructors
├── core/               # SDK core — client, types, pipeline
├── entity/             # Entity implementations
├── feature/            # Built-in features (Base, Test, Log)
├── utility/            # Utility functions and struct library
└── test/               # Test suites
```

The root package (`github.com/voxgig-sdk/statuspage-sdk/go`) re-exports everything needed
for normal use. Import sub-packages only when you need specific types
like `core.ToMapAny`.

### Entity state

Entity instances are stateful. After a successful `Load`, the entity
stores the returned data and match criteria internally.

```go
postmortem := client.Postmortem(nil)
postmortem.Load(map[string]any{"incident_id": "example", "page_id": "example"}, nil)

// postmortem.Data() now returns the postmortem data from the last load
// postmortem.Match() returns the last match criteria
```

Call `Make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

`Direct()` gives full control over the HTTP request. Use it for
non-standard endpoints, bulk operations, or any path not modelled as
an entity. `Prepare()` builds the request without sending it — useful
for debugging or custom transport.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
