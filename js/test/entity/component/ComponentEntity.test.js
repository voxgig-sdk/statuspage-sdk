
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


describe('ComponentEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.Component()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const component_ref01_ent = client.Component()
    let component_ref01_data = setup.data.new.component['component_ref01']
    component_ref01_data['page_access_group_id'] = setup.idmap['page_access_group01']
    component_ref01_data['page_access_user_id'] = setup.idmap['page_access_user01']
    component_ref01_data['page_id'] = setup.idmap['page01']

    component_ref01_data = await component_ref01_ent.create(component_ref01_data)
    assert(null != component_ref01_data.id)


    // LIST
    const component_ref01_match = {}
    component_ref01_match['page_id'] = setup.idmap['page01']

    const component_ref01_list = await component_ref01_ent.list(component_ref01_match)

    assert(!isempty(select(component_ref01_list, { id: component_ref01_data.id })))


    // UPDATE
    const component_ref01_data_up0 = {}
    component_ref01_data_up0.id = component_ref01_data.id
    component_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const component_ref01_markdef_up0 = { name: 'automation_email', value: 'Mark01-component_ref01_' + setup.now }
    component_ref01_data_up0 [component_ref01_markdef_up0.name] = component_ref01_markdef_up0.value

    const component_ref01_resdata_up0 = await component_ref01_ent.update(component_ref01_data_up0)
    assert(component_ref01_resdata_up0.id === component_ref01_data_up0.id)

    assert(component_ref01_resdata_up0[component_ref01_markdef_up0.name] === component_ref01_markdef_up0.value)


    // LOAD
    const component_ref01_match_dt0 = {}
    component_ref01_match_dt0.id = component_ref01_data.id
    const component_ref01_data_dt0 = await component_ref01_ent.load(component_ref01_match_dt0)
    assert(component_ref01_data_dt0.id === component_ref01_data.id)


    // REMOVE
    const component_ref01_match_rm0 = {}
    component_ref01_match_rm0.id = component_ref01_data.id
    await component_ref01_ent.remove(component_ref01_match_rm0)
  

    // LIST
    const component_ref01_match_rt0 = {}
    component_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const component_ref01_list_rt0 = await component_ref01_ent.list(component_ref01_match_rt0)

    assert(isempty(select(component_ref01_list_rt0, { id: component_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/component/ComponentTestData.json')

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
    ['component01','component02','component03','page01','page02','page03','page01','page02','page03','page_access_group01','page_access_group02','page_access_group03','page01','page02','page03','page_access_user01','page_access_user02','page_access_user03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_COMPONENT_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_COMPONENT_ENTID']

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
  
