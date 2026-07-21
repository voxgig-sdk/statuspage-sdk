-- Typed models for the Statuspage SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Component
---@field automation_email? string
---@field component? table
---@field created_at? string
---@field description? string
---@field group? boolean
---@field group_id? string
---@field id? string
---@field major_outage? number
---@field name? string
---@field only_show_if_degraded? boolean
---@field page_id? string
---@field partial_outage? number
---@field position? number
---@field range_end? string
---@field range_start? string
---@field related_event? table
---@field showcase? boolean
---@field start_date? string
---@field status? string
---@field updated_at? string
---@field uptime_percentage? number
---@field warning? string

---@class ComponentLoadMatch
---@field id string
---@field page_id string

---@class ComponentListMatch
---@field page_access_group_id? string
---@field page_id string
---@field page_access_user_id? string

---@class ComponentCreateData
---@field page_id string

---@class ComponentUpdateData
---@field id string
---@field page_id string

---@class ComponentRemoveMatch
---@field id string
---@field page_id string

---@class ComponentGroupUptime
---@field id? string
---@field major_outage? number
---@field name? string
---@field partial_outage? number
---@field range_end? string
---@field range_start? string
---@field related_event? table
---@field uptime_percentage? number
---@field warning? string

---@class ComponentGroupUptimeLoadMatch
---@field id string
---@field page_id string

---@class GroupComponent
---@field component? string
---@field component_group table
---@field created_at? string
---@field description? string
---@field id? string
---@field name? string
---@field page_id? string
---@field position? string
---@field updated_at? string

---@class GroupComponentLoadMatch
---@field id string
---@field page_id string

---@class GroupComponentListMatch
---@field page_id string

---@class GroupComponentCreateData
---@field page_id string

---@class GroupComponentUpdateData
---@field id string
---@field page_id string

---@class GroupComponentRemoveMatch
---@field id string
---@field page_id string

---@class Incident
---@field auto_transition_deliver_notifications_at_end? boolean
---@field auto_transition_deliver_notifications_at_start? boolean
---@field auto_transition_to_maintenance_state? boolean
---@field auto_transition_to_operational_state? boolean
---@field component? table
---@field created_at? string
---@field id? string
---@field impact? string
---@field impact_override? string
---@field incident table
---@field incident_impact? table
---@field incident_update? table
---@field metadata? table
---@field monitoring_at? string
---@field name? string
---@field page_id? string
---@field postmortem_body? string
---@field postmortem_body_last_updated_at? string
---@field postmortem_ignored? boolean
---@field postmortem_notified_subscriber? boolean
---@field postmortem_notified_twitter? boolean
---@field postmortem_published_at? boolean
---@field reminder_interval? string
---@field resolved_at? string
---@field scheduled_auto_completed? boolean
---@field scheduled_auto_in_progress? boolean
---@field scheduled_for? string
---@field scheduled_remind_prior? boolean
---@field scheduled_reminded_at? string
---@field scheduled_until? string
---@field shortlink? string
---@field status? string
---@field updated_at? string

---@class IncidentLoadMatch
---@field id string
---@field page_id string

---@class IncidentListMatch
---@field page_id string

---@class IncidentCreateData
---@field page_id string

---@class IncidentUpdateData
---@field id string
---@field page_id string

---@class IncidentRemoveMatch
---@field id string
---@field page_id string

---@class IncidentPostmortem

---@class IncidentPostmortemRemoveMatch
---@field id string
---@field page_id string

---@class IncidentSubscriber

---@class IncidentSubscriberCreateData
---@field incident_id string
---@field page_id string
---@field subscriber_id string

---@class IncidentTemplate
---@field body? string
---@field component? table
---@field group_id? string
---@field id? string
---@field name? string
---@field should_send_notification? boolean
---@field should_tweet? boolean
---@field template table
---@field title? string
---@field update_status? string

---@class IncidentTemplateListMatch
---@field page_id string

---@class IncidentTemplateCreateData
---@field page_id string

---@class IncidentUpdate
---@field affected_component? table
---@field body? string
---@field created_at? string
---@field custom_tweet? string
---@field deliver_notification? boolean
---@field display_at? string
---@field id? string
---@field incident_id? string
---@field incident_update? table
---@field status? string
---@field tweet_id? string
---@field twitter_updated_at? string
---@field updated_at? string
---@field wants_twitter_update? boolean

---@class IncidentUpdateUpdateData
---@field id string
---@field incident_id string
---@field page_id string

---@class Metric
---@field backfill_percentage? number
---@field backfilled? boolean
---@field created_at? string
---@field data table
---@field decimal_place? number
---@field display? boolean
---@field id? string
---@field last_fetched_at? string
---@field metric? table
---@field metric_identifier? string
---@field metrics_provider_id? string
---@field most_recent_data_at? string
---@field name? string
---@field reference_name? string
---@field suffix? string
---@field tooltip_description? string
---@field updated_at? string
---@field y_axis_hidden? boolean
---@field y_axis_max? number
---@field y_axis_min? number

---@class MetricLoadMatch
---@field metrics_provider_id? string
---@field page_id string
---@field id? string

---@class MetricListMatch
---@field page_access_user_id string
---@field page_id string

---@class MetricCreateData
---@field metrics_provider_id string
---@field page_id string

---@class MetricUpdateData
---@field id string
---@field page_id string

---@class MetricRemoveMatch
---@field id string
---@field page_id string

---@class MetricsProvider
---@field created_at? string
---@field disabled? boolean
---@field id? string
---@field last_revalidated_at? string
---@field metric_base_uri? string
---@field metrics_provider? table
---@field page_id? number
---@field type? string
---@field updated_at? string

---@class MetricsProviderLoadMatch
---@field id string
---@field page_id string

---@class MetricsProviderListMatch
---@field page_id string

---@class MetricsProviderCreateData
---@field page_id string

---@class MetricsProviderUpdateData
---@field id string
---@field page_id string

---@class MetricsProviderRemoveMatch
---@field id string
---@field page_id string

---@class Page
---@field activity_score? number
---@field allow_email_subscriber? boolean
---@field allow_incident_subscriber? boolean
---@field allow_page_subscriber? boolean
---@field allow_rss_atom_feed? boolean
---@field allow_sms_subscriber? boolean
---@field allow_webhook_subscriber? boolean
---@field branding? string
---@field city? string
---@field country? string
---@field created_at? string
---@field css_blue? string
---@field css_body_background_color? string
---@field css_border_color? string
---@field css_font_color? string
---@field css_graph_color? string
---@field css_green? string
---@field css_light_font_color? string
---@field css_link_color? string
---@field css_no_data? string
---@field css_orange? string
---@field css_red? string
---@field css_yellow? string
---@field domain? string
---@field email_logo? string
---@field favicon_logo? string
---@field headline? string
---@field hero_cover? string
---@field hidden_from_search? boolean
---@field id? string
---@field ip_restriction? string
---@field name? string
---@field notifications_email_footer? string
---@field notifications_from_email? string
---@field page? table
---@field page_description? string
---@field state? string
---@field subdomain? string
---@field support_url? string
---@field time_zone? string
---@field transactional_logo? string
---@field twitter_logo? string
---@field twitter_username? string
---@field updated_at? string
---@field url? string
---@field viewers_must_be_team_member? boolean

---@class PageLoadMatch
---@field id string

---@class PageListMatch
---@field activity_score? number
---@field allow_email_subscriber? boolean
---@field allow_incident_subscriber? boolean
---@field allow_page_subscriber? boolean
---@field allow_rss_atom_feed? boolean
---@field allow_sms_subscriber? boolean
---@field allow_webhook_subscriber? boolean
---@field branding? string
---@field city? string
---@field country? string
---@field created_at? string
---@field css_blue? string
---@field css_body_background_color? string
---@field css_border_color? string
---@field css_font_color? string
---@field css_graph_color? string
---@field css_green? string
---@field css_light_font_color? string
---@field css_link_color? string
---@field css_no_data? string
---@field css_orange? string
---@field css_red? string
---@field css_yellow? string
---@field domain? string
---@field email_logo? string
---@field favicon_logo? string
---@field headline? string
---@field hero_cover? string
---@field hidden_from_search? boolean
---@field id? string
---@field ip_restriction? string
---@field name? string
---@field notifications_email_footer? string
---@field notifications_from_email? string
---@field page? table
---@field page_description? string
---@field state? string
---@field subdomain? string
---@field support_url? string
---@field time_zone? string
---@field transactional_logo? string
---@field twitter_logo? string
---@field twitter_username? string
---@field updated_at? string
---@field url? string
---@field viewers_must_be_team_member? boolean

---@class PageUpdateData
---@field id string

---@class PageAccessGroup
---@field component_id? table
---@field created_at? string
---@field external_identifier? string
---@field id? string
---@field metric_id? table
---@field name? string
---@field page_access_group? table
---@field page_access_user_id? table
---@field page_id? string
---@field updated_at? string

---@class PageAccessGroupLoadMatch
---@field id string
---@field page_id string

---@class PageAccessGroupListMatch
---@field id string

---@class PageAccessGroupCreateData
---@field id string

---@class PageAccessGroupUpdateData
---@field id string
---@field page_id string

---@class PageAccessGroupRemoveMatch
---@field component_id? string
---@field id string
---@field page_id string

---@class PageAccessUser
---@field component_id table
---@field created_at? string
---@field email? string
---@field external_login? string
---@field id? string
---@field metric_id table
---@field page_access_group_id? string
---@field page_access_user? table
---@field page_id? string
---@field updated_at? string

---@class PageAccessUserLoadMatch
---@field id string
---@field page_id string

---@class PageAccessUserListMatch
---@field id string

---@class PageAccessUserCreateData
---@field id string

---@class PageAccessUserUpdateData
---@field id string
---@field page_id string

---@class PageAccessUserRemoveMatch
---@field component_id? string
---@field id string
---@field page_id string
---@field metric_id? string

---@class Permission
---@field data? table
---@field page? table

---@class PermissionLoadMatch
---@field id string
---@field organization_id string

---@class PermissionUpdateData
---@field id string
---@field organization_id string

---@class Postmortem
---@field body? string
---@field body_draft? string
---@field body_draft_updated_at? string
---@field body_updated_at? string
---@field created_at? string
---@field custom_tweet? string
---@field notify_subscriber? boolean
---@field notify_twitter? boolean
---@field postmortem table
---@field preview_key? string
---@field published_at? string
---@field updated_at? string

---@class PostmortemLoadMatch
---@field incident_id string
---@field page_id string

---@class PostmortemUpdateData
---@field incident_id string
---@field page_id string

---@class StatusEmbedConfig
---@field incident_background_color? string
---@field incident_text_color? string
---@field maintenance_background_color? string
---@field maintenance_text_color? string
---@field page_id? string
---@field position? string
---@field status_embed_config? table

---@class StatusEmbedConfigLoadMatch
---@field page_id string

---@class StatusEmbedConfigUpdateData
---@field page_id string

---@class Subscriber
---@field component? string
---@field component_id? table
---@field created_at? string
---@field display_phone_number? string
---@field email? string
---@field endpoint? string
---@field id? string
---@field integration_partner? number
---@field mode? string
---@field obfuscated_channel_name? string
---@field page_access_user_id? string
---@field phone_country? string
---@field phone_number? string
---@field purge_at? string
---@field quarantined_at? string
---@field skip_confirmation_notification? boolean
---@field skip_unsubscription_notification? boolean
---@field slack? number
---@field sms? number
---@field state? string
---@field subscriber? table
---@field team? number
---@field type? string
---@field webhook? number
---@field workspace_name? string

---@class SubscriberLoadMatch
---@field id string
---@field incident_id? string
---@field page_id string

---@class SubscriberListMatch
---@field page_id string
---@field incident_id? string

---@class SubscriberCreateData
---@field incident_id? string
---@field page_id string

---@class SubscriberUpdateData
---@field id string
---@field page_id string

---@class SubscriberRemoveMatch
---@field id string
---@field incident_id? string
---@field page_id string

---@class User
---@field created_at? string
---@field email? string
---@field first_name? string
---@field id? string
---@field last_name? string
---@field organization_id? string
---@field updated_at? string
---@field user table

---@class UserListMatch
---@field organization_id string

---@class UserCreateData
---@field organization_id string

---@class UserRemoveMatch
---@field id string
---@field organization_id string

local M = {}

return M
