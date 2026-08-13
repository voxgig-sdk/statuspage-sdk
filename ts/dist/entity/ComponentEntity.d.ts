import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Component, ComponentLoadMatch, ComponentListMatch, ComponentCreateData, ComponentUpdateData, ComponentRemoveMatch } from '../StatuspageTypes';
declare class ComponentEntity extends StatuspageEntityBase<Component> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: ComponentEntity): ComponentEntity;
    load(this: any, reqmatch?: ComponentLoadMatch, ctrl?: Control): Promise<Component>;
    list(this: any, reqmatch?: ComponentListMatch, ctrl?: Control): Promise<Component[]>;
    create(this: any, reqdata?: ComponentCreateData, ctrl?: Control): Promise<Component>;
    update(this: any, reqdata?: ComponentUpdateData, ctrl?: Control): Promise<Component>;
    remove(this: any, reqmatch?: ComponentRemoveMatch, ctrl?: Control): Promise<Component>;
}
export { ComponentEntity };
