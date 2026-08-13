
const envlocal = __dirname + '/../../../.env.local'
require('dotenv').config({ quiet: true, path: [envlocal] })

const { test, describe } = require('node:test')
const assert = require('node:assert')


const { StatuspageSDK } = require('../../..')

const {
  envOverride,
} = require('../../utility')


describe('ComponentDirect', async () => {

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


  test('direct-load-component', async () => {
    const setup = directSetup({ id: 'direct01' })
    const { client, calls } = setup

    const params = {}
    if (setup.live) {
      const listResult = await client.direct({
        path: 'pages/{page_id}/page_access_groups/{page_access_group_id}/components',
        method: 'GET',
        params: {
        page_access_group_id: setup.idmap['page_access_group01'],
        page_id: setup.idmap['page01'],
        },
      })
      assert(listResult.ok === true)
      const listData = listResult.data
      if (!Array.isArray(listData) || listData.length === 0) {
        return // skip: no entities to load in live mode
      }
      params.id = listData[0].id
      params.page_id = setup.idmap['page01']
    } else {
      params.id = 'direct01'
      params.page_id = 'direct02'
    }

    const result = await client.direct({
      path: 'pages/{page_id}/components/{id}/uptime',
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

  test('direct-list-component', async () => {
    const setup = directSetup([{ id: 'direct01' }, { id: 'direct02' }])
    const { client, calls } = setup

    const params = {}
    if (setup.live) {
      params.page_access_group_id = setup.idmap['page_access_group01']
      params.page_id = setup.idmap['page01']
    } else {
      params.page_access_group_id = 'direct01'
      params.page_id = 'direct02'
    }

    const result = await client.direct({
      path: 'pages/{page_id}/page_access_groups/{page_access_group_id}/components',
      method: 'GET',
      params,
    })

    assert(result.ok === true)
    assert(result.status === 200)
    assert(Array.isArray(result.data))

    if (!setup.live) {
      assert(result.data.length === 2)
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
    'STATUSPAGE_TEST_COMPONENT_ENTID': {},
    'STATUSPAGE_TEST_LIVE': 'FALSE',
    'STATUSPAGE_APIKEY': 'NONE',
  })

  const live = 'TRUE' === env.STATUSPAGE_TEST_LIVE

  if (live) {
    const client = new StatuspageSDK({
      apikey: env.STATUSPAGE_APIKEY,
    })

    let idmap = env['STATUSPAGE_TEST_COMPONENT_ENTID']
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
  
