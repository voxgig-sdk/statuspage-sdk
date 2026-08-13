// Statuspage Js SDK

const { ComponentEntity } = require('./entity/ComponentEntity')
const { ComponentGroupUptimeEntity } = require('./entity/ComponentGroupUptimeEntity')
const { GroupComponentEntity } = require('./entity/GroupComponentEntity')
const { IncidentEntity } = require('./entity/IncidentEntity')
const { IncidentPostmortemEntity } = require('./entity/IncidentPostmortemEntity')
const { IncidentSubscriberEntity } = require('./entity/IncidentSubscriberEntity')
const { IncidentTemplateEntity } = require('./entity/IncidentTemplateEntity')
const { IncidentUpdateEntity } = require('./entity/IncidentUpdateEntity')
const { MetricEntity } = require('./entity/MetricEntity')
const { MetricsProviderEntity } = require('./entity/MetricsProviderEntity')
const { PageEntity } = require('./entity/PageEntity')
const { PageAccessGroupEntity } = require('./entity/PageAccessGroupEntity')
const { PageAccessUserEntity } = require('./entity/PageAccessUserEntity')
const { PermissionEntity } = require('./entity/PermissionEntity')
const { PostmortemEntity } = require('./entity/PostmortemEntity')
const { StatusEmbedConfigEntity } = require('./entity/StatusEmbedConfigEntity')
const { SubscriberEntity } = require('./entity/SubscriberEntity')
const { UserEntity } = require('./entity/UserEntity')


const { inspect } = require('node:util')

const { config } = require('./Config')
const { Utility } = require('./utility/Utility')
const { StatuspageEntityBase } = require('./StatuspageEntityBase')


const { BaseFeature } = require('./feature/base/BaseFeature')


const stdutil = new Utility()


class StatuspageSDK {
  _mode = 'live'
  _options
  _utility = new Utility()
  _features
  _rootctx

  constructor(options) {

    this._rootctx = this._utility.makeContext({
      client: this,
      utility: this._utility,
      config,
      options,
      shared: new WeakMap()
    })

    this._options = this._utility.makeOptions(this._rootctx)

    const struct = this._utility.struct
    const getpath = struct.getpath

    if (true === getpath(this._options.feature, 'test.active')) {
      this._mode = 'test'
    }

    this._rootctx.options = this._options

    this._features = []

    const featureAdd = this._utility.featureAdd
    const featureInit = this._utility.featureInit

    // Add features in the resolved order (makeOptions puts an explicit
    // array order first, else defaults to test-first). Ordering matters:
    // the `test` feature installs the base mock transport and the transport
    // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
    // so `test` must be added before them to sit at the base of the chain.
    const featureorder = getpath(this._options, '__derived__.featureorder') || []
    for (const fname of featureorder) {
      const fopts = this._options.feature[fname] || {}
      if (fopts.active) {
        featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname))
      }
    }

    if (null != this._options.extend) {
      for (let f of this._options.extend) {
        featureAdd(this._rootctx, f)
      }
    }

    for (let f of this._features) {
      featureInit(this._rootctx, f)
    }

    const featureHook = this._utility.featureHook
    featureHook(this._rootctx, 'PostConstruct')
  }


  options() {
    return this._utility.struct.clone(this._options)
  }


  utility() {
    return this._utility.struct.clone(this._utility)
  }


  async prepare(fetchargs) {
    const utility = this._utility
    const struct = utility.struct
    const clone = struct.clone

    const {
      makeContext,
      makeFetchDef,
      prepareHeaders,
      prepareAuth,
    } = utility

    fetchargs = fetchargs || {}

    let ctx = makeContext({
      opname: 'prepare',
      ctrl: fetchargs.ctrl || {},
    }, this._rootctx)

    const options = this._options

    // Build spec directly from SDK options + user-provided fetch args.
    const spec = {
      base: options.base,
      prefix: options.prefix,
      suffix: options.suffix,
      path: fetchargs.path || '',
      method: fetchargs.method || 'GET',
      params: fetchargs.params || {},
      query: fetchargs.query || {},
      headers: prepareHeaders(ctx),
      body: fetchargs.body,
      step: 'start',
    }

    ctx.spec = spec

    // Merge user-provided headers over SDK defaults.
    if (fetchargs.headers) {
      const uheaders = fetchargs.headers
      for (let key in uheaders) {
        spec.headers[key] = uheaders[key]
      }
    }

    // Apply SDK auth (apikey, auth prefix, etc.)
    const authResult = prepareAuth(ctx)
    if (authResult instanceof Error) {
      return authResult
    }

    return makeFetchDef(ctx)
  }


  // Raw endpoint access is operator-controllable, like every entity op.
  // Blocking it means denying BOTH the 'direct' and 'graphql' tokens, since
  // either one reaches the same endpoint.
  async direct(fetchargs) {
    if (!this._options.allow.op.includes('direct')) {
      return {
        ok: false,
        err: new Error('StatuspageSDK: direct: operation not allowed by' +
          ' SDK option allow.op value: "' + this._options.allow.op + '"'),
      }
    }

    return this._rawRequest(fetchargs)
  }


  // Ungated request path shared by direct() and graphql(), each of which
  // checks its own allow.op token first. Private, rather than a flag on
  // fetchargs: a caller-supplied marker would let anyone opt straight back
  // out of the gate by passing it.
  async _rawRequest(fetchargs) {
    const utility = this._utility

    const fetcher = utility.fetcher
    const makeContext = utility.makeContext

    const fetchdef = await this.prepare(fetchargs)
    if (fetchdef instanceof Error) {
      return fetchdef
    }

    let ctx = makeContext({
      opname: 'direct',
      ctrl: (fetchargs || {}).ctrl || {},
    }, this._rootctx)

    try {
      const fetched = await fetcher(ctx, fetchdef.url, fetchdef)

      if (null == fetched) {
        return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') }
      }
      else if (fetched instanceof Error) {
        return { ok: false, err: fetched }
      }

      const status = fetched.status
      const json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json

      return {
        ok: status >= 200 && status < 300,
        status,
        headers: fetched.headers,
        data: json,
      }
    }
    catch (err) {
      return { ok: false, err }
    }
  }



  // Raw GraphQL access: the pressure valve that makes the generated
  // surface's deliberate omissions (per-call selection sets, typed filter
  // builders, batching, subscriptions) livable — the whole schema stays
  // reachable.
  //
  // Thin wrapper over the same prepare/fetch path `direct` uses, with the
  // one thing raw `direct` cannot do for GraphQL: a GraphQL failure rides
  // HTTP 200 as a top-level `errors` array, so status alone would report a
  // failed query as ok.
  //
  // NOTE: like `direct`, this bypasses the feature pipeline — no retry,
  // ratelimit or paging features apply.
  async graphql(query, variables, ctrl) {
    const options = this._options

    if (!options.allow.op.includes('graphql')) {
      return {
        ok: false,
        err: new Error('StatuspageSDK: graphql: operation not allowed by' +
          ' SDK option allow.op value: "' + options.allow.op + '"'),
      }
    }

    const res = await this._rawRequest({
      method: 'POST',
      headers: { 'content-type': 'application/json' },
      body: { query, variables: variables || {} },
      ctrl,
    })

    if (res instanceof Error) {
      return res
    }

    // Errors are read BEFORE any status check: a GraphQL parse or validation
    // failure comes back as HTTP 400 carrying the standard { errors: [...] }
    // body, and the raw path represents a non-2xx as { ok: false } with no
    // err — so returning early on status would discard the server's own
    // diagnostics, which are the only useful part of that response.
    const errors = null == res.data ? undefined : res.data.errors

    if (null != errors && Array.isArray(errors) && 0 < errors.length) {
      const first = errors[0] || {}
      const err = new Error('StatuspageSDK: graphql: ' +
        (first.message || 'graphql error'))
      err.graphql = errors
      return { ok: false, status: res.status, headers: res.headers, err, data: res.data }
    }

    return res
  }



  // Entity access: `client.Component().list()` / `client.Component().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Component(entopts) {
    const self = this
    return new ComponentEntity(self, entopts)
  }


  // Entity access: `client.ComponentGroupUptime().list()` / `client.ComponentGroupUptime().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  ComponentGroupUptime(entopts) {
    const self = this
    return new ComponentGroupUptimeEntity(self, entopts)
  }


  // Entity access: `client.GroupComponent().list()` / `client.GroupComponent().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  GroupComponent(entopts) {
    const self = this
    return new GroupComponentEntity(self, entopts)
  }


  // Entity access: `client.Incident().list()` / `client.Incident().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Incident(entopts) {
    const self = this
    return new IncidentEntity(self, entopts)
  }


  // Entity access: `client.IncidentPostmortem().list()` / `client.IncidentPostmortem().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IncidentPostmortem(entopts) {
    const self = this
    return new IncidentPostmortemEntity(self, entopts)
  }


  // Entity access: `client.IncidentSubscriber().list()` / `client.IncidentSubscriber().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IncidentSubscriber(entopts) {
    const self = this
    return new IncidentSubscriberEntity(self, entopts)
  }


  // Entity access: `client.IncidentTemplate().list()` / `client.IncidentTemplate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IncidentTemplate(entopts) {
    const self = this
    return new IncidentTemplateEntity(self, entopts)
  }


  // Entity access: `client.IncidentUpdate().list()` / `client.IncidentUpdate().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  IncidentUpdate(entopts) {
    const self = this
    return new IncidentUpdateEntity(self, entopts)
  }


  // Entity access: `client.Metric().list()` / `client.Metric().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Metric(entopts) {
    const self = this
    return new MetricEntity(self, entopts)
  }


  // Entity access: `client.MetricsProvider().list()` / `client.MetricsProvider().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  MetricsProvider(entopts) {
    const self = this
    return new MetricsProviderEntity(self, entopts)
  }


  // Entity access: `client.Page().list()` / `client.Page().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Page(entopts) {
    const self = this
    return new PageEntity(self, entopts)
  }


  // Entity access: `client.PageAccessGroup().list()` / `client.PageAccessGroup().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PageAccessGroup(entopts) {
    const self = this
    return new PageAccessGroupEntity(self, entopts)
  }


  // Entity access: `client.PageAccessUser().list()` / `client.PageAccessUser().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  PageAccessUser(entopts) {
    const self = this
    return new PageAccessUserEntity(self, entopts)
  }


  // Entity access: `client.Permission().list()` / `client.Permission().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Permission(entopts) {
    const self = this
    return new PermissionEntity(self, entopts)
  }


  // Entity access: `client.Postmortem().list()` / `client.Postmortem().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Postmortem(entopts) {
    const self = this
    return new PostmortemEntity(self, entopts)
  }


  // Entity access: `client.StatusEmbedConfig().list()` / `client.StatusEmbedConfig().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  StatusEmbedConfig(entopts) {
    const self = this
    return new StatusEmbedConfigEntity(self, entopts)
  }


  // Entity access: `client.Subscriber().list()` / `client.Subscriber().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  Subscriber(entopts) {
    const self = this
    return new SubscriberEntity(self, entopts)
  }


  // Entity access: `client.User().list()` / `client.User().load({ id })`.
  // The argument is the entity OPTIONS object (passed to the entity
  // constructor as entopts), not initial entity data.
  User(entopts) {
    const self = this
    return new UserEntity(self, entopts)
  }




  static test(testoptsarg, sdkoptsarg) {
    const struct = stdutil.struct
    const setpath = struct.setpath
    const getdef = struct.getdef
    const clone = struct.clone
    const setprop = struct.setprop

    const sdkopts = getdef(clone(sdkoptsarg), {})
    const testopts = getdef(clone(testoptsarg), {})
    setprop(testopts, 'active', true)
    setpath(sdkopts, 'feature.test', testopts)

    const testsdk = new StatuspageSDK(sdkopts)
    testsdk._mode = 'test'

    return testsdk
  }


  tester(testopts, sdkopts) {
    return StatuspageSDK.test(testopts, sdkopts)
  }


  toJSON() {
    return { name: 'Statuspage' }
  }

  toString() {
    return 'Statuspage ' + this._utility.struct.jsonify(this.toJSON())
  }

  [inspect.custom]() {
    return this.toString()
  }

}




const SDK = StatuspageSDK


module.exports = {
  stdutil,
  config,

  BaseFeature,
  StatuspageEntityBase,

  StatuspageSDK,
  SDK,
}

