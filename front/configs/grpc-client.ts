import * as grpc from "@grpc/grpc-js";
import * as protoLoader from "@grpc/proto-loader";
import path from "path";

const HOST_SERVICE_COURSE = "service-course:3000";

const PROTO_LOADER_CONFIGS = {
  keepCase: true,
  defaults: true,
  oneofs: true,
};

const PROTO_PATH = path.join(process.cwd(), "./protos/users.proto");
const FILES_PROTO_PATH = path.join(process.cwd(), "./protos/files.proto");
const COURSES_PROTO_PATH = path.join(process.cwd(), "./protos/courses.proto");

// suggested options for similarity to loading grpc.load behavior
const packageDefinition = protoLoader.loadSync(
  PROTO_PATH,
  PROTO_LOADER_CONFIGS
);

const filesPackageDefinition = protoLoader.loadSync(
  FILES_PROTO_PATH,
  PROTO_LOADER_CONFIGS
);

const coursesPackageDefinition = protoLoader.loadSync(
  COURSES_PROTO_PATH,
  PROTO_LOADER_CONFIGS
);

export const UsersService = (
  grpc.loadPackageDefinition(packageDefinition) as any
).UsersService;
export const usersService = new UsersService(
  process.env.SERVICE_COURSE_URL,
  grpc.credentials.createInsecure()
);

export const FilesUploadService = (
  grpc.loadPackageDefinition(filesPackageDefinition) as unknown as any
).FileUploadService;
export const filesGrpcClient = new FilesUploadService(
  HOST_SERVICE_COURSE,
  grpc.credentials.createInsecure()
);

const CoursesService = (
  grpc.loadPackageDefinition(coursesPackageDefinition) as unknown as any
).CoursesService;
export const coursesClient = new CoursesService(
  HOST_SERVICE_COURSE,
  grpc.credentials.createInsecure()
);
