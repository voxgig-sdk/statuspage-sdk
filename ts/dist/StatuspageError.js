"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.StatuspageError = void 0;
class StatuspageError extends Error {
    isStatuspageError = true;
    sdk = 'Statuspage';
    code;
    ctx;
    status = -1;
    // `err.notFound` rather than a magic number at every call site.
    get notFound() { return 404 === this.status; }
    constructor(code, msg, ctx) {
        super(msg);
        this.code = code;
        this.ctx = ctx;
    }
}
exports.StatuspageError = StatuspageError;
//# sourceMappingURL=StatuspageError.js.map