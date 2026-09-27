

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


loadEnvLocal(__dirname + '/../../../.env.local')


describe('UserEntity', async () => {

  // Per-test live pacing. Delay is read from sdk-test-control.json's
  // `test.live.delayMs`; only sleeps when STATUSPAGE_TEST_LIVE=TRUE.
  afterEach(liveDelay('STATUSPAGE_TEST_LIVE'))

  test('instance', async () => {
    const testsdk = StatuspageSDK.test()
    const ent = testsdk.User()
    assert(null != ent)
  })


  test('basic', async (t) => {

    const live = 'TRUE' === process.env.STATUSPAGE_TEST_LIVE
    for (const op of ['create', 'list', 'remove']) {
      if (!live && maybeSkipControl(t, 'entityOp', 'user.' + op, live)) return
    }

    
    const setup = basicSetup()
    if (setup.live) {
      return runLiveEntity(setup, {"active":true,"alias":{"field":{}},"fields":{"created_at":{"a":true,"fo":"date-time","h":"Created At","n":"created_at","r":false,"t":"`$STRING`","key$":"created_at","index$":0},"email":{"a":true,"h":"Email","n":"email","r":false,"sh":"Email address for the team member","t":"`$STRING`","key$":"email","index$":1},"first_name":{"a":true,"h":"First Name","n":"first_name","r":false,"t":"`$STRING`","key$":"first_name","index$":2},"id":{"a":true,"h":"Id","n":"id","r":false,"sh":"User identifier","t":"`$STRING`","key$":"id","index$":3},"last_name":{"a":true,"h":"Last Name","n":"last_name","r":false,"t":"`$STRING`","key$":"last_name","index$":4},"organization_id":{"a":true,"h":"Organization Id","n":"organization_id","r":false,"sh":"Organization identifier","t":"`$STRING`","key$":"organization_id","index$":5},"updated_at":{"a":true,"fo":"date-time","h":"Updated At","n":"updated_at","r":false,"t":"`$STRING`","key$":"updated_at","index$":6},"user":{"a":true,"h":"User","n":"user","r":true,"t":"`$OBJECT`","key$":"user","index$":7}},"id":{"field":"id","name":"id"},"name":"user","op":{"create":{"input":"data","name":"create","points":[{"a":true,"co":{"id":"POST /organizations/{organization_id}/users","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"organization_id","or":"organization_id","r":true,"t":"`$STRING`","index$":0}]},"k":"http","m":"POST","o":"/organizations/{organization_id}/users","q":{"exist":["organization_id"]},"r":{},"s":[{"lit":"organizations"},{"var":"organization_id"},{"lit":"users"}],"t":{"req":{"user":"`reqdata`"},"res":"`body`"},"index$":0}],"key$":"create"},"list":{"input":"data","name":"list","points":[{"a":true,"co":{"id":"GET /organizations/{organization_id}/users","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"organization_id","or":"organization_id","r":true,"t":"`$STRING`","index$":0}],"query":[{"a":true,"k":"query","n":"page","or":"page","r":false,"t":"`$INTEGER`","index$":0},{"a":true,"k":"query","n":"per_page","or":"per_page","r":false,"t":"`$INTEGER`","index$":1}]},"k":"http","m":"GET","o":"/organizations/{organization_id}/users","q":{"exist":["organization_id","page","per_page"]},"r":{},"s":[{"lit":"organizations"},{"var":"organization_id"},{"lit":"users"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"list"},"remove":{"input":"data","name":"remove","points":[{"a":true,"co":{"id":"DELETE /organizations/{organization_id}/users/{user_id}","source":"openapi3","version":2},"g":{"params":[{"a":true,"k":"param","n":"id","or":"user_id","r":true,"t":"`$STRING`","index$":0},{"a":true,"k":"param","n":"organization_id","or":"organization_id","r":true,"t":"`$STRING`","index$":1}]},"k":"http","m":"DELETE","o":"/organizations/{organization_id}/users/{user_id}","q":{"exist":["id","organization_id"]},"r":{"param":{"user_id":"id"}},"s":[{"lit":"organizations"},{"var":"organization_id"},{"lit":"users"},{"var":"id"}],"t":{"req":"`reqdata`","res":"`body`"},"index$":0}],"key$":"remove"}},"relations":{"ancestors":[]},"key$":"user","name__orig":"user","Name":"User","name_":"user","name-":"user","NAME":"USER","index$":16}, {"active":true,"entity":"user","key$":"BasicUserFlow","kind":"basic","name":"BasicUserFlow","param":{},"step":[{"a":true,"d":{},"i":{"ref":"user_ref01"},"m":{"organization_id":"organization01"},"o":"create","s":[],"v":[],"index$":0},{"a":true,"d":{},"i":{},"m":{"organization_id":"organization01"},"o":"list","s":[],"v":[{"apply":"ItemExists","def":{"ref":"user_ref01"}}],"index$":1},{"a":true,"d":{},"i":{"ref":"user_ref01","suffix":"_rm0"},"m":{"id":"user01","organization_id":"organization01"},"o":"remove","s":[],"v":[],"index$":2},{"a":true,"d":{},"i":{"suffix":"_rt0"},"m":{"organization_id":"organization01"},"o":"list","s":[],"v":[{"apply":"ItemNotExists","def":{"ref":"user_ref01"}}],"index$":3}]}, 'User', {"POST /organizations/{organization_id}/users":{"protocol":"http","requestBody":{"content":{"application/json":{"schema":{"type":"object","properties":{"user":{"type":"object","properties":{"email":{"type":"string","description":"Email address for the team member"},"password":{"type":"string","description":"Password the team member uses to access the site"},"first_name":{"type":"string"},"last_name":{"type":"string"}},"key$":"user"}},"required":["user"],"description":"Create a user","x-ref":"#/components/schemas/postOrganizationsOrganizationIdUsers","index$":1}}},"required":true},"parameters":[{"in":"path","name":"organization_id","description":"Organization Identifier","required":true,"schema":{"type":"string"},"index$":0}]},"GET /organizations/{organization_id}/users":{"protocol":"http","parameters":[{"in":"path","name":"organization_id","description":"Organization Identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"query","name":"page","description":"Page offset to fetch. Beginning February 28, 2023, this endpoint will return paginated data even if this query parameter is not provided.","required":false,"schema":{"type":"integer","format":"int32"},"index$":1},{"in":"query","name":"per_page","description":"Number of results to return per page. Beginning February 28, 2023, a default and maximum limit of 100 will be imposed and this endpoint will return paginated data even if this query parameter is not provided.","required":false,"schema":{"type":"integer","format":"int32"},"index$":2}]},"DELETE /organizations/{organization_id}/users/{user_id}":{"protocol":"http","parameters":[{"in":"path","name":"organization_id","description":"Organization Identifier","required":true,"schema":{"type":"string"},"index$":0},{"in":"path","name":"user_id","description":"User Identifier","required":true,"schema":{"type":"string"},"index$":1}]}})
    }
    const client = setup.client
    const struct = setup.struct

    const isempty = struct.isempty
    const select = struct.select


    // CREATE
    const user_ref01_ent = client.User()
    let user_ref01_data = setup.data.new.user['user_ref01']
    user_ref01_data['organization_id'] = setup.idmap['organization01']

    user_ref01_data = (await user_ref01_ent.create(user_ref01_data)).data()
    assert(null != user_ref01_data.id)


    // LIST
    const user_ref01_match: any = {}
    user_ref01_match['organization_id'] = setup.idmap['organization01']

    const user_ref01_list = (await user_ref01_ent.list(user_ref01_match)).map((e: any) => e.data())

    assert(!isempty(select(user_ref01_list, { id: user_ref01_data.id })))


    // REMOVE
    const user_ref01_match_rm0: any = { id: user_ref01_data.id }
    await user_ref01_ent.remove(user_ref01_match_rm0)
  

    // LIST
    const user_ref01_match_rt0: any = {}
    user_ref01_match_rt0['organization_id'] = setup.idmap['organization01']

    const user_ref01_list_rt0 = (await user_ref01_ent.list(user_ref01_match_rt0)).map((e: any) => e.data())

    assert(isempty(select(user_ref01_list_rt0, { id: user_ref01_data.id })))


  })
})



function basicSetup(extra?: any) {
  // TODO: fix test def options
  const options: any = {} // null

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
    ['user01','user02','user03','organization01'],
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
    'STATUSPAGE_APIKEY': '',
  })

  idmap = env['STATUSPAGE_TEST_USER_ENTID']

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  const transport = createLiveTransport()
  if (live) {
    const rawIds = process.env['STATUSPAGE_TEST_USER_ENTID']
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
  
