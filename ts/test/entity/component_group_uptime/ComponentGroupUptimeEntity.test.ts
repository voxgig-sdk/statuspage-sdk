

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
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"component_id":{"a":true,"h":"Component Id","n":"component_id","r":false,"sh":"Component identifier","t":"`$STRING`","key$":"component_id","index$":0},"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":1},"incidents":{"a":true,"h":"Incidents","n":"incidents","r":false,"sh":"Related incidents","t":"`$OBJECT`","key$":"incidents","index$":2}},"id":{"field":"id","name":"id"},"name":"component_group_uptime","op":{"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /pages/{page_id}/component-groups/{id}/uptime","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":1}],"query":[{"a":true,"k":"query","n":"end","or":"end","r":false,"t":"Any","index$":0},{"a":true,"k":"query","n":"start","or":"start","r":false,"t":"Any","index$":1}]},"k":"http","m":"GET","o":"/pages/{page_id}/component-groups/{id}/uptime","q":{"exist":["end","id","page_id","start"]},"r":{},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"component-groups"},{"var":"id"},{"lit":"uptime"}],"t":{"req":"`reqdata`","res":"`body.related_events`"},"index$":0}],"key$":"load"}},"relations":{"ancestors":[["$.main.kit.entity.page"]]},"key$":"component_group_uptime","name__orig":"component_group_uptime","Name":"ComponentGroupUptime","name_":"component_group_uptime","name-":"component-group-uptime","NAME":"COMPONENT_GROUP_UPTIME","index$":1}, {"active":true,"entity":"component_group_uptime","key$":"BasicComponentGroupUptimeFlow","kind":"basic","name":"BasicComponentGroupUptimeFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"component_group_uptime_ref01","srcdatavar":"component_group_uptime_ref01_data","suffix":"_dt0"},"m":{"id":"component_group_uptime01","page_id":"page01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-component_group_uptime_ref01"}}],"index$":0}]}, 'ComponentGroupUptime', {"GET /pages/{page_id}/component-groups/{id}/uptime":{"protocol":"http","parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"id","description":"Component group identifier","required":true,"schema":{"type":"string"},"index$":1},{"in":"query","name":"start","description":"The start date for uptime calculation (defaults to the date of the component in the group with the earliest start_date, or 90 days ago, whichever is more recent).\nThe maximum supported date range is six calendar months. If the year is given, the date defaults to the first day of the year.\nIf the year and month are given, the start date defaults to the first day of that month.\nThe earliest supported date is January 1, 1970.\n","required":false,"schema":{"type":"PartialStartDate"},"index$":2},{"in":"query","name":"end","description":"The end date for uptime calculation (defaults to today in the page's time zone). The maximum supported date range is six calendar months.\nIf the year is given, the date defaults to the last day of the year. If the year and month are given, the date defaults to the last day of that month.\nThe earliest supported date is January 1, 1970.\n","required":false,"schema":{"type":"PartialEndDate"},"index$":3}]}})
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
  
