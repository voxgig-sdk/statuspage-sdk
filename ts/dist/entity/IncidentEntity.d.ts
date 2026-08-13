import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Incident, IncidentLoadMatch, IncidentListMatch, IncidentCreateData, IncidentUpdateData, IncidentRemoveMatch } from '../StatuspageTypes';
declare class IncidentEntity extends StatuspageEntityBase<Incident> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: IncidentEntity): IncidentEntity;
    load(this: any, reqmatch?: IncidentLoadMatch, ctrl?: Control): Promise<Incident>;
    list(this: any, reqmatch?: IncidentListMatch, ctrl?: Control): Promise<Incident[]>;
    create(this: any, reqdata?: IncidentCreateData, ctrl?: Control): Promise<Incident>;
    update(this: any, reqdata?: IncidentUpdateData, ctrl?: Control): Promise<Incident>;
    remove(this: any, reqmatch?: IncidentRemoveMatch, ctrl?: Control): Promise<Incident>;
}
export { IncidentEntity };
