package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewComponentEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewComponentGroupUptimeEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewGroupComponentEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewIncidentEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewIncidentPostmortemEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewIncidentSubscriberEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewIncidentTemplateEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewIncidentUpdateEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewMetricEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewMetricsProviderEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewPageEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewPageAccessGroupEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewPageAccessUserEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewPermissionEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewPostmortemEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewStatusEmbedConfigEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewSubscriberEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

var NewUserEntityFunc func(client *StatuspageSDK, entopts map[string]any) StatuspageEntity

