import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { ComponentGroupUptime, ComponentGroupUptimeLoadMatch } from '../StatuspageTypes';
declare class ComponentGroupUptimeEntity extends StatuspageEntityBase<ComponentGroupUptime> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: ComponentGroupUptimeEntity): ComponentGroupUptimeEntity;
    load(this: any, reqmatch?: ComponentGroupUptimeLoadMatch, ctrl?: Control): Promise<ComponentGroupUptime>;
}
export { ComponentGroupUptimeEntity };
