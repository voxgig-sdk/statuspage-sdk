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
(0, node_test_1.describe)('IncidentUpdateEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.IncidentUpdate();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['update']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'incident_update.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": [{ "active": true, "name": "affected_components", "req": false, "short": "Affected components associated with the incident update.", "type": "`$ARRAY`", "index$": 0 }, { "active": true, "name": "body", "req": false, "short": "Incident update body.", "type": "`$STRING`", "index$": 1 }, { "active": true, "format": "date-time", "name": "created_at", "req": false, "short": "The timestamp when the incident update was created at.", "type": "`$STRING`", "index$": 2 }, { "active": true, "name": "custom_tweet", "req": false, "short": "An optional customized tweet message for incident postmortem.", "type": "`$STRING`", "index$": 3 }, { "active": true, "name": "deliver_notifications", "req": false, "short": "Controls whether to delivery notifications.", "type": "`$BOOLEAN`", "index$": 4 }, { "active": true, "format": "date-time", "name": "display_at", "req": false, "short": "Timestamp when incident update is happened.", "type": "`$STRING`", "index$": 5 }, { "active": true, "name": "id", "req": false, "short": "Incident Update Identifier.", "type": "`$STRING`", "index$": 6 }, { "active": true, "name": "incident_id", "req": false, "short": "Incident Identifier.", "type": "`$STRING`", "index$": 7 }, { "active": true, "name": "incident_update", "req": false, "type": "`$OBJECT`", "index$": 8 }, { "active": true, "name": "status", "req": false, "short": "The incident status.", "type": "`$STRING`", "index$": 9 }, { "active": true, "name": "tweet_id", "req": false, "short": "Tweet identifier associated to this incident update.", "type": "`$STRING`", "index$": 10 }, { "active": true, "format": "date-time", "name": "twitter_updated_at", "req": false, "short": "The timestamp when twitter updated at.", "type": "`$STRING`", "index$": 11 }, { "active": true, "format": "date-time", "name": "updated_at", "req": false, "short": "The timestamp when the incident update is updated.", "type": "`$STRING`", "index$": 12 }, { "active": true, "name": "wants_twitter_update", "req": false, "short": "Controls whether to create twitter update.", "type": "`$BOOLEAN`", "index$": 13 }], "id": { "field": "id", "name": "id" }, "name": "incident_update", "op": { "patch": { "input": "data", "name": "patch", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "incident_update_id", "reqd": true, "type": "`$STRING`" }, { "active": true, "kind": "param", "name": "incident_id", "orig": "incident_id", "reqd": true, "type": "`$STRING`" }, { "active": true, "kind": "param", "name": "page_id", "orig": "page_id", "reqd": true, "type": "`$STRING`" }] }, "contract": { "id": "PATCH /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}", "json": "{\"operationId\":\"patchPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Identifier\",\"in\":\"path\",\"name\":\"incident_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Update Identifier\",\"in\":\"path\",\"name\":\"incident_update_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update a previous incident update\",\"properties\":{\"incident_update\":{\"properties\":{\"body\":{\"description\":\"Incident update body.\",\"type\":\"string\"},\"deliver_notifications\":{\"description\":\"Controls whether to delivery notifications.\",\"type\":\"boolean\"},\"display_at\":{\"description\":\"Timestamp when incident update is happened.\",\"format\":\"date-time\",\"type\":\"string\"},\"wants_twitter_update\":{\"description\":\"Controls whether to create twitter update.\",\"type\":\"boolean\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update a previous incident update\",\"properties\":{\"affected_components\":{\"description\":\"Affected components associated with the incident update.\",\"items\":{\"example\":{\"code\":\"string\",\"name\":\"string\",\"new_status\":\"operational\",\"old_status\":\"operational\"},\"type\":\"object\"},\"type\":\"array\"},\"body\":{\"description\":\"Incident update body.\",\"type\":\"string\"},\"created_at\":{\"description\":\"The timestamp when the incident update was created at.\",\"format\":\"date-time\",\"type\":\"string\"},\"custom_tweet\":{\"description\":\"An optional customized tweet message for incident postmortem.\",\"type\":\"string\"},\"deliver_notifications\":{\"description\":\"Controls whether to delivery notifications.\",\"type\":\"boolean\"},\"display_at\":{\"description\":\"Timestamp when incident update is happened.\",\"format\":\"date-time\",\"type\":\"string\"},\"id\":{\"description\":\"Incident Update Identifier.\",\"type\":\"string\"},\"incident_id\":{\"description\":\"Incident Identifier.\",\"type\":\"string\"},\"status\":{\"description\":\"The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed.\",\"enum\":[\"investigating\",\"identified\",\"monitoring\",\"resolved\",\"scheduled\",\"in_progress\",\"verifying\",\"completed\"],\"type\":\"string\"},\"tweet_id\":{\"description\":\"Tweet identifier associated to this incident update.\",\"type\":\"string\"},\"twitter_updated_at\":{\"description\":\"The timestamp when twitter updated at.\",\"format\":\"date-time\",\"type\":\"string\"},\"updated_at\":{\"description\":\"The timestamp when the incident update is updated.\",\"format\":\"date-time\",\"type\":\"string\"},\"wants_twitter_update\":{\"description\":\"Controls whether to create twitter update.\",\"type\":\"boolean\"}},\"type\":\"object\"}}},\"description\":\"Update a previous incident update\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "PATCH", "orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}", "rename": { "param": { "incident_update_id": "id" } }, "segments": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incidents" }, { "var": "incident_id" }, { "lit": "incident_updates" }, { "var": "id" }], "select": { "exist": ["id", "incident_id", "page_id"] }, "transform": { "req": { "incident_update": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "patch" }, "update": { "input": "data", "name": "update", "points": [{ "active": true, "args": { "params": [{ "active": true, "kind": "param", "name": "id", "orig": "incident_update_id", "reqd": true, "type": "`$STRING`", "index$": 0 }, { "active": true, "kind": "param", "name": "incident_id", "orig": "incident_id", "reqd": true, "type": "`$STRING`", "index$": 1 }, { "active": true, "kind": "param", "name": "page_id", "orig": "page_id", "reqd": true, "type": "`$STRING`", "index$": 2 }] }, "contract": { "id": "PUT /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}", "json": "{\"operationId\":\"putPagesPageIdIncidentsIncidentIdIncidentUpdatesIncidentUpdateId\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Identifier\",\"in\":\"path\",\"name\":\"incident_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Incident Update Identifier\",\"in\":\"path\",\"name\":\"incident_update_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update a previous incident update\",\"properties\":{\"incident_update\":{\"properties\":{\"body\":{\"description\":\"Incident update body.\",\"type\":\"string\"},\"deliver_notifications\":{\"description\":\"Controls whether to delivery notifications.\",\"type\":\"boolean\"},\"display_at\":{\"description\":\"Timestamp when incident update is happened.\",\"format\":\"date-time\",\"type\":\"string\"},\"wants_twitter_update\":{\"description\":\"Controls whether to create twitter update.\",\"type\":\"boolean\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update a previous incident update\",\"properties\":{\"affected_components\":{\"description\":\"Affected components associated with the incident update.\",\"items\":{\"example\":{\"code\":\"string\",\"name\":\"string\",\"new_status\":\"operational\",\"old_status\":\"operational\"},\"type\":\"object\"},\"type\":\"array\"},\"body\":{\"description\":\"Incident update body.\",\"type\":\"string\"},\"created_at\":{\"description\":\"The timestamp when the incident update was created at.\",\"format\":\"date-time\",\"type\":\"string\"},\"custom_tweet\":{\"description\":\"An optional customized tweet message for incident postmortem.\",\"type\":\"string\"},\"deliver_notifications\":{\"description\":\"Controls whether to delivery notifications.\",\"type\":\"boolean\"},\"display_at\":{\"description\":\"Timestamp when incident update is happened.\",\"format\":\"date-time\",\"type\":\"string\"},\"id\":{\"description\":\"Incident Update Identifier.\",\"type\":\"string\"},\"incident_id\":{\"description\":\"Incident Identifier.\",\"type\":\"string\"},\"status\":{\"description\":\"The incident status. For realtime incidents, valid values are investigating, identified, monitoring, and resolved. For scheduled incidents, valid values are scheduled, in_progress, verifying, and completed.\",\"enum\":[\"investigating\",\"identified\",\"monitoring\",\"resolved\",\"scheduled\",\"in_progress\",\"verifying\",\"completed\"],\"type\":\"string\"},\"tweet_id\":{\"description\":\"Tweet identifier associated to this incident update.\",\"type\":\"string\"},\"twitter_updated_at\":{\"description\":\"The timestamp when twitter updated at.\",\"format\":\"date-time\",\"type\":\"string\"},\"updated_at\":{\"description\":\"The timestamp when the incident update is updated.\",\"format\":\"date-time\",\"type\":\"string\"},\"wants_twitter_update\":{\"description\":\"Controls whether to create twitter update.\",\"type\":\"boolean\"}},\"type\":\"object\"}}},\"description\":\"Update a previous incident update\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}", "source": "openapi3", "version": 1 }, "kind": "http", "method": "PUT", "orig": "/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}", "rename": { "param": { "incident_update_id": "id" } }, "segments": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "incidents" }, { "var": "incident_id" }, { "lit": "incident_updates" }, { "var": "id" }], "select": { "exist": ["id", "incident_id", "page_id"] }, "transform": { "req": { "incident_update": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["page", "incident"]] }, "key$": "incident_update", "name__orig": "incident_update", "Name": "IncidentUpdate", "name_": "incident_update", "name-": "incident-update", "NAME": "INCIDENT_UPDATE", "index$": 7 }, { "active": true, "entity": "incident_update", "key$": "BasicIncidentUpdateFlow", "kind": "basic", "name": "BasicIncidentUpdateFlow", "param": {}, "step": [{ "active": true, "data": { "incident_id": "incident01", "page_id": "page01" }, "input": { "ref": "incident_update_ref01", "srcdatavar": "incident_update_ref01_data", "suffix": "_up0", "textfield": "body" }, "match": {}, "op": "update", "spec": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-incident_update_ref01" } }], "valid": [], "index$": 0 }] }, 'IncidentUpdate');
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let incident_update_ref01_data = Object.values(setup.data.existing.incident_update)[0];
        // UPDATE
        const incident_update_ref01_ent = client.IncidentUpdate();
        const incident_update_ref01_data_up0 = {};
        incident_update_ref01_data_up0.id = incident_update_ref01_data.id;
        incident_update_ref01_data_up0['incident_id'] = setup.idmap['incident_id'];
        incident_update_ref01_data_up0['page_id'] = setup.idmap['page_id'];
        const incident_update_ref01_markdef_up0 = { name: 'body', value: 'Mark01-incident_update_ref01_' + setup.now };
        incident_update_ref01_data_up0[incident_update_ref01_markdef_up0.name] = incident_update_ref01_markdef_up0.value;
        const incident_update_ref01_resdata_up0 = (await incident_update_ref01_ent.update(incident_update_ref01_data_up0)).data();
        (0, node_assert_1.default)(incident_update_ref01_resdata_up0.id === incident_update_ref01_data_up0.id);
        (0, node_assert_1.default)(incident_update_ref01_resdata_up0[incident_update_ref01_markdef_up0.name] === incident_update_ref01_markdef_up0.value);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/incident_update/IncidentUpdateTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['incident_update01', 'incident_update02', 'incident_update03', 'page01', 'page02', 'page03', 'incident01', 'incident02', 'incident03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID'];
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
//# sourceMappingURL=IncidentUpdateEntity.test.js.map