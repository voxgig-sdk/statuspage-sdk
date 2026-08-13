import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Subscriber, SubscriberLoadMatch, SubscriberListMatch, SubscriberCreateData, SubscriberUpdateData, SubscriberRemoveMatch } from '../StatuspageTypes';
declare class SubscriberEntity extends StatuspageEntityBase<Subscriber> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: SubscriberEntity): SubscriberEntity;
    load(this: any, reqmatch?: SubscriberLoadMatch, ctrl?: Control): Promise<Subscriber>;
    list(this: any, reqmatch?: SubscriberListMatch, ctrl?: Control): Promise<Subscriber[]>;
    create(this: any, reqdata?: SubscriberCreateData, ctrl?: Control): Promise<Subscriber>;
    update(this: any, reqdata?: SubscriberUpdateData, ctrl?: Control): Promise<Subscriber>;
    remove(this: any, reqmatch?: SubscriberRemoveMatch, ctrl?: Control): Promise<Subscriber>;
}
export { SubscriberEntity };
