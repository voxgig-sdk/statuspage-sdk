import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Postmortem, PostmortemLoadMatch, PostmortemUpdateData } from '../StatuspageTypes';
declare class PostmortemEntity extends StatuspageEntityBase<Postmortem> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: PostmortemEntity): PostmortemEntity;
    load(this: any, reqmatch?: PostmortemLoadMatch, ctrl?: Control): Promise<PostmortemEntity>;
    update(this: any, reqdata?: PostmortemUpdateData, ctrl?: Control): Promise<PostmortemEntity>;
}
export { PostmortemEntity };
