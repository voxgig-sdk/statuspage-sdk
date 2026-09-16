

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { StatuspageSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('IncidentTemplateEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.IncidentTemplate()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['create', 'list']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'incident_template.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"body","req":false,"short":"Body of the incident or maintenance update to be applied when selecting this template","type":"`$STRING`","index$":0},{"active":true,"name":"components","req":false,"short":"Affected components","type":"`$ARRAY`","index$":1},{"active":true,"name":"group_id","req":false,"short":"Identifier of Template Group this template belongs to","type":"`$STRING`","index$":2},{"active":true,"name":"id","req":false,"short":"Incident Template Identifier","type":"`$STRING`","index$":3},{"active":true,"name":"name","req":false,"short":"Name of the template, as shown in the list on the \"Templates\" tab of the \"Incidents\" page","type":"`$STRING`","index$":4},{"active":true,"name":"should_send_notifications","req":false,"short":"Whether the \"deliver notifications\" checkbox should be selected when selecting this template","type":"`$BOOLEAN`","index$":5},{"active":true,"name":"should_tweet","req":false,"short":"Whether the \"tweet update\" checkbox should be selected when selecting this template","type":"`$BOOLEAN`","index$":6},{"active":true,"name":"template","req":true,"type":"`$OBJECT`","index$":7},{"active":true,"name":"title","req":false,"short":"Title to be applied to the incident or maintenance when selecting this template","type":"`$STRING`","index$":8},{"active":true,"name":"update_status","req":false,"short":"The status the incident or maintenance should transition to when selecting this template","type":"`$STRING`","index$":9}],"id":{"field":"id","name":"id"},"name":"incident_template","op":{"create":{"input":"data","name":"create","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"page_id","orig":"page_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"POST /pages/{page_id}/incident_templates","json":"{\"operationId\":\"postPagesPageIdIncidentTemplates\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Create a template\",\"properties\":{\"template\":{\"properties\":{\"body\":{\"description\":\"The initial message, created as the first incident or maintenance update.\",\"type\":\"string\"},\"component_ids\":{\"description\":\"List of component_ids affected by this incident\",\"items\":{\"default\":\"1ss1vgkq6034\",\"type\":\"string\"},\"type\":\"array\"},\"group_id\":{\"description\":\"Identifier of Template Group this template belongs to\",\"type\":\"string\"},\"name\":{\"description\":\"Name of the template, as shown in the list on the \\\"Templates\\\" tab of the \\\"Incidents\\\" page\",\"type\":\"string\"},\"should_send_notifications\":{\"description\":\"Whether the \\\"deliver notifications\\\" checkbox should be selected when selecting this template\",\"type\":\"boolean\"},\"should_tweet\":{\"description\":\"Whether the \\\"tweet update\\\" checkbox should be selected when selecting this template\",\"type\":\"boolean\"},\"title\":{\"description\":\"Title to be applied to the incident or maintenance when selecting this template\",\"type\":\"string\"},\"update_status\":{\"description\":\"The status the incident or maintenance should transition to when selecting this template\",\"enum\":[\"investigating\",\"identified\",\"monitoring\",\"resolved\",\"scheduled\",\"in_progress\",\"verifying\",\"completed\"],\"type\":\"string\"}},\"required\":[\"name\",\"title\",\"body\"],\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"201\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of templates\",\"properties\":{\"body\":{\"description\":\"Body of the incident or maintenance update to be applied when selecting this template\",\"type\":\"string\"},\"components\":{\"description\":\"Affected components\",\"items\":{\"description\":\"Add page access groups to a component\",\"properties\":{\"automation_email\":{\"description\":\"Requires a special feature flag to be enabled\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"description\":{\"description\":\"More detailed description for component\",\"type\":\"string\"},\"group\":{\"description\":\"Is this component a group\",\"type\":\"boolean\"},\"group_id\":{\"description\":\"Component Group identifier\",\"type\":\"string\"},\"id\":{\"description\":\"Identifier for component\",\"type\":\"string\"},\"name\":{\"description\":\"Display name for component\",\"type\":\"string\"},\"only_show_if_degraded\":{\"description\":\"Requires a special feature flag to be enabled\",\"type\":\"boolean\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"},\"position\":{\"description\":\"Order the component will appear on the page\",\"format\":\"int32\",\"type\":\"integer\"},\"showcase\":{\"description\":\"Should this component be showcased\",\"type\":\"boolean\"},\"start_date\":{\"description\":\"The date this component started being used\",\"format\":\"date\",\"type\":\"string\"},\"status\":{\"description\":\"Status of component\",\"enum\":[\"operational\",\"under_maintenance\",\"degraded_performance\",\"partial_outage\",\"major_outage\",\"\"],\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"},\"group_id\":{\"description\":\"Identifier of Template Group this template belongs to\",\"example\":\"zljczr8kgxth\",\"type\":\"string\"},\"id\":{\"description\":\"Incident Template Identifier\",\"example\":\"m4f04px479hx\",\"type\":\"string\"},\"name\":{\"description\":\"Name of the template, as shown in the list on the \\\"Templates\\\" tab of the \\\"Incidents\\\" page\",\"type\":\"string\"},\"should_send_notifications\":{\"description\":\"Whether the \\\"deliver notifications\\\" checkbox should be selected when selecting this template\",\"type\":\"boolean\"},\"should_tweet\":{\"description\":\"Whether the \\\"tweet update\\\" checkbox should be selected when selecting this template\",\"type\":\"boolean\"},\"title\":{\"description\":\"Title to be applied to the incident or maintenance when selecting this template\",\"type\":\"string\"},\"update_status\":{\"description\":\"The status the incident or maintenance should transition to when selecting this template\",\"enum\":[\"investigating\",\"identified\",\"monitoring\",\"resolved\",\"scheduled\",\"in_progress\",\"verifying\",\"completed\"],\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Create a template\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"},\"422\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Unprocessable entity\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"POST","orig":"/pages/{page_id}/incident_templates","segments":[{"lit":"pages"},{"var":"page_id"},{"lit":"incident_templates"}],"select":{"exist":["page_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"page_id","orig":"page_id","reqd":true,"type":"`$STRING`","index$":0}],"query":[{"active":true,"example":1,"kind":"query","name":"page","orig":"page","reqd":false,"type":"`$INTEGER`","index$":0},{"active":true,"example":100,"kind":"query","name":"per_page","orig":"per_page","reqd":false,"type":"`$INTEGER`","index$":1}]},"contract":{"id":"GET /pages/{page_id}/incident_templates","json":"{\"operationId\":\"getPagesPageIdIncidentTemplates\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Page offset to fetch.\",\"in\":\"query\",\"name\":\"page\",\"required\":false,\"schema\":{\"default\":1,\"format\":\"int32\",\"type\":\"integer\"}},{\"description\":\"Number of results to return per page.\",\"in\":\"query\",\"name\":\"per_page\",\"required\":false,\"schema\":{\"default\":100,\"format\":\"int32\",\"type\":\"integer\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"items\":{\"description\":\"Get a list of templates\",\"properties\":{\"body\":{\"description\":\"Body of the incident or maintenance update to be applied when selecting this template\",\"type\":\"string\"},\"components\":{\"description\":\"Affected components\",\"items\":{\"description\":\"Add page access groups to a component\",\"properties\":{\"automation_email\":{\"description\":\"Requires a special feature flag to be enabled\",\"type\":\"string\"},\"created_at\":{\"format\":\"date-time\",\"type\":\"string\"},\"description\":{\"description\":\"More detailed description for component\",\"type\":\"string\"},\"group\":{\"description\":\"Is this component a group\",\"type\":\"boolean\"},\"group_id\":{\"description\":\"Component Group identifier\",\"type\":\"string\"},\"id\":{\"description\":\"Identifier for component\",\"type\":\"string\"},\"name\":{\"description\":\"Display name for component\",\"type\":\"string\"},\"only_show_if_degraded\":{\"description\":\"Requires a special feature flag to be enabled\",\"type\":\"boolean\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"},\"position\":{\"description\":\"Order the component will appear on the page\",\"format\":\"int32\",\"type\":\"integer\"},\"showcase\":{\"description\":\"Should this component be showcased\",\"type\":\"boolean\"},\"start_date\":{\"description\":\"The date this component started being used\",\"format\":\"date\",\"type\":\"string\"},\"status\":{\"description\":\"Status of component\",\"enum\":[\"operational\",\"under_maintenance\",\"degraded_performance\",\"partial_outage\",\"major_outage\",\"\"],\"type\":\"string\"},\"updated_at\":{\"format\":\"date-time\",\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"},\"group_id\":{\"description\":\"Identifier of Template Group this template belongs to\",\"example\":\"zljczr8kgxth\",\"type\":\"string\"},\"id\":{\"description\":\"Incident Template Identifier\",\"example\":\"m4f04px479hx\",\"type\":\"string\"},\"name\":{\"description\":\"Name of the template, as shown in the list on the \\\"Templates\\\" tab of the \\\"Incidents\\\" page\",\"type\":\"string\"},\"should_send_notifications\":{\"description\":\"Whether the \\\"deliver notifications\\\" checkbox should be selected when selecting this template\",\"type\":\"boolean\"},\"should_tweet\":{\"description\":\"Whether the \\\"tweet update\\\" checkbox should be selected when selecting this template\",\"type\":\"boolean\"},\"title\":{\"description\":\"Title to be applied to the incident or maintenance when selecting this template\",\"type\":\"string\"},\"update_status\":{\"description\":\"The status the incident or maintenance should transition to when selecting this template\",\"enum\":[\"investigating\",\"identified\",\"monitoring\",\"resolved\",\"scheduled\",\"in_progress\",\"verifying\",\"completed\"],\"type\":\"string\"}},\"type\":\"object\"},\"type\":\"array\"}}},\"description\":\"Get a list of templates\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/pages/{page_id}/incident_templates","segments":[{"lit":"pages"},{"var":"page_id"},{"lit":"incident_templates"}],"select":{"exist":["page","page_id","per_page"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"}},"relations":{"ancestors":[["page"]]},"key$":"incident_template","name__orig":"incident_template","Name":"IncidentTemplate","name_":"incident_template","name-":"incident-template","NAME":"INCIDENT_TEMPLATE","index$":6}, {"active":true,"entity":"incident_template","key$":"BasicIncidentTemplateFlow","kind":"basic","name":"BasicIncidentTemplateFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"incident_template_ref01"},"match":{"page_id":"page01"},"op":"create","spec":[],"valid":[],"index$":0},{"active":true,"data":{},"input":{},"match":{"page_id":"page01"},"op":"list","spec":[],"valid":[{"apply":"ItemExists","def":{"ref":"incident_template_ref01"}}],"index$":1}]}, 'IncidentTemplate')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const incident_template_ref01_ent = client.IncidentTemplate()
    let incident_template_ref01_data = setup.data.new.incident_template['incident_template_ref01']
    incident_template_ref01_data['page_id'] = setup.idmap['page01']

    incident_template_ref01_data = (await incident_template_ref01_ent.create(incident_template_ref01_data)).data()
    assert(null != incident_template_ref01_data.id)


    // LIST
    const incident_template_ref01_match: any = {}
    incident_template_ref01_match['page_id'] = setup.idmap['page01']

    const incident_template_ref01_list = (await incident_template_ref01_ent.list(incident_template_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(incident_template_ref01_list, { id: incident_template_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/incident_template/IncidentTemplateTestData.json')

  // TODO: file ready util needed?
  const entityDataSource = Fs.readFileSync(entityDataFile).toString('utf8')

  // TODO: need a xlang JSON parse utility in voxgig/struct with better error msgs
  const entityData = JSON.parse(entityDataSource)

  options.entity = entityData.existing

  let client = StatuspageSDK.test(options, extra)
  const struct = client.utility().struct
  const merge = struct.merge
  const transform = struct.transform

  let idmap = transform(
    ['incident_template01','incident_template02','incident_template03','page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new StatuspageSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
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
    ]))
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
  }

  return setup
}
  
