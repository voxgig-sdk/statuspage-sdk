# Statuspage JavaScript SDK



The JavaScript SDK for the Statuspage API — an entity-oriented client with full async/await support.

The API is exposed as capitalised, semantic **Entities** — e.g.
`client.Component()` — each with a small set of operations (`list`, `load`, `create`, `update`, `remove`, `patch`)
instead of raw URL paths and query parameters. This keeps the surface
predictable and low-friction for both humans and AI agents.

> Other languages, the CLI, and MCP server live alongside this one — see
> the [top-level README](../README.md).


## Install
```js
npm install statuspage
```
## Tutorial: your first API call

This tutorial walks through creating a client, listing entities, and
loading a specific record.


### Create a Client

```js
const { StatuspageSDK } = require('@voxgig-sdk/statuspage-js')

const client = new StatuspageSDK({
  apikey: process.env.STATUSPAGE_APIKEY,
})
```

### Load a Component

```js
const component = await client.Component().load({ id: 'component_id', page_id: 'example_page_id' })
console.log(component)
```

### List Component Records

```js
const components = await client.Component().list()
for (const component of components) {
  console.log(component)
}
```

### Create a Component

```js
const created = await client.Component().create({
  page_id: 'example_page_id',
})
console.log(created)
```

### Update a Component

```js
const updated = await client.Component().update({
  id: 'component_id',
  page_id: 'example_page_id',
  automation_email: 'example_automation_email',
})
console.log(updated)
```

### Remove a Component

```js
await client.Component().remove({ id: 'component_id', page_id: 'example_page_id' })
```

### Direct API Access

Use `client.direct()` to call any API endpoint directly:

```js
const result = await client.direct({
  path: '/custom/endpoint/{id}',
  method: 'GET',
  params: { id: 'abc123' },
})

if (result.ok) {
  console.log(result.data)
}
```


## Error handling

Entity operations reject on failure, so wrap them in `try` / `catch`:

```ts
try {
  const postmortem = await client.Postmortem().load({ incident_id: "example", page_id: "example" })
  console.log(postmortem)
} catch (err) {
  console.error('load failed:', err)
}
```

The low-level `direct()` method does **not** throw — it returns the
value or an `Error`, so check the result before using it:

```ts
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example_id' },
})

if (result instanceof Error) {
  throw result
}
```


## How-to guides

### Make a direct HTTP request

For endpoints not covered by entity methods:

```js
const result = await client.direct({
  path: '/api/resource/{id}',
  method: 'GET',
  params: { id: 'example' },
})

if (result instanceof Error) {
  throw result
}
if (result.ok) {
  console.log(result.status)  // 200
  console.log(result.data)    // response body
}
```

### Prepare a request without sending it

```js
const fetchdef = await client.prepare({
  path: '/api/resource/{id}',
  method: 'DELETE',
  params: { id: 'example' },
})

// Inspect before sending
console.log(fetchdef.url)
console.log(fetchdef.method)
console.log(fetchdef.headers)
```

### Use test mode

Create a mock client for unit testing — no server required:

```js
const client = StatuspageSDK.test()

const postmortem = await client.Postmortem().load({ incident_id: 'example_incident_id', page_id: 'example_page_id' })
// postmortem is a bare entity populated with mock response data
console.log(postmortem)
```

You can also use the instance method:

```js
const client = new StatuspageSDK({ apikey: '...' })
const testClient = client.tester()
```

### Retain entity state across calls

Entity instances remember their last match and data:

```js
const entity = client.Postmortem()

// First call runs the operation and stores its result
await entity.load({ incident_id: 'example_incident_id', page_id: 'example_page_id' })

// Subsequent calls reuse the stored state
const data = entity.data()
console.log(data)
```

### Add custom middleware

Pass features via the `extend` option:

```js
const logger = {
  hooks: {
    PreRequest: (ctx) => {
      console.log('Requesting:', ctx.spec.method, ctx.spec.path)
    },
    PreResponse: (ctx) => {
      console.log('Status:', ctx.out.request?.status)
    },
  },
}

const client = new StatuspageSDK({
  apikey: '...',
  extend: [logger],
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
cd js && npm test
```


## Reference

### StatuspageSDK

#### Constructor

```js
new StatuspageSDK(options?)
```

| Option | Type | Description |
| --- | --- | --- |
| `apikey` | `string` | API key for authentication. |
| `base` | `string` | Base URL of the API server. |
| `prefix` | `string` | URL path prefix prepended to all requests. |
| `suffix` | `string` | URL path suffix appended to all requests. |
| `feature` | `object` | Feature activation flags (e.g. `{ test: { active: true } }`). |
| `extend` | `Feature[]` | Additional feature instances to load. |

#### Methods

| Method | Returns | Description |
| --- | --- | --- |
| `options()` | `object` | Deep copy of current SDK options. |
| `utility()` | `Utility` | Deep copy of the SDK utility object. |
| `prepare(fetchargs?)` | `Promise<FetchDef>` | Build an HTTP request definition without sending it. |
| `direct(fetchargs?)` | `Promise<DirectResult>` | Build and send an HTTP request. |
| `Component(data?)` | `ComponentEntity` | Create a Component entity instance. |
| `ComponentGroupUptime(data?)` | `ComponentGroupUptimeEntity` | Create a ComponentGroupUptime entity instance. |
| `GroupComponent(data?)` | `GroupComponentEntity` | Create a GroupComponent entity instance. |
| `Incident(data?)` | `IncidentEntity` | Create an Incident entity instance. |
| `IncidentPostmortem(data?)` | `IncidentPostmortemEntity` | Create an IncidentPostmortem entity instance. |
| `IncidentSubscriber(data?)` | `IncidentSubscriberEntity` | Create an IncidentSubscriber entity instance. |
| `IncidentTemplate(data?)` | `IncidentTemplateEntity` | Create an IncidentTemplate entity instance. |
| `IncidentUpdate(data?)` | `IncidentUpdateEntity` | Create an IncidentUpdate entity instance. |
| `Metric(data?)` | `MetricEntity` | Create a Metric entity instance. |
| `MetricsProvider(data?)` | `MetricsProviderEntity` | Create a MetricsProvider entity instance. |
| `Page(data?)` | `PageEntity` | Create a Page entity instance. |
| `PageAccessGroup(data?)` | `PageAccessGroupEntity` | Create a PageAccessGroup entity instance. |
| `PageAccessUser(data?)` | `PageAccessUserEntity` | Create a PageAccessUser entity instance. |
| `Permission(data?)` | `PermissionEntity` | Create a Permission entity instance. |
| `Postmortem(data?)` | `PostmortemEntity` | Create a Postmortem entity instance. |
| `StatusEmbedConfig(data?)` | `StatusEmbedConfigEntity` | Create a StatusEmbedConfig entity instance. |
| `Subscriber(data?)` | `SubscriberEntity` | Create a Subscriber entity instance. |
| `User(data?)` | `UserEntity` | Create an User entity instance. |
| `tester(testopts?, sdkopts?)` | `StatuspageSDK` | Create a test-mode client instance. |

#### Static methods

| Method | Returns | Description |
| --- | --- | --- |
| `StatuspageSDK.test(testopts?, sdkopts?)` | `StatuspageSDK` | Create a test-mode client. |

### Entity interface

All entities share the same interface.

#### Methods

| Method | Signature | Description |
| --- | --- | --- |
| `load` | `load(reqmatch?, ctrl?): Promise<Entity>` | Load a single entity by match criteria. |
| `list` | `list(reqmatch?, ctrl?): Promise<Entity[]>` | List entities matching the criteria. |
| `create` | `create(reqdata?, ctrl?): Promise<Entity>` | Create a new entity. |
| `update` | `update(reqdata?, ctrl?): Promise<Entity>` | Update an existing entity. |
| `remove` | `remove(reqmatch?, ctrl?): Promise<void>` | Remove an entity. |
| `data` | `data(data?: Partial<Entity>): Entity` | Get or set entity data. |
| `match` | `match(match?: Partial<Entity>): Partial<Entity>` | Get or set entity match criteria. |
| `make` | `make(): Entity` | Create a new instance with the same options. |
| `client` | `client(): StatuspageSDK` | Return the parent SDK client. |
| `entopts` | `entopts(): object` | Return a copy of the entity options. |

#### Return values

Entity operations resolve to the entity data directly — there is no
result envelope:

- `load`, `create` and `update` resolve to a single entity object.
- `list` resolves to an **array** of entity objects (iterate it directly;
  there is no `.data` and no `.ok`).
- `remove` resolves to `undefined`.

On a failed request these methods **throw**, so wrap calls in
`try`/`catch` to handle errors. Only `direct()` returns the result
envelope described below.

### DirectResult shape

The `direct()` method returns:

```js
{
  ok: true,
  status: 200,
  headers: {},
  data: {}
}
```

On error, `ok` is `false` and an `err` property contains the error.

### FetchDef shape

The `prepare()` method returns:

```js
{
  url: 'string',
  method: 'string',
  headers: {},
  body: undefined
}
```

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
| `name` |  |
| `only_show_if_degraded` |  |
| `page_id` |  |
| `position` |  |
| `showcase` |  |
| `start_date` |  |
| `status` |  |
| `updated_at` |  |

Operations: create, list, load, patch, remove, update.

API path: `/pages/{page_id}/components/{component_id}/page_access_groups`

#### ComponentGroupUptime

| Field | Description |
| --- | --- |
| `component_id` |  |
| `incidents` |  |

Operations: load.

API path: `/pages/{page_id}/component-groups/{id}/uptime`

#### GroupComponent

| Field | Description |
| --- | --- |
| `component_group` |  |
| `components` |  |
| `created_at` |  |
| `description` |  |
| `id` |  |
| `name` |  |
| `page_id` |  |
| `position` |  |
| `updated_at` |  |

Operations: create, list, load, patch, remove, update.

API path: `/pages/{page_id}/component-groups`

#### Incident

| Field | Description |
| --- | --- |
| `auto_transition_deliver_notifications_at_end` |  |
| `auto_transition_deliver_notifications_at_start` |  |
| `auto_transition_to_maintenance_state` |  |
| `auto_transition_to_operational_state` |  |
| `components` |  |
| `created_at` |  |
| `id` |  |
| `impact` |  |
| `impact_override` |  |
| `incident` |  |
| `incident_updates` |  |
| `metadata` |  |
| `monitoring_at` |  |
| `name` |  |
| `page_id` |  |
| `postmortem_body` |  |
| `postmortem_body_last_updated_at` |  |
| `postmortem_ignored` |  |
| `postmortem_notified_subscribers` |  |
| `postmortem_notified_twitter` |  |
| `postmortem_published_at` |  |
| `reminder_intervals` |  |
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

Operations: create, list, load, patch, remove, update.

API path: `/pages/{page_id}/incidents`

#### IncidentPostmortem

| Field | Description |
| --- | --- |

Operations: remove.

API path: `/pages/{page_id}/incidents/{incident_id}/postmortem`

#### IncidentSubscriber

| Field | Description |
| --- | --- |

Operations: create.

API path: `/pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation`

#### IncidentTemplate

| Field | Description |
| --- | --- |
| `body` |  |
| `components` |  |
| `group_id` |  |
| `id` |  |
| `name` |  |
| `should_send_notifications` |  |
| `should_tweet` |  |
| `template` |  |
| `title` |  |
| `update_status` |  |

Operations: create, list.

API path: `/pages/{page_id}/incident_templates`

#### IncidentUpdate

| Field | Description |
| --- | --- |
| `affected_components` |  |
| `body` |  |
| `created_at` |  |
| `custom_tweet` |  |
| `deliver_notifications` |  |
| `display_at` |  |
| `id` |  |
| `incident_id` |  |
| `incident_update` |  |
| `status` |  |
| `tweet_id` |  |
| `twitter_updated_at` |  |
| `updated_at` |  |
| `wants_twitter_update` |  |

Operations: patch, update.

API path: `/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}`

#### Metric

| Field | Description |
| --- | --- |
| `backfill_percentage` |  |
| `backfilled` |  |
| `created_at` |  |
| `data` |  |
| `decimal_places` |  |
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

Operations: create, list, load, patch, remove, update.

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

Operations: create, list, load, patch, remove, update.

API path: `/pages/{page_id}/metrics_providers`

#### Page

| Field | Description |
| --- | --- |
| `activity_score` |  |
| `allow_email_subscribers` |  |
| `allow_incident_subscribers` |  |
| `allow_page_subscribers` |  |
| `allow_rss_atom_feeds` |  |
| `allow_sms_subscribers` |  |
| `allow_webhook_subscribers` |  |
| `branding` |  |
| `city` |  |
| `country` |  |
| `created_at` |  |
| `css_blues` |  |
| `css_body_background_color` |  |
| `css_border_color` |  |
| `css_font_color` |  |
| `css_graph_color` |  |
| `css_greens` |  |
| `css_light_font_color` |  |
| `css_link_color` |  |
| `css_no_data` |  |
| `css_oranges` |  |
| `css_reds` |  |
| `css_yellows` |  |
| `domain` |  |
| `email_logo` |  |
| `favicon_logo` |  |
| `headline` |  |
| `hero_cover` |  |
| `hidden_from_search` |  |
| `id` |  |
| `ip_restrictions` |  |
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
| `viewers_must_be_team_members` |  |

Operations: list, load, patch, update.

API path: `/pages`

#### PageAccessGroup

| Field | Description |
| --- | --- |
| `component_ids` |  |
| `created_at` |  |
| `external_identifier` |  |
| `id` |  |
| `metric_ids` |  |
| `name` |  |
| `page_access_group` |  |
| `page_access_user_ids` |  |
| `page_id` |  |
| `updated_at` |  |

Operations: create, list, load, patch, remove, update.

API path: `/pages/{page_id}/page_access_groups/{page_access_group_id}/components`

#### PageAccessUser

| Field | Description |
| --- | --- |
| `component_ids` |  |
| `created_at` |  |
| `email` |  |
| `external_login` |  |
| `id` |  |
| `metric_ids` |  |
| `page_access_group_id` |  |
| `page_access_group_ids` |  |
| `page_access_user` |  |
| `page_id` |  |
| `updated_at` |  |

Operations: create, list, load, patch, remove, update.

API path: `/pages/{page_id}/page_access_users/{page_access_user_id}/components`

#### Permission

| Field | Description |
| --- | --- |
| `pages` |  |
| `user_id` |  |

Operations: load, update.

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
| `notify_subscribers` |  |
| `notify_twitter` |  |
| `postmortem` |  |
| `preview_key` |  |
| `published_at` |  |
| `updated_at` |  |

Operations: load, update.

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

Operations: load, patch, update.

API path: `/pages/{page_id}/status_embed_config`

#### Subscriber

| Field | Description |
| --- | --- |
| `component_ids` |  |
| `components` |  |
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
| `subscribers` |  |
| `teams` |  |
| `type` |  |
| `webhook` |  |
| `workspace_name` |  |

Operations: create, list, load, remove, update.

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

Operations: create, list, remove.

API path: `/organizations/{organization_id}/users`



## Entities


### Component

Create an instance: `const component = client.Component()`

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
| `component` | `Object` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `group` | `boolean` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `only_show_if_degraded` | `boolean` |  |
| `page_id` | `string` |  |
| `position` | `number` |  |
| `showcase` | `boolean` |  |
| `start_date` | `string` |  |
| `status` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const component = await client.Component().load({ id: 'component_id', page_id: 'page_id' })
```

#### Example: List

```ts
const components = await client.Component().list()
```

#### Example: Create

```ts
const component = await client.Component().create({
  page_id: 'example_page_id',
})
```


### ComponentGroupUptime

Create an instance: `const component_group_uptime = client.ComponentGroupUptime()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `component_id` | `string` |  |
| `incidents` | `Object` |  |

#### Example: Load

```ts
const component_group_uptime = await client.ComponentGroupUptime().load({ id: 'component_group_uptime_id', page_id: 'page_id' })
```


### GroupComponent

Create an instance: `const group_component = client.GroupComponent()`

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
| `component_group` | `Object` |  |
| `components` | `string` |  |
| `created_at` | `string` |  |
| `description` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `position` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const group_component = await client.GroupComponent().load({ id: 'group_component_id', page_id: 'page_id' })
```

#### Example: List

```ts
const group_components = await client.GroupComponent().list()
```

#### Example: Create

```ts
const group_component = await client.GroupComponent().create({
  page_id: 'example_page_id',
  component_group: {},
})
```


### Incident

Create an instance: `const incident = client.Incident()`

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
| `components` | `Array` |  |
| `created_at` | `string` |  |
| `id` | `string` |  |
| `impact` | `string` |  |
| `impact_override` | `string` |  |
| `incident` | `Object` |  |
| `incident_updates` | `Array` |  |
| `metadata` | `Object` |  |
| `monitoring_at` | `string` |  |
| `name` | `string` |  |
| `page_id` | `string` |  |
| `postmortem_body` | `string` |  |
| `postmortem_body_last_updated_at` | `string` |  |
| `postmortem_ignored` | `boolean` |  |
| `postmortem_notified_subscribers` | `boolean` |  |
| `postmortem_notified_twitter` | `boolean` |  |
| `postmortem_published_at` | `boolean` |  |
| `reminder_intervals` | `string` |  |
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

```ts
const incident = await client.Incident().load({ id: 'incident_id', page_id: 'page_id' })
```

#### Example: List

```ts
const incidents = await client.Incident().list()
```

#### Example: Create

```ts
const incident = await client.Incident().create({
  page_id: 'example_page_id',
  incident: {},
})
```


### IncidentPostmortem

Create an instance: `const incident_postmortem = client.IncidentPostmortem()`

#### Operations

| Method | Description |
| --- | --- |
| `remove(match)` | Remove the matching entity. |


### IncidentSubscriber

Create an instance: `const incident_subscriber = client.IncidentSubscriber()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |

#### Example: Create

```ts
const incident_subscriber = await client.IncidentSubscriber().create({
  incident_id: 'example_incident_id',
  page_id: 'example_page_id',
  subscriber_id: 'example_subscriber_id',
})
```


### IncidentTemplate

Create an instance: `const incident_template = client.IncidentTemplate()`

#### Operations

| Method | Description |
| --- | --- |
| `create(data)` | Create a new entity with the given data. |
| `list(match)` | List entities matching the criteria. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `body` | `string` |  |
| `components` | `Array` |  |
| `group_id` | `string` |  |
| `id` | `string` |  |
| `name` | `string` |  |
| `should_send_notifications` | `boolean` |  |
| `should_tweet` | `boolean` |  |
| `template` | `Object` |  |
| `title` | `string` |  |
| `update_status` | `string` |  |

#### Example: List

```ts
const incident_templates = await client.IncidentTemplate().list()
```

#### Example: Create

```ts
const incident_template = await client.IncidentTemplate().create({
  page_id: 'example_page_id',
  template: {},
})
```


### IncidentUpdate

Create an instance: `const incident_update = client.IncidentUpdate()`

#### Operations

| Method | Description |
| --- | --- |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `affected_components` | `Array` |  |
| `body` | `string` |  |
| `created_at` | `string` |  |
| `custom_tweet` | `string` |  |
| `deliver_notifications` | `boolean` |  |
| `display_at` | `string` |  |
| `id` | `string` |  |
| `incident_id` | `string` |  |
| `incident_update` | `Object` |  |
| `status` | `string` |  |
| `tweet_id` | `string` |  |
| `twitter_updated_at` | `string` |  |
| `updated_at` | `string` |  |
| `wants_twitter_update` | `boolean` |  |


### Metric

Create an instance: `const metric = client.Metric()`

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
| `data` | `Object` |  |
| `decimal_places` | `number` |  |
| `display` | `boolean` |  |
| `id` | `string` |  |
| `last_fetched_at` | `string` |  |
| `metric` | `Object` |  |
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

```ts
const metric = await client.Metric().load({ id: 'metric_id', page_id: 'page_id' })
```

#### Example: List

```ts
const metrics = await client.Metric().list()
```

#### Example: Create

```ts
const metric = await client.Metric().create({
  metrics_provider_id: 'example_metrics_provider_id',
  page_id: 'example_page_id',
  data: {},
})
```


### MetricsProvider

Create an instance: `const metrics_provider = client.MetricsProvider()`

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
| `metrics_provider` | `Object` |  |
| `page_id` | `number` |  |
| `type` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const metrics_provider = await client.MetricsProvider().load({ id: 'metrics_provider_id', page_id: 'page_id' })
```

#### Example: List

```ts
const metrics_providers = await client.MetricsProvider().list()
```

#### Example: Create

```ts
const metrics_provider = await client.MetricsProvider().create({
  page_id: 'example_page_id',
})
```


### Page

Create an instance: `const page = client.Page()`

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
| `allow_email_subscribers` | `boolean` |  |
| `allow_incident_subscribers` | `boolean` |  |
| `allow_page_subscribers` | `boolean` |  |
| `allow_rss_atom_feeds` | `boolean` |  |
| `allow_sms_subscribers` | `boolean` |  |
| `allow_webhook_subscribers` | `boolean` |  |
| `branding` | `string` |  |
| `city` | `string` |  |
| `country` | `string` |  |
| `created_at` | `string` |  |
| `css_blues` | `string` |  |
| `css_body_background_color` | `string` |  |
| `css_border_color` | `string` |  |
| `css_font_color` | `string` |  |
| `css_graph_color` | `string` |  |
| `css_greens` | `string` |  |
| `css_light_font_color` | `string` |  |
| `css_link_color` | `string` |  |
| `css_no_data` | `string` |  |
| `css_oranges` | `string` |  |
| `css_reds` | `string` |  |
| `css_yellows` | `string` |  |
| `domain` | `string` |  |
| `email_logo` | `string` |  |
| `favicon_logo` | `string` |  |
| `headline` | `string` |  |
| `hero_cover` | `string` |  |
| `hidden_from_search` | `boolean` |  |
| `id` | `string` |  |
| `ip_restrictions` | `string` |  |
| `name` | `string` |  |
| `notifications_email_footer` | `string` |  |
| `notifications_from_email` | `string` |  |
| `page` | `Object` |  |
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
| `viewers_must_be_team_members` | `boolean` |  |

#### Example: Load

```ts
const page = await client.Page().load({ id: 'page_id' })
```

#### Example: List

```ts
const pages = await client.Page().list()
```


### PageAccessGroup

Create an instance: `const page_access_group = client.PageAccessGroup()`

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
| `component_ids` | `Array` |  |
| `created_at` | `string` |  |
| `external_identifier` | `string` |  |
| `id` | `string` |  |
| `metric_ids` | `Array` |  |
| `name` | `string` |  |
| `page_access_group` | `Object` |  |
| `page_access_user_ids` | `Array` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const page_access_group = await client.PageAccessGroup().load({ id: 'page_access_group_id', page_id: 'page_id' })
```

#### Example: List

```ts
const page_access_groups = await client.PageAccessGroup().list()
```

#### Example: Create

```ts
const page_access_group = await client.PageAccessGroup().create({
  id: 'example_id',
})
```


### PageAccessUser

Create an instance: `const page_access_user = client.PageAccessUser()`

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
| `component_ids` | `Array` |  |
| `created_at` | `string` |  |
| `email` | `string` |  |
| `external_login` | `string` |  |
| `id` | `string` |  |
| `metric_ids` | `Array` |  |
| `page_access_group_id` | `string` |  |
| `page_access_group_ids` | `string` |  |
| `page_access_user` | `Object` |  |
| `page_id` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const page_access_user = await client.PageAccessUser().load({ id: 'page_access_user_id', page_id: 'page_id' })
```

#### Example: List

```ts
const page_access_users = await client.PageAccessUser().list()
```

#### Example: Create

```ts
const page_access_user = await client.PageAccessUser().create({
  id: 'example_id',
  component_ids: [],
  metric_ids: [],
})
```


### Permission

Create an instance: `const permission = client.Permission()`

#### Operations

| Method | Description |
| --- | --- |
| `load(match)` | Load a single entity by match criteria. |
| `update(data)` | Update an existing entity. |

#### Fields

| Field | Type | Description |
| --- | --- | --- |
| `pages` | `Object` |  |
| `user_id` | `string` |  |

#### Example: Load

```ts
const permission = await client.Permission().load({ id: 'permission_id', organization_id: 'organization_id' })
```


### Postmortem

Create an instance: `const postmortem = client.Postmortem()`

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
| `notify_subscribers` | `boolean` |  |
| `notify_twitter` | `boolean` |  |
| `postmortem` | `Object` |  |
| `preview_key` | `string` |  |
| `published_at` | `string` |  |
| `updated_at` | `string` |  |

#### Example: Load

```ts
const postmortem = await client.Postmortem().load({ incident_id: 'incident_id', page_id: 'page_id' })
```


### StatusEmbedConfig

Create an instance: `const status_embed_config = client.StatusEmbedConfig()`

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
| `status_embed_config` | `Object` |  |

#### Example: Load

```ts
const status_embed_config = await client.StatusEmbedConfig().load({ page_id: 'page_id' })
```


### Subscriber

Create an instance: `const subscriber = client.Subscriber()`

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
| `component_ids` | `Array` |  |
| `components` | `string` |  |
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
| `subscriber` | `Object` |  |
| `subscribers` | `string` |  |
| `teams` | `number` |  |
| `type` | `string` |  |
| `webhook` | `number` |  |
| `workspace_name` | `string` |  |

#### Example: Load

```ts
const subscriber = await client.Subscriber().load({ id: 'subscriber_id', page_id: 'page_id' })
```

#### Example: List

```ts
const subscribers = await client.Subscriber().list()
```

#### Example: Create

```ts
const subscriber = await client.Subscriber().create({
  page_id: 'example_page_id',
  subscribers: 'example_subscribers',
})
```


### User

Create an instance: `const user = client.User()`

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
| `user` | `Object` |  |

#### Example: List

```ts
const users = await client.User().list()
```

#### Example: Create

```ts
const user = await client.User().create({
  organization_id: 'example_organization_id',
  user: {},
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

Features are the extension mechanism. A feature is an object with a
`hooks` map. Each hook key is a pipeline stage name, and the value is
a function that receives the context.

The SDK ships with built-in features:

- **TestFeature**: In-memory mock transport for testing without a live server

Features are initialized in order. Hooks fire in the order features
were added, so later features can override earlier ones.

### Module structure

```
statuspage/
├── src/
│   ├── StatuspageSDK.js        # Main SDK class
│   ├── entity/             # Entity implementations
│   ├── feature/            # Built-in features (Base, Test, Log)
│   └── utility/            # Utility functions
└── test/                   # Test suites
```

Import the SDK from the package root:

```js
const { StatuspageSDK } = require('@voxgig-sdk/statuspage-js')
```

### Entity state

Entity instances are stateful. After a successful `load`, the entity
stores the returned data and match criteria internally. Subsequent
calls on the same instance can rely on this state.

```ts
const postmortem = client.Postmortem()
await postmortem.load({ incident_id: "example", page_id: "example" })

// postmortem.data() now returns the postmortem data from the last `load`
// postmortem.match() returns the last match criteria
```

Call `make()` to create a fresh instance with the same configuration
but no stored state.

### Direct vs entity access

The entity interface handles URL construction, parameter placement,
and response parsing automatically. Use it for standard CRUD operations.

The `direct` method gives full control over the HTTP request. Use it
for non-standard endpoints, bulk operations, or any path not modelled
as an entity. The `prepare` method is useful for debugging — it
shows exactly what `direct` would send.


## Full Reference

See [REFERENCE.md](REFERENCE.md) for complete API reference
documentation including all method signatures, entity field schemas,
and detailed usage examples.
