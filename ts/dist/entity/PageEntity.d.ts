import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Page, PageLoadMatch, PageListMatch, PageUpdateData } from '../StatuspageTypes';
declare class PageEntity extends StatuspageEntityBase<Page> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: PageEntity): PageEntity;
    load(this: any, reqmatch?: PageLoadMatch, ctrl?: Control): Promise<PageEntity>;
    list(this: any, reqmatch?: PageListMatch, ctrl?: Control): Promise<PageEntity[]>;
    update(this: any, reqdata?: PageUpdateData, ctrl?: Control): Promise<PageEntity>;
}
export { PageEntity };
