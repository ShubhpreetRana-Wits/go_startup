package services

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Define the request message for consuming a user and message
type ConsumeMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	User          string                 `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"` // User object
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Medium        string                 `protobuf:"bytes,3,opt,name=medium,proto3" json:"medium,omitempty"` // Message object
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ConsumeMessageRequest) Reset() {
	*x = ConsumeMessageRequest{}
	mi := &file_grpc_proto_send_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ConsumeMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConsumeMessageRequest) ProtoMessage() {}

func (x *ConsumeMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_send_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConsumeMessageRequest.ProtoReflect.Descriptor instead.
func (*ConsumeMessageRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_send_message_proto_rawDescGZIP(), []int{0}
}

func (x *ConsumeMessageRequest) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ConsumeMessageRequest) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ConsumeMessageRequest) GetMedium() string {
	if x != nil {
		return x.Medium
	}
	return ""
}

// Define the response message for sent message
type SentMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Status        string                 `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // Status of the message sending operation (e.g., success or failure)// Unique message ID for tracking
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SentMessageResponse) Reset() {
	*x = SentMessageResponse{}
	mi := &file_grpc_proto_send_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SentMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SentMessageResponse) ProtoMessage() {}

func (x *SentMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_send_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SentMessageResponse.ProtoReflect.Descriptor instead.
func (*SentMessageResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_send_message_proto_rawDescGZIP(), []int{1}
}

func (x *SentMessageResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_grpc_proto_send_message_proto protoreflect.FileDescriptor

var file_grpc_proto_send_message_proto_rawDesc = string([]byte{
	0x0a, 0x1d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x6e,
	0x64, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5d, 0x0a,
	0x15, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x64, 0x69, 0x75, 0x6d, 0x22, 0x2d, 0x0a, 0x13,
	0x53, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x6c, 0x0a, 0x13, 0x4e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x55, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x6e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x6e,
	0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_grpc_proto_send_message_proto_rawDescOnce sync.Once
	file_grpc_proto_send_message_proto_rawDescData []byte
)

func file_grpc_proto_send_message_proto_rawDescGZIP() []byte {
	file_grpc_proto_send_message_proto_rawDescOnce.Do(func() {
		file_grpc_proto_send_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_grpc_proto_send_message_proto_rawDesc), len(file_grpc_proto_send_message_proto_rawDesc)))
	})
	return file_grpc_proto_send_message_proto_rawDescData
}

var file_grpc_proto_send_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_send_message_proto_goTypes = []any{
	(*ConsumeMessageRequest)(nil), // 0: notification.ConsumeMessageRequest
	(*SentMessageResponse)(nil),   // 1: notification.SentMessageResponse
}
var file_grpc_proto_send_message_proto_depIdxs = []int32{
	0, // 0: notification.NotificationService.SendMessage:input_type -> notification.ConsumeMessageRequest
	1, // 1: notification.NotificationService.SendMessage:output_type -> notification.SentMessageResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpc_proto_send_message_proto_init() }
func file_grpc_proto_send_message_proto_init() {
	if File_grpc_proto_send_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_grpc_proto_send_message_proto_rawDesc), len(file_grpc_proto_send_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_send_message_proto_goTypes,
		DependencyIndexes: file_grpc_proto_send_message_proto_depIdxs,
		MessageInfos:      file_grpc_proto_send_message_proto_msgTypes,
	}.Build()
	File_grpc_proto_send_message_proto = out.File
	file_grpc_proto_send_message_proto_goTypes = nil
	file_grpc_proto_send_message_proto_depIdxs = nil
}
