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
| `automation_email` | `string` | No | Requires a special feature flag to be enabled |
| `component` | `Record<string, any>` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No | More detailed description for component |
| `group` | `boolean` | No | Is this component a group |
| `group_id` | `string` | No | Component Group identifier |
| `id` | `string` | No | Incident identifier |
| `name` | `string` | No | Display name for component |
| `only_show_if_degraded` | `boolean` | No | Requires a special feature flag to be enabled |
| `page_id` | `string` | No | Page identifier |
| `position` | `number` | No | Order the component will appear on the page |
| `showcase` | `boolean` | No | Should this component be showcased |
| `start_date` | `string` | No | The date this component started being used |
| `status` | `string` | No | Status of component |
| `updated_at` | `string` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `page_access_group` | `/pages/{page_id}/components/{component_id}/page_access_groups` | `client.Component().create({ $action: 'page_access_group', ... })` |
| `page_access_user` | `/pages/{page_id}/components/{component_id}/page_access_users` | `client.Component().create({ $action: 'page_access_user', ... })` |
| `uptime` | `/pages/{page_id}/components/{component_id}/uptime` | `client.Component().load({ $action: 'uptime', ... })` |
| `page_access_group` | `/pages/{page_id}/components/{component_id}/page_access_groups` | `client.Component().remove({ $action: 'page_access_group', ... })` |
| `page_access_user` | `/pages/{page_id}/components/{component_id}/page_access_users` | `client.Component().remove({ $action: 'page_access_user', ... })` |

An action returns that action's OWN response, which is not necessarily a
Component record — check the API definition for its shape.

```ts
const result = await client.Component().create({
  $action: 'page_access_group',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Component().create({
  page_id: 'example_page_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Component().list({ page_id: "example" })
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
| `component_id` | `string` | No | Component identifier |
| `incidents` | `Record<string, any>` | No | Related incidents |

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
| `component_group` | `Record<string, any>` | Yes |  |
| `components` | `string` | No |  |
| `created_at` | `string` | No |  |
| `description` | `string` | No | Description of the component group. |
| `id` | `string` | No | Component Group Identifier |
| `name` | `string` | No |  |
| `page_id` | `string` | No |  |
| `position` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.GroupComponent().create({
  page_id: 'example_page_id',
  component_group: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.GroupComponent().list({ page_id: "example" })
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
| `auto_transition_deliver_notifications_at_end` | `boolean` | No | Controls whether send notification when scheduled maintenances auto transition to completed. |
| `auto_transition_deliver_notifications_at_start` | `boolean` | No | Controls whether send notification when scheduled maintenances auto transition to started. |
| `auto_transition_to_maintenance_state` | `boolean` | No | Controls whether change components status to under_maintenance once scheduled maintenance is in progress. |
| `auto_transition_to_operational_state` | `boolean` | No | Controls whether change components status to operational once scheduled maintenance completes. |
| `components` | `any[]` | No | Incident components |
| `created_at` | `string` | No | The timestamp when the incident was created at. |
| `id` | `string` | No | Incident Identifier |
| `impact` | `string` | No | The impact of the incident. |
| `impact_override` | `string` | No | value to override calculated impact value |
| `incident` | `Record<string, any>` | Yes |  |
| `incident_updates` | `any[]` | No | The incident updates for incident. |
| `metadata` | `Record<string, any>` | No | Metadata attached to the incident. |
| `monitoring_at` | `string` | No | The timestamp when incident entered monitoring state. |
| `name` | `string` | No | Incident Name. |
| `page_id` | `string` | No | Incident Page Identifier |
| `postmortem_body` | `string` | No | Body of the Postmortem. |
| `postmortem_body_last_updated_at` | `string` | No | The timestamp when the incident postmortem body was last updated at. |
| `postmortem_ignored` | `boolean` | No | Controls whether the incident will have postmortem. |
| `postmortem_notified_subscribers` | `boolean` | No | Indicates whether subscribers are already notificed about postmortem. |
| `postmortem_notified_twitter` | `boolean` | No | Controls whether to decide if notify postmortem on twitter. |
| `postmortem_published_at` | `boolean` | No | The timestamp when the postmortem was published. |
| `reminder_intervals` | `string` | No | Custom reminder intervals for unresolved/open incidents. |
| `resolved_at` | `string` | No | The timestamp when incident was resolved. |
| `scheduled_auto_completed` | `boolean` | No | Controls whether the incident is scheduled to automatically change to complete. |
| `scheduled_auto_in_progress` | `boolean` | No | Controls whether the incident is scheduled to automatically change to in progress. |
| `scheduled_for` | `string` | No | The timestamp the incident is scheduled for. |
| `scheduled_remind_prior` | `boolean` | No | Controls whether to remind subscribers prior to scheduled incidents. |
| `scheduled_reminded_at` | `string` | No | The timestamp when the scheduled incident reminder was sent at. |
| `scheduled_until` | `string` | No | The timestamp the incident is scheduled until. |
| `shortlink` | `string` | No | Incident Shortlink. |
| `status` | `string` | No | The incident status. |
| `updated_at` | `string` | No | The timestamp when the incident was updated at. |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `auto_transition_deliver_notifications_at_end` | - | - | - | - | - |
| `auto_transition_deliver_notifications_at_start` | - | - | - | - | - |
| `auto_transition_to_maintenance_state` | - | - | - | - | - |
| `auto_transition_to_operational_state` | - | - | - | - | - |
| `components` | - | - | - | - | - |
| `created_at` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `impact` | - | - | - | - | - |
| `impact_override` | - | - | - | - | - |
| `incident` | - | - | - | Yes | - |
| `incident_updates` | - | - | - | - | - |
| `metadata` | - | - | - | - | - |
| `monitoring_at` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `page_id` | - | - | - | - | - |
| `postmortem_body` | - | - | - | - | - |
| `postmortem_body_last_updated_at` | - | - | - | - | - |
| `postmortem_ignored` | - | - | - | - | - |
| `postmortem_notified_subscribers` | - | - | - | - | - |
| `postmortem_notified_twitter` | - | - | - | - | - |
| `postmortem_published_at` | - | - | - | - | - |
| `reminder_intervals` | - | - | - | - | - |
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

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `active_maintenance` | `/pages/{page_id}/incidents/active_maintenance` | `client.Incident().list({ $action: 'active_maintenance', ... })` |
| `scheduled` | `/pages/{page_id}/incidents/scheduled` | `client.Incident().list({ $action: 'scheduled', ... })` |
| `unresolved` | `/pages/{page_id}/incidents/unresolved` | `client.Incident().list({ $action: 'unresolved', ... })` |
| `upcoming` | `/pages/{page_id}/incidents/upcoming` | `client.Incident().list({ $action: 'upcoming', ... })` |

An action returns that action's OWN response, which is not necessarily a
Incident record — check the API definition for its shape.

```ts
const result = await client.Incident().list({
  $action: 'active_maintenance',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Incident().create({
  page_id: 'example_page_id',
  incident: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Incident().list({ page_id: "example" })
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
  incident_id: 'example_incident_id',
  page_id: 'example_page_id',
  subscriber_id: 'example_subscriber_id',
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
| `body` | `string` | No | Body of the incident or maintenance update to be applied when selecting this template |
| `components` | `any[]` | No | Affected components |
| `group_id` | `string` | No | Identifier of Template Group this template belongs to |
| `id` | `string` | No | Incident Template Identifier |
| `name` | `string` | No | Name of the template, as shown in the list on the "Templates" tab of the "Incidents" page |
| `should_send_notifications` | `boolean` | No | Whether the "deliver notifications" checkbox should be selected when selecting this template |
| `should_tweet` | `boolean` | No | Whether the "tweet update" checkbox should be selected when selecting this template |
| `template` | `Record<string, any>` | Yes |  |
| `title` | `string` | No | Title to be applied to the incident or maintenance when selecting this template |
| `update_status` | `string` | No | The status the incident or maintenance should transition to when selecting this template |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.IncidentTemplate().create({
  page_id: 'example_page_id',
  template: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.IncidentTemplate().list({ page_id: "example" })
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
| `affected_components` | `any[]` | No | Affected components associated with the incident update. |
| `body` | `string` | No | Incident update body. |
| `created_at` | `string` | No | The timestamp when the incident update was created at. |
| `custom_tweet` | `string` | No | An optional customized tweet message for incident postmortem. |
| `deliver_notifications` | `boolean` | No | Controls whether to delivery notifications. |
| `display_at` | `string` | No | Timestamp when incident update is happened. |
| `id` | `string` | No | Incident Update Identifier. |
| `incident_id` | `string` | No | Incident Identifier. |
| `incident_update` | `Record<string, any>` | No |  |
| `status` | `string` | No | The incident status. |
| `tweet_id` | `string` | No | Tweet identifier associated to this incident update. |
| `twitter_updated_at` | `string` | No | The timestamp when twitter updated at. |
| `updated_at` | `string` | No | The timestamp when the incident update is updated. |
| `wants_twitter_update` | `boolean` | No | Controls whether to create twitter update. |

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
| `data` | `Record<string, any>` | Yes | Add data points to metrics |
| `decimal_places` | `number` | No |  |
| `display` | `boolean` | No | Should the metric be displayed |
| `id` | `string` | No | Metric identifier |
| `last_fetched_at` | `string` | No |  |
| `metric` | `Record<string, any>` | No |  |
| `metric_identifier` | `string` | No | Metric Display identifier used to look up the metric data from the provider |
| `metrics_provider_id` | `string` | No | Metric Provider identifier |
| `most_recent_data_at` | `string` | No |  |
| `name` | `string` | No | Name of metric |
| `reference_name` | `string` | No |  |
| `suffix` | `string` | No | Suffix to describe the units on the graph |
| `tooltip_description` | `string` | No |  |
| `updated_at` | `string` | No |  |
| `y_axis_hidden` | `boolean` | No | Should the values on the y axis be hidden on render |
| `y_axis_max` | `number` | No |  |
| `y_axis_min` | `number` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `data` | `/pages/{page_id}/metrics/{metric_id}/data` | `client.Metric().create({ $action: 'data', ... })` |
| `data` | `/pages/{page_id}/metrics/data` | `client.Metric().create({ $action: 'data', ... })` |
| `data` | `/pages/{page_id}/metrics/{metric_id}/data` | `client.Metric().remove({ $action: 'data', ... })` |

An action returns that action's OWN response, which is not necessarily a
Metric record — check the API definition for its shape.

```ts
const result = await client.Metric().create({
  $action: 'data',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Metric().create({
  metrics_provider_id: 'example_metrics_provider_id',
  page_id: 'example_page_id',
  data: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Metric().list({ page_access_user_id: "example", page_id: "example" })
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
| `id` | `string` | No | Identifier for Metrics Provider |
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
  page_id: 'example_page_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.MetricsProvider().list({ page_id: "example" })
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
| `allow_email_subscribers` | `boolean` | No | Can your users choose to receive notifications via email |
| `allow_incident_subscribers` | `boolean` | No | Can your users subscribe to notifications for a single incident |
| `allow_page_subscribers` | `boolean` | No | Can your users subscribe to all notifications on the page |
| `allow_rss_atom_feeds` | `boolean` | No | Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages) |
| `allow_sms_subscribers` | `boolean` | No | Can your users choose to receive notifications via SMS |
| `allow_webhook_subscribers` | `boolean` | No | Can your users choose to receive notifications via Webhooks |
| `branding` | `string` | No | The main template your statuspage will use |
| `city` | `string` | No |  |
| `country` | `string` | No |  |
| `created_at` | `string` | No | Timestamp the record was created |
| `css_blues` | `string` | No | CSS Color |
| `css_body_background_color` | `string` | No | CSS Color |
| `css_border_color` | `string` | No | CSS Color |
| `css_font_color` | `string` | No | CSS Color |
| `css_graph_color` | `string` | No | CSS Color |
| `css_greens` | `string` | No | CSS Color |
| `css_light_font_color` | `string` | No | CSS Color |
| `css_link_color` | `string` | No | CSS Color |
| `css_no_data` | `string` | No | CSS Color |
| `css_oranges` | `string` | No | CSS Color |
| `css_reds` | `string` | No | CSS Color |
| `css_yellows` | `string` | No | CSS Color |
| `domain` | `string` | No | CNAME alias for your status page |
| `email_logo` | `string` | No |  |
| `favicon_logo` | `string` | No |  |
| `headline` | `string` | No |  |
| `hero_cover` | `string` | No |  |
| `hidden_from_search` | `boolean` | No | Should your page hide itself from search engines |
| `id` | `string` | No | Page identifier |
| `ip_restrictions` | `string` | No |  |
| `name` | `string` | No | Name of your page to be displayed |
| `notifications_email_footer` | `string` | No | Allows you to customize the footer appearing on your notification emails. |
| `notifications_from_email` | `string` | No | Allows you to customize the email address your page notifications come from |
| `page` | `Record<string, any>` | No |  |
| `page_description` | `string` | No |  |
| `state` | `string` | No |  |
| `subdomain` | `string` | No | Subdomain at which to access your status page |
| `support_url` | `string` | No |  |
| `time_zone` | `string` | No | Timezone configured for your page |
| `transactional_logo` | `string` | No |  |
| `twitter_logo` | `string` | No |  |
| `twitter_username` | `string` | No |  |
| `updated_at` | `string` | No | Timestamp the record was last updated |
| `url` | `string` | No | Website of your page. |
| `viewers_must_be_team_members` | `boolean` | No |  |

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
| `component_ids` | `any[]` | No | List of components codes to set on the page access group |
| `created_at` | `string` | No |  |
| `external_identifier` | `string` | No | Associates group with external group. |
| `id` | `string` | No | Page Access Group Identifier |
| `metric_ids` | `any[]` | No |  |
| `name` | `string` | No | Name for this Group. |
| `page_access_group` | `Record<string, any>` | No |  |
| `page_access_user_ids` | `any[]` | No |  |
| `page_id` | `string` | No | Page Identifier. |
| `updated_at` | `string` | No |  |

### Field Usage by Operation

| Field | load | list | create | update | remove |
| --- | --- | --- | --- | --- | --- |
| `component_ids` | - | - | Yes | - | - |
| `created_at` | - | - | - | - | - |
| `external_identifier` | - | - | - | - | - |
| `id` | - | - | - | - | - |
| `metric_ids` | - | - | - | - | - |
| `name` | - | - | - | - | - |
| `page_access_group` | - | - | - | - | - |
| `page_access_user_ids` | - | - | - | - | - |
| `page_id` | - | - | - | - | - |
| `updated_at` | - | - | - | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `component` | `/pages/{page_id}/page_access_groups/{page_access_group_id}/components` | `client.PageAccessGroup().create({ $action: 'component', ... })` |
| `component` | `/pages/{page_id}/page_access_groups/{page_access_group_id}/components` | `client.PageAccessGroup().patch({ $action: 'component', ... })` |
| `component` | `/pages/{page_id}/page_access_groups/{page_access_group_id}/components` | `client.PageAccessGroup().remove({ $action: 'component', ... })` |
| `component` | `/pages/{page_id}/page_access_groups/{page_access_group_id}/components` | `client.PageAccessGroup().update({ $action: 'component', ... })` |

An action returns that action's OWN response, which is not necessarily a
PageAccessGroup record — check the API definition for its shape.

```ts
const result = await client.PageAccessGroup().create({
  $action: 'component',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PageAccessGroup().create({
  id: 'example_id',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PageAccessGroup().list({ id: "example_id" })
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
| `component_ids` | `any[]` | Yes | List of component codes to allow access to |
| `created_at` | `string` | No |  |
| `email` | `string` | No |  |
| `external_login` | `string` | No | IDP login user id. |
| `id` | `string` | No | Page Access User Identifier |
| `metric_ids` | `any[]` | Yes | List of metrics to add |
| `page_access_group_id` | `string` | No |  |
| `page_access_group_ids` | `string` | No |  |
| `page_access_user` | `Record<string, any>` | No |  |
| `page_id` | `string` | No |  |
| `updated_at` | `string` | No |  |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `component` | `/pages/{page_id}/page_access_users/{page_access_user_id}/components` | `client.PageAccessUser().create({ $action: 'component', ... })` |
| `metric` | `/pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | `client.PageAccessUser().create({ $action: 'metric', ... })` |
| `component` | `/pages/{page_id}/page_access_users/{page_access_user_id}/components` | `client.PageAccessUser().patch({ $action: 'component', ... })` |
| `metric` | `/pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | `client.PageAccessUser().patch({ $action: 'metric', ... })` |
| `component` | `/pages/{page_id}/page_access_users/{page_access_user_id}/components` | `client.PageAccessUser().remove({ $action: 'component', ... })` |
| `metric` | `/pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | `client.PageAccessUser().remove({ $action: 'metric', ... })` |
| `component` | `/pages/{page_id}/page_access_users/{page_access_user_id}/components` | `client.PageAccessUser().update({ $action: 'component', ... })` |
| `metric` | `/pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | `client.PageAccessUser().update({ $action: 'metric', ... })` |

An action returns that action's OWN response, which is not necessarily a
PageAccessUser record — check the API definition for its shape.

```ts
const result = await client.PageAccessUser().create({
  $action: 'component',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.PageAccessUser().create({
  id: 'example_id',
  component_ids: [],
  metric_ids: [],
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.PageAccessUser().list({ id: "example_id" })
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
| `pages` | `Record<string, any>` | No | Pages accessible by the user. |
| `user_id` | `string` | No | User identifier |

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
| `body` | `string` | No | Postmortem body |
| `body_draft` | `string` | No | Body draft |
| `body_draft_updated_at` | `string` | No |  |
| `body_updated_at` | `string` | No |  |
| `created_at` | `string` | No |  |
| `custom_tweet` | `string` | No | Custom tweet for Incident Postmortem |
| `notify_subscribers` | `boolean` | No | Should email subscribers be notified. |
| `notify_twitter` | `boolean` | No | Should Twitter followers be notified. |
| `postmortem` | `Record<string, any>` | Yes |  |
| `preview_key` | `string` | No | Preview Key |
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
| `notify_subscribers` | - | - |
| `notify_twitter` | - | - |
| `postmortem` | - | Yes |
| `preview_key` | - | - |
| `published_at` | - | - |
| `updated_at` | - | - |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `publish` | `/pages/{page_id}/incidents/{incident_id}/postmortem/publish` | `client.Postmortem().update({ $action: 'publish', ... })` |
| `revert` | `/pages/{page_id}/incidents/{incident_id}/postmortem/revert` | `client.Postmortem().update({ $action: 'revert', ... })` |

An action returns that action's OWN response, which is not necessarily a
Postmortem record — check the API definition for its shape.

```ts
const result = await client.Postmortem().update({
  $action: 'publish',
  /* ...the action's own arguments */
})
```

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
| `incident_background_color` | `string` | No | Color of status embed iframe background when displaying incident |
| `incident_text_color` | `string` | No | Color of status embed iframe text when displaying incident |
| `maintenance_background_color` | `string` | No | Color of status embed iframe background when displaying maintenance |
| `maintenance_text_color` | `string` | No | Color of status embed iframe text when displaying maintenance |
| `page_id` | `string` | No | Page identifier |
| `position` | `string` | No | Corner where status embed iframe will appear on page |
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
| `component_ids` | `any[]` | No | A list of component ids for which the subscriber should recieve updates for. |
| `components` | `string` | No | The components for which the subscriber has elected to receive updates. |
| `created_at` | `string` | No |  |
| `display_phone_number` | `string` | No | A formatted version of the phone_number and phone_country pair, nicely formatted for display. |
| `email` | `string` | No | The email address to use to contact the subscriber. |
| `endpoint` | `string` | No | The URL where a webhook subscriber elects to receive updates. |
| `id` | `string` | No | Subscriber Identifier |
| `integration_partner` | `number` | No | The number of integration partners found by the query. |
| `mode` | `string` | No | The communication mode of the subscriber. |
| `obfuscated_channel_name` | `string` | No | Obfuscated slack channel name |
| `page_access_user_id` | `string` | No | The Page Access user this subscriber belongs to (only for audience-specific pages). |
| `phone_country` | `string` | No | The two-character country code representing the country of which the phone_number is a part. |
| `phone_number` | `string` | No | The phone number used to contact an SMS subscriber |
| `purge_at` | `string` | No | The timestamp when a quarantined subscriber will be purged (unsubscribed). |
| `quarantined_at` | `string` | No | The timestamp when the subscriber was quarantined due to an issue reaching them. |
| `skip_confirmation_notification` | `boolean` | No | If this is true, do not notify the user with changes to their subscription. |
| `skip_unsubscription_notification` | `boolean` | No | If skip_unsubscription_notification is true, the subscribers do not receive any notifications when they are unsubscribed. |
| `slack` | `number` | No | The number of Slack subscribers found by the query. |
| `sms` | `number` | No | The number of Webhook subscribers found by the query. |
| `state` | `string` | No | If this is present, only unsubscribe subscribers in this state. |
| `subscriber` | `Record<string, any>` | No |  |
| `subscribers` | `string` | Yes | The array of quarantined subscriber codes to reactivate, or "all" to reactivate all quarantined subscribers. |
| `teams` | `number` | No | The number of MS teams subscribers found by the query. |
| `type` | `string` | No | If this is present, only reactivate subscribers of this type. |
| `webhook` | `number` | No | The number of SMS subscribers found by the query. |
| `workspace_name` | `string` | No | The workspace name of the slack subscriber. |

### Actions

This entity exposes custom API actions in addition to the standard
operations. Select one with `$action` in the call's argument; the
remaining keys are sent as that action's payload.

| Action | Route | Call |
| --- | --- | --- |
| `reactivate` | `/pages/{page_id}/subscribers/reactivate` | `client.Subscriber().create({ $action: 'reactivate', ... })` |
| `resend_confirmation` | `/pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation` | `client.Subscriber().create({ $action: 'resend_confirmation', ... })` |
| `resend_confirmation` | `/pages/{page_id}/subscribers/resend_confirmation` | `client.Subscriber().create({ $action: 'resend_confirmation', ... })` |
| `unsubscribe` | `/pages/{page_id}/subscribers/unsubscribe` | `client.Subscriber().create({ $action: 'unsubscribe', ... })` |
| `unsubscribed` | `/pages/{page_id}/subscribers/unsubscribed` | `client.Subscriber().list({ $action: 'unsubscribed', ... })` |
| `count` | `/pages/{page_id}/subscribers/count` | `client.Subscriber().load({ $action: 'count', ... })` |
| `histogram_by_state` | `/pages/{page_id}/subscribers/histogram_by_state` | `client.Subscriber().load({ $action: 'histogram_by_state', ... })` |

An action returns that action's OWN response, which is not necessarily a
Subscriber record — check the API definition for its shape.

```ts
const result = await client.Subscriber().create({
  $action: 'reactivate',
  /* ...the action's own arguments */
})
```

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.Subscriber().create({
  page_id: 'example_page_id',
  subscribers: 'example_subscribers',
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.Subscriber().list({ page_id: "example" })
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
| `email` | `string` | No | Email address for the team member |
| `first_name` | `string` | No |  |
| `id` | `string` | No | User identifier |
| `last_name` | `string` | No |  |
| `organization_id` | `string` | No | Organization identifier |
| `updated_at` | `string` | No |  |
| `user` | `Record<string, any>` | Yes |  |

### Operations

#### `create(data: object, ctrl?: object)`

Create a new entity with the given data.

```ts
const result = await client.User().create({
  organization_id: 'example_organization_id',
  user: {},
})
```

#### `list(match: object, ctrl?: object)`

List entities matching the given criteria. Returns an array.

```ts
const results = await client.User().list({ organization_id: "example" })
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

