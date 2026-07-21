-- Statuspage SDK error

local StatuspageError = {}
StatuspageError.__index = StatuspageError


function StatuspageError.new(code, msg, ctx)
  local self = setmetatable({}, StatuspageError)
  self.is_sdk_error = true
  self.sdk = "Statuspage"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function StatuspageError:error()
  return self.msg
end


function StatuspageError:__tostring()
  return self.msg
end


return StatuspageError
