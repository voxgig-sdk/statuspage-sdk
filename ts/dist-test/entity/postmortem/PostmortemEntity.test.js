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
(0, node_test_1.describe)('PostmortemEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.Postmortem();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['update', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'postmortem.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "body", "req": false, "short": "Postmortem body", "type": "`$STRING`", "index$": 0 }, { "active": true, "name": "body_draft", "req": false, "short": "Body draft", "type": "`$STRING`", "index$": 1 }, { "active": true, "format": "date-time", "name": "body_draft_updated_at", "req": false, "type": "`$STRING`", "index$": 2 }, { "active": true, "format": "date-time", "name": "body_updated_at", "req": false, "type": "`$STRING`", "index$": 3 }, { "active": true, "format": "date-time", "name": "created_at", "req": false, "type": "`$STRING`", "index$": 4 }, { "active": true, "name": "custom_tweet", "req": false, "short": "Custom tweet for Incident Postmortem", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "notify_subscribers", "req": false, "short": "Should email subscribers be notified.", "type": "`$BOOLEAN`", "index$": 6 }, { "active": true, "name": "notify_twitter", "req": false, "short": "Should Twitter followers be notified.", "type": "`$BOOLEAN`", "index$": 7 }, { "active": true, "name": "postmortem", "req": true, "type": "`$OBJECT`", "index$": 8 }, { "active": true, "name": "preview_key", "req": false, "short": "Preview Key", "type": "`$STRING`", "index$": 9 }, { "active": true, "format": "date-time", "name": "published_at", "req": false, "type": "`$STRING`", "index$": 10 }, { "active": true, "format": "date-time", "name": "updated_at", "req": false, "type": "`$STRING`", "index$": 11 }], "name": "postmortem", "op": { "load": { "input": "data", "name": "load", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "incident_id", "orig": "incident_id", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "param", "name": "page_id", "orig": "page_id", "reqd": true, "type": "`$STRING`", "index$": 1 }] }, "contract": { "id": "GET /pages/{page_id}/incidents/{incident_id}/postmortem", "json": "{\"operationId\":\"getPagesPageIdIncidentsIncidentIdPostmortem\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Identifier\",\"in\":\"path\",\"name\":\"incident_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Revert Postmortem\",\"properties\":{\"body\":{\"description\":\"Postmortem body\",\"type\":\"string\"},\"body_draft\":{\"description\":\"Body draft\",\"type\":\"string\"},\"body_draft_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"body_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"custom_tweet\":{\"description\":\"Custom tweet for Incident Postmortem\",\"type\":\"string\"},\"notify_subscribers\":{\"description\":\"Should email subscribers be notified.\",\"type\":\"boolean\"},\"notify_twitter\":{\"description\":\"Should Twitter followers be notified.\",\"type\":\"boolean\"},\"preview_key\":{\"description\":\"Preview Key\",\"type\":\"string\"},\"published_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Get Postmortem\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "GET", "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem", "segments": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incidents" }, { "var": "incident_id" }, { "lit": "postmortem" }], "select": { "exist": ["incident_id", "page_id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "update": { "input": "data", "name": "update", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "incident_id", "orig": "incident_id", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "param", "name": "page_id", "orig": "page_id", "reqd": true, "type": "`$STRING`", "index$": 1 }] }, "contract": { "id": "PUT /pages/{page_id}/incidents/{incident_id}/postmortem", "json": "{\"operationId\":\"putPagesPageIdIncidentsIncidentIdPostmortem\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Identifier\",\"in\":\"path\",\"name\":\"incident_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Create Postmortem\",\"properties\":{\"postmortem\":{\"properties\":{\"body_draft\":{\"description\":\"Body of Postmortem to create.\",\"type\":\"string\"}},\"required\":[\"body_draft\"],\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Revert Postmortem\",\"properties\":{\"body\":{\"description\":\"Postmortem body\",\"type\":\"string\"},\"body_draft\":{\"description\":\"Body draft\",\"type\":\"string\"},\"body_draft_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"body_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"custom_tweet\":{\"description\":\"Custom tweet for Incident Postmortem\",\"type\":\"string\"},\"notify_subscribers\":{\"description\":\"Should email subscribers be notified.\",\"type\":\"boolean\"},\"notify_twitter\":{\"description\":\"Should Twitter followers be notified.\",\"type\":\"boolean\"},\"preview_key\":{\"description\":\"Preview Key\",\"type\":\"string\"},\"published_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Create Postmortem\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "PUT", "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem", "segments": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incidents" }, { "var": "incident_id" }, { "lit": "postmortem" }], "select": { "exist": ["incident_id", "page_id"] }, "transform": { "req": { "postmortem": "`reqdata`" }, "res": "`body`" }, "index$": 0 }, { "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "incident_id", "orig": "incident_id", "reqd": true, "type": "`$STRING`" }, { "active": true, "kind": "param", "name": "page_id", "orig": "page_id", "reqd": true, "type": "`$STRING`" }] }, "contract": { "id": "PUT /pages/{page_id}/incidents/{incident_id}/postmortem/publish", "json": "{\"operationId\":\"putPagesPageIdIncidentsIncidentIdPostmortemPublish\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Identifier\",\"in\":\"path\",\"name\":\"incident_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Publish Postmortem\",\"properties\":{\"postmortem\":{\"properties\":{\"custom_tweet\":{\"description\":\"Custom postmortem tweet to publish\",\"type\":\"string\"},\"notify_subscribers\":{\"description\":\"Whether to notify e-mail subscribers\",\"type\":\"boolean\"},\"notify_twitter\":{\"description\":\"Whether to notify Twitter followers\",\"type\":\"boolean\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Revert Postmortem\",\"properties\":{\"body\":{\"description\":\"Postmortem body\",\"type\":\"string\"},\"body_draft\":{\"description\":\"Body draft\",\"type\":\"string\"},\"body_draft_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"body_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"custom_tweet\":{\"description\":\"Custom tweet for Incident Postmortem\",\"type\":\"string\"},\"notify_subscribers\":{\"description\":\"Should email subscribers be notified.\",\"type\":\"boolean\"},\"notify_twitter\":{\"description\":\"Should Twitter followers be notified.\",\"type\":\"boolean\"},\"preview_key\":{\"description\":\"Preview Key\",\"type\":\"string\"},\"published_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Publish Postmortem\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "PUT", "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/publish", "segments": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incidents" }, { "var": "incident_id" }, { "lit": "postmortem" }, { "lit": "publish" }], "select": { "$action": "publish", "exist": ["incident_id", "page_id"] }, "transform": { "req": { "postmortem": "`reqdata`" }, "res": "`body`" }, "index$": 1 }, { "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "incident_id", "orig": "incident_id", "reqd": true, "type": "`$STRING`" }, { "active": true, "kind": "param", "name": "page_id", "orig": "page_id", "reqd": true, "type": "`$STRING`" }] }, "contract": { "id": "PUT /pages/{page_id}/incidents/{incident_id}/postmortem/revert", "json": "{\"operationId\":\"putPagesPageIdIncidentsIncidentIdPostmortemRevert\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Identifier\",\"in\":\"path\",\"name\":\"incident_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Revert Postmortem\",\"properties\":{\"body\":{\"description\":\"Postmortem body\",\"type\":\"string\"},\"body_draft\":{\"description\":\"Body draft\",\"type\":\"string\"},\"body_draft_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"body_updated_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"custom_tweet\":{\"description\":\"Custom tweet for Incident Postmortem\",\"type\":\"string\"},\"notify_subscribers\":{\"description\":\"Should email subscribers be notified.\",\"type\":\"boolean\"},\"notify_twitter\":{\"description\":\"Should Twitter followers be notified.\",\"type\":\"boolean\"},\"preview_key\":{\"description\":\"Preview Key\",\"type\":\"string\"},\"published_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Revert Postmortem\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "PUT", "orig": "/pages/{page_id}/incidents/{incident_id}/postmortem/revert", "segments": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incidents" }, { "var": "incident_id" }, { "lit": "postmortem" }, { "lit": "revert" }], "select": { "$action": "revert", "exist": ["incident_id", "page_id"] }, "transform": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "update" } }, "relations": { "ancestors": [["page", "incident"]] }, "key$": "postmortem", "name__orig": "postmortem", "Name": "Postmortem", "name_": "postmortem", "name-": "postmortem", "NAME": "POSTMORTEM", "index$": 14 }, { "active": true, "entity": "postmortem", "key$": "BasicPostmortemFlow", "kind": "basic", "name": "BasicPostmortemFlow", "param": {}, "step": [{ "active": true, "data": { "page_id": "page01" }, "input": { "ref": "postmortem_ref01", "srcdatavar": "postmortem_ref01_data", "suffix": "_up0", "textfield": "body" }, "match": {}, "op": "update", "spec": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-postmortem_ref01" } }], "valid": [], "index$": 0 }, { "active": true, "data": {}, "input": { "ref": "postmortem_ref01", "srcdatavar": "postmortem_ref01_data", "suffix": "_dt0" }, "match": { "id": "postmortem01", "page_id": "page01" }, "op": "load", "spec": [], "valid": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-postmortem_ref01" } }], "index$": 1 }] }, 'Postmortem');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let postmortem_ref01_data = Object.values(setup.data.existing.postmortem)[0];
        // UPDATE
        const postmortem_ref01_ent = client.Postmortem();
        const postmortem_ref01_data_up0 = {};
        postmortem_ref01_data_up0['page_id'] = setup.idmap['page_id'];
        const postmortem_ref01_markdef_up0 = { name: 'body', value: 'Mark01-postmortem_ref01_' + setup.now };
        postmortem_ref01_data_up0[postmortem_ref01_markdef_up0.name] = postmortem_ref01_markdef_up0.value;
        const postmortem_ref01_resdata_up0 = (await postmortem_ref01_ent.update(postmortem_ref01_data_up0)).data();
        (0, node_assert_1.default)(null != postmortem_ref01_resdata_up0);
        (0, node_assert_1.default)(postmortem_ref01_resdata_up0[postmortem_ref01_markdef_up0.name] === postmortem_ref01_markdef_up0.value);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/postmortem/PostmortemTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['postmortem01', 'postmortem02', 'postmortem03', 'page01', 'page02', 'page03', 'incident01', 'incident02', 'incident03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_POSTMORTEM_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_POSTMORTEM_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_POSTMORTEM_ENTID'];
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
//# sourceMappingURL=PostmortemEntity.test.js.map