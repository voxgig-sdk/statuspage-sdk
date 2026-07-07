# frozen_string_literal: true

# Typed models for the Statuspage SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Component entity data model.
#
# @!attribute [rw] automation_email
#   @return [String, nil]
#
# @!attribute [rw] component
#   @return [Hash, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] group
#   @return [Boolean, nil]
#
# @!attribute [rw] group_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] major_outage
#   @return [Integer, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] only_show_if_degraded
#   @return [Boolean, nil]
#
# @!attribute [rw] page_id
#   @return [String, nil]
#
# @!attribute [rw] partial_outage
#   @return [Integer, nil]
#
# @!attribute [rw] position
#   @return [Integer, nil]
#
# @!attribute [rw] range_end
#   @return [String, nil]
#
# @!attribute [rw] range_start
#   @return [String, nil]
#
# @!attribute [rw] related_event
#   @return [Hash, nil]
#
# @!attribute [rw] showcase
#   @return [Boolean, nil]
#
# @!attribute [rw] start_date
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] uptime_percentage
#   @return [Float, nil]
#
# @!attribute [rw] warning
#   @return [String, nil]
Component = Struct.new(
  :automation_email,
  :component,
  :created_at,
  :description,
  :group,
  :group_id,
  :id,
  :major_outage,
  :name,
  :only_show_if_degraded,
  :page_id,
  :partial_outage,
  :position,
  :range_end,
  :range_start,
  :related_event,
  :showcase,
  :start_date,
  :status,
  :updated_at,
  :uptime_percentage,
  :warning,
  keyword_init: true
)

# Request payload for Component#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
ComponentLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for Component#list.
#
# @!attribute [rw] page_access_group_id
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String]
#
# @!attribute [rw] page_access_user_id
#   @return [String, nil]
ComponentListMatch = Struct.new(
  :page_access_group_id,
  :page_id,
  :page_access_user_id,
  keyword_init: true
)

# Request payload for Component#create.
#
# @!attribute [rw] page_id
#   @return [String]
ComponentCreateData = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for Component#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
ComponentUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for Component#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
ComponentRemoveMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# ComponentGroupUptime entity data model.
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] major_outage
#   @return [Integer, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] partial_outage
#   @return [Integer, nil]
#
# @!attribute [rw] range_end
#   @return [String, nil]
#
# @!attribute [rw] range_start
#   @return [String, nil]
#
# @!attribute [rw] related_event
#   @return [Hash, nil]
#
# @!attribute [rw] uptime_percentage
#   @return [Float, nil]
#
# @!attribute [rw] warning
#   @return [String, nil]
ComponentGroupUptime = Struct.new(
  :id,
  :major_outage,
  :name,
  :partial_outage,
  :range_end,
  :range_start,
  :related_event,
  :uptime_percentage,
  :warning,
  keyword_init: true
)

# Request payload for ComponentGroupUptime#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
ComponentGroupUptimeLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# GroupComponent entity data model.
#
# @!attribute [rw] component
#   @return [String, nil]
#
# @!attribute [rw] component_group
#   @return [Hash]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] description
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String, nil]
#
# @!attribute [rw] position
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
GroupComponent = Struct.new(
  :component,
  :component_group,
  :created_at,
  :description,
  :id,
  :name,
  :page_id,
  :position,
  :updated_at,
  keyword_init: true
)

# Request payload for GroupComponent#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
GroupComponentLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for GroupComponent#list.
#
# @!attribute [rw] page_id
#   @return [String]
GroupComponentListMatch = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for GroupComponent#create.
#
# @!attribute [rw] page_id
#   @return [String]
GroupComponentCreateData = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for GroupComponent#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
GroupComponentUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for GroupComponent#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
GroupComponentRemoveMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Incident entity data model.
#
# @!attribute [rw] auto_transition_deliver_notifications_at_end
#   @return [Boolean, nil]
#
# @!attribute [rw] auto_transition_deliver_notifications_at_start
#   @return [Boolean, nil]
#
# @!attribute [rw] auto_transition_to_maintenance_state
#   @return [Boolean, nil]
#
# @!attribute [rw] auto_transition_to_operational_state
#   @return [Boolean, nil]
#
# @!attribute [rw] component
#   @return [Array, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] impact
#   @return [String, nil]
#
# @!attribute [rw] impact_override
#   @return [String, nil]
#
# @!attribute [rw] incident
#   @return [Hash]
#
# @!attribute [rw] incident_impact
#   @return [Array, nil]
#
# @!attribute [rw] incident_update
#   @return [Array, nil]
#
# @!attribute [rw] metadata
#   @return [Hash, nil]
#
# @!attribute [rw] monitoring_at
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String, nil]
#
# @!attribute [rw] postmortem_body
#   @return [String, nil]
#
# @!attribute [rw] postmortem_body_last_updated_at
#   @return [String, nil]
#
# @!attribute [rw] postmortem_ignored
#   @return [Boolean, nil]
#
# @!attribute [rw] postmortem_notified_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] postmortem_notified_twitter
#   @return [Boolean, nil]
#
# @!attribute [rw] postmortem_published_at
#   @return [Boolean, nil]
#
# @!attribute [rw] reminder_interval
#   @return [String, nil]
#
# @!attribute [rw] resolved_at
#   @return [String, nil]
#
# @!attribute [rw] scheduled_auto_completed
#   @return [Boolean, nil]
#
# @!attribute [rw] scheduled_auto_in_progress
#   @return [Boolean, nil]
#
# @!attribute [rw] scheduled_for
#   @return [String, nil]
#
# @!attribute [rw] scheduled_remind_prior
#   @return [Boolean, nil]
#
# @!attribute [rw] scheduled_reminded_at
#   @return [String, nil]
#
# @!attribute [rw] scheduled_until
#   @return [String, nil]
#
# @!attribute [rw] shortlink
#   @return [String, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
Incident = Struct.new(
  :auto_transition_deliver_notifications_at_end,
  :auto_transition_deliver_notifications_at_start,
  :auto_transition_to_maintenance_state,
  :auto_transition_to_operational_state,
  :component,
  :created_at,
  :id,
  :impact,
  :impact_override,
  :incident,
  :incident_impact,
  :incident_update,
  :metadata,
  :monitoring_at,
  :name,
  :page_id,
  :postmortem_body,
  :postmortem_body_last_updated_at,
  :postmortem_ignored,
  :postmortem_notified_subscriber,
  :postmortem_notified_twitter,
  :postmortem_published_at,
  :reminder_interval,
  :resolved_at,
  :scheduled_auto_completed,
  :scheduled_auto_in_progress,
  :scheduled_for,
  :scheduled_remind_prior,
  :scheduled_reminded_at,
  :scheduled_until,
  :shortlink,
  :status,
  :updated_at,
  keyword_init: true
)

# Request payload for Incident#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
IncidentLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for Incident#list.
#
# @!attribute [rw] page_id
#   @return [String]
IncidentListMatch = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for Incident#create.
#
# @!attribute [rw] page_id
#   @return [String]
IncidentCreateData = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for Incident#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
IncidentUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for Incident#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
IncidentRemoveMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# IncidentPostmortem entity data model.
class IncidentPostmortem
end

# Request payload for IncidentPostmortem#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
IncidentPostmortemRemoveMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# IncidentSubscriber entity data model.
class IncidentSubscriber
end

# Request payload for IncidentSubscriber#create.
#
# @!attribute [rw] incident_id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
#
# @!attribute [rw] subscriber_id
#   @return [String]
IncidentSubscriberCreateData = Struct.new(
  :incident_id,
  :page_id,
  :subscriber_id,
  keyword_init: true
)

# IncidentTemplate entity data model.
#
# @!attribute [rw] body
#   @return [String, nil]
#
# @!attribute [rw] component
#   @return [Array, nil]
#
# @!attribute [rw] group_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] should_send_notification
#   @return [Boolean, nil]
#
# @!attribute [rw] should_tweet
#   @return [Boolean, nil]
#
# @!attribute [rw] template
#   @return [Hash]
#
# @!attribute [rw] title
#   @return [String, nil]
#
# @!attribute [rw] update_status
#   @return [String, nil]
IncidentTemplate = Struct.new(
  :body,
  :component,
  :group_id,
  :id,
  :name,
  :should_send_notification,
  :should_tweet,
  :template,
  :title,
  :update_status,
  keyword_init: true
)

# Request payload for IncidentTemplate#list.
#
# @!attribute [rw] page_id
#   @return [String]
IncidentTemplateListMatch = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for IncidentTemplate#create.
#
# @!attribute [rw] page_id
#   @return [String]
IncidentTemplateCreateData = Struct.new(
  :page_id,
  keyword_init: true
)

# IncidentUpdate entity data model.
#
# @!attribute [rw] affected_component
#   @return [Array, nil]
#
# @!attribute [rw] body
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] custom_tweet
#   @return [String, nil]
#
# @!attribute [rw] deliver_notification
#   @return [Boolean, nil]
#
# @!attribute [rw] display_at
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] incident_id
#   @return [String, nil]
#
# @!attribute [rw] incident_update
#   @return [Hash, nil]
#
# @!attribute [rw] status
#   @return [String, nil]
#
# @!attribute [rw] tweet_id
#   @return [String, nil]
#
# @!attribute [rw] twitter_updated_at
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] wants_twitter_update
#   @return [Boolean, nil]
IncidentUpdate = Struct.new(
  :affected_component,
  :body,
  :created_at,
  :custom_tweet,
  :deliver_notification,
  :display_at,
  :id,
  :incident_id,
  :incident_update,
  :status,
  :tweet_id,
  :twitter_updated_at,
  :updated_at,
  :wants_twitter_update,
  keyword_init: true
)

# Request payload for IncidentUpdate#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] incident_id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
IncidentUpdateUpdateData = Struct.new(
  :id,
  :incident_id,
  :page_id,
  keyword_init: true
)

# Metric entity data model.
#
# @!attribute [rw] backfill_percentage
#   @return [Integer, nil]
#
# @!attribute [rw] backfilled
#   @return [Boolean, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] data
#   @return [Hash]
#
# @!attribute [rw] decimal_place
#   @return [Integer, nil]
#
# @!attribute [rw] display
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] last_fetched_at
#   @return [String, nil]
#
# @!attribute [rw] metric
#   @return [Hash, nil]
#
# @!attribute [rw] metric_identifier
#   @return [String, nil]
#
# @!attribute [rw] metrics_provider_id
#   @return [String, nil]
#
# @!attribute [rw] most_recent_data_at
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] reference_name
#   @return [String, nil]
#
# @!attribute [rw] suffix
#   @return [String, nil]
#
# @!attribute [rw] tooltip_description
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] y_axis_hidden
#   @return [Boolean, nil]
#
# @!attribute [rw] y_axis_max
#   @return [Float, nil]
#
# @!attribute [rw] y_axis_min
#   @return [Float, nil]
Metric = Struct.new(
  :backfill_percentage,
  :backfilled,
  :created_at,
  :data,
  :decimal_place,
  :display,
  :id,
  :last_fetched_at,
  :metric,
  :metric_identifier,
  :metrics_provider_id,
  :most_recent_data_at,
  :name,
  :reference_name,
  :suffix,
  :tooltip_description,
  :updated_at,
  :y_axis_hidden,
  :y_axis_max,
  :y_axis_min,
  keyword_init: true
)

# Request payload for Metric#load.
#
# @!attribute [rw] metrics_provider_id
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String]
#
# @!attribute [rw] id
#   @return [String, nil]
MetricLoadMatch = Struct.new(
  :metrics_provider_id,
  :page_id,
  :id,
  keyword_init: true
)

# Request payload for Metric#list.
#
# @!attribute [rw] page_access_user_id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricListMatch = Struct.new(
  :page_access_user_id,
  :page_id,
  keyword_init: true
)

# Request payload for Metric#create.
#
# @!attribute [rw] metrics_provider_id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricCreateData = Struct.new(
  :metrics_provider_id,
  :page_id,
  keyword_init: true
)

# Request payload for Metric#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for Metric#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricRemoveMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# MetricsProvider entity data model.
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] disabled
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] last_revalidated_at
#   @return [String, nil]
#
# @!attribute [rw] metric_base_uri
#   @return [String, nil]
#
# @!attribute [rw] metrics_provider
#   @return [Hash, nil]
#
# @!attribute [rw] page_id
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
MetricsProvider = Struct.new(
  :created_at,
  :disabled,
  :id,
  :last_revalidated_at,
  :metric_base_uri,
  :metrics_provider,
  :page_id,
  :type,
  :updated_at,
  keyword_init: true
)

# Request payload for MetricsProvider#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricsProviderLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for MetricsProvider#list.
#
# @!attribute [rw] page_id
#   @return [String]
MetricsProviderListMatch = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for MetricsProvider#create.
#
# @!attribute [rw] page_id
#   @return [String]
MetricsProviderCreateData = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for MetricsProvider#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricsProviderUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for MetricsProvider#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
MetricsProviderRemoveMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Page entity data model.
#
# @!attribute [rw] activity_score
#   @return [Float, nil]
#
# @!attribute [rw] allow_email_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_incident_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_page_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_rss_atom_feed
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_sms_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_webhook_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] branding
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] css_blue
#   @return [String, nil]
#
# @!attribute [rw] css_body_background_color
#   @return [String, nil]
#
# @!attribute [rw] css_border_color
#   @return [String, nil]
#
# @!attribute [rw] css_font_color
#   @return [String, nil]
#
# @!attribute [rw] css_graph_color
#   @return [String, nil]
#
# @!attribute [rw] css_green
#   @return [String, nil]
#
# @!attribute [rw] css_light_font_color
#   @return [String, nil]
#
# @!attribute [rw] css_link_color
#   @return [String, nil]
#
# @!attribute [rw] css_no_data
#   @return [String, nil]
#
# @!attribute [rw] css_orange
#   @return [String, nil]
#
# @!attribute [rw] css_red
#   @return [String, nil]
#
# @!attribute [rw] css_yellow
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] email_logo
#   @return [String, nil]
#
# @!attribute [rw] favicon_logo
#   @return [String, nil]
#
# @!attribute [rw] headline
#   @return [String, nil]
#
# @!attribute [rw] hero_cover
#   @return [String, nil]
#
# @!attribute [rw] hidden_from_search
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ip_restriction
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notifications_email_footer
#   @return [String, nil]
#
# @!attribute [rw] notifications_from_email
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Hash, nil]
#
# @!attribute [rw] page_description
#   @return [String, nil]
#
# @!attribute [rw] state
#   @return [String, nil]
#
# @!attribute [rw] subdomain
#   @return [String, nil]
#
# @!attribute [rw] support_url
#   @return [String, nil]
#
# @!attribute [rw] time_zone
#   @return [String, nil]
#
# @!attribute [rw] transactional_logo
#   @return [String, nil]
#
# @!attribute [rw] twitter_logo
#   @return [String, nil]
#
# @!attribute [rw] twitter_username
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
#
# @!attribute [rw] viewers_must_be_team_member
#   @return [Boolean, nil]
Page = Struct.new(
  :activity_score,
  :allow_email_subscriber,
  :allow_incident_subscriber,
  :allow_page_subscriber,
  :allow_rss_atom_feed,
  :allow_sms_subscriber,
  :allow_webhook_subscriber,
  :branding,
  :city,
  :country,
  :created_at,
  :css_blue,
  :css_body_background_color,
  :css_border_color,
  :css_font_color,
  :css_graph_color,
  :css_green,
  :css_light_font_color,
  :css_link_color,
  :css_no_data,
  :css_orange,
  :css_red,
  :css_yellow,
  :domain,
  :email_logo,
  :favicon_logo,
  :headline,
  :hero_cover,
  :hidden_from_search,
  :id,
  :ip_restriction,
  :name,
  :notifications_email_footer,
  :notifications_from_email,
  :page,
  :page_description,
  :state,
  :subdomain,
  :support_url,
  :time_zone,
  :transactional_logo,
  :twitter_logo,
  :twitter_username,
  :updated_at,
  :url,
  :viewers_must_be_team_member,
  keyword_init: true
)

# Request payload for Page#load.
#
# @!attribute [rw] id
#   @return [String]
PageLoadMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for Page#list.
#
# @!attribute [rw] activity_score
#   @return [Float, nil]
#
# @!attribute [rw] allow_email_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_incident_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_page_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_rss_atom_feed
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_sms_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] allow_webhook_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] branding
#   @return [String, nil]
#
# @!attribute [rw] city
#   @return [String, nil]
#
# @!attribute [rw] country
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] css_blue
#   @return [String, nil]
#
# @!attribute [rw] css_body_background_color
#   @return [String, nil]
#
# @!attribute [rw] css_border_color
#   @return [String, nil]
#
# @!attribute [rw] css_font_color
#   @return [String, nil]
#
# @!attribute [rw] css_graph_color
#   @return [String, nil]
#
# @!attribute [rw] css_green
#   @return [String, nil]
#
# @!attribute [rw] css_light_font_color
#   @return [String, nil]
#
# @!attribute [rw] css_link_color
#   @return [String, nil]
#
# @!attribute [rw] css_no_data
#   @return [String, nil]
#
# @!attribute [rw] css_orange
#   @return [String, nil]
#
# @!attribute [rw] css_red
#   @return [String, nil]
#
# @!attribute [rw] css_yellow
#   @return [String, nil]
#
# @!attribute [rw] domain
#   @return [String, nil]
#
# @!attribute [rw] email_logo
#   @return [String, nil]
#
# @!attribute [rw] favicon_logo
#   @return [String, nil]
#
# @!attribute [rw] headline
#   @return [String, nil]
#
# @!attribute [rw] hero_cover
#   @return [String, nil]
#
# @!attribute [rw] hidden_from_search
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] ip_restriction
#   @return [String, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] notifications_email_footer
#   @return [String, nil]
#
# @!attribute [rw] notifications_from_email
#   @return [String, nil]
#
# @!attribute [rw] page
#   @return [Hash, nil]
#
# @!attribute [rw] page_description
#   @return [String, nil]
#
# @!attribute [rw] state
#   @return [String, nil]
#
# @!attribute [rw] subdomain
#   @return [String, nil]
#
# @!attribute [rw] support_url
#   @return [String, nil]
#
# @!attribute [rw] time_zone
#   @return [String, nil]
#
# @!attribute [rw] transactional_logo
#   @return [String, nil]
#
# @!attribute [rw] twitter_logo
#   @return [String, nil]
#
# @!attribute [rw] twitter_username
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] url
#   @return [String, nil]
#
# @!attribute [rw] viewers_must_be_team_member
#   @return [Boolean, nil]
PageListMatch = Struct.new(
  :activity_score,
  :allow_email_subscriber,
  :allow_incident_subscriber,
  :allow_page_subscriber,
  :allow_rss_atom_feed,
  :allow_sms_subscriber,
  :allow_webhook_subscriber,
  :branding,
  :city,
  :country,
  :created_at,
  :css_blue,
  :css_body_background_color,
  :css_border_color,
  :css_font_color,
  :css_graph_color,
  :css_green,
  :css_light_font_color,
  :css_link_color,
  :css_no_data,
  :css_orange,
  :css_red,
  :css_yellow,
  :domain,
  :email_logo,
  :favicon_logo,
  :headline,
  :hero_cover,
  :hidden_from_search,
  :id,
  :ip_restriction,
  :name,
  :notifications_email_footer,
  :notifications_from_email,
  :page,
  :page_description,
  :state,
  :subdomain,
  :support_url,
  :time_zone,
  :transactional_logo,
  :twitter_logo,
  :twitter_username,
  :updated_at,
  :url,
  :viewers_must_be_team_member,
  keyword_init: true
)

# Request payload for Page#update.
#
# @!attribute [rw] id
#   @return [String]
PageUpdateData = Struct.new(
  :id,
  keyword_init: true
)

# PageAccessGroup entity data model.
#
# @!attribute [rw] component_id
#   @return [Array, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] external_identifier
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] metric_id
#   @return [Array, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] page_access_group
#   @return [Hash, nil]
#
# @!attribute [rw] page_access_user_id
#   @return [Array, nil]
#
# @!attribute [rw] page_id
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
PageAccessGroup = Struct.new(
  :component_id,
  :created_at,
  :external_identifier,
  :id,
  :metric_id,
  :name,
  :page_access_group,
  :page_access_user_id,
  :page_id,
  :updated_at,
  keyword_init: true
)

# Request payload for PageAccessGroup#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PageAccessGroupLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for PageAccessGroup#list.
#
# @!attribute [rw] id
#   @return [String]
PageAccessGroupListMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for PageAccessGroup#create.
#
# @!attribute [rw] id
#   @return [String]
PageAccessGroupCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for PageAccessGroup#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PageAccessGroupUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for PageAccessGroup#remove.
#
# @!attribute [rw] component_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PageAccessGroupRemoveMatch = Struct.new(
  :component_id,
  :id,
  :page_id,
  keyword_init: true
)

# PageAccessUser entity data model.
#
# @!attribute [rw] component_id
#   @return [Array]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] external_login
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] metric_id
#   @return [Array]
#
# @!attribute [rw] page_access_group_id
#   @return [String, nil]
#
# @!attribute [rw] page_access_user
#   @return [Hash, nil]
#
# @!attribute [rw] page_id
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
PageAccessUser = Struct.new(
  :component_id,
  :created_at,
  :email,
  :external_login,
  :id,
  :metric_id,
  :page_access_group_id,
  :page_access_user,
  :page_id,
  :updated_at,
  keyword_init: true
)

# Request payload for PageAccessUser#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PageAccessUserLoadMatch = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for PageAccessUser#list.
#
# @!attribute [rw] id
#   @return [String]
PageAccessUserListMatch = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for PageAccessUser#create.
#
# @!attribute [rw] id
#   @return [String]
PageAccessUserCreateData = Struct.new(
  :id,
  keyword_init: true
)

# Request payload for PageAccessUser#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PageAccessUserUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for PageAccessUser#remove.
#
# @!attribute [rw] component_id
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
#
# @!attribute [rw] metric_id
#   @return [String, nil]
PageAccessUserRemoveMatch = Struct.new(
  :component_id,
  :id,
  :page_id,
  :metric_id,
  keyword_init: true
)

# Permission entity data model.
#
# @!attribute [rw] data
#   @return [Hash, nil]
#
# @!attribute [rw] page
#   @return [Hash, nil]
Permission = Struct.new(
  :data,
  :page,
  keyword_init: true
)

# Request payload for Permission#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] organization_id
#   @return [String]
PermissionLoadMatch = Struct.new(
  :id,
  :organization_id,
  keyword_init: true
)

# Request payload for Permission#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] organization_id
#   @return [String]
PermissionUpdateData = Struct.new(
  :id,
  :organization_id,
  keyword_init: true
)

# Postmortem entity data model.
#
# @!attribute [rw] body
#   @return [String, nil]
#
# @!attribute [rw] body_draft
#   @return [String, nil]
#
# @!attribute [rw] body_draft_updated_at
#   @return [String, nil]
#
# @!attribute [rw] body_updated_at
#   @return [String, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] custom_tweet
#   @return [String, nil]
#
# @!attribute [rw] notify_subscriber
#   @return [Boolean, nil]
#
# @!attribute [rw] notify_twitter
#   @return [Boolean, nil]
#
# @!attribute [rw] postmortem
#   @return [Hash]
#
# @!attribute [rw] preview_key
#   @return [String, nil]
#
# @!attribute [rw] published_at
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
Postmortem = Struct.new(
  :body,
  :body_draft,
  :body_draft_updated_at,
  :body_updated_at,
  :created_at,
  :custom_tweet,
  :notify_subscriber,
  :notify_twitter,
  :postmortem,
  :preview_key,
  :published_at,
  :updated_at,
  keyword_init: true
)

# Request payload for Postmortem#load.
#
# @!attribute [rw] incident_id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PostmortemLoadMatch = Struct.new(
  :incident_id,
  :page_id,
  keyword_init: true
)

# Request payload for Postmortem#update.
#
# @!attribute [rw] incident_id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
PostmortemUpdateData = Struct.new(
  :incident_id,
  :page_id,
  keyword_init: true
)

# StatusEmbedConfig entity data model.
#
# @!attribute [rw] incident_background_color
#   @return [String, nil]
#
# @!attribute [rw] incident_text_color
#   @return [String, nil]
#
# @!attribute [rw] maintenance_background_color
#   @return [String, nil]
#
# @!attribute [rw] maintenance_text_color
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String, nil]
#
# @!attribute [rw] position
#   @return [String, nil]
#
# @!attribute [rw] status_embed_config
#   @return [Hash, nil]
StatusEmbedConfig = Struct.new(
  :incident_background_color,
  :incident_text_color,
  :maintenance_background_color,
  :maintenance_text_color,
  :page_id,
  :position,
  :status_embed_config,
  keyword_init: true
)

# Request payload for StatusEmbedConfig#load.
#
# @!attribute [rw] page_id
#   @return [String]
StatusEmbedConfigLoadMatch = Struct.new(
  :page_id,
  keyword_init: true
)

# Request payload for StatusEmbedConfig#update.
#
# @!attribute [rw] page_id
#   @return [String]
StatusEmbedConfigUpdateData = Struct.new(
  :page_id,
  keyword_init: true
)

# Subscriber entity data model.
#
# @!attribute [rw] component
#   @return [String, nil]
#
# @!attribute [rw] component_id
#   @return [Array, nil]
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] display_phone_number
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] endpoint
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] integration_partner
#   @return [Integer, nil]
#
# @!attribute [rw] mode
#   @return [String, nil]
#
# @!attribute [rw] obfuscated_channel_name
#   @return [String, nil]
#
# @!attribute [rw] page_access_user_id
#   @return [String, nil]
#
# @!attribute [rw] phone_country
#   @return [String, nil]
#
# @!attribute [rw] phone_number
#   @return [String, nil]
#
# @!attribute [rw] purge_at
#   @return [String, nil]
#
# @!attribute [rw] quarantined_at
#   @return [String, nil]
#
# @!attribute [rw] skip_confirmation_notification
#   @return [Boolean, nil]
#
# @!attribute [rw] skip_unsubscription_notification
#   @return [Boolean, nil]
#
# @!attribute [rw] slack
#   @return [Integer, nil]
#
# @!attribute [rw] sms
#   @return [Integer, nil]
#
# @!attribute [rw] state
#   @return [String, nil]
#
# @!attribute [rw] subscriber
#   @return [Hash, nil]
#
# @!attribute [rw] team
#   @return [Integer, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
#
# @!attribute [rw] webhook
#   @return [Integer, nil]
#
# @!attribute [rw] workspace_name
#   @return [String, nil]
Subscriber = Struct.new(
  :component,
  :component_id,
  :created_at,
  :display_phone_number,
  :email,
  :endpoint,
  :id,
  :integration_partner,
  :mode,
  :obfuscated_channel_name,
  :page_access_user_id,
  :phone_country,
  :phone_number,
  :purge_at,
  :quarantined_at,
  :skip_confirmation_notification,
  :skip_unsubscription_notification,
  :slack,
  :sms,
  :state,
  :subscriber,
  :team,
  :type,
  :webhook,
  :workspace_name,
  keyword_init: true
)

# Request payload for Subscriber#load.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] incident_id
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String]
SubscriberLoadMatch = Struct.new(
  :id,
  :incident_id,
  :page_id,
  keyword_init: true
)

# Request payload for Subscriber#list.
#
# @!attribute [rw] page_id
#   @return [String]
#
# @!attribute [rw] incident_id
#   @return [String, nil]
SubscriberListMatch = Struct.new(
  :page_id,
  :incident_id,
  keyword_init: true
)

# Request payload for Subscriber#create.
#
# @!attribute [rw] incident_id
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String]
SubscriberCreateData = Struct.new(
  :incident_id,
  :page_id,
  keyword_init: true
)

# Request payload for Subscriber#update.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] page_id
#   @return [String]
SubscriberUpdateData = Struct.new(
  :id,
  :page_id,
  keyword_init: true
)

# Request payload for Subscriber#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] incident_id
#   @return [String, nil]
#
# @!attribute [rw] page_id
#   @return [String]
SubscriberRemoveMatch = Struct.new(
  :id,
  :incident_id,
  :page_id,
  keyword_init: true
)

# User entity data model.
#
# @!attribute [rw] created_at
#   @return [String, nil]
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] first_name
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] last_name
#   @return [String, nil]
#
# @!attribute [rw] organization_id
#   @return [String, nil]
#
# @!attribute [rw] updated_at
#   @return [String, nil]
#
# @!attribute [rw] user
#   @return [Hash]
User = Struct.new(
  :created_at,
  :email,
  :first_name,
  :id,
  :last_name,
  :organization_id,
  :updated_at,
  :user,
  keyword_init: true
)

# Request payload for User#list.
#
# @!attribute [rw] organization_id
#   @return [String]
UserListMatch = Struct.new(
  :organization_id,
  keyword_init: true
)

# Request payload for User#create.
#
# @!attribute [rw] organization_id
#   @return [String]
UserCreateData = Struct.new(
  :organization_id,
  keyword_init: true
)

# Request payload for User#remove.
#
# @!attribute [rw] id
#   @return [String]
#
# @!attribute [rw] organization_id
#   @return [String]
UserRemoveMatch = Struct.new(
  :id,
  :organization_id,
  keyword_init: true
)

