<?php
declare(strict_types=1);

// Typed models for the Statuspage SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Component entity data model. */
class Component
{
    public ?string $automation_email = null;
    public ?array $component = null;
    public ?string $created_at = null;
    public ?string $description = null;
    public ?bool $group = null;
    public ?string $group_id = null;
    public ?string $id = null;
    public ?int $major_outage = null;
    public ?string $name = null;
    public ?bool $only_show_if_degraded = null;
    public ?string $page_id = null;
    public ?int $partial_outage = null;
    public ?int $position = null;
    public ?string $range_end = null;
    public ?string $range_start = null;
    public ?array $related_event = null;
    public ?bool $showcase = null;
    public ?string $start_date = null;
    public ?string $status = null;
    public ?string $updated_at = null;
    public ?float $uptime_percentage = null;
    public ?string $warning = null;
}

/** Request payload for Component#load. */
class ComponentLoadMatch
{
    public string $id;
    public string $page_id;
}

/** Request payload for Component#list. */
class ComponentListMatch
{
    public ?string $page_access_group_id = null;
    public string $page_id;
    public ?string $page_access_user_id = null;
}

/** Request payload for Component#create. */
class ComponentCreateData
{
    public string $page_id;
}

/** Request payload for Component#update. */
class ComponentUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for Component#remove. */
class ComponentRemoveMatch
{
    public string $id;
    public string $page_id;
}

/** ComponentGroupUptime entity data model. */
class ComponentGroupUptime
{
    public ?string $id = null;
    public ?int $major_outage = null;
    public ?string $name = null;
    public ?int $partial_outage = null;
    public ?string $range_end = null;
    public ?string $range_start = null;
    public ?array $related_event = null;
    public ?float $uptime_percentage = null;
    public ?string $warning = null;
}

/** Request payload for ComponentGroupUptime#load. */
class ComponentGroupUptimeLoadMatch
{
    public string $id;
    public string $page_id;
}

/** GroupComponent entity data model. */
class GroupComponent
{
    public ?string $component = null;
    public array $component_group;
    public ?string $created_at = null;
    public ?string $description = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?string $page_id = null;
    public ?string $position = null;
    public ?string $updated_at = null;
}

/** Request payload for GroupComponent#load. */
class GroupComponentLoadMatch
{
    public string $id;
    public string $page_id;
}

/** Request payload for GroupComponent#list. */
class GroupComponentListMatch
{
    public string $page_id;
}

/** Request payload for GroupComponent#create. */
class GroupComponentCreateData
{
    public string $page_id;
}

/** Request payload for GroupComponent#update. */
class GroupComponentUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for GroupComponent#remove. */
class GroupComponentRemoveMatch
{
    public string $id;
    public string $page_id;
}

/** Incident entity data model. */
class Incident
{
    public ?bool $auto_transition_deliver_notifications_at_end = null;
    public ?bool $auto_transition_deliver_notifications_at_start = null;
    public ?bool $auto_transition_to_maintenance_state = null;
    public ?bool $auto_transition_to_operational_state = null;
    public ?array $component = null;
    public ?string $created_at = null;
    public ?string $id = null;
    public ?string $impact = null;
    public ?string $impact_override = null;
    public array $incident;
    public ?array $incident_impact = null;
    public ?array $incident_update = null;
    public ?array $metadata = null;
    public ?string $monitoring_at = null;
    public ?string $name = null;
    public ?string $page_id = null;
    public ?string $postmortem_body = null;
    public ?string $postmortem_body_last_updated_at = null;
    public ?bool $postmortem_ignored = null;
    public ?bool $postmortem_notified_subscriber = null;
    public ?bool $postmortem_notified_twitter = null;
    public ?bool $postmortem_published_at = null;
    public ?string $reminder_interval = null;
    public ?string $resolved_at = null;
    public ?bool $scheduled_auto_completed = null;
    public ?bool $scheduled_auto_in_progress = null;
    public ?string $scheduled_for = null;
    public ?bool $scheduled_remind_prior = null;
    public ?string $scheduled_reminded_at = null;
    public ?string $scheduled_until = null;
    public ?string $shortlink = null;
    public ?string $status = null;
    public ?string $updated_at = null;
}

/** Request payload for Incident#load. */
class IncidentLoadMatch
{
    public string $id;
    public string $page_id;
}

/** Request payload for Incident#list. */
class IncidentListMatch
{
    public string $page_id;
}

/** Request payload for Incident#create. */
class IncidentCreateData
{
    public string $page_id;
}

/** Request payload for Incident#update. */
class IncidentUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for Incident#remove. */
class IncidentRemoveMatch
{
    public string $id;
    public string $page_id;
}

/** IncidentPostmortem entity data model. */
class IncidentPostmortem
{
}

/** Request payload for IncidentPostmortem#remove. */
class IncidentPostmortemRemoveMatch
{
    public string $id;
    public string $page_id;
}

/** IncidentSubscriber entity data model. */
class IncidentSubscriber
{
}

/** Request payload for IncidentSubscriber#create. */
class IncidentSubscriberCreateData
{
    public string $incident_id;
    public string $page_id;
    public string $subscriber_id;
}

/** IncidentTemplate entity data model. */
class IncidentTemplate
{
    public ?string $body = null;
    public ?array $component = null;
    public ?string $group_id = null;
    public ?string $id = null;
    public ?string $name = null;
    public ?bool $should_send_notification = null;
    public ?bool $should_tweet = null;
    public array $template;
    public ?string $title = null;
    public ?string $update_status = null;
}

/** Request payload for IncidentTemplate#list. */
class IncidentTemplateListMatch
{
    public string $page_id;
}

/** Request payload for IncidentTemplate#create. */
class IncidentTemplateCreateData
{
    public string $page_id;
}

/** IncidentUpdate entity data model. */
class IncidentUpdate
{
    public ?array $affected_component = null;
    public ?string $body = null;
    public ?string $created_at = null;
    public ?string $custom_tweet = null;
    public ?bool $deliver_notification = null;
    public ?string $display_at = null;
    public ?string $id = null;
    public ?string $incident_id = null;
    public ?array $incident_update = null;
    public ?string $status = null;
    public ?string $tweet_id = null;
    public ?string $twitter_updated_at = null;
    public ?string $updated_at = null;
    public ?bool $wants_twitter_update = null;
}

/** Request payload for IncidentUpdate#update. */
class IncidentUpdateUpdateData
{
    public string $id;
    public string $incident_id;
    public string $page_id;
}

/** Metric entity data model. */
class Metric
{
    public ?int $backfill_percentage = null;
    public ?bool $backfilled = null;
    public ?string $created_at = null;
    public array $data;
    public ?int $decimal_place = null;
    public ?bool $display = null;
    public ?string $id = null;
    public ?string $last_fetched_at = null;
    public ?array $metric = null;
    public ?string $metric_identifier = null;
    public ?string $metrics_provider_id = null;
    public ?string $most_recent_data_at = null;
    public ?string $name = null;
    public ?string $reference_name = null;
    public ?string $suffix = null;
    public ?string $tooltip_description = null;
    public ?string $updated_at = null;
    public ?bool $y_axis_hidden = null;
    public ?float $y_axis_max = null;
    public ?float $y_axis_min = null;
}

/** Request payload for Metric#load. */
class MetricLoadMatch
{
    public ?string $metrics_provider_id = null;
    public string $page_id;
    public ?string $id = null;
}

/** Request payload for Metric#list. */
class MetricListMatch
{
    public string $page_access_user_id;
    public string $page_id;
}

/** Request payload for Metric#create. */
class MetricCreateData
{
    public string $metrics_provider_id;
    public string $page_id;
}

/** Request payload for Metric#update. */
class MetricUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for Metric#remove. */
class MetricRemoveMatch
{
    public string $id;
    public string $page_id;
}

/** MetricsProvider entity data model. */
class MetricsProvider
{
    public ?string $created_at = null;
    public ?bool $disabled = null;
    public ?string $id = null;
    public ?string $last_revalidated_at = null;
    public ?string $metric_base_uri = null;
    public ?array $metrics_provider = null;
    public ?int $page_id = null;
    public ?string $type = null;
    public ?string $updated_at = null;
}

/** Request payload for MetricsProvider#load. */
class MetricsProviderLoadMatch
{
    public string $id;
    public string $page_id;
}

/** Request payload for MetricsProvider#list. */
class MetricsProviderListMatch
{
    public string $page_id;
}

/** Request payload for MetricsProvider#create. */
class MetricsProviderCreateData
{
    public string $page_id;
}

/** Request payload for MetricsProvider#update. */
class MetricsProviderUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for MetricsProvider#remove. */
class MetricsProviderRemoveMatch
{
    public string $id;
    public string $page_id;
}

/** Page entity data model. */
class Page
{
    public ?float $activity_score = null;
    public ?bool $allow_email_subscriber = null;
    public ?bool $allow_incident_subscriber = null;
    public ?bool $allow_page_subscriber = null;
    public ?bool $allow_rss_atom_feed = null;
    public ?bool $allow_sms_subscriber = null;
    public ?bool $allow_webhook_subscriber = null;
    public ?string $branding = null;
    public ?string $city = null;
    public ?string $country = null;
    public ?string $created_at = null;
    public ?string $css_blue = null;
    public ?string $css_body_background_color = null;
    public ?string $css_border_color = null;
    public ?string $css_font_color = null;
    public ?string $css_graph_color = null;
    public ?string $css_green = null;
    public ?string $css_light_font_color = null;
    public ?string $css_link_color = null;
    public ?string $css_no_data = null;
    public ?string $css_orange = null;
    public ?string $css_red = null;
    public ?string $css_yellow = null;
    public ?string $domain = null;
    public ?string $email_logo = null;
    public ?string $favicon_logo = null;
    public ?string $headline = null;
    public ?string $hero_cover = null;
    public ?bool $hidden_from_search = null;
    public ?string $id = null;
    public ?string $ip_restriction = null;
    public ?string $name = null;
    public ?string $notifications_email_footer = null;
    public ?string $notifications_from_email = null;
    public ?array $page = null;
    public ?string $page_description = null;
    public ?string $state = null;
    public ?string $subdomain = null;
    public ?string $support_url = null;
    public ?string $time_zone = null;
    public ?string $transactional_logo = null;
    public ?string $twitter_logo = null;
    public ?string $twitter_username = null;
    public ?string $updated_at = null;
    public ?string $url = null;
    public ?bool $viewers_must_be_team_member = null;
}

/** Request payload for Page#load. */
class PageLoadMatch
{
    public string $id;
}

/** Request payload for Page#list. */
class PageListMatch
{
    public ?float $activity_score = null;
    public ?bool $allow_email_subscriber = null;
    public ?bool $allow_incident_subscriber = null;
    public ?bool $allow_page_subscriber = null;
    public ?bool $allow_rss_atom_feed = null;
    public ?bool $allow_sms_subscriber = null;
    public ?bool $allow_webhook_subscriber = null;
    public ?string $branding = null;
    public ?string $city = null;
    public ?string $country = null;
    public ?string $created_at = null;
    public ?string $css_blue = null;
    public ?string $css_body_background_color = null;
    public ?string $css_border_color = null;
    public ?string $css_font_color = null;
    public ?string $css_graph_color = null;
    public ?string $css_green = null;
    public ?string $css_light_font_color = null;
    public ?string $css_link_color = null;
    public ?string $css_no_data = null;
    public ?string $css_orange = null;
    public ?string $css_red = null;
    public ?string $css_yellow = null;
    public ?string $domain = null;
    public ?string $email_logo = null;
    public ?string $favicon_logo = null;
    public ?string $headline = null;
    public ?string $hero_cover = null;
    public ?bool $hidden_from_search = null;
    public ?string $id = null;
    public ?string $ip_restriction = null;
    public ?string $name = null;
    public ?string $notifications_email_footer = null;
    public ?string $notifications_from_email = null;
    public ?array $page = null;
    public ?string $page_description = null;
    public ?string $state = null;
    public ?string $subdomain = null;
    public ?string $support_url = null;
    public ?string $time_zone = null;
    public ?string $transactional_logo = null;
    public ?string $twitter_logo = null;
    public ?string $twitter_username = null;
    public ?string $updated_at = null;
    public ?string $url = null;
    public ?bool $viewers_must_be_team_member = null;
}

/** Request payload for Page#update. */
class PageUpdateData
{
    public string $id;
}

/** PageAccessGroup entity data model. */
class PageAccessGroup
{
    public ?array $component_id = null;
    public ?string $created_at = null;
    public ?string $external_identifier = null;
    public ?string $id = null;
    public ?array $metric_id = null;
    public ?string $name = null;
    public ?array $page_access_group = null;
    public ?array $page_access_user_id = null;
    public ?string $page_id = null;
    public ?string $updated_at = null;
}

/** Request payload for PageAccessGroup#load. */
class PageAccessGroupLoadMatch
{
    public string $id;
    public string $page_id;
}

/** Request payload for PageAccessGroup#list. */
class PageAccessGroupListMatch
{
    public string $id;
}

/** Request payload for PageAccessGroup#create. */
class PageAccessGroupCreateData
{
    public string $id;
}

/** Request payload for PageAccessGroup#update. */
class PageAccessGroupUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for PageAccessGroup#remove. */
class PageAccessGroupRemoveMatch
{
    public ?string $component_id = null;
    public string $id;
    public string $page_id;
}

/** PageAccessUser entity data model. */
class PageAccessUser
{
    public array $component_id;
    public ?string $created_at = null;
    public ?string $email = null;
    public ?string $external_login = null;
    public ?string $id = null;
    public array $metric_id;
    public ?string $page_access_group_id = null;
    public ?array $page_access_user = null;
    public ?string $page_id = null;
    public ?string $updated_at = null;
}

/** Request payload for PageAccessUser#load. */
class PageAccessUserLoadMatch
{
    public string $id;
    public string $page_id;
}

/** Request payload for PageAccessUser#list. */
class PageAccessUserListMatch
{
    public string $id;
}

/** Request payload for PageAccessUser#create. */
class PageAccessUserCreateData
{
    public string $id;
}

/** Request payload for PageAccessUser#update. */
class PageAccessUserUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for PageAccessUser#remove. */
class PageAccessUserRemoveMatch
{
    public ?string $component_id = null;
    public string $id;
    public string $page_id;
    public ?string $metric_id = null;
}

/** Permission entity data model. */
class Permission
{
    public ?array $data = null;
    public ?array $page = null;
}

/** Request payload for Permission#load. */
class PermissionLoadMatch
{
    public string $id;
    public string $organization_id;
}

/** Request payload for Permission#update. */
class PermissionUpdateData
{
    public string $id;
    public string $organization_id;
}

/** Postmortem entity data model. */
class Postmortem
{
    public ?string $body = null;
    public ?string $body_draft = null;
    public ?string $body_draft_updated_at = null;
    public ?string $body_updated_at = null;
    public ?string $created_at = null;
    public ?string $custom_tweet = null;
    public ?bool $notify_subscriber = null;
    public ?bool $notify_twitter = null;
    public array $postmortem;
    public ?string $preview_key = null;
    public ?string $published_at = null;
    public ?string $updated_at = null;
}

/** Request payload for Postmortem#load. */
class PostmortemLoadMatch
{
    public string $incident_id;
    public string $page_id;
}

/** Request payload for Postmortem#update. */
class PostmortemUpdateData
{
    public string $incident_id;
    public string $page_id;
}

/** StatusEmbedConfig entity data model. */
class StatusEmbedConfig
{
    public ?string $incident_background_color = null;
    public ?string $incident_text_color = null;
    public ?string $maintenance_background_color = null;
    public ?string $maintenance_text_color = null;
    public ?string $page_id = null;
    public ?string $position = null;
    public ?array $status_embed_config = null;
}

/** Request payload for StatusEmbedConfig#load. */
class StatusEmbedConfigLoadMatch
{
    public string $page_id;
}

/** Request payload for StatusEmbedConfig#update. */
class StatusEmbedConfigUpdateData
{
    public string $page_id;
}

/** Subscriber entity data model. */
class Subscriber
{
    public ?string $component = null;
    public ?array $component_id = null;
    public ?string $created_at = null;
    public ?string $display_phone_number = null;
    public ?string $email = null;
    public ?string $endpoint = null;
    public ?string $id = null;
    public ?int $integration_partner = null;
    public ?string $mode = null;
    public ?string $obfuscated_channel_name = null;
    public ?string $page_access_user_id = null;
    public ?string $phone_country = null;
    public ?string $phone_number = null;
    public ?string $purge_at = null;
    public ?string $quarantined_at = null;
    public ?bool $skip_confirmation_notification = null;
    public ?bool $skip_unsubscription_notification = null;
    public ?int $slack = null;
    public ?int $sms = null;
    public ?string $state = null;
    public ?array $subscriber = null;
    public ?int $team = null;
    public ?string $type = null;
    public ?int $webhook = null;
    public ?string $workspace_name = null;
}

/** Request payload for Subscriber#load. */
class SubscriberLoadMatch
{
    public string $id;
    public ?string $incident_id = null;
    public string $page_id;
}

/** Request payload for Subscriber#list. */
class SubscriberListMatch
{
    public string $page_id;
    public ?string $incident_id = null;
}

/** Request payload for Subscriber#create. */
class SubscriberCreateData
{
    public ?string $incident_id = null;
    public string $page_id;
}

/** Request payload for Subscriber#update. */
class SubscriberUpdateData
{
    public string $id;
    public string $page_id;
}

/** Request payload for Subscriber#remove. */
class SubscriberRemoveMatch
{
    public string $id;
    public ?string $incident_id = null;
    public string $page_id;
}

/** User entity data model. */
class User
{
    public ?string $created_at = null;
    public ?string $email = null;
    public ?string $first_name = null;
    public ?string $id = null;
    public ?string $last_name = null;
    public ?string $organization_id = null;
    public ?string $updated_at = null;
    public array $user;
}

/** Request payload for User#list. */
class UserListMatch
{
    public string $organization_id;
}

/** Request payload for User#create. */
class UserCreateData
{
    public string $organization_id;
}

/** Request payload for User#remove. */
class UserRemoveMatch
{
    public string $id;
    public string $organization_id;
}

