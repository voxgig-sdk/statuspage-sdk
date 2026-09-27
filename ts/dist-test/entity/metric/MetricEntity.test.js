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
(0, node_test_1.describe)('MetricEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.Metric();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['create', 'list', 'update', 'load', 'remove']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'metric.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "backfill_percentage": { "a": true, "fo": "int32", "h": "Backfill Percentage", "n": "backfill_percentage", "r": false, "t": "`$INTEGER`", "key$": "backfill_percentage", "index$": 0 }, "backfilled": { "a": true, "h": "Backfilled", "n": "backfilled", "r": false, "t": "`$BOOLEAN`", "key$": "backfilled", "index$": 1 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": false, "t": "`$STRING`", "key$": "created_at", "index$": 2 }, "decimal_places": { "a": true, "fo": "int32", "h": "Decimal Places", "n": "decimal_places", "r": false, "t": "`$INTEGER`", "key$": "decimal_places", "index$": 3 }, "display": { "a": true, "h": "Display", "n": "display", "r": false, "sh": "Should the metric be displayed", "t": "`$BOOLEAN`", "key$": "display", "index$": 4 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Metric identifier", "t": "`$STRING`", "key$": "id", "index$": 5 }, "last_fetched_at": { "a": true, "fo": "date-time", "h": "Last Fetched At", "n": "last_fetched_at", "r": false, "t": "`$STRING`", "key$": "last_fetched_at", "index$": 6 }, "metric": { "a": true, "h": "Metric", "n": "metric", "r": false, "t": "`$OBJECT`", "key$": "metric", "index$": 7 }, "metric_identifier": { "a": true, "h": "Metric Identifier", "n": "metric_identifier", "r": false, "sh": "Metric Display identifier used to look up the metric data from the provider", "t": "`$STRING`", "key$": "metric_identifier", "index$": 8 }, "metrics_provider_id": { "a": true, "h": "Metrics Provider Id", "n": "metrics_provider_id", "r": false, "sh": "Metric Provider identifier", "t": "`$STRING`", "key$": "metrics_provider_id", "index$": 9 }, "most_recent_data_at": { "a": true, "fo": "date-time", "h": "Most Recent Data At", "n": "most_recent_data_at", "r": false, "t": "`$STRING`", "key$": "most_recent_data_at", "index$": 10 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "sh": "Name of metric", "t": "`$STRING`", "key$": "name", "index$": 11 }, "reference_name": { "a": true, "h": "Reference Name", "n": "reference_name", "r": false, "t": "`$STRING`", "key$": "reference_name", "index$": 12 }, "suffix": { "a": true, "h": "Suffix", "n": "suffix", "r": false, "sh": "Suffix to describe the units on the graph", "t": "`$STRING`", "key$": "suffix", "index$": 13 }, "tooltip_description": { "a": true, "h": "Tooltip Description", "n": "tooltip_description", "r": false, "t": "`$STRING`", "key$": "tooltip_description", "index$": 14 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": false, "t": "`$STRING`", "key$": "updated_at", "index$": 15 }, "y_axis_hidden": { "a": true, "h": "Y Axis Hidden", "n": "y_axis_hidden", "r": false, "sh": "Should the values on the y axis be hidden on render", "t": "`$BOOLEAN`", "key$": "y_axis_hidden", "index$": 16 }, "y_axis_max": { "a": true, "fo": "float", "h": "Y Axis Max", "n": "y_axis_max", "r": false, "t": "`$NUMBER`", "key$": "y_axis_max", "index$": 17 }, "y_axis_min": { "a": true, "fo": "float", "h": "Y Axis Min", "n": "y_axis_min", "r": false, "t": "`$NUMBER`", "key$": "y_axis_min", "index$": 18 } }, "id": { "field": "id", "name": "id" }, "name": "metric", "op": { "create": { "input": "data", "name": "create", "points": [{ "a": true, "co": { "id": "POST /pages/{page_id}/metrics/{metric_id}/data", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "metric_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/metrics/{metric_id}/data", "q": { "$action": "data", "exist": ["id", "page_id"] }, "r": { "param": { "metric_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "var": "id" }, { "lit": "data" }], "t": { "req": "`reqdata`", "res": "`body.data`" }, "index$": 0 }, { "a": true, "co": { "id": "POST /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "metrics_provider_id", "or": "metrics_provider_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics", "q": { "exist": ["metrics_provider_id", "page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics_providers" }, { "var": "metrics_provider_id" }, { "lit": "metrics" }], "t": { "req": { "metric": "`reqdata`" }, "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "POST /pages/{page_id}/metrics/data", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "POST", "o": "/pages/{page_id}/metrics/data", "q": { "$action": "data", "exist": ["page_id"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "lit": "data" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "create" }, "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/page_access_users/{page_access_user_id}/metrics", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_access_user_id", "or": "page_access_user_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/page_access_users/{page_access_user_id}/metrics", "q": { "exist": ["page", "page_access_user_id", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "page_access_users" }, { "var": "page_access_user_id" }, { "lit": "metrics" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "metrics_provider_id", "or": "metrics_provider_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics", "q": { "exist": ["metrics_provider_id", "page", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics_providers" }, { "var": "metrics_provider_id" }, { "lit": "metrics" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "GET /pages/{page_id}/metrics", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }], "query": [{ "a": true, "k": "query", "n": "page", "or": "page", "r": false, "t": "`$INTEGER`", "index$": 0 }, { "a": true, "k": "query", "n": "per_page", "or": "per_page", "r": false, "t": "`$INTEGER`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/metrics", "q": { "exist": ["page", "page_id", "per_page"] }, "r": {}, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }, { "a": true, "co": { "id": "GET /pages/{page_id}/metrics/{metric_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "metric_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}/metrics/{metric_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "metric_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 2 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /pages/{page_id}/metrics/{metric_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "metric_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}/metrics/{metric_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "metric_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "var": "id" }], "t": { "req": { "metric": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "patch" }, "remove": { "input": "data", "name": "remove", "points": [{ "a": true, "co": { "id": "DELETE /pages/{page_id}/metrics/{metric_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "metric_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/metrics/{metric_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "metric_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }, { "a": true, "co": { "id": "DELETE /pages/{page_id}/metrics/{metric_id}/data", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "metric_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "DELETE", "o": "/pages/{page_id}/metrics/{metric_id}/data", "q": { "$action": "data", "exist": ["id", "page_id"] }, "r": { "param": { "metric_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "var": "id" }, { "lit": "data" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 1 }], "key$": "remove" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /pages/{page_id}/metrics/{metric_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "metric_id", "r": true, "t": "`$STRING`", "index$": 0 }, { "a": true, "k": "param", "n": "page_id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 1 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}/metrics/{metric_id}", "q": { "exist": ["id", "page_id"] }, "r": { "param": { "metric_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "page_id" }, { "lit": "metrics" }, { "var": "id" }], "t": { "req": { "metric": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [["$.main.kit.entity.page"], ["$.main.kit.entity.page", "$.main.kit.entity.metrics_provider"], ["$.main.kit.entity.page", "$.main.kit.entity.page_access_user"]] }, "key$": "metric", "name__orig": "metric", "Name": "Metric", "name_": "metric", "name-": "metric", "NAME": "METRIC", "index$": 7 }, { "active": true, "entity": "metric", "key$": "BasicMetricFlow", "kind": "basic", "name": "BasicMetricFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": { "ref": "metric_ref01" }, "m": { "page_access_user_id": "page_access_user01", "page_id": "page01" }, "o": "create", "s": [], "v": [], "index$": 0 }, { "a": true, "d": {}, "i": {}, "m": { "page_access_user_id": "page_access_user01", "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "metric_ref01" } }], "index$": 1 }, { "a": true, "d": { "page_id": "page01" }, "i": { "ref": "metric_ref01", "srcdatavar": "metric_ref01_data", "suffix": "_up0", "textfield": "created_at" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-metric_ref01" } }], "v": [], "index$": 2 }, { "a": true, "d": {}, "i": { "ref": "metric_ref01", "srcdatavar": "metric_ref01_data", "suffix": "_dt0" }, "m": { "id": "metric01", "page_id": "page01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-metric_ref01" } }], "index$": 3 }, { "a": true, "d": {}, "i": { "ref": "metric_ref01", "suffix": "_rm0" }, "m": { "id": "metric01", "page_id": "page01" }, "o": "remove", "s": [], "v": [], "index$": 4 }, { "a": true, "d": {}, "i": { "suffix": "_rt0" }, "m": { "page_access_user_id": "page_access_user01", "page_id": "page01" }, "o": "list", "s": [], "v": [{ "apply": "ItemNotExists", "def": { "ref": "metric_ref01" } }], "index$": 5 }] }, 'Metric', { "POST /pages/{page_id}/metrics/{metric_id}/data": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "data": { "type": "object", "properties": { "timestamp": { "type": "integer", "format": "int32", "description": "Time to store the metric against" }, "value": { "type": "number", "format": "float" } } } }, "required": ["data"], "description": "Add data to a metric", "x-ref": "#/components/schemas/postPagesPageIdMetricsMetricIdData" } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metric_id", "description": "Metric Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "POST /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "metric": { "type": "object", "properties": { "name": { "type": "string", "description": "Name of metric" }, "metric_identifier": { "type": "string", "description": "The identifier used to look up the metric data from the provider" }, "transform": { "type": "string", "description": "The transform to apply to metric before pulling into Statuspage. One of: \"average\", \"count\", \"max\", \"min\", or \"sum\"" }, "application_id": { "type": "string", "description": "The Identifier for new relic application. Required in the case of NewRelic only" }, "suffix": { "type": "string", "description": "Suffix to describe the units on the graph" }, "y_axis_min": { "type": "integer", "format": "int32", "description": "The lower bound of the y axis" }, "y_axis_max": { "type": "integer", "format": "int32", "description": "The upper bound of the y axis" }, "y_axis_hidden": { "type": "boolean", "description": "Should the values on the y axis be hidden on render" }, "display": { "type": "boolean", "description": "Should the metric be displayed" }, "decimal_places": { "type": "integer", "format": "int32", "description": "How many decimal places to render on the graph" }, "tooltip_description": { "type": "string" } }, "key$": "metric" } }, "description": "Create a metric for a metric provider", "x-ref": "#/components/schemas/postPagesPageIdMetricsProvidersMetricsProviderIdMetrics", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metrics_provider_id", "description": "Metric Provider Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "POST /pages/{page_id}/metrics/data": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "data": { "type": "object", "properties": { "metric_id": { "type": "array", "items": { "type": "object", "properties": {} }, "description": "Metric identifier to add data to" } }, "description": "Add data points to metrics", "x-ref": "#/components/schemas/MetricAddResponse" } }, "required": ["data"], "description": "Add data points to metrics", "x-ref": "#/components/schemas/postPagesPageIdMetricsData" } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "GET /pages/{page_id}/page_access_users/{page_access_user_id}/metrics": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "page_access_user_id", "description": "Page Access User Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 3 }] }, "GET /pages/{page_id}/metrics_providers/{metrics_provider_id}/metrics": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metrics_provider_id", "description": "Metric Provider Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 3 }] }, "GET /pages/{page_id}/metrics": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "query", "name": "page", "description": "Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 1 }, { "in": "query", "name": "per_page", "description": "Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.", "required": false, "schema": { "type": "integer", "format": "int32" }, "index$": 2 }] }, "GET /pages/{page_id}/metrics/{metric_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metric_id", "description": "Metric Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PATCH /pages/{page_id}/metrics/{metric_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "metric": { "type": "object", "properties": { "name": { "type": "string", "description": "Name of metric" }, "metric_identifier": { "type": "string", "description": "Metric Display identifier used to look up the metric data from the provider" } }, "key$": "metric" } }, "description": "Update a metric", "x-ref": "#/components/schemas/patchPagesPageIdMetrics", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metric_id", "description": "Metric Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/metrics/{metric_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metric_id", "description": "Metric Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "DELETE /pages/{page_id}/metrics/{metric_id}/data": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metric_id", "description": "Metric Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] }, "PUT /pages/{page_id}/metrics/{metric_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "metric": { "type": "object", "properties": { "name": { "type": "string", "description": "Name of metric" }, "metric_identifier": { "type": "string", "description": "Metric Display identifier used to look up the metric data from the provider" } }, "key$": "metric" } }, "description": "Update a metric", "x-ref": "#/components/schemas/putPagesPageIdMetrics", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }, { "in": "path", "name": "metric_id", "description": "Metric Identifier", "required": true, "schema": { "type": "string" }, "index$": 1 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        // CREATE
        const metric_ref01_ent = client.Metric();
        let metric_ref01_data = setup.data.new.metric['metric_ref01'];
        metric_ref01_data['page_access_user_id'] = setup.idmap['page_access_user01'];
        metric_ref01_data['page_id'] = setup.idmap['page01'];
        metric_ref01_data = (await metric_ref01_ent.create(metric_ref01_data)).data();
        (0, node_assert_1.default)(null != metric_ref01_data.id);
        // LIST
        const metric_ref01_match = {};
        metric_ref01_match['page_access_user_id'] = setup.idmap['page_access_user01'];
        metric_ref01_match['page_id'] = setup.idmap['page01'];
        const metric_ref01_list = (await metric_ref01_ent.list(metric_ref01_match)).map((e) => e.data());
        (0, node_assert_1.default)(!isempty(select(metric_ref01_list, { id: metric_ref01_data.id })));
        // UPDATE
        const metric_ref01_data_up0 = {};
        metric_ref01_data_up0.id = metric_ref01_data.id;
        metric_ref01_data_up0['page_id'] = setup.idmap['page_id'];
        const metric_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-metric_ref01_' + setup.now };
        metric_ref01_data_up0[metric_ref01_markdef_up0.name] = metric_ref01_markdef_up0.value;
        const metric_ref01_resdata_up0 = (await metric_ref01_ent.update(metric_ref01_data_up0)).data();
        (0, node_assert_1.default)(metric_ref01_resdata_up0.id === metric_ref01_data_up0.id);
        (0, node_assert_1.default)(metric_ref01_resdata_up0[metric_ref01_markdef_up0.name] === metric_ref01_markdef_up0.value);
        // LOAD
        const metric_ref01_match_dt0 = {};
        metric_ref01_match_dt0.id = metric_ref01_data.id;
        const metric_ref01_data_dt0 = (await metric_ref01_ent.load(metric_ref01_match_dt0)).data();
        (0, node_assert_1.default)(metric_ref01_data_dt0.id === metric_ref01_data.id);
        // REMOVE
        const metric_ref01_match_rm0 = { id: metric_ref01_data.id };
        await metric_ref01_ent.remove(metric_ref01_match_rm0);
        // LIST
        const metric_ref01_match_rt0 = {};
        metric_ref01_match_rt0['page_access_user_id'] = setup.idmap['page_access_user01'];
        metric_ref01_match_rt0['page_id'] = setup.idmap['page01'];
        const metric_ref01_list_rt0 = (await metric_ref01_ent.list(metric_ref01_match_rt0)).map((e) => e.data());
        (0, node_assert_1.default)(isempty(select(metric_ref01_list_rt0, { id: metric_ref01_data.id })));
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/metric/MetricTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['metric01', 'metric02', 'metric03', 'page01', 'page02', 'page03', 'metrics_provider01', 'metrics_provider02', 'metrics_provider03', 'page_access_user01', 'page_access_user02', 'page_access_user03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_METRIC_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_METRIC_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_METRIC_ENTID'];
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
//# sourceMappingURL=MetricEntity.test.js.map