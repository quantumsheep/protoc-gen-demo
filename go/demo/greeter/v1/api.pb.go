// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: demo/greeter/v1/api.proto

package greeterv1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_greeter_v1_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_greeter_v1_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_demo_greeter_v1_api_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_greeter_v1_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_greeter_v1_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_demo_greeter_v1_api_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SayGoodbyeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *SayGoodbyeRequest) Reset() {
	*x = SayGoodbyeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_greeter_v1_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodbyeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodbyeRequest) ProtoMessage() {}

func (x *SayGoodbyeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_demo_greeter_v1_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodbyeRequest.ProtoReflect.Descriptor instead.
func (*SayGoodbyeRequest) Descriptor() ([]byte, []int) {
	return file_demo_greeter_v1_api_proto_rawDescGZIP(), []int{2}
}

func (x *SayGoodbyeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type SayGoodbyeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SayGoodbyeResponse) Reset() {
	*x = SayGoodbyeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_demo_greeter_v1_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SayGoodbyeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayGoodbyeResponse) ProtoMessage() {}

func (x *SayGoodbyeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_demo_greeter_v1_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayGoodbyeResponse.ProtoReflect.Descriptor instead.
func (*SayGoodbyeResponse) Descriptor() ([]byte, []int) {
	return file_demo_greeter_v1_api_proto_rawDescGZIP(), []int{3}
}

func (x *SayGoodbyeResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_demo_greeter_v1_api_proto protoreflect.FileDescriptor

var file_demo_greeter_v1_api_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76,
	0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x64, 0x65, 0x6d,
	0x6f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0f, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x30, 0x0a, 0x11, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x2e, 0x0a, 0x12, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0xf6, 0x01, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x08, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x12, 0x20, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a,
	0x01, 0x2a, 0x22, 0x11, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x75, 0x0a, 0x0a, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64,
	0x62, 0x79, 0x65, 0x12, 0x22, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x64, 0x65, 0x6d, 0x6f, 0x2e, 0x67,
	0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x47, 0x6f, 0x6f,
	0x64, 0x62, 0x79, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x18, 0x3a, 0x01, 0x2a, 0x22, 0x13, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x65,
	0x72, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x6f, 0x64, 0x62, 0x79, 0x65, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x63, 0x61, 0x6c, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x6f, 0x2f,
	0x67, 0x72, 0x65, 0x65, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_demo_greeter_v1_api_proto_rawDescOnce sync.Once
	file_demo_greeter_v1_api_proto_rawDescData = file_demo_greeter_v1_api_proto_rawDesc
)

func file_demo_greeter_v1_api_proto_rawDescGZIP() []byte {
	file_demo_greeter_v1_api_proto_rawDescOnce.Do(func() {
		file_demo_greeter_v1_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_demo_greeter_v1_api_proto_rawDescData)
	})
	return file_demo_greeter_v1_api_proto_rawDescData
}

var file_demo_greeter_v1_api_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_demo_greeter_v1_api_proto_goTypes = []interface{}{
	(*SayHelloRequest)(nil),    // 0: demo.greeter.v1.SayHelloRequest
	(*SayHelloResponse)(nil),   // 1: demo.greeter.v1.SayHelloResponse
	(*SayGoodbyeRequest)(nil),  // 2: demo.greeter.v1.SayGoodbyeRequest
	(*SayGoodbyeResponse)(nil), // 3: demo.greeter.v1.SayGoodbyeResponse
}
var file_demo_greeter_v1_api_proto_depIdxs = []int32{
	0, // 0: demo.greeter.v1.GreeterService.SayHello:input_type -> demo.greeter.v1.SayHelloRequest
	2, // 1: demo.greeter.v1.GreeterService.SayGoodbye:input_type -> demo.greeter.v1.SayGoodbyeRequest
	1, // 2: demo.greeter.v1.GreeterService.SayHello:output_type -> demo.greeter.v1.SayHelloResponse
	3, // 3: demo.greeter.v1.GreeterService.SayGoodbye:output_type -> demo.greeter.v1.SayGoodbyeResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_demo_greeter_v1_api_proto_init() }
func file_demo_greeter_v1_api_proto_init() {
	if File_demo_greeter_v1_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_demo_greeter_v1_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloRequest); i {
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
		file_demo_greeter_v1_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayHelloResponse); i {
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
		file_demo_greeter_v1_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodbyeRequest); i {
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
		file_demo_greeter_v1_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SayGoodbyeResponse); i {
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
			RawDescriptor: file_demo_greeter_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_demo_greeter_v1_api_proto_goTypes,
		DependencyIndexes: file_demo_greeter_v1_api_proto_depIdxs,
		MessageInfos:      file_demo_greeter_v1_api_proto_msgTypes,
	}.Build()
	File_demo_greeter_v1_api_proto = out.File
	file_demo_greeter_v1_api_proto_rawDesc = nil
	file_demo_greeter_v1_api_proto_goTypes = nil
	file_demo_greeter_v1_api_proto_depIdxs = nil
}