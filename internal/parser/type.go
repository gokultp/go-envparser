package parser

import (
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"io/ioutil"
	"reflect"
	"strings"
)

// Type is the type defenition struct
type Type struct {
	FileName string
	Name     string
	Fields   []Field
	Package  string
}

// NewType returns a new instance of Type with given name
func NewType(name string) *Type {
	return &Type{
		Name: name,
	}
}

// getStruct returns the struct metadata if the given node is a struct
func getStruct(nodeType ast.Node) *ast.StructType {
	switch node := nodeType.(type) {
	case *ast.StructType:
		return node
	default:
		return nil
	}
}

// Parse parses Type metadata from given file using go parser & ast
func (t *Type) Parse(fileName string) error {
	t.FileName = fileName
	fset := token.NewFileSet()
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}
	f, err := parser.ParseFile(fset, fileName, content, parser.ParseComments)
	if err != nil {
		return err
	}
	t.astInspect(f)
	return nil
}

// astInspect will inspect the go-ast of the struct and get the meta data
func (t *Type) astInspect(f *ast.File) {
	ast.Inspect(f, func(node ast.Node) bool {
		switch nodeType := node.(type) {
		case *ast.TypeSpec:
			if node := getStruct(nodeType.Type); node != nil {
				// Fetch info for only the given struct
				if t.Name == nodeType.Name.String() {
					// Helper to populate struct's field and tags info
					t.Package = f.Name.String()
					t.Fields = getFields(node)
					return false
				}
			}
		}
		return true
	})
}

// getFields will transforms the field metadata returned by go ast to the template's format
func getFields(node *ast.StructType) []Field {
	var fields []Field
	for _, field := range node.Fields.List {
		var tags reflect.StructTag
		if field.Tag != nil {
			tags = reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
		}
		if len(field.Names) == 0 {
			fieldType := types.ExprString(field.Type)
			fields = append(fields, NewField(nil, fieldType, tags))
			continue
		}
		for _, fieldName := range field.Names {
			fieldType := types.ExprString(field.Type)
			fields = append(fields, NewField(fieldName, fieldType, tags))
		}
	}
	return fields
}
