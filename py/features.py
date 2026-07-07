# Statuspage SDK feature factory

from feature.base_feature import StatuspageBaseFeature
from feature.test_feature import StatuspageTestFeature


def _make_feature(name):
    features = {
        "base": lambda: StatuspageBaseFeature(),
        "test": lambda: StatuspageTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
