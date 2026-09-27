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
(0, node_test_1.describe)('PageEntity', async () => {
    // Per-test live pacing. Delay is read from sdk-test-control.json's
    // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
    (0, node_test_1.afterEach)((0, utility_1.liveDelay)('STATUSPAGE_TEST_LIVE'));
    (0, node_test_1.test)('instance', async () => {
        const testsdk = __1.StatuspageSDK.test();
        const ent = testsdk.Page();
        (0, node_assert_1.default)(null != ent);
    });
    (0, node_test_1.test)('basic', async (t) => {
        const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE;
        for (const op of ['list', 'update', 'load']) {
            if (!live && (0, utility_1.maybeSkipControl)(t, 'entityOp', 'page.' + op, live))
                return;
        }
        const setup = basicSetup();
        if (setup.live) {
            return (0, live_entity_1.runLiveEntity)(setup, { "active": true, "alias": { "field": {} }, "fields": { "activity_score": { "a": true, "fo": "float", "h": "Activity Score", "n": "activity_score", "r": false, "t": "`$NUMBER`", "key$": "activity_score", "index$": 0 }, "allow_email_subscribers": { "a": true, "h": "Allow Email Subscribers", "n": "allow_email_subscribers", "r": false, "sh": "Can your users choose to receive notifications via email", "t": "`$BOOLEAN`", "key$": "allow_email_subscribers", "index$": 1 }, "allow_incident_subscribers": { "a": true, "h": "Allow Incident Subscribers", "n": "allow_incident_subscribers", "r": false, "sh": "Can your users subscribe to notifications for a single incident", "t": "`$BOOLEAN`", "key$": "allow_incident_subscribers", "index$": 2 }, "allow_page_subscribers": { "a": true, "h": "Allow Page Subscribers", "n": "allow_page_subscribers", "r": false, "sh": "Can your users subscribe to all notifications on the page", "t": "`$BOOLEAN`", "key$": "allow_page_subscribers", "index$": 3 }, "allow_rss_atom_feeds": { "a": true, "h": "Allow Rss Atom Feeds", "n": "allow_rss_atom_feeds", "r": false, "sh": "Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)", "t": "`$BOOLEAN`", "key$": "allow_rss_atom_feeds", "index$": 4 }, "allow_sms_subscribers": { "a": true, "h": "Allow Sms Subscribers", "n": "allow_sms_subscribers", "r": false, "sh": "Can your users choose to receive notifications via SMS", "t": "`$BOOLEAN`", "key$": "allow_sms_subscribers", "index$": 5 }, "allow_webhook_subscribers": { "a": true, "h": "Allow Webhook Subscribers", "n": "allow_webhook_subscribers", "r": false, "sh": "Can your users choose to receive notifications via Webhooks", "t": "`$BOOLEAN`", "key$": "allow_webhook_subscribers", "index$": 6 }, "branding": { "a": true, "h": "Branding", "n": "branding", "r": false, "sh": "The main template your statuspage will use", "t": "`$STRING`", "key$": "branding", "index$": 7 }, "city": { "a": true, "h": "City", "n": "city", "r": false, "t": "`$STRING`", "key$": "city", "index$": 8 }, "country": { "a": true, "h": "Country", "n": "country", "r": false, "t": "`$STRING`", "key$": "country", "index$": 9 }, "created_at": { "a": true, "fo": "date-time", "h": "Created At", "n": "created_at", "r": false, "sh": "Timestamp the record was created", "t": "`$STRING`", "key$": "created_at", "index$": 10 }, "css_blues": { "a": true, "h": "Css Blues", "n": "css_blues", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_blues", "index$": 11 }, "css_body_background_color": { "a": true, "h": "Css Body Background Color", "n": "css_body_background_color", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_body_background_color", "index$": 12 }, "css_border_color": { "a": true, "h": "Css Border Color", "n": "css_border_color", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_border_color", "index$": 13 }, "css_font_color": { "a": true, "h": "Css Font Color", "n": "css_font_color", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_font_color", "index$": 14 }, "css_graph_color": { "a": true, "h": "Css Graph Color", "n": "css_graph_color", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_graph_color", "index$": 15 }, "css_greens": { "a": true, "h": "Css Greens", "n": "css_greens", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_greens", "index$": 16 }, "css_light_font_color": { "a": true, "h": "Css Light Font Color", "n": "css_light_font_color", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_light_font_color", "index$": 17 }, "css_link_color": { "a": true, "h": "Css Link Color", "n": "css_link_color", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_link_color", "index$": 18 }, "css_no_data": { "a": true, "h": "Css No Data", "n": "css_no_data", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_no_data", "index$": 19 }, "css_oranges": { "a": true, "h": "Css Oranges", "n": "css_oranges", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_oranges", "index$": 20 }, "css_reds": { "a": true, "h": "Css Reds", "n": "css_reds", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_reds", "index$": 21 }, "css_yellows": { "a": true, "h": "Css Yellows", "n": "css_yellows", "r": false, "sh": "CSS Color", "t": "`$STRING`", "key$": "css_yellows", "index$": 22 }, "domain": { "a": true, "h": "Domain", "n": "domain", "r": false, "sh": "CNAME alias for your status page", "t": "`$STRING`", "key$": "domain", "index$": 23 }, "email_logo": { "a": true, "h": "Email Logo", "n": "email_logo", "r": false, "t": "`$STRING`", "key$": "email_logo", "index$": 24 }, "favicon_logo": { "a": true, "h": "Favicon Logo", "n": "favicon_logo", "r": false, "t": "`$STRING`", "key$": "favicon_logo", "index$": 25 }, "headline": { "a": true, "h": "Headline", "n": "headline", "r": false, "t": "`$STRING`", "key$": "headline", "index$": 26 }, "hero_cover": { "a": true, "h": "Hero Cover", "n": "hero_cover", "r": false, "t": "`$STRING`", "key$": "hero_cover", "index$": 27 }, "hidden_from_search": { "a": true, "h": "Hidden From Search", "n": "hidden_from_search", "r": false, "sh": "Should your page hide itself from search engines", "t": "`$BOOLEAN`", "key$": "hidden_from_search", "index$": 28 }, "id": { "a": true, "h": "Id", "n": "id", "r": false, "sh": "Page identifier", "t": "`$STRING`", "key$": "id", "index$": 29 }, "ip_restrictions": { "a": true, "h": "Ip Restrictions", "n": "ip_restrictions", "r": false, "t": "`$STRING`", "key$": "ip_restrictions", "index$": 30 }, "name": { "a": true, "h": "Name", "n": "name", "r": false, "sh": "Name of your page to be displayed", "t": "`$STRING`", "key$": "name", "index$": 31 }, "notifications_email_footer": { "a": true, "h": "Notifications Email Footer", "n": "notifications_email_footer", "r": false, "sh": "Allows you to customize the footer appearing on your notification emails.", "t": "`$STRING`", "key$": "notifications_email_footer", "index$": 32 }, "notifications_from_email": { "a": true, "h": "Notifications From Email", "n": "notifications_from_email", "r": false, "sh": "Allows you to customize the email address your page notifications come from", "t": "`$STRING`", "key$": "notifications_from_email", "index$": 33 }, "page": { "a": true, "h": "Page", "n": "page", "r": false, "t": "`$OBJECT`", "key$": "page", "index$": 34 }, "page_description": { "a": true, "h": "Page Description", "n": "page_description", "r": false, "t": "`$STRING`", "key$": "page_description", "index$": 35 }, "state": { "a": true, "h": "State", "n": "state", "r": false, "t": "`$STRING`", "key$": "state", "index$": 36 }, "subdomain": { "a": true, "h": "Subdomain", "n": "subdomain", "r": false, "sh": "Subdomain at which to access your status page", "t": "`$STRING`", "key$": "subdomain", "index$": 37 }, "support_url": { "a": true, "h": "Support Url", "n": "support_url", "r": false, "t": "`$STRING`", "key$": "support_url", "index$": 38 }, "time_zone": { "a": true, "h": "Time Zone", "n": "time_zone", "r": false, "sh": "Timezone configured for your page", "t": "`$STRING`", "key$": "time_zone", "index$": 39 }, "transactional_logo": { "a": true, "h": "Transactional Logo", "n": "transactional_logo", "r": false, "t": "`$STRING`", "key$": "transactional_logo", "index$": 40 }, "twitter_logo": { "a": true, "h": "Twitter Logo", "n": "twitter_logo", "r": false, "t": "`$STRING`", "key$": "twitter_logo", "index$": 41 }, "twitter_username": { "a": true, "h": "Twitter Username", "n": "twitter_username", "r": false, "t": "`$STRING`", "key$": "twitter_username", "index$": 42 }, "updated_at": { "a": true, "fo": "date-time", "h": "Updated At", "n": "updated_at", "r": false, "sh": "Timestamp the record was last updated", "t": "`$STRING`", "key$": "updated_at", "index$": 43 }, "url": { "a": true, "h": "Url", "n": "url", "r": false, "sh": "Website of your page.", "t": "`$STRING`", "key$": "url", "index$": 44 }, "viewers_must_be_team_members": { "a": true, "h": "Viewers Must Be Team Members", "n": "viewers_must_be_team_members", "r": false, "t": "`$BOOLEAN`", "key$": "viewers_must_be_team_members", "index$": 45 } }, "id": { "field": "id", "name": "id" }, "name": "page", "op": { "list": { "input": "data", "name": "list", "points": [{ "a": true, "co": { "id": "GET /pages", "source": "openapi3", "version": 2 }, "g": {}, "k": "http", "m": "GET", "o": "/pages", "q": {}, "r": {}, "s": [{ "lit": "pages" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "list" }, "load": { "input": "data", "name": "load", "points": [{ "a": true, "co": { "id": "GET /pages/{page_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "GET", "o": "/pages/{page_id}", "q": { "exist": ["id"] }, "r": { "param": { "page_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "id" }], "t": { "req": "`reqdata`", "res": "`body`" }, "index$": 0 }], "key$": "load" }, "patch": { "input": "data", "name": "patch", "points": [{ "a": true, "co": { "id": "PATCH /pages/{page_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PATCH", "o": "/pages/{page_id}", "q": { "exist": ["id"] }, "r": { "param": { "page_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "id" }], "t": { "req": { "page": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "patch" }, "update": { "input": "data", "name": "update", "points": [{ "a": true, "co": { "id": "PUT /pages/{page_id}", "source": "openapi3", "version": 2 }, "g": { "params": [{ "a": true, "k": "param", "n": "id", "or": "page_id", "r": true, "t": "`$STRING`", "index$": 0 }] }, "k": "http", "m": "PUT", "o": "/pages/{page_id}", "q": { "exist": ["id"] }, "r": { "param": { "page_id": "id" } }, "s": [{ "lit": "pages" }, { "var": "id" }], "t": { "req": { "page": "`reqdata`" }, "res": "`body`" }, "index$": 0 }], "key$": "update" } }, "relations": { "ancestors": [] }, "key$": "page", "name__orig": "page", "Name": "Page", "name_": "page", "name-": "page", "NAME": "PAGE", "index$": 9 }, { "active": true, "entity": "page", "key$": "BasicPageFlow", "kind": "basic", "name": "BasicPageFlow", "param": {}, "step": [{ "a": true, "d": {}, "i": {}, "m": {}, "o": "list", "s": [], "v": [{ "apply": "ItemExists", "def": { "ref": "page_ref01" } }], "index$": 0 }, { "a": true, "d": {}, "i": { "ref": "page_ref01", "srcdatavar": "page_ref01_data", "suffix": "_up0", "textfield": "branding" }, "m": {}, "o": "update", "s": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-page_ref01" } }], "v": [], "index$": 1 }, { "a": true, "d": {}, "i": { "ref": "page_ref01", "srcdatavar": "page_ref01_data", "suffix": "_dt0" }, "m": { "id": "page01" }, "o": "load", "s": [], "v": [{ "apply": "TextFieldMark", "def": { "mark": "Mark01-page_ref01" } }], "index$": 2 }] }, 'Page', { "GET /pages": { "protocol": "http", "parameters": [] }, "GET /pages/{page_id}": { "protocol": "http", "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "PATCH /pages/{page_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "page": { "type": "object", "properties": { "name": { "type": "string", "description": "Name of your page to be displayed" }, "domain": { "type": "string", "description": "CNAME alias for your status page" }, "subdomain": { "type": "string", "description": "Subdomain at which to access your status page" }, "url": { "type": "string", "description": "Website of your page.  Clicking on your statuspage image will link here." }, "branding": { "type": "string", "description": "The main template your statuspage will use", "enum": ["premium", "basic"] }, "css_body_background_color": { "type": "string", "description": "CSS Color" }, "css_font_color": { "type": "string", "description": "CSS Color" }, "css_light_font_color": { "type": "string", "description": "CSS Color" }, "css_greens": { "type": "string", "description": "CSS Color" }, "css_yellows": { "type": "string", "description": "CSS Color" }, "css_oranges": { "type": "string", "description": "CSS Color" }, "css_reds": { "type": "string", "description": "CSS Color" }, "css_blues": { "type": "string", "description": "CSS Color" }, "css_border_color": { "type": "string", "description": "CSS Color" }, "css_graph_color": { "type": "string", "description": "CSS Color" }, "css_link_color": { "type": "string", "description": "CSS Color" }, "css_no_data": { "type": "string", "description": "CSS Color" }, "hidden_from_search": { "type": "boolean", "description": "Should your page hide itself from search engines" }, "viewers_must_be_team_members": { "type": "boolean" }, "allow_page_subscribers": { "type": "boolean", "description": "Can your users subscribe to all notifications on the page" }, "allow_incident_subscribers": { "type": "boolean", "description": "Can your users subscribe to notifications for a single incident" }, "allow_email_subscribers": { "type": "boolean", "description": "Can your users choose to receive notifications via email" }, "allow_sms_subscribers": { "type": "boolean", "description": "Can your users choose to receive notifications via SMS" }, "allow_rss_atom_feeds": { "type": "boolean", "description": "Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)" }, "allow_webhook_subscribers": { "type": "boolean", "description": "Can your users choose to receive notifications via Webhooks" }, "notifications_from_email": { "type": "string", "description": "Allows you to customize the email address your page notifications come from" }, "time_zone": { "type": "string", "description": "Timezone configured for your page" }, "notifications_email_footer": { "type": "string", "description": "Allows you to customize the footer appearing on your notification emails.  Accepts Markdown for formatting" } }, "key$": "page" } }, "description": "Update a page", "x-ref": "#/components/schemas/patchPages", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] }, "PUT /pages/{page_id}": { "protocol": "http", "requestBody": { "content": { "application/json": { "schema": { "type": "object", "properties": { "page": { "type": "object", "properties": { "name": { "type": "string", "description": "Name of your page to be displayed" }, "domain": { "type": "string", "description": "CNAME alias for your status page" }, "subdomain": { "type": "string", "description": "Subdomain at which to access your status page" }, "url": { "type": "string", "description": "Website of your page.  Clicking on your statuspage image will link here." }, "branding": { "type": "string", "description": "The main template your statuspage will use", "enum": ["premium", "basic"] }, "css_body_background_color": { "type": "string", "description": "CSS Color" }, "css_font_color": { "type": "string", "description": "CSS Color" }, "css_light_font_color": { "type": "string", "description": "CSS Color" }, "css_greens": { "type": "string", "description": "CSS Color" }, "css_yellows": { "type": "string", "description": "CSS Color" }, "css_oranges": { "type": "string", "description": "CSS Color" }, "css_reds": { "type": "string", "description": "CSS Color" }, "css_blues": { "type": "string", "description": "CSS Color" }, "css_border_color": { "type": "string", "description": "CSS Color" }, "css_graph_color": { "type": "string", "description": "CSS Color" }, "css_link_color": { "type": "string", "description": "CSS Color" }, "css_no_data": { "type": "string", "description": "CSS Color" }, "hidden_from_search": { "type": "boolean", "description": "Should your page hide itself from search engines" }, "viewers_must_be_team_members": { "type": "boolean" }, "allow_page_subscribers": { "type": "boolean", "description": "Can your users subscribe to all notifications on the page" }, "allow_incident_subscribers": { "type": "boolean", "description": "Can your users subscribe to notifications for a single incident" }, "allow_email_subscribers": { "type": "boolean", "description": "Can your users choose to receive notifications via email" }, "allow_sms_subscribers": { "type": "boolean", "description": "Can your users choose to receive notifications via SMS" }, "allow_rss_atom_feeds": { "type": "boolean", "description": "Can your users choose to access incident feeds via RSS/Atom (not functional on Audience-Specific pages)" }, "allow_webhook_subscribers": { "type": "boolean", "description": "Can your users choose to receive notifications via Webhooks" }, "notifications_from_email": { "type": "string", "description": "Allows you to customize the email address your page notifications come from" }, "time_zone": { "type": "string", "description": "Timezone configured for your page" }, "notifications_email_footer": { "type": "string", "description": "Allows you to customize the footer appearing on your notification emails.  Accepts Markdown for formatting" } }, "key$": "page" } }, "description": "Update a page", "x-ref": "#/components/schemas/putPages", "index$": 1 } } }, "required": true }, "parameters": [{ "in": "path", "name": "page_id", "description": "Page identifier", "required": true, "schema": { "type": "string" }, "index$": 0 }] } });
        }
        const client = setup.client;
        const struct = setup.struct;
        const isempty = struct.isempty;
        const select = struct.select;
        let page_ref01_data = Object.values(setup.data.existing.page)[0];
        // LIST
        const page_ref01_ent = client.Page();
        const page_ref01_match = {};
        const page_ref01_list = (await page_ref01_ent.list(page_ref01_match)).map((e) => e.data());
        // UPDATE
        const page_ref01_data_up0 = {};
        page_ref01_data_up0.id = page_ref01_data.id;
        const page_ref01_markdef_up0 = { name: 'branding', value: 'Mark01-page_ref01_' + setup.now };
        page_ref01_data_up0[page_ref01_markdef_up0.name] = page_ref01_markdef_up0.value;
        const page_ref01_resdata_up0 = (await page_ref01_ent.update(page_ref01_data_up0)).data();
        (0, node_assert_1.default)(page_ref01_resdata_up0.id === page_ref01_data_up0.id);
        (0, node_assert_1.default)(page_ref01_resdata_up0[page_ref01_markdef_up0.name] === page_ref01_markdef_up0.value);
        // LOAD
        const page_ref01_match_dt0 = {};
        page_ref01_match_dt0.id = page_ref01_data.id;
        const page_ref01_data_dt0 = (await page_ref01_ent.load(page_ref01_match_dt0)).data();
        (0, node_assert_1.default)(page_ref01_data_dt0.id === page_ref01_data.id);
    });
});
function basicSetup(extra) {
    // TODO: fix test def options
    const options = {}; // null
    // TODO: needs test utility to resolve path
    const entityDataFile = node_path_1.default.resolve(__dirname, '../../../../.sdk/test/entity/page/PageTestData.json');
    // TODO: file ready util needed?
    const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8');
    // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
    const entityData = JSON.parse(entityDataSource);
    options.entity = entityData.existing;
    let client = __1.StatuspageSDK.test(options, extra);
    const struct = client.utility().struct;
    const merge = struct.merge;
    const transform = struct.transform;
    let idmap = transform(['page01', 'page02', 'page03'], {
        '`$PACK`': ['', {
                '`$KEY`': '`$COPY`',
                '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
            }]
    });
    const env = (0, utility_1.envOverride)({
        'STATUSPAGE_TEST_PAGE_ENTID': idmap,
        'STATUSPAGE_TEST_LIVE': 'FALSE',
        'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
        'STATUSPAGE_APIKEY': '',
    });
    idmap = env['STATUSPAGE_TEST_PAGE_ENTID'];
    const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE;
    const transport = (0, live_runner_1.createLiveTransport)();
    if (live) {
        const rawIds = process.env['STATUSPAGE_TEST_PAGE_ENTID'];
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
//# sourceMappingURL=PageEntity.test.js.map