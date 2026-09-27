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
(0, node_test_1.describe)('PageAccessGroupEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.PageAccessGroup();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['create', 'list', 'update', 'load', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'page_access_group.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "component_ids": { "a": true, "h": "Component Ids", "n": "component_ids", "r": false, "t": "`$ARRAY`", "key$": "component_ids", "index$": 0 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": false, "t": "`$STRING`", "key$": "created_at", "index$": 1 }, "external_identifier": { "a": true, "h": "External Identifier", "n": "external_identifier", "r": false, "sh": "Associates group with external group.", "t": "`$STRING`", "key$": "external_identifier", "index$": 2 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Page Access Group Identifier", "t": "`$STRING`", "key$": "id", "index$": 3 }, "metric_ids": { "a": true, "h": "Metric Ids", "n": "metric_ids", "r": false, "t": "`$ARRAY`", "key$": "metric_ids", "index$": 4 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "sh": "Name for this Group.", "t": "`$STRING`", "key$": "name", "index$": 5 }, "page_access_group": { "a": true, "h": "Page Access Group", "n": "page_access_group", "r": false, "t": "`$OBJECT`", "key$": "page_access_group", "index$": 6 }, "page_access_user_ids": { "a": true, "h": "Page Access User Ids", "n": "page_access_user_ids", "r": false, "t": "`$ARRAY`", "key$": "page_access_user_ids", "index$": 7 }, "page_id": { "a": true, "h": "Page Id", "n": "page_id", "r": false, "sh": "Page Identifier.", "t": "`$STRING`", "key$": "page_id", "index$": 8 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": false, "t": "`$STRING`", "key$": "updated_at", "index$": 9 } }, "id": { "field": "id", "name": "id" }, "name": "page_access_group", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /pages/{page_id}/page_access_groups/{page_access_group_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components", "q": { "$action": "component", "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "POST /pages/{page_id}/page_access_groups", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/page_access_groups", "q": { "exist": ["id"] }, "r": { "param": { "page_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "id" }, { "lit": "page_access_groups" }], "t": { "req": { "page_access_group": "`reqdata`" }, "res": "`body`" }, "index$": 1 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/page_access_groups", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/page_access_groups", "q": { "exist": ["id", "page", "per_page"] }, "r": { "param": { "page_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "id" }, { "lit": "page_access_groups" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/page_access_groups/{page_access_group_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /pages/{page_id}/page_access_groups/{page_access_group_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }], "t": { "req": { "page_access_group": "`reqdata`" }, "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "PATCH /pages/{page_id}/page_access_groups/{page_access_group_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components", "q": { "$action": "component", "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "patch" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "component_id", "or": "component_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 1 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 2 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}", "q": { "exist": ["component_id", "id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }, { "lit": "components" }, { "var": "component_id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components", "q": { "$action": "component", "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /pages/{page_id}/page_access_groups/{page_access_group_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }], "t": { "req": { "page_access_group": "`reqdata`" }, "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "PUT /pages/{page_id}/page_access_groups/{page_access_group_id}/components", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_access_group_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}/page_access_groups/{page_access_group_id}/components", "q": { "$action": "component", "exist": ["id", "page_id"] }, "r": { "param": { "page_access_group_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_groups" }, { "var": "id" }, { "lit": "components" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.page"], ["$.main.kit.entity.page", "$.main.kit.entity.component"]] }, "key$": "page_access_group", "name__orig": "page_access_group", "Name": "PageAccessGroup", "name_": "page_access_group", "name-": "page-access-group", "NAME": "PAGE_ACCESS_GROUP", "index$": 10 }, { "active": true, "entity": "page_access_group", "key$": "BasicPageAccessGroupFlow", "kind": "basic", "name": "BasicPageAccessGroupFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "page_access_group_ref01" }, "m": { "page_id": "page01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "page_access_group_ref01" } }], "index$": 1 }, { "a": true, "d": { "page_id": "page01" }, "i": { "ref": "page_access_group_ref01", "srcdatavar": "page_access_group_ref01_data", "suffix": "_up0", "textfield": "created_at" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-page_access_group_ref01" } }], "v": [], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "page_access_group_ref01", "srcdatavar": "page_access_group_ref01_data", "suffix": "_dt0" }, "m": { "id": "page_access_group01", "page_id": "page01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-page_access_group_ref01" } }], "index$": 3 }, { "a": true, "d": {}, "i": { "ref": "page_access_group_ref01", "suffix": "_rm0" }, "m": { "id": "page_access_group01", "page_id": "page01" }, "o": "remove", "s": [], "v": [], "index$": 4 }, { "a": true, "d": {}, "i": { "suffix": "_rt0" }, "m": { "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemNotExists", "def": { "ref": "page_access_group_ref01" } }], "index$": 5 }] }, 'PageAccessGroup', { "POST /pages/{page_id}/page_access_groups/{page_access_group_id}/components": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component_ids": { "type": "array", "description": "List of components codes to set on the page access group", "items": { "type": "string" } } }, "required": ["component_ids"], "description": "Replace components for a page access group", "x-ref": "#/components/schemas/postPagesPageIdPageAccessGroupsPageAccessGroupIdComponents" } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "POST /pages/{page_id}/page_access_groups": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "page_access_group": { "type": "object", "properties": { "name": { "type": "string", "description": "Name for this Group." }, "external_identifier": { "type": "string", "description": "Associates group with external group." }, "component_ids": { "type": "array", "items": { "type": "string" } }, "metric_ids": { "type": "array", "items": { "type": "string" } }, "page_access_user_ids": { "type": "array", "items": { "type": "string" } } }, "key$": "page_access_group" } }, "description": "Create a page access group", "x-ref": "#/components/schemas/postPagesPageIdPageAccessGroups", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "GET /pages/{page_id}/page_access_groups": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 1 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }] }, "GET /pages/{page_id}/page_access_groups/{page_access_group_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PATCH /pages/{page_id}/page_access_groups/{page_access_group_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "page_access_group": { "type": "object", "properties": { "name": { "type": "string", "description": "Name for this Group." }, "external_identifier": { "type": "string", "description": "Associates group with external group." }, "component_ids": { "type": "array", "items": { "type": "string" } }, "metric_ids": { "type": "array", "items": { "type": "string" } }, "page_access_user_ids": { "type": "array", "items": { "type": "string" } } }, "key$": "page_access_group" } }, "description": "Update a page access group", "x-ref": "#/components/schemas/patchPagesPageIdPageAccessGroups", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PATCH /pages/{page_id}/page_access_groups/{page_access_group_id}/components": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component_ids": { "type": "array", "description": "List of Component identifiers", "items": { "type": "string" } } }, "description": "Add components to page access group", "x-ref": "#/components/schemas/patchPagesPageIdPageAccessGroupsPageAccessGroupIdComponents" } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}/components/{component_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }, { "in": "path", "name": "component_id", "description": "Component identifier", "required": true, "schema": { "type": "string" }, "index$": 2 }] }, "DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/page_access_groups/{page_access_group_id}/components": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component_ids": { "type": "array", "items": { "type": "string" } } }, "description": "Delete components for a page access group", "x-ref": "#/components/schemas/deletePagesPageIdPageAccessGroupsPageAccessGroupIdComponents" } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PUT /pages/{page_id}/page_access_groups/{page_access_group_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "page_access_group": { "type": "object", "properties": { "name": { "type": "string", "description": "Name for this Group." }, "external_identifier": { "type": "string", "description": "Associates group with external group." }, "component_ids": { "type": "array", "items": { "type": "string" } }, "metric_ids": { "type": "array", "items": { "type": "string" } }, "page_access_user_ids": { "type": "array", "items": { "type": "string" } } }, "key$": "page_access_group" } }, "description": "Update a page access group", "x-ref": "#/components/schemas/putPagesPageIdPageAccessGroups", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PUT /pages/{page_id}/page_access_groups/{page_access_group_id}/components": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "component_ids": { "type": "array", "description": "List of Component identifiers", "items": { "type": "string" } } }, "description": "Add components to page access group", "x-ref": "#/components/schemas/putPagesPageIdPageAccessGroupsPageAccessGroupIdComponents" } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_group_id", "description": "Page Access Group Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const page_access_group_ref01_ent = client.PageAccessGroup();
        let page_access_group_ref01_data = setup.data.new.page_access_group['page_access_group_ref01'];
        page_access_group_ref01_data['page_id'] = setup.idmap['page01'];
        page_access_group_ref01_data = (await page_access_group_ref01_ent.create(page_access_group_ref01_data)).data();
        (0, node_assert_1.default)(null != page_access_group_ref01_data.id);
        // LIST
        const page_access_group_ref01_match = {};
        page_access_group_ref01_match['page_id'] = setup.idmap['page01'];
        const page_access_group_ref01_list = (await page_access_group_ref01_ent.list(page_access_group_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(page_access_group_ref01_list, { id: page_access_group_ref01_data.id })));
        // UPDATE
        const page_access_group_ref01_data_up0 = {};
        page_access_group_ref01_data_up0.id = page_access_group_ref01_data.id;
        page_access_group_ref01_data_up0['page_id'] = setup.idmap['page_id'];
        const page_access_group_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-page_access_group_ref01_' + setup.now };
        page_access_group_ref01_data_up0[page_access_group_ref01_markdef_up0.name] = page_access_group_ref01_markdef_up0.value;
        const page_access_group_ref01_resdata_up0 = (await page_access_group_ref01_ent.update(page_access_group_ref01_data_up0)).data();
        (0, node_assert_1.default)(page_access_group_ref01_resdata_up0.id === page_access_group_ref01_data_up0.id);
        (0, node_assert_1.default)(page_access_group_ref01_resdata_up0[page_access_group_ref01_markdef_up0.name] === page_access_group_ref01_markdef_up0.value);
        // LOAD
        const page_access_group_ref01_match_dt0 = {};
        page_access_group_ref01_match_dt0.id = page_access_group_ref01_data.id;
        const page_access_group_ref01_data_dt0 = (await page_access_group_ref01_ent.load(page_access_group_ref01_match_dt0)).data();
        (0, node_assert_1.default)(page_access_group_ref01_data_dt0.id === page_access_group_ref01_data.id);
        // REMOVE
        const page_access_group_ref01_match_rm0 = { id: page_access_group_ref01_data.id };
        await page_access_group_ref01_ent.remove(page_access_group_ref01_match_rm0);
        // LIST
        const page_access_group_ref01_match_rt0 = {};
        page_access_group_ref01_match_rt0['page_id'] = setup.idmap['page01'];
        const page_access_group_ref01_list_rt0 = (await page_access_group_ref01_ent.list(page_access_group_ref01_match_rt0)).map((e) => e.data());
        (0, node_assert_1.default)(isempty(select(page_access_group_ref01_list_rt0, { id: page_access_group_ref01_data.id })));
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/page_access_group/PageAccessGroupTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['page_access_group01', 'page_access_group02', 'page_access_group03', 'page01', 'page02', 'page03', 'component01', 'component02', 'component03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID'];
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
//# sourceMappingURL=PageAccessGroupEntity.test.js.map