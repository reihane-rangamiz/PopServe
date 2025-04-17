package utils

import (
	"reflect"
	"strings"
)

func Capitalize(s string) string {
	if s == "" {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func InferGoType(value interface{}) string {
	switch value.(type) {
	case string:
		return "string"
	case int, int32, int64:
		return "int"
	case float32, float64:
		return "float64"
	case bool:
		return "bool"
	case []interface{}:
		return "[]interface{}"
	case map[string]interface{}:
		return "map[string]interface{}"
	default:
		return "interface{}"
	}
}

func DetectType(field string, value interface{}) string {
	switch {
	case field == "id" || strings.HasSuffix(field, "_id"):
		return "uint"
	case field == "created_at" || field == "updated_at" || field == "deleted_at":
		return "time.Time"
	case reflect.TypeOf(value).Kind() == reflect.String:
		return "string"
	case reflect.TypeOf(value).Kind() == reflect.Bool:
		return "bool"
	case reflect.TypeOf(value).Kind() == reflect.Float64:
		// You can decide whether to round it to int or keep float64
		return "float64"
	default:
		return "interface{}"
	}
}
