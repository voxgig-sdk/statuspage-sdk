
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { StatuspageSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = StatuspageSDK.test()
    equal(testsdk instanceof StatuspageSDK, true,
      'StatuspageSDK.test() must return a client synchronously')
  })

})
