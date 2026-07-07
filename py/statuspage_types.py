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
    major_outage: int
    name: str
    only_show_if_degraded: bool
    page_id: str
    partial_outage: int
    position: int
    range_end: str
    range_start: str
    related_event: dict
    showcase: bool
    start_date: str
    status: str
    updated_at: str
    uptime_percentage: float
    warning: str


class ComponentLoadMatch(TypedDict):
    id: str
    page_id: str


class ComponentListMatchRequired(TypedDict):
    page_id: str


class ComponentListMatch(ComponentListMatchRequired, total=False):
    page_access_group_id: str
    page_access_user_id: str


class ComponentCreateData(TypedDict):
    page_id: str


class ComponentUpdateData(TypedDict):
    id: str
    page_id: str


class ComponentRemoveMatch(TypedDict):
    id: str
    page_id: str


class ComponentGroupUptime(TypedDict, total=False):
    id: str
    major_outage: int
    name: str
    partial_outage: int
    range_end: str
    range_start: str
    related_event: dict
    uptime_percentage: float
    warning: str


class ComponentGroupUptimeLoadMatch(TypedDict):
    id: str
    page_id: str


class GroupComponentRequired(TypedDict):
    component_group: dict


class GroupComponent(GroupComponentRequired, total=False):
    component: str
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


class GroupComponentCreateData(TypedDict):
    page_id: str


class GroupComponentUpdateData(TypedDict):
    id: str
    page_id: str


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
    component: list
    created_at: str
    id: str
    impact: str
    impact_override: str
    incident_impact: list
    incident_update: list
    metadata: dict
    monitoring_at: str
    name: str
    page_id: str
    postmortem_body: str
    postmortem_body_last_updated_at: str
    postmortem_ignored: bool
    postmortem_notified_subscriber: bool
    postmortem_notified_twitter: bool
    postmortem_published_at: bool
    reminder_interval: str
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


class IncidentCreateData(TypedDict):
    page_id: str


class IncidentUpdateData(TypedDict):
    id: str
    page_id: str


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
    component: list
    group_id: str
    id: str
    name: str
    should_send_notification: bool
    should_tweet: bool
    title: str
    update_status: str


class IncidentTemplateListMatch(TypedDict):
    page_id: str


class IncidentTemplateCreateData(TypedDict):
    page_id: str


class IncidentUpdate(TypedDict, total=False):
    affected_component: list
    body: str
    created_at: str
    custom_tweet: str
    deliver_notification: bool
    display_at: str
    id: str
    incident_id: str
    incident_update: dict
    status: str
    tweet_id: str
    twitter_updated_at: str
    updated_at: str
    wants_twitter_update: bool


class IncidentUpdateUpdateData(TypedDict):
    id: str
    incident_id: str
    page_id: str


class MetricRequired(TypedDict):
    data: dict


class Metric(MetricRequired, total=False):
    backfill_percentage: int
    backfilled: bool
    created_at: str
    decimal_place: int
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


class MetricCreateData(TypedDict):
    metrics_provider_id: str
    page_id: str


class MetricUpdateData(TypedDict):
    id: str
    page_id: str


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


class MetricsProviderCreateData(TypedDict):
    page_id: str


class MetricsProviderUpdateData(TypedDict):
    id: str
    page_id: str


class MetricsProviderRemoveMatch(TypedDict):
    id: str
    page_id: str


class Page(TypedDict, total=False):
    activity_score: float
    allow_email_subscriber: bool
    allow_incident_subscriber: bool
    allow_page_subscriber: bool
    allow_rss_atom_feed: bool
    allow_sms_subscriber: bool
    allow_webhook_subscriber: bool
    branding: str
    city: str
    country: str
    created_at: str
    css_blue: str
    css_body_background_color: str
    css_border_color: str
    css_font_color: str
    css_graph_color: str
    css_green: str
    css_light_font_color: str
    css_link_color: str
    css_no_data: str
    css_orange: str
    css_red: str
    css_yellow: str
    domain: str
    email_logo: str
    favicon_logo: str
    headline: str
    hero_cover: str
    hidden_from_search: bool
    id: str
    ip_restriction: str
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
    viewers_must_be_team_member: bool


class PageLoadMatch(TypedDict):
    id: str


class PageListMatch(TypedDict, total=False):
    activity_score: float
    allow_email_subscriber: bool
    allow_incident_subscriber: bool
    allow_page_subscriber: bool
    allow_rss_atom_feed: bool
    allow_sms_subscriber: bool
    allow_webhook_subscriber: bool
    branding: str
    city: str
    country: str
    created_at: str
    css_blue: str
    css_body_background_color: str
    css_border_color: str
    css_font_color: str
    css_graph_color: str
    css_green: str
    css_light_font_color: str
    css_link_color: str
    css_no_data: str
    css_orange: str
    css_red: str
    css_yellow: str
    domain: str
    email_logo: str
    favicon_logo: str
    headline: str
    hero_cover: str
    hidden_from_search: bool
    id: str
    ip_restriction: str
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
    viewers_must_be_team_member: bool


class PageUpdateData(TypedDict):
    id: str


class PageAccessGroup(TypedDict, total=False):
    component_id: list
    created_at: str
    external_identifier: str
    id: str
    metric_id: list
    name: str
    page_access_group: dict
    page_access_user_id: list
    page_id: str
    updated_at: str


class PageAccessGroupLoadMatch(TypedDict):
    id: str
    page_id: str


class PageAccessGroupListMatch(TypedDict):
    id: str


class PageAccessGroupCreateData(TypedDict):
    id: str


class PageAccessGroupUpdateData(TypedDict):
    id: str
    page_id: str


class PageAccessGroupRemoveMatchRequired(TypedDict):
    id: str
    page_id: str


class PageAccessGroupRemoveMatch(PageAccessGroupRemoveMatchRequired, total=False):
    component_id: str


class PageAccessUserRequired(TypedDict):
    component_id: list
    metric_id: list


class PageAccessUser(PageAccessUserRequired, total=False):
    created_at: str
    email: str
    external_login: str
    id: str
    page_access_group_id: str
    page_access_user: dict
    page_id: str
    updated_at: str


class PageAccessUserLoadMatch(TypedDict):
    id: str
    page_id: str


class PageAccessUserListMatch(TypedDict):
    id: str


class PageAccessUserCreateData(TypedDict):
    id: str


class PageAccessUserUpdateData(TypedDict):
    id: str
    page_id: str


class PageAccessUserRemoveMatchRequired(TypedDict):
    id: str
    page_id: str


class PageAccessUserRemoveMatch(PageAccessUserRemoveMatchRequired, total=False):
    component_id: str
    metric_id: str


class Permission(TypedDict, total=False):
    data: dict
    page: dict


class PermissionLoadMatch(TypedDict):
    id: str
    organization_id: str


class PermissionUpdateData(TypedDict):
    id: str
    organization_id: str


class PostmortemRequired(TypedDict):
    postmortem: dict


class Postmortem(PostmortemRequired, total=False):
    body: str
    body_draft: str
    body_draft_updated_at: str
    body_updated_at: str
    created_at: str
    custom_tweet: str
    notify_subscriber: bool
    notify_twitter: bool
    preview_key: str
    published_at: str
    updated_at: str


class PostmortemLoadMatch(TypedDict):
    incident_id: str
    page_id: str


class PostmortemUpdateData(TypedDict):
    incident_id: str
    page_id: str


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


class StatusEmbedConfigUpdateData(TypedDict):
    page_id: str


class Subscriber(TypedDict, total=False):
    component: str
    component_id: list
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
    team: int
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


class SubscriberCreateData(SubscriberCreateDataRequired, total=False):
    incident_id: str


class SubscriberUpdateData(TypedDict):
    id: str
    page_id: str


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


class UserCreateData(TypedDict):
    organization_id: str


class UserRemoveMatch(TypedDict):
    id: str
    organization_id: str
