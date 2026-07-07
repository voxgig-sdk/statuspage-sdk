# Statuspage SDK exists test

require "minitest/autorun"
require_relative "../Statuspage_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = StatuspageSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
