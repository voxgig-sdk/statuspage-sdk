
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


describe('MetricsProviderEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.MetricsProvider()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['create', 'list', 'update', 'load', 'remove']) {
      if (maybeSkipControl(t, 'entityOp', 'metrics_provider.' + op, live)) return
    }

    const setup = basicSetup()
    // The basic flow consumes synthetic IDs and field values from the
    // fixture (entity TestData.json). Those don't exist on the live API.
    // Skip live runs unless the user provided a real ENTID env override.
    if (setup.syntheticOnly) {
      t.skip('live entity test uses synthetic IDs from fixture — set STATUSPAGE_TEST_METRICS_PROVIDER_ENTID JSON to run live')
      return
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const metrics_provider_ref01_ent = client.MetricsProvider()
    let metrics_provider_ref01_data = setup.data.new.metrics_provider['metrics_provider_ref01']
    metrics_provider_ref01_data['page_id'] = setup.idmap['page01']

    metrics_provider_ref01_data = await metrics_provider_ref01_ent.create(metrics_provider_ref01_data)
    assert(null != metrics_provider_ref01_data.id)


    // LIST
    const metrics_provider_ref01_match: any = {}
    metrics_provider_ref01_match['page_id'] = setup.idmap['page01']

    const metrics_provider_ref01_list = await metrics_provider_ref01_ent.list(metrics_provider_ref01_match)

    assert(!isempty(select(metrics_provider_ref01_list, { id: metrics_provider_ref01_data.id })))


    // UPDATE
    const metrics_provider_ref01_data_up0: any = {}
    metrics_provider_ref01_data_up0.id = metrics_provider_ref01_data.id
    metrics_provider_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const metrics_provider_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-metrics_provider_ref01_' + setup.now }
    ;(metrics_provider_ref01_data_up0 as any)[metrics_provider_ref01_markdef_up0.name] = metrics_provider_ref01_markdef_up0.value

    const metrics_provider_ref01_resdata_up0 = await metrics_provider_ref01_ent.update(metrics_provider_ref01_data_up0)
    assert(metrics_provider_ref01_resdata_up0.id === metrics_provider_ref01_data_up0.id)

    assert((metrics_provider_ref01_resdata_up0 as any)[metrics_provider_ref01_markdef_up0.name] === metrics_provider_ref01_markdef_up0.value)


    // LOAD
    const metrics_provider_ref01_match_dt0: any = {}
    metrics_provider_ref01_match_dt0.id = metrics_provider_ref01_data.id
    const metrics_provider_ref01_data_dt0 = await metrics_provider_ref01_ent.load(metrics_provider_ref01_match_dt0)
    assert(metrics_provider_ref01_data_dt0.id === metrics_provider_ref01_data.id)


    // REMOVE
    const metrics_provider_ref01_match_rm0: any = { id: metrics_provider_ref01_data.id }
    await metrics_provider_ref01_ent.remove(metrics_provider_ref01_match_rm0)
  

    // LIST
    const metrics_provider_ref01_match_rt0: any = {}
    metrics_provider_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const metrics_provider_ref01_list_rt0 = await metrics_provider_ref01_ent.list(metrics_provider_ref01_match_rt0)

    assert(isempty(select(metrics_provider_ref01_list_rt0, { id: metrics_provider_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname, 
      '../../../../.sdk/test/entity/metrics_provider/MetricsProviderTestData.json')

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
    ['metrics_provider01','metrics_provider02','metrics_provider03','page01','page02','page03'],
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
  const idmapEnvVal = process.env['STATUSPAGE_TEST_METRICS_PROVIDER_ENTID']
  const idmapOverridden = null != idmapEnvVal && idmapEnvVal.trim().startsWith('{')

  const env = envOverride({
    'STATUSPAGE_TEST_METRICS_PROVIDER_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_METRICS_PROVIDER_ENTID']

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
  
