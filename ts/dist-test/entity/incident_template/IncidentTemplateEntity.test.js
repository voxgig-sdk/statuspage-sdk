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
(0, node_test_1.describe)('IncidentTemplateEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.IncidentTemplate();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['create', 'list']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'incident_template.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "body": { "a": true, "h": "Body", "n": "body", "r": false, "sh": "Body of the incident or maintenance update to be applied when selecting this template", "t": "`$STRING`", "key$": "body", "index$": 0 }, "components": { "a": true, "h": "Components", "n": "components", "r": false, "sh": "Affected components", "t": "`$ARRAY`", "key$": "components", "index$": 1 }, "group_id": { "a": true, "h": "Group Id", "n": "group_id", "r": false, "sh": "Identifier of Template Group this template belongs to", "t": "`$STRING`", "key$": "group_id", "index$": 2 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Incident Template Identifier", "t": "`$STRING`", "key$": "id", "index$": 3 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "sh": "Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page", "t": "`$STRING`", "key$": "name", "index$": 4 }, "should_send_notifications": { "a": true, "h": "Should Send Notifications", "n": "should_send_notifications", "r": false, "sh": "Whether the \"deliver notifications\" checkbox should be selected when selecting this template", "t": "`$BOOLEAN`", "key$": "should_send_notifications", "index$": 5 }, "should_tweet": { "a": true, "h": "Should Tweet", "n": "should_tweet", "r": false, "sh": "Whether the \"tweet update\" checkbox should be selected when selecting this template", "t": "`$BOOLEAN`", "key$": "should_tweet", "index$": 6 }, "template": { "a": true, "h": "Template", "n": "template", "r": true, "t": "`$OBJECT`", "key$": "template", "index$": 7 }, "title": { "a": true, "h": "Title", "n": "title", "r": false, "sh": "Title to be applied to the incident or maintenance when selecting this template", "t": "`$STRING`", "key$": "title", "index$": 8 }, "update_status": { "a": true, "h": "Update Status", "n": "update_status", "r": false, "sh": "The status the incident or maintenance should transition to when selecting this template", "t": "`$STRING`", "key$": "update_status", "index$": 9 } }, "id": { "field": "id", "name": "id" }, "name": "incident_template", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /pages/{page_id}/incident_templates", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/incident_templates", "q": { "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incident_templates" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/incident_templates", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "ex": 1, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "ex": 100, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/incident_templates", "q": { "exist": ["page", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incident_templates" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" } }, "relations": { "ancestors": [["$.main.kit.entity.page"]] }, "key$": "incident_template", "name__orig": "incident_template", "Name": "IncidentTemplate", "name_": "incident_template", "name-": "incident-template", "NAME": "INCIDENT_TEMPLATE", "index$": 5 }, { "active": true, "entity": "incident_template", "key$": "BasicIncidentTemplateFlow", "kind": "basic", "name": "BasicIncidentTemplateFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "incident_template_ref01" }, "m": { "page_id": "page01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "incident_template_ref01" } }], "index$": 1 }] }, 'IncidentTemplate', { "POST /pages/{page_id}/incident_templates": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "template": { "type": "object", "properties": { "name": { "type": "string", "description": "Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page" }, "title": { "type": "string", "description": "Title to be applied to the incident or maintenance when selecting this template" }, "body": { "type": "string", "description": "The initial message, created as the first incident or maintenance update." }, "group_id": { "type": "string", "description": "Identifier of Template Group this template belongs to" }, "update_status": { "type": "string", "description": "The status the incident or maintenance should transition to when selecting this template", "enum": ["investigating", "identified", "monitoring", "resolved", "scheduled", "in_progress", "verifying", "completed"] }, "should_tweet": { "type": "boolean", "description": "Whether the \"tweet update\" checkbox should be selected when selecting this template" }, "should_send_notifications": { "type": "boolean", "description": "Whether the \"deliver notifications\" checkbox should be selected when selecting this template" }, "component_ids": { "type": "array", "description": "List of component_ids affected by this incident", "items": { "type": "string", "default": "1ss1vgkq6034" } } }, "required": ["name", "title", "body"], "key$": "template" } }, "description": "Create a template", "x-ref": "#/components/schemas/postPagesPageIdIncidentTemplates", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "GET /pages/{page_id}/incident_templates": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "query", "name": "page", "description": "Page offset to fetch.", "required": false, "schema": { "type": "integer", "format": "int32", "default": 1 }, "index$": 1 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page.", "required": false, "schema": { "type": "integer", "format": "int32", "default": 100 }, "index$": 2 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const incident_template_ref01_ent = client.IncidentTemplate();
        let incident_template_ref01_data = setup.data.new.incident_template['incident_template_ref01'];
        incident_template_ref01_data['page_id'] = setup.idmap['page01'];
        incident_template_ref01_data = (await incident_template_ref01_ent.create(incident_template_ref01_data)).data();
        (0, node_assert_1.default)(null != incident_template_ref01_data.id);
        // LIST
        const incident_template_ref01_match = {};
        incident_template_ref01_match['page_id'] = setup.idmap['page01'];
        const incident_template_ref01_list = (await incident_template_ref01_ent.list(incident_template_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(incident_template_ref01_list, { id: incident_template_ref01_data.id })));
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/incident_template/IncidentTemplateTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['incident_template01', 'incident_template02', 'incident_template03', 'page01', 'page02', 'page03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID'];
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
//# sourceMappingURL=IncidentTemplateEntity.test.js.map