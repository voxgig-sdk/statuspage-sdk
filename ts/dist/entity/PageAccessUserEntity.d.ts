import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { PageAccessUser, PageAccessUserLoadMatch, PageAccessUserListMatch, PageAccessUserCreateData, PageAccessUserUpdateData, PageAccessUserRemoveMatch } from '../StatuspageTypes';
declare class PageAccessUserEntity extends StatuspageEntityBase<PageAccessUser> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: PageAccessUserEntity): PageAccessUserEntity;
    load(this: any, reqmatch?: PageAccessUserLoadMatch, ctrl?: Control): Promise<PageAccessUser>;
    list(this: any, reqmatch?: PageAccessUserListMatch, ctrl?: Control): Promise<PageAccessUser[]>;
    create(this: any, reqdata?: PageAccessUserCreateData, ctrl?: Control): Promise<PageAccessUser>;
    update(this: any, reqdata?: PageAccessUserUpdateData, ctrl?: Control): Promise<PageAccessUser>;
    remove(this: any, reqmatch?: PageAccessUserRemoveMatch, ctrl?: Control): Promise<PageAccessUser>;
}
export { PageAccessUserEntity };
