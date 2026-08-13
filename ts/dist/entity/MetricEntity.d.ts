import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Metric, MetricLoadMatch, MetricListMatch, MetricCreateData, MetricUpdateData, MetricRemoveMatch } from '../StatuspageTypes';
declare class MetricEntity extends StatuspageEntityBase<Metric> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: MetricEntity): MetricEntity;
    load(this: any, reqmatch?: MetricLoadMatch, ctrl?: Control): Promise<Metric>;
    list(this: any, reqmatch?: MetricListMatch, ctrl?: Control): Promise<Metric[]>;
    create(this: any, reqdata?: MetricCreateData, ctrl?: Control): Promise<Metric>;
    update(this: any, reqdata?: MetricUpdateData, ctrl?: Control): Promise<Metric>;
    remove(this: any, reqmatch?: MetricRemoveMatch, ctrl?: Control): Promise<Metric>;
}
export { MetricEntity };
