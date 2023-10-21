// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: service-course/protos/courses.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CoursesServiceClient is the client API for CoursesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CoursesServiceClient interface {
	Create(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error)
	Update(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*Course, error)
	Delete(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
	CreateCourseSection(ctx context.Context, in *CreateCourseSectionRequest, opts ...grpc.CallOption) (*Course, error)
	Get(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*GetCoursesResponse, error)
}

type coursesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursesServiceClient(cc grpc.ClientConnInterface) CoursesServiceClient {
	return &coursesServiceClient{cc}
}

func (c *coursesServiceClient) Create(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) Update(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) Delete(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error) {
	out := new(DeleteCourseResponse)
	err := c.cc.Invoke(ctx, "/CoursesService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) CreateCourseSection(ctx context.Context, in *CreateCourseSectionRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/CreateCourseSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) Get(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*GetCoursesResponse, error) {
	out := new(GetCoursesResponse)
	err := c.cc.Invoke(ctx, "/CoursesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursesServiceServer is the server API for CoursesService service.
// All implementations must embed UnimplementedCoursesServiceServer
// for forward compatibility
type CoursesServiceServer interface {
	Create(context.Context, *CreateCourseRequest) (*Course, error)
	Update(context.Context, *UpdateCourseRequest) (*Course, error)
	Delete(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error)
	CreateCourseSection(context.Context, *CreateCourseSectionRequest) (*Course, error)
	Get(context.Context, *GetCoursesRequest) (*GetCoursesResponse, error)
	mustEmbedUnimplementedCoursesServiceServer()
}

// UnimplementedCoursesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoursesServiceServer struct {
}

func (UnimplementedCoursesServiceServer) Create(context.Context, *CreateCourseRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedCoursesServiceServer) Update(context.Context, *UpdateCourseRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedCoursesServiceServer) Delete(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCoursesServiceServer) CreateCourseSection(context.Context, *CreateCourseSectionRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCourseSection not implemented")
}
func (UnimplementedCoursesServiceServer) Get(context.Context, *GetCoursesRequest) (*GetCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCoursesServiceServer) mustEmbedUnimplementedCoursesServiceServer() {}

// UnsafeCoursesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CoursesServiceServer will
// result in compilation errors.
type UnsafeCoursesServiceServer interface {
	mustEmbedUnimplementedCoursesServiceServer()
}

func RegisterCoursesServiceServer(s grpc.ServiceRegistrar, srv CoursesServiceServer) {
	s.RegisterService(&CoursesService_ServiceDesc, srv)
}

func _CoursesService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).Create(ctx, req.(*CreateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).Update(ctx, req.(*UpdateCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).Delete(ctx, req.(*DeleteCourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_CreateCourseSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).CreateCourseSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/CreateCourseSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).CreateCourseSection(ctx, req.(*CreateCourseSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCoursesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).Get(ctx, req.(*GetCoursesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CoursesService_ServiceDesc is the grpc.ServiceDesc for CoursesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CoursesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CoursesService",
	HandlerType: (*CoursesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _CoursesService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _CoursesService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CoursesService_Delete_Handler,
		},
		{
			MethodName: "CreateCourseSection",
			Handler:    _CoursesService_CreateCourseSection_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CoursesService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-course/protos/courses.proto",
}
