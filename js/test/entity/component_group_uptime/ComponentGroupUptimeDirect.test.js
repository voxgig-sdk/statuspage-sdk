
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

const { test, describe } = require('node:test')
const assert = require('node:assert')


const { StatuspageSDK } = require('../../..')

const {
  envOverride,
} = require('../../utility')


describe('ComponentGroupUptimeDirect', async () => {

  test('direct-exists', async () => {
    const sdk = new StatuspageSDK({
      // Concrete base: a live construction must satisfy any server
      // variables a templated base URL declares; overriding base with a
      // literal (as the direct flow tests do) sidesteps the requirement.
      base: 'http://localhost:8080',
      system: { fetch: async () => ({}) }
    })
    assert('function' === typeof sdk.direct)
    assert('function' === typeof sdk.prepare)
  })


  test('direct-load-component_group_uptime', async () => {
    const setup = directSetup({ id: 'direct01' })
    const { client, calls } = setup

    const params = {}
    if (!setup.live) {
      params.id = 'direct01'
      params.page_id = 'direct02'
    }

    const result = await client.direct({
      path: 'pages/{page_id}/component-groups/{id}/uptime',
      method: 'GET',
      params,
    })

    assert(result.ok === true)
    assert(result.status === 200)
    assert(null != result.data)

    if (!setup.live) {
      assert(result.data.id === 'direct01')
      assert(calls.length === 1)
      assert(calls[0].init.method === 'GET')
      assert(calls[0].url.includes('direct01'))
      assert(calls[0].url.includes('direct02'))
    }
  })

})



function directSetup(mockres) {
  const calls = []

  const env = envOverride({
    'STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID': {},
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  if (live) {
    const client = new StatuspageSDK({
      apikey: env.STATUSPAGE_APIKEY,
    })

    let idmap = env['STATUSPAGE_TEST_COMPONENT_GROUP_UPTIME_ENTID']
    if ('string' === typeof idmap && idmap.startsWith('{')) {
      idmap = JSON.parse(idmap)
    }

    return { client, calls, live, idmap }
  }

  const mockFetch = async (url, init) => {
    calls.push({ url, init })
    return {
      status: 200,
      statusText: 'OK',
      headers: {},
      json: async () => (null != mockres ? mockres : { id: 'direct01' }),
    }
  }

  const client = new StatuspageSDK({
    base: 'http://localhost:8080',
    system: { fetch: mockFetch },
  })

  return { client, calls, live, idmap: {} }
}
  
