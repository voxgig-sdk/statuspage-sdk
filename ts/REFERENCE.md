# Statuspage TypeScript SDK Reference

Complete API reference for the Statuspage TypeScript SDK.


## StatuspageSDK

### Constructor

```ts
new StatuspageSDK(options?: object)
```

Create a new SDK client instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `options` | `object` | SDK configuration options. |
| `options.apikey` | `string` | API key for authentication. |
| `options.base` | `string` | Base URL for API requests. |
| `options.prefix` | `string` | URL prefix appended after base. |
| `options.suffix` | `string` | URL suffix appended after path. |
| `options.headers` | `object` | Custom headers for all requests. |
| `options.feature` | `object` | Feature configuration. |
| `options.system` | `object` | System overrides (e.g. custom fetch). |


### Static Methods

#### `StatuspageSDK.test(testopts?, sdkopts?)`

Create a test client with mock features active.

```ts
const client = StatuspageSDK.test()
```

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `testopts` | `object` | Test feature options. |
| `sdkopts` | `object` | Additional SDK options merged with test defaults. |

**Returns:** `StatuspageSDK` instance in test mode.


### Instance Methods

#### `Component(data?: object)`

Create a new `Component` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ComponentEntity` instance.

#### `ComponentGroupUptime(data?: object)`

Create a new `ComponentGroupUptime` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `ComponentGroupUptimeEntity` instance.

#### `GroupComponent(data?: object)`

Create a new `GroupComponent` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `GroupComponentEntity` instance.

#### `Incident(data?: object)`

Create a new `Incident` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IncidentEntity` instance.

#### `IncidentPostmortem(data?: object)`

Create a new `IncidentPostmortem` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IncidentPostmortemEntity` instance.

#### `IncidentSubscriber(data?: object)`

Create a new `IncidentSubscriber` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IncidentSubscriberEntity` instance.

#### `IncidentTemplate(data?: object)`

Create a new `IncidentTemplate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IncidentTemplateEntity` instance.

#### `IncidentUpdate(data?: object)`

Create a new `IncidentUpdate` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `IncidentUpdateEntity` instance.

#### `Metric(data?: object)`

Create a new `Metric` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MetricEntity` instance.

#### `MetricsProvider(data?: object)`

Create a new `MetricsProvider` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `MetricsProviderEntity` instance.

#### `Page(data?: object)`

Create a new `Page` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PageEntity` instance.

#### `PageAccessGroup(data?: object)`

Create a new `PageAccessGroup` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PageAccessGroupEntity` instance.

#### `PageAccessUser(data?: object)`

Create a new `PageAccessUser` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PageAccessUserEntity` instance.

#### `Permission(data?: object)`

Create a new `Permission` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PermissionEntity` instance.

#### `Postmortem(data?: object)`

Create a new `Postmortem` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `PostmortemEntity` instance.

#### `StatusEmbedConfig(data?: object)`

Create a new `StatusEmbedConfig` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `StatusEmbedConfigEntity` instance.

#### `Subscriber(data?: object)`

Create a new `Subscriber` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `SubscriberEntity` instance.

#### `User(data?: object)`

Create a new `User` entity instance.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `data` | `object` | Initial entity data. |

**Returns:** `UserEntity` instance.

#### `options()`

Return a deep copy of the current SDK options.

**Returns:** `object`

#### `utility()`

Return a copy of the SDK utility object.

**Returns:** `object`

#### `direct(fetchargs?: object)`

Make a direct HTTP request to any API endpoint.

**Parameters:**

| Name | Type | Description |
| --- | --- | --- |
| `fetchargs.path` | `string` | URL path with optional `{param}` placeholders. |
| `fetchargs.method` | `string` | HTTP method (default: `GET`). |
| `fetchargs.params` | `object` | Path parameter values for `{param}` substitution. |
| `fetchargs.query` | `object` | Query string parameters. |
| `fetchargs.headers` | `object` | Request headers (merged with defaults). |
| `fetchargs.body` | `any` | Request body (objects are JSON-serialized). |
| `fetchargs.ctrl` | `object` | Control options (e.g. `{ explain: true }`). |

**Returns:** `Promise<{ ok, status, headers, data } | Error>`

#### `prepare(fetchargs?: object)`

Prepare a fetch definition without sending the request. Accepts the
same parameters as `direct()`.

**Returns:** `Promise<{ url, method, headers, body } | Error>`

#### `tester(testopts?, sdkopts?)`

Alias for `StatuspageSDK.test()`.

**Returns:** `StatuspageSDK` instance in test mode.


---

## ComponentEntity

```ts
const component = client.Component()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `automation_email` | `string` | No |  |
| `component` | `Record<string, any>` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No |  |
| `group` | `boolean` | No |  |
| `group_id` | `string` | No |  |
| `id` | `string` | No |  |
| `major_outage` | `number` | No |  |
| `name` | `string` | No |  |
| `only_show_if_degraded` | `boolean` | No |  |
| `page_id` | `string` | No |  |
| `partial_outage` | `number` | No |  |
| `position` | `number` | No |  |
| `range_end` | `string` | No |  |
| `range_start` | `string` | No |  |
| `related_event` | `Record<string, any>` | No |  |
| `showcase` | `boolean` | No |  |
| `start_date` | `string` | No |  |
| `status` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `uptime_percentage` | `number` | No |  |
| `warning` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Component().create({
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Component().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Component().load({ id: 'component_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Component().remove({ id: 'component_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Component().update({
  id: 'component_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ComponentEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## ComponentGroupUptimeEntity

```ts
const component_group_uptime = client.ComponentGroupUptime()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `id` | `string` | No |  |
| `major_outage` | `number` | No |  |
| `name` | `string` | No |  |
| `partial_outage` | `number` | No |  |
| `range_end` | `string` | No |  |
| `range_start` | `string` | No |  |
| `related_event` | `Record<string, any>` | No |  |
| `uptime_percentage` | `number` | No |  |
| `warning` | `string` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.ComponentGroupUptime().load({ id: 'component_group_uptime_id', page_id: 'page_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `ComponentGroupUptimeEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## GroupComponentEntity

```ts
const group_component = client.GroupComponent()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `string` | No |  |
| `component_group` | `Record<string, any>` | Yes |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No |  |
| `id` | `string` | No |  |
| `name` | `string` | No |  |
| `page_id` | `string` | No |  |
| `position` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GroupComponent().create({
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.GroupComponent().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.GroupComponent().load({ id: 'group_component_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.GroupComponent().remove({ id: 'group_component_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.GroupComponent().update({
  id: 'group_component_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `GroupComponentEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IncidentEntity

```ts
const incident = client.Incident()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | `boolean` | No |  |
| `auto_transition_deliver_notifications_at_start` | `boolean` | No |  |
| `auto_transition_to_maintenance_state` | `boolean` | No |  |
| `auto_transition_to_operational_state` | `boolean` | No |  |
| `component` | `any[]` | No |  |
| `created_at` | `string` | No |  |
| `id` | `string` | No |  |
| `impact` | `string` | No |  |
| `impact_override` | `string` | No |  |
| `incident` | `Record<string, any>` | Yes |  |
| `incident_impact` | `any[]` | No |  |
| `incident_update` | `any[]` | No |  |
| `metadata` | `Record<string, any>` | No |  |
| `monitoring_at` | `string` | No |  |
| `name` | `string` | No |  |
| `page_id` | `string` | No |  |
| `postmortem_body` | `string` | No |  |
| `postmortem_body_last_updated_at` | `string` | No |  |
| `postmortem_ignored` | `boolean` | No |  |
| `postmortem_notified_subscriber` | `boolean` | No |  |
| `postmortem_notified_twitter` | `boolean` | No |  |
| `postmortem_published_at` | `boolean` | No |  |
| `reminder_interval` | `string` | No |  |
| `resolved_at` | `string` | No |  |
| `scheduled_auto_completed` | `boolean` | No |  |
| `scheduled_auto_in_progress` | `boolean` | No |  |
| `scheduled_for` | `string` | No |  |
| `scheduled_remind_prior` | `boolean` | No |  |
| `scheduled_reminded_at` | `string` | No |  |
| `scheduled_until` | `string` | No |  |
| `shortlink` | `string` | No |  |
| `status` | `string` | No |  |
| `updated_at` | `string` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Incident().create({
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Incident().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Incident().load({ id: 'incident_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Incident().remove({ id: 'incident_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Incident().update({
  id: 'incident_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IncidentEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IncidentPostmortemEntity

```ts
const incident_postmortem = client.IncidentPostmortem()
```

### Operations

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.IncidentPostmortem().remove({ id: 'id', page_id: 'page_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IncidentPostmortemEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IncidentSubscriberEntity

```ts
const incident_subscriber = client.IncidentSubscriber()
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IncidentSubscriber().create({
  incident_id: /* string */,
  page_id: /* string */,
  subscriber_id: /* string */,
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IncidentSubscriberEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IncidentTemplateEntity

```ts
const incident_template = client.IncidentTemplate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `string` | No |  |
| `component` | `any[]` | No |  |
| `group_id` | `string` | No |  |
| `id` | `string` | No |  |
| `name` | `string` | No |  |
| `should_send_notification` | `boolean` | No |  |
| `should_tweet` | `boolean` | No |  |
| `template` | `Record<string, any>` | Yes |  |
| `title` | `string` | No |  |
| `update_status` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IncidentTemplate().create({
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.IncidentTemplate().list()
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IncidentTemplateEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## IncidentUpdateEntity

```ts
const incident_update = client.IncidentUpdate()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `affected_component` | `any[]` | No |  |
| `body` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No |  |
| `deliver_notification` | `boolean` | No |  |
| `display_at` | `string` | No |  |
| `id` | `string` | No |  |
| `incident_id` | `string` | No |  |
| `incident_update` | `Record<string, any>` | No |  |
| `status` | `string` | No |  |
| `tweet_id` | `string` | No |  |
| `twitter_updated_at` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `wants_twitter_update` | `boolean` | No |  |

### Operations

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.IncidentUpdate().update({
  id: 'id',
  incident_id: 'incident_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `IncidentUpdateEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MetricEntity

```ts
const metric = client.Metric()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `backfill_percentage` | `number` | No |  |
| `backfilled` | `boolean` | No |  |
| `created_at` | `string` | No |  |
| `data` | `Record<string, any>` | Yes |  |
| `decimal_place` | `number` | No |  |
| `display` | `boolean` | No |  |
| `id` | `string` | No |  |
| `last_fetched_at` | `string` | No |  |
| `metric` | `Record<string, any>` | No |  |
| `metric_identifier` | `string` | No |  |
| `metrics_provider_id` | `string` | No |  |
| `most_recent_data_at` | `string` | No |  |
| `name` | `string` | No |  |
| `reference_name` | `string` | No |  |
| `suffix` | `string` | No |  |
| `tooltip_description` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `y_axis_hidden` | `boolean` | No |  |
| `y_axis_max` | `number` | No |  |
| `y_axis_min` | `number` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Metric().create({
  metrics_provider_id: /* string */,
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Metric().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Metric().load({ id: 'metric_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Metric().remove({ id: 'metric_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Metric().update({
  id: 'metric_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MetricEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## MetricsProviderEntity

```ts
const metrics_provider = client.MetricsProvider()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `disabled` | `boolean` | No |  |
| `id` | `string` | No |  |
| `last_revalidated_at` | `string` | No |  |
| `metric_base_uri` | `string` | No |  |
| `metrics_provider` | `Record<string, any>` | No |  |
| `page_id` | `number` | No |  |
| `type` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.MetricsProvider().create({
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MetricsProvider().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.MetricsProvider().load({ id: 'metrics_provider_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.MetricsProvider().remove({ id: 'metrics_provider_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.MetricsProvider().update({
  id: 'metrics_provider_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `MetricsProviderEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PageEntity

```ts
const page = client.Page()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `activity_score` | `number` | No |  |
| `allow_email_subscriber` | `boolean` | No |  |
| `allow_incident_subscriber` | `boolean` | No |  |
| `allow_page_subscriber` | `boolean` | No |  |
| `allow_rss_atom_feed` | `boolean` | No |  |
| `allow_sms_subscriber` | `boolean` | No |  |
| `allow_webhook_subscriber` | `boolean` | No |  |
| `branding` | `string` | No |  |
| `city` | `string` | No |  |
| `country` | `string` | No |  |
| `created_at` | `string` | No |  |
| `css_blue` | `string` | No |  |
| `css_body_background_color` | `string` | No |  |
| `css_border_color` | `string` | No |  |
| `css_font_color` | `string` | No |  |
| `css_graph_color` | `string` | No |  |
| `css_green` | `string` | No |  |
| `css_light_font_color` | `string` | No |  |
| `css_link_color` | `string` | No |  |
| `css_no_data` | `string` | No |  |
| `css_orange` | `string` | No |  |
| `css_red` | `string` | No |  |
| `css_yellow` | `string` | No |  |
| `domain` | `string` | No |  |
| `email_logo` | `string` | No |  |
| `favicon_logo` | `string` | No |  |
| `headline` | `string` | No |  |
| `hero_cover` | `string` | No |  |
| `hidden_from_search` | `boolean` | No |  |
| `id` | `string` | No |  |
| `ip_restriction` | `string` | No |  |
| `name` | `string` | No |  |
| `notifications_email_footer` | `string` | No |  |
| `notifications_from_email` | `string` | No |  |
| `page` | `Record<string, any>` | No |  |
| `page_description` | `string` | No |  |
| `state` | `string` | No |  |
| `subdomain` | `string` | No |  |
| `support_url` | `string` | No |  |
| `time_zone` | `string` | No |  |
| `transactional_logo` | `string` | No |  |
| `twitter_logo` | `string` | No |  |
| `twitter_username` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `url` | `string` | No |  |
| `viewers_must_be_team_member` | `boolean` | No |  |

### Operations

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Page().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Page().load({ id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Page().update({
  id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PageEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PageAccessGroupEntity

```ts
const page_access_group = client.PageAccessGroup()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `any[]` | No |  |
| `created_at` | `string` | No |  |
| `external_identifier` | `string` | No |  |
| `id` | `string` | No |  |
| `metric_id` | `any[]` | No |  |
| `name` | `string` | No |  |
| `page_access_group` | `Record<string, any>` | No |  |
| `page_access_user_id` | `any[]` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PageAccessGroup().create({
  id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PageAccessGroup().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PageAccessGroup().load({ id: 'page_access_group_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PageAccessGroup().remove({ id: 'page_access_group_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.PageAccessGroup().update({
  id: 'page_access_group_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PageAccessGroupEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PageAccessUserEntity

```ts
const page_access_user = client.PageAccessUser()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component_id` | `any[]` | Yes |  |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `external_login` | `string` | No |  |
| `id` | `string` | No |  |
| `metric_id` | `any[]` | Yes |  |
| `page_access_group_id` | `string` | No |  |
| `page_access_user` | `Record<string, any>` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PageAccessUser().create({
  id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PageAccessUser().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.PageAccessUser().load({ id: 'page_access_user_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.PageAccessUser().remove({ id: 'page_access_user_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.PageAccessUser().update({
  id: 'page_access_user_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PageAccessUserEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PermissionEntity

```ts
const permission = client.Permission()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `data` | `Record<string, any>` | No |  |
| `page` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Permission().load({ id: 'permission_id', organization_id: 'organization_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Permission().update({
  id: 'permission_id',
  organization_id: 'organization_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PermissionEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## PostmortemEntity

```ts
const postmortem = client.Postmortem()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `body` | `string` | No |  |
| `body_draft` | `string` | No |  |
| `body_draft_updated_at` | `string` | No |  |
| `body_updated_at` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No |  |
| `notify_subscriber` | `boolean` | No |  |
| `notify_twitter` | `boolean` | No |  |
| `postmortem` | `Record<string, any>` | Yes |  |
| `preview_key` | `string` | No |  |
| `published_at` | `string` | No |  |
| `updated_at` | `string` | No |  |

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

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Postmortem().load({ incident_id: 'incident_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Postmortem().update({
  incident_id: 'incident_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `PostmortemEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## StatusEmbedConfigEntity

```ts
const status_embed_config = client.StatusEmbedConfig()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `incident_background_color` | `string` | No |  |
| `incident_text_color` | `string` | No |  |
| `maintenance_background_color` | `string` | No |  |
| `maintenance_text_color` | `string` | No |  |
| `page_id` | `string` | No |  |
| `position` | `string` | No |  |
| `status_embed_config` | `Record<string, any>` | No |  |

### Operations

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.StatusEmbedConfig().load({ page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.StatusEmbedConfig().update({
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `StatusEmbedConfigEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## SubscriberEntity

```ts
const subscriber = client.Subscriber()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `component` | `string` | No |  |
| `component_id` | `any[]` | No |  |
| `created_at` | `string` | No |  |
| `display_phone_number` | `string` | No |  |
| `email` | `string` | No |  |
| `endpoint` | `string` | No |  |
| `id` | `string` | No |  |
| `integration_partner` | `number` | No |  |
| `mode` | `string` | No |  |
| `obfuscated_channel_name` | `string` | No |  |
| `page_access_user_id` | `string` | No |  |
| `phone_country` | `string` | No |  |
| `phone_number` | `string` | No |  |
| `purge_at` | `string` | No |  |
| `quarantined_at` | `string` | No |  |
| `skip_confirmation_notification` | `boolean` | No |  |
| `skip_unsubscription_notification` | `boolean` | No |  |
| `slack` | `number` | No |  |
| `sms` | `number` | No |  |
| `state` | `string` | No |  |
| `subscriber` | `Record<string, any>` | No |  |
| `team` | `number` | No |  |
| `type` | `string` | No |  |
| `webhook` | `number` | No |  |
| `workspace_name` | `string` | No |  |

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

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Subscriber().create({
  page_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Subscriber().list()
```

#### `load(match: object, ctrl?: object)`

Load a single entity matching the given criteria.

```ts
const result = await client.Subscriber().load({ id: 'subscriber_id', page_id: 'page_id' })
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.Subscriber().remove({ id: 'subscriber_id', page_id: 'page_id' })
```

#### `update(data: object, ctrl?: object)`

Update an existing entity. The data must include the entity `id`.

```ts
const result = await client.Subscriber().update({
  id: 'subscriber_id',
  page_id: 'page_id',
  // Fields to update
})
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `SubscriberEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## UserEntity

```ts
const user = client.User()
```

### Fields

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `first_name` | `string` | No |  |
| `id` | `string` | No |  |
| `last_name` | `string` | No |  |
| `organization_id` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.User().create({
  organization_id: /* string */,
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.User().list()
```

#### `remove(match: object, ctrl?: object)`

Remove the entity matching the given criteria.

```ts
const result = await client.User().remove({ id: 'id', organization_id: 'organization_id' })
```

### Common Methods

#### `data(data?: object)`

Get or set the entity data. When called with data, sets the entity's
internal data and returns the current data. When called without
arguments, returns a copy of the current data.

#### `match(match?: object)`

Get or set the entity match criteria. Works the same as `data()`.

#### `make()`

Create a new `UserEntity` instance with the same client and
options.

#### `client()`

Return the parent `StatuspageSDK` instance.

#### `entopts()`

Return a copy of the entity options.


---

## Features

| Feature | Version | Description |
| --- | --- | --- |
| `test` | 0.0.1 | In-memory mock transport for testing without a live server |


Features are activated via the `feature` option:

```ts
const client = new StatuspageSDK({
  feature: {
    test: { active: true },
  }
})
```

