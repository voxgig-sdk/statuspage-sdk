
import { Context } from './Context'


class StatuspageError extends Error {

  isStatuspageError = true

  sdk = 'Statuspage'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  StatuspageError
}

