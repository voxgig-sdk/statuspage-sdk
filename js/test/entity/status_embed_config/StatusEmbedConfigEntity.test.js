
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


describe('StatusEmbedConfigEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.StatusEmbedConfig()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let status_embed_config_ref01_data = Object.values(setup.data.existing.status_embed_config)[0]

    // UPDATE
    const status_embed_config_ref01_ent = client.StatusEmbedConfig()
    const status_embed_config_ref01_data_up0 = {}

    const status_embed_config_ref01_markdef_up0 = { name: 'incident_background_color', value: 'Mark01-status_embed_config_ref01_' + setup.now }
    status_embed_config_ref01_data_up0 [status_embed_config_ref01_markdef_up0.name] = status_embed_config_ref01_markdef_up0.value

    const status_embed_config_ref01_resdata_up0 = await status_embed_config_ref01_ent.update(status_embed_config_ref01_data_up0)
    assert(null != status_embed_config_ref01_resdata_up0)

    assert(status_embed_config_ref01_resdata_up0[status_embed_config_ref01_markdef_up0.name] === status_embed_config_ref01_markdef_up0.value)


    // LOAD
    const status_embed_config_ref01_match_dt0 = {}
    const status_embed_config_ref01_data_dt0 = await status_embed_config_ref01_ent.load(status_embed_config_ref01_match_dt0)
    assert(null != status_embed_config_ref01_data_dt0)


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/status_embed_config/StatusEmbedConfigTestData.json')

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
    ['status_embed_config01','status_embed_config02','status_embed_config03','page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID']

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
  
