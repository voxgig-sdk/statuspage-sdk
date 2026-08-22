import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { MetricsProvider, MetricsProviderLoadMatch, MetricsProviderListMatch, MetricsProviderCreateData, MetricsProviderUpdateData, MetricsProviderRemoveMatch } from '../StatuspageTypes';
declare class MetricsProviderEntity extends StatuspageEntityBase<MetricsProvider> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: MetricsProviderEntity): MetricsProviderEntity;
    load(this: any, reqmatch?: MetricsProviderLoadMatch, ctrl?: Control): Promise<MetricsProviderEntity>;
    list(this: any, reqmatch?: MetricsProviderListMatch, ctrl?: Control): Promise<MetricsProviderEntity[]>;
    create(this: any, reqdata?: MetricsProviderCreateData, ctrl?: Control): Promise<MetricsProviderEntity>;
    update(this: any, reqdata?: MetricsProviderUpdateData, ctrl?: Control): Promise<MetricsProviderEntity>;
    remove(this: any, reqmatch?: MetricsProviderRemoveMatch, ctrl?: Control): Promise<MetricsProviderEntity>;
}
export { MetricsProviderEntity };
