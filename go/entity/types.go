// Typed models for the Statuspage SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/statuspage-sdk/go/core"
)

// Component is the typed data model for the component entity.
type Component struct {
	AutomationEmail *string `json:"automation_email,omitempty"`
	Component *map[string]any `json:"component,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Group *bool `json:"group,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OnlyShowIfDegraded *bool `json:"only_show_if_degraded,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	Position *int `json:"position,omitempty"`
	Showcase *bool `json:"showcase,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// ComponentLoadMatch is the typed request payload for Component.LoadTyped.
type ComponentLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// ComponentListMatch is the typed request payload for Component.ListTyped.
type ComponentListMatch struct {
	PageAccessGroupId *string `json:"page_access_group_id,omitempty"`
	PageId string `json:"page_id"`
	PageAccessUserId *string `json:"page_access_user_id,omitempty"`
}

// ComponentCreateData is the typed request payload for Component.CreateTyped.
type ComponentCreateData struct {
	PageId string `json:"page_id"`
	AutomationEmail *string `json:"automation_email,omitempty"`
	Component *map[string]any `json:"component,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Group *bool `json:"group,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	OnlyShowIfDegraded *bool `json:"only_show_if_degraded,omitempty"`
	Position *int `json:"position,omitempty"`
	Showcase *bool `json:"showcase,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// ComponentUpdateData is the typed request payload for Component.UpdateTyped.
type ComponentUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	AutomationEmail *string `json:"automation_email,omitempty"`
	Component *map[string]any `json:"component,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Group *bool `json:"group,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Name *string `json:"name,omitempty"`
	OnlyShowIfDegraded *bool `json:"only_show_if_degraded,omitempty"`
	Position *int `json:"position,omitempty"`
	Showcase *bool `json:"showcase,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// ComponentRemoveMatch is the typed request payload for Component.RemoveTyped.
type ComponentRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// ComponentGroupUptime is the typed data model for the component_group_uptime entity.
type ComponentGroupUptime struct {
	ComponentId *string `json:"component_id,omitempty"`
	Incidents *map[string]any `json:"incidents,omitempty"`
}

// ComponentGroupUptimeLoadMatch is the typed request payload for ComponentGroupUptime.LoadTyped.
type ComponentGroupUptimeLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// GroupComponent is the typed data model for the group_component entity.
type GroupComponent struct {
	ComponentGroup map[string]any `json:"component_group"`
	Components *string `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	Position *string `json:"position,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// GroupComponentLoadMatch is the typed request payload for GroupComponent.LoadTyped.
type GroupComponentLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// GroupComponentListMatch is the typed request payload for GroupComponent.ListTyped.
type GroupComponentListMatch struct {
	PageId string `json:"page_id"`
}

// GroupComponentCreateData is the typed request payload for GroupComponent.CreateTyped.
type GroupComponentCreateData struct {
	PageId string `json:"page_id"`
	ComponentGroup map[string]any `json:"component_group"`
	Components *string `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Position *string `json:"position,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// GroupComponentUpdateData is the typed request payload for GroupComponent.UpdateTyped.
type GroupComponentUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	ComponentGroup *map[string]any `json:"component_group,omitempty"`
	Components *string `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Name *string `json:"name,omitempty"`
	Position *string `json:"position,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// GroupComponentRemoveMatch is the typed request payload for GroupComponent.RemoveTyped.
type GroupComponentRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// Incident is the typed data model for the incident entity.
type Incident struct {
	AutoTransitionDeliverNotificationsAtEnd *bool `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	AutoTransitionDeliverNotificationsAtStart *bool `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	AutoTransitionToMaintenanceState *bool `json:"auto_transition_to_maintenance_state,omitempty"`
	AutoTransitionToOperationalState *bool `json:"auto_transition_to_operational_state,omitempty"`
	Components *[]any `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Impact *string `json:"impact,omitempty"`
	ImpactOverride *string `json:"impact_override,omitempty"`
	Incident map[string]any `json:"incident"`
	IncidentUpdates *[]any `json:"incident_updates,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	MonitoringAt *string `json:"monitoring_at,omitempty"`
	Name *string `json:"name,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	PostmortemBody *string `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt *string `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored *bool `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers *bool `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter *bool `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt *bool `json:"postmortem_published_at,omitempty"`
	ReminderIntervals *string `json:"reminder_intervals,omitempty"`
	ResolvedAt *string `json:"resolved_at,omitempty"`
	ScheduledAutoCompleted *bool `json:"scheduled_auto_completed,omitempty"`
	ScheduledAutoInProgress *bool `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledFor *string `json:"scheduled_for,omitempty"`
	ScheduledRemindPrior *bool `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt *string `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil *string `json:"scheduled_until,omitempty"`
	Shortlink *string `json:"shortlink,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// IncidentLoadMatch is the typed request payload for Incident.LoadTyped.
type IncidentLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// IncidentListMatch is the typed request payload for Incident.ListTyped.
type IncidentListMatch struct {
	PageId string `json:"page_id"`
}

// IncidentCreateData is the typed request payload for Incident.CreateTyped.
type IncidentCreateData struct {
	PageId string `json:"page_id"`
	AutoTransitionDeliverNotificationsAtEnd *bool `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	AutoTransitionDeliverNotificationsAtStart *bool `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	AutoTransitionToMaintenanceState *bool `json:"auto_transition_to_maintenance_state,omitempty"`
	AutoTransitionToOperationalState *bool `json:"auto_transition_to_operational_state,omitempty"`
	Components *[]any `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Impact *string `json:"impact,omitempty"`
	ImpactOverride *string `json:"impact_override,omitempty"`
	Incident map[string]any `json:"incident"`
	IncidentUpdates *[]any `json:"incident_updates,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	MonitoringAt *string `json:"monitoring_at,omitempty"`
	Name *string `json:"name,omitempty"`
	PostmortemBody *string `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt *string `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored *bool `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers *bool `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter *bool `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt *bool `json:"postmortem_published_at,omitempty"`
	ReminderIntervals *string `json:"reminder_intervals,omitempty"`
	ResolvedAt *string `json:"resolved_at,omitempty"`
	ScheduledAutoCompleted *bool `json:"scheduled_auto_completed,omitempty"`
	ScheduledAutoInProgress *bool `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledFor *string `json:"scheduled_for,omitempty"`
	ScheduledRemindPrior *bool `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt *string `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil *string `json:"scheduled_until,omitempty"`
	Shortlink *string `json:"shortlink,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// IncidentUpdateData is the typed request payload for Incident.UpdateTyped.
type IncidentUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	AutoTransitionDeliverNotificationsAtEnd *bool `json:"auto_transition_deliver_notifications_at_end,omitempty"`
	AutoTransitionDeliverNotificationsAtStart *bool `json:"auto_transition_deliver_notifications_at_start,omitempty"`
	AutoTransitionToMaintenanceState *bool `json:"auto_transition_to_maintenance_state,omitempty"`
	AutoTransitionToOperationalState *bool `json:"auto_transition_to_operational_state,omitempty"`
	Components *[]any `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Impact *string `json:"impact,omitempty"`
	ImpactOverride *string `json:"impact_override,omitempty"`
	Incident *map[string]any `json:"incident,omitempty"`
	IncidentUpdates *[]any `json:"incident_updates,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	MonitoringAt *string `json:"monitoring_at,omitempty"`
	Name *string `json:"name,omitempty"`
	PostmortemBody *string `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt *string `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored *bool `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscribers *bool `json:"postmortem_notified_subscribers,omitempty"`
	PostmortemNotifiedTwitter *bool `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt *bool `json:"postmortem_published_at,omitempty"`
	ReminderIntervals *string `json:"reminder_intervals,omitempty"`
	ResolvedAt *string `json:"resolved_at,omitempty"`
	ScheduledAutoCompleted *bool `json:"scheduled_auto_completed,omitempty"`
	ScheduledAutoInProgress *bool `json:"scheduled_auto_in_progress,omitempty"`
	ScheduledFor *string `json:"scheduled_for,omitempty"`
	ScheduledRemindPrior *bool `json:"scheduled_remind_prior,omitempty"`
	ScheduledRemindedAt *string `json:"scheduled_reminded_at,omitempty"`
	ScheduledUntil *string `json:"scheduled_until,omitempty"`
	Shortlink *string `json:"shortlink,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// IncidentRemoveMatch is the typed request payload for Incident.RemoveTyped.
type IncidentRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// IncidentPostmortem is the typed data model for the incident_postmortem entity.
type IncidentPostmortem struct {
}

// IncidentPostmortemRemoveMatch is the typed request payload for IncidentPostmortem.RemoveTyped.
type IncidentPostmortemRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// IncidentSubscriber is the typed data model for the incident_subscriber entity.
type IncidentSubscriber struct {
}

// IncidentSubscriberCreateData is the typed request payload for IncidentSubscriber.CreateTyped.
type IncidentSubscriberCreateData struct {
	IncidentId string `json:"incident_id"`
	PageId string `json:"page_id"`
	SubscriberId string `json:"subscriber_id"`
}

// IncidentTemplate is the typed data model for the incident_template entity.
type IncidentTemplate struct {
	Body *string `json:"body,omitempty"`
	Components *[]any `json:"components,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ShouldSendNotifications *bool `json:"should_send_notifications,omitempty"`
	ShouldTweet *bool `json:"should_tweet,omitempty"`
	Template map[string]any `json:"template"`
	Title *string `json:"title,omitempty"`
	UpdateStatus *string `json:"update_status,omitempty"`
}

// IncidentTemplateListMatch is the typed request payload for IncidentTemplate.ListTyped.
type IncidentTemplateListMatch struct {
	PageId string `json:"page_id"`
}

// IncidentTemplateCreateData is the typed request payload for IncidentTemplate.CreateTyped.
type IncidentTemplateCreateData struct {
	PageId string `json:"page_id"`
	Body *string `json:"body,omitempty"`
	Components *[]any `json:"components,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ShouldSendNotifications *bool `json:"should_send_notifications,omitempty"`
	ShouldTweet *bool `json:"should_tweet,omitempty"`
	Template map[string]any `json:"template"`
	Title *string `json:"title,omitempty"`
	UpdateStatus *string `json:"update_status,omitempty"`
}

// IncidentUpdate is the typed data model for the incident_update entity.
type IncidentUpdate struct {
	AffectedComponents *[]any `json:"affected_components,omitempty"`
	Body *string `json:"body,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomTweet *string `json:"custom_tweet,omitempty"`
	DeliverNotifications *bool `json:"deliver_notifications,omitempty"`
	DisplayAt *string `json:"display_at,omitempty"`
	Id *string `json:"id,omitempty"`
	IncidentId *string `json:"incident_id,omitempty"`
	IncidentUpdate *map[string]any `json:"incident_update,omitempty"`
	Status *string `json:"status,omitempty"`
	TweetId *string `json:"tweet_id,omitempty"`
	TwitterUpdatedAt *string `json:"twitter_updated_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	WantsTwitterUpdate *bool `json:"wants_twitter_update,omitempty"`
}

// IncidentUpdateUpdateData is the typed request payload for IncidentUpdate.UpdateTyped.
type IncidentUpdateUpdateData struct {
	Id string `json:"id"`
	IncidentId string `json:"incident_id"`
	PageId string `json:"page_id"`
	AffectedComponents *[]any `json:"affected_components,omitempty"`
	Body *string `json:"body,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomTweet *string `json:"custom_tweet,omitempty"`
	DeliverNotifications *bool `json:"deliver_notifications,omitempty"`
	DisplayAt *string `json:"display_at,omitempty"`
	IncidentUpdate *map[string]any `json:"incident_update,omitempty"`
	Status *string `json:"status,omitempty"`
	TweetId *string `json:"tweet_id,omitempty"`
	TwitterUpdatedAt *string `json:"twitter_updated_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	WantsTwitterUpdate *bool `json:"wants_twitter_update,omitempty"`
}

// Metric is the typed data model for the metric entity.
type Metric struct {
	BackfillPercentage *int `json:"backfill_percentage,omitempty"`
	Backfilled *bool `json:"backfilled,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Data map[string]any `json:"data"`
	DecimalPlaces *int `json:"decimal_places,omitempty"`
	Display *bool `json:"display,omitempty"`
	Id *string `json:"id,omitempty"`
	LastFetchedAt *string `json:"last_fetched_at,omitempty"`
	Metric *map[string]any `json:"metric,omitempty"`
	MetricIdentifier *string `json:"metric_identifier,omitempty"`
	MetricsProviderId *string `json:"metrics_provider_id,omitempty"`
	MostRecentDataAt *string `json:"most_recent_data_at,omitempty"`
	Name *string `json:"name,omitempty"`
	ReferenceName *string `json:"reference_name,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
	TooltipDescription *string `json:"tooltip_description,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	YAxisHidden *bool `json:"y_axis_hidden,omitempty"`
	YAxisMax *float64 `json:"y_axis_max,omitempty"`
	YAxisMin *float64 `json:"y_axis_min,omitempty"`
}

// MetricLoadMatch is the typed request payload for Metric.LoadTyped.
type MetricLoadMatch struct {
	MetricsProviderId *string `json:"metrics_provider_id,omitempty"`
	PageId string `json:"page_id"`
	Id *string `json:"id,omitempty"`
}

// MetricListMatch is the typed request payload for Metric.ListTyped.
type MetricListMatch struct {
	PageAccessUserId string `json:"page_access_user_id"`
	PageId string `json:"page_id"`
}

// MetricCreateData is the typed request payload for Metric.CreateTyped.
type MetricCreateData struct {
	MetricsProviderId string `json:"metrics_provider_id"`
	PageId string `json:"page_id"`
	BackfillPercentage *int `json:"backfill_percentage,omitempty"`
	Backfilled *bool `json:"backfilled,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Data map[string]any `json:"data"`
	DecimalPlaces *int `json:"decimal_places,omitempty"`
	Display *bool `json:"display,omitempty"`
	Id *string `json:"id,omitempty"`
	LastFetchedAt *string `json:"last_fetched_at,omitempty"`
	Metric *map[string]any `json:"metric,omitempty"`
	MetricIdentifier *string `json:"metric_identifier,omitempty"`
	MostRecentDataAt *string `json:"most_recent_data_at,omitempty"`
	Name *string `json:"name,omitempty"`
	ReferenceName *string `json:"reference_name,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
	TooltipDescription *string `json:"tooltip_description,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	YAxisHidden *bool `json:"y_axis_hidden,omitempty"`
	YAxisMax *float64 `json:"y_axis_max,omitempty"`
	YAxisMin *float64 `json:"y_axis_min,omitempty"`
}

// MetricUpdateData is the typed request payload for Metric.UpdateTyped.
type MetricUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	BackfillPercentage *int `json:"backfill_percentage,omitempty"`
	Backfilled *bool `json:"backfilled,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Data *map[string]any `json:"data,omitempty"`
	DecimalPlaces *int `json:"decimal_places,omitempty"`
	Display *bool `json:"display,omitempty"`
	LastFetchedAt *string `json:"last_fetched_at,omitempty"`
	Metric *map[string]any `json:"metric,omitempty"`
	MetricIdentifier *string `json:"metric_identifier,omitempty"`
	MetricsProviderId *string `json:"metrics_provider_id,omitempty"`
	MostRecentDataAt *string `json:"most_recent_data_at,omitempty"`
	Name *string `json:"name,omitempty"`
	ReferenceName *string `json:"reference_name,omitempty"`
	Suffix *string `json:"suffix,omitempty"`
	TooltipDescription *string `json:"tooltip_description,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	YAxisHidden *bool `json:"y_axis_hidden,omitempty"`
	YAxisMax *float64 `json:"y_axis_max,omitempty"`
	YAxisMin *float64 `json:"y_axis_min,omitempty"`
}

// MetricRemoveMatch is the typed request payload for Metric.RemoveTyped.
type MetricRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// MetricsProvider is the typed data model for the metrics_provider entity.
type MetricsProvider struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Id *string `json:"id,omitempty"`
	LastRevalidatedAt *string `json:"last_revalidated_at,omitempty"`
	MetricBaseUri *string `json:"metric_base_uri,omitempty"`
	MetricsProvider *map[string]any `json:"metrics_provider,omitempty"`
	PageId *int `json:"page_id,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// MetricsProviderLoadMatch is the typed request payload for MetricsProvider.LoadTyped.
type MetricsProviderLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// MetricsProviderListMatch is the typed request payload for MetricsProvider.ListTyped.
type MetricsProviderListMatch struct {
	PageId string `json:"page_id"`
}

// MetricsProviderCreateData is the typed request payload for MetricsProvider.CreateTyped.
type MetricsProviderCreateData struct {
	PageId string `json:"page_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	Id *string `json:"id,omitempty"`
	LastRevalidatedAt *string `json:"last_revalidated_at,omitempty"`
	MetricBaseUri *string `json:"metric_base_uri,omitempty"`
	MetricsProvider *map[string]any `json:"metrics_provider,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// MetricsProviderUpdateData is the typed request payload for MetricsProvider.UpdateTyped.
type MetricsProviderUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Disabled *bool `json:"disabled,omitempty"`
	LastRevalidatedAt *string `json:"last_revalidated_at,omitempty"`
	MetricBaseUri *string `json:"metric_base_uri,omitempty"`
	MetricsProvider *map[string]any `json:"metrics_provider,omitempty"`
	Type *string `json:"type,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// MetricsProviderRemoveMatch is the typed request payload for MetricsProvider.RemoveTyped.
type MetricsProviderRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// Page is the typed data model for the page entity.
type Page struct {
	ActivityScore *float64 `json:"activity_score,omitempty"`
	AllowEmailSubscribers *bool `json:"allow_email_subscribers,omitempty"`
	AllowIncidentSubscribers *bool `json:"allow_incident_subscribers,omitempty"`
	AllowPageSubscribers *bool `json:"allow_page_subscribers,omitempty"`
	AllowRssAtomFeeds *bool `json:"allow_rss_atom_feeds,omitempty"`
	AllowSmsSubscribers *bool `json:"allow_sms_subscribers,omitempty"`
	AllowWebhookSubscribers *bool `json:"allow_webhook_subscribers,omitempty"`
	Branding *string `json:"branding,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CssBlues *string `json:"css_blues,omitempty"`
	CssBodyBackgroundColor *string `json:"css_body_background_color,omitempty"`
	CssBorderColor *string `json:"css_border_color,omitempty"`
	CssFontColor *string `json:"css_font_color,omitempty"`
	CssGraphColor *string `json:"css_graph_color,omitempty"`
	CssGreens *string `json:"css_greens,omitempty"`
	CssLightFontColor *string `json:"css_light_font_color,omitempty"`
	CssLinkColor *string `json:"css_link_color,omitempty"`
	CssNoData *string `json:"css_no_data,omitempty"`
	CssOranges *string `json:"css_oranges,omitempty"`
	CssReds *string `json:"css_reds,omitempty"`
	CssYellows *string `json:"css_yellows,omitempty"`
	Domain *string `json:"domain,omitempty"`
	EmailLogo *string `json:"email_logo,omitempty"`
	FaviconLogo *string `json:"favicon_logo,omitempty"`
	Headline *string `json:"headline,omitempty"`
	HeroCover *string `json:"hero_cover,omitempty"`
	HiddenFromSearch *bool `json:"hidden_from_search,omitempty"`
	Id *string `json:"id,omitempty"`
	IpRestrictions *string `json:"ip_restrictions,omitempty"`
	Name *string `json:"name,omitempty"`
	NotificationsEmailFooter *string `json:"notifications_email_footer,omitempty"`
	NotificationsFromEmail *string `json:"notifications_from_email,omitempty"`
	Page *map[string]any `json:"page,omitempty"`
	PageDescription *string `json:"page_description,omitempty"`
	State *string `json:"state,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	SupportUrl *string `json:"support_url,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
	TransactionalLogo *string `json:"transactional_logo,omitempty"`
	TwitterLogo *string `json:"twitter_logo,omitempty"`
	TwitterUsername *string `json:"twitter_username,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
	ViewersMustBeTeamMembers *bool `json:"viewers_must_be_team_members,omitempty"`
}

// PageLoadMatch is the typed request payload for Page.LoadTyped.
type PageLoadMatch struct {
	Id string `json:"id"`
}

// PageListMatch is the typed request payload for Page.ListTyped.
type PageListMatch struct {
	ActivityScore *float64 `json:"activity_score,omitempty"`
	AllowEmailSubscribers *bool `json:"allow_email_subscribers,omitempty"`
	AllowIncidentSubscribers *bool `json:"allow_incident_subscribers,omitempty"`
	AllowPageSubscribers *bool `json:"allow_page_subscribers,omitempty"`
	AllowRssAtomFeeds *bool `json:"allow_rss_atom_feeds,omitempty"`
	AllowSmsSubscribers *bool `json:"allow_sms_subscribers,omitempty"`
	AllowWebhookSubscribers *bool `json:"allow_webhook_subscribers,omitempty"`
	Branding *string `json:"branding,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CssBlues *string `json:"css_blues,omitempty"`
	CssBodyBackgroundColor *string `json:"css_body_background_color,omitempty"`
	CssBorderColor *string `json:"css_border_color,omitempty"`
	CssFontColor *string `json:"css_font_color,omitempty"`
	CssGraphColor *string `json:"css_graph_color,omitempty"`
	CssGreens *string `json:"css_greens,omitempty"`
	CssLightFontColor *string `json:"css_light_font_color,omitempty"`
	CssLinkColor *string `json:"css_link_color,omitempty"`
	CssNoData *string `json:"css_no_data,omitempty"`
	CssOranges *string `json:"css_oranges,omitempty"`
	CssReds *string `json:"css_reds,omitempty"`
	CssYellows *string `json:"css_yellows,omitempty"`
	Domain *string `json:"domain,omitempty"`
	EmailLogo *string `json:"email_logo,omitempty"`
	FaviconLogo *string `json:"favicon_logo,omitempty"`
	Headline *string `json:"headline,omitempty"`
	HeroCover *string `json:"hero_cover,omitempty"`
	HiddenFromSearch *bool `json:"hidden_from_search,omitempty"`
	Id *string `json:"id,omitempty"`
	IpRestrictions *string `json:"ip_restrictions,omitempty"`
	Name *string `json:"name,omitempty"`
	NotificationsEmailFooter *string `json:"notifications_email_footer,omitempty"`
	NotificationsFromEmail *string `json:"notifications_from_email,omitempty"`
	Page *map[string]any `json:"page,omitempty"`
	PageDescription *string `json:"page_description,omitempty"`
	State *string `json:"state,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	SupportUrl *string `json:"support_url,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
	TransactionalLogo *string `json:"transactional_logo,omitempty"`
	TwitterLogo *string `json:"twitter_logo,omitempty"`
	TwitterUsername *string `json:"twitter_username,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
	ViewersMustBeTeamMembers *bool `json:"viewers_must_be_team_members,omitempty"`
}

// PageUpdateData is the typed request payload for Page.UpdateTyped.
type PageUpdateData struct {
	Id string `json:"id"`
	ActivityScore *float64 `json:"activity_score,omitempty"`
	AllowEmailSubscribers *bool `json:"allow_email_subscribers,omitempty"`
	AllowIncidentSubscribers *bool `json:"allow_incident_subscribers,omitempty"`
	AllowPageSubscribers *bool `json:"allow_page_subscribers,omitempty"`
	AllowRssAtomFeeds *bool `json:"allow_rss_atom_feeds,omitempty"`
	AllowSmsSubscribers *bool `json:"allow_sms_subscribers,omitempty"`
	AllowWebhookSubscribers *bool `json:"allow_webhook_subscribers,omitempty"`
	Branding *string `json:"branding,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CssBlues *string `json:"css_blues,omitempty"`
	CssBodyBackgroundColor *string `json:"css_body_background_color,omitempty"`
	CssBorderColor *string `json:"css_border_color,omitempty"`
	CssFontColor *string `json:"css_font_color,omitempty"`
	CssGraphColor *string `json:"css_graph_color,omitempty"`
	CssGreens *string `json:"css_greens,omitempty"`
	CssLightFontColor *string `json:"css_light_font_color,omitempty"`
	CssLinkColor *string `json:"css_link_color,omitempty"`
	CssNoData *string `json:"css_no_data,omitempty"`
	CssOranges *string `json:"css_oranges,omitempty"`
	CssReds *string `json:"css_reds,omitempty"`
	CssYellows *string `json:"css_yellows,omitempty"`
	Domain *string `json:"domain,omitempty"`
	EmailLogo *string `json:"email_logo,omitempty"`
	FaviconLogo *string `json:"favicon_logo,omitempty"`
	Headline *string `json:"headline,omitempty"`
	HeroCover *string `json:"hero_cover,omitempty"`
	HiddenFromSearch *bool `json:"hidden_from_search,omitempty"`
	IpRestrictions *string `json:"ip_restrictions,omitempty"`
	Name *string `json:"name,omitempty"`
	NotificationsEmailFooter *string `json:"notifications_email_footer,omitempty"`
	NotificationsFromEmail *string `json:"notifications_from_email,omitempty"`
	Page *map[string]any `json:"page,omitempty"`
	PageDescription *string `json:"page_description,omitempty"`
	State *string `json:"state,omitempty"`
	Subdomain *string `json:"subdomain,omitempty"`
	SupportUrl *string `json:"support_url,omitempty"`
	TimeZone *string `json:"time_zone,omitempty"`
	TransactionalLogo *string `json:"transactional_logo,omitempty"`
	TwitterLogo *string `json:"twitter_logo,omitempty"`
	TwitterUsername *string `json:"twitter_username,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	Url *string `json:"url,omitempty"`
	ViewersMustBeTeamMembers *bool `json:"viewers_must_be_team_members,omitempty"`
}

// PageAccessGroup is the typed data model for the page_access_group entity.
type PageAccessGroup struct {
	ComponentIds *[]any `json:"component_ids,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	Id *string `json:"id,omitempty"`
	MetricIds *[]any `json:"metric_ids,omitempty"`
	Name *string `json:"name,omitempty"`
	PageAccessGroup *map[string]any `json:"page_access_group,omitempty"`
	PageAccessUserIds *[]any `json:"page_access_user_ids,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PageAccessGroupLoadMatch is the typed request payload for PageAccessGroup.LoadTyped.
type PageAccessGroupLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// PageAccessGroupListMatch is the typed request payload for PageAccessGroup.ListTyped.
type PageAccessGroupListMatch struct {
	Id string `json:"id"`
}

// PageAccessGroupCreateData is the typed request payload for PageAccessGroup.CreateTyped.
type PageAccessGroupCreateData struct {
	Id string `json:"id"`
	ComponentIds *[]any `json:"component_ids,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	MetricIds *[]any `json:"metric_ids,omitempty"`
	Name *string `json:"name,omitempty"`
	PageAccessGroup *map[string]any `json:"page_access_group,omitempty"`
	PageAccessUserIds *[]any `json:"page_access_user_ids,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PageAccessGroupUpdateData is the typed request payload for PageAccessGroup.UpdateTyped.
type PageAccessGroupUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	ComponentIds *[]any `json:"component_ids,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	MetricIds *[]any `json:"metric_ids,omitempty"`
	Name *string `json:"name,omitempty"`
	PageAccessGroup *map[string]any `json:"page_access_group,omitempty"`
	PageAccessUserIds *[]any `json:"page_access_user_ids,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PageAccessGroupRemoveMatch is the typed request payload for PageAccessGroup.RemoveTyped.
type PageAccessGroupRemoveMatch struct {
	ComponentId *string `json:"component_id,omitempty"`
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// PageAccessUser is the typed data model for the page_access_user entity.
type PageAccessUser struct {
	ComponentIds []any `json:"component_ids"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalLogin *string `json:"external_login,omitempty"`
	Id *string `json:"id,omitempty"`
	MetricIds []any `json:"metric_ids"`
	PageAccessGroupId *string `json:"page_access_group_id,omitempty"`
	PageAccessGroupIds *string `json:"page_access_group_ids,omitempty"`
	PageAccessUser *map[string]any `json:"page_access_user,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PageAccessUserLoadMatch is the typed request payload for PageAccessUser.LoadTyped.
type PageAccessUserLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// PageAccessUserListMatch is the typed request payload for PageAccessUser.ListTyped.
type PageAccessUserListMatch struct {
	Id string `json:"id"`
}

// PageAccessUserCreateData is the typed request payload for PageAccessUser.CreateTyped.
type PageAccessUserCreateData struct {
	Id string `json:"id"`
	ComponentIds []any `json:"component_ids"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalLogin *string `json:"external_login,omitempty"`
	MetricIds []any `json:"metric_ids"`
	PageAccessGroupId *string `json:"page_access_group_id,omitempty"`
	PageAccessGroupIds *string `json:"page_access_group_ids,omitempty"`
	PageAccessUser *map[string]any `json:"page_access_user,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PageAccessUserUpdateData is the typed request payload for PageAccessUser.UpdateTyped.
type PageAccessUserUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	ComponentIds *[]any `json:"component_ids,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalLogin *string `json:"external_login,omitempty"`
	MetricIds *[]any `json:"metric_ids,omitempty"`
	PageAccessGroupId *string `json:"page_access_group_id,omitempty"`
	PageAccessGroupIds *string `json:"page_access_group_ids,omitempty"`
	PageAccessUser *map[string]any `json:"page_access_user,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PageAccessUserRemoveMatch is the typed request payload for PageAccessUser.RemoveTyped.
type PageAccessUserRemoveMatch struct {
	ComponentId *string `json:"component_id,omitempty"`
	Id string `json:"id"`
	PageId string `json:"page_id"`
	MetricId *string `json:"metric_id,omitempty"`
}

// Permission is the typed data model for the permission entity.
type Permission struct {
	Pages *map[string]any `json:"pages,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// PermissionLoadMatch is the typed request payload for Permission.LoadTyped.
type PermissionLoadMatch struct {
	Id string `json:"id"`
	OrganizationId string `json:"organization_id"`
}

// PermissionUpdateData is the typed request payload for Permission.UpdateTyped.
type PermissionUpdateData struct {
	Id string `json:"id"`
	OrganizationId string `json:"organization_id"`
	Pages *map[string]any `json:"pages,omitempty"`
	UserId *string `json:"user_id,omitempty"`
}

// Postmortem is the typed data model for the postmortem entity.
type Postmortem struct {
	Body *string `json:"body,omitempty"`
	BodyDraft *string `json:"body_draft,omitempty"`
	BodyDraftUpdatedAt *string `json:"body_draft_updated_at,omitempty"`
	BodyUpdatedAt *string `json:"body_updated_at,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomTweet *string `json:"custom_tweet,omitempty"`
	NotifySubscribers *bool `json:"notify_subscribers,omitempty"`
	NotifyTwitter *bool `json:"notify_twitter,omitempty"`
	Postmortem map[string]any `json:"postmortem"`
	PreviewKey *string `json:"preview_key,omitempty"`
	PublishedAt *string `json:"published_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// PostmortemLoadMatch is the typed request payload for Postmortem.LoadTyped.
type PostmortemLoadMatch struct {
	IncidentId string `json:"incident_id"`
	PageId string `json:"page_id"`
}

// PostmortemUpdateData is the typed request payload for Postmortem.UpdateTyped.
type PostmortemUpdateData struct {
	IncidentId string `json:"incident_id"`
	PageId string `json:"page_id"`
	Body *string `json:"body,omitempty"`
	BodyDraft *string `json:"body_draft,omitempty"`
	BodyDraftUpdatedAt *string `json:"body_draft_updated_at,omitempty"`
	BodyUpdatedAt *string `json:"body_updated_at,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomTweet *string `json:"custom_tweet,omitempty"`
	NotifySubscribers *bool `json:"notify_subscribers,omitempty"`
	NotifyTwitter *bool `json:"notify_twitter,omitempty"`
	Postmortem *map[string]any `json:"postmortem,omitempty"`
	PreviewKey *string `json:"preview_key,omitempty"`
	PublishedAt *string `json:"published_at,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// StatusEmbedConfig is the typed data model for the status_embed_config entity.
type StatusEmbedConfig struct {
	IncidentBackgroundColor *string `json:"incident_background_color,omitempty"`
	IncidentTextColor *string `json:"incident_text_color,omitempty"`
	MaintenanceBackgroundColor *string `json:"maintenance_background_color,omitempty"`
	MaintenanceTextColor *string `json:"maintenance_text_color,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	Position *string `json:"position,omitempty"`
	StatusEmbedConfig *map[string]any `json:"status_embed_config,omitempty"`
}

// StatusEmbedConfigLoadMatch is the typed request payload for StatusEmbedConfig.LoadTyped.
type StatusEmbedConfigLoadMatch struct {
	PageId string `json:"page_id"`
}

// StatusEmbedConfigUpdateData is the typed request payload for StatusEmbedConfig.UpdateTyped.
type StatusEmbedConfigUpdateData struct {
	PageId string `json:"page_id"`
	IncidentBackgroundColor *string `json:"incident_background_color,omitempty"`
	IncidentTextColor *string `json:"incident_text_color,omitempty"`
	MaintenanceBackgroundColor *string `json:"maintenance_background_color,omitempty"`
	MaintenanceTextColor *string `json:"maintenance_text_color,omitempty"`
	Position *string `json:"position,omitempty"`
	StatusEmbedConfig *map[string]any `json:"status_embed_config,omitempty"`
}

// Subscriber is the typed data model for the subscriber entity.
type Subscriber struct {
	ComponentIds *[]any `json:"component_ids,omitempty"`
	Components *string `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DisplayPhoneNumber *string `json:"display_phone_number,omitempty"`
	Email *string `json:"email,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Id *string `json:"id,omitempty"`
	IntegrationPartner *int `json:"integration_partner,omitempty"`
	Mode *string `json:"mode,omitempty"`
	ObfuscatedChannelName *string `json:"obfuscated_channel_name,omitempty"`
	PageAccessUserId *string `json:"page_access_user_id,omitempty"`
	PhoneCountry *string `json:"phone_country,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PurgeAt *string `json:"purge_at,omitempty"`
	QuarantinedAt *string `json:"quarantined_at,omitempty"`
	SkipConfirmationNotification *bool `json:"skip_confirmation_notification,omitempty"`
	SkipUnsubscriptionNotification *bool `json:"skip_unsubscription_notification,omitempty"`
	Slack *int `json:"slack,omitempty"`
	Sms *int `json:"sms,omitempty"`
	State *string `json:"state,omitempty"`
	Subscriber *map[string]any `json:"subscriber,omitempty"`
	Subscribers string `json:"subscribers"`
	Teams *int `json:"teams,omitempty"`
	Type *string `json:"type,omitempty"`
	Webhook *int `json:"webhook,omitempty"`
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

// SubscriberLoadMatch is the typed request payload for Subscriber.LoadTyped.
type SubscriberLoadMatch struct {
	Id string `json:"id"`
	IncidentId *string `json:"incident_id,omitempty"`
	PageId string `json:"page_id"`
}

// SubscriberListMatch is the typed request payload for Subscriber.ListTyped.
type SubscriberListMatch struct {
	PageId string `json:"page_id"`
	IncidentId *string `json:"incident_id,omitempty"`
}

// SubscriberCreateData is the typed request payload for Subscriber.CreateTyped.
type SubscriberCreateData struct {
	IncidentId *string `json:"incident_id,omitempty"`
	PageId string `json:"page_id"`
	ComponentIds *[]any `json:"component_ids,omitempty"`
	Components *string `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DisplayPhoneNumber *string `json:"display_phone_number,omitempty"`
	Email *string `json:"email,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Id *string `json:"id,omitempty"`
	IntegrationPartner *int `json:"integration_partner,omitempty"`
	Mode *string `json:"mode,omitempty"`
	ObfuscatedChannelName *string `json:"obfuscated_channel_name,omitempty"`
	PageAccessUserId *string `json:"page_access_user_id,omitempty"`
	PhoneCountry *string `json:"phone_country,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PurgeAt *string `json:"purge_at,omitempty"`
	QuarantinedAt *string `json:"quarantined_at,omitempty"`
	SkipConfirmationNotification *bool `json:"skip_confirmation_notification,omitempty"`
	SkipUnsubscriptionNotification *bool `json:"skip_unsubscription_notification,omitempty"`
	Slack *int `json:"slack,omitempty"`
	Sms *int `json:"sms,omitempty"`
	State *string `json:"state,omitempty"`
	Subscriber *map[string]any `json:"subscriber,omitempty"`
	Subscribers string `json:"subscribers"`
	Teams *int `json:"teams,omitempty"`
	Type *string `json:"type,omitempty"`
	Webhook *int `json:"webhook,omitempty"`
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

// SubscriberUpdateData is the typed request payload for Subscriber.UpdateTyped.
type SubscriberUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
	ComponentIds *[]any `json:"component_ids,omitempty"`
	Components *string `json:"components,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	DisplayPhoneNumber *string `json:"display_phone_number,omitempty"`
	Email *string `json:"email,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	IntegrationPartner *int `json:"integration_partner,omitempty"`
	Mode *string `json:"mode,omitempty"`
	ObfuscatedChannelName *string `json:"obfuscated_channel_name,omitempty"`
	PageAccessUserId *string `json:"page_access_user_id,omitempty"`
	PhoneCountry *string `json:"phone_country,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	PurgeAt *string `json:"purge_at,omitempty"`
	QuarantinedAt *string `json:"quarantined_at,omitempty"`
	SkipConfirmationNotification *bool `json:"skip_confirmation_notification,omitempty"`
	SkipUnsubscriptionNotification *bool `json:"skip_unsubscription_notification,omitempty"`
	Slack *int `json:"slack,omitempty"`
	Sms *int `json:"sms,omitempty"`
	State *string `json:"state,omitempty"`
	Subscriber *map[string]any `json:"subscriber,omitempty"`
	Subscribers *string `json:"subscribers,omitempty"`
	Teams *int `json:"teams,omitempty"`
	Type *string `json:"type,omitempty"`
	Webhook *int `json:"webhook,omitempty"`
	WorkspaceName *string `json:"workspace_name,omitempty"`
}

// SubscriberRemoveMatch is the typed request payload for Subscriber.RemoveTyped.
type SubscriberRemoveMatch struct {
	Id string `json:"id"`
	IncidentId *string `json:"incident_id,omitempty"`
	PageId string `json:"page_id"`
}

// User is the typed data model for the user entity.
type User struct {
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	OrganizationId *string `json:"organization_id,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	User map[string]any `json:"user"`
}

// UserListMatch is the typed request payload for User.ListTyped.
type UserListMatch struct {
	OrganizationId string `json:"organization_id"`
}

// UserCreateData is the typed request payload for User.CreateTyped.
type UserCreateData struct {
	OrganizationId string `json:"organization_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	Id *string `json:"id,omitempty"`
	LastName *string `json:"last_name,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	User map[string]any `json:"user"`
}

// UserRemoveMatch is the typed request payload for User.RemoveTyped.
type UserRemoveMatch struct {
	Id string `json:"id"`
	OrganizationId string `json:"organization_id"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
