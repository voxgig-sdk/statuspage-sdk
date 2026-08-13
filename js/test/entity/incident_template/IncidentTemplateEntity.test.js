
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


describe('IncidentTemplateEntity', async () => {

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.IncidentTemplate()
    assert(null != ent)
  })


  test('basic', async () => {

    const setup = basicSetup()
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const incident_template_ref01_ent = client.IncidentTemplate()
    let incident_template_ref01_data = setup.data.new.incident_template['incident_template_ref01']
    incident_template_ref01_data['page_id'] = setup.idmap['page01']

    incident_template_ref01_data = await incident_template_ref01_ent.create(incident_template_ref01_data)
    assert(null != incident_template_ref01_data.id)


    // LIST
    const incident_template_ref01_match = {}
    incident_template_ref01_match['page_id'] = setup.idmap['page01']

    const incident_template_ref01_list = await incident_template_ref01_ent.list(incident_template_ref01_match)

    assert(!isempty(select(incident_template_ref01_list, { id: incident_template_ref01_data.id })))


  })
})



function basicSetup(extra) {
  // TODO: fix test def options
  const options = {} // null

  // TODO: needs test utility to resolve path
  const entityDataFile =
    Path.resolve(__dirname,
      '../../../../.sdk/test/entity/incident_template/IncidentTemplateTestData.json')

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
    ['incident_template01','incident_template02','incident_template03','page01','page02','page03'],
    {
      '`$PACK`': ['', {
        '`$KEY`': '`$COPY`',
        '`$VAL`': ['`$FORMAT`', 'upper', '`$COPY`']
      }]
    })

  const env = envOverride({
    'STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID': idmap,
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_TEST_EXPLAIN': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  idmap = env['STATUSPAGE_TEST_INCIDENT_TEMPLATE_ENTID']

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
  
