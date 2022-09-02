// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: config.proto

package proto

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

type OptionType int32

const (
	OptionType_OPTION_TYPE_ANY          OptionType = 0
	OptionType_OPTION_TYPE_STRING       OptionType = 1
	OptionType_OPTION_TYPE_STRING_ARRAY OptionType = 2
	OptionType_OPTION_TYPE_INT          OptionType = 3
	OptionType_OPTION_TYPE_BOOL         OptionType = 4
)

// Enum value maps for OptionType.
var (
	OptionType_name = map[int32]string{
		0: "OPTION_TYPE_ANY",
		1: "OPTION_TYPE_STRING",
		2: "OPTION_TYPE_STRING_ARRAY",
		3: "OPTION_TYPE_INT",
		4: "OPTION_TYPE_BOOL",
	}
	OptionType_value = map[string]int32{
		"OPTION_TYPE_ANY":          0,
		"OPTION_TYPE_STRING":       1,
		"OPTION_TYPE_STRING_ARRAY": 2,
		"OPTION_TYPE_INT":          3,
		"OPTION_TYPE_BOOL":         4,
	}
)

func (x OptionType) Enum() *OptionType {
	p := new(OptionType)
	*p = x
	return p
}

func (x OptionType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OptionType) Descriptor() protoreflect.EnumDescriptor {
	return file_config_proto_enumTypes[0].Descriptor()
}

func (OptionType) Type() protoreflect.EnumType {
	return &file_config_proto_enumTypes[0]
}

func (x OptionType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OptionType.Descriptor instead.
func (OptionType) EnumDescriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int         int64    `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
	Bool        bool     `protobuf:"varint,2,opt,name=bool,proto3" json:"bool,omitempty"`
	Float       float32  `protobuf:"fixed32,3,opt,name=float,proto3" json:"float,omitempty"`
	String_     string   `protobuf:"bytes,4,opt,name=string,proto3" json:"string,omitempty"`
	StringArray []string `protobuf:"bytes,5,rep,name=string_array,json=stringArray,proto3" json:"string_array,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *Value) GetInt() int64 {
	if x != nil {
		return x.Int
	}
	return 0
}

func (x *Value) GetBool() bool {
	if x != nil {
		return x.Bool
	}
	return false
}

func (x *Value) GetFloat() float32 {
	if x != nil {
		return x.Float
	}
	return 0
}

func (x *Value) GetString_() string {
	if x != nil {
		return x.String_
	}
	return ""
}

func (x *Value) GetStringArray() []string {
	if x != nil {
		return x.StringArray
	}
	return nil
}

type Option struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Help        string            `protobuf:"bytes,2,opt,name=help,proto3" json:"help,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Key         string            `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty"`
	Default     *Value            `protobuf:"bytes,5,opt,name=default,proto3" json:"default,omitempty"`
	OptionType  OptionType        `protobuf:"varint,6,opt,name=option_type,json=optionType,proto3,enum=safing.portmaster.plugin.proto.OptionType" json:"option_type,omitempty"`
	Annotations map[string]*Value `protobuf:"bytes,7,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Option) Reset() {
	*x = Option{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Option) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Option) ProtoMessage() {}

func (x *Option) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Option.ProtoReflect.Descriptor instead.
func (*Option) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *Option) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Option) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *Option) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Option) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Option) GetDefault() *Value {
	if x != nil {
		return x.Default
	}
	return nil
}

func (x *Option) GetOptionType() OptionType {
	if x != nil {
		return x.OptionType
	}
	return OptionType_OPTION_TYPE_ANY
}

func (x *Option) GetAnnotations() map[string]*Value {
	if x != nil {
		return x.Annotations
	}
	return nil
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Option *Option `protobuf:"bytes,1,opt,name=option,proto3" json:"option,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetOption() *Option {
	if x != nil {
		return x.Option
	}
	return nil
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{3}
}

type GetValueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *GetValueRequest) Reset() {
	*x = GetValueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueRequest) ProtoMessage() {}

func (x *GetValueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueRequest.ProtoReflect.Descriptor instead.
func (*GetValueRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetValueRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetValueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *Value `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *GetValueResponse) Reset() {
	*x = GetValueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetValueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetValueResponse) ProtoMessage() {}

func (x *GetValueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetValueResponse.ProtoReflect.Descriptor instead.
func (*GetValueResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetValueResponse) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

type WatchChangesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Keys []string `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
}

func (x *WatchChangesRequest) Reset() {
	*x = WatchChangesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchChangesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchChangesRequest) ProtoMessage() {}

func (x *WatchChangesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchChangesRequest.ProtoReflect.Descriptor instead.
func (*WatchChangesRequest) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{6}
}

func (x *WatchChangesRequest) GetKeys() []string {
	if x != nil {
		return x.Keys
	}
	return nil
}

type WatchChangesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *Value `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WatchChangesResponse) Reset() {
	*x = WatchChangesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WatchChangesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WatchChangesResponse) ProtoMessage() {}

func (x *WatchChangesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WatchChangesResponse.ProtoReflect.Descriptor instead.
func (*WatchChangesResponse) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{7}
}

func (x *WatchChangesResponse) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *WatchChangesResponse) GetValue() *Value {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e,
	0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x7e,
	0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x6f,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x62, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a,
	0x05, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x21, 0x0a, 0x0c, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x41, 0x72, 0x72, 0x61, 0x79, 0x22, 0xb4,
	0x03, 0x0a, 0x06, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x65, 0x6c,
	0x70, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3f, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x4b, 0x0a, 0x0b, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x73, 0x61,
	0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0a, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x59, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x65,
	0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72,
	0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x51, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x06, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x0f,
	0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x4f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f,
	0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0x29, 0x0a, 0x13, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x65, 0x79,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x65, 0x79, 0x73, 0x22, 0x65, 0x0a,
	0x14, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e,
	0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x2a, 0x82, 0x01, 0x0a, 0x0a, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x41, 0x4e, 0x59, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x4f, 0x50, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x1c, 0x0a, 0x18, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x02, 0x12, 0x13,
	0x0a, 0x0f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e,
	0x54, 0x10, 0x03, 0x12, 0x14, 0x0a, 0x10, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x04, 0x32, 0xf5, 0x02, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x75, 0x0a, 0x0e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2f, 0x2e,
	0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65,
	0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x6f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x2f,
	0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74,
	0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x30, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x47, 0x65, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x0b, 0x57, 0x61, 0x74, 0x63, 0x68, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x12, 0x33, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x6f, 0x72, 0x74,
	0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x34, 0x2e, 0x73, 0x61, 0x66, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x6f, 0x72, 0x74, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x6c, 0x75, 0x67,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x61, 0x74, 0x63, 0x68, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_config_proto_goTypes = []interface{}{
	(OptionType)(0),              // 0: safing.portmaster.plugin.proto.OptionType
	(*Value)(nil),                // 1: safing.portmaster.plugin.proto.Value
	(*Option)(nil),               // 2: safing.portmaster.plugin.proto.Option
	(*RegisterRequest)(nil),      // 3: safing.portmaster.plugin.proto.RegisterRequest
	(*RegisterResponse)(nil),     // 4: safing.portmaster.plugin.proto.RegisterResponse
	(*GetValueRequest)(nil),      // 5: safing.portmaster.plugin.proto.GetValueRequest
	(*GetValueResponse)(nil),     // 6: safing.portmaster.plugin.proto.GetValueResponse
	(*WatchChangesRequest)(nil),  // 7: safing.portmaster.plugin.proto.WatchChangesRequest
	(*WatchChangesResponse)(nil), // 8: safing.portmaster.plugin.proto.WatchChangesResponse
	nil,                          // 9: safing.portmaster.plugin.proto.Option.AnnotationsEntry
}
var file_config_proto_depIdxs = []int32{
	1,  // 0: safing.portmaster.plugin.proto.Option.default:type_name -> safing.portmaster.plugin.proto.Value
	0,  // 1: safing.portmaster.plugin.proto.Option.option_type:type_name -> safing.portmaster.plugin.proto.OptionType
	9,  // 2: safing.portmaster.plugin.proto.Option.annotations:type_name -> safing.portmaster.plugin.proto.Option.AnnotationsEntry
	2,  // 3: safing.portmaster.plugin.proto.RegisterRequest.option:type_name -> safing.portmaster.plugin.proto.Option
	1,  // 4: safing.portmaster.plugin.proto.GetValueResponse.value:type_name -> safing.portmaster.plugin.proto.Value
	1,  // 5: safing.portmaster.plugin.proto.WatchChangesResponse.value:type_name -> safing.portmaster.plugin.proto.Value
	1,  // 6: safing.portmaster.plugin.proto.Option.AnnotationsEntry.value:type_name -> safing.portmaster.plugin.proto.Value
	3,  // 7: safing.portmaster.plugin.proto.ConfigService.RegisterOption:input_type -> safing.portmaster.plugin.proto.RegisterRequest
	5,  // 8: safing.portmaster.plugin.proto.ConfigService.GetValue:input_type -> safing.portmaster.plugin.proto.GetValueRequest
	7,  // 9: safing.portmaster.plugin.proto.ConfigService.WatchValues:input_type -> safing.portmaster.plugin.proto.WatchChangesRequest
	4,  // 10: safing.portmaster.plugin.proto.ConfigService.RegisterOption:output_type -> safing.portmaster.plugin.proto.RegisterResponse
	6,  // 11: safing.portmaster.plugin.proto.ConfigService.GetValue:output_type -> safing.portmaster.plugin.proto.GetValueResponse
	8,  // 12: safing.portmaster.plugin.proto.ConfigService.WatchValues:output_type -> safing.portmaster.plugin.proto.WatchChangesResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Option); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
		file_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueRequest); i {
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
		file_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetValueResponse); i {
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
		file_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchChangesRequest); i {
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
		file_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WatchChangesResponse); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		EnumInfos:         file_config_proto_enumTypes,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}