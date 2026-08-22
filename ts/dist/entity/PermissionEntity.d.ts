import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Permission, PermissionLoadMatch, PermissionUpdateData } from '../StatuspageTypes';
declare class PermissionEntity extends StatuspageEntityBase<Permission> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: PermissionEntity): PermissionEntity;
    load(this: any, reqmatch?: PermissionLoadMatch, ctrl?: Control): Promise<PermissionEntity>;
    update(this: any, reqdata?: PermissionUpdateData, ctrl?: Control): Promise<PermissionEntity>;
}
export { PermissionEntity };
