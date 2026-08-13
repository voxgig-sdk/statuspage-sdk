
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

const Path = require('node:path')
const Fs = require('node:fs')

const { test, describe } = require('node:test')
const assert = require('node:assert')


const { StatuspageSDK, BaseFeature, stdutil, config } = require('../../..')

const {
  envOverride,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
} = require('../../utility')


describe('MetricsProviderEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.MetricsProvider()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
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
    const metrics_provider_ref01_match = {}
    metrics_provider_ref01_match['page_id'] = setup.idmap['page01']

    const metrics_provider_ref01_list = await metrics_provider_ref01_ent.list(metrics_provider_ref01_match)

    assert(!isempty(select(metrics_provider_ref01_list, { id: metrics_provider_ref01_data.id })))


    // UPDATE
    const metrics_provider_ref01_data_up0 = {}
    metrics_provider_ref01_data_up0.id = metrics_provider_ref01_data.id
    metrics_provider_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const metrics_provider_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-metrics_provider_ref01_' + setup.now }
    metrics_provider_ref01_data_up0 [metrics_provider_ref01_markdef_up0.name] = metrics_provider_ref01_markdef_up0.value

    const metrics_provider_ref01_resdata_up0 = await metrics_provider_ref01_ent.update(metrics_provider_ref01_data_up0)
    assert(metrics_provider_ref01_resdata_up0.id === metrics_provider_ref01_data_up0.id)

    assert(metrics_provider_ref01_resdata_up0[metrics_provider_ref01_markdef_up0.name] === metrics_provider_ref01_markdef_up0.value)


    // LOAD
    const metrics_provider_ref01_match_dt0 = {}
    metrics_provider_ref01_match_dt0.id = metrics_provider_ref01_data.id
    const metrics_provider_ref01_data_dt0 = await metrics_provider_ref01_ent.load(metrics_provider_ref01_match_dt0)
    assert(metrics_provider_ref01_data_dt0.id === metrics_provider_ref01_data.id)


    // REMOVE
    const metrics_provider_ref01_match_rm0 = {}
    metrics_provider_ref01_match_rm0.id = metrics_provider_ref01_data.id
    await metrics_provider_ref01_ent.remove(metrics_provider_ref01_match_rm0)
  

    // LIST
    const metrics_provider_ref01_match_rt0 = {}
    metrics_provider_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const metrics_provider_ref01_list_rt0 = await metrics_provider_ref01_ent.list(metrics_provider_ref01_match_rt0)

    assert(isempty(select(metrics_provider_ref01_list_rt0, { id: metrics_provider_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

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

  const env = envOverride({
    'STATUSPAGE_TEST_METRICS_PROVIDER_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_METRICS_PROVIDER_ENTID']

  if ('TRUE' === env.STATUSPAGE_TEST_LIVE) {
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
    now: Date.now(),
  }

  return setup
}
  
