# Statuspage API

# Code of Conduct Please don&#39;t abuse the API, and please report all feature requests and issues to https://support.atlassian.com/contact # Rate Limiting Each API token is limited to 1 request / second as measured on a 60 second rolling window. To get this limit increased, please contact us at https://support.atlassian.com/contact Error codes 420 or 429 indicate that you have exceeded the rate limit and the request has been rejected. # Basics ## HTTPS It&#39;s required ## URL Prefix In order to maintain version integrity into the future, the API is versioned. All calls currently begin with the following prefix: https://api.statuspage.io/v1/ ## RESTful Interface Wherever possible, the API seeks to implement repeatable patterns with logical, representative URLs and descriptive HTTP verbs. Below are some examples and conventions you will see throughout the documentation. * Collections are buckets: https://api.statuspage.io/v1/pages/asdf123/incidents.json * Elements have unique IDs: https://api.statuspage.io/v1/pages/asdf123/incidents/jklm456.json * GET will retrieve information about a collection/element * POST will create an element in a collection * PATCH will update a single element * PUT will replace a single element in a collection (rarely used) * DELETE will destroy a single element ## Sending Data Information can be sent in the body as form urlencoded or JSON, but make sure the Content-Type header matches the body structure or the server gremlins will be angry. All examples are provided in JSON format, however they can easily be converted to form encoding if required. Some examples of how to convert things are below: // JSON &#123; &quot;incident&quot;: &#123; &quot;name&quot;: &quot;test incident&quot;, &quot;components&quot;: [&quot;8kbf7d35c070&quot;, &quot;vtnh60py4yd7&quot;] &#125; &#125; // Form Encoded (using curl as an example): curl -X POST https://api.statuspage.io/v1/example \ -d &quot;incident[name]=test incident&quot; \ -d &quot;incident[components][]=8kbf7d35c070&quot; \ -d &quot;incident[components][]=vtnh60py4yd7&quot; # Authentication &lt;!-- ReDoc-Inject: &lt;security-definitions&gt; --&gt;

## Start here

This guide introduces the API, the client libraries, and the companion tools in this repository. Start with the API capabilities, choose a client for your application, and use the linked reference when you need exact request and response details.

The selected API surface contains 18 entities and 112 HTTP routes. There are 6 SDK targets and 2 companion tools.

An entity groups related API operations. An operation can have several routes with different inputs or authentication requirements. The SDK exposes the entity and its operations using the conventions of the selected language.

## What the API provides

### [Component](docs/api/component.html)

Results: Add page access groups to a component; Add page access users to a component; Create a component; List components for a page access group; Get components for page access user; Get a list of components; Get uptime data for a component that has uptime showcase enabled; Get a component; Update a component; Delete a component; Remove page access groups from a component; Remove page access users from component.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `automation_email`: Requires a special feature flag to be enabled
- `description`: More detailed description for component
- `group`: Is this component a group
- `group_id`: Component Group identifier
- `id`: Identifier for component

### [ComponentGroupUptime](docs/api/component_group_uptime.html)

Results: Get uptime data for a component group that has uptime showcase enabled for at least one component.

SDK operations: `load`.

Key fields to recognise:

- `component_id`: Component identifier
- `incidents`: Related incidents

### [GroupComponent](docs/api/group_component.html)

Results: Create a component group; Get a list of component groups; Get a component group; Update a component group; Delete a component group.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `description`: Description of the component group.
- `id`: Component Group Identifier

### [Incident](docs/api/incident.html)

Results: Create an incident; Get a list of incidents; Get a list of active maintenances; Get a list of scheduled incidents; Get a list of unresolved incidents; Get a list of upcoming incidents; Get an incident; Update an incident; Delete an incident.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `auto_transition_deliver_notifications_at_end`: Controls whether send notification when scheduled maintenances auto transition to completed.
- `auto_transition_deliver_notifications_at_start`: Controls whether send notification when scheduled maintenances auto transition to started.
- `auto_transition_to_maintenance_state`: Controls whether change components status to under_maintenance once scheduled maintenance is in progress.
- `auto_transition_to_operational_state`: Controls whether change components status to operational once scheduled maintenance completes.
- `components`: Incident components

### [IncidentPostmortem](docs/api/incident_postmortem.html)

Results: Delete Postmortem.

SDK operations: `remove`.

### [IncidentSubscriber](docs/api/incident_subscriber.html)

Results: Resend confirmation to an incident subscriber.

SDK operations: `create`.

### [IncidentTemplate](docs/api/incident_template.html)

Results: Create a template; Get a list of templates.

SDK operations: `create`, `list`.

Key fields to recognise:

- `body`: Body of the incident or maintenance update to be applied when selecting this template
- `components`: Affected components
- `group_id`: Identifier of Template Group this template belongs to
- `id`: Incident Template Identifier
- `name`: Name of the template, as shown in the list on the &quot;Templates&quot; tab of the &quot;Incidents&quot; page

### [IncidentUpdate](docs/api/incident_update.html)

Results: Update a previous incident update.

SDK operations: `patch`, `update`.

Key fields to recognise:

- `affected_components`: Affected components associated with the incident update.
- `body`: Incident update body.
- `created_at`: The timestamp when the incident update was created at.
- `custom_tweet`: An optional customized tweet message for incident postmortem.
- `deliver_notifications`: Controls whether to delivery notifications.

### [Metric](docs/api/metric.html)

Results: Add data to a metric; Create a metric for a metric provider; Data Point is submitted and is currently being added to the metrics; Get metrics for page access user; List metrics for a metric provider; Get a list of metrics; Get a metric; Update a metric; Delete a metric; Reset data for a metric.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `display`: Should the metric be displayed
- `id`: Metric identifier
- `metric_identifier`: Metric Display identifier used to look up the metric data from the provider
- `metrics_provider_id`: Metric Provider identifier
- `name`: Name of metric

### [MetricsProvider](docs/api/metrics_provider.html)

Results: Create a metric provider; Get a list of metric providers; Get a metric provider; Update a metric provider; Delete a metric provider.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `id`: Identifier for Metrics Provider

### [Page](docs/api/page.html)

Results: Get a list of pages; Get a page; Update a page.

SDK operations: `list`, `load`, `patch`, `update`.

Key fields to recognise:

- `allow_email_subscribers`: Can your users choose to receive notifications via email
- `allow_incident_subscribers`: Can your users subscribe to notifications for a single incident
- `allow_page_subscribers`: Can your users subscribe to all notifications on the page
- `allow_rss_atom_feeds`: Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)
- `allow_sms_subscribers`: Can your users choose to receive notifications via SMS

### [PageAccessGroup](docs/api/page_access_group.html)

Results: Replace components for a page access group; Create a page access group; Get a list of page access groups; Get a page access group; Update a page access group; Add components to page access group; Remove a component from a page access group; Remove a page access group; Delete components for a page access group.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `external_identifier`: Associates group with external group.
- `id`: Page Access Group Identifier
- `name`: Name for this Group.
- `page_id`: Page Identifier.

### [PageAccessUser](docs/api/page_access_user.html)

Results: Replace components for page access user; Replace metrics for page access user; Add a page access user; Get a list of page access users; Get page access user; Update page access user; Add components for page access user; Add metrics for page access user; Remove component for page access user; Delete metric for page access user; Delete page access user; Remove components for page access user; Delete metrics for page access user.

SDK operations: `create`, `list`, `load`, `patch`, `remove`, `update`.

Key fields to recognise:

- `external_login`: IDP login user id. Key is typically &quot;uid&quot;.
- `id`: Page Access User Identifier

### [Permission](docs/api/permission.html)

Results: Get a user&#39;s permissions; Update a user&#39;s role permissions. Payload should contain a mapping of pages to a set of the desired roles, if the page has Role Based Access Control. Otherwise, the pages should map to an empty hash. User will lose access to any pages omitted from the payload.

SDK operations: `load`, `update`.

Key fields to recognise:

- `pages`: Pages accessible by the user.
- `user_id`: User identifier

### [Postmortem](docs/api/postmortem.html)

Results: Get Postmortem; Create Postmortem; Publish Postmortem; Revert Postmortem.

SDK operations: `load`, `update`.

Key fields to recognise:

- `body`: Postmortem body
- `body_draft`: Body draft
- `custom_tweet`: Custom tweet for Incident Postmortem
- `notify_subscribers`: Should email subscribers be notified.
- `notify_twitter`: Should Twitter followers be notified.

### [StatusEmbedConfig](docs/api/status_embed_config.html)

Results: Get status embed config settings; Update status embed config settings.

SDK operations: `load`, `patch`, `update`.

Key fields to recognise:

- `incident_background_color`: Color of status embed iframe background when displaying incident
- `incident_text_color`: Color of status embed iframe text when displaying incident
- `maintenance_background_color`: Color of status embed iframe background when displaying maintenance
- `maintenance_text_color`: Color of status embed iframe text when displaying maintenance
- `page_id`: Page identifier

### [Subscriber](docs/api/subscriber.html)

Results: Resend confirmation to a subscriber; Create an incident subscriber; Create a subscriber. Not applicable for Slack subscribers.; Reactivate a list of quarantined subscribers; Resend confirmations to a list of subscribers; Unsubscribe a list of subscribers; Get a list of subscribers; Get a list of incident subscribers; Get a list of unsubscribed subscribers; Get an incident subscriber; Get a count of subscribers by type; Get a subscriber; Get a histogram of subscribers by type and then state; Unsubscribe an incident subscriber; Unsubscribe a subscriber; Update a subscriber.

SDK operations: `create`, `list`, `load`, `remove`, `update`.

Key fields to recognise:

- `component_ids`: A list of component ids for which the subscriber should recieve updates for.
- `components`: The components for which the subscriber has elected to receive updates.
- `display_phone_number`: A formatted version of the phone_number and phone_country pair, nicely formatted for display.
- `email`: The email address to use to contact the subscriber. Used for Email and Webhook subscribers.
- `endpoint`: The URL where a webhook subscriber elects to receive updates.

### [User](docs/api/user.html)

Results: Create a user; Get a list of users; Delete a user.

SDK operations: `create`, `list`, `remove`.

Key fields to recognise:

- `email`: Email address for the team member
- `id`: User identifier
- `organization_id`: Organization identifier

### Route map

Use this map to locate a capability. Consult the entity reference before supplying request data; routes for the same operation can require different fields.

| Entity | SDK operation | HTTP route | Authentication |
| --- | --- | --- | --- |
| [Component](docs/api/component.html) | `create` | `POST /pages/{page_id}/components/{component_id}/page_access_groups` | Required |
| [Component](docs/api/component.html) | `create` | `POST /pages/{page_id}/components/{component_id}/page_access_users` | Required |
| [Component](docs/api/component.html) | `create` | `POST /pages/{page_id}/components` | Required |
| [Component](docs/api/component.html) | `list` | `GET /pages/{page_id}/page_access_groups/{page_access_group_id}/components` | Required |
| [Component](docs/api/component.html) | `list` | `GET /pages/{page_id}/page_access_users/{page_access_user_id}/components` | Required |
| [Component](docs/api/component.html) | `list` | `GET /pages/{page_id}/components` | Required |
| [Component](docs/api/component.html) | `load` | `GET /pages/{page_id}/components/{component_id}/uptime` | Required |
| [Component](docs/api/component.html) | `load` | `GET /pages/{page_id}/components/{component_id}` | Required |
| [Component](docs/api/component.html) | `patch` | `PATCH /pages/{page_id}/components/{component_id}` | Required |
| [Component](docs/api/component.html) | `remove` | `DELETE /pages/{page_id}/components/{component_id}` | Required |
| [Component](docs/api/component.html) | `remove` | `DELETE /pages/{page_id}/components/{component_id}/page_access_groups` | Required |
| [Component](docs/api/component.html) | `remove` | `DELETE /pages/{page_id}/components/{component_id}/page_access_users` | Required |
| [Component](docs/api/component.html) | `update` | `PUT /pages/{page_id}/components/{component_id}` | Required |
| [ComponentGroupUptime](docs/api/component_group_uptime.html) | `load` | `GET /pages/{page_id}/component-groups/{id}/uptime` | Required |
| [GroupComponent](docs/api/group_component.html) | `create` | `POST /pages/{page_id}/component-groups` | Required |
| [GroupComponent](docs/api/group_component.html) | `list` | `GET /pages/{page_id}/component-groups` | Required |
| [GroupComponent](docs/api/group_component.html) | `load` | `GET /pages/{page_id}/component-groups/{id}` | Required |
| [GroupComponent](docs/api/group_component.html) | `patch` | `PATCH /pages/{page_id}/component-groups/{id}` | Required |
| [GroupComponent](docs/api/group_component.html) | `remove` | `DELETE /pages/{page_id}/component-groups/{id}` | Required |
| [GroupComponent](docs/api/group_component.html) | `update` | `PUT /pages/{page_id}/component-groups/{id}` | Required |
| [Incident](docs/api/incident.html) | `create` | `POST /pages/{page_id}/incidents` | Required |
| [Incident](docs/api/incident.html) | `list` | `GET /pages/{page_id}/incidents` | Required |
| [Incident](docs/api/incident.html) | `list` | `GET /pages/{page_id}/incidents/active_maintenance` | Required |
| [Incident](docs/api/incident.html) | `list` | `GET /pages/{page_id}/incidents/scheduled` | Required |
| [Incident](docs/api/incident.html) | `list` | `GET /pages/{page_id}/incidents/unresolved` | Required |
| [Incident](docs/api/incident.html) | `list` | `GET /pages/{page_id}/incidents/upcoming` | Required |
| [Incident](docs/api/incident.html) | `load` | `GET /pages/{page_id}/incidents/{incident_id}` | Required |
| [Incident](docs/api/incident.html) | `patch` | `PATCH /pages/{page_id}/incidents/{incident_id}` | Required |
| [Incident](docs/api/incident.html) | `remove` | `DELETE /pages/{page_id}/incidents/{incident_id}` | Required |
| [Incident](docs/api/incident.html) | `update` | `PUT /pages/{page_id}/incidents/{incident_id}` | Required |
| [IncidentPostmortem](docs/api/incident_postmortem.html) | `remove` | `DELETE /pages/{page_id}/incidents/{incident_id}/postmortem` | Required |
| [IncidentSubscriber](docs/api/incident_subscriber.html) | `create` | `POST /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}/resend_confirmation` | Required |
| [IncidentTemplate](docs/api/incident_template.html) | `create` | `POST /pages/{page_id}/incident_templates` | Required |
| [IncidentTemplate](docs/api/incident_template.html) | `list` | `GET /pages/{page_id}/incident_templates` | Required |
| [IncidentUpdate](docs/api/incident_update.html) | `patch` | `PATCH /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}` | Required |
| [IncidentUpdate](docs/api/incident_update.html) | `update` | `PUT /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}` | Required |
| [Metric](docs/api/metric.html) | `create` | `POST /pages/{page_id}/metrics/{metric_id}/data` | Required |
| [Metric](docs/api/metric.html) | `create` | `POST /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics` | Required |
| [Metric](docs/api/metric.html) | `create` | `POST /pages/{page_id}/metrics/data` | Required |
| [Metric](docs/api/metric.html) | `list` | `GET /pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | Required |
| [Metric](docs/api/metric.html) | `load` | `GET /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics` | Required |
| [Metric](docs/api/metric.html) | `load` | `GET /pages/{page_id}/metrics` | Required |
| [Metric](docs/api/metric.html) | `load` | `GET /pages/{page_id}/metrics/{metric_id}` | Required |
| [Metric](docs/api/metric.html) | `patch` | `PATCH /pages/{page_id}/metrics/{metric_id}` | Required |
| [Metric](docs/api/metric.html) | `remove` | `DELETE /pages/{page_id}/metrics/{metric_id}` | Required |
| [Metric](docs/api/metric.html) | `remove` | `DELETE /pages/{page_id}/metrics/{metric_id}/data` | Required |
| [Metric](docs/api/metric.html) | `update` | `PUT /pages/{page_id}/metrics/{metric_id}` | Required |
| [MetricsProvider](docs/api/metrics_provider.html) | `create` | `POST /pages/{page_id}/metrics_providers` | Required |
| [MetricsProvider](docs/api/metrics_provider.html) | `list` | `GET /pages/{page_id}/metrics_providers` | Required |
| [MetricsProvider](docs/api/metrics_provider.html) | `load` | `GET /pages/{page_id}/metrics_providers/{metrics_provider_id}` | Required |
| [MetricsProvider](docs/api/metrics_provider.html) | `patch` | `PATCH /pages/{page_id}/metrics_providers/{metrics_provider_id}` | Required |
| [MetricsProvider](docs/api/metrics_provider.html) | `remove` | `DELETE /pages/{page_id}/metrics_providers/{metrics_provider_id}` | Required |
| [MetricsProvider](docs/api/metrics_provider.html) | `update` | `PUT /pages/{page_id}/metrics_providers/{metrics_provider_id}` | Required |
| [Page](docs/api/page.html) | `list` | `GET /pages` | Required |
| [Page](docs/api/page.html) | `load` | `GET /pages/{page_id}` | Required |
| [Page](docs/api/page.html) | `patch` | `PATCH /pages/{page_id}` | Required |
| [Page](docs/api/page.html) | `update` | `PUT /pages/{page_id}` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `create` | `POST /pages/{page_id}/page_access_groups/{page_access_group_id}/components` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `create` | `POST /pages/{page_id}/page_access_groups` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `list` | `GET /pages/{page_id}/page_access_groups` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `load` | `GET /pages/{page_id}/page_access_groups/{page_access_group_id}` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `patch` | `PATCH /pages/{page_id}/page_access_groups/{page_access_group_id}` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `patch` | `PATCH /pages/{page_id}/page_access_groups/{page_access_group_id}/components` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `remove` | `DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `remove` | `DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `remove` | `DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}/components` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `update` | `PUT /pages/{page_id}/page_access_groups/{page_access_group_id}` | Required |
| [PageAccessGroup](docs/api/page_access_group.html) | `update` | `PUT /pages/{page_id}/page_access_groups/{page_access_group_id}/components` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `create` | `POST /pages/{page_id}/page_access_users/{page_access_user_id}/components` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `create` | `POST /pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `create` | `POST /pages/{page_id}/page_access_users` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `list` | `GET /pages/{page_id}/page_access_users` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `load` | `GET /pages/{page_id}/page_access_users/{page_access_user_id}` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `patch` | `PATCH /pages/{page_id}/page_access_users/{page_access_user_id}` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `patch` | `PATCH /pages/{page_id}/page_access_users/{page_access_user_id}/components` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `patch` | `PATCH /pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `remove` | `DELETE /pages/{page_id}/page_access_users/{page_access_user_id}/components/{component_id}` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `remove` | `DELETE /pages/{page_id}/page_access_users/{page_access_user_id}/metrics/{metric_id}` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `remove` | `DELETE /pages/{page_id}/page_access_users/{page_access_user_id}` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `remove` | `DELETE /pages/{page_id}/page_access_users/{page_access_user_id}/components` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `remove` | `DELETE /pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `update` | `PUT /pages/{page_id}/page_access_users/{page_access_user_id}` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `update` | `PUT /pages/{page_id}/page_access_users/{page_access_user_id}/components` | Required |
| [PageAccessUser](docs/api/page_access_user.html) | `update` | `PUT /pages/{page_id}/page_access_users/{page_access_user_id}/metrics` | Required |
| [Permission](docs/api/permission.html) | `load` | `GET /organizations/{organization_id}/permissions/{user_id}` | Required |
| [Permission](docs/api/permission.html) | `update` | `PUT /organizations/{organization_id}/permissions/{user_id}` | Required |
| [Postmortem](docs/api/postmortem.html) | `load` | `GET /pages/{page_id}/incidents/{incident_id}/postmortem` | Required |
| [Postmortem](docs/api/postmortem.html) | `update` | `PUT /pages/{page_id}/incidents/{incident_id}/postmortem` | Required |
| [Postmortem](docs/api/postmortem.html) | `update` | `PUT /pages/{page_id}/incidents/{incident_id}/postmortem/publish` | Required |
| [Postmortem](docs/api/postmortem.html) | `update` | `PUT /pages/{page_id}/incidents/{incident_id}/postmortem/revert` | Required |
| [StatusEmbedConfig](docs/api/status_embed_config.html) | `load` | `GET /pages/{page_id}/status_embed_config` | Required |
| [StatusEmbedConfig](docs/api/status_embed_config.html) | `patch` | `PATCH /pages/{page_id}/status_embed_config` | Required |
| [StatusEmbedConfig](docs/api/status_embed_config.html) | `update` | `PUT /pages/{page_id}/status_embed_config` | Required |
| [Subscriber](docs/api/subscriber.html) | `create` | `POST /pages/{page_id}/subscribers/{subscriber_id}/resend_confirmation` | Required |
| [Subscriber](docs/api/subscriber.html) | `create` | `POST /pages/{page_id}/incidents/{incident_id}/subscribers` | Required |
| [Subscriber](docs/api/subscriber.html) | `create` | `POST /pages/{page_id}/subscribers` | Required |
| [Subscriber](docs/api/subscriber.html) | `create` | `POST /pages/{page_id}/subscribers/reactivate` | Required |
| [Subscriber](docs/api/subscriber.html) | `create` | `POST /pages/{page_id}/subscribers/resend_confirmation` | Required |
| [Subscriber](docs/api/subscriber.html) | `create` | `POST /pages/{page_id}/subscribers/unsubscribe` | Required |
| [Subscriber](docs/api/subscriber.html) | `list` | `GET /pages/{page_id}/subscribers` | Required |
| [Subscriber](docs/api/subscriber.html) | `list` | `GET /pages/{page_id}/incidents/{incident_id}/subscribers` | Required |
| [Subscriber](docs/api/subscriber.html) | `list` | `GET /pages/{page_id}/subscribers/unsubscribed` | Required |
| [Subscriber](docs/api/subscriber.html) | `load` | `GET /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}` | Required |
| [Subscriber](docs/api/subscriber.html) | `load` | `GET /pages/{page_id}/subscribers/count` | Required |
| [Subscriber](docs/api/subscriber.html) | `load` | `GET /pages/{page_id}/subscribers/{subscriber_id}` | Required |
| [Subscriber](docs/api/subscriber.html) | `load` | `GET /pages/{page_id}/subscribers/histogram_by_state` | Required |
| [Subscriber](docs/api/subscriber.html) | `remove` | `DELETE /pages/{page_id}/incidents/{incident_id}/subscribers/{subscriber_id}` | Required |
| [Subscriber](docs/api/subscriber.html) | `remove` | `DELETE /pages/{page_id}/subscribers/{subscriber_id}` | Required |
| [Subscriber](docs/api/subscriber.html) | `update` | `PATCH /pages/{page_id}/subscribers/{subscriber_id}` | Required |
| [User](docs/api/user.html) | `create` | `POST /organizations/{organization_id}/users` | Required |
| [User](docs/api/user.html) | `list` | `GET /organizations/{organization_id}/users` | Required |
| [User](docs/api/user.html) | `remove` | `DELETE /organizations/{organization_id}/users/{user_id}` | Required |

## Connect to the API

- API server: `https://api.statuspage.io/v1`

The default credential is sent in the `Authorization` header with the `OAuth` prefix.

#### Obtaining your API Key Authentication is done via an API token provided in the Statuspage management interface. 1. Log in to your account at https://manage.statuspage.io/login. 2. Click on your avatar in the bottom left of your screen to access the user menu. 3. Click **API info**. ### Passing your API key in an authorization header The following example authenticates you with the Statuspage API. Along with the Page ID listed on the API page, we can fetch your page profile. curl -H &quot;Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1&quot; \ https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json ### Passing your API key in a query param curl &quot;https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1&quot;

Check authentication for the route you plan to call. A route that declares no authentication can be used without credentials; this does not change the requirements of other routes. Keep credentials in environment variables or a configured secret provider, and keep them out of source control and logs.

## Make a first request

1. Choose the API server and an operation that matches your task.
2. Check the operation’s required input and authentication. Use values valid for your account and environment.
3. Send one request and inspect the returned data before adding retries, concurrency, or a larger batch.

For an SDK call, install or build the chosen client, create a client instance with its documented configuration, and call the required entity operation. Language references describe the argument shape, asynchronous behaviour, and returned values.

## Choose an SDK

Choose the language already used by your application or service. The clients represent the same API model, while package setup, naming, and return types follow each language. Check the selected client’s reference and tests before integrating it into an existing application.

| Client | Repository directory | Distribution |
| --- | --- | --- |
| [Golang](docs/sdks/go.html) | `go/` | Build from source |
| [Lua](docs/sdks/lua.html) | `lua/` | Build from source |
| [PHP](docs/sdks/php.html) | `php/` | Build from source |
| [Python](docs/sdks/py.html) | `py/` | Build from source |
| [Ruby](docs/sdks/rb.html) | `rb/` | Build from source |
| [TypeScript](docs/sdks/ts.html) | `ts/` | Build from source |

Build-from-source entries are not marked as published in the project model. Follow the build instructions in that target’s README, then consume the resulting package using your language’s local dependency mechanism. Published entries give the installation command recorded for that client.

## Companion tools

These targets provide another way to use the API. Their available commands or tools can cover a smaller set of operations than the client libraries.

### [Go CLI](docs/tools/go-cli.html)

Use the command-line interface for shell-based tasks and scripts.

Repository directory: `go-cli/`. Not published. Build from the go-cli directory.


### [Go MCP server](docs/tools/go-mcp.html)

Use the MCP server to expose supported API operations to an MCP client.

Repository directory: `go-mcp/`. Not published. Build from the go-mcp directory.

- `statuspage_list`: List records for an entity. Supported entities: `component`, `group_component`, `incident`, `incident_template`, `metric`, `metrics_provider`, `page`, `page_access_group`, `page_access_user`, `subscriber`, `user`.
- `statuspage_load`: Load one record for an entity. Supported entities: `component`, `component_group_uptime`, `group_component`, `incident`, `metric`, `metrics_provider`, `page`, `page_access_group`, `page_access_user`, `permission`, `postmortem`, `status_embed_config`, `subscriber`.

## Operational features

Features supply behaviour around API calls, such as request handling, diagnostics, or local testing. Inclusion in this project does not mean a feature is enabled at runtime. Check the selected SDK’s supported features and configuration defaults, then enable the behaviour your application needs.

- [`debug`](docs/features/debug.html): Request/response capture ring buffer for debugging
- [`idempotency`](docs/features/idempotency.html): Idempotency keys for safe retries of mutating operations
- [`metrics`](docs/features/metrics.html): Statistics capture: per-operation counters and latency
- [`paging`](docs/features/paging.html): Pagination signals for list operations
- [`ratelimit`](docs/features/ratelimit.html): Client-side rate limiting via a token bucket
- [`retry`](docs/features/retry.html): Automatic retry of transient failures with exponential backoff
- [`test`](docs/features/test.html): In-memory mock transport for testing without a live server
- [`timeout`](docs/features/timeout.html): Per-request timeout with transport abort

Start with the default client configuration. Add request limits and diagnostics as needed, test error paths, and review retry behaviour before using operations that change data. A retry can repeat an operation unless the API provides a suitable guarantee.

## Continue with the documentation

- Follow the [first-call guide](docs/guides/first-call.html) for the setup sequence.
- Read the [authentication guide](docs/guides/authentication.html) before using protected routes.
- Use the [API reference](docs/api/index.html) for request schemas, response formats, and status codes.
- Check the chosen SDK or companion tool reference for its configuration and supported operations.

