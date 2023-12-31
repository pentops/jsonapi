package structure

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/pentops/jsonapi/gen/v1/jsonapi_pb"
	"github.com/pentops/jsonapi/jsonapi"
	"github.com/pentops/o5-runtime-sidecar/testproto"
	"github.com/stretchr/testify/assert"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	pathMessage = 4
	pathField   = 2
)

func TestBuild(t *testing.T) {

	descriptors := &descriptorpb.FileDescriptorProto{
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String("github.com/pentops/jsonapi/test_pb"),
		},
		Name:    proto.String("test.proto"),
		Package: proto.String("test.v1"),
		Service: []*descriptorpb.ServiceDescriptorProto{{
			Name: proto.String("TestService"),
			Method: []*descriptorpb.MethodDescriptorProto{
				testproto.BuildHTTPMethod("Test", &annotations.HttpRule{
					Pattern: &annotations.HttpRule_Get{
						Get: "/test/{test_field}",
					},
				}),
			},
		}},
		MessageType: []*descriptorpb.DescriptorProto{{
			Name: proto.String("TestRequest"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:   proto.String("test_field"),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Number: proto.Int32(1),
			}},
		}, {
			Name: proto.String("TestResponse"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:   proto.String("test_field"),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Number: proto.Int32(1),
			}, {
				Name:     proto.String("msg"),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
				Number:   proto.Int32(2),
				TypeName: proto.String(".test.v1.Nested"),
			}},
		}, {
			Name: proto.String("Nested"),
			Field: []*descriptorpb.FieldDescriptorProto{{
				Name:   proto.String("nested_field"),
				Type:   descriptorpb.FieldDescriptorProto_TYPE_STRING.Enum(),
				Number: proto.Int32(1),
			}, {
				Name:     proto.String("enum"),
				Type:     descriptorpb.FieldDescriptorProto_TYPE_ENUM.Enum(),
				Number:   proto.Int32(3),
				TypeName: proto.String(".test.v1.TestEnum"),
			}},
		}},
		EnumType: []*descriptorpb.EnumDescriptorProto{{
			Name: proto.String("TestEnum"),
			Value: []*descriptorpb.EnumValueDescriptorProto{{
				Name:   proto.String("TEST_ENUM_UNSPECIFIED"),
				Number: proto.Int32(0),
			}, {
				Name:   proto.String("TEST_ENUM_FOO"),
				Number: proto.Int32(1),
			}},
		}},

		SourceCodeInfo: &descriptorpb.SourceCodeInfo{
			Location: []*descriptorpb.SourceCodeInfo_Location{{
				LeadingComments: proto.String("Message Comment"),
				Path:            []int32{pathMessage, 2}, // 4 = Message, 2 = Nested
				Span:            []int32{1, 1, 1},        // Single line comment
			}, {
				LeadingComments: proto.String("Field Comment"),
				Path:            []int32{pathMessage, 2, pathField, 0}, // 4 = Message, 2 = Nested, 1 = Field
				Span:            []int32{2, 1, 2},                      // Single line comment
			}},
		},
	}

	want := &Built{
		Packages: []*Package{{
			Label: "Test",
			Name:  "test.v1",
			Methods: []*Method{{
				GrpcServiceName: "TestService",
				FullGrpcName:    "/test.v1.TestService/Test",
				GrpcMethodName:  "Test",
				HTTPMethod:      "get",
				HTTPPath:        "/test/:testField",
				ResponseBody: &jsonapi.SchemaItem{
					ItemType: &jsonapi.ObjectItem{
						FullProtoName: "test.v1.TestResponse",
						ProtoName:     "TestResponse",
						ObjectRules:   jsonapi.ObjectRules{},
						IsOneof:       false,
						GoPackageName: "github.com/pentops/jsonapi/test_pb",
						GoTypeName:    "TestResponse",
						GRPCPackage:   "test.v1",
						Properties: []*jsonapi.ObjectProperty{{
							Name:             "testField",
							Description:      "",
							ProtoFieldName:   "test_field",
							ProtoFieldNumber: 1,
							Skip:             false,
							Optional:         true,
							Required:         false,
							WriteOnly:        false,
							SchemaItem: jsonapi.SchemaItem{
								Description: "",
								ItemType:    jsonapi.StringItem{},
							},
						}, {
							Name:             "msg",
							Description:      "",
							ProtoFieldName:   "msg",
							ProtoFieldNumber: 2,
							Required:         false,
							Skip:             false,
							Optional:         true,
							SchemaItem: jsonapi.SchemaItem{
								Ref: "test.v1.Nested",
							},
						}},
					},
				},
				PathParameters: []*Parameter{{
					Name:        "testField",
					Description: "",
					Required:    true,
					Schema: jsonapi.SchemaItem{
						ItemType: jsonapi.StringItem{},
					},
				}},
			}},
		}},
	}

	built, err := BuildFromDescriptors(&jsonapi_pb.Config{
		Packages: []*jsonapi_pb.PackageConfig{{
			Label: "Test",
			Name:  "test.v1",
		}},
		Options: &jsonapi_pb.CodecOptions{
			ShortEnums: &jsonapi_pb.ShortEnumOptions{},
		},
	}, &descriptorpb.FileDescriptorSet{
		File: []*descriptorpb.FileDescriptorProto{descriptors},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	bb, err := json.MarshalIndent(built, "", "  ")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(bb))

	// Packages are controlled by this package, should equal in full. Schema
	// tests are in the jsonapi package.
	assert.Equal(t, want.Packages, built.Packages)

	if _, ok := built.Schemas["test.v1.TestRequest"]; ok {
		t.Fatal("TestRequest should not be registered as a schema, but was")
	}

	refSchema, ok := built.Schemas["test.v1.Nested"]
	if !ok {
		t.Fatal("schema not found")
	}

	if tn := refSchema.ItemType.TypeName(); tn != "object" {
		t.Fatalf("unexpected type: %s", tn)
	}

	if refSchema.Description != "Message Comment" {
		t.Errorf("unexpected description: '%s'", refSchema.Description)
	}

	asObject := refSchema.ItemType.(*jsonapi.ObjectItem)
	if len(asObject.Properties) != 2 {
		t.Fatalf("unexpected properties: %d", len(asObject.Properties))
	}

	f1 := asObject.Properties[0]
	if f1.Name != "nestedField" {
		t.Errorf("unexpected field name: '%s'", f1.Name)
	}

	if f1.Description != "Field Comment" {
		t.Errorf("unexpected description: '%s'", f1.Description)
	}

	fEnum := asObject.Properties[1]
	if fEnum.Name != "enum" {
		t.Errorf("unexpected field name: '%s'", fEnum.Name)
	}

	ref := fEnum.SchemaItem.Ref
	if ref == "" {
		t.Fatal("ref is nil")
	}

	schemaEnum, ok := built.Schemas[strings.TrimPrefix(ref, "#/components/schemas/")]
	if !ok {
		t.Fatalf("schema not found: '%s'", ref)
	}

	enumType, ok := schemaEnum.ItemType.(jsonapi.EnumItem)
	if !ok {
		t.Fatalf("unexpected type: %T", fEnum.SchemaItem.ItemType)
	}

	if enumType.Enum[0] != "UNSPECIFIED" {
		t.Errorf("unexpected enum value: '%s'", enumType.Enum[0])
	}
	if enumType.Enum[1] != "FOO" {
		t.Errorf("unexpected enum value: '%s'", enumType.Enum[1])
	}

}
