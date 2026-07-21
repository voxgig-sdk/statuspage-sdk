# Statuspage SDK

require_relative 'utility/struct/voxgig_struct'
require_relative 'core/utility_type'
require_relative 'core/spec'
require_relative 'core/helpers'

# Load utility registration
require_relative 'utility/register'

# Load config and features
require_relative 'config'
require_relative 'feature/base_feature'
require_relative 'features'

# Load typed models (Struct value objects).
require_relative 'Statuspage_types'


class StatuspageSDK
  attr_accessor :mode, :features, :options

  def initialize(options = {})
    @mode = "live"
    @features = []
    @options = nil

    utility = StatuspageUtility.new
    @_utility = utility

    config = StatuspageConfig.make_config

    @_rootctx = utility.make_context.call({
      "client" => self,
      "utility" => utility,
      "config" => config,
      "options" => options || {},
      "shared" => {},
    }, nil)

    @options = utility.make_options.call(@_rootctx)

    if VoxgigStruct.getpath(@options, "feature.test.active") == true
      @mode = "test"
    end

    @_rootctx.options = @options

    # Add features in the resolved order (make_options puts an explicit array
    # order first, else defaults to test-first). Ordering matters: the `test`
    # feature installs the base mock transport and the transport features
    # (retry/cache/netsim/proxy/ratelimit) wrap whatever is current, so `test`
    # must be added before them to sit at the base of the chain.
    feature_opts = StatuspageHelpers.to_map(VoxgigStruct.getprop(@options, "feature"))
    if feature_opts
      featureorder = VoxgigStruct.getpath(@options, "__derived__.featureorder")
      if featureorder.is_a?(Array)
        featureorder.each do |fname|
          fopts = StatuspageHelpers.to_map(feature_opts[fname])
          if fopts && fopts["active"] == true
            utility.feature_add.call(@_rootctx, StatuspageFeatures.make_feature(fname))
          end
        end
      end
    end

    # Add extension features.
    extend_val = VoxgigStruct.getprop(@options, "extend")
    if extend_val.is_a?(Array)
      extend_val.each do |f|
        if f.respond_to?(:get_name)
          utility.feature_add.call(@_rootctx, f)
        end
      end
    end

    # Initialize features.
    @features.each do |f|
      utility.feature_init.call(@_rootctx, f)
    end

    utility.feature_hook.call(@_rootctx, "PostConstruct")
  end

  def options_map
    out = VoxgigStruct.clone(@options)
    out.is_a?(Hash) ? out : {}
  end

  def get_utility
    StatuspageUtility.copy(@_utility)
  end

  def get_root_ctx
    @_rootctx
  end

  def prepare(fetchargs = {})
    utility = @_utility
    fetchargs ||= {}

    ctrl = StatuspageHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "prepare",
      "ctrl" => ctrl,
    }, @_rootctx)

    opts = @options
    path = VoxgigStruct.getprop(fetchargs, "path") || ""
    path = "" unless path.is_a?(String)
    method_val = VoxgigStruct.getprop(fetchargs, "method") || "GET"
    method_val = "GET" unless method_val.is_a?(String)
    params = StatuspageHelpers.to_map(VoxgigStruct.getprop(fetchargs, "params")) || {}
    query = StatuspageHelpers.to_map(VoxgigStruct.getprop(fetchargs, "query")) || {}
    headers = utility.prepare_headers.call(ctx)

    base = VoxgigStruct.getprop(opts, "base") || ""
    base = "" unless base.is_a?(String)
    prefix = VoxgigStruct.getprop(opts, "prefix") || ""
    prefix = "" unless prefix.is_a?(String)
    suffix = VoxgigStruct.getprop(opts, "suffix") || ""
    suffix = "" unless suffix.is_a?(String)

    ctx.spec = StatuspageSpec.new({
      "base" => base, "prefix" => prefix, "suffix" => suffix,
      "path" => path, "method" => method_val,
      "params" => params, "query" => query, "headers" => headers,
      "body" => VoxgigStruct.getprop(fetchargs, "body"),
      "step" => "start",
    })

    # Merge user-provided headers.
    uh = VoxgigStruct.getprop(fetchargs, "headers")
    if uh.is_a?(Hash)
      uh.each { |k, v| ctx.spec.headers[k] = v }
    end

    _, err = utility.prepare_auth.call(ctx)
    raise err if err

    # make_fetch_def returns a (fetchdef, err) tuple; destructure it and
    # return just the fetchdef Hash (raising on error) so callers — including
    # direct(), which indexes fetchdef["url"] — receive a Hash, mirroring the
    # ts/py prepare().
    fetchdef, fd_err = utility.make_fetch_def.call(ctx)
    raise fd_err if fd_err

    fetchdef
  end

  def direct(fetchargs = {})
    utility = @_utility

    # direct() is the raw-HTTP escape hatch: it always returns a result hash
    # ({ "ok" => ..., ... }) and never raises. prepare() raises on error, so
    # trap that and surface it in the hash.
    begin
      fetchdef = prepare(fetchargs)
    rescue StatuspageError => err
      return { "ok" => false, "err" => err }
    end

    fetchargs ||= {}
    ctrl = StatuspageHelpers.to_map(VoxgigStruct.getprop(fetchargs, "ctrl")) || {}

    ctx = utility.make_context.call({
      "opname" => "direct",
      "ctrl" => ctrl,
    }, @_rootctx)

    url = fetchdef["url"] || ""
    fetched, fetch_err = utility.fetcher.call(ctx, url, fetchdef)

    return { "ok" => false, "err" => fetch_err } if fetch_err

    if fetched.nil?
      return {
        "ok" => false,
        "err" => ctx.make_error("direct_no_response", "response: undefined"),
      }
    end

    if fetched.is_a?(Hash)
      status = StatuspageHelpers.to_int(VoxgigStruct.getprop(fetched, "status"))
      headers = VoxgigStruct.getprop(fetched, "headers") || {}

      # No-body responses (204, 304) and explicit zero content-length must
      # skip JSON parsing — calling json() on an empty body errors.
      content_length = headers.is_a?(Hash) ? headers["content-length"] : nil
      no_body = status == 204 || status == 304 || content_length.to_s == "0"

      json_data = nil
      unless no_body
        jf = VoxgigStruct.getprop(fetched, "json")
        if jf.is_a?(Proc)
          begin
            json_data = jf.call
          rescue StandardError
            # Non-JSON body — leave data nil, keep status/headers.
            json_data = nil
          end
        end
      end

      return {
        "ok" => status >= 200 && status < 300,
        "status" => status,
        "headers" => headers,
        "data" => json_data,
      }
    end

    return {
      "ok" => false,
      "err" => ctx.make_error("direct_invalid", "invalid response type"),
    }
  end


  # Canonical facade: client.Component.list / client.Component.load({ "id" => ... })
  def Component(data = nil)
    require_relative 'entity/component_entity'
    ComponentEntity.new(self, data)
  end


  # Canonical facade: client.ComponentGroupUptime.list / client.ComponentGroupUptime.load({ "id" => ... })
  def ComponentGroupUptime(data = nil)
    require_relative 'entity/component_group_uptime_entity'
    ComponentGroupUptimeEntity.new(self, data)
  end


  # Canonical facade: client.GroupComponent.list / client.GroupComponent.load({ "id" => ... })
  def GroupComponent(data = nil)
    require_relative 'entity/group_component_entity'
    GroupComponentEntity.new(self, data)
  end


  # Canonical facade: client.Incident.list / client.Incident.load({ "id" => ... })
  def Incident(data = nil)
    require_relative 'entity/incident_entity'
    IncidentEntity.new(self, data)
  end


  # Canonical facade: client.IncidentPostmortem.list / client.IncidentPostmortem.load({ "id" => ... })
  def IncidentPostmortem(data = nil)
    require_relative 'entity/incident_postmortem_entity'
    IncidentPostmortemEntity.new(self, data)
  end


  # Canonical facade: client.IncidentSubscriber.list / client.IncidentSubscriber.load({ "id" => ... })
  def IncidentSubscriber(data = nil)
    require_relative 'entity/incident_subscriber_entity'
    IncidentSubscriberEntity.new(self, data)
  end


  # Canonical facade: client.IncidentTemplate.list / client.IncidentTemplate.load({ "id" => ... })
  def IncidentTemplate(data = nil)
    require_relative 'entity/incident_template_entity'
    IncidentTemplateEntity.new(self, data)
  end


  # Canonical facade: client.IncidentUpdate.list / client.IncidentUpdate.load({ "id" => ... })
  def IncidentUpdate(data = nil)
    require_relative 'entity/incident_update_entity'
    IncidentUpdateEntity.new(self, data)
  end


  # Canonical facade: client.Metric.list / client.Metric.load({ "id" => ... })
  def Metric(data = nil)
    require_relative 'entity/metric_entity'
    MetricEntity.new(self, data)
  end


  # Canonical facade: client.MetricsProvider.list / client.MetricsProvider.load({ "id" => ... })
  def MetricsProvider(data = nil)
    require_relative 'entity/metrics_provider_entity'
    MetricsProviderEntity.new(self, data)
  end


  # Canonical facade: client.Page.list / client.Page.load({ "id" => ... })
  def Page(data = nil)
    require_relative 'entity/page_entity'
    PageEntity.new(self, data)
  end


  # Canonical facade: client.PageAccessGroup.list / client.PageAccessGroup.load({ "id" => ... })
  def PageAccessGroup(data = nil)
    require_relative 'entity/page_access_group_entity'
    PageAccessGroupEntity.new(self, data)
  end


  # Canonical facade: client.PageAccessUser.list / client.PageAccessUser.load({ "id" => ... })
  def PageAccessUser(data = nil)
    require_relative 'entity/page_access_user_entity'
    PageAccessUserEntity.new(self, data)
  end


  # Canonical facade: client.Permission.list / client.Permission.load({ "id" => ... })
  def Permission(data = nil)
    require_relative 'entity/permission_entity'
    PermissionEntity.new(self, data)
  end


  # Canonical facade: client.Postmortem.list / client.Postmortem.load({ "id" => ... })
  def Postmortem(data = nil)
    require_relative 'entity/postmortem_entity'
    PostmortemEntity.new(self, data)
  end


  # Canonical facade: client.StatusEmbedConfig.list / client.StatusEmbedConfig.load({ "id" => ... })
  def StatusEmbedConfig(data = nil)
    require_relative 'entity/status_embed_config_entity'
    StatusEmbedConfigEntity.new(self, data)
  end


  # Canonical facade: client.Subscriber.list / client.Subscriber.load({ "id" => ... })
  def Subscriber(data = nil)
    require_relative 'entity/subscriber_entity'
    SubscriberEntity.new(self, data)
  end


  # Canonical facade: client.User.list / client.User.load({ "id" => ... })
  def User(data = nil)
    require_relative 'entity/user_entity'
    UserEntity.new(self, data)
  end



  def self.test(testopts = nil, sdkopts = nil)
    sdkopts = sdkopts || {}
    sdkopts = VoxgigStruct.clone(sdkopts)
    sdkopts = {} unless sdkopts.is_a?(Hash)

    testopts = testopts || {}
    testopts = VoxgigStruct.clone(testopts)
    testopts = {} unless testopts.is_a?(Hash)
    testopts["active"] = true

    VoxgigStruct.setpath(sdkopts, "feature.test", testopts)

    sdk = StatuspageSDK.new(sdkopts)
    sdk.mode = "test"
    sdk
  end
end
