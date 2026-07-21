-- Statuspage SDK

local vs = require("utility.struct.struct")
local Utility = require("core.utility_type")
local Spec = require("core.spec")
local helpers = require("core.helpers")

-- Load utility registration (populates Utility._registrar)
require("utility.register")

-- Load features
local BaseFeature = require("feature.base_feature")
local features_factory = require("features")


local StatuspageSDK = {}
StatuspageSDK.__index = StatuspageSDK


local function _make_feature(name)
  local factory = features_factory[name]
  if factory ~= nil then
    return factory()
  end
  return features_factory.base()
end

StatuspageSDK._make_feature = _make_feature


function StatuspageSDK.new(options)
  local self = setmetatable({}, StatuspageSDK)
  self.mode = "live"
  self.features = {}
  self.options = nil

  local utility = Utility.new()
  self._utility = utility

  local config = require("config")()

  self._rootctx = utility.make_context({
    client = self,
    utility = utility,
    config = config,
    options = options or {},
    shared = {},
  }, nil)

  self.options = utility.make_options(self._rootctx)

  if vs.getpath(self.options, "feature.test.active") == true then
    self.mode = "test"
  end

  self._rootctx.options = self.options

  -- Add features in the resolved order (make_options puts an explicit list
  -- order first, else defaults to test-first). Ordering matters: the `test`
  -- feature installs the base mock transport and the transport features
  -- (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
  -- must be added before them to sit at the base of the chain.
  local feature_opts = helpers.to_map(vs.getprop(self.options, "feature"))
  if feature_opts ~= nil then
    local featureorder = vs.getpath(self.options, "__derived__.featureorder")
    if type(featureorder) == "table" then
      for _, fname in ipairs(featureorder) do
        local fopts = helpers.to_map(feature_opts[fname])
        if fopts ~= nil and fopts["active"] == true then
          utility.feature_add(self._rootctx, _make_feature(fname))
        end
      end
    end
  end

  -- Add extension features.
  local extend = vs.getprop(self.options, "extend")
  if type(extend) == "table" then
    for _, f in ipairs(extend) do
      if type(f) == "table" and type(f.get_name) == "function" then
        utility.feature_add(self._rootctx, f)
      end
    end
  end

  -- Initialize features.
  for _, f in ipairs(self.features) do
    utility.feature_init(self._rootctx, f)
  end

  utility.feature_hook(self._rootctx, "PostConstruct")

  -- #BuildFeatures

  return self
end


function StatuspageSDK:options_map()
  local out = vs.clone(self.options)
  if type(out) == "table" then
    return out
  end
  return {}
end


function StatuspageSDK:get_utility()
  return Utility.copy(self._utility)
end


function StatuspageSDK:get_root_ctx()
  return self._rootctx
end


function StatuspageSDK:prepare(fetchargs)
  local utility = self._utility

  fetchargs = fetchargs or {}

  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "prepare",
    ctrl = ctrl,
  }, self._rootctx)

  local options = self.options

  local path = vs.getprop(fetchargs, "path") or ""
  if type(path) ~= "string" then path = "" end

  local method = vs.getprop(fetchargs, "method") or "GET"
  if type(method) ~= "string" then method = "GET" end

  local params = helpers.to_map(vs.getprop(fetchargs, "params")) or {}
  local query = helpers.to_map(vs.getprop(fetchargs, "query")) or {}

  local headers = utility.prepare_headers(ctx)

  local base = vs.getprop(options, "base") or ""
  if type(base) ~= "string" then base = "" end
  local prefix = vs.getprop(options, "prefix") or ""
  if type(prefix) ~= "string" then prefix = "" end
  local suffix = vs.getprop(options, "suffix") or ""
  if type(suffix) ~= "string" then suffix = "" end

  ctx.spec = Spec.new({
    base = base,
    prefix = prefix,
    suffix = suffix,
    path = path,
    method = method,
    params = params,
    query = query,
    headers = headers,
    body = vs.getprop(fetchargs, "body"),
    step = "start",
  })

  -- Merge user-provided headers.
  local uh = vs.getprop(fetchargs, "headers")
  if type(uh) == "table" then
    for k, v in pairs(uh) do
      ctx.spec.headers[k] = v
    end
  end

  local _, err = utility.prepare_auth(ctx)
  if err ~= nil then
    return nil, err
  end

  return utility.make_fetch_def(ctx)
end


function StatuspageSDK:direct(fetchargs)
  local utility = self._utility

  local fetchdef, err = self:prepare(fetchargs)
  if err ~= nil then
    return { ok = false, err = err }, nil
  end

  fetchargs = fetchargs or {}
  local ctrl = helpers.to_map(vs.getprop(fetchargs, "ctrl")) or {}

  local ctx = utility.make_context({
    opname = "direct",
    ctrl = ctrl,
  }, self._rootctx)

  local url = fetchdef["url"] or ""
  local fetched, fetch_err = utility.fetcher(ctx, url, fetchdef)

  if fetch_err ~= nil then
    return { ok = false, err = fetch_err }, nil
  end

  if fetched == nil then
    return {
      ok = false,
      err = ctx:make_error("direct_no_response", "response: undefined"),
    }, nil
  end

  if type(fetched) == "table" then
    local status = helpers.to_int(vs.getprop(fetched, "status"))
    local headers = vs.getprop(fetched, "headers") or {}

    -- No-body responses (204, 304) and explicit zero content-length
    -- must skip JSON parsing — calling json() on an empty body errors.
    local content_length = nil
    if type(headers) == "table" then
      content_length = headers["content-length"]
    end
    local no_body = status == 204 or status == 304 or tostring(content_length) == "0"

    local json_data = nil
    if not no_body then
      local jf = vs.getprop(fetched, "json")
      if type(jf) == "function" then
        local ok, result = pcall(jf)
        if ok then
          json_data = result
        end
        -- Non-JSON body: json_data stays nil, status/headers preserved.
      end
    end

    return {
      ok = status >= 200 and status < 300,
      status = status,
      headers = headers,
      data = json_data,
    }, nil
  end

  return {
    ok = false,
    err = ctx:make_error("direct_invalid", "invalid response type"),
  }, nil
end



-- Idiomatic facade: client:Component():list() / client:Component():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Component(data)
  local EntityMod = require("entity.component_entity")
  if data == nil then
    if self._component == nil then
      self._component = EntityMod.new(self, nil)
    end
    return self._component
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:ComponentGroupUptime():list() / client:ComponentGroupUptime():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:ComponentGroupUptime(data)
  local EntityMod = require("entity.component_group_uptime_entity")
  if data == nil then
    if self._component_group_uptime == nil then
      self._component_group_uptime = EntityMod.new(self, nil)
    end
    return self._component_group_uptime
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:GroupComponent():list() / client:GroupComponent():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:GroupComponent(data)
  local EntityMod = require("entity.group_component_entity")
  if data == nil then
    if self._group_component == nil then
      self._group_component = EntityMod.new(self, nil)
    end
    return self._group_component
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Incident():list() / client:Incident():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Incident(data)
  local EntityMod = require("entity.incident_entity")
  if data == nil then
    if self._incident == nil then
      self._incident = EntityMod.new(self, nil)
    end
    return self._incident
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IncidentPostmortem():list() / client:IncidentPostmortem():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:IncidentPostmortem(data)
  local EntityMod = require("entity.incident_postmortem_entity")
  if data == nil then
    if self._incident_postmortem == nil then
      self._incident_postmortem = EntityMod.new(self, nil)
    end
    return self._incident_postmortem
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IncidentSubscriber():list() / client:IncidentSubscriber():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:IncidentSubscriber(data)
  local EntityMod = require("entity.incident_subscriber_entity")
  if data == nil then
    if self._incident_subscriber == nil then
      self._incident_subscriber = EntityMod.new(self, nil)
    end
    return self._incident_subscriber
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IncidentTemplate():list() / client:IncidentTemplate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:IncidentTemplate(data)
  local EntityMod = require("entity.incident_template_entity")
  if data == nil then
    if self._incident_template == nil then
      self._incident_template = EntityMod.new(self, nil)
    end
    return self._incident_template
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:IncidentUpdate():list() / client:IncidentUpdate():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:IncidentUpdate(data)
  local EntityMod = require("entity.incident_update_entity")
  if data == nil then
    if self._incident_update == nil then
      self._incident_update = EntityMod.new(self, nil)
    end
    return self._incident_update
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Metric():list() / client:Metric():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Metric(data)
  local EntityMod = require("entity.metric_entity")
  if data == nil then
    if self._metric == nil then
      self._metric = EntityMod.new(self, nil)
    end
    return self._metric
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:MetricsProvider():list() / client:MetricsProvider():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:MetricsProvider(data)
  local EntityMod = require("entity.metrics_provider_entity")
  if data == nil then
    if self._metrics_provider == nil then
      self._metrics_provider = EntityMod.new(self, nil)
    end
    return self._metrics_provider
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Page():list() / client:Page():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Page(data)
  local EntityMod = require("entity.page_entity")
  if data == nil then
    if self._page == nil then
      self._page = EntityMod.new(self, nil)
    end
    return self._page
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PageAccessGroup():list() / client:PageAccessGroup():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:PageAccessGroup(data)
  local EntityMod = require("entity.page_access_group_entity")
  if data == nil then
    if self._page_access_group == nil then
      self._page_access_group = EntityMod.new(self, nil)
    end
    return self._page_access_group
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:PageAccessUser():list() / client:PageAccessUser():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:PageAccessUser(data)
  local EntityMod = require("entity.page_access_user_entity")
  if data == nil then
    if self._page_access_user == nil then
      self._page_access_user = EntityMod.new(self, nil)
    end
    return self._page_access_user
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Permission():list() / client:Permission():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Permission(data)
  local EntityMod = require("entity.permission_entity")
  if data == nil then
    if self._permission == nil then
      self._permission = EntityMod.new(self, nil)
    end
    return self._permission
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Postmortem():list() / client:Postmortem():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Postmortem(data)
  local EntityMod = require("entity.postmortem_entity")
  if data == nil then
    if self._postmortem == nil then
      self._postmortem = EntityMod.new(self, nil)
    end
    return self._postmortem
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:StatusEmbedConfig():list() / client:StatusEmbedConfig():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:StatusEmbedConfig(data)
  local EntityMod = require("entity.status_embed_config_entity")
  if data == nil then
    if self._status_embed_config == nil then
      self._status_embed_config = EntityMod.new(self, nil)
    end
    return self._status_embed_config
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:Subscriber():list() / client:Subscriber():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:Subscriber(data)
  local EntityMod = require("entity.subscriber_entity")
  if data == nil then
    if self._subscriber == nil then
      self._subscriber = EntityMod.new(self, nil)
    end
    return self._subscriber
  end
  return EntityMod.new(self, data)
end


-- Idiomatic facade: client:User():list() / client:User():load({ id = ... })
-- Entity access is capitalised (PascalCase) for parity with the other SDKs.
function StatuspageSDK:User(data)
  local EntityMod = require("entity.user_entity")
  if data == nil then
    if self._user == nil then
      self._user = EntityMod.new(self, nil)
    end
    return self._user
  end
  return EntityMod.new(self, data)
end




function StatuspageSDK.test(testopts, sdkopts)
  sdkopts = sdkopts or {}
  sdkopts = vs.clone(sdkopts)
  if type(sdkopts) ~= "table" then
    sdkopts = {}
  end

  testopts = testopts or {}
  testopts = vs.clone(testopts)
  if type(testopts) ~= "table" then
    testopts = {}
  end
  testopts["active"] = true

  vs.setpath(sdkopts, "feature.test", testopts)

  local sdk = StatuspageSDK.new(sdkopts)
  sdk.mode = "test"

  return sdk
end


return StatuspageSDK
