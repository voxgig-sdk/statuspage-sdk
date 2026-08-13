import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { StatusEmbedConfig, StatusEmbedConfigLoadMatch, StatusEmbedConfigUpdateData } from '../StatuspageTypes';
declare class StatusEmbedConfigEntity extends StatuspageEntityBase<StatusEmbedConfig> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: StatusEmbedConfigEntity): StatusEmbedConfigEntity;
    load(this: any, reqmatch?: StatusEmbedConfigLoadMatch, ctrl?: Control): Promise<StatusEmbedConfig>;
    update(this: any, reqdata?: StatusEmbedConfigUpdateData, ctrl?: Control): Promise<StatusEmbedConfig>;
}
export { StatusEmbedConfigEntity };
