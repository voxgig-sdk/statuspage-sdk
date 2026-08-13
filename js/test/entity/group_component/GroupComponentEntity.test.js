
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


describe('GroupComponentEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.GroupComponent()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const group_component_ref01_ent = client.GroupComponent()
    let group_component_ref01_data = setup.data.new.group_component['group_component_ref01']
    group_component_ref01_data['page_id'] = setup.idmap['page01']

    group_component_ref01_data = await group_component_ref01_ent.create(group_component_ref01_data)
    assert(null != group_component_ref01_data.id)


    // LIST
    const group_component_ref01_match = {}
    group_component_ref01_match['page_id'] = setup.idmap['page01']

    const group_component_ref01_list = await group_component_ref01_ent.list(group_component_ref01_match)

    assert(!isempty(select(group_component_ref01_list, { id: group_component_ref01_data.id })))


    // UPDATE
    const group_component_ref01_data_up0 = {}
    group_component_ref01_data_up0.id = group_component_ref01_data.id
    group_component_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const group_component_ref01_markdef_up0 = { name: 'components', value: 'Mark01-group_component_ref01_' + setup.now }
    group_component_ref01_data_up0 [group_component_ref01_markdef_up0.name] = group_component_ref01_markdef_up0.value

    const group_component_ref01_resdata_up0 = await group_component_ref01_ent.update(group_component_ref01_data_up0)
    assert(group_component_ref01_resdata_up0.id === group_component_ref01_data_up0.id)

    assert(group_component_ref01_resdata_up0[group_component_ref01_markdef_up0.name] === group_component_ref01_markdef_up0.value)


    // LOAD
    const group_component_ref01_match_dt0 = {}
    group_component_ref01_match_dt0.id = group_component_ref01_data.id
    const group_component_ref01_data_dt0 = await group_component_ref01_ent.load(group_component_ref01_match_dt0)
    assert(group_component_ref01_data_dt0.id === group_component_ref01_data.id)


    // REMOVE
    const group_component_ref01_match_rm0 = {}
    group_component_ref01_match_rm0.id = group_component_ref01_data.id
    await group_component_ref01_ent.remove(group_component_ref01_match_rm0)
  

    // LIST
    const group_component_ref01_match_rt0 = {}
    group_component_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const group_component_ref01_list_rt0 = await group_component_ref01_ent.list(group_component_ref01_match_rt0)

    assert(isempty(select(group_component_ref01_list_rt0, { id: group_component_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/group_component/GroupComponentTestData.json')

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
    ['group_component01','group_component02','group_component03','page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_GROUP_COMPONENT_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_GROUP_COMPONENT_ENTID']

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
  
