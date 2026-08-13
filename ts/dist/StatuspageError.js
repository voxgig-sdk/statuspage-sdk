"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.StatuspageError = void 0;
class StatuspageError extends Error {
    isStatuspageError = true;
    sdk = 'Statuspage';
    code;
    ctx;
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.StatuspageError = StatuspageError;
//# sourceMappingURL=StatuspageError.js.map