
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


describe('PageEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.Page()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let page_ref01_data = Object.values(setup.data.existing.page)[0]

    // LIST
    const page_ref01_ent = client.Page()
    const page_ref01_match = {}

    const page_ref01_list = await page_ref01_ent.list(page_ref01_match)


    // UPDATE
    const page_ref01_data_up0 = {}
    page_ref01_data_up0.id = page_ref01_data.id

    const page_ref01_markdef_up0 = { name: 'branding', value: 'Mark01-page_ref01_' + setup.now }
    page_ref01_data_up0 [page_ref01_markdef_up0.name] = page_ref01_markdef_up0.value

    const page_ref01_resdata_up0 = await page_ref01_ent.update(page_ref01_data_up0)
    assert(page_ref01_resdata_up0.id === page_ref01_data_up0.id)

    assert(page_ref01_resdata_up0[page_ref01_markdef_up0.name] === page_ref01_markdef_up0.value)


    // LOAD
    const page_ref01_match_dt0 = {}
    page_ref01_match_dt0.id = page_ref01_data.id
    const page_ref01_data_dt0 = await page_ref01_ent.load(page_ref01_match_dt0)
    assert(page_ref01_data_dt0.id === page_ref01_data.id)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/page/PageTestData.json')

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
    ['page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_PAGE_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_PAGE_ENTID']

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
  
