import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { IncidentPostmortem, IncidentPostmortemRemoveMatch } from '../StatuspageTypes';
declare class IncidentPostmortemEntity extends StatuspageEntityBase<IncidentPostmortem> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: IncidentPostmortemEntity): IncidentPostmortemEntity;
    remove(this: any, reqmatch?: IncidentPostmortemRemoveMatch, ctrl?: Control): Promise<IncidentPostmortemEntity>;
}
export { IncidentPostmortemEntity };
