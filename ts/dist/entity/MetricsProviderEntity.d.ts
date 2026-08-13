import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { MetricsProvider, MetricsProviderLoadMatch, MetricsProviderListMatch, MetricsProviderCreateData, MetricsProviderUpdateData, MetricsProviderRemoveMatch } from '../StatuspageTypes';
declare class MetricsProviderEntity extends StatuspageEntityBase<MetricsProvider> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: MetricsProviderEntity): MetricsProviderEntity;
    load(this: any, reqmatch?: MetricsProviderLoadMatch, ctrl?: Control): Promise<MetricsProvider>;
    list(this: any, reqmatch?: MetricsProviderListMatch, ctrl?: Control): Promise<MetricsProvider[]>;
    create(this: any, reqdata?: MetricsProviderCreateData, ctrl?: Control): Promise<MetricsProvider>;
    update(this: any, reqdata?: MetricsProviderUpdateData, ctrl?: Control): Promise<MetricsProvider>;
    remove(this: any, reqmatch?: MetricsProviderRemoveMatch, ctrl?: Control): Promise<MetricsProvider>;
}
export { MetricsProviderEntity };
