# Typed models for the Statuspage SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Component(TypedDict, total=False):
    automation_email: str
    component: dict
    created_at: str
    description: str
    group: bool
    group_id: str
    id: str
    name: str
    only_show_if_degraded: bool
    page_id: str
    position: int
    showcase: bool
    start_date: str
    status: str
    updated_at: str


class ComponentLoadMatch(TypedDict):
    id: str
    page_id: str


class ComponentListMatchRequired(TypedDict):
    page_id: str


class ComponentListMatch(ComponentListMatchRequired, total=False):
    page_access_group_id: str
    page_access_user_id: str


class ComponentCreateDataRequired(TypedDict):
    page_id: str


class ComponentCreateData(ComponentCreateDataRequired, total=False):
    automation_email: str
    component: dict
    created_at: str
    description: str
    group: bool
    group_id: str
    id: str
    name: str
    only_show_if_degraded: bool
    position: int
    showcase: bool
    start_date: str
    status: str
    updated_at: str


class ComponentUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class ComponentUpdateData(ComponentUpdateDataRequired, total=False):
    automation_email: str
    component: dict
    created_at: str
    description: str
    group: bool
    group_id: str
    name: str
    only_show_if_degraded: bool
    position: int
    showcase: bool
    start_date: str
    status: str
    updated_at: str


class ComponentRemoveMatch(TypedDict):
    id: str
    page_id: str


class ComponentGroupUptime(TypedDict, total=False):
    component_id: str
    incidents: dict


class ComponentGroupUptimeLoadMatch(TypedDict):
    id: str
    page_id: str


class GroupComponentRequired(TypedDict):
    component_group: dict


class GroupComponent(GroupComponentRequired, total=False):
    components: str
    created_at: str
    description: str
    id: str
    name: str
    page_id: str
    position: str
    updated_at: str


class GroupComponentLoadMatch(TypedDict):
    id: str
    page_id: str


class GroupComponentListMatch(TypedDict):
    page_id: str


class GroupComponentCreateDataRequired(TypedDict):
    page_id: str
    component_group: dict


class GroupComponentCreateData(GroupComponentCreateDataRequired, total=False):
    components: str
    created_at: str
    description: str
    id: str
    name: str
    position: str
    updated_at: str


class GroupComponentUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class GroupComponentUpdateData(GroupComponentUpdateDataRequired, total=False):
    component_group: dict
    components: str
    created_at: str
    description: str
    name: str
    position: str
    updated_at: str


class GroupComponentRemoveMatch(TypedDict):
    id: str
    page_id: str


class IncidentRequired(TypedDict):
    incident: dict


class Incident(IncidentRequired, total=False):
    auto_transition_deliver_notifications_at_end: bool
    auto_transition_deliver_notifications_at_start: bool
    auto_transition_to_maintenance_state: bool
    auto_transition_to_operational_state: bool
    components: list
    created_at: str
    id: str
    impact: str
    impact_override: str
    incident_updates: list
    metadata: dict
    monitoring_at: str
    name: str
    page_id: str
    postmortem_body: str
    postmortem_body_last_updated_at: str
    postmortem_ignored: bool
    postmortem_notified_subscribers: bool
    postmortem_notified_twitter: bool
    postmortem_published_at: bool
    reminder_intervals: str
    resolved_at: str
    scheduled_auto_completed: bool
    scheduled_auto_in_progress: bool
    scheduled_for: str
    scheduled_remind_prior: bool
    scheduled_reminded_at: str
    scheduled_until: str
    shortlink: str
    status: str
    updated_at: str


class IncidentLoadMatch(TypedDict):
    id: str
    page_id: str


class IncidentListMatch(TypedDict):
    page_id: str


class IncidentCreateDataRequired(TypedDict):
    page_id: str
    incident: dict


class IncidentCreateData(IncidentCreateDataRequired, total=False):
    auto_transition_deliver_notifications_at_end: bool
    auto_transition_deliver_notifications_at_start: bool
    auto_transition_to_maintenance_state: bool
    auto_transition_to_operational_state: bool
    components: list
    created_at: str
    id: str
    impact: str
    impact_override: str
    incident_updates: list
    metadata: dict
    monitoring_at: str
    name: str
    postmortem_body: str
    postmortem_body_last_updated_at: str
    postmortem_ignored: bool
    postmortem_notified_subscribers: bool
    postmortem_notified_twitter: bool
    postmortem_published_at: bool
    reminder_intervals: str
    resolved_at: str
    scheduled_auto_completed: bool
    scheduled_auto_in_progress: bool
    scheduled_for: str
    scheduled_remind_prior: bool
    scheduled_reminded_at: str
    scheduled_until: str
    shortlink: str
    status: str
    updated_at: str


class IncidentUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class IncidentUpdateData(IncidentUpdateDataRequired, total=False):
    auto_transition_deliver_notifications_at_end: bool
    auto_transition_deliver_notifications_at_start: bool
    auto_transition_to_maintenance_state: bool
    auto_transition_to_operational_state: bool
    components: list
    created_at: str
    impact: str
    impact_override: str
    incident: dict
    incident_updates: list
    metadata: dict
    monitoring_at: str
    name: str
    postmortem_body: str
    postmortem_body_last_updated_at: str
    postmortem_ignored: bool
    postmortem_notified_subscribers: bool
    postmortem_notified_twitter: bool
    postmortem_published_at: bool
    reminder_intervals: str
    resolved_at: str
    scheduled_auto_completed: bool
    scheduled_auto_in_progress: bool
    scheduled_for: str
    scheduled_remind_prior: bool
    scheduled_reminded_at: str
    scheduled_until: str
    shortlink: str
    status: str
    updated_at: str


class IncidentRemoveMatch(TypedDict):
    id: str
    page_id: str


class IncidentPostmortem(TypedDict):
    pass


class IncidentPostmortemRemoveMatch(TypedDict):
    id: str
    page_id: str


class IncidentSubscriber(TypedDict):
    pass


class IncidentSubscriberCreateData(TypedDict):
    incident_id: str
    page_id: str
    subscriber_id: str


class IncidentTemplateRequired(TypedDict):
    template: dict


class IncidentTemplate(IncidentTemplateRequired, total=False):
    body: str
    components: list
    group_id: str
    id: str
    name: str
    should_send_notifications: bool
    should_tweet: bool
    title: str
    update_status: str


class IncidentTemplateListMatch(TypedDict):
    page_id: str


class IncidentTemplateCreateDataRequired(TypedDict):
    page_id: str
    template: dict


class IncidentTemplateCreateData(IncidentTemplateCreateDataRequired, total=False):
    body: str
    components: list
    group_id: str
    id: str
    name: str
    should_send_notifications: bool
    should_tweet: bool
    title: str
    update_status: str


class IncidentUpdate(TypedDict, total=False):
    affected_components: list
    body: str
    created_at: str
    custom_tweet: str
    deliver_notifications: bool
    display_at: str
    id: str
    incident_id: str
    incident_update: dict
    status: str
    tweet_id: str
    twitter_updated_at: str
    updated_at: str
    wants_twitter_update: bool


class IncidentUpdateUpdateDataRequired(TypedDict):
    id: str
    incident_id: str
    page_id: str


class IncidentUpdateUpdateData(IncidentUpdateUpdateDataRequired, total=False):
    affected_components: list
    body: str
    created_at: str
    custom_tweet: str
    deliver_notifications: bool
    display_at: str
    incident_update: dict
    status: str
    tweet_id: str
    twitter_updated_at: str
    updated_at: str
    wants_twitter_update: bool


class MetricRequired(TypedDict):
    data: dict


class Metric(MetricRequired, total=False):
    backfill_percentage: int
    backfilled: bool
    created_at: str
    decimal_places: int
    display: bool
    id: str
    last_fetched_at: str
    metric: dict
    metric_identifier: str
    metrics_provider_id: str
    most_recent_data_at: str
    name: str
    reference_name: str
    suffix: str
    tooltip_description: str
    updated_at: str
    y_axis_hidden: bool
    y_axis_max: float
    y_axis_min: float


class MetricLoadMatchRequired(TypedDict):
    page_id: str


class MetricLoadMatch(MetricLoadMatchRequired, total=False):
    metrics_provider_id: str
    id: str


class MetricListMatch(TypedDict):
    page_access_user_id: str
    page_id: str


class MetricCreateDataRequired(TypedDict):
    metrics_provider_id: str
    page_id: str
    data: dict


class MetricCreateData(MetricCreateDataRequired, total=False):
    backfill_percentage: int
    backfilled: bool
    created_at: str
    decimal_places: int
    display: bool
    id: str
    last_fetched_at: str
    metric: dict
    metric_identifier: str
    most_recent_data_at: str
    name: str
    reference_name: str
    suffix: str
    tooltip_description: str
    updated_at: str
    y_axis_hidden: bool
    y_axis_max: float
    y_axis_min: float


class MetricUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class MetricUpdateData(MetricUpdateDataRequired, total=False):
    backfill_percentage: int
    backfilled: bool
    created_at: str
    data: dict
    decimal_places: int
    display: bool
    last_fetched_at: str
    metric: dict
    metric_identifier: str
    metrics_provider_id: str
    most_recent_data_at: str
    name: str
    reference_name: str
    suffix: str
    tooltip_description: str
    updated_at: str
    y_axis_hidden: bool
    y_axis_max: float
    y_axis_min: float


class MetricRemoveMatch(TypedDict):
    id: str
    page_id: str


class MetricsProvider(TypedDict, total=False):
    created_at: str
    disabled: bool
    id: str
    last_revalidated_at: str
    metric_base_uri: str
    metrics_provider: dict
    page_id: int
    type: str
    updated_at: str


class MetricsProviderLoadMatch(TypedDict):
    id: str
    page_id: str


class MetricsProviderListMatch(TypedDict):
    page_id: str


class MetricsProviderCreateDataRequired(TypedDict):
    page_id: str


class MetricsProviderCreateData(MetricsProviderCreateDataRequired, total=False):
    created_at: str
    disabled: bool
    id: str
    last_revalidated_at: str
    metric_base_uri: str
    metrics_provider: dict
    type: str
    updated_at: str


class MetricsProviderUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class MetricsProviderUpdateData(MetricsProviderUpdateDataRequired, total=False):
    created_at: str
    disabled: bool
    last_revalidated_at: str
    metric_base_uri: str
    metrics_provider: dict
    type: str
    updated_at: str


class MetricsProviderRemoveMatch(TypedDict):
    id: str
    page_id: str


class Page(TypedDict, total=False):
    activity_score: float
    allow_email_subscribers: bool
    allow_incident_subscribers: bool
    allow_page_subscribers: bool
    allow_rss_atom_feeds: bool
    allow_sms_subscribers: bool
    allow_webhook_subscribers: bool
    branding: str
    city: str
    country: str
    created_at: str
    css_blues: str
    css_body_background_color: str
    css_border_color: str
    css_font_color: str
    css_graph_color: str
    css_greens: str
    css_light_font_color: str
    css_link_color: str
    css_no_data: str
    css_oranges: str
    css_reds: str
    css_yellows: str
    domain: str
    email_logo: str
    favicon_logo: str
    headline: str
    hero_cover: str
    hidden_from_search: bool
    id: str
    ip_restrictions: str
    name: str
    notifications_email_footer: str
    notifications_from_email: str
    page: dict
    page_description: str
    state: str
    subdomain: str
    support_url: str
    time_zone: str
    transactional_logo: str
    twitter_logo: str
    twitter_username: str
    updated_at: str
    url: str
    viewers_must_be_team_members: bool


class PageLoadMatch(TypedDict):
    id: str


class PageListMatch(TypedDict, total=False):
    activity_score: float
    allow_email_subscribers: bool
    allow_incident_subscribers: bool
    allow_page_subscribers: bool
    allow_rss_atom_feeds: bool
    allow_sms_subscribers: bool
    allow_webhook_subscribers: bool
    branding: str
    city: str
    country: str
    created_at: str
    css_blues: str
    css_body_background_color: str
    css_border_color: str
    css_font_color: str
    css_graph_color: str
    css_greens: str
    css_light_font_color: str
    css_link_color: str
    css_no_data: str
    css_oranges: str
    css_reds: str
    css_yellows: str
    domain: str
    email_logo: str
    favicon_logo: str
    headline: str
    hero_cover: str
    hidden_from_search: bool
    id: str
    ip_restrictions: str
    name: str
    notifications_email_footer: str
    notifications_from_email: str
    page: dict
    page_description: str
    state: str
    subdomain: str
    support_url: str
    time_zone: str
    transactional_logo: str
    twitter_logo: str
    twitter_username: str
    updated_at: str
    url: str
    viewers_must_be_team_members: bool


class PageUpdateDataRequired(TypedDict):
    id: str


class PageUpdateData(PageUpdateDataRequired, total=False):
    activity_score: float
    allow_email_subscribers: bool
    allow_incident_subscribers: bool
    allow_page_subscribers: bool
    allow_rss_atom_feeds: bool
    allow_sms_subscribers: bool
    allow_webhook_subscribers: bool
    branding: str
    city: str
    country: str
    created_at: str
    css_blues: str
    css_body_background_color: str
    css_border_color: str
    css_font_color: str
    css_graph_color: str
    css_greens: str
    css_light_font_color: str
    css_link_color: str
    css_no_data: str
    css_oranges: str
    css_reds: str
    css_yellows: str
    domain: str
    email_logo: str
    favicon_logo: str
    headline: str
    hero_cover: str
    hidden_from_search: bool
    ip_restrictions: str
    name: str
    notifications_email_footer: str
    notifications_from_email: str
    page: dict
    page_description: str
    state: str
    subdomain: str
    support_url: str
    time_zone: str
    transactional_logo: str
    twitter_logo: str
    twitter_username: str
    updated_at: str
    url: str
    viewers_must_be_team_members: bool


class PageAccessGroup(TypedDict, total=False):
    component_ids: list
    created_at: str
    external_identifier: str
    id: str
    metric_ids: list
    name: str
    page_access_group: dict
    page_access_user_ids: list
    page_id: str
    updated_at: str


class PageAccessGroupLoadMatch(TypedDict):
    id: str
    page_id: str


class PageAccessGroupListMatch(TypedDict):
    id: str


class PageAccessGroupCreateDataRequired(TypedDict):
    id: str


class PageAccessGroupCreateData(PageAccessGroupCreateDataRequired, total=False):
    component_ids: list
    created_at: str
    external_identifier: str
    metric_ids: list
    name: str
    page_access_group: dict
    page_access_user_ids: list
    page_id: str
    updated_at: str


class PageAccessGroupUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class PageAccessGroupUpdateData(PageAccessGroupUpdateDataRequired, total=False):
    component_ids: list
    created_at: str
    external_identifier: str
    metric_ids: list
    name: str
    page_access_group: dict
    page_access_user_ids: list
    updated_at: str


class PageAccessGroupRemoveMatchRequired(TypedDict):
    id: str
    page_id: str


class PageAccessGroupRemoveMatch(PageAccessGroupRemoveMatchRequired, total=False):
    component_id: str


class PageAccessUserRequired(TypedDict):
    component_ids: list
    metric_ids: list


class PageAccessUser(PageAccessUserRequired, total=False):
    created_at: str
    email: str
    external_login: str
    id: str
    page_access_group_id: str
    page_access_group_ids: str
    page_access_user: dict
    page_id: str
    updated_at: str


class PageAccessUserLoadMatch(TypedDict):
    id: str
    page_id: str


class PageAccessUserListMatch(TypedDict):
    id: str


class PageAccessUserCreateDataRequired(TypedDict):
    id: str
    component_ids: list
    metric_ids: list


class PageAccessUserCreateData(PageAccessUserCreateDataRequired, total=False):
    created_at: str
    email: str
    external_login: str
    page_access_group_id: str
    page_access_group_ids: str
    page_access_user: dict
    page_id: str
    updated_at: str


class PageAccessUserUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class PageAccessUserUpdateData(PageAccessUserUpdateDataRequired, total=False):
    component_ids: list
    created_at: str
    email: str
    external_login: str
    metric_ids: list
    page_access_group_id: str
    page_access_group_ids: str
    page_access_user: dict
    updated_at: str


class PageAccessUserRemoveMatchRequired(TypedDict):
    id: str
    page_id: str


class PageAccessUserRemoveMatch(PageAccessUserRemoveMatchRequired, total=False):
    component_id: str
    metric_id: str


class Permission(TypedDict, total=False):
    pages: dict
    user_id: str


class PermissionLoadMatch(TypedDict):
    id: str
    organization_id: str


class PermissionUpdateDataRequired(TypedDict):
    id: str
    organization_id: str


class PermissionUpdateData(PermissionUpdateDataRequired, total=False):
    pages: dict
    user_id: str


class PostmortemRequired(TypedDict):
    postmortem: dict


class Postmortem(PostmortemRequired, total=False):
    body: str
    body_draft: str
    body_draft_updated_at: str
    body_updated_at: str
    created_at: str
    custom_tweet: str
    notify_subscribers: bool
    notify_twitter: bool
    preview_key: str
    published_at: str
    updated_at: str


class PostmortemLoadMatch(TypedDict):
    incident_id: str
    page_id: str


class PostmortemUpdateDataRequired(TypedDict):
    incident_id: str
    page_id: str


class PostmortemUpdateData(PostmortemUpdateDataRequired, total=False):
    body: str
    body_draft: str
    body_draft_updated_at: str
    body_updated_at: str
    created_at: str
    custom_tweet: str
    notify_subscribers: bool
    notify_twitter: bool
    postmortem: dict
    preview_key: str
    published_at: str
    updated_at: str


class StatusEmbedConfig(TypedDict, total=False):
    incident_background_color: str
    incident_text_color: str
    maintenance_background_color: str
    maintenance_text_color: str
    page_id: str
    position: str
    status_embed_config: dict


class StatusEmbedConfigLoadMatch(TypedDict):
    page_id: str


class StatusEmbedConfigUpdateDataRequired(TypedDict):
    page_id: str


class StatusEmbedConfigUpdateData(StatusEmbedConfigUpdateDataRequired, total=False):
    incident_background_color: str
    incident_text_color: str
    maintenance_background_color: str
    maintenance_text_color: str
    position: str
    status_embed_config: dict


class SubscriberRequired(TypedDict):
    subscribers: str


class Subscriber(SubscriberRequired, total=False):
    component_ids: list
    components: str
    created_at: str
    display_phone_number: str
    email: str
    endpoint: str
    id: str
    integration_partner: int
    mode: str
    obfuscated_channel_name: str
    page_access_user_id: str
    phone_country: str
    phone_number: str
    purge_at: str
    quarantined_at: str
    skip_confirmation_notification: bool
    skip_unsubscription_notification: bool
    slack: int
    sms: int
    state: str
    subscriber: dict
    teams: int
    type: str
    webhook: int
    workspace_name: str


class SubscriberLoadMatchRequired(TypedDict):
    id: str
    page_id: str


class SubscriberLoadMatch(SubscriberLoadMatchRequired, total=False):
    incident_id: str


class SubscriberListMatchRequired(TypedDict):
    page_id: str


class SubscriberListMatch(SubscriberListMatchRequired, total=False):
    incident_id: str


class SubscriberCreateDataRequired(TypedDict):
    page_id: str
    subscribers: str


class SubscriberCreateData(SubscriberCreateDataRequired, total=False):
    incident_id: str
    component_ids: list
    components: str
    created_at: str
    display_phone_number: str
    email: str
    endpoint: str
    id: str
    integration_partner: int
    mode: str
    obfuscated_channel_name: str
    page_access_user_id: str
    phone_country: str
    phone_number: str
    purge_at: str
    quarantined_at: str
    skip_confirmation_notification: bool
    skip_unsubscription_notification: bool
    slack: int
    sms: int
    state: str
    subscriber: dict
    teams: int
    type: str
    webhook: int
    workspace_name: str


class SubscriberUpdateDataRequired(TypedDict):
    id: str
    page_id: str


class SubscriberUpdateData(SubscriberUpdateDataRequired, total=False):
    component_ids: list
    components: str
    created_at: str
    display_phone_number: str
    email: str
    endpoint: str
    integration_partner: int
    mode: str
    obfuscated_channel_name: str
    page_access_user_id: str
    phone_country: str
    phone_number: str
    purge_at: str
    quarantined_at: str
    skip_confirmation_notification: bool
    skip_unsubscription_notification: bool
    slack: int
    sms: int
    state: str
    subscriber: dict
    subscribers: str
    teams: int
    type: str
    webhook: int
    workspace_name: str


class SubscriberRemoveMatchRequired(TypedDict):
    id: str
    page_id: str


class SubscriberRemoveMatch(SubscriberRemoveMatchRequired, total=False):
    incident_id: str


class UserRequired(TypedDict):
    user: dict


class User(UserRequired, total=False):
    created_at: str
    email: str
    first_name: str
    id: str
    last_name: str
    organization_id: str
    updated_at: str


class UserListMatch(TypedDict):
    organization_id: str


class UserCreateDataRequired(TypedDict):
    organization_id: str
    user: dict


class UserCreateData(UserCreateDataRequired, total=False):
    created_at: str
    email: str
    first_name: str
    id: str
    last_name: str
    updated_at: str


class UserRemoveMatch(TypedDict):
    id: str
    organization_id: str
