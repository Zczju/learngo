// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: types.proto

package apis

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type PersonalInformationList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*PersonalInformation `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *PersonalInformationList) Reset() {
	*x = PersonalInformationList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalInformationList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalInformationList) ProtoMessage() {}

func (x *PersonalInformationList) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalInformationList.ProtoReflect.Descriptor instead.
func (*PersonalInformationList) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{0}
}

func (x *PersonalInformationList) GetItems() []*PersonalInformation {
	if x != nil {
		return x.Items
	}
	return nil
}

type PersonalInformation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: gorm:"primaryKey; column:id" //注意这个一定要写 不要忘了Key
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// @gotags: gorm:"column:name" // 这个要手写么？ 反正必须要有 protoc-go-inject-tag -input="*.pb.go"就需要读取这个tag
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// @gotags: gorm:"column:sex"
	Sex string `protobuf:"bytes,3,opt,name=sex,proto3" json:"sex,omitempty"`
	// @gotags: gorm:"column:tall"
	Tall float32 `protobuf:"fixed32,4,opt,name=tall,proto3" json:"tall,omitempty"`
	// @gotags: gorm:"column:weight"
	Weight float32 `protobuf:"fixed32,5,opt,name=weight,proto3" json:"weight,omitempty"`
	// @gotags: gorm:"column:age"
	Age int64 `protobuf:"varint,6,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *PersonalInformation) Reset() {
	*x = PersonalInformation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PersonalInformation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PersonalInformation) ProtoMessage() {}

func (x *PersonalInformation) ProtoReflect() protoreflect.Message {
	mi := &file_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PersonalInformation.ProtoReflect.Descriptor instead.
func (*PersonalInformation) Descriptor() ([]byte, []int) {
	return file_types_proto_rawDescGZIP(), []int{1}
}

func (x *PersonalInformation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PersonalInformation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PersonalInformation) GetSex() string {
	if x != nil {
		return x.Sex
	}
	return ""
}

func (x *PersonalInformation) GetTall() float32 {
	if x != nil {
		return x.Tall
	}
	return 0
}

func (x *PersonalInformation) GetWeight() float32 {
	if x != nil {
		return x.Weight
	}
	return 0
}

func (x *PersonalInformation) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_types_proto protoreflect.FileDescriptor

var file_types_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61,
	0x70, 0x69, 0x73, 0x22, 0x4a, 0x0a, 0x17, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22,
	0x89, 0x01, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f,
	0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73,
	0x65, 0x78, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x65, 0x78, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x61, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x74, 0x61, 0x6c,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x06, 0x77, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x61, 0x67, 0x65, 0x32, 0x51, 0x0a, 0x0b, 0x52,
	0x61, 0x6e, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61,
	0x6c, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_types_proto_rawDescOnce sync.Once
	file_types_proto_rawDescData = file_types_proto_rawDesc
)

func file_types_proto_rawDescGZIP() []byte {
	file_types_proto_rawDescOnce.Do(func() {
		file_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_types_proto_rawDescData)
	})
	return file_types_proto_rawDescData
}

var file_types_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_types_proto_goTypes = []interface{}{
	(*PersonalInformationList)(nil), // 0: apis.PersonalInformationList
	(*PersonalInformation)(nil),     // 1: apis.PersonalInformation
}
var file_types_proto_depIdxs = []int32{
	1, // 0: apis.PersonalInformationList.items:type_name -> apis.PersonalInformation
	1, // 1: apis.RankService.Register:input_type -> apis.PersonalInformation
	1, // 2: apis.RankService.Register:output_type -> apis.PersonalInformation
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_types_proto_init() }
func file_types_proto_init() {
	if File_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalInformationList); i {
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
		file_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PersonalInformation); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_types_proto_goTypes,
		DependencyIndexes: file_types_proto_depIdxs,
		MessageInfos:      file_types_proto_msgTypes,
	}.Build()
	File_types_proto = out.File
	file_types_proto_rawDesc = nil
	file_types_proto_goTypes = nil
	file_types_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RankServiceClient is the client API for RankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RankServiceClient interface {
	Register(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalInformation, error)
}

type rankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRankServiceClient(cc grpc.ClientConnInterface) RankServiceClient {
	return &rankServiceClient{cc}
}

func (c *rankServiceClient) Register(ctx context.Context, in *PersonalInformation, opts ...grpc.CallOption) (*PersonalInformation, error) {
	out := new(PersonalInformation)
	err := c.cc.Invoke(ctx, "/apis.RankService/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RankServiceServer is the server API for RankService service.
type RankServiceServer interface {
	Register(context.Context, *PersonalInformation) (*PersonalInformation, error)
}

// UnimplementedRankServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRankServiceServer struct {
}

func (*UnimplementedRankServiceServer) Register(context.Context, *PersonalInformation) (*PersonalInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRankServiceServer(s *grpc.Server, srv RankServiceServer) {
	s.RegisterService(&_RankService_serviceDesc, srv)
}

func _RankService_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PersonalInformation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RankServiceServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apis.RankService/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RankServiceServer).Register(ctx, req.(*PersonalInformation))
	}
	return interceptor(ctx, in, info, handler)
}

var _RankService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "apis.RankService",
	HandlerType: (*RankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _RankService_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}