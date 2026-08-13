import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Postmortem, PostmortemLoadMatch, PostmortemUpdateData } from '../StatuspageTypes';
declare class PostmortemEntity extends StatuspageEntityBase<Postmortem> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: PostmortemEntity): PostmortemEntity;
    load(this: any, reqmatch?: PostmortemLoadMatch, ctrl?: Control): Promise<Postmortem>;
    update(this: any, reqdata?: PostmortemUpdateData, ctrl?: Control): Promise<Postmortem>;
}
export { PostmortemEntity };
