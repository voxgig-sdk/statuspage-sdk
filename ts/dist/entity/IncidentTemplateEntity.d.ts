import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { IncidentTemplate, IncidentTemplateListMatch, IncidentTemplateCreateData } from '../StatuspageTypes';
declare class IncidentTemplateEntity extends StatuspageEntityBase<IncidentTemplate> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: IncidentTemplateEntity): IncidentTemplateEntity;
    list(this: any, reqmatch?: IncidentTemplateListMatch, ctrl?: Control): Promise<IncidentTemplate[]>;
    create(this: any, reqdata?: IncidentTemplateCreateData, ctrl?: Control): Promise<IncidentTemplate>;
}
export { IncidentTemplateEntity };
