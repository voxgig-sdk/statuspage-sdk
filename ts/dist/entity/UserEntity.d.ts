import { StatuspageEntityBase } from '../StatuspageEntityBase';
import type { StatuspageSDK } from '../StatuspageSDK';
import type { Control } from '../types';
import type { User, UserListMatch, UserCreateData, UserRemoveMatch } from '../StatuspageTypes';
declare class UserEntity extends StatuspageEntityBase<User> {
    constructor(client: StatuspageSDK, entopts: any);
    make(this: UserEntity): UserEntity;
    list(this: any, reqmatch?: UserListMatch, ctrl?: Control): Promise<User[]>;
    create(this: any, reqdata?: UserCreateData, ctrl?: Control): Promise<User>;
    remove(this: any, reqmatch?: UserRemoveMatch, ctrl?: Control): Promise<User>;
}
export { UserEntity };
