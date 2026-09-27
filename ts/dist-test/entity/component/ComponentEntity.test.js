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
(0, node_test_1.describe)('ComponentEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.Component();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['create', 'list', 'update', 'load', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'component.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "automation_email": { "a": true, "h": "Automation Email", "n": "automation_email", "r": false, "sh": "Requires a special feature flag to be enabled", "t": "`$STRING`", "key$": "automation_email", "index$": 0 }, "component": { "a": true, "h": "Component", "n": "component", "r": false, "t": "`$OBJECT`", "key$": "component", "index$": 1 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": false, "t": "`$STRING`", "key$": "created_at", "index$": 2 }, "description": { "a": true, "h": "Description", "n": "description", "r": false, "sh": "More detailed description for component", "t": "`$STRING`", "key$": "description", "index$": 3 }, "group": { "a": true, "h": "Group", "n": "group", "r": false, "sh": "Is this component a group", "t": "`$BOOLEAN`", "key$": "group", "index$": 4 }, "group_id": { "a": true, "h": "Group Id", "n": "group_id", "r": false, "sh": "Component Group identifier", "t": "`$STRING`", "key$": "group_id", "index$": 5 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Identifier for component", "t": "`$STRING`", "key$": "id", "index$": 6 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "sh": "Display name for component", "t": "`$STRING`", "key$": "name", "index$": 7 }, "only_show_if_degraded": { "a": true, "h": "Only Show If Degraded", "n": "only_show_if_degraded", "r": false, "sh": "Requires a special feature flag to be enabled", "t": "`$BOOLEAN`", "key$": "only_show_if_degraded", "index$": 8 }, "page_id": { "a": true, "h": "Page Id", "n": "page_id", "r": false, "sh": "Page identifier", "t": "`$STRING`", "key$": "page_id", "index$": 9 }, "position": { "a": true, "fo": "int32", "h": "Position", "n": "position", "r": false, "sh": "Order the component will appear on the page", "t": "`$INTEGER`", "key$": "position", "index$": 10 }, "showcase": { "a": true, "h": "Showcase", "n": "showcase", "r": false, "sh": "Should this component be showcased", "t": "`$BOOLEAN`", "key$": "showcase", "index$": 11 }, "start_date": { "a": true, "fo": "date", "h": "Start Date", "n": "start_date", "r": false, "sh": "The date this component started being used", "t": "`$STRING`", "key$": "start_date", "index$": 12 }, "status": { "a": true, "h": "Status", "n": "status", "r": false, "sh": "Status of component", "t": "`$STRING`", "key$": "status", "index$": 13 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": false, "t": "`$STRING`", "key$": "updated_at", "index$": 14 } }, "id": { "field": "id", "name": "id" }, "name": "component", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /pages/{page_id}/components/{component_id}/page_access_groups", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/components/{component_id}/page_access_groups", "q": { "$action": "page_access_group", "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }, { "lit": "page_access_groups" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "POST /pages/{page_id}/components/{component_id}/page_access_users", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/components/{component_id}/page_access_users", "q": { "$action": "page_access_user", "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }, { "lit": "page_access_users" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "POST /pages/{page_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/components", "q": { "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }], "t": { "req": { "component": "`reqdata`" }, "res": "`body`" }, "index$": 2 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/page_access_groups/{page_access_group_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_access_group_id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components", "q": { "exist": ["page", "page_access_group_id", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "page_access_group_id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "GET /pages/{page_id}/page_access_users/{page_access_user_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_access_user_id", "or": "page_access_user_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/page_access_users/{page_access_user_id}/components", "q": { "exist": ["page", "page_access_user_id", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_users" }, { "var": "page_access_user_id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "GET /pages/{page_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/components", "q": { "exist": ["page", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/components/{component_id}/uptime", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }], "query": [{ "a": true, "k": "query", "n": "end", "or": "end", "r": false, "t": "Any", "index$": 0 }, { "a": true, "k": "query", "n": "start", "or": "start", "r": false, "t": "Any", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/components/{component_id}/uptime", "q": { "$action": "uptime", "exist": ["end", "id", "page_id", "start"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }, { "lit": "uptime" }], "t": { "req": "`reqdata`", "res": "`body.related_events`" }, "index$": 0 }, { "a": true, "co": { "id": "GET /pages/{page_id}/components/{component_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/components/{component_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /pages/{page_id}/components/{component_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}/components/{component_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }], "t": { "req": { "component": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "patch" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /pages/{page_id}/components/{component_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/components/{component_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "DELETE /pages/{page_id}/components/{component_id}/page_access_groups", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/components/{component_id}/page_access_groups", "q": { "$action": "page_access_group", "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }, { "lit": "page_access_groups" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "DELETE /pages/{page_id}/components/{component_id}/page_access_users", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/components/{component_id}/page_access_users", "q": { "$action": "page_access_user", "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }, { "lit": "page_access_users" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /pages/{page_id}/components/{component_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}/components/{component_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "component_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "components" }, { "var": "id" }], "t": { "req": { "component": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.page"], ["$.main.kit.entity.page", "$.main.kit.entity.page_access_group"], ["$.main.kit.entity.page", "$.main.kit.entity.page_access_user"]] }, "key$": "component", "name__orig": "component", "Name": "Component", "name_": "component", "name-": "component", "NAME": "COMPONENT", "index$": 0 }, { "active": true, "entity": "component", "key$": "BasicComponentFlow", "kind": "basic", "name": "BasicComponentFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "component_ref01" }, "m": { "page_access_group_id": "page_access_group01", "page_access_user_id": "page_access_user01", "page_id": "page01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "component_ref01" } }], "index$": 1 }, { "a": true, "d": { "page_id": "page01" }, "i": { "ref": "component_ref01", "srcdatavar": "component_ref01_data", "suffix": "_up0", "textfield": "automation_email" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-component_ref01" } }], "v": [], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "component_ref01", "srcdatavar": "component_ref01_data", "suffix": "_dt0" }, "m": { "id": "component01", "page_id": "page01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-component_ref01" } }], "index$": 3 }, { "a": true, "d": {}, "i": { "ref": "component_ref01", "suffix": "_rm0" }, "m": { "id": "component01", "page_id": "page01" }, "o": "remove", "s": [], "v": [], "index$": 4 }, { "a": true, "d": {}, "i": { "suffix": "_rt0" }, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemNotExists", "def": { "ref": "component_ref01" } }], "index$": 5 }] }, 'Component', { "POST /pages/{page_id}/components/{component_id}/page_access_groups": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "POST /pages/{page_id}/components/{component_id}/page_access_users": { "protocol": "http", "requestBody": { "content": { "application/x-www-form-urlencoded": { "schema": { "type": "object", "properties": { "page_access_user_ids": { "description": "List of page access users to add to component", "type": "array", "items": { "type": "string" } } }, "required": ["page_access_user_ids"] } } } }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "POST /pages/{page_id}/components": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component": { "type": "object", "properties": { "description": { "type": "string", "description": "More detailed description for component" }, "status": { "type": "string", "description": "Status of component", "enum": ["operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage", ""] }, "name": { "type": "string", "description": "Display name for component" }, "only_show_if_degraded": { "type": "boolean", "description": "Requires a special feature flag to be enabled" }, "group_id": { "type": "string", "description": "Component Group identifier" }, "showcase": { "type": "boolean", "description": "Should this component be showcased" }, "start_date": { "type": "string", "format": "date", "description": "The date this component started being used" } }, "key$": "component" } }, "description": "Create a component", "x-ref": "#/components/schemas/postPagesPageIdComponents", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "GET /pages/{page_id}/page_access_groups/{page_access_group_id}/components": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 3 }] }, "GET /pages/{page_id}/page_access_users/{page_access_user_id}/components": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_user_id", "description": "Page Access User Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 3 }] }, "GET /pages/{page_id}/components": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 1 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32", "minimum": 1, "maximum": 1100 }, "index$": 2 }] }, "GET /pages/{page_id}/components/{component_id}/uptime": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }, { "in": "query", "name": "start", "description": "The start date for uptime calculation (defaults to the component's start_date field or 90 days ago, whichever is more recent).\nThe maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year.\nIf the year and month are given, the start date defaults to the first day of that month.\nThe earliest supported date is January 1, 1970.\n", "required": false, "schema": { "type": "PartialStartDate" }, "index$": 2 }, { "in": "query", "name": "end", "description": "The end date for uptime calculation (defaults to today in the page's time zone). The maximum supported date range is six calendar months.\nIf the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month.\nThe earliest supported date is January 1, 1970.\n", "required": false, "schema": { "type": "PartialEndDate" }, "index$": 3 }] }, "GET /pages/{page_id}/components/{component_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PATCH /pages/{page_id}/components/{component_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component": { "type": "object", "properties": { "description": { "type": "string", "description": "More detailed description for component" }, "status": { "type": "string", "description": "Status of component", "enum": ["operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage", ""] }, "name": { "type": "string", "description": "Display name for component" }, "only_show_if_degraded": { "type": "boolean", "description": "Requires a special feature flag to be enabled" }, "group_id": { "type": "string", "description": "Component Group identifier" }, "showcase": { "type": "boolean", "description": "Should this component be showcased" }, "start_date": { "type": "string", "format": "date", "description": "The date this component started being used" } }, "key$": "component" } }, "description": "Update a component", "x-ref": "#/components/schemas/patchPagesPageIdComponents", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/components/{component_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/components/{component_id}/page_access_groups": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/components/{component_id}/page_access_users": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PUT /pages/{page_id}/components/{component_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component": { "type": "object", "properties": { "description": { "type": "string", "description": "More detailed description for component" }, "status": { "type": "string", "description": "Status of component", "enum": ["operational", "under_maintenance", "degraded_performance", "partial_outage", "major_outage", ""] }, "name": { "type": "string", "description": "Display name for component" }, "only_show_if_degraded": { "type": "boolean", "description": "Requires a special feature flag to be enabled" }, "group_id": { "type": "string", "description": "Component Group identifier" }, "showcase": { "type": "boolean", "description": "Should this component be showcased" }, "start_date": { "type": "string", "format": "date", "description": "The date this component started being used" } }, "key$": "component" } }, "description": "Update a component", "x-ref": "#/components/schemas/putPagesPageIdComponents", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const component_ref01_ent = client.Component();
        let component_ref01_data = setup.data.new.component['component_ref01'];
        component_ref01_data['page_access_group_id'] = setup.idmap['page_access_group01'];
        component_ref01_data['page_access_user_id'] = setup.idmap['page_access_user01'];
        component_ref01_data['page_id'] = setup.idmap['page01'];
        component_ref01_data = (await component_ref01_ent.create(component_ref01_data)).data();
        (0, node_assert_1.default)(null != component_ref01_data.id);
        // LIST
        const component_ref01_match = {};
        component_ref01_match['page_id'] = setup.idmap['page01'];
        const component_ref01_list = (await component_ref01_ent.list(component_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(component_ref01_list, { id: component_ref01_data.id })));
        // UPDATE
        const component_ref01_data_up0 = {};
        component_ref01_data_up0.id = component_ref01_data.id;
        component_ref01_data_up0['page_id'] = setup.idmap['page_id'];
        const component_ref01_markdef_up0 = { name: 'automation_email', value: 'Mark01-component_ref01_' + setup.now };
        component_ref01_data_up0[component_ref01_markdef_up0.name] = component_ref01_markdef_up0.value;
        const component_ref01_resdata_up0 = (await component_ref01_ent.update(component_ref01_data_up0)).data();
        (0, node_assert_1.default)(component_ref01_resdata_up0.id === component_ref01_data_up0.id);
        (0, node_assert_1.default)(component_ref01_resdata_up0[component_ref01_markdef_up0.name] === component_ref01_markdef_up0.value);
        // LOAD
        const component_ref01_match_dt0 = {};
        component_ref01_match_dt0.id = component_ref01_data.id;
        const component_ref01_data_dt0 = (await component_ref01_ent.load(component_ref01_match_dt0)).data();
        (0, node_assert_1.default)(component_ref01_data_dt0.id === component_ref01_data.id);
        // REMOVE
        const component_ref01_match_rm0 = { id: component_ref01_data.id };
        await component_ref01_ent.remove(component_ref01_match_rm0);
        // LIST
        const component_ref01_match_rt0 = {};
        component_ref01_match_rt0['page_id'] = setup.idmap['page01'];
        const component_ref01_list_rt0 = (await component_ref01_ent.list(component_ref01_match_rt0)).map((e) => e.data());
        (0, node_assert_1.default)(isempty(select(component_ref01_list_rt0, { id: component_ref01_data.id })));
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/component/ComponentTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['component01', 'component02', 'component03', 'page01', 'page02', 'page03', 'page_access_group01', 'page_access_group02', 'page_access_group03', 'page_access_user01', 'page_access_user02', 'page_access_user03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_COMPONENT_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_COMPONENT_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_COMPONENT_ENTID'];
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
//# sourceMappingURL=ComponentEntity.test.js.map