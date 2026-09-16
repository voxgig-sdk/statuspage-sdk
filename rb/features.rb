# Statuspage SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/debug_feature'
require_relative 'feature/idempotency_feature'
require_relative 'feature/metrics_feature'
require_relative 'feature/paging_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module StatuspageFeatures
  def self.make_feature(name)
    case name
    when "base"
      StatuspageBaseFeature.new
    when "debug"
      StatuspageDebugFeature.new
    when "idempotency"
      StatuspageIdempotencyFeature.new
    when "metrics"
      StatuspageMetricsFeature.new
    when "paging"
      StatuspagePagingFeature.new
    when "ratelimit"
      StatuspageRatelimitFeature.new
    when "retry"
      StatuspageRetryFeature.new
    when "test"
      StatuspageTestFeature.new
    when "timeout"
      StatuspageTimeoutFeature.new
    else
      StatuspageBaseFeature.new
    end
  end
end
