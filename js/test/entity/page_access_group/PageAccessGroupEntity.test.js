
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


describe('PageAccessGroupEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.PageAccessGroup()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const page_access_group_ref01_ent = client.PageAccessGroup()
    let page_access_group_ref01_data = setup.data.new.page_access_group['page_access_group_ref01']
    page_access_group_ref01_data['page_id'] = setup.idmap['page01']

    page_access_group_ref01_data = await page_access_group_ref01_ent.create(page_access_group_ref01_data)
    assert(null != page_access_group_ref01_data.id)


    // LIST
    const page_access_group_ref01_match = {}
    page_access_group_ref01_match['page_id'] = setup.idmap['page01']

    const page_access_group_ref01_list = await page_access_group_ref01_ent.list(page_access_group_ref01_match)

    assert(!isempty(select(page_access_group_ref01_list, { id: page_access_group_ref01_data.id })))


    // UPDATE
    const page_access_group_ref01_data_up0 = {}
    page_access_group_ref01_data_up0.id = page_access_group_ref01_data.id
    page_access_group_ref01_data_up0 ['page_id'] = setup.idmap['page_id']

    const page_access_group_ref01_markdef_up0 = { name: 'created_at', value: 'Mark01-page_access_group_ref01_' + setup.now }
    page_access_group_ref01_data_up0 [page_access_group_ref01_markdef_up0.name] = page_access_group_ref01_markdef_up0.value

    const page_access_group_ref01_resdata_up0 = await page_access_group_ref01_ent.update(page_access_group_ref01_data_up0)
    assert(page_access_group_ref01_resdata_up0.id === page_access_group_ref01_data_up0.id)

    assert(page_access_group_ref01_resdata_up0[page_access_group_ref01_markdef_up0.name] === page_access_group_ref01_markdef_up0.value)


    // LOAD
    const page_access_group_ref01_match_dt0 = {}
    page_access_group_ref01_match_dt0.id = page_access_group_ref01_data.id
    const page_access_group_ref01_data_dt0 = await page_access_group_ref01_ent.load(page_access_group_ref01_match_dt0)
    assert(page_access_group_ref01_data_dt0.id === page_access_group_ref01_data.id)


    // REMOVE
    const page_access_group_ref01_match_rm0 = {}
    page_access_group_ref01_match_rm0.id = page_access_group_ref01_data.id
    await page_access_group_ref01_ent.remove(page_access_group_ref01_match_rm0)
  

    // LIST
    const page_access_group_ref01_match_rt0 = {}
    page_access_group_ref01_match_rt0['page_id'] = setup.idmap['page01']

    const page_access_group_ref01_list_rt0 = await page_access_group_ref01_ent.list(page_access_group_ref01_match_rt0)

    assert(isempty(select(page_access_group_ref01_list_rt0, { id: page_access_group_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/page_access_group/PageAccessGroupTestData.json')

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
    ['page_access_group01','page_access_group02','page_access_group03','page01','page02','page03','page01','page02','page03','component01','component02','component03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_PAGE_ACCESS_GROUP_ENTID']

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
  
