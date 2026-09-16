

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


describe('ComponentGroupUptimeEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.ComponentGroupUptime()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'component_group_uptime.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"component_id","req":false,"short":"Component identifier","type":"`$STRING`","index$":0},{"active":true,"name":"id","req":false,"type":"`$STRING`","index$":1},{"active":true,"name":"incidents","req":false,"short":"Related incidents","type":"`$OBJECT`","index$":2}],"id":{"field":"id","name":"id"},"name":"component_group_uptime","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"id","orig":"id","reqd":true,"type":"`$STRING`","index$":0},{"active":true,"kind":"param","name":"page_id","orig":"page_id","reqd":true,"type":"`$STRING`","index$":1}],"query":[{"active":true,"kind":"query","name":"end","orig":"end","reqd":false,"type":"Any","index$":0},{"active":true,"kind":"query","name":"start","orig":"start","reqd":false,"type":"Any","index$":1}]},"contract":{"id":"GET /pages/{page_id}/component-groups/{id}/uptime","json":"{\"operationId\":\"getPagesPageIdComponentGroupsIdUptime\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"Component group identifier\",\"in\":\"path\",\"name\":\"id\",\"required\":true,\"schema\":{\"type\":\"string\"}},{\"description\":\"The start date for uptime calculation (defaults to the date of the component in the group with the earliest start_date, or 90 days ago, whichever is more recent).\\nThe maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year.\\nIf the year and month are given, the start date defaults to the first day of that month.\\nThe earliest supported date is January 1, 1970.\\n\",\"in\":\"query\",\"name\":\"start\",\"required\":false,\"schema\":{\"type\":\"PartialStartDate\"}},{\"description\":\"The end date for uptime calculation (defaults to today in the page's time zone). The maximum supported date range is six calendar months.\\nIf the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month.\\nThe earliest supported date is January 1, 1970.\\n\",\"in\":\"query\",\"name\":\"end\",\"required\":false,\"schema\":{\"type\":\"PartialEndDate\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get uptime data for a component group that has uptime showcase enabled for at least one component.\",\"properties\":{\"id\":{\"description\":\"Component group identifier\",\"type\":\"string\"},\"major_outage\":{\"description\":\"Seconds of major outage\",\"example\":86400,\"format\":\"int32\",\"type\":\"integer\"},\"name\":{\"description\":\"Component group display name\",\"type\":\"string\"},\"partial_outage\":{\"description\":\"Seconds of partial outage\",\"example\":0,\"format\":\"int32\",\"type\":\"integer\"},\"range_end\":{\"description\":\"End date used for uptime calculation (see the warnings field in the response if this value does not match the end parameter you provided).\",\"example\":\"2020-02-15\",\"format\":\"date-time\",\"type\":\"string\"},\"range_start\":{\"description\":\"Start date used for uptime calculation (see the warnings field in the response if this value does not match the start parameter you provided).\",\"example\":\"2020-01-15\",\"format\":\"date-time\",\"type\":\"string\"},\"related_events\":{\"description\":\"Related incidents by component\",\"properties\":{\"component_id\":{\"description\":\"Component identifier\",\"type\":\"string\"},\"incidents\":{\"description\":\"Related incidents\",\"properties\":{\"id\":{\"description\":\"Incident identifier\",\"type\":\"string\"}},\"type\":\"object\"}},\"type\":\"object\"},\"uptime_percentage\":{\"description\":\"Uptime percentage for a component\",\"example\":96.67,\"format\":\"float\",\"type\":\"number\"},\"warnings\":{\"description\":\"Warning messages related to the uptime query that may occur\",\"example\":[\"End date was adjusted to today. See range_end field in response for end date used.\"],\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Get uptime data for a component group that has uptime showcase enabled for at least one component.\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"},\"422\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Unprocessable entity\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/pages/{page_id}/component-groups/{id}/uptime","segments":[{"lit":"pages"},{"var":"page_id"},{"lit":"component-groups"},{"var":"id"},{"lit":"uptime"}],"select":{"exist":["end","id","page_id","start"]},"transform":{"req":"`reqdata`","res":"`body.related_events`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[["page"]]},"key$":"component_group_uptime","name__orig":"component_group_uptime","Name":"ComponentGroupUptime","name_":"component_group_uptime","name-":"component-group-uptime","NAME":"COMPONENT_GROUP_UPTIME","index$":1}, {"active":true,"entity":"component_group_uptime","key$":"BasicComponentGroupUptimeFlow","kind":"basic","name":"BasicComponentGroupUptimeFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"component_group_uptime_ref01","srcdatavar":"component_group_uptime_ref01_data","suffix":"_dt0"},"match":{"id":"component_group_uptime01","page_id":"page01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-component_group_uptime_ref01"}}],"index$":0}]}, 'ComponentGroupUptime')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let component_group_uptime_ref01_data = Object.values(setup.data.existing.component_group_uptime)[0] as any

    // LOAD
    const component_group_uptime_ref01_ent = client.ComponentGroupUptime()
    const component_group_uptime_ref01_match_dt0: any = {}
    component_group_uptime_ref01_match_dt0.id = component_group_uptime_ref01_data.id
    const component_group_uptime_ref01_data_dt0 = (await component_group_uptime_ref01_ent.load(component_group_uptime_ref01_match_dt0)).data()
    assert(component_group_uptime_ref01_data_dt0.id === component_group_uptime_ref01_data.id)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/component_group_uptime/ComponentGroupUptimeTestData.json')

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
    ['component_group_uptime01','component_group_uptime02','component_group_uptime03','page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID']
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
  
