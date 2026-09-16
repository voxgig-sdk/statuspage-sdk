package voxgigstatuspagesdk

import (
	"github.com/voxgig-sdk/statuspage-sdk/go/core"
	"github.com/voxgig-sdk/statuspage-sdk/go/entity"
	"github.com/voxgig-sdk/statuspage-sdk/go/feature"
	_ "github.com/voxgig-sdk/statuspage-sdk/go/utility"
)

// Type aliases preserve external API.
type StatuspageSDK = core.StatuspageSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type StatuspageEntity = core.StatuspageEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type StatuspageError = core.StatuspageError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewDebugFeatureFunc = func() core.Feature {
		return feature.NewDebugFeature()
	}
	core.NewIdempotencyFeatureFunc = func() core.Feature {
		return feature.NewIdempotencyFeature()
	}
	core.NewMetricsFeatureFunc = func() core.Feature {
		return feature.NewMetricsFeature()
	}
	core.NewPagingFeatureFunc = func() core.Feature {
		return feature.NewPagingFeature()
	}
	core.NewRatelimitFeatureFunc = func() core.Feature {
		return feature.NewRatelimitFeature()
	}
	core.NewRetryFeatureFunc = func() core.Feature {
		return feature.NewRetryFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewTimeoutFeatureFunc = func() core.Feature {
		return feature.NewTimeoutFeature()
	}
	core.NewComponentEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewComponentEntity(client, entopts)
	}
	core.NewComponentGroupUptimeEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewComponentGroupUptimeEntity(client, entopts)
	}
	core.NewGroupComponentEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewGroupComponentEntity(client, entopts)
	}
	core.NewIncidentEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewIncidentEntity(client, entopts)
	}
	core.NewIncidentPostmortemEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewIncidentPostmortemEntity(client, entopts)
	}
	core.NewIncidentSubscriberEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewIncidentSubscriberEntity(client, entopts)
	}
	core.NewIncidentTemplateEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewIncidentTemplateEntity(client, entopts)
	}
	core.NewIncidentUpdateEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewIncidentUpdateEntity(client, entopts)
	}
	core.NewMetricEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewMetricEntity(client, entopts)
	}
	core.NewMetricsProviderEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewMetricsProviderEntity(client, entopts)
	}
	core.NewPageEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewPageEntity(client, entopts)
	}
	core.NewPageAccessGroupEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewPageAccessGroupEntity(client, entopts)
	}
	core.NewPageAccessUserEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewPageAccessUserEntity(client, entopts)
	}
	core.NewPermissionEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewPermissionEntity(client, entopts)
	}
	core.NewPostmortemEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewPostmortemEntity(client, entopts)
	}
	core.NewStatusEmbedConfigEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewStatusEmbedConfigEntity(client, entopts)
	}
	core.NewSubscriberEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewSubscriberEntity(client, entopts)
	}
	core.NewUserEntityFunc = func(client *core.StatuspageSDK, entopts map[string]any) core.StatuspageEntity {
		return entity.NewUserEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewStatuspageSDK = core.NewStatuspageSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewStatuspageSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *StatuspageSDK  { return NewStatuspageSDK(nil) }
func Test() *StatuspageSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewDebugFeature = feature.NewDebugFeature
var NewIdempotencyFeature = feature.NewIdempotencyFeature
var NewMetricsFeature = feature.NewMetricsFeature
var NewPagingFeature = feature.NewPagingFeature
var NewRatelimitFeature = feature.NewRatelimitFeature
var NewRetryFeature = feature.NewRetryFeature
var NewTestFeature = feature.NewTestFeature
var NewTimeoutFeature = feature.NewTimeoutFeature
