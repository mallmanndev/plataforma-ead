// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: service-course/protos/files.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type VideoUploadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*VideoUploadRequest_Info
	//	*VideoUploadRequest_Chunk
	Data isVideoUploadRequest_Data `protobuf_oneof:"data"`
}

func (x *VideoUploadRequest) Reset() {
	*x = VideoUploadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_course_protos_files_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoUploadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoUploadRequest) ProtoMessage() {}

func (x *VideoUploadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_course_protos_files_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoUploadRequest.ProtoReflect.Descriptor instead.
func (*VideoUploadRequest) Descriptor() ([]byte, []int) {
	return file_service_course_protos_files_proto_rawDescGZIP(), []int{0}
}

func (m *VideoUploadRequest) GetData() isVideoUploadRequest_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *VideoUploadRequest) GetInfo() *VideoInfo {
	if x, ok := x.GetData().(*VideoUploadRequest_Info); ok {
		return x.Info
	}
	return nil
}

func (x *VideoUploadRequest) GetChunk() []byte {
	if x, ok := x.GetData().(*VideoUploadRequest_Chunk); ok {
		return x.Chunk
	}
	return nil
}

type isVideoUploadRequest_Data interface {
	isVideoUploadRequest_Data()
}

type VideoUploadRequest_Info struct {
	Info *VideoInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type VideoUploadRequest_Chunk struct {
	Chunk []byte `protobuf:"bytes,2,opt,name=chunk,proto3,oneof"`
}

func (*VideoUploadRequest_Info) isVideoUploadRequest_Data() {}

func (*VideoUploadRequest_Chunk) isVideoUploadRequest_Data() {}

type VideoInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *VideoInfo) Reset() {
	*x = VideoInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_course_protos_files_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoInfo) ProtoMessage() {}

func (x *VideoInfo) ProtoReflect() protoreflect.Message {
	mi := &file_service_course_protos_files_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoInfo.ProtoReflect.Descriptor instead.
func (*VideoInfo) Descriptor() ([]byte, []int) {
	return file_service_course_protos_files_proto_rawDescGZIP(), []int{1}
}

func (x *VideoInfo) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *VideoInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

type VideoUploadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *VideoUploadResponse) Reset() {
	*x = VideoUploadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_course_protos_files_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoUploadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoUploadResponse) ProtoMessage() {}

func (x *VideoUploadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_course_protos_files_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoUploadResponse.ProtoReflect.Descriptor instead.
func (*VideoUploadResponse) Descriptor() ([]byte, []int) {
	return file_service_course_protos_files_proto_rawDescGZIP(), []int{2}
}

func (x *VideoUploadResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type VideoResolution struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resolution         string `protobuf:"bytes,1,opt,name=resolution,proto3" json:"resolution,omitempty"`
	CompleteResolution string `protobuf:"bytes,2,opt,name=complete_resolution,json=completeResolution,proto3" json:"complete_resolution,omitempty"`
	Url                string `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *VideoResolution) Reset() {
	*x = VideoResolution{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_course_protos_files_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoResolution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoResolution) ProtoMessage() {}

func (x *VideoResolution) ProtoReflect() protoreflect.Message {
	mi := &file_service_course_protos_files_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoResolution.ProtoReflect.Descriptor instead.
func (*VideoResolution) Descriptor() ([]byte, []int) {
	return file_service_course_protos_files_proto_rawDescGZIP(), []int{3}
}

func (x *VideoResolution) GetResolution() string {
	if x != nil {
		return x.Resolution
	}
	return ""
}

func (x *VideoResolution) GetCompleteResolution() string {
	if x != nil {
		return x.CompleteResolution
	}
	return ""
}

func (x *VideoResolution) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetVideoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetVideoRequest) Reset() {
	*x = GetVideoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_course_protos_files_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoRequest) ProtoMessage() {}

func (x *GetVideoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_course_protos_files_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoRequest.ProtoReflect.Descriptor instead.
func (*GetVideoRequest) Descriptor() ([]byte, []int) {
	return file_service_course_protos_files_proto_rawDescGZIP(), []int{4}
}

func (x *GetVideoRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetVideoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Type        string             `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Status      string             `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Size        int64              `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	Resolutions []*VideoResolution `protobuf:"bytes,5,rep,name=resolutions,proto3" json:"resolutions,omitempty"`
	CreatedAt   string             `protobuf:"bytes,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   string             `protobuf:"bytes,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Url         string             `protobuf:"bytes,8,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *GetVideoResponse) Reset() {
	*x = GetVideoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_course_protos_files_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVideoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVideoResponse) ProtoMessage() {}

func (x *GetVideoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_course_protos_files_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVideoResponse.ProtoReflect.Descriptor instead.
func (*GetVideoResponse) Descriptor() ([]byte, []int) {
	return file_service_course_protos_files_proto_rawDescGZIP(), []int{5}
}

func (x *GetVideoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetVideoResponse) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *GetVideoResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetVideoResponse) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetVideoResponse) GetResolutions() []*VideoResolution {
	if x != nil {
		return x.Resolutions
	}
	return nil
}

func (x *GetVideoResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetVideoResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *GetVideoResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_service_course_protos_files_proto protoreflect.FileDescriptor

var file_service_course_protos_files_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x56, 0x0a, 0x12, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x04, 0x69, 0x6e, 0x66,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x49,
	0x6e, 0x66, 0x6f, 0x48, 0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x05, 0x63,
	0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x05, 0x63, 0x68,
	0x75, 0x6e, 0x6b, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x33, 0x0a, 0x09, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x25, 0x0a, 0x13, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x74, 0x0a, 0x0f, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x21, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x22, 0xe4, 0x01, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x32, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x32, 0x84, 0x01, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65,
	0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a,
	0x0b, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x13, 0x2e, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x14, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x31, 0x0a, 0x08, 0x47,
	0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64,
	0x65, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x56,
	0x69, 0x64, 0x65, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_course_protos_files_proto_rawDescOnce sync.Once
	file_service_course_protos_files_proto_rawDescData = file_service_course_protos_files_proto_rawDesc
)

func file_service_course_protos_files_proto_rawDescGZIP() []byte {
	file_service_course_protos_files_proto_rawDescOnce.Do(func() {
		file_service_course_protos_files_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_course_protos_files_proto_rawDescData)
	})
	return file_service_course_protos_files_proto_rawDescData
}

var file_service_course_protos_files_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_service_course_protos_files_proto_goTypes = []interface{}{
	(*VideoUploadRequest)(nil),  // 0: VideoUploadRequest
	(*VideoInfo)(nil),           // 1: VideoInfo
	(*VideoUploadResponse)(nil), // 2: VideoUploadResponse
	(*VideoResolution)(nil),     // 3: VideoResolution
	(*GetVideoRequest)(nil),     // 4: GetVideoRequest
	(*GetVideoResponse)(nil),    // 5: GetVideoResponse
}
var file_service_course_protos_files_proto_depIdxs = []int32{
	1, // 0: VideoUploadRequest.info:type_name -> VideoInfo
	3, // 1: GetVideoResponse.resolutions:type_name -> VideoResolution
	0, // 2: FileUploadService.VideoUpload:input_type -> VideoUploadRequest
	4, // 3: FileUploadService.GetVideo:input_type -> GetVideoRequest
	2, // 4: FileUploadService.VideoUpload:output_type -> VideoUploadResponse
	5, // 5: FileUploadService.GetVideo:output_type -> GetVideoResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_service_course_protos_files_proto_init() }
func file_service_course_protos_files_proto_init() {
	if File_service_course_protos_files_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_course_protos_files_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoUploadRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_course_protos_files_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_course_protos_files_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoUploadResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_course_protos_files_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoResolution); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_course_protos_files_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_service_course_protos_files_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVideoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_service_course_protos_files_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*VideoUploadRequest_Info)(nil),
		(*VideoUploadRequest_Chunk)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_service_course_protos_files_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_course_protos_files_proto_goTypes,
		DependencyIndexes: file_service_course_protos_files_proto_depIdxs,
		MessageInfos:      file_service_course_protos_files_proto_msgTypes,
	}.Build()
	File_service_course_protos_files_proto = out.File
	file_service_course_protos_files_proto_rawDesc = nil
	file_service_course_protos_files_proto_goTypes = nil
	file_service_course_protos_files_proto_depIdxs = nil
}
