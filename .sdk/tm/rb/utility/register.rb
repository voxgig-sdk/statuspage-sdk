# Statuspage SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

StatuspageUtility.registrar = ->(u) {
  u.clean = StatuspageUtilities::Clean
  u.done = StatuspageUtilities::Done
  u.make_error = StatuspageUtilities::MakeError
  u.feature_add = StatuspageUtilities::FeatureAdd
  u.feature_hook = StatuspageUtilities::FeatureHook
  u.feature_init = StatuspageUtilities::FeatureInit
  u.fetcher = StatuspageUtilities::Fetcher
  u.make_fetch_def = StatuspageUtilities::MakeFetchDef
  u.make_context = StatuspageUtilities::MakeContext
  u.make_options = StatuspageUtilities::MakeOptions
  u.make_request = StatuspageUtilities::MakeRequest
  u.make_response = StatuspageUtilities::MakeResponse
  u.make_result = StatuspageUtilities::MakeResult
  u.make_point = StatuspageUtilities::MakePoint
  u.make_spec = StatuspageUtilities::MakeSpec
  u.make_url = StatuspageUtilities::MakeUrl
  u.param = StatuspageUtilities::Param
  u.prepare_auth = StatuspageUtilities::PrepareAuth
  u.prepare_body = StatuspageUtilities::PrepareBody
  u.prepare_headers = StatuspageUtilities::PrepareHeaders
  u.prepare_method = StatuspageUtilities::PrepareMethod
  u.prepare_params = StatuspageUtilities::PrepareParams
  u.prepare_path = StatuspageUtilities::PreparePath
  u.prepare_query = StatuspageUtilities::PrepareQuery
  u.result_basic = StatuspageUtilities::ResultBasic
  u.result_body = StatuspageUtilities::ResultBody
  u.result_headers = StatuspageUtilities::ResultHeaders
  u.transform_request = StatuspageUtilities::TransformRequest
  u.transform_response = StatuspageUtilities::TransformResponse
}
