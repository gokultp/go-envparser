package generator

import (
	"fmt"
	"strings"
	"text/template"
)

const (
	typeBool       = "bool"
	typeByte       = "byte"
	typeComplex128 = "complex128"
	typeComplex64  = "complex64"
	typeFloat32    = "float32"
	typeFloat64    = "float64"
	typeInt        = "int"
	typeInt8       = "int8"
	typeInt16      = "int16"
	typeInt32      = "int32"
	typeInt64      = "int64"
	typeRune       = "rune"
	typeString     = "string"
	typeUInt       = "uint"
	typeUInt8      = "uint8"
	typeUInt16     = "uint16"
	typeUInt32     = "uint32"
	typeUint64     = "uint64"
	typeUIntPtr    = "uintptr"
)

func isBuiltInType(typ string) bool {
	fmt.Println("hhhhh", typ)
	switch typ {
	case typeBool, typeByte, typeComplex128, typeComplex64:
	case typeFloat32, typeFloat64:
	case typeInt, typeInt8, typeInt16, typeInt32, typeInt64:
	case typeRune, typeString:
	case typeUInt, typeUInt8, typeUInt16, typeUInt32, typeUint64, typeUIntPtr:
	default:
		return false
	}
	return true
}

func parseFunc(typ string, arg string) string {
	base := baseType(typ)
	switch base {
	case typeInt:
		bits := strings.Replace(typ, typeInt, "", 1)
		if bits == "" {
			bits = "32"
		}
		return fmt.Sprintf("strconv.ParseInt(_rec%sStr, %d, %s)", arg, 10, bits)
	case typeUInt:
		bits := strings.Replace(typ, typeUInt, "", 1)
		if bits == "" {
			bits = "32"
		}
		return fmt.Sprintf("strconv.ParseInt(%sStr, %d, %s)", varName(arg), 10, bits)
	}
	return ""
}

func getFuncMap() template.FuncMap {
	return template.FuncMap{
		"is_builtin":           isBuiltInType,
		"basetype":             baseType,
		"parsefunc":            parseFunc,
		"is_notsupported_type": isNotSupportedType,
		"varname":              varName,
	}
}

func isNotSupportedType(t string) bool {
	// complex is not supported
	return isBuiltInType(t) && (t == typeComplex128 || t == typeComplex64)
}

func baseType(typ string) string {
	switch typ {
	case typeInt, typeInt8, typeInt16, typeInt32, typeInt64:
		return typeInt
	case typeUInt, typeUInt8, typeUInt16, typeUInt32, typeUint64, typeUIntPtr:
		return typeUInt
	default:
		return typ
	}
}

func varName(tag string) string {
	return fmt.Sprintf("_rec%s", strings.ToTitle(tag))
}
