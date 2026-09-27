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
(0, node_test_1.describe)('GroupComponentEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.GroupComponent();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['create', 'list', 'update', 'load', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'group_component.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "component_group": { "a": true, "h": "Component Group", "n": "component_group", "r": true, "t": "`$OBJECT`", "key$": "component_group", "index$": 0 }, "components": { "a": true, "h": "Components", "n": "components", "r": false, "t": "`$STRING`", "key$": "components", "index$": 1 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": false, "t": "`$STRING`", "key$": "created_at", "index$": 2 }, "description": { "a": true, "h": "Description", "n": "description", "r": false, "sh": "Description of the component group.", "t": "`$STRING`", "key$": "description", "index$": 3 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Component Group Identifier", "t": "`$STRING`", "key$": "id", "index$": 4 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "t": "`$STRING`", "key$": "name", "index$": 5 }, "page_id": { "a": true, "h": "Page Id", "n": "page_id", "r": false, "t": "`$STRING`", "key$": "page_id", "index$": 6 }, "position": { "a": true, "h": "Position", "n": "position", "r": false, "t": "`$STRING`", "key$": "position", "index$": 7 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": false, "t": "`$STRING`", "key$": "updated_at", "index$": 8 } }, "id": { "field": "id", "name": "id" }, "name": "group_component", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /pages/{page_id}/component-groups", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/component-groups", "q": { "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "component-groups" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/component-groups", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/component-groups", "q": { "exist": ["page", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "component-groups" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/component-groups/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/component-groups/{id}", "q": { "exist": ["id", "page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "component-groups" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /pages/{page_id}/component-groups/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}/component-groups/{id}", "q": { "exist": ["id", "page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "component-groups" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "patch" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /pages/{page_id}/component-groups/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/component-groups/{id}", "q": { "exist": ["id", "page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "component-groups" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /pages/{page_id}/component-groups/{id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}/component-groups/{id}", "q": { "exist": ["id", "page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "component-groups" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.page"]] }, "key$": "group_component", "name__orig": "group_component", "Name": "GroupComponent", "name_": "group_component", "name-": "group-component", "NAME": "GROUP_COMPONENT", "index$": 2 }, { "active": true, "entity": "group_component", "key$": "BasicGroupComponentFlow", "kind": "basic", "name": "BasicGroupComponentFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "group_component_ref01" }, "m": { "page_id": "page01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "group_component_ref01" } }], "index$": 1 }, { "a": true, "d": { "page_id": "page01" }, "i": { "ref": "group_component_ref01", "srcdatavar": "group_component_ref01_data", "suffix": "_up0", "textfield": "components" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-group_component_ref01" } }], "v": [], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "group_component_ref01", "srcdatavar": "group_component_ref01_data", "suffix": "_dt0" }, "m": { "id": "group_component01", "page_id": "page01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-group_component_ref01" } }], "index$": 3 }, { "a": true, "d": {}, "i": { "ref": "group_component_ref01", "suffix": "_rm0" }, "m": { "id": "group_component01", "page_id": "page01" }, "o": "remove", "s": [], "v": [], "index$": 4 }, { "a": true, "d": {}, "i": { "suffix": "_rt0" }, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemNotExists", "def": { "ref": "group_component_ref01" } }], "index$": 5 }] }, 'GroupComponent', { "POST /pages/{page_id}/component-groups": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "description": { "type": "string", "description": "Description of the component group.", "key$": "description" }, "component_group": { "type": "object", "properties": { "components": { "type": "array", "items": { "type": "string" } }, "name": { "type": "string" } }, "required": ["components", "name"], "key$": "component_group" } }, "description": "Create a component group", "x-ref": "#/components/schemas/postPagesPageIdComponentGroups", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "GET /pages/{page_id}/component-groups": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 1 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }] }, "GET /pages/{page_id}/component-groups/{id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "id", "description": "Component group identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PATCH /pages/{page_id}/component-groups/{id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "description": { "type": "string", "description": "Updated description of the component group.", "key$": "description" }, "component_group": { "type": "object", "properties": { "components": { "type": "array", "items": { "type": "string" } }, "name": { "type": "string" } }, "required": ["components", "name"], "key$": "component_group" } }, "description": "Update a component group", "x-ref": "#/components/schemas/patchPagesPageIdComponentGroups", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "id", "description": "Component group identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/component-groups/{id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "id", "description": "Component group identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PUT /pages/{page_id}/component-groups/{id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "description": { "type": "string", "description": "Updated description of the component group.", "key$": "description" }, "component_group": { "type": "object", "properties": { "components": { "type": "array", "items": { "type": "string" } }, "name": { "type": "string" } }, "required": ["components", "name"], "key$": "component_group" } }, "description": "Update a component group", "x-ref": "#/components/schemas/putPagesPageIdComponentGroups", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "id", "description": "Component group identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const group_component_ref01_ent = client.GroupComponent();
        let group_component_ref01_data = setup.data.new.group_component['group_component_ref01'];
        group_component_ref01_data['page_id'] = setup.idmap['page01'];
        group_component_ref01_data = (await group_component_ref01_ent.create(group_component_ref01_data)).data();
        (0, node_assert_1.default)(null != group_component_ref01_data.id);
        // LIST
        const group_component_ref01_match = {};
        group_component_ref01_match['page_id'] = setup.idmap['page01'];
        const group_component_ref01_list = (await group_component_ref01_ent.list(group_component_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(group_component_ref01_list, { id: group_component_ref01_data.id })));
        // UPDATE
        const group_component_ref01_data_up0 = {};
        group_component_ref01_data_up0.id = group_component_ref01_data.id;
        group_component_ref01_data_up0['page_id'] = setup.idmap['page_id'];
        const group_component_ref01_markdef_up0 = { name: 'components', value: 'Mark01-group_component_ref01_' + setup.now };
        group_component_ref01_data_up0[group_component_ref01_markdef_up0.name] = group_component_ref01_markdef_up0.value;
        const group_component_ref01_resdata_up0 = (await group_component_ref01_ent.update(group_component_ref01_data_up0)).data();
        (0, node_assert_1.default)(group_component_ref01_resdata_up0.id === group_component_ref01_data_up0.id);
        (0, node_assert_1.default)(group_component_ref01_resdata_up0[group_component_ref01_markdef_up0.name] === group_component_ref01_markdef_up0.value);
        // LOAD
        const group_component_ref01_match_dt0 = {};
        group_component_ref01_match_dt0.id = group_component_ref01_data.id;
        const group_component_ref01_data_dt0 = (await group_component_ref01_ent.load(group_component_ref01_match_dt0)).data();
        (0, node_assert_1.default)(group_component_ref01_data_dt0.id === group_component_ref01_data.id);
        // REMOVE
        const group_component_ref01_match_rm0 = { id: group_component_ref01_data.id };
        await group_component_ref01_ent.remove(group_component_ref01_match_rm0);
        // LIST
        const group_component_ref01_match_rt0 = {};
        group_component_ref01_match_rt0['page_id'] = setup.idmap['page01'];
        const group_component_ref01_list_rt0 = (await group_component_ref01_ent.list(group_component_ref01_match_rt0)).map((e) => e.data());
        (0, node_assert_1.default)(isempty(select(group_component_ref01_list_rt0, { id: group_component_ref01_data.id })));
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/group_component/GroupComponentTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['group_component01', 'group_component02', 'group_component03', 'page01', 'page02', 'page03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_GROUP_COMPONENT_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_GROUP_COMPONENT_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_GROUP_COMPONENT_ENTID'];
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
//# sourceMappingURL=GroupComponentEntity.test.js.map