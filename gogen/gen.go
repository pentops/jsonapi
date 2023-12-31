package gogen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/pentops/jsonapi/jsonapi"
	"github.com/pentops/jsonapi/structure"
)

type Options struct {
	TrimPackagePrefix string
	AddGoPrefix       string
}

func (o Options) ToGoPackage(pkg string) (string, error) {
	if pkg == "" {
		return "", fmt.Errorf("empty package")
	}

	if o.TrimPackagePrefix != "" {
		pkg = strings.TrimPrefix(pkg, o.TrimPackagePrefix)
	}

	pkg = strings.TrimSuffix(pkg, ".service")
	pkg = strings.TrimSuffix(pkg, ".topic")
	pkg = strings.TrimSuffix(pkg, ".sandbox")

	parts := strings.Split(pkg, ".")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid package: %s", pkg)
	}
	nextName := parts[len(parts)-2]
	parts = append(parts, nextName)

	pkg = strings.Join(parts, "/")

	if o.AddGoPrefix != "" {
		pkg = path.Join(o.AddGoPrefix, pkg)
	}
	return pkg, nil
}

func scalarTypeName(item jsonapi.SchemaItem) (*DataType, error) {
	switch item := item.ItemType.(type) {
	case jsonapi.StringItem:
		switch item.Format {
		case "", "uuid", "date", "email":
			return &DataType{
				Name:    "string",
				Pointer: false,
			}, nil
		case "date-time":
			return &DataType{
				Name:    "Time",
				Pointer: true,
				Package: "time",
			}, nil
		case "byte":
			return &DataType{
				Name:    "[]byte",
				Pointer: false,
			}, nil
		default:
			return nil, fmt.Errorf("Unknown string format: %s", item.Format)
		}

	case jsonapi.NumberItem:
		return &DataType{
			Name:    "float64",
			Pointer: false,
		}, nil

	case jsonapi.IntegerItem:
		return &DataType{
			Name:    "int64",
			Pointer: false,
		}, nil

	case jsonapi.BooleanItem:
		return &DataType{
			Name:    "bool",
			Pointer: false,
		}, nil

	case jsonapi.EnumItem:
		return &DataType{
			Name:    "string",
			Pointer: false,
		}, nil

	default:
		return nil, fmt.Errorf("Unknown type: %T\n", item)
	}
}

type builder struct {
	document *structure.Built
	fileSet  *FileSet
	options  Options
}

type FileWriter interface {
	WriteFile(name string, data []byte) error
}

type DirFileWriter string

func (fw DirFileWriter) WriteFile(relPath string, data []byte) error {
	fullPath := filepath.Join(string(fw), relPath)
	dirName := filepath.Dir(fullPath)
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return fmt.Errorf("mkdirall for %s: %w", fullPath, err)
	}
	if err := os.WriteFile(fullPath, data, 0644); err != nil {
		return fmt.Errorf("writefile for %s: %w", fullPath, err)
	}
	return nil
}

func WriteGoCode(document *structure.Built, output FileWriter, options Options) error {

	fileSet := NewFileSet(options.AddGoPrefix)

	bb := &builder{
		document: document,
		fileSet:  fileSet,
		options:  options,
	}

	err := bb.root()
	if err != nil {
		return err
	}

	return fileSet.WriteAll(output)
}

func (bb *builder) buildTypeName(schema jsonapi.SchemaItem) (*DataType, error) {

	if schema.Ref != "" {
		refVal, ok := bb.document.Schemas[schema.Ref]
		if !ok {
			return nil, fmt.Errorf("Unknown ref: %s", schema.Ref)
		}

		switch referredType := refVal.ItemType.(type) {
		case jsonapi.EnumItem:
			return &DataType{
				Name:    "string",
				Pointer: false,
			}, nil

		case *jsonapi.ObjectItem:

			if err := bb.addObject(referredType); err != nil {
				return nil, fmt.Errorf("referencedType in %s: %w", schema.Ref, err)
			}

			objectPackage, err := bb.options.ToGoPackage(referredType.GRPCPackage)
			if err != nil {
				return nil, fmt.Errorf("referredType in %s: %w", schema.Ref, err)
			}

			return &DataType{
				Name:    referredType.GoTypeName,
				Package: objectPackage,
				Pointer: true,
			}, nil

		default:
			return nil, fmt.Errorf("Unknown ref type: %T", referredType)
		}
	}

	if arrayType, ok := schema.ItemType.(jsonapi.ArrayItem); ok {
		itemType, err := bb.buildTypeName(arrayType.Items)
		if err != nil {
			return nil, err
		}

		return &DataType{
			Name:    itemType.Name,
			Pointer: itemType.Pointer,
			Slice:   true,
		}, nil
	}

	if mapType, ok := schema.ItemType.(jsonapi.MapItem); ok {
		keyType, err := bb.buildTypeName(mapType.KeyProperty)
		if err != nil {
			return nil, fmt.Errorf("map key: %w", err)
		}
		valueType, err := bb.buildTypeName(mapType.ValueProperty)
		if err != nil {
			return nil, fmt.Errorf("map value: %w", err)
		}

		return &DataType{
			Name:    fmt.Sprintf("map[%s]%s", keyType.Name, valueType.Name),
			Pointer: false,
		}, nil

	}

	if objectType, ok := schema.ItemType.(*jsonapi.ObjectItem); ok {

		if err := bb.addObject(objectType); err != nil {
			return nil, fmt.Errorf("directObject: %w", err)
		}

		return &DataType{
			Name:    objectType.GoTypeName,
			Pointer: true,
		}, nil

	}

	return scalarTypeName(schema)

}

func (bb *builder) jsonField(property *jsonapi.ObjectProperty) (*Field, error) {

	tags := map[string]string{}

	tags["json"] = property.Name
	if !property.Required {
		tags["json"] += ",omitempty"
	}

	dataType, err := bb.buildTypeName(property.SchemaItem)
	if err != nil {
		return nil, fmt.Errorf("building field %s: %w", property.Name, err)
	}

	if !dataType.Pointer && !dataType.Slice && property.Optional {
		dataType.Pointer = true
	}

	return &Field{
		Name:     GoName(property.Name),
		DataType: *dataType,
		Tags:     tags,
	}, nil

}

func (bb *builder) addObject(object *jsonapi.ObjectItem) error {
	objectPackage, err := bb.options.ToGoPackage(object.GRPCPackage)
	if err != nil {
		return fmt.Errorf("object package name '%s': %w", object.GoTypeName, err)
	}
	gen, err := bb.fileSet.File(objectPackage, filepath.Base(objectPackage))
	if err != nil {
		return err
	}

	_, ok := gen.types[object.GoTypeName]
	if ok {
		return nil
	}

	structType := &Struct{
		Name: object.GoTypeName,
		Comment: fmt.Sprintf(
			"Proto: %s",
			object.FullProtoName,
		),
	}
	gen.types[object.GoTypeName] = structType

	for _, property := range object.Properties {
		field, err := bb.jsonField(property)
		if err != nil {
			return fmt.Errorf("object %s: %w", object.GoTypeName, err)
		}
		structType.Fields = append(structType.Fields, field)
	}

	return nil
}

func (bb *builder) root() error {

	for _, pkgsss := range bb.document.Packages {
		fullGoPackage, err := bb.options.ToGoPackage(pkgsss.Name)
		if err != nil {
			return err
		}
		for _, operation := range pkgsss.Methods {
			if err := bb.addOperation(fullGoPackage, operation); err != nil {
				return err
			}
		}
	}
	return nil
}

func (bb *builder) addOperation(fullGoPackage string, operation *structure.Method) error {

	goPackageName := path.Base(fullGoPackage)
	gen, err := bb.fileSet.File(fullGoPackage, goPackageName)
	if err != nil {
		return err
	}

	gen.EnsureInterface(&Interface{
		Name: "Requester",
		Methods: []*Function{{
			Name: "Request",
			Parameters: []*Parameter{{
				Name: "ctx",
				DataType: DataType{
					Package: "context",
					Name:    "Context",
				},
			}, {
				Name:     "method",
				DataType: DataType{Name: "string"},
			}, {
				Name:     "path",
				DataType: DataType{Name: "string"},
			}, {
				Name:     "body",
				DataType: DataType{Name: "interface{}"},
			}, {
				Name:     "response",
				DataType: DataType{Name: "interface{}"},
			}},
			Returns: []*Parameter{{
				DataType: DataType{
					Name: "error",
				}},
			},
		}},
	})

	service := gen.Service(operation.GrpcServiceName)

	requestType := fmt.Sprintf("%sRequest", operation.GrpcMethodName)
	responseType := fmt.Sprintf("%sResponse", operation.GrpcMethodName)

	requestStruct := &Struct{
		Name: requestType,
	}

	if _, ok := gen.types[requestType]; ok {
		return fmt.Errorf("response type %q already exists", responseType)
	}
	gen.types[requestType] = requestStruct

	pathParameters := map[string]*Field{}

	for _, parameter := range operation.PathParameters {
		typeName, err := bb.buildTypeName(parameter.Schema)
		if err != nil {
			return err
		}

		field := &Field{
			Name:     GoName(parameter.Name),
			DataType: *typeName,
			Tags: map[string]string{
				"path": parameter.Name,
				"json": "-",
			},
		}
		requestStruct.Fields = append(requestStruct.Fields, field)
		pathParameters[parameter.Name] = field
	}

	queryMethod := &Function{
		Name:       "QueryParameters",
		Parameters: []*Parameter{},
		Returns: []*Parameter{{
			DataType: DataType{
				Name:    "Values",
				Package: "net/url",
			},
		}, {
			DataType: DataType{
				Name: "error",
			},
		}},
		StringGen: gen.ChildGen(),
	}

	queryMethod.P("  values := ", DataType{Package: "net/url", Name: "Values"}, "{}")

	usedQuery := false

	for _, parameter := range operation.QueryParameters {
		usedQuery = true
		dataType, err := bb.buildTypeName(parameter.Schema)
		if err != nil {
			return err
		}

		if !dataType.Pointer && !dataType.Slice && !parameter.Required {
			dataType.Pointer = true
		}

		requestStruct.Fields = append(requestStruct.Fields, &Field{
			Name:     GoName(parameter.Name),
			DataType: *dataType,
			Tags: map[string]string{
				"query": parameter.Name,
				"json":  "-",
			},
		})

		schema := parameter.Schema
		if parameter.Schema.Ref != "" {
			si, ok := bb.document.Schemas[parameter.Schema.Ref]
			if !ok {
				return fmt.Errorf("Unknown ref: %s", parameter.Schema.Ref)
			}
			schema = *si
		}

		switch schema.ItemType.(type) {
		case jsonapi.StringItem:
			if parameter.Required {
				queryMethod.P("  values.Set(\"", parameter.Name, "\", s.", GoName(parameter.Name), ")")
			} else {
				queryMethod.P("  if s.", GoName(parameter.Name), " != nil {")
				queryMethod.P("    values.Set(\"", parameter.Name, "\", *s.", GoName(parameter.Name), ")")
				queryMethod.P("  }")
			}

		case *jsonapi.ObjectItem:
			// include as JSON
			queryMethod.P("  if s.", GoName(parameter.Name), " != nil {")
			queryMethod.P("    bb, err := ", DataType{Package: "encoding/json", Name: "Marshal"}, "(s.", GoName(parameter.Name), ")")
			queryMethod.P("    if err != nil {")
			queryMethod.P("      return nil, err")
			queryMethod.P("    }")
			queryMethod.P("    values.Set(\"", parameter.Name, "\", string(bb))")
			queryMethod.P("  }")

		default:
			queryMethod.P(" // Skipping query parameter ", parameter.Name, " of type ", dataType.Name)
			//queryMethod.P("    values.Set(\"", parameter.Name, "\", fmt.Sprintf(\"%v\", *s.", GoName(parameter.Name), "))")
		}
	}

	queryMethod.P("  return values, nil")

	if usedQuery {
		requestStruct.Methods = append(requestStruct.Methods, queryMethod)
	}

	if operation.RequestBody != nil {
		requestSchema := operation.RequestBody.ItemType.(*jsonapi.ObjectItem)
		for _, property := range requestSchema.Properties {
			field, err := bb.jsonField(property)
			if err != nil {
				return err
			}
			requestStruct.Fields = append(requestStruct.Fields, field)
		}
	}

	responseStruct := &Struct{
		Name: responseType,
	}

	if _, ok := gen.types[responseType]; ok {
		return fmt.Errorf("response type %q already exists", responseType)
	}

	gen.types[responseType] = responseStruct

	responseSchema := operation.ResponseBody.ItemType.(*jsonapi.ObjectItem)
	for _, property := range responseSchema.Properties {
		field, err := bb.jsonField(property)
		if err != nil {
			return fmt.Errorf("%s.ResponseBody: %w", operation.FullGrpcName, err)
		}
		responseStruct.Fields = append(responseStruct.Fields, field)
	}

	requestMethod := &Function{
		Name: operation.GrpcMethodName,
		Parameters: []*Parameter{{
			Name: "ctx",
			DataType: DataType{
				Package: "context",
				Name:    "Context",
			},
		}, {
			Name: "req",
			DataType: DataType{
				Name:    requestType,
				Pointer: true,
			},
		}},
		Returns: []*Parameter{{
			DataType: DataType{
				Name:    responseType,
				Pointer: true,
			},
		}, {
			DataType: DataType{
				Name: "error",
			},
		}},
		StringGen: gen.ChildGen(),
	}

	pathParts := strings.Split(operation.HTTPPath, "/")
	pathParams := make([]string, 0)
	for idx, part := range pathParts {
		if len(part) == 0 {
			continue
		}
		if part[0] == ':' {
			name := part[1:]
			field, ok := pathParameters[name]
			if !ok {
				return fmt.Errorf("path parameter %q not found in request object %s", name, requestType)
			}

			pathParts[idx] = "%s"

			pathParams = append(pathParams, fmt.Sprintf("req.%s", field.Name))
		}
	}

	requestMethod.P("  path := ", DataType{Package: "fmt", Name: "Sprintf"}, "(\"", strings.Join(pathParts, "/"), "\", ")
	for _, param := range pathParams {
		requestMethod.P("   ", param, ", ")
	}
	requestMethod.P("  )")

	requestMethod.P("  resp := &", responseType, "{}")
	requestMethod.P("  err := s.Request(ctx, \"", strings.ToUpper(operation.HTTPMethod), "\", path, req, resp)")
	requestMethod.P("  if err != nil {")
	requestMethod.P("    return nil, err")
	requestMethod.P("  }")

	requestMethod.P("  return resp, nil")

	service.Methods = append(service.Methods, requestMethod)

	return nil

}

var reUnsafe = regexp.MustCompile(`[^a-zA-Z0-9_]`)

// GoName exports the field name
func GoName(name string) string {
	name = reUnsafe.ReplaceAllString(name, "_")
	return strings.ToUpper(name[0:1]) + name[1:]
}
