// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.20.0-devel
// 	protoc        v3.10.1
// source: internal/errors.proto

package internal

import (
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Error is the generic error returned from unary RPCs.
type Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	// This is to make the error more compatible with users that expect errors to be Status objects:
	// https://github.com/grpc/grpc/blob/master/src/proto/grpc/status/status.proto
	// It should be the exact same message as the Error field.
	Code    int32      `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Message string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Details []*any.Any `protobuf:"bytes,4,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Error) Reset() {
	*x = Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Error) ProtoMessage() {}

func (x *Error) ProtoReflect() protoreflect.Message {
	mi := &file_internal_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Error.ProtoReflect.Descriptor instead.
func (*Error) Descriptor() ([]byte, []int) {
	return file_internal_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Error) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Error) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Error) GetDetails() []*any.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

// StreamError is a response type which is returned when
// streaming rpc returns an error.
type StreamError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GrpcCode   int32      `protobuf:"varint,1,opt,name=grpc_code,json=grpcCode,proto3" json:"grpc_code,omitempty"`
	HttpCode   int32      `protobuf:"varint,2,opt,name=http_code,json=httpCode,proto3" json:"http_code,omitempty"`
	Message    string     `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	HttpStatus string     `protobuf:"bytes,4,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	Details    []*any.Any `protobuf:"bytes,5,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *StreamError) Reset() {
	*x = StreamError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_errors_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamError) ProtoMessage() {}

func (x *StreamError) ProtoReflect() protoreflect.Message {
	mi := &file_internal_errors_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamError.ProtoReflect.Descriptor instead.
func (*StreamError) Descriptor() ([]byte, []int) {
	return file_internal_errors_proto_rawDescGZIP(), []int{1}
}

func (x *StreamError) GetGrpcCode() int32 {
	if x != nil {
		return x.GrpcCode
	}
	return 0
}

func (x *StreamError) GetHttpCode() int32 {
	if x != nil {
		return x.HttpCode
	}
	return 0
}

func (x *StreamError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *StreamError) GetHttpStatus() string {
	if x != nil {
		return x.HttpStatus
	}
	return ""
}

func (x *StreamError) GetDetails() []*any.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_internal_errors_proto protoreflect.FileDescriptor

var file_internal_errors_proto_rawDesc = []byte{
	0x0a, 0x15, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x1a, 0x19, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61,
	0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7b, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0xb2, 0x01, 0x0a, 0x0b, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x67, 0x72, 0x70, 0x63, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x68, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x74, 0x74,
	0x70, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x68, 0x74, 0x74, 0x70, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x07, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0x0a, 0x5a, 0x08, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_errors_proto_rawDescOnce sync.Once
	file_internal_errors_proto_rawDescData = file_internal_errors_proto_rawDesc
)

func file_internal_errors_proto_rawDescGZIP() []byte {
	file_internal_errors_proto_rawDescOnce.Do(func() {
		file_internal_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_errors_proto_rawDescData)
	})
	return file_internal_errors_proto_rawDescData
}

var file_internal_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_errors_proto_goTypes = []interface{}{
	(*Error)(nil),       // 0: grpc.gateway.runtime.Error
	(*StreamError)(nil), // 1: grpc.gateway.runtime.StreamError
	(*any.Any)(nil),     // 2: google.protobuf.Any
}
var file_internal_errors_proto_depIdxs = []int32{
	2, // 0: grpc.gateway.runtime.Error.details:type_name -> google.protobuf.Any
	2, // 1: grpc.gateway.runtime.StreamError.details:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_errors_proto_init() }
func file_internal_errors_proto_init() {
	if File_internal_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Error); i {
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
		file_internal_errors_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamError); i {
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
			RawDescriptor: file_internal_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_errors_proto_goTypes,
		DependencyIndexes: file_internal_errors_proto_depIdxs,
		MessageInfos:      file_internal_errors_proto_msgTypes,
	}.Build()
	File_internal_errors_proto = out.File
	file_internal_errors_proto_rawDesc = nil
	file_internal_errors_proto_goTypes = nil
	file_internal_errors_proto_depIdxs = nil
}