

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


describe('IncidentPostmortemEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.IncidentPostmortem()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of []) {
      if (!live && maybeSkipControl(t, 'entityOp', 'incident_postmortem.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"id":{"a":true,"h":"Id","n":"id","r":false,"t":"`$STRING`","key$":"id","index$":0}},"id":{"field":"id","name":"id"},"name":"incident_postmortem","op":{"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /pages/{page_id}/incidents/{incident_id}/postmortem","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"incident_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/pages/{page_id}/incidents/{incident_id}/postmortem","q":{"exist":["id","page_id"]},"r":{"param":{"incident_id":"id"}},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"id"},{"lit":"postmortem"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"}},"relations":{"ancestors":[["$.main.kit.entity.page"]]},"key$":"incident_postmortem","name__orig":"incident_postmortem","Name":"IncidentPostmortem","name_":"incident_postmortem","name-":"incident-postmortem","NAME":"INCIDENT_POSTMORTEM","index$":4}, {"active":true,"entity":"incident_postmortem","key$":"BasicIncidentPostmortemFlow","kind":"basic","name":"BasicIncidentPostmortemFlow","param":{},"step":[]}, 'IncidentPostmortem', {"DELETE /pages/{page_id}/incidents/{incident_id}/postmortem":{"protocol":"http","parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let incident_postmortem_ref01_data = Object.values(setup.data.existing.incident_postmortem)[0] as any

  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/incident_postmortem/IncidentPostmortemTestData.json')

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
    ['incident_postmortem01','incident_postmortem02','incident_postmortem03','page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_INCIDENT_POSTMORTEM_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_INCIDENT_POSTMORTEM_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_INCIDENT_POSTMORTEM_ENTID']
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
  
