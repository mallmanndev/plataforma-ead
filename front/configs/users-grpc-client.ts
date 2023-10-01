import * as grpc from '@grpc/grpc-js';
import * as protoLoader from '@grpc/proto-loader';
import path from 'path';
import {ProtoGrpcType} from "@/protos/generated/users";

const PROTO_PATH = path.join(process.cwd(), './protos/users.proto');

// suggested options for similarity to loading grpc.load behavior
const packageDefinition = protoLoader.loadSync(PROTO_PATH, {
    keepCase: true,
    defaults: true,
    oneofs: true,
});

const UsersService = (
    grpc.loadPackageDefinition(packageDefinition) as unknown as ProtoGrpcType
).UsersService;

export default UsersService;