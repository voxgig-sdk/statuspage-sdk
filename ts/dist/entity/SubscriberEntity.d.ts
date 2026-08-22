import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Subscriber, SubscriberLoadMatch, SubscriberListMatch, SubscriberCreateData, SubscriberUpdateData, SubscriberRemoveMatch } from '../StatuspageTypes';
declare class SubscriberEntity extends StatuspageEntityBase<Subscriber> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: SubscriberEntity): SubscriberEntity;
    load(this: any, reqmatch?: SubscriberLoadMatch, ctrl?: Control): Promise<SubscriberEntity>;
    list(this: any, reqmatch?: SubscriberListMatch, ctrl?: Control): Promise<SubscriberEntity[]>;
    create(this: any, reqdata?: SubscriberCreateData, ctrl?: Control): Promise<SubscriberEntity>;
    update(this: any, reqdata?: SubscriberUpdateData, ctrl?: Control): Promise<SubscriberEntity>;
    remove(this: any, reqmatch?: SubscriberRemoveMatch, ctrl?: Control): Promise<SubscriberEntity>;
}
export { SubscriberEntity };
