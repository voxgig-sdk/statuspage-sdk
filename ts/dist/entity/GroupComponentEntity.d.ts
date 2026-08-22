import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { GroupComponent, GroupComponentLoadMatch, GroupComponentListMatch, GroupComponentCreateData, GroupComponentUpdateData, GroupComponentRemoveMatch } from '../StatuspageTypes';
declare class GroupComponentEntity extends StatuspageEntityBase<GroupComponent> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: GroupComponentEntity): GroupComponentEntity;
    load(this: any, reqmatch?: GroupComponentLoadMatch, ctrl?: Control): Promise<GroupComponentEntity>;
    list(this: any, reqmatch?: GroupComponentListMatch, ctrl?: Control): Promise<GroupComponentEntity[]>;
    create(this: any, reqdata?: GroupComponentCreateData, ctrl?: Control): Promise<GroupComponentEntity>;
    update(this: any, reqdata?: GroupComponentUpdateData, ctrl?: Control): Promise<GroupComponentEntity>;
    remove(this: any, reqmatch?: GroupComponentRemoveMatch, ctrl?: Control): Promise<GroupComponentEntity>;
}
export { GroupComponentEntity };
