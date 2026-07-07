# ProjectName SDK exists test

import pytest
from statuspage_sdk import StatuspageSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = StatuspageSDK.test(None, None)
        assert testsdk is not None
