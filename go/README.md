# Statuspage Golang SDK



The Golang SDK for the Statuspage API — an entity-oriented client using standard Go conventions. No generics required; data flows as `map[string]any`.

It exposes the API as capitalised, semantic **Entities** — e.g. `client.Component(nil)` — each with the same small set of operations (`List`, `Load`, `Create`, `Update`, `Remove`, `Patch`) instead of raw URL paths and query strings. You call meaning, not endpoints, which keeps the cognitive load low.

> Other languages, the CLI, and MCP server live alongside this one — see
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
    component, err := client.Component(nil).Load(map[string]any{"id": "example"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(component)

    // Create a component.
    created, err := client.Component(nil).Create(map[string]any{"page_id": "example"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(created)

    // Update a component.
    updated, err := client.Component(nil).Update(map[string]any{"id": "example", "page_id": "example"}, nil)
    if err != nil {
        panic(err)
    }
    fmt.Println(updated)

    // Remove a component.
    removed, err := client.Component(nil).Remove(map[string]any{"id": "example"}, nil)
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
components, err := client.Component(nil).List(nil, nil)
if err != nil {
    // handle err
    return
}
_ = components
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

component, err := client.Component(nil).List(
    nil, nil,
)
if err != nil {
    panic(err)
}
fmt.Println(component) // the returned mock data
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
| `"automation_email"` |  |
| `"component"` |  |
| `"created_at"` |  |
| `"description"` |  |
| `"group"` |  |
| `"group_id"` |  |
| `"id"` |  |
| `"major_outage"` |  |
| `"name"` |  |
| `"only_show_if_degraded"` |  |
| `"page_id"` |  |
| `"partial_outage"` |  |
| `"position"` |  |
| `"range_end"` |  |
| `"range_start"` |  |
| `"related_event"` |  |
| `"showcase"` |  |
| `"start_date"` |  |
| `"status"` |  |
| `"updated_at"` |  |
| `"uptime_percentage"` |  |
| `"warning"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/components/{component_id}/page_access_groups`

#### ComponentGroupUptime

| Field | Description |
| --- | --- |
| `"id"` |  |
| `"major_outage"` |  |
| `"name"` |  |
| `"partial_outage"` |  |
| `"range_end"` |  |
| `"range_start"` |  |
| `"related_event"` |  |
| `"uptime_percentage"` |  |
| `"warning"` |  |

Operations: Load.

API path: `/pages/{page_id}/component-groups/{id}/uptime`

#### GroupComponent

| Field | Description |
| --- | --- |
| `"component"` |  |
| `"component_group"` |  |
| `"created_at"` |  |
| `"description"` |  |
| `"id"` |  |
| `"name"` |  |
| `"page_id"` |  |
| `"position"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/component-groups`

#### Incident

| Field | Description |
| --- | --- |
| `"auto_transition_deliver_notifications_at_end"` |  |
| `"auto_transition_deliver_notifications_at_start"` |  |
| `"auto_transition_to_maintenance_state"` |  |
| `"auto_transition_to_operational_state"` |  |
| `"component"` |  |
| `"created_at"` |  |
| `"id"` |  |
| `"impact"` |  |
| `"impact_override"` |  |
| `"incident"` |  |
| `"incident_impact"` |  |
| `"incident_update"` |  |
| `"metadata"` |  |
| `"monitoring_at"` |  |
| `"name"` |  |
| `"page_id"` |  |
| `"postmortem_body"` |  |
| `"postmortem_body_last_updated_at"` |  |
| `"postmortem_ignored"` |  |
| `"postmortem_notified_subscriber"` |  |
| `"postmortem_notified_twitter"` |  |
| `"postmortem_published_at"` |  |
| `"reminder_interval"` |  |
| `"resolved_at"` |  |
| `"scheduled_auto_completed"` |  |
| `"scheduled_auto_in_progress"` |  |
| `"scheduled_for"` |  |
| `"scheduled_remind_prior"` |  |
| `"scheduled_reminded_at"` |  |
| `"scheduled_until"` |  |
| `"shortlink"` |  |
| `"status"` |  |
| `"updated_at"` |  |

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
| `"body"` |  |
| `"component"` |  |
| `"group_id"` |  |
| `"id"` |  |
| `"name"` |  |
| `"should_send_notification"` |  |
| `"should_tweet"` |  |
| `"template"` |  |
| `"title"` |  |
| `"update_status"` |  |

Operations: Create, List.

API path: `/pages/{page_id}/incident_templates`

#### IncidentUpdate

| Field | Description |
| --- | --- |
| `"affected_component"` |  |
| `"body"` |  |
| `"created_at"` |  |
| `"custom_tweet"` |  |
| `"deliver_notification"` |  |
| `"display_at"` |  |
| `"id"` |  |
| `"incident_id"` |  |
| `"incident_update"` |  |
| `"status"` |  |
| `"tweet_id"` |  |
| `"twitter_updated_at"` |  |
| `"updated_at"` |  |
| `"wants_twitter_update"` |  |

Operations: Patch, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}`

#### Metric

| Field | Description |
| --- | --- |
| `"backfill_percentage"` |  |
| `"backfilled"` |  |
| `"created_at"` |  |
| `"data"` |  |
| `"decimal_place"` |  |
| `"display"` |  |
| `"id"` |  |
| `"last_fetched_at"` |  |
| `"metric"` |  |
| `"metric_identifier"` |  |
| `"metrics_provider_id"` |  |
| `"most_recent_data_at"` |  |
| `"name"` |  |
| `"reference_name"` |  |
| `"suffix"` |  |
| `"tooltip_description"` |  |
| `"updated_at"` |  |
| `"y_axis_hidden"` |  |
| `"y_axis_max"` |  |
| `"y_axis_min"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/metrics/{metric_id}/data`

#### MetricsProvider

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"disabled"` |  |
| `"id"` |  |
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
| `"allow_email_subscriber"` |  |
| `"allow_incident_subscriber"` |  |
| `"allow_page_subscriber"` |  |
| `"allow_rss_atom_feed"` |  |
| `"allow_sms_subscriber"` |  |
| `"allow_webhook_subscriber"` |  |
| `"branding"` |  |
| `"city"` |  |
| `"country"` |  |
| `"created_at"` |  |
| `"css_blue"` |  |
| `"css_body_background_color"` |  |
| `"css_border_color"` |  |
| `"css_font_color"` |  |
| `"css_graph_color"` |  |
| `"css_green"` |  |
| `"css_light_font_color"` |  |
| `"css_link_color"` |  |
| `"css_no_data"` |  |
| `"css_orange"` |  |
| `"css_red"` |  |
| `"css_yellow"` |  |
| `"domain"` |  |
| `"email_logo"` |  |
| `"favicon_logo"` |  |
| `"headline"` |  |
| `"hero_cover"` |  |
| `"hidden_from_search"` |  |
| `"id"` |  |
| `"ip_restriction"` |  |
| `"name"` |  |
| `"notifications_email_footer"` |  |
| `"notifications_from_email"` |  |
| `"page"` |  |
| `"page_description"` |  |
| `"state"` |  |
| `"subdomain"` |  |
| `"support_url"` |  |
| `"time_zone"` |  |
| `"transactional_logo"` |  |
| `"twitter_logo"` |  |
| `"twitter_username"` |  |
| `"updated_at"` |  |
| `"url"` |  |
| `"viewers_must_be_team_member"` |  |

Operations: List, Load, Patch, Update.

API path: `/pages`

#### PageAccessGroup

| Field | Description |
| --- | --- |
| `"component_id"` |  |
| `"created_at"` |  |
| `"external_identifier"` |  |
| `"id"` |  |
| `"metric_id"` |  |
| `"name"` |  |
| `"page_access_group"` |  |
| `"page_access_user_id"` |  |
| `"page_id"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_groups/{page_access_group_id}/components`

#### PageAccessUser

| Field | Description |
| --- | --- |
| `"component_id"` |  |
| `"created_at"` |  |
| `"email"` |  |
| `"external_login"` |  |
| `"id"` |  |
| `"metric_id"` |  |
| `"page_access_group_id"` |  |
| `"page_access_user"` |  |
| `"page_id"` |  |
| `"updated_at"` |  |

Operations: Create, List, Load, Patch, Remove, Update.

API path: `/pages/{page_id}/page_access_users/{page_access_user_id}/components`

#### Permission

| Field | Description |
| --- | --- |
| `"data"` |  |
| `"page"` |  |

Operations: Load, Update.

API path: `/organizations/{organization_id}/permissions/{user_id}`

#### Postmortem

| Field | Description |
| --- | --- |
| `"body"` |  |
| `"body_draft"` |  |
| `"body_draft_updated_at"` |  |
| `"body_updated_at"` |  |
| `"created_at"` |  |
| `"custom_tweet"` |  |
| `"notify_subscriber"` |  |
| `"notify_twitter"` |  |
| `"postmortem"` |  |
| `"preview_key"` |  |
| `"published_at"` |  |
| `"updated_at"` |  |

Operations: Load, Update.

API path: `/pages/{page_id}/incidents/{incident_id}/postmortem`

#### StatusEmbedConfig

| Field | Description |
| --- | --- |
| `"incident_background_color"` |  |
| `"incident_text_color"` |  |
| `"maintenance_background_color"` |  |
| `"maintenance_text_color"` |  |
| `"page_id"` |  |
| `"position"` |  |
| `"status_embed_config"` |  |

Operations: Load, Patch, Update.

API path: `/pages/{page_id}/status_embed_config`

#### Subscriber

| Field | Description |
| --- | --- |
| `"component"` |  |
| `"component_id"` |  |
| `"created_at"` |  |
| `"display_phone_number"` |  |
| `"email"` |  |
| `"endpoint"` |  |
| `"id"` |  |
| `"integration_partner"` |  |
| `"mode"` |  |
| `"obfuscated_channel_name"` |  |
| `"page_access_user_id"` |  |
| `"phone_country"` |  |
| `"phone_number"` |  |
| `"purge_at"` |  |
| `"quarantined_at"` |  |
| `"skip_confirmation_notification"` |  |
| `"skip_unsubscription_notification"` |  |
| `"slack"` |  |
| `"sms"` |  |
| `"state"` |  |
| `"subscriber"` |  |
| `"team"` |  |
| `"type"` |  |
| `"webhook"` |  |
| `"workspace_name"` |  |

Operations: Create, List, Load, Remove, Update.

API path: `/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation`

#### User

| Field | Description |
| --- | --- |
| `"created_at"` |  |
| `"email"` |  |
| `"first_name"` |  |
| `"id"` |  |
| `"last_name"` |  |
| `"organization_id"` |  |
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
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `automation_email` | `string` |  |
| `component` | `map[string]any` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `group` | `bool` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `major_outage` | `int` |  |
| `name` | `string` |  |
| `only_show_if_degraded` | `bool` |  |
| `page_id` | `string` |  |
| `partial_outage` | `int` |  |
| `position` | `int` |  |
| `range_end` | `string` |  |
| `range_start` | `string` |  |
| `related_event` | `map[string]any` |  |
| `showcase` | `bool` |  |
| `start_date` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |
| `uptime_percentage` | `float64` |  |
| `warning` | `string` |  |

#### Example: Load

```go
component, err := client.Component(nil).Load(map[string]any{"id": "component_id"}, nil)
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
}, nil)
```


### ComponentGroupUptime

Create an instance: `component_group_uptime := client.ComponentGroupUptime(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `id` | `string` |  |
| `major_outage` | `int` |  |
| `name` | `string` |  |
| `partial_outage` | `int` |  |
| `range_end` | `string` |  |
| `range_start` | `string` |  |
| `related_event` | `map[string]any` |  |
| `uptime_percentage` | `float64` |  |
| `warning` | `string` |  |

#### Example: Load

```go
component_group_uptime, err := client.ComponentGroupUptime(nil).Load(map[string]any{"id": "component_group_uptime_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(component_group_uptime) // the loaded record
```


### GroupComponent

Create an instance: `group_component := client.GroupComponent(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `string` |  |
| `component_group` | `map[string]any` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
group_component, err := client.GroupComponent(nil).Load(map[string]any{"id": "group_component_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(group_component) // the loaded record
```

#### Example: List

```go
group_components, err := client.GroupComponent(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(group_components) // the array of records
```

#### Example: Create

```go
result, err := client.GroupComponent(nil).Create(map[string]any{
    "component_group": /* map[string]any */,
}, nil)
```


### Incident

Create an instance: `incident := client.Incident(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `bool` |  |
| `auto_transition_deliver_notifications_at_start` | `bool` |  |
| `auto_transition_to_maintenance_state` | `bool` |  |
| `auto_transition_to_operational_state` | `bool` |  |
| `component` | `[]any` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `impact` | `string` |  |
| `impact_override` | `string` |  |
| `incident` | `map[string]any` |  |
| `incident_impact` | `[]any` |  |
| `incident_update` | `[]any` |  |
| `metadata` | `map[string]any` |  |
| `monitoring_at` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `postmortem_body` | `string` |  |
| `postmortem_body_last_updated_at` | `string` |  |
| `postmortem_ignored` | `bool` |  |
| `postmortem_notified_subscriber` | `bool` |  |
| `postmortem_notified_twitter` | `bool` |  |
| `postmortem_published_at` | `bool` |  |
| `reminder_interval` | `string` |  |
| `resolved_at` | `string` |  |
| `scheduled_auto_completed` | `bool` |  |
| `scheduled_auto_in_progress` | `bool` |  |
| `scheduled_for` | `string` |  |
| `scheduled_remind_prior` | `bool` |  |
| `scheduled_reminded_at` | `string` |  |
| `scheduled_until` | `string` |  |
| `shortlink` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
incident, err := client.Incident(nil).Load(map[string]any{"id": "incident_id"}, nil)
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
    "incident": /* map[string]any */,
}, nil)
```


### IncidentPostmortem

Create an instance: `incident_postmortem := client.IncidentPostmortem(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Remove(match, ctrl)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `incident_subscriber := client.IncidentSubscriber(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |

#### Example: Create

```go
result, err := client.IncidentSubscriber(nil).Create(map[string]any{
}, nil)
```


### IncidentTemplate

Create an instance: `incident_template := client.IncidentTemplate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` |  |
| `component` | `[]any` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `should_send_notification` | `bool` |  |
| `should_tweet` | `bool` |  |
| `template` | `map[string]any` |  |
| `title` | `string` |  |
| `update_status` | `string` |  |

#### Example: List

```go
incident_templates, err := client.IncidentTemplate(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(incident_templates) // the array of records
```

#### Example: Create

```go
result, err := client.IncidentTemplate(nil).Create(map[string]any{
    "template": /* map[string]any */,
}, nil)
```


### IncidentUpdate

Create an instance: `incident_update := client.IncidentUpdate(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_component` | `[]any` |  |
| `body` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `deliver_notification` | `bool` |  |
| `display_at` | `string` |  |
| `id` | `string` |  |
| `incident_id` | `string` |  |
| `incident_update` | `map[string]any` |  |
| `status` | `string` |  |
| `tweet_id` | `string` |  |
| `twitter_updated_at` | `string` |  |
| `updated_at` | `string` |  |
| `wants_twitter_update` | `bool` |  |


### Metric

Create an instance: `metric := client.Metric(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `backfill_percentage` | `int` |  |
| `backfilled` | `bool` |  |
| `created_at` | `string` |  |
| `data` | `map[string]any` |  |
| `decimal_place` | `int` |  |
| `display` | `bool` |  |
| `id` | `string` |  |
| `last_fetched_at` | `string` |  |
| `metric` | `map[string]any` |  |
| `metric_identifier` | `string` |  |
| `metrics_provider_id` | `string` |  |
| `most_recent_data_at` | `string` |  |
| `name` | `string` |  |
| `reference_name` | `string` |  |
| `suffix` | `string` |  |
| `tooltip_description` | `string` |  |
| `updated_at` | `string` |  |
| `y_axis_hidden` | `bool` |  |
| `y_axis_max` | `float64` |  |
| `y_axis_min` | `float64` |  |

#### Example: Load

```go
metric, err := client.Metric(nil).Load(map[string]any{"id": "metric_id"}, nil)
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
    "data": /* map[string]any */,
}, nil)
```


### MetricsProvider

Create an instance: `metrics_provider := client.MetricsProvider(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `created_at` | `string` |  |
| `disabled` | `bool` |  |
| `id` | `string` |  |
| `last_revalidated_at` | `string` |  |
| `metric_base_uri` | `string` |  |
| `metrics_provider` | `map[string]any` |  |
| `page_id` | `int` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
metrics_provider, err := client.MetricsProvider(nil).Load(map[string]any{"id": "metrics_provider_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(metrics_provider) // the loaded record
```

#### Example: List

```go
metrics_providers, err := client.MetricsProvider(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(metrics_providers) // the array of records
```

#### Example: Create

```go
result, err := client.MetricsProvider(nil).Create(map[string]any{
}, nil)
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
| `allow_email_subscriber` | `bool` |  |
| `allow_incident_subscriber` | `bool` |  |
| `allow_page_subscriber` | `bool` |  |
| `allow_rss_atom_feed` | `bool` |  |
| `allow_sms_subscriber` | `bool` |  |
| `allow_webhook_subscriber` | `bool` |  |
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
| `hidden_from_search` | `bool` |  |
| `id` | `string` |  |
| `ip_restriction` | `string` |  |
| `name` | `string` |  |
| `notifications_email_footer` | `string` |  |
| `notifications_from_email` | `string` |  |
| `page` | `map[string]any` |  |
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
| `viewers_must_be_team_member` | `bool` |  |

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

Create an instance: `page_access_group := client.PageAccessGroup(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `[]any` |  |
| `created_at` | `string` |  |
| `external_identifier` | `string` |  |
| `id` | `string` |  |
| `metric_id` | `[]any` |  |
| `name` | `string` |  |
| `page_access_group` | `map[string]any` |  |
| `page_access_user_id` | `[]any` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
page_access_group, err := client.PageAccessGroup(nil).Load(map[string]any{"id": "page_access_group_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(page_access_group) // the loaded record
```

#### Example: List

```go
page_access_groups, err := client.PageAccessGroup(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(page_access_groups) // the array of records
```

#### Example: Create

```go
result, err := client.PageAccessGroup(nil).Create(map[string]any{
}, nil)
```


### PageAccessUser

Create an instance: `page_access_user := client.PageAccessUser(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `[]any` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `external_login` | `string` |  |
| `id` | `string` |  |
| `metric_id` | `[]any` |  |
| `page_access_group_id` | `string` |  |
| `page_access_user` | `map[string]any` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
page_access_user, err := client.PageAccessUser(nil).Load(map[string]any{"id": "page_access_user_id"}, nil)
if err != nil {
    panic(err)
}
fmt.Println(page_access_user) // the loaded record
```

#### Example: List

```go
page_access_users, err := client.PageAccessUser(nil).List(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(page_access_users) // the array of records
```

#### Example: Create

```go
result, err := client.PageAccessUser(nil).Create(map[string]any{
    "component_id": /* []any */,
    "metric_id": /* []any */,
}, nil)
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
| `data` | `map[string]any` |  |
| `page` | `map[string]any` |  |

#### Example: Load

```go
permission, err := client.Permission(nil).Load(map[string]any{"id": "permission_id"}, nil)
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
| `body` | `string` |  |
| `body_draft` | `string` |  |
| `body_draft_updated_at` | `string` |  |
| `body_updated_at` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `notify_subscriber` | `bool` |  |
| `notify_twitter` | `bool` |  |
| `postmortem` | `map[string]any` |  |
| `preview_key` | `string` |  |
| `published_at` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```go
postmortem, err := client.Postmortem(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(postmortem) // the loaded record
```


### StatusEmbedConfig

Create an instance: `status_embed_config := client.StatusEmbedConfig(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `incident_background_color` | `string` |  |
| `incident_text_color` | `string` |  |
| `maintenance_background_color` | `string` |  |
| `maintenance_text_color` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `status_embed_config` | `map[string]any` |  |

#### Example: Load

```go
status_embed_config, err := client.StatusEmbedConfig(nil).Load(nil, nil)
if err != nil {
    panic(err)
}
fmt.Println(status_embed_config) // the loaded record
```


### Subscriber

Create an instance: `subscriber := client.Subscriber(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Load(match, ctrl)` | Load a single entity by match criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |
| `Update(data, ctrl)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component` | `string` |  |
| `component_id` | `[]any` |  |
| `created_at` | `string` |  |
| `display_phone_number` | `string` |  |
| `email` | `string` |  |
| `endpoint` | `string` |  |
| `id` | `string` |  |
| `integration_partner` | `int` |  |
| `mode` | `string` |  |
| `obfuscated_channel_name` | `string` |  |
| `page_access_user_id` | `string` |  |
| `phone_country` | `string` |  |
| `phone_number` | `string` |  |
| `purge_at` | `string` |  |
| `quarantined_at` | `string` |  |
| `skip_confirmation_notification` | `bool` |  |
| `skip_unsubscription_notification` | `bool` |  |
| `slack` | `int` |  |
| `sms` | `int` |  |
| `state` | `string` |  |
| `subscriber` | `map[string]any` |  |
| `team` | `int` |  |
| `type` | `string` |  |
| `webhook` | `int` |  |
| `workspace_name` | `string` |  |

#### Example: Load

```go
subscriber, err := client.Subscriber(nil).Load(map[string]any{"id": "subscriber_id"}, nil)
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
}, nil)
```


### User

Create an instance: `user := client.User(nil)`

#### Operations

| Method | Description |
| --- | --- |
| `Create(data, ctrl)` | Create a new entity with the given data. |
| `List(match, ctrl)` | List entities matching the criteria. |
| `Remove(match, ctrl)` | Remove the matching entity. |

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
    "user": /* map[string]any */,
}, nil)
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

Entity instances are stateful. After a successful `List`, the entity
stores the returned data and match criteria internally.

```go
component := client.Component(nil)
component.List(nil, nil)

// component.Data() now returns the component data from the last list
// component.Match() returns the last match criteria
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
