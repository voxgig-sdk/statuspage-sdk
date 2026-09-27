"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || (function () {
    var ownKeys = function(o) {
        ownKeys = Object.getOwnPropertyNames || function (o) {
            var ar = [];
            for (var k in o) if (Object.prototype.hasOwnProperty.call(o, k)) ar[ar.length] = k;
            return ar;
        };
        return ownKeys(o);
    };
    return function (mod) {
        if (mod && mod.__esModule) return mod;
        var result = {};
        if (mod != null) for (var k = ownKeys(mod), i = 0; i < k.length; i++) if (k[i] !== "default") __createBinding(result, mod, k[i]);
        __setModuleDefault(result, mod);
        return result;
    };
})();
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const node_path_1 = __importDefault(require("node:path"));
const Fs = __importStar(require("node:fs"));
const node_test_1 = require("node:test");
const node_assert_1 = __importDefault(require("node:assert"));
const live_runner_1 = require("../../live-runner");
const live_entity_1 = require("../../live-entity");
const __1 = require("../../..");
const utility_1 = require("../../utility");
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('StatusEmbedConfigEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.StatusEmbedConfig();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['update', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'status_embed_config.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "incident_background_color": { "a": true, "h": "Incident Background Color", "n": "incident_background_color", "r": false, "sh": "Color of status embed iframe background when displaying incident", "t": "`$STRING`", "key$": "incident_background_color", "index$": 0 }, "incident_text_color": { "a": true, "h": "Incident Text Color", "n": "incident_text_color", "r": false, "sh": "Color of status embed iframe text when displaying incident", "t": "`$STRING`", "key$": "incident_text_color", "index$": 1 }, "maintenance_background_color": { "a": true, "h": "Maintenance Background Color", "n": "maintenance_background_color", "r": false, "sh": "Color of status embed iframe background when displaying maintenance", "t": "`$STRING`", "key$": "maintenance_background_color", "index$": 2 }, "maintenance_text_color": { "a": true, "h": "Maintenance Text Color", "n": "maintenance_text_color", "r": false, "sh": "Color of status embed iframe text when displaying maintenance", "t": "`$STRING`", "key$": "maintenance_text_color", "index$": 3 }, "page_id": { "a": true, "h": "Page Id", "n": "page_id", "r": false, "sh": "Page identifier", "t": "`$STRING`", "key$": "page_id", "index$": 4 }, "position": { "a": true, "h": "Position", "n": "position", "r": false, "sh": "Corner where status embed iframe will appear on page", "t": "`$STRING`", "key$": "position", "index$": 5 }, "status_embed_config": { "a": true, "h": "Status Embed Config", "n": "status_embed_config", "r": false, "t": "`$OBJECT`", "key$": "status_embed_config", "index$": 6 } }, "name": "status_embed_config", "op": { "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/status_embed_config", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/status_embed_config", "q": { "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "status_embed_config" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /pages/{page_id}/status_embed_config", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}/status_embed_config", "q": { "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "status_embed_config" }], "t": { "req": { "status_embed_config": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "patch" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /pages/{page_id}/status_embed_config", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}/status_embed_config", "q": { "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "status_embed_config" }], "t": { "req": { "status_embed_config": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.page"]] }, "key$": "status_embed_config", "name__orig": "status_embed_config", "Name": "StatusEmbedConfig", "name_": "status_embed_config", "name-": "status-embed-config", "NAME": "STATUS_EMBED_CONFIG", "index$": 14 }, { "active": true, "entity": "status_embed_config", "key$": "BasicStatusEmbedConfigFlow", "kind": "basic", "name": "BasicStatusEmbedConfigFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "status_embed_config_ref01", "srcdatavar": "status_embed_config_ref01_data", "suffix": "_up0", "textfield": "incident_background_color" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-status_embed_config_ref01" } }], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "status_embed_config_ref01", "srcdatavar": "status_embed_config_ref01_data", "suffix": "_dt0" }, "m": { "id": "status_embed_config01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-status_embed_config_ref01" } }], "index$": 1 }] }, 'StatusEmbedConfig', { "GET /pages/{page_id}/status_embed_config": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "PATCH /pages/{page_id}/status_embed_config": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "status_embed_config": { "type": "object", "properties": { "position": { "type": "string", "description": "Corner where status embed iframe will appear on page" }, "incident_background_color": { "type": "string", "description": "Color of status embed iframe background when displaying incident" }, "incident_text_color": { "type": "string", "description": "Color of status embed iframe text when displaying incident" }, "maintenance_background_color": { "type": "string", "description": "Color of status embed iframe background when displaying maintenance" }, "maintenance_text_color": { "type": "string", "description": "Color of status embed iframe text when displaying maintenance" } }, "key$": "status_embed_config" } }, "description": "Update status embed config settings", "x-ref": "#/components/schemas/patchPagesPageIdStatusEmbedConfig", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "PUT /pages/{page_id}/status_embed_config": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "status_embed_config": { "type": "object", "properties": { "position": { "type": "string", "description": "Corner where status embed iframe will appear on page" }, "incident_background_color": { "type": "string", "description": "Color of status embed iframe background when displaying incident" }, "incident_text_color": { "type": "string", "description": "Color of status embed iframe text when displaying incident" }, "maintenance_background_color": { "type": "string", "description": "Color of status embed iframe background when displaying maintenance" }, "maintenance_text_color": { "type": "string", "description": "Color of status embed iframe text when displaying maintenance" } }, "key$": "status_embed_config" } }, "description": "Update status embed config settings", "x-ref": "#/components/schemas/putPagesPageIdStatusEmbedConfig", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let status_embed_config_ref01_data = Object.values(setup.data.existing.status_embed_config)[0];
        // UPDATE
        const status_embed_config_ref01_ent = client.StatusEmbedConfig();
        const status_embed_config_ref01_data_up0 = {};
        const status_embed_config_ref01_markdef_up0 = { name: 'incident_background_color', value: 'Mark01-status_embed_config_ref01_' + setup.now };
        status_embed_config_ref01_data_up0[status_embed_config_ref01_markdef_up0.name] = status_embed_config_ref01_markdef_up0.value;
        const status_embed_config_ref01_resdata_up0 = (await status_embed_config_ref01_ent.update(status_embed_config_ref01_data_up0)).data();
        (0, node_assert_1.default)(null != status_embed_config_ref01_resdata_up0);
        (0, node_assert_1.default)(status_embed_config_ref01_resdata_up0[status_embed_config_ref01_markdef_up0.name] === status_embed_config_ref01_markdef_up0.value);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/status_embed_config/StatusEmbedConfigTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['status_embed_config01', 'status_embed_config02', 'status_embed_config03', 'page01', 'page02', 'page03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID'];
        idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {};
        if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
            throw new Error('Live ENTID must be a JSON object');
        }
        client = new __1.StatuspageSDK(merge([
            // FIRST, so the generated fields below win: sdk-test-control.json's
            // test.client.options adds to the live client, it does not redirect it.
            (0, utility_1.liveClientOptions)(),
            {
                apikey: env.STATUSPAGE_APIKEY,
            },
            // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
            // last entry is undefined, and basicSetup is normally called with no
            // argument at all - so a bare 'extra' silently discarded the apikey
            // and server values above and handed the SDK undefined. Harmless
            // while there was nothing in that object; not harmless now.
            extra || {},
            { system: { fetch: transport.fetch } }
        ]));
    }
    const setup = {
        idmap,
        env,
        options,
        client,
        struct,
        data: entityData,
        explain: 'TRUE' === env.STATUSPAGE_TEST_EXPLAIN,
        live,
        transport,
        now: Date.now(),
    };
    return setup;
}
//# sourceMappingURL=StatusEmbedConfigEntity.test.js.map