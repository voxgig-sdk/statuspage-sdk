

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


describe('PostmortemEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.Postmortem()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['update', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'postmortem.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"body":{"a":true,"h":"Body","n":"body","r":false,"sh":"Postmortem body","t":"`$STRING`","key$":"body","index$":0},"body_draft":{"a":true,"h":"Body Draft","n":"body_draft","r":false,"sh":"Body draft","t":"`$STRING`","key$":"body_draft","index$":1},"body_draft_updated_at":{"a":true,"fo":"date-time","h":"Body Draft Updated At","n":"body_draft_updated_at","r":false,"t":"`$STRING`","key$":"body_draft_updated_at","index$":2},"body_updated_at":{"a":true,"fo":"date-time","h":"Body Updated At","n":"body_updated_at","r":false,"t":"`$STRING`","key$":"body_updated_at","index$":3},"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":false,"t":"`$STRING`","key$":"created_at","index$":4},"custom_tweet":{"a":true,"h":"Custom Tweet","n":"custom_tweet","r":false,"sh":"Custom tweet for Incident Postmortem","t":"`$STRING`","key$":"custom_tweet","index$":5},"notify_subscribers":{"a":true,"h":"Notify Subscribers","n":"notify_subscribers","r":false,"sh":"Should email subscribers be notified.","t":"`$BOOLEAN`","key$":"notify_subscribers","index$":6},"notify_twitter":{"a":true,"h":"Notify Twitter","n":"notify_twitter","r":false,"sh":"Should Twitter followers be notified.","t":"`$BOOLEAN`","key$":"notify_twitter","index$":7},"postmortem":{"a":true,"h":"Postmortem","n":"postmortem","r":true,"t":"`$OBJECT`","key$":"postmortem","index$":8},"preview_key":{"a":true,"h":"Preview Key","n":"preview_key","r":false,"sh":"Preview Key","t":"`$STRING`","key$":"preview_key","index$":9},"published_at":{"a":true,"fo":"date-time","h":"Published At","n":"published_at","r":false,"t":"`$STRING`","key$":"published_at","index$":10},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":false,"t":"`$STRING`","key$":"updated_at","index$":11}},"name":"postmortem","op":{"load":{"input":"data","name":"load","points":[{"a":true,"co":{"id":"GET /pages/{page_id}/incidents/{incident_id}/postmortem","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"incident_id","or":"incident_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"GET","o":"/pages/{page_id}/incidents/{incident_id}/postmortem","q":{"exist":["incident_id","page_id"]},"r":{},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"incident_id"},{"lit":"postmortem"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PUT /pages/{page_id}/incidents/{incident_id}/postmortem","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"incident_id","or":"incident_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PUT","o":"/pages/{page_id}/incidents/{incident_id}/postmortem","q":{"exist":["incident_id","page_id"]},"r":{},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"incident_id"},{"lit":"postmortem"}],"t":{"req":{"postmortem":"`reqdata`"},"res":"`body`"},"index$":0},{"a":true,"co":{"id":"PUT /pages/{page_id}/incidents/{incident_id}/postmortem/publish","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"incident_id","or":"incident_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PUT","o":"/pages/{page_id}/incidents/{incident_id}/postmortem/publish","q":{"$action":"publish","exist":["incident_id","page_id"]},"r":{},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"incident_id"},{"lit":"postmortem"},{"lit":"publish"}],"t":{"req":{"postmortem":"`reqdata`"},"res":"`body`"},"index$":1},{"a":true,"co":{"id":"PUT /pages/{page_id}/incidents/{incident_id}/postmortem/revert","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"incident_id","or":"incident_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"PUT","o":"/pages/{page_id}/incidents/{incident_id}/postmortem/revert","q":{"$action":"revert","exist":["incident_id","page_id"]},"r":{},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"incident_id"},{"lit":"postmortem"},{"lit":"revert"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":2}],"key$":"update"}},"relations":{"ancestors":[["$.main.kit.entity.page","$.main.kit.entity.incident"]]},"key$":"postmortem","name__orig":"postmortem","Name":"Postmortem","name_":"postmortem","name-":"postmortem","NAME":"POSTMORTEM","index$":13}, {"active":true,"entity":"postmortem","key$":"BasicPostmortemFlow","kind":"basic","name":"BasicPostmortemFlow","param":{},"step":[{"a":true,"d":{"page_id":"page01"},"i":{"ref":"postmortem_ref01","srcdatavar":"postmortem_ref01_data","suffix":"_up0","textfield":"body"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-postmortem_ref01"}}],"v":[],"index$":0},{"a":true,"d":{},"i":{"ref":"postmortem_ref01","srcdatavar":"postmortem_ref01_data","suffix":"_dt0"},"m":{"id":"postmortem01","page_id":"page01"},"o":"load","s":[],"v":[{"apply":"TextFieldMark","def":{"mark":"Mark01-postmortem_ref01"}}],"index$":1}]}, 'Postmortem', {"GET /pages/{page_id}/incidents/{incident_id}/postmortem":{"protocol":"http","parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1}]},"PUT /pages/{page_id}/incidents/{incident_id}/postmortem":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"postmortem":{"type":"object","properties":{"body_draft":{"type":"string","description":"Body of Postmortem to create."}},"required":["body_draft"],"key$":"postmortem"}},"description":"Create Postmortem","x-ref":"#/components/schemas/putPagesPageIdIncidentsIncidentIdPostmortem","index$":1}}},"required":true},"parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1}]},"PUT /pages/{page_id}/incidents/{incident_id}/postmortem/publish":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"postmortem":{"type":"object","properties":{"notify_twitter":{"type":"boolean","description":"Whether to notify Twitter followers"},"notify_subscribers":{"type":"boolean","description":"Whether to notify e-mail subscribers"},"custom_tweet":{"type":"string","description":"Custom postmortem tweet to publish"}}}},"description":"Publish Postmortem","x-ref":"#/components/schemas/putPagesPageIdIncidentsIncidentIdPostmortemPublish"}}},"required":true},"parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1}]},"PUT /pages/{page_id}/incidents/{incident_id}/postmortem/revert":{"protocol":"http","parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let postmortem_ref01_data = Object.values(setup.data.existing.postmortem)[0] as any

    // UPDATE
    const postmortem_ref01_ent = client.Postmortem()
    const postmortem_ref01_data_up0: any = {}
    postmortem_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const postmortem_ref01_markdef_up0 = { name: 'body', value: 'Mark01-postmortem_ref01_' + setup.now }
    ;(postmortem_ref01_data_up0 as any)[postmortem_ref01_markdef_up0.name] = postmortem_ref01_markdef_up0.value

    const postmortem_ref01_resdata_up0 = (await postmortem_ref01_ent.update(postmortem_ref01_data_up0)).data()
    assert(null != postmortem_ref01_resdata_up0)

    assert((postmortem_ref01_resdata_up0 as any)[postmortem_ref01_markdef_up0.name] === postmortem_ref01_markdef_up0.value)



  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/postmortem/PostmortemTestData.json')

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
    ['postmortem01','postmortem02','postmortem03','page01','page02','page03','incident01','incident02','incident03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_POSTMORTEM_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_POSTMORTEM_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_POSTMORTEM_ENTID']
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
  
