import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { IncidentUpdate, IncidentUpdateUpdateData } from '../StatuspageTypes';
declare class IncidentUpdateEntity extends StatuspageEntityBase<IncidentUpdate> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: IncidentUpdateEntity): IncidentUpdateEntity;
    update(this: any, reqdata?: IncidentUpdateUpdateData, ctrl?: Control): Promise<IncidentUpdateEntity>;
}
export { IncidentUpdateEntity };
