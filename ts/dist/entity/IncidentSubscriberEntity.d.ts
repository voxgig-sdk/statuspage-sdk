import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { IncidentSubscriber, IncidentSubscriberCreateData } from '../StatuspageTypes';
declare class IncidentSubscriberEntity extends StatuspageEntityBase<IncidentSubscriber> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: IncidentSubscriberEntity): IncidentSubscriberEntity;
    create(this: any, reqdata?: IncidentSubscriberCreateData, ctrl?: Control): Promise<IncidentSubscriber>;
}
export { IncidentSubscriberEntity };
