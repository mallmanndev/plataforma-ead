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
	Get(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*GetCoursesResponse, error)
	Create(ctx context.Context, in *CreateCourseRequest, opts ...grpc.CallOption) (*Course, error)
	Update(ctx context.Context, in *UpdateCourseRequest, opts ...grpc.CallOption) (*Course, error)
	Delete(ctx context.Context, in *DeleteCourseRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
	CreateSection(ctx context.Context, in *CreateCourseSectionRequest, opts ...grpc.CallOption) (*Course, error)
	UpdateSection(ctx context.Context, in *UpdateCourseSectionRequest, opts ...grpc.CallOption) (*Course, error)
	DeleteSection(ctx context.Context, in *DeleteCourseSectionRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error)
	GetSection(ctx context.Context, in *GetSectionRequest, opts ...grpc.CallOption) (*CourseSection, error)
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*Course, error)
	UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*Course, error)
	DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*Course, error)
	GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*CourseItem, error)
}

type coursesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCoursesServiceClient(cc grpc.ClientConnInterface) CoursesServiceClient {
	return &coursesServiceClient{cc}
}

func (c *coursesServiceClient) Get(ctx context.Context, in *GetCoursesRequest, opts ...grpc.CallOption) (*GetCoursesResponse, error) {
	out := new(GetCoursesResponse)
	err := c.cc.Invoke(ctx, "/CoursesService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *coursesServiceClient) CreateSection(ctx context.Context, in *CreateCourseSectionRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/CreateSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) UpdateSection(ctx context.Context, in *UpdateCourseSectionRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/UpdateSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) DeleteSection(ctx context.Context, in *DeleteCourseSectionRequest, opts ...grpc.CallOption) (*DeleteCourseResponse, error) {
	out := new(DeleteCourseResponse)
	err := c.cc.Invoke(ctx, "/CoursesService/DeleteSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) GetSection(ctx context.Context, in *GetSectionRequest, opts ...grpc.CallOption) (*CourseSection, error) {
	out := new(CourseSection)
	err := c.cc.Invoke(ctx, "/CoursesService/GetSection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) UpdateItem(ctx context.Context, in *UpdateItemRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/UpdateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) DeleteItem(ctx context.Context, in *DeleteItemRequest, opts ...grpc.CallOption) (*Course, error) {
	out := new(Course)
	err := c.cc.Invoke(ctx, "/CoursesService/DeleteItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *coursesServiceClient) GetItem(ctx context.Context, in *GetItemRequest, opts ...grpc.CallOption) (*CourseItem, error) {
	out := new(CourseItem)
	err := c.cc.Invoke(ctx, "/CoursesService/GetItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CoursesServiceServer is the server API for CoursesService service.
// All implementations must embed UnimplementedCoursesServiceServer
// for forward compatibility
type CoursesServiceServer interface {
	Get(context.Context, *GetCoursesRequest) (*GetCoursesResponse, error)
	Create(context.Context, *CreateCourseRequest) (*Course, error)
	Update(context.Context, *UpdateCourseRequest) (*Course, error)
	Delete(context.Context, *DeleteCourseRequest) (*DeleteCourseResponse, error)
	CreateSection(context.Context, *CreateCourseSectionRequest) (*Course, error)
	UpdateSection(context.Context, *UpdateCourseSectionRequest) (*Course, error)
	DeleteSection(context.Context, *DeleteCourseSectionRequest) (*DeleteCourseResponse, error)
	GetSection(context.Context, *GetSectionRequest) (*CourseSection, error)
	CreateItem(context.Context, *CreateItemRequest) (*Course, error)
	UpdateItem(context.Context, *UpdateItemRequest) (*Course, error)
	DeleteItem(context.Context, *DeleteItemRequest) (*Course, error)
	GetItem(context.Context, *GetItemRequest) (*CourseItem, error)
	mustEmbedUnimplementedCoursesServiceServer()
}

// UnimplementedCoursesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCoursesServiceServer struct {
}

func (UnimplementedCoursesServiceServer) Get(context.Context, *GetCoursesRequest) (*GetCoursesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
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
func (UnimplementedCoursesServiceServer) CreateSection(context.Context, *CreateCourseSectionRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSection not implemented")
}
func (UnimplementedCoursesServiceServer) UpdateSection(context.Context, *UpdateCourseSectionRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSection not implemented")
}
func (UnimplementedCoursesServiceServer) DeleteSection(context.Context, *DeleteCourseSectionRequest) (*DeleteCourseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSection not implemented")
}
func (UnimplementedCoursesServiceServer) GetSection(context.Context, *GetSectionRequest) (*CourseSection, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSection not implemented")
}
func (UnimplementedCoursesServiceServer) CreateItem(context.Context, *CreateItemRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateItem not implemented")
}
func (UnimplementedCoursesServiceServer) UpdateItem(context.Context, *UpdateItemRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateItem not implemented")
}
func (UnimplementedCoursesServiceServer) DeleteItem(context.Context, *DeleteItemRequest) (*Course, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteItem not implemented")
}
func (UnimplementedCoursesServiceServer) GetItem(context.Context, *GetItemRequest) (*CourseItem, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetItem not implemented")
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

func _CoursesService_CreateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCourseSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).CreateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/CreateSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).CreateSection(ctx, req.(*CreateCourseSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_UpdateSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCourseSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).UpdateSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/UpdateSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).UpdateSection(ctx, req.(*UpdateCourseSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_DeleteSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCourseSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).DeleteSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/DeleteSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).DeleteSection(ctx, req.(*DeleteCourseSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_GetSection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).GetSection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/GetSection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).GetSection(ctx, req.(*GetSectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_UpdateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).UpdateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/UpdateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).UpdateItem(ctx, req.(*UpdateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_DeleteItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).DeleteItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/DeleteItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).DeleteItem(ctx, req.(*DeleteItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CoursesService_GetItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CoursesServiceServer).GetItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CoursesService/GetItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CoursesServiceServer).GetItem(ctx, req.(*GetItemRequest))
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
			MethodName: "Get",
			Handler:    _CoursesService_Get_Handler,
		},
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
			MethodName: "CreateSection",
			Handler:    _CoursesService_CreateSection_Handler,
		},
		{
			MethodName: "UpdateSection",
			Handler:    _CoursesService_UpdateSection_Handler,
		},
		{
			MethodName: "DeleteSection",
			Handler:    _CoursesService_DeleteSection_Handler,
		},
		{
			MethodName: "GetSection",
			Handler:    _CoursesService_GetSection_Handler,
		},
		{
			MethodName: "CreateItem",
			Handler:    _CoursesService_CreateItem_Handler,
		},
		{
			MethodName: "UpdateItem",
			Handler:    _CoursesService_UpdateItem_Handler,
		},
		{
			MethodName: "DeleteItem",
			Handler:    _CoursesService_DeleteItem_Handler,
		},
		{
			MethodName: "GetItem",
			Handler:    _CoursesService_GetItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-course/protos/courses.proto",
}
