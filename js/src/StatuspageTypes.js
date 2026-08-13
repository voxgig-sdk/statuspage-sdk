// Typed models for the Statuspage SDK (JSDoc typedefs).
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
// edit by hand.

/**
 * @typedef {Object} Component
 * @property {string} [automation_email]
 * @property {Object} [component]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {boolean} [group]
 * @property {string} [group_id]
 * @property {string} [id]
 * @property {string} [name]
 * @property {boolean} [only_show_if_degraded]
 * @property {string} [page_id]
 * @property {number} [position]
 * @property {boolean} [showcase]
 * @property {string} [start_date]
 * @property {string} [status]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} ComponentLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} ComponentListMatch
 * @property {string} [page_access_group_id]
 * @property {string} page_id
 * @property {string} [page_access_user_id]
 */

/**
 * @typedef {Object} ComponentCreateData
 * @property {string} page_id
 * @property {string} [automation_email]
 * @property {Object} [component]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {boolean} [group]
 * @property {string} [group_id]
 * @property {string} [id]
 * @property {string} [name]
 * @property {boolean} [only_show_if_degraded]
 * @property {number} [position]
 * @property {boolean} [showcase]
 * @property {string} [start_date]
 * @property {string} [status]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} ComponentUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {string} [automation_email]
 * @property {Object} [component]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {boolean} [group]
 * @property {string} [group_id]
 * @property {string} [name]
 * @property {boolean} [only_show_if_degraded]
 * @property {number} [position]
 * @property {boolean} [showcase]
 * @property {string} [start_date]
 * @property {string} [status]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} ComponentRemoveMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} ComponentGroupUptime
 * @property {string} [component_id]
 * @property {Object} [incidents]
 */

/**
 * @typedef {Object} ComponentGroupUptimeLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} GroupComponent
 * @property {Object} component_group
 * @property {string} [components]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [page_id]
 * @property {string} [position]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} GroupComponentLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} GroupComponentListMatch
 * @property {string} page_id
 */

/**
 * @typedef {Object} GroupComponentCreateData
 * @property {string} page_id
 * @property {Object} component_group
 * @property {string} [components]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {string} [id]
 * @property {string} [name]
 * @property {string} [position]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} GroupComponentUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {Object} [component_group]
 * @property {string} [components]
 * @property {string} [created_at]
 * @property {string} [description]
 * @property {string} [name]
 * @property {string} [position]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} GroupComponentRemoveMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} Incident
 * @property {boolean} [auto_transition_deliver_notifications_at_end]
 * @property {boolean} [auto_transition_deliver_notifications_at_start]
 * @property {boolean} [auto_transition_to_maintenance_state]
 * @property {boolean} [auto_transition_to_operational_state]
 * @property {Array} [components]
 * @property {string} [created_at]
 * @property {string} [id]
 * @property {string} [impact]
 * @property {string} [impact_override]
 * @property {Object} incident
 * @property {Array} [incident_updates]
 * @property {Object} [metadata]
 * @property {string} [monitoring_at]
 * @property {string} [name]
 * @property {string} [page_id]
 * @property {string} [postmortem_body]
 * @property {string} [postmortem_body_last_updated_at]
 * @property {boolean} [postmortem_ignored]
 * @property {boolean} [postmortem_notified_subscribers]
 * @property {boolean} [postmortem_notified_twitter]
 * @property {boolean} [postmortem_published_at]
 * @property {string} [reminder_intervals]
 * @property {string} [resolved_at]
 * @property {boolean} [scheduled_auto_completed]
 * @property {boolean} [scheduled_auto_in_progress]
 * @property {string} [scheduled_for]
 * @property {boolean} [scheduled_remind_prior]
 * @property {string} [scheduled_reminded_at]
 * @property {string} [scheduled_until]
 * @property {string} [shortlink]
 * @property {string} [status]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} IncidentLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} IncidentListMatch
 * @property {string} page_id
 */

/**
 * @typedef {Object} IncidentCreateData
 * @property {string} page_id
 * @property {boolean} [auto_transition_deliver_notifications_at_end]
 * @property {boolean} [auto_transition_deliver_notifications_at_start]
 * @property {boolean} [auto_transition_to_maintenance_state]
 * @property {boolean} [auto_transition_to_operational_state]
 * @property {Array} [components]
 * @property {string} [created_at]
 * @property {string} [id]
 * @property {string} [impact]
 * @property {string} [impact_override]
 * @property {Object} incident
 * @property {Array} [incident_updates]
 * @property {Object} [metadata]
 * @property {string} [monitoring_at]
 * @property {string} [name]
 * @property {string} [postmortem_body]
 * @property {string} [postmortem_body_last_updated_at]
 * @property {boolean} [postmortem_ignored]
 * @property {boolean} [postmortem_notified_subscribers]
 * @property {boolean} [postmortem_notified_twitter]
 * @property {boolean} [postmortem_published_at]
 * @property {string} [reminder_intervals]
 * @property {string} [resolved_at]
 * @property {boolean} [scheduled_auto_completed]
 * @property {boolean} [scheduled_auto_in_progress]
 * @property {string} [scheduled_for]
 * @property {boolean} [scheduled_remind_prior]
 * @property {string} [scheduled_reminded_at]
 * @property {string} [scheduled_until]
 * @property {string} [shortlink]
 * @property {string} [status]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} IncidentUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {boolean} [auto_transition_deliver_notifications_at_end]
 * @property {boolean} [auto_transition_deliver_notifications_at_start]
 * @property {boolean} [auto_transition_to_maintenance_state]
 * @property {boolean} [auto_transition_to_operational_state]
 * @property {Array} [components]
 * @property {string} [created_at]
 * @property {string} [impact]
 * @property {string} [impact_override]
 * @property {Object} [incident]
 * @property {Array} [incident_updates]
 * @property {Object} [metadata]
 * @property {string} [monitoring_at]
 * @property {string} [name]
 * @property {string} [postmortem_body]
 * @property {string} [postmortem_body_last_updated_at]
 * @property {boolean} [postmortem_ignored]
 * @property {boolean} [postmortem_notified_subscribers]
 * @property {boolean} [postmortem_notified_twitter]
 * @property {boolean} [postmortem_published_at]
 * @property {string} [reminder_intervals]
 * @property {string} [resolved_at]
 * @property {boolean} [scheduled_auto_completed]
 * @property {boolean} [scheduled_auto_in_progress]
 * @property {string} [scheduled_for]
 * @property {boolean} [scheduled_remind_prior]
 * @property {string} [scheduled_reminded_at]
 * @property {string} [scheduled_until]
 * @property {string} [shortlink]
 * @property {string} [status]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} IncidentRemoveMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} IncidentPostmortem
 */

/**
 * @typedef {Object} IncidentPostmortemRemoveMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} IncidentSubscriber
 */

/**
 * @typedef {Object} IncidentSubscriberCreateData
 * @property {string} incident_id
 * @property {string} page_id
 * @property {string} subscriber_id
 */

/**
 * @typedef {Object} IncidentTemplate
 * @property {string} [body]
 * @property {Array} [components]
 * @property {string} [group_id]
 * @property {string} [id]
 * @property {string} [name]
 * @property {boolean} [should_send_notifications]
 * @property {boolean} [should_tweet]
 * @property {Object} template
 * @property {string} [title]
 * @property {string} [update_status]
 */

/**
 * @typedef {Object} IncidentTemplateListMatch
 * @property {string} page_id
 */

/**
 * @typedef {Object} IncidentTemplateCreateData
 * @property {string} page_id
 * @property {string} [body]
 * @property {Array} [components]
 * @property {string} [group_id]
 * @property {string} [id]
 * @property {string} [name]
 * @property {boolean} [should_send_notifications]
 * @property {boolean} [should_tweet]
 * @property {Object} template
 * @property {string} [title]
 * @property {string} [update_status]
 */

/**
 * @typedef {Object} IncidentUpdate
 * @property {Array} [affected_components]
 * @property {string} [body]
 * @property {string} [created_at]
 * @property {string} [custom_tweet]
 * @property {boolean} [deliver_notifications]
 * @property {string} [display_at]
 * @property {string} [id]
 * @property {string} [incident_id]
 * @property {Object} [incident_update]
 * @property {string} [status]
 * @property {string} [tweet_id]
 * @property {string} [twitter_updated_at]
 * @property {string} [updated_at]
 * @property {boolean} [wants_twitter_update]
 */

/**
 * @typedef {Object} IncidentUpdateUpdateData
 * @property {string} id
 * @property {string} incident_id
 * @property {string} page_id
 * @property {Array} [affected_components]
 * @property {string} [body]
 * @property {string} [created_at]
 * @property {string} [custom_tweet]
 * @property {boolean} [deliver_notifications]
 * @property {string} [display_at]
 * @property {Object} [incident_update]
 * @property {string} [status]
 * @property {string} [tweet_id]
 * @property {string} [twitter_updated_at]
 * @property {string} [updated_at]
 * @property {boolean} [wants_twitter_update]
 */

/**
 * @typedef {Object} Metric
 * @property {number} [backfill_percentage]
 * @property {boolean} [backfilled]
 * @property {string} [created_at]
 * @property {Object} data
 * @property {number} [decimal_places]
 * @property {boolean} [display]
 * @property {string} [id]
 * @property {string} [last_fetched_at]
 * @property {Object} [metric]
 * @property {string} [metric_identifier]
 * @property {string} [metrics_provider_id]
 * @property {string} [most_recent_data_at]
 * @property {string} [name]
 * @property {string} [reference_name]
 * @property {string} [suffix]
 * @property {string} [tooltip_description]
 * @property {string} [updated_at]
 * @property {boolean} [y_axis_hidden]
 * @property {number} [y_axis_max]
 * @property {number} [y_axis_min]
 */

/**
 * @typedef {Object} MetricLoadMatch
 * @property {string} [metrics_provider_id]
 * @property {string} page_id
 * @property {string} [id]
 */

/**
 * @typedef {Object} MetricListMatch
 * @property {string} page_access_user_id
 * @property {string} page_id
 */

/**
 * @typedef {Object} MetricCreateData
 * @property {string} metrics_provider_id
 * @property {string} page_id
 * @property {number} [backfill_percentage]
 * @property {boolean} [backfilled]
 * @property {string} [created_at]
 * @property {Object} data
 * @property {number} [decimal_places]
 * @property {boolean} [display]
 * @property {string} [id]
 * @property {string} [last_fetched_at]
 * @property {Object} [metric]
 * @property {string} [metric_identifier]
 * @property {string} [most_recent_data_at]
 * @property {string} [name]
 * @property {string} [reference_name]
 * @property {string} [suffix]
 * @property {string} [tooltip_description]
 * @property {string} [updated_at]
 * @property {boolean} [y_axis_hidden]
 * @property {number} [y_axis_max]
 * @property {number} [y_axis_min]
 */

/**
 * @typedef {Object} MetricUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {number} [backfill_percentage]
 * @property {boolean} [backfilled]
 * @property {string} [created_at]
 * @property {Object} [data]
 * @property {number} [decimal_places]
 * @property {boolean} [display]
 * @property {string} [last_fetched_at]
 * @property {Object} [metric]
 * @property {string} [metric_identifier]
 * @property {string} [metrics_provider_id]
 * @property {string} [most_recent_data_at]
 * @property {string} [name]
 * @property {string} [reference_name]
 * @property {string} [suffix]
 * @property {string} [tooltip_description]
 * @property {string} [updated_at]
 * @property {boolean} [y_axis_hidden]
 * @property {number} [y_axis_max]
 * @property {number} [y_axis_min]
 */

/**
 * @typedef {Object} MetricRemoveMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} MetricsProvider
 * @property {string} [created_at]
 * @property {boolean} [disabled]
 * @property {string} [id]
 * @property {string} [last_revalidated_at]
 * @property {string} [metric_base_uri]
 * @property {Object} [metrics_provider]
 * @property {number} [page_id]
 * @property {string} [type]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} MetricsProviderLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} MetricsProviderListMatch
 * @property {string} page_id
 */

/**
 * @typedef {Object} MetricsProviderCreateData
 * @property {string} page_id
 * @property {string} [created_at]
 * @property {boolean} [disabled]
 * @property {string} [id]
 * @property {string} [last_revalidated_at]
 * @property {string} [metric_base_uri]
 * @property {Object} [metrics_provider]
 * @property {string} [type]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} MetricsProviderUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {string} [created_at]
 * @property {boolean} [disabled]
 * @property {string} [last_revalidated_at]
 * @property {string} [metric_base_uri]
 * @property {Object} [metrics_provider]
 * @property {string} [type]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} MetricsProviderRemoveMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} Page
 * @property {number} [activity_score]
 * @property {boolean} [allow_email_subscribers]
 * @property {boolean} [allow_incident_subscribers]
 * @property {boolean} [allow_page_subscribers]
 * @property {boolean} [allow_rss_atom_feeds]
 * @property {boolean} [allow_sms_subscribers]
 * @property {boolean} [allow_webhook_subscribers]
 * @property {string} [branding]
 * @property {string} [city]
 * @property {string} [country]
 * @property {string} [created_at]
 * @property {string} [css_blues]
 * @property {string} [css_body_background_color]
 * @property {string} [css_border_color]
 * @property {string} [css_font_color]
 * @property {string} [css_graph_color]
 * @property {string} [css_greens]
 * @property {string} [css_light_font_color]
 * @property {string} [css_link_color]
 * @property {string} [css_no_data]
 * @property {string} [css_oranges]
 * @property {string} [css_reds]
 * @property {string} [css_yellows]
 * @property {string} [domain]
 * @property {string} [email_logo]
 * @property {string} [favicon_logo]
 * @property {string} [headline]
 * @property {string} [hero_cover]
 * @property {boolean} [hidden_from_search]
 * @property {string} [id]
 * @property {string} [ip_restrictions]
 * @property {string} [name]
 * @property {string} [notifications_email_footer]
 * @property {string} [notifications_from_email]
 * @property {Object} [page]
 * @property {string} [page_description]
 * @property {string} [state]
 * @property {string} [subdomain]
 * @property {string} [support_url]
 * @property {string} [time_zone]
 * @property {string} [transactional_logo]
 * @property {string} [twitter_logo]
 * @property {string} [twitter_username]
 * @property {string} [updated_at]
 * @property {string} [url]
 * @property {boolean} [viewers_must_be_team_members]
 */

/**
 * @typedef {Object} PageLoadMatch
 * @property {string} id
 */

/**
 * @typedef {Object} PageListMatch
 * @property {number} [activity_score]
 * @property {boolean} [allow_email_subscribers]
 * @property {boolean} [allow_incident_subscribers]
 * @property {boolean} [allow_page_subscribers]
 * @property {boolean} [allow_rss_atom_feeds]
 * @property {boolean} [allow_sms_subscribers]
 * @property {boolean} [allow_webhook_subscribers]
 * @property {string} [branding]
 * @property {string} [city]
 * @property {string} [country]
 * @property {string} [created_at]
 * @property {string} [css_blues]
 * @property {string} [css_body_background_color]
 * @property {string} [css_border_color]
 * @property {string} [css_font_color]
 * @property {string} [css_graph_color]
 * @property {string} [css_greens]
 * @property {string} [css_light_font_color]
 * @property {string} [css_link_color]
 * @property {string} [css_no_data]
 * @property {string} [css_oranges]
 * @property {string} [css_reds]
 * @property {string} [css_yellows]
 * @property {string} [domain]
 * @property {string} [email_logo]
 * @property {string} [favicon_logo]
 * @property {string} [headline]
 * @property {string} [hero_cover]
 * @property {boolean} [hidden_from_search]
 * @property {string} [id]
 * @property {string} [ip_restrictions]
 * @property {string} [name]
 * @property {string} [notifications_email_footer]
 * @property {string} [notifications_from_email]
 * @property {Object} [page]
 * @property {string} [page_description]
 * @property {string} [state]
 * @property {string} [subdomain]
 * @property {string} [support_url]
 * @property {string} [time_zone]
 * @property {string} [transactional_logo]
 * @property {string} [twitter_logo]
 * @property {string} [twitter_username]
 * @property {string} [updated_at]
 * @property {string} [url]
 * @property {boolean} [viewers_must_be_team_members]
 */

/**
 * @typedef {Object} PageUpdateData
 * @property {string} id
 * @property {number} [activity_score]
 * @property {boolean} [allow_email_subscribers]
 * @property {boolean} [allow_incident_subscribers]
 * @property {boolean} [allow_page_subscribers]
 * @property {boolean} [allow_rss_atom_feeds]
 * @property {boolean} [allow_sms_subscribers]
 * @property {boolean} [allow_webhook_subscribers]
 * @property {string} [branding]
 * @property {string} [city]
 * @property {string} [country]
 * @property {string} [created_at]
 * @property {string} [css_blues]
 * @property {string} [css_body_background_color]
 * @property {string} [css_border_color]
 * @property {string} [css_font_color]
 * @property {string} [css_graph_color]
 * @property {string} [css_greens]
 * @property {string} [css_light_font_color]
 * @property {string} [css_link_color]
 * @property {string} [css_no_data]
 * @property {string} [css_oranges]
 * @property {string} [css_reds]
 * @property {string} [css_yellows]
 * @property {string} [domain]
 * @property {string} [email_logo]
 * @property {string} [favicon_logo]
 * @property {string} [headline]
 * @property {string} [hero_cover]
 * @property {boolean} [hidden_from_search]
 * @property {string} [ip_restrictions]
 * @property {string} [name]
 * @property {string} [notifications_email_footer]
 * @property {string} [notifications_from_email]
 * @property {Object} [page]
 * @property {string} [page_description]
 * @property {string} [state]
 * @property {string} [subdomain]
 * @property {string} [support_url]
 * @property {string} [time_zone]
 * @property {string} [transactional_logo]
 * @property {string} [twitter_logo]
 * @property {string} [twitter_username]
 * @property {string} [updated_at]
 * @property {string} [url]
 * @property {boolean} [viewers_must_be_team_members]
 */

/**
 * @typedef {Object} PageAccessGroup
 * @property {Array} [component_ids]
 * @property {string} [created_at]
 * @property {string} [external_identifier]
 * @property {string} [id]
 * @property {Array} [metric_ids]
 * @property {string} [name]
 * @property {Object} [page_access_group]
 * @property {Array} [page_access_user_ids]
 * @property {string} [page_id]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PageAccessGroupLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} PageAccessGroupListMatch
 * @property {string} id
 */

/**
 * @typedef {Object} PageAccessGroupCreateData
 * @property {string} id
 * @property {Array} [component_ids]
 * @property {string} [created_at]
 * @property {string} [external_identifier]
 * @property {Array} [metric_ids]
 * @property {string} [name]
 * @property {Object} [page_access_group]
 * @property {Array} [page_access_user_ids]
 * @property {string} [page_id]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PageAccessGroupUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {Array} [component_ids]
 * @property {string} [created_at]
 * @property {string} [external_identifier]
 * @property {Array} [metric_ids]
 * @property {string} [name]
 * @property {Object} [page_access_group]
 * @property {Array} [page_access_user_ids]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PageAccessGroupRemoveMatch
 * @property {string} [component_id]
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} PageAccessUser
 * @property {Array} component_ids
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {string} [external_login]
 * @property {string} [id]
 * @property {Array} metric_ids
 * @property {string} [page_access_group_id]
 * @property {string} [page_access_group_ids]
 * @property {Object} [page_access_user]
 * @property {string} [page_id]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PageAccessUserLoadMatch
 * @property {string} id
 * @property {string} page_id
 */

/**
 * @typedef {Object} PageAccessUserListMatch
 * @property {string} id
 */

/**
 * @typedef {Object} PageAccessUserCreateData
 * @property {string} id
 * @property {Array} component_ids
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {string} [external_login]
 * @property {Array} metric_ids
 * @property {string} [page_access_group_id]
 * @property {string} [page_access_group_ids]
 * @property {Object} [page_access_user]
 * @property {string} [page_id]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PageAccessUserUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {Array} [component_ids]
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {string} [external_login]
 * @property {Array} [metric_ids]
 * @property {string} [page_access_group_id]
 * @property {string} [page_access_group_ids]
 * @property {Object} [page_access_user]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PageAccessUserRemoveMatch
 * @property {string} [component_id]
 * @property {string} id
 * @property {string} page_id
 * @property {string} [metric_id]
 */

/**
 * @typedef {Object} Permission
 * @property {Object} [pages]
 * @property {string} [user_id]
 */

/**
 * @typedef {Object} PermissionLoadMatch
 * @property {string} id
 * @property {string} organization_id
 */

/**
 * @typedef {Object} PermissionUpdateData
 * @property {string} id
 * @property {string} organization_id
 * @property {Object} [pages]
 * @property {string} [user_id]
 */

/**
 * @typedef {Object} Postmortem
 * @property {string} [body]
 * @property {string} [body_draft]
 * @property {string} [body_draft_updated_at]
 * @property {string} [body_updated_at]
 * @property {string} [created_at]
 * @property {string} [custom_tweet]
 * @property {boolean} [notify_subscribers]
 * @property {boolean} [notify_twitter]
 * @property {Object} postmortem
 * @property {string} [preview_key]
 * @property {string} [published_at]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} PostmortemLoadMatch
 * @property {string} incident_id
 * @property {string} page_id
 */

/**
 * @typedef {Object} PostmortemUpdateData
 * @property {string} incident_id
 * @property {string} page_id
 * @property {string} [body]
 * @property {string} [body_draft]
 * @property {string} [body_draft_updated_at]
 * @property {string} [body_updated_at]
 * @property {string} [created_at]
 * @property {string} [custom_tweet]
 * @property {boolean} [notify_subscribers]
 * @property {boolean} [notify_twitter]
 * @property {Object} [postmortem]
 * @property {string} [preview_key]
 * @property {string} [published_at]
 * @property {string} [updated_at]
 */

/**
 * @typedef {Object} StatusEmbedConfig
 * @property {string} [incident_background_color]
 * @property {string} [incident_text_color]
 * @property {string} [maintenance_background_color]
 * @property {string} [maintenance_text_color]
 * @property {string} [page_id]
 * @property {string} [position]
 * @property {Object} [status_embed_config]
 */

/**
 * @typedef {Object} StatusEmbedConfigLoadMatch
 * @property {string} page_id
 */

/**
 * @typedef {Object} StatusEmbedConfigUpdateData
 * @property {string} page_id
 * @property {string} [incident_background_color]
 * @property {string} [incident_text_color]
 * @property {string} [maintenance_background_color]
 * @property {string} [maintenance_text_color]
 * @property {string} [position]
 * @property {Object} [status_embed_config]
 */

/**
 * @typedef {Object} Subscriber
 * @property {Array} [component_ids]
 * @property {string} [components]
 * @property {string} [created_at]
 * @property {string} [display_phone_number]
 * @property {string} [email]
 * @property {string} [endpoint]
 * @property {string} [id]
 * @property {number} [integration_partner]
 * @property {string} [mode]
 * @property {string} [obfuscated_channel_name]
 * @property {string} [page_access_user_id]
 * @property {string} [phone_country]
 * @property {string} [phone_number]
 * @property {string} [purge_at]
 * @property {string} [quarantined_at]
 * @property {boolean} [skip_confirmation_notification]
 * @property {boolean} [skip_unsubscription_notification]
 * @property {number} [slack]
 * @property {number} [sms]
 * @property {string} [state]
 * @property {Object} [subscriber]
 * @property {string} subscribers
 * @property {number} [teams]
 * @property {string} [type]
 * @property {number} [webhook]
 * @property {string} [workspace_name]
 */

/**
 * @typedef {Object} SubscriberLoadMatch
 * @property {string} id
 * @property {string} [incident_id]
 * @property {string} page_id
 */

/**
 * @typedef {Object} SubscriberListMatch
 * @property {string} page_id
 * @property {string} [incident_id]
 */

/**
 * @typedef {Object} SubscriberCreateData
 * @property {string} [incident_id]
 * @property {string} page_id
 * @property {Array} [component_ids]
 * @property {string} [components]
 * @property {string} [created_at]
 * @property {string} [display_phone_number]
 * @property {string} [email]
 * @property {string} [endpoint]
 * @property {string} [id]
 * @property {number} [integration_partner]
 * @property {string} [mode]
 * @property {string} [obfuscated_channel_name]
 * @property {string} [page_access_user_id]
 * @property {string} [phone_country]
 * @property {string} [phone_number]
 * @property {string} [purge_at]
 * @property {string} [quarantined_at]
 * @property {boolean} [skip_confirmation_notification]
 * @property {boolean} [skip_unsubscription_notification]
 * @property {number} [slack]
 * @property {number} [sms]
 * @property {string} [state]
 * @property {Object} [subscriber]
 * @property {string} subscribers
 * @property {number} [teams]
 * @property {string} [type]
 * @property {number} [webhook]
 * @property {string} [workspace_name]
 */

/**
 * @typedef {Object} SubscriberUpdateData
 * @property {string} id
 * @property {string} page_id
 * @property {Array} [component_ids]
 * @property {string} [components]
 * @property {string} [created_at]
 * @property {string} [display_phone_number]
 * @property {string} [email]
 * @property {string} [endpoint]
 * @property {number} [integration_partner]
 * @property {string} [mode]
 * @property {string} [obfuscated_channel_name]
 * @property {string} [page_access_user_id]
 * @property {string} [phone_country]
 * @property {string} [phone_number]
 * @property {string} [purge_at]
 * @property {string} [quarantined_at]
 * @property {boolean} [skip_confirmation_notification]
 * @property {boolean} [skip_unsubscription_notification]
 * @property {number} [slack]
 * @property {number} [sms]
 * @property {string} [state]
 * @property {Object} [subscriber]
 * @property {string} [subscribers]
 * @property {number} [teams]
 * @property {string} [type]
 * @property {number} [webhook]
 * @property {string} [workspace_name]
 */

/**
 * @typedef {Object} SubscriberRemoveMatch
 * @property {string} id
 * @property {string} [incident_id]
 * @property {string} page_id
 */

/**
 * @typedef {Object} User
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {string} [first_name]
 * @property {string} [id]
 * @property {string} [last_name]
 * @property {string} [organization_id]
 * @property {string} [updated_at]
 * @property {Object} user
 */

/**
 * @typedef {Object} UserListMatch
 * @property {string} organization_id
 */

/**
 * @typedef {Object} UserCreateData
 * @property {string} organization_id
 * @property {string} [created_at]
 * @property {string} [email]
 * @property {string} [first_name]
 * @property {string} [id]
 * @property {string} [last_name]
 * @property {string} [updated_at]
 * @property {Object} user
 */

/**
 * @typedef {Object} UserRemoveMatch
 * @property {string} id
 * @property {string} organization_id
 */

