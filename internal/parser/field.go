package parser

import (
	"go/ast"
	"reflect"
	"strings"
)

// Field encapsulates the struct field metadata needed for the template execution
type Field struct {
	Name      string
	Type      string
	EnvTag    string
	IsPointer bool
	IsArray   bool
}

// NewField creates a new instance of field
func NewField(field *ast.Ident, fieldType string, tags reflect.StructTag) Field {
	if field == nil {
		return Field{
			Name:      "",
			Type:      cleanTypeStr(fieldType),
			IsPointer: isPointer(fieldType),
			IsArray:   isArray(fieldType),
			EnvTag:    getEnvSourceTag(tags, fieldType),
		}
	}
	return Field{
		Name:      field.Name,
		Type:      cleanTypeStr(fieldType),
		IsPointer: isPointer(fieldType),
		IsArray:   isArray(fieldType),
		EnvTag:    getEnvSourceTag(tags, field.Name),
	}
}

// isPointer checks if a given type is a pointer or not
func isPointer(typeName string) bool {
	return len(typeName) > 0 && typeName[0] == '*'
}

// isArray checks is a given type is an array or not
func isArray(typeName string) bool {
	return len(typeName) > 2 && typeName[:2] == "[]"
}

// getEnvSourceTag will says what is ths env variable to be used to fetch the data
func getEnvSourceTag(tags reflect.StructTag, fieldName string) string {
	tag, ok := tags.Lookup("env")
	if !ok {
		return strings.ToUpper(fieldName)
	}
	return tag
}

// cleanTypeStr will strip all unwanted space and other characters to return the type name
func cleanTypeStr(typ string) string {
	typ = strings.TrimSpace(typ)
	typ = strings.TrimLeft(typ, "*")
	typ = strings.TrimLeft(typ, "[]")
	return typ
}
