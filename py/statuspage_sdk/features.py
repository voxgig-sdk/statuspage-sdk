# Statuspage SDK feature factory

from statuspage_sdk.feature.base_feature import StatuspageBaseFeature
from statuspage_sdk.feature.debug_feature import StatuspageDebugFeature
from statuspage_sdk.feature.idempotency_feature import StatuspageIdempotencyFeature
from statuspage_sdk.feature.metrics_feature import StatuspageMetricsFeature
from statuspage_sdk.feature.paging_feature import StatuspagePagingFeature
from statuspage_sdk.feature.ratelimit_feature import StatuspageRatelimitFeature
from statuspage_sdk.feature.retry_feature import StatuspageRetryFeature
from statuspage_sdk.feature.test_feature import StatuspageTestFeature
from statuspage_sdk.feature.timeout_feature import StatuspageTimeoutFeature


_FEATURES = {
    "base": lambda: StatuspageBaseFeature(),
    "debug": lambda: StatuspageDebugFeature(),
    "idempotency": lambda: StatuspageIdempotencyFeature(),
    "metrics": lambda: StatuspageMetricsFeature(),
    "paging": lambda: StatuspagePagingFeature(),
    "ratelimit": lambda: StatuspageRatelimitFeature(),
    "retry": lambda: StatuspageRetryFeature(),
    "test": lambda: StatuspageTestFeature(),
    "timeout": lambda: StatuspageTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
