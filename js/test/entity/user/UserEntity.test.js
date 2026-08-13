
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


describe('UserEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.User()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const user_ref01_ent = client.User()
    let user_ref01_data = setup.data.new.user['user_ref01']
    user_ref01_data['organization_id'] = setup.idmap['organization01']

    user_ref01_data = await user_ref01_ent.create(user_ref01_data)
    assert(null != user_ref01_data.id)


    // LIST
    const user_ref01_match = {}
    user_ref01_match['organization_id'] = setup.idmap['organization01']

    const user_ref01_list = await user_ref01_ent.list(user_ref01_match)

    assert(!isempty(select(user_ref01_list, { id: user_ref01_data.id })))


    // REMOVE
    const user_ref01_match_rm0 = {}
    user_ref01_match_rm0.id = user_ref01_data.id
    await user_ref01_ent.remove(user_ref01_match_rm0)
  

    // LIST
    const user_ref01_match_rt0 = {}
    user_ref01_match_rt0['organization_id'] = setup.idmap['organization01']

    const user_ref01_list_rt0 = await user_ref01_ent.list(user_ref01_match_rt0)

    assert(isempty(select(user_ref01_list_rt0, { id: user_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/user/UserTestData.json')

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
    ['user01','user02','user03','organization01','organization02','organization03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_USER_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_USER_ENTID']

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
  
