// Typed models for the Statuspage SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Component is the typed data model for the component entity.
type Component struct {
	AutomationEmail *string `json:"automation_email,omitempty"`
	Component *map[string]any `json:"component,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Description *string `json:"description,omitempty"`
	Group *bool `json:"group,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Id *string `json:"id,omitempty"`
	MajorOutage *int `json:"major_outage,omitempty"`
	Name *string `json:"name,omitempty"`
	OnlyShowIfDegraded *bool `json:"only_show_if_degraded,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	PartialOutage *int `json:"partial_outage,omitempty"`
	Position *int `json:"position,omitempty"`
	RangeEnd *string `json:"range_end,omitempty"`
	RangeStart *string `json:"range_start,omitempty"`
	RelatedEvent *map[string]any `json:"related_event,omitempty"`
	Showcase *bool `json:"showcase,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
	Status *string `json:"status,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
	UptimePercentage *float64 `json:"uptime_percentage,omitempty"`
	Warning *string `json:"warning,omitempty"`
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
}

// ComponentUpdateData is the typed request payload for Component.UpdateTyped.
type ComponentUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// ComponentRemoveMatch is the typed request payload for Component.RemoveTyped.
type ComponentRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// ComponentGroupUptime is the typed data model for the component_group_uptime entity.
type ComponentGroupUptime struct {
	Id *string `json:"id,omitempty"`
	MajorOutage *int `json:"major_outage,omitempty"`
	Name *string `json:"name,omitempty"`
	PartialOutage *int `json:"partial_outage,omitempty"`
	RangeEnd *string `json:"range_end,omitempty"`
	RangeStart *string `json:"range_start,omitempty"`
	RelatedEvent *map[string]any `json:"related_event,omitempty"`
	UptimePercentage *float64 `json:"uptime_percentage,omitempty"`
	Warning *string `json:"warning,omitempty"`
}

// ComponentGroupUptimeLoadMatch is the typed request payload for ComponentGroupUptime.LoadTyped.
type ComponentGroupUptimeLoadMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// GroupComponent is the typed data model for the group_component entity.
type GroupComponent struct {
	Component *string `json:"component,omitempty"`
	ComponentGroup map[string]any `json:"component_group"`
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
}

// GroupComponentUpdateData is the typed request payload for GroupComponent.UpdateTyped.
type GroupComponentUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
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
	Component *[]any `json:"component,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Id *string `json:"id,omitempty"`
	Impact *string `json:"impact,omitempty"`
	ImpactOverride *string `json:"impact_override,omitempty"`
	Incident map[string]any `json:"incident"`
	IncidentImpact *[]any `json:"incident_impact,omitempty"`
	IncidentUpdate *[]any `json:"incident_update,omitempty"`
	Metadata *map[string]any `json:"metadata,omitempty"`
	MonitoringAt *string `json:"monitoring_at,omitempty"`
	Name *string `json:"name,omitempty"`
	PageId *string `json:"page_id,omitempty"`
	PostmortemBody *string `json:"postmortem_body,omitempty"`
	PostmortemBodyLastUpdatedAt *string `json:"postmortem_body_last_updated_at,omitempty"`
	PostmortemIgnored *bool `json:"postmortem_ignored,omitempty"`
	PostmortemNotifiedSubscriber *bool `json:"postmortem_notified_subscriber,omitempty"`
	PostmortemNotifiedTwitter *bool `json:"postmortem_notified_twitter,omitempty"`
	PostmortemPublishedAt *bool `json:"postmortem_published_at,omitempty"`
	ReminderInterval *string `json:"reminder_interval,omitempty"`
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
}

// IncidentUpdateData is the typed request payload for Incident.UpdateTyped.
type IncidentUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
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
	Component *[]any `json:"component,omitempty"`
	GroupId *string `json:"group_id,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	ShouldSendNotification *bool `json:"should_send_notification,omitempty"`
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
}

// IncidentUpdate is the typed data model for the incident_update entity.
type IncidentUpdate struct {
	AffectedComponent *[]any `json:"affected_component,omitempty"`
	Body *string `json:"body,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomTweet *string `json:"custom_tweet,omitempty"`
	DeliverNotification *bool `json:"deliver_notification,omitempty"`
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
}

// Metric is the typed data model for the metric entity.
type Metric struct {
	BackfillPercentage *int `json:"backfill_percentage,omitempty"`
	Backfilled *bool `json:"backfilled,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	Data map[string]any `json:"data"`
	DecimalPlace *int `json:"decimal_place,omitempty"`
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
}

// MetricUpdateData is the typed request payload for Metric.UpdateTyped.
type MetricUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
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
}

// MetricsProviderUpdateData is the typed request payload for MetricsProvider.UpdateTyped.
type MetricsProviderUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// MetricsProviderRemoveMatch is the typed request payload for MetricsProvider.RemoveTyped.
type MetricsProviderRemoveMatch struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// Page is the typed data model for the page entity.
type Page struct {
	ActivityScore *float64 `json:"activity_score,omitempty"`
	AllowEmailSubscriber *bool `json:"allow_email_subscriber,omitempty"`
	AllowIncidentSubscriber *bool `json:"allow_incident_subscriber,omitempty"`
	AllowPageSubscriber *bool `json:"allow_page_subscriber,omitempty"`
	AllowRssAtomFeed *bool `json:"allow_rss_atom_feed,omitempty"`
	AllowSmsSubscriber *bool `json:"allow_sms_subscriber,omitempty"`
	AllowWebhookSubscriber *bool `json:"allow_webhook_subscriber,omitempty"`
	Branding *string `json:"branding,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CssBlue *string `json:"css_blue,omitempty"`
	CssBodyBackgroundColor *string `json:"css_body_background_color,omitempty"`
	CssBorderColor *string `json:"css_border_color,omitempty"`
	CssFontColor *string `json:"css_font_color,omitempty"`
	CssGraphColor *string `json:"css_graph_color,omitempty"`
	CssGreen *string `json:"css_green,omitempty"`
	CssLightFontColor *string `json:"css_light_font_color,omitempty"`
	CssLinkColor *string `json:"css_link_color,omitempty"`
	CssNoData *string `json:"css_no_data,omitempty"`
	CssOrange *string `json:"css_orange,omitempty"`
	CssRed *string `json:"css_red,omitempty"`
	CssYellow *string `json:"css_yellow,omitempty"`
	Domain *string `json:"domain,omitempty"`
	EmailLogo *string `json:"email_logo,omitempty"`
	FaviconLogo *string `json:"favicon_logo,omitempty"`
	Headline *string `json:"headline,omitempty"`
	HeroCover *string `json:"hero_cover,omitempty"`
	HiddenFromSearch *bool `json:"hidden_from_search,omitempty"`
	Id *string `json:"id,omitempty"`
	IpRestriction *string `json:"ip_restriction,omitempty"`
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
	ViewersMustBeTeamMember *bool `json:"viewers_must_be_team_member,omitempty"`
}

// PageLoadMatch is the typed request payload for Page.LoadTyped.
type PageLoadMatch struct {
	Id string `json:"id"`
}

// PageListMatch is the typed request payload for Page.ListTyped.
type PageListMatch struct {
	ActivityScore *float64 `json:"activity_score,omitempty"`
	AllowEmailSubscriber *bool `json:"allow_email_subscriber,omitempty"`
	AllowIncidentSubscriber *bool `json:"allow_incident_subscriber,omitempty"`
	AllowPageSubscriber *bool `json:"allow_page_subscriber,omitempty"`
	AllowRssAtomFeed *bool `json:"allow_rss_atom_feed,omitempty"`
	AllowSmsSubscriber *bool `json:"allow_sms_subscriber,omitempty"`
	AllowWebhookSubscriber *bool `json:"allow_webhook_subscriber,omitempty"`
	Branding *string `json:"branding,omitempty"`
	City *string `json:"city,omitempty"`
	Country *string `json:"country,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CssBlue *string `json:"css_blue,omitempty"`
	CssBodyBackgroundColor *string `json:"css_body_background_color,omitempty"`
	CssBorderColor *string `json:"css_border_color,omitempty"`
	CssFontColor *string `json:"css_font_color,omitempty"`
	CssGraphColor *string `json:"css_graph_color,omitempty"`
	CssGreen *string `json:"css_green,omitempty"`
	CssLightFontColor *string `json:"css_light_font_color,omitempty"`
	CssLinkColor *string `json:"css_link_color,omitempty"`
	CssNoData *string `json:"css_no_data,omitempty"`
	CssOrange *string `json:"css_orange,omitempty"`
	CssRed *string `json:"css_red,omitempty"`
	CssYellow *string `json:"css_yellow,omitempty"`
	Domain *string `json:"domain,omitempty"`
	EmailLogo *string `json:"email_logo,omitempty"`
	FaviconLogo *string `json:"favicon_logo,omitempty"`
	Headline *string `json:"headline,omitempty"`
	HeroCover *string `json:"hero_cover,omitempty"`
	HiddenFromSearch *bool `json:"hidden_from_search,omitempty"`
	Id *string `json:"id,omitempty"`
	IpRestriction *string `json:"ip_restriction,omitempty"`
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
	ViewersMustBeTeamMember *bool `json:"viewers_must_be_team_member,omitempty"`
}

// PageUpdateData is the typed request payload for Page.UpdateTyped.
type PageUpdateData struct {
	Id string `json:"id"`
}

// PageAccessGroup is the typed data model for the page_access_group entity.
type PageAccessGroup struct {
	ComponentId *[]any `json:"component_id,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	ExternalIdentifier *string `json:"external_identifier,omitempty"`
	Id *string `json:"id,omitempty"`
	MetricId *[]any `json:"metric_id,omitempty"`
	Name *string `json:"name,omitempty"`
	PageAccessGroup *map[string]any `json:"page_access_group,omitempty"`
	PageAccessUserId *[]any `json:"page_access_user_id,omitempty"`
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
}

// PageAccessGroupUpdateData is the typed request payload for PageAccessGroup.UpdateTyped.
type PageAccessGroupUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// PageAccessGroupRemoveMatch is the typed request payload for PageAccessGroup.RemoveTyped.
type PageAccessGroupRemoveMatch struct {
	ComponentId *string `json:"component_id,omitempty"`
	Id string `json:"id"`
	PageId string `json:"page_id"`
}

// PageAccessUser is the typed data model for the page_access_user entity.
type PageAccessUser struct {
	ComponentId []any `json:"component_id"`
	CreatedAt *string `json:"created_at,omitempty"`
	Email *string `json:"email,omitempty"`
	ExternalLogin *string `json:"external_login,omitempty"`
	Id *string `json:"id,omitempty"`
	MetricId []any `json:"metric_id"`
	PageAccessGroupId *string `json:"page_access_group_id,omitempty"`
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
}

// PageAccessUserUpdateData is the typed request payload for PageAccessUser.UpdateTyped.
type PageAccessUserUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
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
	Data *map[string]any `json:"data,omitempty"`
	Page *map[string]any `json:"page,omitempty"`
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
}

// Postmortem is the typed data model for the postmortem entity.
type Postmortem struct {
	Body *string `json:"body,omitempty"`
	BodyDraft *string `json:"body_draft,omitempty"`
	BodyDraftUpdatedAt *string `json:"body_draft_updated_at,omitempty"`
	BodyUpdatedAt *string `json:"body_updated_at,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	CustomTweet *string `json:"custom_tweet,omitempty"`
	NotifySubscriber *bool `json:"notify_subscriber,omitempty"`
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
}

// Subscriber is the typed data model for the subscriber entity.
type Subscriber struct {
	Component *string `json:"component,omitempty"`
	ComponentId *[]any `json:"component_id,omitempty"`
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
	Team *int `json:"team,omitempty"`
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
}

// SubscriberUpdateData is the typed request payload for Subscriber.UpdateTyped.
type SubscriberUpdateData struct {
	Id string `json:"id"`
	PageId string `json:"page_id"`
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

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
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

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
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
