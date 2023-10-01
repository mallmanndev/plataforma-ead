import type * as grpc from '@grpc/grpc-js';
import type { MessageTypeDefinition } from '@grpc/proto-loader';

import type { UsersServiceClient as _UsersServiceClient, UsersServiceDefinition as _UsersServiceDefinition } from './UsersService';

type SubtypeConstructor<Constructor extends new (...args: any) => any, Subtype> = {
  new(...args: ConstructorParameters<Constructor>): Subtype;
};

export interface ProtoGrpcType {
  CreateUserRequest: MessageTypeDefinition
  LoginRequest: MessageTypeDefinition
  User: MessageTypeDefinition
  UsersService: SubtypeConstructor<typeof grpc.Client, _UsersServiceClient> & { service: _UsersServiceDefinition }
}

