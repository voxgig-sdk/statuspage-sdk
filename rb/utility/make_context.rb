# Statuspage SDK utility: make_context
require_relative '../core/context'
module StatuspageUtilities
  MakeContext = ->(ctxmap, basectx) {
    StatuspageContext.new(ctxmap, basectx)
  }
end
