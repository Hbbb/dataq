package main

import (
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + "_index.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)

	g.P("// Code generated by protoc-gen-dataq-index. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()

	for _, message := range file.Messages {
		generateIndexMethods(g, message)
	}
}

func generateIndexMethods(g *protogen.GeneratedFile, message *protogen.Message) {
	// Generate SchemaKind method
	g.P("func (m *", message.GoIdent.GoName, ") SchemaKind() string {")
	g.P(`    return "`, message.GoIdent.GoName, `"`)
	g.P("}")
	g.P()

	// Generate Metadata method
	g.P("func (m *", message.GoIdent.GoName, ") SchemaMetadata() map[string]interface{} {")
	g.P("    metadata := make(map[string]interface{})")
	g.P()

	for _, field := range message.Fields {
		g.P("    if m.", field.GoName, " != ", zeroValue(field), " {")
		g.P(`        metadata["`, field.Desc.TextName(), `"] = m.`, field.GoName)
		g.P("    }")
	}

	g.P("    return metadata")
	g.P("}")
	g.P()
}

func zeroValue(field *protogen.Field) string {
	switch field.Desc.Kind() {
	case protoreflect.StringKind:
		return `""`
	case protoreflect.BoolKind:
		return "false"
	case protoreflect.Int32Kind, protoreflect.Int64Kind, protoreflect.Sint32Kind,
		protoreflect.Sint64Kind, protoreflect.Uint32Kind, protoreflect.Uint64Kind,
		protoreflect.Fixed32Kind, protoreflect.Fixed64Kind, protoreflect.Sfixed32Kind,
		protoreflect.Sfixed64Kind:
		return "0"
	case protoreflect.FloatKind, protoreflect.DoubleKind:
		return "0"
	case protoreflect.BytesKind:
		return "nil"
	case protoreflect.MessageKind, protoreflect.GroupKind:
		return "nil"
	case protoreflect.EnumKind:
		return "0"
	default:
		return "nil"
	}
}
