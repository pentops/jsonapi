// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: test/v1/test.proto

package testpb

import (
	_ "github.com/pentops/sugar-go/v1/sugar_pb"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Enum int32

const (
	Enum_ENUM_UNSPECIFIED Enum = 0
	Enum_ENUM_VALUE1      Enum = 1
	Enum_ENUM_VALUE2      Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ENUM_UNSPECIFIED",
		1: "ENUM_VALUE1",
		2: "ENUM_VALUE2",
	}
	Enum_value = map[string]int32{
		"ENUM_UNSPECIFIED": 0,
		"ENUM_VALUE1":      1,
		"ENUM_VALUE2":      2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_test_v1_test_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_test_v1_test_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{0}
}

type GetFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Number  int64     `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
	Numbers []float32 `protobuf:"fixed32,3,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
}

func (x *GetFooRequest) Reset() {
	*x = GetFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooRequest) ProtoMessage() {}

func (x *GetFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooRequest.ProtoReflect.Descriptor instead.
func (*GetFooRequest) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{0}
}

func (x *GetFooRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetFooRequest) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *GetFooRequest) GetNumbers() []float32 {
	if x != nil {
		return x.Numbers
	}
	return nil
}

type GetFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *GetFooResponse) Reset() {
	*x = GetFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFooResponse) ProtoMessage() {}

func (x *GetFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFooResponse.ProtoReflect.Descriptor instead.
func (*GetFooResponse) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{1}
}

func (x *GetFooResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetFooResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFooResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type PostFooRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SString         string                   `protobuf:"bytes,1,opt,name=s_string,json=sString,proto3" json:"s_string,omitempty"`
	OString         *string                  `protobuf:"bytes,2,opt,name=o_string,json=oString,proto3,oneof" json:"o_string,omitempty"`
	RString         []string                 `protobuf:"bytes,3,rep,name=r_string,json=rString,proto3" json:"r_string,omitempty"`
	SFloat          float32                  `protobuf:"fixed32,4,opt,name=s_float,json=sFloat,proto3" json:"s_float,omitempty"`
	OFloat          *float32                 `protobuf:"fixed32,5,opt,name=o_float,json=oFloat,proto3,oneof" json:"o_float,omitempty"`
	RFloat          []float32                `protobuf:"fixed32,6,rep,packed,name=r_float,json=rFloat,proto3" json:"r_float,omitempty"`
	Ts              *timestamppb.Timestamp   `protobuf:"bytes,7,opt,name=ts,proto3" json:"ts,omitempty"`
	RTs             []*timestamppb.Timestamp `protobuf:"bytes,8,rep,name=r_ts,json=rTs,proto3" json:"r_ts,omitempty"`
	SBar            *Bar                     `protobuf:"bytes,9,opt,name=s_bar,json=sBar,proto3" json:"s_bar,omitempty"`
	RBars           []*Bar                   `protobuf:"bytes,10,rep,name=r_bars,json=rBars,proto3" json:"r_bars,omitempty"`
	Enum            Enum                     `protobuf:"varint,11,opt,name=enum,proto3,enum=test.v1.Enum" json:"enum,omitempty"`
	REnum           []Enum                   `protobuf:"varint,12,rep,packed,name=r_enum,json=rEnum,proto3,enum=test.v1.Enum" json:"r_enum,omitempty"`
	SBytes          []byte                   `protobuf:"bytes,13,opt,name=s_bytes,json=sBytes,proto3" json:"s_bytes,omitempty"`
	RBytes          [][]byte                 `protobuf:"bytes,14,rep,name=r_bytes,json=rBytes,proto3" json:"r_bytes,omitempty"`
	MapStringString map[string]string        `protobuf:"bytes,15,rep,name=map_string_string,json=mapStringString,proto3" json:"map_string_string,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are assignable to NakedOneof:
	//
	//	*PostFooRequest_OneofString
	//	*PostFooRequest_OneBar
	//	*PostFooRequest_OneofFloat
	//	*PostFooRequest_OneofEnum
	NakedOneof   isPostFooRequest_NakedOneof `protobuf_oneof:"naked_oneof"`
	WrappedOneof *WrappedOneof               `protobuf:"bytes,16,opt,name=wrapped_oneof,json=wrappedOneof,proto3" json:"wrapped_oneof,omitempty"`
}

func (x *PostFooRequest) Reset() {
	*x = PostFooRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFooRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFooRequest) ProtoMessage() {}

func (x *PostFooRequest) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFooRequest.ProtoReflect.Descriptor instead.
func (*PostFooRequest) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{2}
}

func (x *PostFooRequest) GetSString() string {
	if x != nil {
		return x.SString
	}
	return ""
}

func (x *PostFooRequest) GetOString() string {
	if x != nil && x.OString != nil {
		return *x.OString
	}
	return ""
}

func (x *PostFooRequest) GetRString() []string {
	if x != nil {
		return x.RString
	}
	return nil
}

func (x *PostFooRequest) GetSFloat() float32 {
	if x != nil {
		return x.SFloat
	}
	return 0
}

func (x *PostFooRequest) GetOFloat() float32 {
	if x != nil && x.OFloat != nil {
		return *x.OFloat
	}
	return 0
}

func (x *PostFooRequest) GetRFloat() []float32 {
	if x != nil {
		return x.RFloat
	}
	return nil
}

func (x *PostFooRequest) GetTs() *timestamppb.Timestamp {
	if x != nil {
		return x.Ts
	}
	return nil
}

func (x *PostFooRequest) GetRTs() []*timestamppb.Timestamp {
	if x != nil {
		return x.RTs
	}
	return nil
}

func (x *PostFooRequest) GetSBar() *Bar {
	if x != nil {
		return x.SBar
	}
	return nil
}

func (x *PostFooRequest) GetRBars() []*Bar {
	if x != nil {
		return x.RBars
	}
	return nil
}

func (x *PostFooRequest) GetEnum() Enum {
	if x != nil {
		return x.Enum
	}
	return Enum_ENUM_UNSPECIFIED
}

func (x *PostFooRequest) GetREnum() []Enum {
	if x != nil {
		return x.REnum
	}
	return nil
}

func (x *PostFooRequest) GetSBytes() []byte {
	if x != nil {
		return x.SBytes
	}
	return nil
}

func (x *PostFooRequest) GetRBytes() [][]byte {
	if x != nil {
		return x.RBytes
	}
	return nil
}

func (x *PostFooRequest) GetMapStringString() map[string]string {
	if x != nil {
		return x.MapStringString
	}
	return nil
}

func (m *PostFooRequest) GetNakedOneof() isPostFooRequest_NakedOneof {
	if m != nil {
		return m.NakedOneof
	}
	return nil
}

func (x *PostFooRequest) GetOneofString() string {
	if x, ok := x.GetNakedOneof().(*PostFooRequest_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (x *PostFooRequest) GetOneBar() *Bar {
	if x, ok := x.GetNakedOneof().(*PostFooRequest_OneBar); ok {
		return x.OneBar
	}
	return nil
}

func (x *PostFooRequest) GetOneofFloat() float32 {
	if x, ok := x.GetNakedOneof().(*PostFooRequest_OneofFloat); ok {
		return x.OneofFloat
	}
	return 0
}

func (x *PostFooRequest) GetOneofEnum() Enum {
	if x, ok := x.GetNakedOneof().(*PostFooRequest_OneofEnum); ok {
		return x.OneofEnum
	}
	return Enum_ENUM_UNSPECIFIED
}

func (x *PostFooRequest) GetWrappedOneof() *WrappedOneof {
	if x != nil {
		return x.WrappedOneof
	}
	return nil
}

type isPostFooRequest_NakedOneof interface {
	isPostFooRequest_NakedOneof()
}

type PostFooRequest_OneofString struct {
	OneofString string `protobuf:"bytes,100,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type PostFooRequest_OneBar struct {
	OneBar *Bar `protobuf:"bytes,101,opt,name=one_bar,json=oneBar,proto3,oneof"`
}

type PostFooRequest_OneofFloat struct {
	OneofFloat float32 `protobuf:"fixed32,102,opt,name=oneof_float,json=oneofFloat,proto3,oneof"`
}

type PostFooRequest_OneofEnum struct {
	OneofEnum Enum `protobuf:"varint,103,opt,name=oneof_enum,json=oneofEnum,proto3,enum=test.v1.Enum,oneof"`
}

func (*PostFooRequest_OneofString) isPostFooRequest_NakedOneof() {}

func (*PostFooRequest_OneBar) isPostFooRequest_NakedOneof() {}

func (*PostFooRequest_OneofFloat) isPostFooRequest_NakedOneof() {}

func (*PostFooRequest_OneofEnum) isPostFooRequest_NakedOneof() {}

type WrappedOneof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//
	//	*WrappedOneof_OneofString
	//	*WrappedOneof_OneBar
	//	*WrappedOneof_OneofFloat
	//	*WrappedOneof_OneofEnum
	Type isWrappedOneof_Type `protobuf_oneof:"type"`
}

func (x *WrappedOneof) Reset() {
	*x = WrappedOneof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WrappedOneof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WrappedOneof) ProtoMessage() {}

func (x *WrappedOneof) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WrappedOneof.ProtoReflect.Descriptor instead.
func (*WrappedOneof) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{3}
}

func (m *WrappedOneof) GetType() isWrappedOneof_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *WrappedOneof) GetOneofString() string {
	if x, ok := x.GetType().(*WrappedOneof_OneofString); ok {
		return x.OneofString
	}
	return ""
}

func (x *WrappedOneof) GetOneBar() *Bar {
	if x, ok := x.GetType().(*WrappedOneof_OneBar); ok {
		return x.OneBar
	}
	return nil
}

func (x *WrappedOneof) GetOneofFloat() float32 {
	if x, ok := x.GetType().(*WrappedOneof_OneofFloat); ok {
		return x.OneofFloat
	}
	return 0
}

func (x *WrappedOneof) GetOneofEnum() Enum {
	if x, ok := x.GetType().(*WrappedOneof_OneofEnum); ok {
		return x.OneofEnum
	}
	return Enum_ENUM_UNSPECIFIED
}

type isWrappedOneof_Type interface {
	isWrappedOneof_Type()
}

type WrappedOneof_OneofString struct {
	OneofString string `protobuf:"bytes,100,opt,name=oneof_string,json=oneofString,proto3,oneof"`
}

type WrappedOneof_OneBar struct {
	OneBar *Bar `protobuf:"bytes,101,opt,name=one_bar,json=oneBar,proto3,oneof"`
}

type WrappedOneof_OneofFloat struct {
	OneofFloat float32 `protobuf:"fixed32,102,opt,name=oneof_float,json=oneofFloat,proto3,oneof"`
}

type WrappedOneof_OneofEnum struct {
	OneofEnum Enum `protobuf:"varint,103,opt,name=oneof_enum,json=oneofEnum,proto3,enum=test.v1.Enum,oneof"`
}

func (*WrappedOneof_OneofString) isWrappedOneof_Type() {}

func (*WrappedOneof_OneBar) isWrappedOneof_Type() {}

func (*WrappedOneof_OneofFloat) isWrappedOneof_Type() {}

func (*WrappedOneof_OneofEnum) isWrappedOneof_Type() {}

type Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *Bar) Reset() {
	*x = Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Bar) ProtoMessage() {}

func (x *Bar) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Bar.ProtoReflect.Descriptor instead.
func (*Bar) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{4}
}

func (x *Bar) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Bar) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Bar) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type PostFooResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *PostFooResponse) Reset() {
	*x = PostFooResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostFooResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostFooResponse) ProtoMessage() {}

func (x *PostFooResponse) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostFooResponse.ProtoReflect.Descriptor instead.
func (*PostFooResponse) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{5}
}

func (x *PostFooResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PostFooResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PostFooResponse) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

type FooMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Field string `protobuf:"bytes,3,opt,name=field,proto3" json:"field,omitempty"`
}

func (x *FooMessage) Reset() {
	*x = FooMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_v1_test_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FooMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FooMessage) ProtoMessage() {}

func (x *FooMessage) ProtoReflect() protoreflect.Message {
	mi := &file_test_v1_test_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FooMessage.ProtoReflect.Descriptor instead.
func (*FooMessage) Descriptor() ([]byte, []int) {
	return file_test_v1_test_proto_rawDescGZIP(), []int{6}
}

func (x *FooMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *FooMessage) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FooMessage) GetField() string {
	if x != nil {
		return x.Field
	}
	return ""
}

var File_test_v1_test_proto protoreflect.FileDescriptor

var file_test_v1_test_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x16, 0x73, 0x75, 0x67, 0x61, 0x72,
	0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x51, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x02, 0x52, 0x07, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x22, 0x4a, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x22, 0xf7, 0x06, 0x0a, 0x0e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x1e,
	0x0a, 0x08, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x01, 0x52, 0x07, 0x6f, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x12, 0x19,
	0x0a, 0x08, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x72, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x66,
	0x6c, 0x6f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x73, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x1c, 0x0a, 0x07, 0x6f, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x02, 0x48, 0x02, 0x52, 0x06, 0x6f, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x88, 0x01, 0x01,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x02, 0x52, 0x06, 0x72, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2a, 0x0a, 0x02, 0x74, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x02, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x04, 0x72, 0x5f, 0x74, 0x73, 0x18, 0x08, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x03, 0x72, 0x54, 0x73, 0x12, 0x21, 0x0a, 0x05, 0x73, 0x5f, 0x62, 0x61, 0x72, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61,
	0x72, 0x52, 0x04, 0x73, 0x42, 0x61, 0x72, 0x12, 0x23, 0x0a, 0x06, 0x72, 0x5f, 0x62, 0x61, 0x72,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x72, 0x52, 0x05, 0x72, 0x42, 0x61, 0x72, 0x73, 0x12, 0x21, 0x0a, 0x04,
	0x65, 0x6e, 0x75, 0x6d, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x12,
	0x24, 0x0a, 0x06, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0e, 0x32,
	0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x05,
	0x72, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x73, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x17,
	0x0a, 0x07, 0x72, 0x5f, 0x62, 0x79, 0x74, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x06, 0x72, 0x42, 0x79, 0x74, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x11, 0x6d, 0x61, 0x70, 0x5f, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x0f, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73,
	0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x61, 0x70, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x0f, 0x6d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x23, 0x0a, 0x0c, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x27, 0x0a, 0x07, 0x6f, 0x6e, 0x65, 0x5f, 0x62, 0x61,
	0x72, 0x18, 0x65, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x42, 0x61, 0x72, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x6e, 0x65, 0x42, 0x61, 0x72, 0x12,
	0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x66,
	0x20, 0x01, 0x28, 0x02, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46, 0x6c, 0x6f,
	0x61, 0x74, 0x12, 0x2e, 0x0a, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d,
	0x18, 0x67, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48, 0x00, 0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x3a, 0x0a, 0x0d, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x5f, 0x6f, 0x6e,
	0x65, 0x6f, 0x66, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66,
	0x52, 0x0c, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x1a, 0x42,
	0x0a, 0x14, 0x4d, 0x61, 0x70, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02,
	0x38, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x6e, 0x61, 0x6b, 0x65, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x6f,
	0x66, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6f, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x6f, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x22, 0xc1, 0x01, 0x0a, 0x0c, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x23, 0x0a, 0x0c, 0x6f,
	0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x27, 0x0a, 0x07, 0x6f, 0x6e, 0x65, 0x5f, 0x62, 0x61, 0x72, 0x18, 0x65, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0c, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x72, 0x48,
	0x00, 0x52, 0x06, 0x6f, 0x6e, 0x65, 0x42, 0x61, 0x72, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x18, 0x66, 0x20, 0x01, 0x28, 0x02, 0x48, 0x00,
	0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x2e, 0x0a, 0x0a,
	0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x18, 0x67, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0d, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x48,
	0x00, 0x52, 0x09, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x45, 0x6e, 0x75, 0x6d, 0x3a, 0x08, 0xd2, 0xa5,
	0xf5, 0xe4, 0x02, 0x02, 0x08, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x3f,
	0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22,
	0x4b, 0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x46, 0x0a, 0x0a,
	0x46, 0x6f, 0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x2a, 0x3e, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x14, 0x0a, 0x10,
	0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45,
	0x31, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x4e, 0x55, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x32, 0x10, 0x02, 0x32, 0xb9, 0x01, 0x0a, 0x0a, 0x46, 0x6f, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x12, 0x16, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x19,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x66, 0x6f, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x55, 0x0a, 0x07, 0x50, 0x6f, 0x73,
	0x74, 0x46, 0x6f, 0x6f, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x6f, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x46, 0x6f, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a,
	0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x6f, 0x6f,
	0x32, 0x40, 0x0a, 0x08, 0x46, 0x6f, 0x6f, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x34, 0x0a, 0x03,
	0x46, 0x6f, 0x6f, 0x12, 0x13, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x6f,
	0x6f, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x22, 0x00, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_v1_test_proto_rawDescOnce sync.Once
	file_test_v1_test_proto_rawDescData = file_test_v1_test_proto_rawDesc
)

func file_test_v1_test_proto_rawDescGZIP() []byte {
	file_test_v1_test_proto_rawDescOnce.Do(func() {
		file_test_v1_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_v1_test_proto_rawDescData)
	})
	return file_test_v1_test_proto_rawDescData
}

var file_test_v1_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_test_v1_test_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_test_v1_test_proto_goTypes = []interface{}{
	(Enum)(0),                     // 0: test.v1.Enum
	(*GetFooRequest)(nil),         // 1: test.v1.GetFooRequest
	(*GetFooResponse)(nil),        // 2: test.v1.GetFooResponse
	(*PostFooRequest)(nil),        // 3: test.v1.PostFooRequest
	(*WrappedOneof)(nil),          // 4: test.v1.WrappedOneof
	(*Bar)(nil),                   // 5: test.v1.Bar
	(*PostFooResponse)(nil),       // 6: test.v1.PostFooResponse
	(*FooMessage)(nil),            // 7: test.v1.FooMessage
	nil,                           // 8: test.v1.PostFooRequest.MapStringStringEntry
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 10: google.protobuf.Empty
}
var file_test_v1_test_proto_depIdxs = []int32{
	9,  // 0: test.v1.PostFooRequest.ts:type_name -> google.protobuf.Timestamp
	9,  // 1: test.v1.PostFooRequest.r_ts:type_name -> google.protobuf.Timestamp
	5,  // 2: test.v1.PostFooRequest.s_bar:type_name -> test.v1.Bar
	5,  // 3: test.v1.PostFooRequest.r_bars:type_name -> test.v1.Bar
	0,  // 4: test.v1.PostFooRequest.enum:type_name -> test.v1.Enum
	0,  // 5: test.v1.PostFooRequest.r_enum:type_name -> test.v1.Enum
	8,  // 6: test.v1.PostFooRequest.map_string_string:type_name -> test.v1.PostFooRequest.MapStringStringEntry
	5,  // 7: test.v1.PostFooRequest.one_bar:type_name -> test.v1.Bar
	0,  // 8: test.v1.PostFooRequest.oneof_enum:type_name -> test.v1.Enum
	4,  // 9: test.v1.PostFooRequest.wrapped_oneof:type_name -> test.v1.WrappedOneof
	5,  // 10: test.v1.WrappedOneof.one_bar:type_name -> test.v1.Bar
	0,  // 11: test.v1.WrappedOneof.oneof_enum:type_name -> test.v1.Enum
	1,  // 12: test.v1.FooService.GetFoo:input_type -> test.v1.GetFooRequest
	3,  // 13: test.v1.FooService.PostFoo:input_type -> test.v1.PostFooRequest
	7,  // 14: test.v1.FooTopic.Foo:input_type -> test.v1.FooMessage
	2,  // 15: test.v1.FooService.GetFoo:output_type -> test.v1.GetFooResponse
	6,  // 16: test.v1.FooService.PostFoo:output_type -> test.v1.PostFooResponse
	10, // 17: test.v1.FooTopic.Foo:output_type -> google.protobuf.Empty
	15, // [15:18] is the sub-list for method output_type
	12, // [12:15] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_test_v1_test_proto_init() }
func file_test_v1_test_proto_init() {
	if File_test_v1_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_v1_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooRequest); i {
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
		file_test_v1_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFooResponse); i {
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
		file_test_v1_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFooRequest); i {
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
		file_test_v1_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WrappedOneof); i {
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
		file_test_v1_test_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Bar); i {
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
		file_test_v1_test_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostFooResponse); i {
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
		file_test_v1_test_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FooMessage); i {
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
	file_test_v1_test_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*PostFooRequest_OneofString)(nil),
		(*PostFooRequest_OneBar)(nil),
		(*PostFooRequest_OneofFloat)(nil),
		(*PostFooRequest_OneofEnum)(nil),
	}
	file_test_v1_test_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*WrappedOneof_OneofString)(nil),
		(*WrappedOneof_OneBar)(nil),
		(*WrappedOneof_OneofFloat)(nil),
		(*WrappedOneof_OneofEnum)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_test_v1_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_test_v1_test_proto_goTypes,
		DependencyIndexes: file_test_v1_test_proto_depIdxs,
		EnumInfos:         file_test_v1_test_proto_enumTypes,
		MessageInfos:      file_test_v1_test_proto_msgTypes,
	}.Build()
	File_test_v1_test_proto = out.File
	file_test_v1_test_proto_rawDesc = nil
	file_test_v1_test_proto_goTypes = nil
	file_test_v1_test_proto_depIdxs = nil
}
