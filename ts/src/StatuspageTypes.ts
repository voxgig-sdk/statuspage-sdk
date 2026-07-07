// Typed models for the Statuspage SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Component {
  automation_email?: string
  component?: Record<string, any>
  created_at?: string
  description?: string
  group?: boolean
  group_id?: string
  id?: string
  major_outage?: number
  name?: string
  only_show_if_degraded?: boolean
  page_id?: string
  partial_outage?: number
  position?: number
  range_end?: string
  range_start?: string
  related_event?: Record<string, any>
  showcase?: boolean
  start_date?: string
  status?: string
  updated_at?: string
  uptime_percentage?: number
  warning?: string
}

export interface ComponentLoadMatch {
  id: string
  page_id: string
}

export interface ComponentListMatch {
  page_access_group_id: string
  page_id: string
  page_access_user_id: string
}

export interface ComponentCreateData {
  id: string
  page_id: string
}

export interface ComponentUpdateData {
  id: string
  page_id: string
}

export interface ComponentRemoveMatch {
  id: string
  page_id: string
}

export interface ComponentGroupUptime {
  id?: string
  major_outage?: number
  name?: string
  partial_outage?: number
  range_end?: string
  range_start?: string
  related_event?: Record<string, any>
  uptime_percentage?: number
  warning?: string
}

export interface ComponentGroupUptimeLoadMatch {
  id: string
  page_id: string
}

export interface GroupComponent {
  component?: string
  component_group: Record<string, any>
  created_at?: string
  description?: string
  id?: string
  name?: string
  page_id?: string
  position?: string
  updated_at?: string
}

export interface GroupComponentLoadMatch {
  id: string
  page_id: string
}

export interface GroupComponentListMatch {
  page_id: string
}

export interface GroupComponentCreateData {
  page_id: string
}

export interface GroupComponentUpdateData {
  id: string
  page_id: string
}

export interface GroupComponentRemoveMatch {
  id: string
  page_id: string
}

export interface Incident {
  auto_transition_deliver_notifications_at_end?: boolean
  auto_transition_deliver_notifications_at_start?: boolean
  auto_transition_to_maintenance_state?: boolean
  auto_transition_to_operational_state?: boolean
  component?: any[]
  created_at?: string
  id?: string
  impact?: string
  impact_override?: string
  incident: Record<string, any>
  incident_impact?: any[]
  incident_update?: any[]
  metadata?: Record<string, any>
  monitoring_at?: string
  name?: string
  page_id?: string
  postmortem_body?: string
  postmortem_body_last_updated_at?: string
  postmortem_ignored?: boolean
  postmortem_notified_subscriber?: boolean
  postmortem_notified_twitter?: boolean
  postmortem_published_at?: boolean
  reminder_interval?: string
  resolved_at?: string
  scheduled_auto_completed?: boolean
  scheduled_auto_in_progress?: boolean
  scheduled_for?: string
  scheduled_remind_prior?: boolean
  scheduled_reminded_at?: string
  scheduled_until?: string
  shortlink?: string
  status?: string
  updated_at?: string
}

export interface IncidentLoadMatch {
  id: string
  page_id: string
}

export interface IncidentListMatch {
  page_id: string
}

export interface IncidentCreateData {
  page_id: string
}

export interface IncidentUpdateData {
  id: string
  page_id: string
}

export interface IncidentRemoveMatch {
  id: string
  page_id: string
}

export interface IncidentPostmortem {
}

export interface IncidentPostmortemRemoveMatch {
  id: string
  page_id: string
}

export interface IncidentSubscriber {
}

export interface IncidentSubscriberCreateData {
  incident_id: string
  page_id: string
  subscriber_id: string
}

export interface IncidentTemplate {
  body?: string
  component?: any[]
  group_id?: string
  id?: string
  name?: string
  should_send_notification?: boolean
  should_tweet?: boolean
  template: Record<string, any>
  title?: string
  update_status?: string
}

export interface IncidentTemplateListMatch {
  page_id: string
}

export interface IncidentTemplateCreateData {
  page_id: string
}

export interface IncidentUpdate {
  affected_component?: any[]
  body?: string
  created_at?: string
  custom_tweet?: string
  deliver_notification?: boolean
  display_at?: string
  id?: string
  incident_id?: string
  incident_update?: Record<string, any>
  status?: string
  tweet_id?: string
  twitter_updated_at?: string
  updated_at?: string
  wants_twitter_update?: boolean
}

export interface IncidentUpdateUpdateData {
  id: string
  incident_id: string
  page_id: string
}

export interface Metric {
  backfill_percentage?: number
  backfilled?: boolean
  created_at?: string
  data: Record<string, any>
  decimal_place?: number
  display?: boolean
  id?: string
  last_fetched_at?: string
  metric?: Record<string, any>
  metric_identifier?: string
  metrics_provider_id?: string
  most_recent_data_at?: string
  name?: string
  reference_name?: string
  suffix?: string
  tooltip_description?: string
  updated_at?: string
  y_axis_hidden?: boolean
  y_axis_max?: number
  y_axis_min?: number
}

export interface MetricLoadMatch {
  metrics_provider_id: string
  page_id: string
  id: string
}

export interface MetricListMatch {
  page_access_user_id: string
  page_id: string
}

export interface MetricCreateData {
  id: string
  page_id: string
  metrics_provider_id: string
}

export interface MetricUpdateData {
  id: string
  page_id: string
}

export interface MetricRemoveMatch {
  id: string
  page_id: string
}

export interface MetricsProvider {
  created_at?: string
  disabled?: boolean
  id?: string
  last_revalidated_at?: string
  metric_base_uri?: string
  metrics_provider?: Record<string, any>
  page_id?: number
  type?: string
  updated_at?: string
}

export interface MetricsProviderLoadMatch {
  id: string
  page_id: string
}

export interface MetricsProviderListMatch {
  page_id: string
}

export interface MetricsProviderCreateData {
  page_id: string
}

export interface MetricsProviderUpdateData {
  id: string
  page_id: string
}

export interface MetricsProviderRemoveMatch {
  id: string
  page_id: string
}

export interface Page {
  activity_score?: number
  allow_email_subscriber?: boolean
  allow_incident_subscriber?: boolean
  allow_page_subscriber?: boolean
  allow_rss_atom_feed?: boolean
  allow_sms_subscriber?: boolean
  allow_webhook_subscriber?: boolean
  branding?: string
  city?: string
  country?: string
  created_at?: string
  css_blue?: string
  css_body_background_color?: string
  css_border_color?: string
  css_font_color?: string
  css_graph_color?: string
  css_green?: string
  css_light_font_color?: string
  css_link_color?: string
  css_no_data?: string
  css_orange?: string
  css_red?: string
  css_yellow?: string
  domain?: string
  email_logo?: string
  favicon_logo?: string
  headline?: string
  hero_cover?: string
  hidden_from_search?: boolean
  id?: string
  ip_restriction?: string
  name?: string
  notifications_email_footer?: string
  notifications_from_email?: string
  page?: Record<string, any>
  page_description?: string
  state?: string
  subdomain?: string
  support_url?: string
  time_zone?: string
  transactional_logo?: string
  twitter_logo?: string
  twitter_username?: string
  updated_at?: string
  url?: string
  viewers_must_be_team_member?: boolean
}

export interface PageLoadMatch {
  id: string
}

export interface PageListMatch {
  activity_score?: number
  allow_email_subscriber?: boolean
  allow_incident_subscriber?: boolean
  allow_page_subscriber?: boolean
  allow_rss_atom_feed?: boolean
  allow_sms_subscriber?: boolean
  allow_webhook_subscriber?: boolean
  branding?: string
  city?: string
  country?: string
  created_at?: string
  css_blue?: string
  css_body_background_color?: string
  css_border_color?: string
  css_font_color?: string
  css_graph_color?: string
  css_green?: string
  css_light_font_color?: string
  css_link_color?: string
  css_no_data?: string
  css_orange?: string
  css_red?: string
  css_yellow?: string
  domain?: string
  email_logo?: string
  favicon_logo?: string
  headline?: string
  hero_cover?: string
  hidden_from_search?: boolean
  id?: string
  ip_restriction?: string
  name?: string
  notifications_email_footer?: string
  notifications_from_email?: string
  page?: Record<string, any>
  page_description?: string
  state?: string
  subdomain?: string
  support_url?: string
  time_zone?: string
  transactional_logo?: string
  twitter_logo?: string
  twitter_username?: string
  updated_at?: string
  url?: string
  viewers_must_be_team_member?: boolean
}

export interface PageUpdateData {
  id: string
}

export interface PageAccessGroup {
  component_id?: any[]
  created_at?: string
  external_identifier?: string
  id?: string
  metric_id?: any[]
  name?: string
  page_access_group?: Record<string, any>
  page_access_user_id?: any[]
  page_id?: string
  updated_at?: string
}

export interface PageAccessGroupLoadMatch {
  id: string
  page_id: string
}

export interface PageAccessGroupListMatch {
  id: string
}

export interface PageAccessGroupCreateData {
  id: string
  page_id: string
}

export interface PageAccessGroupUpdateData {
  id: string
  page_id: string
}

export interface PageAccessGroupRemoveMatch {
  component_id: string
  id: string
  page_id: string
}

export interface PageAccessUser {
  component_id: any[]
  created_at?: string
  email?: string
  external_login?: string
  id?: string
  metric_id: any[]
  page_access_group_id?: string
  page_access_user?: Record<string, any>
  page_id?: string
  updated_at?: string
}

export interface PageAccessUserLoadMatch {
  id: string
  page_id: string
}

export interface PageAccessUserListMatch {
  id: string
}

export interface PageAccessUserCreateData {
  id: string
  page_id: string
}

export interface PageAccessUserUpdateData {
  id: string
  page_id: string
}

export interface PageAccessUserRemoveMatch {
  component_id: string
  id: string
  page_id: string
  metric_id: string
}

export interface Permission {
  data?: Record<string, any>
  page?: Record<string, any>
}

export interface PermissionLoadMatch {
  id: string
  organization_id: string
}

export interface PermissionUpdateData {
  id: string
  organization_id: string
}

export interface Postmortem {
  body?: string
  body_draft?: string
  body_draft_updated_at?: string
  body_updated_at?: string
  created_at?: string
  custom_tweet?: string
  notify_subscriber?: boolean
  notify_twitter?: boolean
  postmortem: Record<string, any>
  preview_key?: string
  published_at?: string
  updated_at?: string
}

export interface PostmortemLoadMatch {
  incident_id: string
  page_id: string
}

export interface PostmortemUpdateData {
  incident_id: string
  page_id: string
}

export interface StatusEmbedConfig {
  incident_background_color?: string
  incident_text_color?: string
  maintenance_background_color?: string
  maintenance_text_color?: string
  page_id?: string
  position?: string
  status_embed_config?: Record<string, any>
}

export interface StatusEmbedConfigLoadMatch {
  page_id: string
}

export interface StatusEmbedConfigUpdateData {
  page_id: string
}

export interface Subscriber {
  component?: string
  component_id?: any[]
  created_at?: string
  display_phone_number?: string
  email?: string
  endpoint?: string
  id?: string
  integration_partner?: number
  mode?: string
  obfuscated_channel_name?: string
  page_access_user_id?: string
  phone_country?: string
  phone_number?: string
  purge_at?: string
  quarantined_at?: string
  skip_confirmation_notification?: boolean
  skip_unsubscription_notification?: boolean
  slack?: number
  sms?: number
  state?: string
  subscriber?: Record<string, any>
  team?: number
  type?: string
  webhook?: number
  workspace_name?: string
}

export interface SubscriberLoadMatch {
  id: string
  incident_id: string
  page_id: string
}

export interface SubscriberListMatch {
  page_id: string
  incident_id: string
}

export interface SubscriberCreateData {
  id: string
  page_id: string
  incident_id: string
}

export interface SubscriberUpdateData {
  id: string
  page_id: string
}

export interface SubscriberRemoveMatch {
  id: string
  incident_id: string
  page_id: string
}

export interface User {
  created_at?: string
  email?: string
  first_name?: string
  id?: string
  last_name?: string
  organization_id?: string
  updated_at?: string
  user: Record<string, any>
}

export interface UserListMatch {
  organization_id: string
}

export interface UserCreateData {
  organization_id: string
}

export interface UserRemoveMatch {
  id: string
  organization_id: string
}

