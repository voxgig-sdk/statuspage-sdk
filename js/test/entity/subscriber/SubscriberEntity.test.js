
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


describe('SubscriberEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.Subscriber()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
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
    const subscriber_ref01_match = {}
    subscriber_ref01_match['page_id'] = setup.idmap['page01']

    const subscriber_ref01_list = await subscriber_ref01_ent.list(subscriber_ref01_match)

    assert(!isempty(select(subscriber_ref01_list, { id: subscriber_ref01_data.id })))


    // UPDATE
    const subscriber_ref01_data_up0 = {}
    subscriber_ref01_data_up0.id = subscriber_ref01_data.id
    subscriber_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const subscriber_ref01_markdef_up0 = { name: 'components', value: 'Mark01-subscriber_ref01_' + setup.now }
    subscriber_ref01_data_up0 [subscriber_ref01_markdef_up0.name] = subscriber_ref01_markdef_up0.value

    const subscriber_ref01_resdata_up0 = await subscriber_ref01_ent.update(subscriber_ref01_data_up0)
    assert(subscriber_ref01_resdata_up0.id === subscriber_ref01_data_up0.id)

    assert(subscriber_ref01_resdata_up0[subscriber_ref01_markdef_up0.name] === subscriber_ref01_markdef_up0.value)


    // LOAD
    const subscriber_ref01_match_dt0 = {}
    subscriber_ref01_match_dt0.id = subscriber_ref01_data.id
    const subscriber_ref01_data_dt0 = await subscriber_ref01_ent.load(subscriber_ref01_match_dt0)
    assert(subscriber_ref01_data_dt0.id === subscriber_ref01_data.id)


    // REMOVE
    const subscriber_ref01_match_rm0 = {}
    subscriber_ref01_match_rm0.id = subscriber_ref01_data.id
    await subscriber_ref01_ent.remove(subscriber_ref01_match_rm0)
  

    // LIST
    const subscriber_ref01_match_rt0 = {}
    subscriber_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const subscriber_ref01_list_rt0 = await subscriber_ref01_ent.list(subscriber_ref01_match_rt0)

    assert(isempty(select(subscriber_ref01_list_rt0, { id: subscriber_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

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

  const env = envOverride({
    'STATUSPAGE_TEST_SUBSCRIBER_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_SUBSCRIBER_ENTID']

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
  
