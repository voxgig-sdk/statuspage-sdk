
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { StatuspageSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await StatuspageSDK.test()
    equal(null !== testsdk, true)
  })

})
