import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Metric, MetricLoadMatch, MetricListMatch, MetricCreateData, MetricUpdateData, MetricRemoveMatch } from '../StatuspageTypes';
declare class MetricEntity extends StatuspageEntityBase<Metric> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: MetricEntity): MetricEntity;
    load(this: any, reqmatch?: MetricLoadMatch, ctrl?: Control): Promise<MetricEntity>;
    list(this: any, reqmatch?: MetricListMatch, ctrl?: Control): Promise<MetricEntity[]>;
    create(this: any, reqdata?: MetricCreateData, ctrl?: Control): Promise<MetricEntity>;
    update(this: any, reqdata?: MetricUpdateData, ctrl?: Control): Promise<MetricEntity>;
    remove(this: any, reqmatch?: MetricRemoveMatch, ctrl?: Control): Promise<MetricEntity>;
}
export { MetricEntity };
