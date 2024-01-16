// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: jsonapi/v1/image.proto

package jsonapi_pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Image struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File     []*descriptorpb.FileDescriptorProto `protobuf:"bytes,1,rep,name=file,proto3" json:"file,omitempty"`
	Packages []*PackageConfig                    `protobuf:"bytes,2,rep,name=packages,proto3" json:"packages,omitempty"`
	Prose    []*ProseFile                        `protobuf:"bytes,3,rep,name=prose,proto3" json:"prose,omitempty"`
	Codec    *CodecOptions                       `protobuf:"bytes,4,opt,name=codec,proto3" json:"codec,omitempty"`
	Registry *RegistryConfig                     `protobuf:"bytes,5,opt,name=registry,proto3" json:"registry,omitempty"`
}

func (x *Image) Reset() {
	*x = Image{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_image_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Image) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Image) ProtoMessage() {}

func (x *Image) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_image_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Image.ProtoReflect.Descriptor instead.
func (*Image) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_image_proto_rawDescGZIP(), []int{0}
}

func (x *Image) GetFile() []*descriptorpb.FileDescriptorProto {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *Image) GetPackages() []*PackageConfig {
	if x != nil {
		return x.Packages
	}
	return nil
}

func (x *Image) GetProse() []*ProseFile {
	if x != nil {
		return x.Prose
	}
	return nil
}

func (x *Image) GetCodec() *CodecOptions {
	if x != nil {
		return x.Codec
	}
	return nil
}

func (x *Image) GetRegistry() *RegistryConfig {
	if x != nil {
		return x.Registry
	}
	return nil
}

type ProseFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content []byte `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ProseFile) Reset() {
	*x = ProseFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_image_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProseFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProseFile) ProtoMessage() {}

func (x *ProseFile) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_image_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProseFile.ProtoReflect.Descriptor instead.
func (*ProseFile) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_image_proto_rawDescGZIP(), []int{1}
}

func (x *ProseFile) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ProseFile) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

type PackageConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=label,proto3" json:"label,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Prose string `protobuf:"bytes,3,opt,name=prose,proto3" json:"prose,omitempty"`
}

func (x *PackageConfig) Reset() {
	*x = PackageConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_image_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PackageConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PackageConfig) ProtoMessage() {}

func (x *PackageConfig) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_image_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PackageConfig.ProtoReflect.Descriptor instead.
func (*PackageConfig) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_image_proto_rawDescGZIP(), []int{2}
}

func (x *PackageConfig) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *PackageConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PackageConfig) GetProse() string {
	if x != nil {
		return x.Prose
	}
	return ""
}

type RegistryConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Organization string `protobuf:"bytes,1,opt,name=organization,proto3" json:"organization,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RegistryConfig) Reset() {
	*x = RegistryConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_image_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistryConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistryConfig) ProtoMessage() {}

func (x *RegistryConfig) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_image_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistryConfig.ProtoReflect.Descriptor instead.
func (*RegistryConfig) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_image_proto_rawDescGZIP(), []int{3}
}

func (x *RegistryConfig) GetOrganization() string {
	if x != nil {
		return x.Organization
	}
	return ""
}

func (x *RegistryConfig) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CodecOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TrimSubPackages []string          `protobuf:"bytes,1,rep,name=trim_sub_packages,json=trimSubPackages,proto3" json:"trim_sub_packages,omitempty"`
	WrapOneof       bool              `protobuf:"varint,2,opt,name=wrap_oneof,json=wrapOneof,proto3" json:"wrap_oneof,omitempty"`
	ShortEnums      *ShortEnumOptions `protobuf:"bytes,3,opt,name=short_enums,json=shortEnums,proto3" json:"short_enums,omitempty"`
}

func (x *CodecOptions) Reset() {
	*x = CodecOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_image_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CodecOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CodecOptions) ProtoMessage() {}

func (x *CodecOptions) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_image_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CodecOptions.ProtoReflect.Descriptor instead.
func (*CodecOptions) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_image_proto_rawDescGZIP(), []int{4}
}

func (x *CodecOptions) GetTrimSubPackages() []string {
	if x != nil {
		return x.TrimSubPackages
	}
	return nil
}

func (x *CodecOptions) GetWrapOneof() bool {
	if x != nil {
		return x.WrapOneof
	}
	return false
}

func (x *CodecOptions) GetShortEnums() *ShortEnumOptions {
	if x != nil {
		return x.ShortEnums
	}
	return nil
}

type ShortEnumOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnspecifiedSuffix string `protobuf:"bytes,1,opt,name=unspecified_suffix,json=unspecifiedSuffix,proto3" json:"unspecified_suffix,omitempty"`
	StrictUnmarshal   bool   `protobuf:"varint,2,opt,name=strict_unmarshal,json=strictUnmarshal,proto3" json:"strict_unmarshal,omitempty"`
}

func (x *ShortEnumOptions) Reset() {
	*x = ShortEnumOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jsonapi_v1_image_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ShortEnumOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortEnumOptions) ProtoMessage() {}

func (x *ShortEnumOptions) ProtoReflect() protoreflect.Message {
	mi := &file_jsonapi_v1_image_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortEnumOptions.ProtoReflect.Descriptor instead.
func (*ShortEnumOptions) Descriptor() ([]byte, []int) {
	return file_jsonapi_v1_image_proto_rawDescGZIP(), []int{5}
}

func (x *ShortEnumOptions) GetUnspecifiedSuffix() string {
	if x != nil {
		return x.UnspecifiedSuffix
	}
	return ""
}

func (x *ShortEnumOptions) GetStrictUnmarshal() bool {
	if x != nil {
		return x.StrictUnmarshal
	}
	return false
}

var File_jsonapi_v1_image_proto protoreflect.FileDescriptor

var file_jsonapi_v1_image_proto_rawDesc = []byte{
	0x0a, 0x16, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x8d, 0x02, 0x0a, 0x05, 0x49, 0x6d, 0x61, 0x67, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6a,
	0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x73, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x72, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72,
	0x6f, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x05, 0x70, 0x72, 0x6f, 0x73, 0x65, 0x12, 0x2e,
	0x0a, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x63,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x05, 0x63, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x36,
	0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x72, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x22, 0x39, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x73, 0x65, 0x46,
	0x69, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x4f, 0x0a, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x6f, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x6f,
	0x73, 0x65, 0x22, 0x48, 0x0a, 0x0e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x98, 0x01, 0x0a,
	0x0c, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2a, 0x0a,
	0x11, 0x74, 0x72, 0x69, 0x6d, 0x5f, 0x73, 0x75, 0x62, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x69, 0x6d, 0x53, 0x75,
	0x62, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x72, 0x61,
	0x70, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x77,
	0x72, 0x61, 0x70, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x3d, 0x0a, 0x0b, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x6a, 0x73, 0x6f, 0x6e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0a, 0x73, 0x68, 0x6f,
	0x72, 0x74, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x22, 0x6c, 0x0a, 0x10, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x75,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x75, 0x66, 0x66, 0x69,
	0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x6e, 0x73, 0x70, 0x65, 0x63, 0x69,
	0x66, 0x69, 0x65, 0x64, 0x53, 0x75, 0x66, 0x66, 0x69, 0x78, 0x12, 0x29, 0x0a, 0x10, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x5f, 0x75, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x55, 0x6e, 0x6d, 0x61,
	0x72, 0x73, 0x68, 0x61, 0x6c, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x65, 0x6e, 0x74, 0x6f, 0x70, 0x73, 0x2f, 0x6a, 0x73, 0x6f, 0x6e,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x61,
	0x70, 0x69, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jsonapi_v1_image_proto_rawDescOnce sync.Once
	file_jsonapi_v1_image_proto_rawDescData = file_jsonapi_v1_image_proto_rawDesc
)

func file_jsonapi_v1_image_proto_rawDescGZIP() []byte {
	file_jsonapi_v1_image_proto_rawDescOnce.Do(func() {
		file_jsonapi_v1_image_proto_rawDescData = protoimpl.X.CompressGZIP(file_jsonapi_v1_image_proto_rawDescData)
	})
	return file_jsonapi_v1_image_proto_rawDescData
}

var file_jsonapi_v1_image_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_jsonapi_v1_image_proto_goTypes = []interface{}{
	(*Image)(nil),                            // 0: jsonapi.v1.Image
	(*ProseFile)(nil),                        // 1: jsonapi.v1.ProseFile
	(*PackageConfig)(nil),                    // 2: jsonapi.v1.PackageConfig
	(*RegistryConfig)(nil),                   // 3: jsonapi.v1.RegistryConfig
	(*CodecOptions)(nil),                     // 4: jsonapi.v1.CodecOptions
	(*ShortEnumOptions)(nil),                 // 5: jsonapi.v1.ShortEnumOptions
	(*descriptorpb.FileDescriptorProto)(nil), // 6: google.protobuf.FileDescriptorProto
}
var file_jsonapi_v1_image_proto_depIdxs = []int32{
	6, // 0: jsonapi.v1.Image.file:type_name -> google.protobuf.FileDescriptorProto
	2, // 1: jsonapi.v1.Image.packages:type_name -> jsonapi.v1.PackageConfig
	1, // 2: jsonapi.v1.Image.prose:type_name -> jsonapi.v1.ProseFile
	4, // 3: jsonapi.v1.Image.codec:type_name -> jsonapi.v1.CodecOptions
	3, // 4: jsonapi.v1.Image.registry:type_name -> jsonapi.v1.RegistryConfig
	5, // 5: jsonapi.v1.CodecOptions.short_enums:type_name -> jsonapi.v1.ShortEnumOptions
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_jsonapi_v1_image_proto_init() }
func file_jsonapi_v1_image_proto_init() {
	if File_jsonapi_v1_image_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jsonapi_v1_image_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Image); i {
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
		file_jsonapi_v1_image_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProseFile); i {
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
		file_jsonapi_v1_image_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PackageConfig); i {
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
		file_jsonapi_v1_image_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistryConfig); i {
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
		file_jsonapi_v1_image_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CodecOptions); i {
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
		file_jsonapi_v1_image_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ShortEnumOptions); i {
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
			RawDescriptor: file_jsonapi_v1_image_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_jsonapi_v1_image_proto_goTypes,
		DependencyIndexes: file_jsonapi_v1_image_proto_depIdxs,
		MessageInfos:      file_jsonapi_v1_image_proto_msgTypes,
	}.Build()
	File_jsonapi_v1_image_proto = out.File
	file_jsonapi_v1_image_proto_rawDesc = nil
	file_jsonapi_v1_image_proto_goTypes = nil
	file_jsonapi_v1_image_proto_depIdxs = nil
}