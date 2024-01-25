package util

import (
	"fmt"
	"strings"
)

func MySQLToGo(sqlType string, unsigned bool) string {
	switch strings.ToUpper(sqlType) {
	// case "BIT":
	// 	return "bool"
	case "BOOL", "BOOLEAN", "TINYINT":
		if unsigned {
			return "uint8"
		} else {
			return "int8"
		}
	case "SMALLINT":
		if unsigned {
			return "uint16"
		} else {
			return "int16"
		}
	case "MEDIUMINT", "INT":
		if unsigned {
			return "uint"
		} else {
			return "int"
		}
	case "BIGINT":
		if unsigned {
			return "uint64"
		} else {
			return "int64"
		}
	case "FLOAT":
		return "float32"
	case "DOUBLE", "DECIMAL", "NUMERIC":
		return "float64"
	case "CHAR", "VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT":
		return "string"
	case "VARBINARY":
		return "[]byte"
	case "DATE":
		return "time.Time"
	case "TIME":
		return "time.Time"
	case "DATETIME":
		return "time.Time"
	case "TIMESTAMP":
		return "time.Time"
	default:
		panic(fmt.Errorf("unknown type: %v", sqlType))
	}
}

func MySQLToKSQL(sqlType string) string {
	switch strings.ToUpper(sqlType) {
	// case "BIT":
	// 	return "BOOLEAN"
	case "BOOL", "BOOLEAN", "TINYINT", "SMALLINT", "MEDIUMINT", "INT":
		return "INT"
	case "BIGINT":
		return "BIGINT"
	case "FLOAT", "DOUBLE", "DECIMAL", "NUMERIC":
		return "DOUBLE"
	case "CHAR", "VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT":
		return "VARCHAR"
	case "VARBINARY":
		return "BYTES"
	case "DATE":
		return "DATE"
	case "TIME":
		return "TIME"
	case "DATETIME", "TIMESTAMP":
		return "TIMESTAMP"
	default:
		panic(fmt.Errorf("unknown type: %v", sqlType))
	}
}

func MySQLToProto(sqlType string, unsigned bool) string {
	switch strings.ToUpper(sqlType) {
	// case "BIT":
	// return "bool"
	case "BOOL", "BOOLEAN", "TINYINT", "SMALLINT", "MEDIUMINT", "INT":
		if unsigned {
			return "uint32"
		} else {
			return "int32"
		}
	case "BIGINT":
		if unsigned {
			return "uint64"
		} else {
			return "int64"
		}
	case "FLOAT":
		return "float"
	case "DOUBLE", "DECIMAL", "NUMERIC":
		return "double"
	case "CHAR", "VARCHAR", "TINYTEXT", "TEXT", "MEDIUMTEXT", "LONGTEXT":
		return "string"
	case "VARBINARY":
		return "bytes"
	case "DATE", "TIME", "DATETIME":
		return "string"
	case "TIMESTAMP":
		return "google.protobuf.Timestamp"
	default:
		panic(fmt.Errorf("unknown type: %v", sqlType))
	}
}

func ProtoToGo(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "bool":
		return "bool"
	case "int32", "sint32", "sfixed32":
		return "int32"
	case "int64", "sint64", "sfixed64":
		return "int64"
	case "uint32", "fixed32":
		return "uint32"
	case "uint64", "fixed64":
		return "uint64"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "string":
		return "string"
	case "bytes":
		return "[]byte"
	default:
		panic(fmt.Errorf("unknown type: %v", sqlType))
	}
}

func ToGraphql(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "bool":
		return "Boolean"
	case "int32", "sint32", "sfixed32":
		return "Int"
	case "int64", "sint64", "sfixed64":
		return "Int"
	case "uint32", "fixed32":
		return "Int"
	case "uint64", "fixed64":
		return "Int"
	case "float", "double":
		return "Float"
	case "string":
		return "String"
	case "bytes":
		return ""
	default:
		return sqlType
	}
}

func IsProtoType(protoType string) bool {
	protoTypeMap := map[string]bool{
		"double":   true,
		"float":    true,
		"int32":    true,
		"int64":    true,
		"uint32":   true,
		"uint64":   true,
		"sint32":   true,
		"sint64":   true,
		"fixed32":  true,
		"fixed64":  true,
		"sfixed32": true,
		"sfixed64": true,
		"bool":     true,
		"string":   true,
		"bytes":    true,
		"[]bytes":  true,
	}
	if _, ok := protoTypeMap[protoType]; ok {
		return true
	}
	return false
}
