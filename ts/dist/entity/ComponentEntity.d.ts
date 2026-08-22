import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { Component, ComponentLoadMatch, ComponentListMatch, ComponentCreateData, ComponentUpdateData, ComponentRemoveMatch } from '../StatuspageTypes';
declare class ComponentEntity extends StatuspageEntityBase<Component> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: ComponentEntity): ComponentEntity;
    load(this: any, reqmatch?: ComponentLoadMatch, ctrl?: Control): Promise<ComponentEntity>;
    list(this: any, reqmatch?: ComponentListMatch, ctrl?: Control): Promise<ComponentEntity[]>;
    create(this: any, reqdata?: ComponentCreateData, ctrl?: Control): Promise<ComponentEntity>;
    update(this: any, reqdata?: ComponentUpdateData, ctrl?: Control): Promise<ComponentEntity>;
    remove(this: any, reqmatch?: ComponentRemoveMatch, ctrl?: Control): Promise<ComponentEntity>;
}
export { ComponentEntity };
