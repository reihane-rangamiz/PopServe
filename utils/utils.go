package utils

import (
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

func DetectType(field string, value interface{}) (goType string, skip bool, tag string) {
	switch val := value.(type) {
	case bool:
		if !val {
			return "", true, ""
		}
		// true: infer based on name
		if field == "created_at" || field == "updated_at" || field == "deleted_at" || strings.HasSuffix(field, "_at") {
			if field == "created_at" {
				return "time.Time", false, "autoCreateTime"
			}
			if field == "updated_at" {
				return "time.Time", false, "autoUpdateTime"
			}
			if field == "deleted_at" {
				return "time.Time", false, "autoCreateTime"
			}
		}
		if val {
			return "string", false, ""
		}

		return "bool", false, ""

	case string:
		switch val {
		case "pk":
			return "uint", false, "primaryKey"
		case "fk":
			return "uint", false, "foreignKey"
		case "uuid":
			return "string", false, "type:uuid"
		case "str", "string":
			return "string", false, ""
		case "int":
			return "int", false, ""
		case "float":
			return "float64", false, ""
		case "jwt":
			return "string", false, ""
		default:
			return "string", false, ""
		}
	case float32, float64:
		return "float64", false, ""
	case int, int32, int64:
		return "int", false, ""
	case uint:
		return "uint", false, ""
	default:
		return "interface{}", false, ""
	}
}
