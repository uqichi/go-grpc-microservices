// Code generated by protoc-gen-go. DO NOT EDIT.
// source: activity/activity.proto

package activity

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	task "github.com/uqichi/go-grpc-microservices/proto/task"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Activity struct {
	Id                   uint64               `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content              *any.Any             `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	UserId               uint64               `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Activity) Reset()         { *m = Activity{} }
func (m *Activity) String() string { return proto.CompactTextString(m) }
func (*Activity) ProtoMessage()    {}
func (*Activity) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2633de44bc656c8, []int{0}
}

func (m *Activity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Activity.Unmarshal(m, b)
}
func (m *Activity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Activity.Marshal(b, m, deterministic)
}
func (m *Activity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Activity.Merge(m, src)
}
func (m *Activity) XXX_Size() int {
	return xxx_messageInfo_Activity.Size(m)
}
func (m *Activity) XXX_DiscardUnknown() {
	xxx_messageInfo_Activity.DiscardUnknown(m)
}

var xxx_messageInfo_Activity proto.InternalMessageInfo

func (m *Activity) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Activity) GetContent() *any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Activity) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Activity) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

type CreateTaskContent struct {
	TaskId               uint64   `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskName             string   `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTaskContent) Reset()         { *m = CreateTaskContent{} }
func (m *CreateTaskContent) String() string { return proto.CompactTextString(m) }
func (*CreateTaskContent) ProtoMessage()    {}
func (*CreateTaskContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2633de44bc656c8, []int{1}
}

func (m *CreateTaskContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTaskContent.Unmarshal(m, b)
}
func (m *CreateTaskContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTaskContent.Marshal(b, m, deterministic)
}
func (m *CreateTaskContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTaskContent.Merge(m, src)
}
func (m *CreateTaskContent) XXX_Size() int {
	return xxx_messageInfo_CreateTaskContent.Size(m)
}
func (m *CreateTaskContent) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTaskContent.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTaskContent proto.InternalMessageInfo

func (m *CreateTaskContent) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *CreateTaskContent) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

type UpdateTaskStatusContent struct {
	TaskId               uint64      `protobuf:"varint,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	TaskName             string      `protobuf:"bytes,2,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskStatus           task.Status `protobuf:"varint,3,opt,name=task_status,json=taskStatus,proto3,enum=task.Status" json:"task_status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *UpdateTaskStatusContent) Reset()         { *m = UpdateTaskStatusContent{} }
func (m *UpdateTaskStatusContent) String() string { return proto.CompactTextString(m) }
func (*UpdateTaskStatusContent) ProtoMessage()    {}
func (*UpdateTaskStatusContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2633de44bc656c8, []int{2}
}

func (m *UpdateTaskStatusContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateTaskStatusContent.Unmarshal(m, b)
}
func (m *UpdateTaskStatusContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateTaskStatusContent.Marshal(b, m, deterministic)
}
func (m *UpdateTaskStatusContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateTaskStatusContent.Merge(m, src)
}
func (m *UpdateTaskStatusContent) XXX_Size() int {
	return xxx_messageInfo_UpdateTaskStatusContent.Size(m)
}
func (m *UpdateTaskStatusContent) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateTaskStatusContent.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateTaskStatusContent proto.InternalMessageInfo

func (m *UpdateTaskStatusContent) GetTaskId() uint64 {
	if m != nil {
		return m.TaskId
	}
	return 0
}

func (m *UpdateTaskStatusContent) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *UpdateTaskStatusContent) GetTaskStatus() task.Status {
	if m != nil {
		return m.TaskStatus
	}
	return task.Status_UNKNOWN
}

type CreateProjectContent struct {
	ProjectId            uint64   `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	ProjectName          string   `protobuf:"bytes,2,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateProjectContent) Reset()         { *m = CreateProjectContent{} }
func (m *CreateProjectContent) String() string { return proto.CompactTextString(m) }
func (*CreateProjectContent) ProtoMessage()    {}
func (*CreateProjectContent) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2633de44bc656c8, []int{3}
}

func (m *CreateProjectContent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateProjectContent.Unmarshal(m, b)
}
func (m *CreateProjectContent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateProjectContent.Marshal(b, m, deterministic)
}
func (m *CreateProjectContent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateProjectContent.Merge(m, src)
}
func (m *CreateProjectContent) XXX_Size() int {
	return xxx_messageInfo_CreateProjectContent.Size(m)
}
func (m *CreateProjectContent) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateProjectContent.DiscardUnknown(m)
}

var xxx_messageInfo_CreateProjectContent proto.InternalMessageInfo

func (m *CreateProjectContent) GetProjectId() uint64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *CreateProjectContent) GetProjectName() string {
	if m != nil {
		return m.ProjectName
	}
	return ""
}

type CreateActivityRequest struct {
	Content              *any.Any `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateActivityRequest) Reset()         { *m = CreateActivityRequest{} }
func (m *CreateActivityRequest) String() string { return proto.CompactTextString(m) }
func (*CreateActivityRequest) ProtoMessage()    {}
func (*CreateActivityRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2633de44bc656c8, []int{4}
}

func (m *CreateActivityRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateActivityRequest.Unmarshal(m, b)
}
func (m *CreateActivityRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateActivityRequest.Marshal(b, m, deterministic)
}
func (m *CreateActivityRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateActivityRequest.Merge(m, src)
}
func (m *CreateActivityRequest) XXX_Size() int {
	return xxx_messageInfo_CreateActivityRequest.Size(m)
}
func (m *CreateActivityRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateActivityRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateActivityRequest proto.InternalMessageInfo

func (m *CreateActivityRequest) GetContent() *any.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

type FindActivitiesResponse struct {
	Activities           []*Activity `protobuf:"bytes,1,rep,name=activities,proto3" json:"activities,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FindActivitiesResponse) Reset()         { *m = FindActivitiesResponse{} }
func (m *FindActivitiesResponse) String() string { return proto.CompactTextString(m) }
func (*FindActivitiesResponse) ProtoMessage()    {}
func (*FindActivitiesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a2633de44bc656c8, []int{5}
}

func (m *FindActivitiesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindActivitiesResponse.Unmarshal(m, b)
}
func (m *FindActivitiesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindActivitiesResponse.Marshal(b, m, deterministic)
}
func (m *FindActivitiesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindActivitiesResponse.Merge(m, src)
}
func (m *FindActivitiesResponse) XXX_Size() int {
	return xxx_messageInfo_FindActivitiesResponse.Size(m)
}
func (m *FindActivitiesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindActivitiesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindActivitiesResponse proto.InternalMessageInfo

func (m *FindActivitiesResponse) GetActivities() []*Activity {
	if m != nil {
		return m.Activities
	}
	return nil
}

func init() {
	proto.RegisterType((*Activity)(nil), "activity.Activity")
	proto.RegisterType((*CreateTaskContent)(nil), "activity.CreateTaskContent")
	proto.RegisterType((*UpdateTaskStatusContent)(nil), "activity.UpdateTaskStatusContent")
	proto.RegisterType((*CreateProjectContent)(nil), "activity.CreateProjectContent")
	proto.RegisterType((*CreateActivityRequest)(nil), "activity.CreateActivityRequest")
	proto.RegisterType((*FindActivitiesResponse)(nil), "activity.FindActivitiesResponse")
}

func init() { proto.RegisterFile("activity/activity.proto", fileDescriptor_a2633de44bc656c8) }

var fileDescriptor_a2633de44bc656c8 = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0xc7, 0xe5, 0xb4, 0x4f, 0x9a, 0x4c, 0xaa, 0x54, 0xcf, 0xaa, 0x34, 0xc6, 0x15, 0xaa, 0xf1,
	0x29, 0x97, 0xd8, 0x92, 0x39, 0x55, 0x9c, 0x42, 0x05, 0xc8, 0x08, 0x21, 0xe4, 0x16, 0x09, 0x71,
	0x89, 0x36, 0xf6, 0x60, 0x96, 0x62, 0xaf, 0xf1, 0x8e, 0x23, 0xf5, 0xc2, 0x47, 0xe1, 0xcc, 0xc7,
	0x44, 0x5e, 0x7b, 0x9d, 0xf4, 0x05, 0x09, 0x89, 0x8b, 0xe5, 0x79, 0xd1, 0xef, 0x3f, 0xff, 0x9d,
	0x5d, 0x98, 0xf1, 0x84, 0xc4, 0x46, 0xd0, 0x4d, 0x60, 0x7e, 0xfc, 0xb2, 0x92, 0x24, 0xd9, 0xc8,
	0xc4, 0xce, 0x11, 0x71, 0x75, 0x1d, 0x34, 0x9f, 0xb6, 0xe4, 0x9c, 0x65, 0x52, 0x66, 0xdf, 0x30,
	0xd0, 0xd1, 0xba, 0xfe, 0x1c, 0x90, 0xc8, 0x51, 0x11, 0xcf, 0xcb, 0xae, 0xe1, 0xf1, 0xdd, 0x06,
	0x5e, 0x74, 0x58, 0xe7, 0xf4, 0x6e, 0x09, 0xf3, 0xd2, 0x68, 0x7a, 0x3f, 0x2d, 0x18, 0x2d, 0x3b,
	0x59, 0x36, 0x85, 0x81, 0x48, 0x6d, 0xcb, 0xb5, 0xe6, 0xfb, 0xf1, 0x40, 0xa4, 0xcc, 0x87, 0x83,
	0x44, 0x16, 0x84, 0x05, 0xd9, 0x7b, 0xae, 0x35, 0x9f, 0x84, 0xc7, 0x7e, 0xcb, 0xf2, 0x0d, 0xcb,
	0x5f, 0x16, 0x37, 0xb1, 0x69, 0x62, 0x33, 0x38, 0xa8, 0x15, 0x56, 0x2b, 0x91, 0xda, 0xfb, 0x1a,
	0x32, 0x6c, 0xc2, 0x28, 0x65, 0xe7, 0x00, 0x49, 0x85, 0x9c, 0x30, 0x5d, 0x71, 0xb2, 0xff, 0xd3,
	0x2c, 0xe7, 0x1e, 0xeb, 0xca, 0x78, 0x8a, 0xc7, 0x5d, 0xf7, 0x92, 0xbc, 0x08, 0xfe, 0xbf, 0xd0,
	0xc1, 0x15, 0x57, 0xd7, 0x17, 0x5b, 0xa1, 0xe6, 0x70, 0x56, 0xfd, 0xb4, 0xc3, 0x26, 0x8c, 0x52,
	0x76, 0x0a, 0x63, 0x5d, 0x28, 0x78, 0x8e, 0xf6, 0xc0, 0xb5, 0xe6, 0xe3, 0x78, 0xd4, 0x24, 0xde,
	0xf1, 0x1c, 0xbd, 0x1f, 0x30, 0xfb, 0x50, 0xa6, 0x1d, 0xea, 0x92, 0x38, 0xd5, 0xea, 0x9f, 0x80,
	0x6c, 0x01, 0x13, 0x5d, 0x54, 0x9a, 0xa5, 0xcf, 0x68, 0x1a, 0x1e, 0xfa, 0x7a, 0x6f, 0x2d, 0x3f,
	0x06, 0xea, 0xb5, 0xbc, 0x8f, 0x70, 0xdc, 0x5a, 0x79, 0x5f, 0xc9, 0xaf, 0x98, 0x90, 0x11, 0x7f,
	0x02, 0x50, 0xb6, 0x99, 0xad, 0xfe, 0xb8, 0xcb, 0x44, 0x29, 0x7b, 0x0a, 0x87, 0xa6, 0xbc, 0x33,
	0xc5, 0xa4, 0xcb, 0x69, 0x67, 0xaf, 0xe1, 0x51, 0x4b, 0x36, 0xab, 0x8c, 0xf1, 0x7b, 0x8d, 0x8a,
	0x76, 0x37, 0x68, 0xfd, 0xc5, 0x06, 0xbd, 0xb7, 0x70, 0xf2, 0x4a, 0x14, 0x69, 0x87, 0x11, 0xa8,
	0x62, 0x54, 0xa5, 0x2c, 0x14, 0xb2, 0x10, 0x80, 0xf7, 0x59, 0xdb, 0x72, 0xf7, 0xe6, 0x93, 0x90,
	0xf9, 0xfd, 0x0d, 0xee, 0x85, 0x77, 0xba, 0xc2, 0x5f, 0x16, 0x1c, 0x99, 0xc2, 0x25, 0x56, 0x1b,
	0x91, 0x20, 0x8b, 0x60, 0x7a, 0x7b, 0x54, 0x76, 0xb6, 0xa5, 0x3c, 0x68, 0xc2, 0x39, 0xb9, 0x37,
	0xf3, 0xcb, 0xe6, 0x06, 0xb3, 0x37, 0x30, 0xbd, 0x3d, 0x2c, 0xfb, 0x43, 0xa7, 0xe3, 0x6e, 0x25,
	0x1e, 0xb6, 0xf7, 0xe2, 0xf9, 0xa7, 0xf3, 0x4c, 0xd0, 0x97, 0x7a, 0xed, 0x27, 0x32, 0x0f, 0x36,
	0x1b, 0x4e, 0xbc, 0xe0, 0x6b, 0x0c, 0x32, 0xb9, 0xc8, 0xaa, 0x32, 0x59, 0xe4, 0x22, 0xa9, 0xa4,
	0x6a, 0x3d, 0xa8, 0xf6, 0x29, 0xf5, 0xcf, 0x77, 0x3d, 0xd4, 0xf1, 0xb3, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x2a, 0x8c, 0xde, 0xdd, 0xda, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityServiceClient interface {
	CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	FindActivities(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FindActivitiesResponse, error)
}

type activityServiceClient struct {
	cc *grpc.ClientConn
}

func NewActivityServiceClient(cc *grpc.ClientConn) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/CreateActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) FindActivities(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*FindActivitiesResponse, error) {
	out := new(FindActivitiesResponse)
	err := c.cc.Invoke(ctx, "/activity.ActivityService/FindActivities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
type ActivityServiceServer interface {
	CreateActivity(context.Context, *CreateActivityRequest) (*empty.Empty, error)
	FindActivities(context.Context, *empty.Empty) (*FindActivitiesResponse, error)
}

// UnimplementedActivityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (*UnimplementedActivityServiceServer) CreateActivity(ctx context.Context, req *CreateActivityRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (*UnimplementedActivityServiceServer) FindActivities(ctx context.Context, req *empty.Empty) (*FindActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindActivities not implemented")
}

func RegisterActivityServiceServer(s *grpc.Server, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/CreateActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_FindActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).FindActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/activity.ActivityService/FindActivities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).FindActivities(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "activity.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateActivity",
			Handler:    _ActivityService_CreateActivity_Handler,
		},
		{
			MethodName: "FindActivities",
			Handler:    _ActivityService_FindActivities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "activity/activity.proto",
}
