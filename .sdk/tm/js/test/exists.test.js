
const { test, describe } = require('node:test')
const { equal } = require('node:assert')


const { StatuspageSDK } = require('..')


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await StatuspageSDK.test()
    equal(null !== testsdk, true)
  })

})
