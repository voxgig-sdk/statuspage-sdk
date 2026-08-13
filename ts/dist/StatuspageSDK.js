"use strict";
// Statuspage Ts SDK
Object.defineProperty(exports, "__esModule", { value: true });
exports.SDK = exports.StatuspageSDK = exports.StatuspageEntityBase = exports.BaseFeature = exports.config = exports.stdutil = void 0;
const ComponentEntity_1 = require("./entity/ComponentEntity");
const ComponentGroupUptimeEntity_1 = require("./entity/ComponentGroupUptimeEntity");
const GroupComponentEntity_1 = require("./entity/GroupComponentEntity");
const IncidentEntity_1 = require("./entity/IncidentEntity");
const IncidentPostmortemEntity_1 = require("./entity/IncidentPostmortemEntity");
const IncidentSubscriberEntity_1 = require("./entity/IncidentSubscriberEntity");
const IncidentTemplateEntity_1 = require("./entity/IncidentTemplateEntity");
const IncidentUpdateEntity_1 = require("./entity/IncidentUpdateEntity");
const MetricEntity_1 = require("./entity/MetricEntity");
const MetricsProviderEntity_1 = require("./entity/MetricsProviderEntity");
const PageEntity_1 = require("./entity/PageEntity");
const PageAccessGroupEntity_1 = require("./entity/PageAccessGroupEntity");
const PageAccessUserEntity_1 = require("./entity/PageAccessUserEntity");
const PermissionEntity_1 = require("./entity/PermissionEntity");
const PostmortemEntity_1 = require("./entity/PostmortemEntity");
const StatusEmbedConfigEntity_1 = require("./entity/StatusEmbedConfigEntity");
const SubscriberEntity_1 = require("./entity/SubscriberEntity");
const UserEntity_1 = require("./entity/UserEntity");
const node_util_1 = require("node:util");
const Config_1 = require("./Config");
Object.defineProperty(exports, "config", { enumerable: true, get: function () { return Config_1.config; } });
const StatuspageEntityBase_1 = require("./StatuspageEntityBase");
Object.defineProperty(exports, "StatuspageEntityBase", { enumerable: true, get: function () { return StatuspageEntityBase_1.StatuspageEntityBase; } });
const Utility_1 = require("./utility/Utility");
const BaseFeature_1 = require("./feature/base/BaseFeature");
Object.defineProperty(exports, "BaseFeature", { enumerable: true, get: function () { return BaseFeature_1.BaseFeature; } });
const stdutil = new Utility_1.Utility();
exports.stdutil = stdutil;
class StatuspageSDK {
    _mode = 'live';
    _options;
    _utility = new Utility_1.Utility();
    _features;
    _rootctx;
    constructor(options) {
        this._rootctx = this._utility.makeContext({
            client: this,
            utility: this._utility,
            config: Config_1.config,
            options,
            shared: new WeakMap()
        });
        this._options = this._utility.makeOptions(this._rootctx);
        const struct = this._utility.struct;
        const getpath = struct.getpath;
        if (true === getpath(this._options.feature, 'test.active')) {
            this._mode = 'test';
        }
        this._rootctx.options = this._options;
        this._features = [];
        const featureAdd = this._utility.featureAdd;
        const featureInit = this._utility.featureInit;
        // Add features in the resolved order (makeOptions puts an explicit
        // array order first, else defaults to test-first). Ordering matters:
        // the `test` feature installs the base mock transport and the transport
        // features (retry/cache/netsim/proxy/ratelimit) wrap whatever is current,
        // so `test` must be added before them to sit at the base of the chain.
        const featureorder = getpath(this._options, '__derived__.featureorder') || [];
        for (const fname of featureorder) {
            const fopts = this._options.feature[fname] || {};
            if (fopts.active) {
                featureAdd(this._rootctx, this._rootctx.config.makeFeature(fname));
            }
        }
        if (null != this._options.extend) {
            for (let f of this._options.extend) {
                featureAdd(this._rootctx, f);
            }
        }
        for (let f of this._features) {
            featureInit(this._rootctx, f);
        }
        const featureHook = this._utility.featureHook;
        featureHook(this._rootctx, 'PostConstruct');
    }
    options() {
        return this._utility.struct.clone(this._options);
    }
    utility() {
        return this._utility.struct.clone(this._utility);
    }
    async prepare(fetchargs) {
        const utility = this._utility;
        const struct = utility.struct;
        const clone = struct.clone;
        const { makeContext, makeFetchDef, prepareHeaders, prepareAuth, } = utility;
        fetchargs = fetchargs || {};
        let ctx = makeContext({
            opname: 'prepare',
            ctrl: fetchargs.ctrl || {},
        }, this._rootctx);
        const options = this._options;
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
        };
        ctx.spec = spec;
        // Merge user-provided headers over SDK defaults.
        if (fetchargs.headers) {
            const uheaders = fetchargs.headers;
            for (let key in uheaders) {
                spec.headers[key] = uheaders[key];
            }
        }
        // Apply SDK auth (apikey, auth prefix, etc.)
        const authResult = prepareAuth(ctx);
        if (authResult instanceof Error) {
            return authResult;
        }
        return makeFetchDef(ctx);
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
            };
        }
        return this._rawRequest(fetchargs);
    }
    // Ungated request path shared by direct() and graphql(), each of which
    // checks its own allow.op token first. Private, rather than a flag on
    // fetchargs: a caller-supplied marker would let anyone opt straight back
    // out of the gate by passing it.
    async _rawRequest(fetchargs) {
        const utility = this._utility;
        const fetcher = utility.fetcher;
        const makeContext = utility.makeContext;
        const fetchdef = await this.prepare(fetchargs);
        if (fetchdef instanceof Error) {
            return fetchdef;
        }
        let ctx = makeContext({
            opname: 'direct',
            ctrl: (fetchargs || {}).ctrl || {},
        }, this._rootctx);
        try {
            const fetched = await fetcher(ctx, fetchdef.url, fetchdef);
            if (null == fetched) {
                return { ok: false, err: ctx.error('direct_no_response', 'response: undefined') };
            }
            else if (fetched instanceof Error) {
                return { ok: false, err: fetched };
            }
            const status = fetched.status;
            // No body responses (204 No Content, 304 Not Modified) and explicit
            // zero content-length must skip JSON parsing — fetched.json() would
            // throw `Unexpected end of JSON input` on an empty body.
            const headers = fetched.headers;
            const contentLength = headers && 'function' === typeof headers.get
                ? headers.get('content-length')
                : (headers || {})['content-length'];
            const noBody = 204 === status || 304 === status || '0' === String(contentLength);
            let json = undefined;
            if (!noBody) {
                try {
                    json = 'function' === typeof fetched.json ? await fetched.json() : fetched.json;
                }
                catch (parseErr) {
                    // Body wasn't valid JSON — surface the raw response rather than
                    // throwing. data stays undefined; callers can inspect status/headers.
                    json = undefined;
                }
            }
            return {
                ok: status >= 200 && status < 300,
                status,
                headers: fetched.headers,
                data: json,
            };
        }
        catch (err) {
            return { ok: false, err };
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
        const options = this._options;
        if (!options.allow.op.includes('graphql')) {
            return {
                ok: false,
                err: new Error('StatuspageSDK: graphql: operation not allowed by' +
                    ' SDK option allow.op value: "' + options.allow.op + '"'),
            };
        }
        const res = await this._rawRequest({
            method: 'POST',
            headers: { 'content-type': 'application/json' },
            body: { query, variables: variables || {} },
            ctrl,
        });
        if (res instanceof Error) {
            return res;
        }
        // Errors are read BEFORE any status check: a GraphQL parse or validation
        // failure comes back as HTTP 400 carrying the standard { errors: [...] }
        // body, and the raw path represents a non-2xx as { ok: false } with no
        // err — so returning early on status would discard the server's own
        // diagnostics, which are the only useful part of that response.
        const errors = null == res.data ? undefined : res.data.errors;
        if (null != errors && Array.isArray(errors) && 0 < errors.length) {
            const first = errors[0] || {};
            const err = new Error('StatuspageSDK: graphql: ' +
                (first.message || 'graphql error'));
            err.graphql = errors;
            return { ok: false, status: res.status, headers: res.headers, err, data: res.data };
        }
        return res;
    }
    // Entity access: `client.Component().list()` / `client.Component().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Component(entopts) {
        const self = this;
        return new ComponentEntity_1.ComponentEntity(self, entopts);
    }
    // Entity access: `client.ComponentGroupUptime().list()` / `client.ComponentGroupUptime().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    ComponentGroupUptime(entopts) {
        const self = this;
        return new ComponentGroupUptimeEntity_1.ComponentGroupUptimeEntity(self, entopts);
    }
    // Entity access: `client.GroupComponent().list()` / `client.GroupComponent().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    GroupComponent(entopts) {
        const self = this;
        return new GroupComponentEntity_1.GroupComponentEntity(self, entopts);
    }
    // Entity access: `client.Incident().list()` / `client.Incident().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Incident(entopts) {
        const self = this;
        return new IncidentEntity_1.IncidentEntity(self, entopts);
    }
    // Entity access: `client.IncidentPostmortem().list()` / `client.IncidentPostmortem().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IncidentPostmortem(entopts) {
        const self = this;
        return new IncidentPostmortemEntity_1.IncidentPostmortemEntity(self, entopts);
    }
    // Entity access: `client.IncidentSubscriber().list()` / `client.IncidentSubscriber().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IncidentSubscriber(entopts) {
        const self = this;
        return new IncidentSubscriberEntity_1.IncidentSubscriberEntity(self, entopts);
    }
    // Entity access: `client.IncidentTemplate().list()` / `client.IncidentTemplate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IncidentTemplate(entopts) {
        const self = this;
        return new IncidentTemplateEntity_1.IncidentTemplateEntity(self, entopts);
    }
    // Entity access: `client.IncidentUpdate().list()` / `client.IncidentUpdate().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    IncidentUpdate(entopts) {
        const self = this;
        return new IncidentUpdateEntity_1.IncidentUpdateEntity(self, entopts);
    }
    // Entity access: `client.Metric().list()` / `client.Metric().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Metric(entopts) {
        const self = this;
        return new MetricEntity_1.MetricEntity(self, entopts);
    }
    // Entity access: `client.MetricsProvider().list()` / `client.MetricsProvider().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    MetricsProvider(entopts) {
        const self = this;
        return new MetricsProviderEntity_1.MetricsProviderEntity(self, entopts);
    }
    // Entity access: `client.Page().list()` / `client.Page().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Page(entopts) {
        const self = this;
        return new PageEntity_1.PageEntity(self, entopts);
    }
    // Entity access: `client.PageAccessGroup().list()` / `client.PageAccessGroup().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PageAccessGroup(entopts) {
        const self = this;
        return new PageAccessGroupEntity_1.PageAccessGroupEntity(self, entopts);
    }
    // Entity access: `client.PageAccessUser().list()` / `client.PageAccessUser().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    PageAccessUser(entopts) {
        const self = this;
        return new PageAccessUserEntity_1.PageAccessUserEntity(self, entopts);
    }
    // Entity access: `client.Permission().list()` / `client.Permission().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Permission(entopts) {
        const self = this;
        return new PermissionEntity_1.PermissionEntity(self, entopts);
    }
    // Entity access: `client.Postmortem().list()` / `client.Postmortem().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Postmortem(entopts) {
        const self = this;
        return new PostmortemEntity_1.PostmortemEntity(self, entopts);
    }
    // Entity access: `client.StatusEmbedConfig().list()` / `client.StatusEmbedConfig().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    StatusEmbedConfig(entopts) {
        const self = this;
        return new StatusEmbedConfigEntity_1.StatusEmbedConfigEntity(self, entopts);
    }
    // Entity access: `client.Subscriber().list()` / `client.Subscriber().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    Subscriber(entopts) {
        const self = this;
        return new SubscriberEntity_1.SubscriberEntity(self, entopts);
    }
    // Entity access: `client.User().list()` / `client.User().load({ id })`.
    // The argument is the entity OPTIONS object (passed to the entity
    // constructor as entopts), not initial entity data.
    User(entopts) {
        const self = this;
        return new UserEntity_1.UserEntity(self, entopts);
    }
    static test(testoptsarg, sdkoptsarg) {
        const struct = stdutil.struct;
        const setpath = struct.setpath;
        const getdef = struct.getdef;
        const clone = struct.clone;
        const setprop = struct.setprop;
        const sdkopts = getdef(clone(sdkoptsarg), {});
        const testopts = getdef(clone(testoptsarg), {});
        setprop(testopts, 'active', true);
        setpath(sdkopts, 'feature.test', testopts);
        const testsdk = new StatuspageSDK(sdkopts);
        testsdk._mode = 'test';
        return testsdk;
    }
    tester(testopts, sdkopts) {
        return StatuspageSDK.test(testopts, sdkopts);
    }
    toJSON() {
        return { name: 'Statuspage' };
    }
    toString() {
        return 'Statuspage ' + this._utility.struct.jsonify(this.toJSON());
    }
    [node_util_1.inspect.custom]() {
        return this.toString();
    }
}
exports.StatuspageSDK = StatuspageSDK;
const SDK = StatuspageSDK;
exports.SDK = SDK;
//# sourceMappingURL=StatuspageSDK.js.map