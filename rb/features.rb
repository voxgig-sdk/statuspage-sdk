# Statuspage SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module StatuspageFeatures
  def self.make_feature(name)
    case name
    when "base"
      StatuspageBaseFeature.new
    when "test"
      StatuspageTestFeature.new
    else
      StatuspageBaseFeature.new
    end
  end
end
