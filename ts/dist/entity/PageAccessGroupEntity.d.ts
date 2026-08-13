import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { PageAccessGroup, PageAccessGroupLoadMatch, PageAccessGroupListMatch, PageAccessGroupCreateData, PageAccessGroupUpdateData, PageAccessGroupRemoveMatch } from '../StatuspageTypes';
declare class PageAccessGroupEntity extends StatuspageEntityBase<PageAccessGroup> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: PageAccessGroupEntity): PageAccessGroupEntity;
    load(this: any, reqmatch?: PageAccessGroupLoadMatch, ctrl?: Control): Promise<PageAccessGroup>;
    list(this: any, reqmatch?: PageAccessGroupListMatch, ctrl?: Control): Promise<PageAccessGroup[]>;
    create(this: any, reqdata?: PageAccessGroupCreateData, ctrl?: Control): Promise<PageAccessGroup>;
    update(this: any, reqdata?: PageAccessGroupUpdateData, ctrl?: Control): Promise<PageAccessGroup>;
    remove(this: any, reqmatch?: PageAccessGroupRemoveMatch, ctrl?: Control): Promise<PageAccessGroup>;
}
export { PageAccessGroupEntity };
