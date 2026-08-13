

class StatuspageError extends Error {

  isStatuspageError = true

  sdk = 'Statuspage'

  constructor(code, msg, ctx) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

module.exports = {
  StatuspageError
}

