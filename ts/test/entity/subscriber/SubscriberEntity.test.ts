
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'


import { StatuspageSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveDelay,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


describe('SubscriberEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.Subscriber()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (maybeSkipControl(t, 'entityOp', 'subscriber.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_SUBSCRIBER_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const subscriber_ref01_ent = client.Subscriber()
    let subscriber_ref01_data = setup.data.new.subscriber['subscriber_ref01']
    subscriber_ref01_data['incident_id'] = setup.idmap['incident01']
    subscriber_ref01_data['page_id'] = setup.idmap['page01']

    subscriber_ref01_data = await subscriber_ref01_ent.create(subscriber_ref01_data)
    assert(null != subscriber_ref01_data.id)


    // LIST
    const subscriber_ref01_match: any = {}
    subscriber_ref01_match['page_id'] = setup.idmap['page01']

    const subscriber_ref01_list = await subscriber_ref01_ent.list(subscriber_ref01_match)

    assert(!isempty(select(subscriber_ref01_list, { id: subscriber_ref01_data.id })))


    // UPDATE
    const subscriber_ref01_data_up0: any = {}
    subscriber_ref01_data_up0.id = subscriber_ref01_data.id
    subscriber_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const subscriber_ref01_markdef_up0 = { name: 'component', value: 'Mark01-subscriber_ref01_' + setup.now }
    ;(subscriber_ref01_data_up0 as any)[subscriber_ref01_markdef_up0.name] = subscriber_ref01_markdef_up0.value

    const subscriber_ref01_resdata_up0 = await subscriber_ref01_ent.update(subscriber_ref01_data_up0)
    assert(subscriber_ref01_resdata_up0.id === subscriber_ref01_data_up0.id)

    assert((subscriber_ref01_resdata_up0 as any)[subscriber_ref01_markdef_up0.name] === subscriber_ref01_markdef_up0.value)


    // LOAD
    const subscriber_ref01_match_dt0: any = {}
    subscriber_ref01_match_dt0.id = subscriber_ref01_data.id
    const subscriber_ref01_data_dt0 = await subscriber_ref01_ent.load(subscriber_ref01_match_dt0)
    assert(subscriber_ref01_data_dt0.id === subscriber_ref01_data.id)


    // REMOVE
    const subscriber_ref01_match_rm0: any = { id: subscriber_ref01_data.id }
    await subscriber_ref01_ent.remove(subscriber_ref01_match_rm0)
  

    // LIST
    const subscriber_ref01_match_rt0: any = {}
    subscriber_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const subscriber_ref01_list_rt0 = await subscriber_ref01_ent.list(subscriber_ref01_match_rt0)

    assert(isempty(select(subscriber_ref01_list_rt0, { id: subscriber_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/subscriber/SubscriberTestData.json')

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
    ['subscriber01','subscriber02','subscriber03','page01','page02','page03','page01','page02','page03','incident01','incident02','incident03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  // Detect whether the user provided a real ENTID JSON via env var. The
  // basic flow consumes synthetic IDs from the fixture file; without an
  // override those synthetic IDs reach the live API and 4xx. Surface this
  // to the test so it can skip rather than fail.
  const idmapEnvVal = process.env['STATUSPAGE_TEST_SUBSCRIBER_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'STATUSPAGE_TEST_SUBSCRIBER_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_SUBSCRIBER_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  if (live) {
    client = new StatuspageSDK(merge([
      {
        apikey: env.STATUSPAGE_APIKEY,
      },
      extra
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
    syntheticOnly: live && !idmapOverridden,
    now: Date.now(),
  }

  return setup
}
  
