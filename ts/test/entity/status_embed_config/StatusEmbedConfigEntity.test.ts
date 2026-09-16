

import Path from 'node:path'
import * as Fs from 'node:fs'

import { test, describe, afterEach } from 'node:test'
import assert from 'node:assert'
import { createLiveTransport } from '../../live-runner'
import { runLiveEntity } from '../../live-entity'


import { StatuspageSDK, BaseFeature, stdutil } from '../../..'

import {
  envOverride,
  liveClientOptions,
  liveDelay,
  loadEnvLocal,
  makeCtrl,
  makeMatch,
  makeReqdata,
  makeStepData,
  makeValid,
  maybeSkipControl,
} from '../../utility'


// AFTER the imports on purpose: TypeScript hoists `import` above any
// statement in the emitted CommonJS, so a loader placed above them would
// run only after every imported module had already been evaluated - and
// anything reading process.env at module scope would miss these values.
loadEnvLocal(__dirname + '/../../../.env.local')


describe('StatusEmbedConfigEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.StatusEmbedConfig()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['update', 'load']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'status_embed_config.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":[{"active":true,"name":"incident_background_color","req":false,"short":"Color of status embed iframe background when displaying incident","type":"`$STRING`","index$":0},{"active":true,"name":"incident_text_color","req":false,"short":"Color of status embed iframe text when displaying incident","type":"`$STRING`","index$":1},{"active":true,"name":"maintenance_background_color","req":false,"short":"Color of status embed iframe background when displaying maintenance","type":"`$STRING`","index$":2},{"active":true,"name":"maintenance_text_color","req":false,"short":"Color of status embed iframe text when displaying maintenance","type":"`$STRING`","index$":3},{"active":true,"name":"page_id","req":false,"short":"Page identifier","type":"`$STRING`","index$":4},{"active":true,"name":"position","req":false,"short":"Corner where status embed iframe will appear on page","type":"`$STRING`","index$":5},{"active":true,"name":"status_embed_config","req":false,"type":"`$OBJECT`","index$":6}],"name":"status_embed_config","op":{"load":{"input":"data","name":"load","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"page_id","orig":"page_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"GET /pages/{page_id}/status_embed_config","json":"{\"operationId\":\"getPagesPageIdStatusEmbedConfig\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update status embed config settings\",\"properties\":{\"incident_background_color\":{\"description\":\"Color of status embed iframe background when displaying incident\",\"type\":\"string\"},\"incident_text_color\":{\"description\":\"Color of status embed iframe text when displaying incident\",\"type\":\"string\"},\"maintenance_background_color\":{\"description\":\"Color of status embed iframe background when displaying maintenance\",\"type\":\"string\"},\"maintenance_text_color\":{\"description\":\"Color of status embed iframe text when displaying maintenance\",\"type\":\"string\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"},\"position\":{\"description\":\"Corner where status embed iframe will appear on page\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Get status embed config settings\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"403\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"You are not authorized to access this resource.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"GET","orig":"/pages/{page_id}/status_embed_config","segments":[{"lit":"pages"},{"var":"page_id"},{"lit":"status_embed_config"}],"select":{"exist":["page_id"]},"transform":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"load"},"patch":{"input":"data","name":"patch","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"page_id","orig":"page_id","reqd":true,"type":"`$STRING`"}]},"contract":{"id":"PATCH /pages/{page_id}/status_embed_config","json":"{\"operationId\":\"patchPagesPageIdStatusEmbedConfig\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update status embed config settings\",\"properties\":{\"status_embed_config\":{\"properties\":{\"incident_background_color\":{\"description\":\"Color of status embed iframe background when displaying incident\",\"type\":\"string\"},\"incident_text_color\":{\"description\":\"Color of status embed iframe text when displaying incident\",\"type\":\"string\"},\"maintenance_background_color\":{\"description\":\"Color of status embed iframe background when displaying maintenance\",\"type\":\"string\"},\"maintenance_text_color\":{\"description\":\"Color of status embed iframe text when displaying maintenance\",\"type\":\"string\"},\"position\":{\"description\":\"Corner where status embed iframe will appear on page\",\"type\":\"string\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update status embed config settings\",\"properties\":{\"incident_background_color\":{\"description\":\"Color of status embed iframe background when displaying incident\",\"type\":\"string\"},\"incident_text_color\":{\"description\":\"Color of status embed iframe text when displaying incident\",\"type\":\"string\"},\"maintenance_background_color\":{\"description\":\"Color of status embed iframe background when displaying maintenance\",\"type\":\"string\"},\"maintenance_text_color\":{\"description\":\"Color of status embed iframe text when displaying maintenance\",\"type\":\"string\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"},\"position\":{\"description\":\"Corner where status embed iframe will appear on page\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Update status embed config settings\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"403\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"You are not authorized to access this resource.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"},\"422\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Unprocessable entity\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"PATCH","orig":"/pages/{page_id}/status_embed_config","segments":[{"lit":"pages"},{"var":"page_id"},{"lit":"status_embed_config"}],"select":{"exist":["page_id"]},"transform":{"req":{"status_embed_config":"`reqdata`"},"res":"`body`"},"index$":0}],"key$":"patch"},"update":{"input":"data","name":"update","points":[{"active":true,"args":{"params":[{"active":true,"kind":"param","name":"page_id","orig":"page_id","reqd":true,"type":"`$STRING`","index$":0}]},"contract":{"id":"PUT /pages/{page_id}/status_embed_config","json":"{\"operationId\":\"putPagesPageIdStatusEmbedConfig\",\"parameters\":[{\"description\":\"Page identifier\",\"in\":\"path\",\"name\":\"page_id\",\"required\":true,\"schema\":{\"type\":\"string\"}}],\"protocol\":\"http\",\"requestBody\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update status embed config settings\",\"properties\":{\"status_embed_config\":{\"properties\":{\"incident_background_color\":{\"description\":\"Color of status embed iframe background when displaying incident\",\"type\":\"string\"},\"incident_text_color\":{\"description\":\"Color of status embed iframe text when displaying incident\",\"type\":\"string\"},\"maintenance_background_color\":{\"description\":\"Color of status embed iframe background when displaying maintenance\",\"type\":\"string\"},\"maintenance_text_color\":{\"description\":\"Color of status embed iframe text when displaying maintenance\",\"type\":\"string\"},\"position\":{\"description\":\"Corner where status embed iframe will appear on page\",\"type\":\"string\"}},\"type\":\"object\"}},\"type\":\"object\"}}},\"required\":true},\"responses\":{\"200\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Update status embed config settings\",\"properties\":{\"incident_background_color\":{\"description\":\"Color of status embed iframe background when displaying incident\",\"type\":\"string\"},\"incident_text_color\":{\"description\":\"Color of status embed iframe text when displaying incident\",\"type\":\"string\"},\"maintenance_background_color\":{\"description\":\"Color of status embed iframe background when displaying maintenance\",\"type\":\"string\"},\"maintenance_text_color\":{\"description\":\"Color of status embed iframe text when displaying maintenance\",\"type\":\"string\"},\"page_id\":{\"description\":\"Page identifier\",\"type\":\"string\"},\"position\":{\"description\":\"Corner where status embed iframe will appear on page\",\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Update status embed config settings\"},\"400\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Bad request\"},\"401\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Could not authenticate\"},\"403\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"You are not authorized to access this resource.\"},\"404\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"The requested resource could not be found.\"},\"422\":{\"content\":{\"application/json\":{\"schema\":{\"description\":\"Get a list of users\",\"properties\":{\"message\":{\"type\":\"string\"}},\"type\":\"object\"}}},\"description\":\"Unprocessable entity\"}},\"security\":[{\"api_key\":[]}],\"securitySchemes\":{\"api_key\":{\"description\":\"#### Obtaining your API Key\\n\\nAuthentication is done via an API token provided in the Statuspage management interface.\\n\\n  1. Log in to your account at https://manage.statuspage.io/login.\\n  2. Click on your avatar in the bottom left of your screen to access the user menu.\\n  3. Click **API info**.\\n\\n### Passing your API key in an authorization header\\n\\nThe following example authenticates you with the Statuspage API.  Along with the Page ID\\nlisted on the API page, we can fetch your page profile.\\n\\n    curl -H \\\"Authorization: OAuth 89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\" \\\\\\n      https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json\\n\\n### Passing your API key in a query param\\n\\n    curl \\\"https://api.statuspage.io/v1/pages/gytm4qzbx9t6.json?api_key=89a229ce1a8dbcf9ff30430fbe35eb4c0426574bca932061892cefd2138aa4b1\\\"\\n\",\"in\":\"header\",\"name\":\"Authorization\",\"type\":\"apiKey\"}},\"securitySource\":\"definition\"}","source":"openapi3","version":1},"kind":"http","method":"PUT","orig":"/pages/{page_id}/status_embed_config","segments":[{"lit":"pages"},{"var":"page_id"},{"lit":"status_embed_config"}],"select":{"exist":["page_id"]},"transform":{"req":{"status_embed_config":"`reqdata`"},"res":"`body`"},"index$":0}],"key$":"update"}},"relations":{"ancestors":[["page"]]},"key$":"status_embed_config","name__orig":"status_embed_config","Name":"StatusEmbedConfig","name_":"status_embed_config","name-":"status-embed-config","NAME":"STATUS_EMBED_CONFIG","index$":15}, {"active":true,"entity":"status_embed_config","key$":"BasicStatusEmbedConfigFlow","kind":"basic","name":"BasicStatusEmbedConfigFlow","param":{},"step":[{"active":true,"data":{},"input":{"ref":"status_embed_config_ref01","srcdatavar":"status_embed_config_ref01_data","suffix":"_up0","textfield":"incident_background_color"},"match":{},"op":"update","spec":[{"apply":"TextFieldMark","def":{"mark":"Mark01-status_embed_config_ref01"}}],"valid":[],"index$":0},{"active":true,"data":{},"input":{"ref":"status_embed_config_ref01","srcdatavar":"status_embed_config_ref01_data","suffix":"_dt0"},"match":{"id":"status_embed_config01"},"op":"load","spec":[],"valid":[{"apply":"TextFieldMark","def":{"mark":"Mark01-status_embed_config_ref01"}}],"index$":1}]}, 'StatusEmbedConfig')
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select

    let status_embed_config_ref01_data = Object.values(setup.data.existing.status_embed_config)[0] as any

    // UPDATE
    const status_embed_config_ref01_ent = client.StatusEmbedConfig()
    const status_embed_config_ref01_data_up0: any = {}

    const status_embed_config_ref01_markdef_up0 = { name: 'incident_background_color', value: 'Mark01-status_embed_config_ref01_' + setup.now }
    ;(status_embed_config_ref01_data_up0 as any)[status_embed_config_ref01_markdef_up0.name] = status_embed_config_ref01_markdef_up0.value

    const status_embed_config_ref01_resdata_up0 = (await status_embed_config_ref01_ent.update(status_embed_config_ref01_data_up0)).data()
    assert(null != status_embed_config_ref01_resdata_up0)

    assert((status_embed_config_ref01_resdata_up0 as any)[status_embed_config_ref01_markdef_up0.name] === status_embed_config_ref01_markdef_up0.value)



  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

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
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_STATUS_EMBED_CONFIG_ENTID']
    idmap = rawIds && rawIds.trim() ? JSON.parse(rawIds) : {}
    if (!idmap || Array.isArray(idmap) || typeof idmap !== 'object') {
      throw new Error('Live ENTID must be a JSON object')
    }
    client = new StatuspageSDK(merge([
      // FIRST, so the generated fields below win: sdk-test-control.json's
      // test.client.options adds to the live client, it does not redirect it.
      liveClientOptions(),
      {
        apikey: env.STATUSPAGE_APIKEY,
      },
      // 'extra || {}', not a bare 'extra': struct.merge returns UNDEFINED when the
      // last entry is undefined, and basicSetup is normally called with no
      // argument at all - so a bare 'extra' silently discarded the apikey
      // and server values above and handed the SDK undefined. Harmless
      // while there was nothing in that object; not harmless now.
      extra || {},
      { system: { fetch: transport.fetch } }
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
    transport,
    now: Date.now(),
  }

  return setup
}
  
