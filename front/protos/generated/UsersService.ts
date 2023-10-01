// Original file: protos/users.proto

import type * as grpc from '@grpc/grpc-js'
import type { MethodDefinition } from '@grpc/proto-loader'
import type { CreateUserRequest as _CreateUserRequest, CreateUserRequest__Output as _CreateUserRequest__Output } from './CreateUserRequest';
import type { LoginRequest as _LoginRequest, LoginRequest__Output as _LoginRequest__Output } from './LoginRequest';
import type { User as _User, User__Output as _User__Output } from './User';

export interface UsersServiceClient extends grpc.Client {
  Create(argument: _CreateUserRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  Create(argument: _CreateUserRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  Create(argument: _CreateUserRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  Create(argument: _CreateUserRequest, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  create(argument: _CreateUserRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  create(argument: _CreateUserRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  create(argument: _CreateUserRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  create(argument: _CreateUserRequest, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  
  Login(argument: _LoginRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  Login(argument: _LoginRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  Login(argument: _LoginRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  Login(argument: _LoginRequest, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  login(argument: _LoginRequest, metadata: grpc.Metadata, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  login(argument: _LoginRequest, metadata: grpc.Metadata, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  login(argument: _LoginRequest, options: grpc.CallOptions, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  login(argument: _LoginRequest, callback: grpc.requestCallback<_User__Output>): grpc.ClientUnaryCall;
  
}

export interface UsersServiceHandlers extends grpc.UntypedServiceImplementation {
  Create: grpc.handleUnaryCall<_CreateUserRequest__Output, _User>;
  
  Login: grpc.handleUnaryCall<_LoginRequest__Output, _User>;
  
}

export interface UsersServiceDefinition extends grpc.ServiceDefinition {
  Create: MethodDefinition<_CreateUserRequest, _User, _CreateUserRequest__Output, _User__Output>
  Login: MethodDefinition<_LoginRequest, _User, _LoginRequest__Output, _User__Output>
}
