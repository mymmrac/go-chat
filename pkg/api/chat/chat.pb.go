// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.2
// source: api/chat/chat.proto

package chat

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

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*ChatRequest_Message
	Type isChatRequest_Type `protobuf_oneof:"type"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{0}
}

func (m *ChatRequest) GetType() isChatRequest_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ChatRequest) GetMessage() string {
	if x, ok := x.GetType().(*ChatRequest_Message); ok {
		return x.Message
	}
	return ""
}

type isChatRequest_Type interface {
	isChatRequest_Type()
}

type ChatRequest_Message struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3,oneof"`
}

func (*ChatRequest_Message) isChatRequest_Type() {}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*ChatResponse_Me
	//	*ChatResponse_Connected
	//	*ChatResponse_Disconnected
	//	*ChatResponse_Error_
	//	*ChatResponse_Message_
	Type isChatResponse_Type `protobuf_oneof:"type"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{1}
}

func (m *ChatResponse) GetType() isChatResponse_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *ChatResponse) GetMe() string {
	if x, ok := x.GetType().(*ChatResponse_Me); ok {
		return x.Me
	}
	return ""
}

func (x *ChatResponse) GetConnected() string {
	if x, ok := x.GetType().(*ChatResponse_Connected); ok {
		return x.Connected
	}
	return ""
}

func (x *ChatResponse) GetDisconnected() string {
	if x, ok := x.GetType().(*ChatResponse_Disconnected); ok {
		return x.Disconnected
	}
	return ""
}

func (x *ChatResponse) GetError() *ChatResponse_Error {
	if x, ok := x.GetType().(*ChatResponse_Error_); ok {
		return x.Error
	}
	return nil
}

func (x *ChatResponse) GetMessage() *ChatResponse_Message {
	if x, ok := x.GetType().(*ChatResponse_Message_); ok {
		return x.Message
	}
	return nil
}

type isChatResponse_Type interface {
	isChatResponse_Type()
}

type ChatResponse_Me struct {
	Me string `protobuf:"bytes,1,opt,name=me,proto3,oneof"`
}

type ChatResponse_Connected struct {
	Connected string `protobuf:"bytes,2,opt,name=connected,proto3,oneof"`
}

type ChatResponse_Disconnected struct {
	Disconnected string `protobuf:"bytes,3,opt,name=disconnected,proto3,oneof"`
}

type ChatResponse_Error_ struct {
	Error *ChatResponse_Error `protobuf:"bytes,4,opt,name=error,proto3,oneof"`
}

type ChatResponse_Message_ struct {
	Message *ChatResponse_Message `protobuf:"bytes,5,opt,name=message,proto3,oneof"`
}

func (*ChatResponse_Me) isChatResponse_Type() {}

func (*ChatResponse_Connected) isChatResponse_Type() {}

func (*ChatResponse_Disconnected) isChatResponse_Type() {}

func (*ChatResponse_Error_) isChatResponse_Type() {}

func (*ChatResponse_Message_) isChatResponse_Type() {}

type ChatResponse_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User string `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Text string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *ChatResponse_Message) Reset() {
	*x = ChatResponse_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse_Message) ProtoMessage() {}

func (x *ChatResponse_Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse_Message.ProtoReflect.Descriptor instead.
func (*ChatResponse_Message) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ChatResponse_Message) GetUser() string {
	if x != nil {
		return x.User
	}
	return ""
}

func (x *ChatResponse_Message) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type ChatResponse_Error struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatResponse_Error) Reset() {
	*x = ChatResponse_Error{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_chat_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse_Error) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse_Error) ProtoMessage() {}

func (x *ChatResponse_Error) ProtoReflect() protoreflect.Message {
	mi := &file_api_chat_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse_Error.ProtoReflect.Descriptor instead.
func (*ChatResponse_Error) Descriptor() ([]byte, []int) {
	return file_api_chat_chat_proto_rawDescGZIP(), []int{1, 1}
}

func (x *ChatResponse_Error) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_chat_chat_proto protoreflect.FileDescriptor

var file_api_chat_chat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68, 0x61, 0x74, 0x22, 0x31, 0x0a, 0x0b, 0x43,
	0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xae,
	0x02, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x10, 0x0a, 0x02, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x02, 0x6d,
	0x65, 0x12, 0x1e, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x12, 0x24, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68,
	0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x48, 0x00, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x36, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x68, 0x61,
	0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x00, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x31, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x1a, 0x21, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x32,
	0x42, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x3a, 0x0a, 0x0d, 0x43, 0x6f, 0x6d, 0x6d, 0x75,
	0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x68,
	0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28,
	0x01, 0x30, 0x01, 0x42, 0x0a, 0x48, 0x01, 0x5a, 0x06, 0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_chat_chat_proto_rawDescOnce sync.Once
	file_api_chat_chat_proto_rawDescData = file_api_chat_chat_proto_rawDesc
)

func file_api_chat_chat_proto_rawDescGZIP() []byte {
	file_api_chat_chat_proto_rawDescOnce.Do(func() {
		file_api_chat_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_chat_chat_proto_rawDescData)
	})
	return file_api_chat_chat_proto_rawDescData
}

var file_api_chat_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_chat_chat_proto_goTypes = []interface{}{
	(*ChatRequest)(nil),          // 0: chat.ChatRequest
	(*ChatResponse)(nil),         // 1: chat.ChatResponse
	(*ChatResponse_Message)(nil), // 2: chat.ChatResponse.Message
	(*ChatResponse_Error)(nil),   // 3: chat.ChatResponse.Error
}
var file_api_chat_chat_proto_depIdxs = []int32{
	3, // 0: chat.ChatResponse.error:type_name -> chat.ChatResponse.Error
	2, // 1: chat.ChatResponse.message:type_name -> chat.ChatResponse.Message
	0, // 2: chat.Chat.Communication:input_type -> chat.ChatRequest
	1, // 3: chat.Chat.Communication:output_type -> chat.ChatResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_api_chat_chat_proto_init() }
func file_api_chat_chat_proto_init() {
	if File_api_chat_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_chat_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_api_chat_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
		file_api_chat_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse_Message); i {
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
		file_api_chat_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse_Error); i {
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
	file_api_chat_chat_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*ChatRequest_Message)(nil),
	}
	file_api_chat_chat_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ChatResponse_Me)(nil),
		(*ChatResponse_Connected)(nil),
		(*ChatResponse_Disconnected)(nil),
		(*ChatResponse_Error_)(nil),
		(*ChatResponse_Message_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_chat_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_chat_chat_proto_goTypes,
		DependencyIndexes: file_api_chat_chat_proto_depIdxs,
		MessageInfos:      file_api_chat_chat_proto_msgTypes,
	}.Build()
	File_api_chat_chat_proto = out.File
	file_api_chat_chat_proto_rawDesc = nil
	file_api_chat_chat_proto_goTypes = nil
	file_api_chat_chat_proto_depIdxs = nil
}
