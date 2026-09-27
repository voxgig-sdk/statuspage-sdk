

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


describe('IncidentUpdateEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.IncidentUpdate()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['update']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'incident_update.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"affected_components":{"a":true,"h":"Affected Components","n":"affected_components","r":false,"sh":"Affected components associated with the incident update.","t":"`$ARRAY`","key$":"affected_components","index$":0},"body":{"a":true,"h":"Body","n":"body","r":false,"sh":"Incident update body.","t":"`$STRING`","key$":"body","index$":1},"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":false,"sh":"The timestamp when the incident update was created at.","t":"`$STRING`","key$":"created_at","index$":2},"custom_tweet":{"a":true,"h":"Custom Tweet","n":"custom_tweet","r":false,"sh":"An optional customized tweet message for incident postmortem.","t":"`$STRING`","key$":"custom_tweet","index$":3},"deliver_notifications":{"a":true,"h":"Deliver Notifications","n":"deliver_notifications","r":false,"sh":"Controls whether to delivery notifications.","t":"`$BOOLEAN`","key$":"deliver_notifications","index$":4},"display_at":{"a":true,"fo":"date-time","h":"Display At","n":"display_at","r":false,"sh":"Timestamp when incident update is happened.","t":"`$STRING`","key$":"display_at","index$":5},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"Incident Update Identifier.","t":"`$STRING`","key$":"id","index$":6},"incident_id":{"a":true,"h":"Incident Id","n":"incident_id","r":false,"sh":"Incident Identifier.","t":"`$STRING`","key$":"incident_id","index$":7},"incident_update":{"a":true,"h":"Incident Update","n":"incident_update","r":false,"t":"`$OBJECT`","key$":"incident_update","index$":8},"status":{"a":true,"h":"Status","n":"status","r":false,"sh":"The incident status.","t":"`$STRING`","key$":"status","index$":9},"tweet_id":{"a":true,"h":"Tweet Id","n":"tweet_id","r":false,"sh":"Tweet identifier associated to this incident update.","t":"`$STRING`","key$":"tweet_id","index$":10},"twitter_updated_at":{"a":true,"fo":"date-time","h":"Twitter Updated At","n":"twitter_updated_at","r":false,"sh":"The timestamp when twitter updated at.","t":"`$STRING`","key$":"twitter_updated_at","index$":11},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":false,"sh":"The timestamp when the incident update is updated.","t":"`$STRING`","key$":"updated_at","index$":12},"wants_twitter_update":{"a":true,"h":"Wants Twitter Update","n":"wants_twitter_update","r":false,"sh":"Controls whether to create twitter update.","t":"`$BOOLEAN`","key$":"wants_twitter_update","index$":13}},"id":{"field":"id","name":"id"},"name":"incident_update","op":{"patch":{"input":"data","name":"patch","points":[{"a":true,"co":{"id":"PATCH /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"incident_update_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"incident_id","or":"incident_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":2}]},"k":"http","m":"PATCH","o":"/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}","q":{"exist":["id","incident_id","page_id"]},"r":{"param":{"incident_update_id":"id"}},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"incident_id"},{"lit":"incident_updates"},{"var":"id"}],"t":{"req":{"incident_update":"`reqdata`"},"res":"`body`"},"index$":0}],"key$":"patch"},"update":{"input":"data","name":"update","points":[{"a":true,"co":{"id":"PUT /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"incident_update_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"incident_id","or":"incident_id","r":true,"t":"`$STRING`","index$":1},{"a":true,"k":"param","n":"page_id","or":"page_id","r":true,"t":"`$STRING`","index$":2}]},"k":"http","m":"PUT","o":"/pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}","q":{"exist":["id","incident_id","page_id"]},"r":{"param":{"incident_update_id":"id"}},"s":[{"lit":"pages"},{"var":"page_id"},{"lit":"incidents"},{"var":"incident_id"},{"lit":"incident_updates"},{"var":"id"}],"t":{"req":{"incident_update":"`reqdata`"},"res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[["$.main.kit.entity.page","$.main.kit.entity.incident"]]},"key$":"incident_update","name__orig":"incident_update","Name":"IncidentUpdate","name_":"incident_update","name-":"incident-update","NAME":"INCIDENT_UPDATE","index$":6}, {"active":true,"entity":"incident_update","key$":"BasicIncidentUpdateFlow","kind":"basic","name":"BasicIncidentUpdateFlow","param":{},"step":[{"a":true,"d":{"incident_id":"incident01","page_id":"page01"},"i":{"ref":"incident_update_ref01","srcdatavar":"incident_update_ref01_data","suffix":"_up0","textfield":"body"},"m":{},"o":"update","s":[{"apply":"TextFieldMark","def":{"mark":"Mark01-incident_update_ref01"}}],"v":[],"index$":0}]}, 'IncidentUpdate', {"PATCH /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"incident_update":{"type":"object","properties":{"wants_twitter_update":{"type":"boolean","description":"Controls whether to create twitter update."},"body":{"type":"string","description":"Incident update body."},"display_at":{"type":"string","format":"date-time","description":"Timestamp when incident update is happened."},"deliver_notifications":{"type":"boolean","description":"Controls whether to delivery notifications."}},"key$":"incident_update"}},"description":"Update a previous incident update","x-ref":"#/components/schemas/patchPagesPageIdIncidentsIncidentIdIncidentUpdates","index$":1}}},"required":true},"parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1},{"in":"path","name":"incident_update_id","description":"Incident Update Identifier","required":true,"schema":{"type":"string"},"index$":2}]},"PUT /pages/{page_id}/incidents/{incident_id}/incident_updates/{incident_update_id}":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"incident_update":{"type":"object","properties":{"wants_twitter_update":{"type":"boolean","description":"Controls whether to create twitter update."},"body":{"type":"string","description":"Incident update body."},"display_at":{"type":"string","format":"date-time","description":"Timestamp when incident update is happened."},"deliver_notifications":{"type":"boolean","description":"Controls whether to delivery notifications."}},"key$":"incident_update"}},"description":"Update a previous incident update","x-ref":"#/components/schemas/putPagesPageIdIncidentsIncidentIdIncidentUpdates","index$":1}}},"required":true},"parameters":[{"in":"path","name":"page_id","description":"Page identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"incident_id","description":"Incident Identifier","required":true,"schema":{"type":"string"},"index$":1},{"in":"path","name":"incident_update_id","description":"Incident Update Identifier","required":true,"schema":{"type":"string"},"index$":2}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let incident_update_ref01_data = Object.values(setup.data.existing.incident_update)[0] as any

    // UPDATE
    const incident_update_ref01_ent = client.IncidentUpdate()
    const incident_update_ref01_data_up0: any = {}
    incident_update_ref01_data_up0.id = incident_update_ref01_data.id
    incident_update_ref01_data_up0 ['incident_id'] = setup.idmap['incident_id']
    incident_update_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const incident_update_ref01_markdef_up0 = { name: 'body', value: 'Mark01-incident_update_ref01_' + setup.now }
    ;(incident_update_ref01_data_up0 as any)[incident_update_ref01_markdef_up0.name] = incident_update_ref01_markdef_up0.value

    const incident_update_ref01_resdata_up0 = (await incident_update_ref01_ent.update(incident_update_ref01_data_up0)).data()
    assert(incident_update_ref01_resdata_up0.id === incident_update_ref01_data_up0.id)

    assert((incident_update_ref01_resdata_up0 as any)[incident_update_ref01_markdef_up0.name] === incident_update_ref01_markdef_up0.value)


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/incident_update/IncidentUpdateTestData.json')

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
    ['incident_update01','incident_update02','incident_update03','page01','page02','page03','incident01','incident02','incident03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_INCIDENT_UPDATE_ENTID']
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
  
