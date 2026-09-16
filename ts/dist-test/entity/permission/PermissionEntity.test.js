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
// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
(0, utility_1.loadEnvLocal)(__dirname + '/../../../.env.local');
(0, node_test_1.describe)('PermissionEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.Permission();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['update', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'permission.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "id", "req": false, "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "pages", "req": false, "short": "Pages accessible by the user.", "type": "`$OBJECT`", "index$": 1 }, { "active": true, "name": "user_id", "req": false, "short": "User identifier", "type": "`$STRING`", "index$": 2 }], "id": { "field": "id", "name": "id" }, "name": "permission", "op": { "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "user_id", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "param", "name": "organization_id", "orig": "organization_id", "reqd": true, "type": "`$STRING`", "index$": 1 }] }, "contract": { "id": "GET /organizations/{organization_id}/permissions/{user_id}", "json": "{\"operationId\":\"getOrganizationsOrganizationIdPermissionsUserId\",\"parameters\":[{\"description\":\"Organization Identifier\",\"in\":\"path\",\"name\":\"organization_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"User identifier\",\"in\":\"path\",\"name\":\"user_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"examples\":{\"response\":{\"value\":{\"pages\":[{\"incident_manager\":true,\"maintenance_manager\":true,\"page_configuration\":false,\"page_id\":\"dlqpc4stl3t6\"},{\"incident_manager\":true,\"maintenance_manager\":true,\"page_configuration\":false,\"page_id\":\"ngdkgwfj7d15\"}],\"user_id\":\"xkjqwf07j7kg\"}}},\"schema\":{\"description\":\"Get a user's permissions\",\"properties\":{\"data\":{\"properties\":{\"pages\":{\"description\":\"Pages accessible by the user.\",\"properties\":{\"incident_manager\":{\"description\":\"User has incident manager role. This field will only be present if the organization has Role Based Access Control.\",\"type\":\"boolean\"},\"maintenance_manager\":{\"description\":\"User has maintenance manager role. This field will only be present if the organization has Role Based Access Control.\",\"type\":\"boolean\"},\"page_configuration\":{\"description\":\"User has page configuration role. This field will only be present if the organization has Role Based Access Control.\",\"type\":\"boolean\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"}},\"type\":\"object\"},\"user_id\":{\"description\":\"User identifier\",\"type\":\"string\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"description\":\"Get a user's permissions\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/organizations/{organization_id}/permissions/{user_id}", "rename": { "param": { "user_id": "id" } }, "segments": [{ "lit": "organizations" }, { "var": "organization_id" }, { "lit": "permissions" }, { "var": "id" }], "select": { "exist": ["id", "organization_id"] }, "transform": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }], "key$": "load" }, "update": { "input": "data", "name": "update", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "user_id", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "param", "name": "organization_id", "orig": "organization_id", "reqd": true, "type": "`$STRING`", "index$": 1 }] }, "contract": { "id": "PUT /organizations/{organization_id}/permissions/{user_id}", "json": "{\"operationId\":\"putOrganizationsOrganizationIdPermissionsUserId\",\"parameters\":[{\"description\":\"Organization Identifier\",\"in\":\"path\",\"name\":\"organization_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"User identifier\",\"in\":\"path\",\"name\":\"user_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update a user's role permissions. Payload should contain a mapping of pages to a set of the desired roles,\\n                  if the page has Role Based Access Control. Otherwise, the pages should map to an empty hash.\\n                  User will lose access to any pages omitted from the payload.\",\"properties\":{\"pages\":{\"properties\":{\"page_id\":{\"properties\":{\"incident_manager\":{\"description\":\"Whether or not user should have incident manager role. This field will only be present for pages with Role Based Access Control.\",\"type\":\"boolean\"},\"maintenance_manager\":{\"description\":\"Whether or not user should have maintenance manager role. This field will only be present for pages with Role Based Access Control.\",\"type\":\"boolean\"},\"page_configuration\":{\"description\":\"Whether or not user should have page configuration role. This field will only be present for pages with Role Based Access Control.\",\"type\":\"boolean\"}},\"type\":\"object\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"examples\":{\"response\":{\"value\":{\"pages\":[{\"incident_manager\":true,\"maintenance_manager\":true,\"page_configuration\":false,\"page_id\":\"04qbfq8lvpp3\"},{\"incident_manager\":true,\"maintenance_manager\":true,\"page_configuration\":false,\"page_id\":\"nz9bqh8d0hnl\"}],\"user_id\":\"wg5rb3fzsg2n\"}}},\"schema\":{\"description\":\"Get a user's permissions\",\"properties\":{\"data\":{\"properties\":{\"pages\":{\"description\":\"Pages accessible by the user.\",\"properties\":{\"incident_manager\":{\"description\":\"User has incident manager role. This field will only be present if the organization has Role Based Access Control.\",\"type\":\"boolean\"},\"maintenance_manager\":{\"description\":\"User has maintenance manager role. This field will only be present if the organization has Role Based Access Control.\",\"type\":\"boolean\"},\"page_configuration\":{\"description\":\"User has page configuration role. This field will only be present if the organization has Role Based Access Control.\",\"type\":\"boolean\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"}},\"type\":\"object\"},\"user_id\":{\"description\":\"User identifier\",\"type\":\"string\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"description\":\"Update a user's role permissions. Payload should contain a mapping of pages to a set of the desired roles,\\n                  if the page has Role Based Access Control. Otherwise, the pages should map to an empty hash.\\n                  User will lose access to any pages omitted from the payload.\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"403\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"You are not authorized to access this resource.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "PUT", "orig": "/organizations/{organization_id}/permissions/{user_id}", "rename": { "param": { "user_id": "id" } }, "segments": [{ "lit": "organizations" }, { "var": "organization_id" }, { "lit": "permissions" }, { "var": "id" }], "select": { "exist": ["id", "organization_id"] }, "transform": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["organization"]] }, "key$": "permission", "name__orig": "permission", "Name": "Permission", "name_": "permission", "name-": "permission", "NAME": "PERMISSION", "index$": 13 }, { "active": true, "entity": "permission", "key$": "BasicPermissionFlow", "kind": "basic", "name": "BasicPermissionFlow", "param": {}, "step": [{ "active": true, "data": { "organization_id": "organization01" }, "input": { "ref": "permission_ref01", "srcdatavar": "permission_ref01_data", "suffix": "_up0", "textfield": "user_id" }, "match": {}, "op": "update", "spec": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-permission_ref01" } }], "valid": [], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "permission_ref01", "srcdatavar": "permission_ref01_data", "suffix": "_dt0" }, "match": { "id": "permission01", "organization_id": "organization01" }, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-permission_ref01" } }], "index$": 1 }] }, 'Permission');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let permission_ref01_data = Object.values(setup.data.existing.permission)[0];
        // UPDATE
        const permission_ref01_ent = client.Permission();
        const permission_ref01_data_up0 = {};
        permission_ref01_data_up0.id = permission_ref01_data.id;
        permission_ref01_data_up0['organization_id'] = setup.idmap['organization_id'];
        const permission_ref01_markdef_up0 = { name: 'user_id', value: 'Mark01-permission_ref01_' + setup.now };
        permission_ref01_data_up0[permission_ref01_markdef_up0.name] = permission_ref01_markdef_up0.value;
        const permission_ref01_resdata_up0 = (await permission_ref01_ent.update(permission_ref01_data_up0)).data();
        (0, node_assert_1.default)(permission_ref01_resdata_up0.id === permission_ref01_data_up0.id);
        (0, node_assert_1.default)(permission_ref01_resdata_up0[permission_ref01_markdef_up0.name] === permission_ref01_markdef_up0.value);
        // LOAD
        const permission_ref01_match_dt0 = {};
        permission_ref01_match_dt0.id = permission_ref01_data.id;
        const permission_ref01_data_dt0 = (await permission_ref01_ent.load(permission_ref01_match_dt0)).data();
        (0, node_assert_1.default)(permission_ref01_data_dt0.id === permission_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/permission/PermissionTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['permission01', 'permission02', 'permission03', 'organization01', 'organization02', 'organization03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_PERMISSION_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_PERMISSION_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_PERMISSION_ENTID'];
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
//# sourceMappingURL=PermissionEntity.test.js.map