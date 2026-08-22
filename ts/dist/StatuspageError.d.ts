import { Context } from './Context';
declare class StatuspageError extends Error {
    isStatuspageError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { StatuspageError };
