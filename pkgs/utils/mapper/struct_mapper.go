package mapper

import (
	"github.com/bytedance/sonic"
	"github.com/jinzhu/copier"
	"reflect"
	"regexp"
	"strings"
)

// Copy - copy struct to struct
func Copy(dest, src interface{}) error {
	return copier.Copy(dest, src)
}

// CopyIgnoreEmpty - copy struct to struct ignore zero value
func CopyIgnoreEmpty(dest, src interface{}) error {
	return copier.CopyWithOption(dest, src, copier.Option{IgnoreEmpty: true})
}

// BindingStruct - biding struct to struct
func BindingStruct(src interface{}, desc interface{}) error {
	// convert to byte
	byteSrc, err := sonic.Marshal(src)
	if err != nil {
		return err
	}
	// binding to desc
	err = sonic.Unmarshal(byteSrc, &desc)
	if err != nil {
		return err
	}
	return nil
}

func BindingAndValidate[T any](detail interface{}, validator func(interface{}) error) (T, error) {
	var model T
	if err := BindingStruct(detail, &model); err != nil {
		return model, err
	}

	if err := validator(model); err != nil {
		return model, err
	}
	return model, nil
}

// Generalized function for converting a struct or a pointer to a struct into a map
func structToMap(input interface{}, ignoreNilField bool, snakeCase bool) map[string]interface{} {
	result := make(map[string]interface{})

	// Ensure input is a struct or a pointer to a struct
	v := reflect.ValueOf(input)
	if v.Kind() == reflect.Pointer {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		panic("structToMap: input must be a struct or a pointer to a struct")
	}

	// Loop through struct fields
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Skip fields with `json:"-"`
		if field.Tag.Get("json") == "-" {
			continue
		}

		fv := v.Field(i)

		// Skip nil pointers if ignoreNilField is true
		if ignoreNilField && fv.Kind() == reflect.Pointer && fv.IsNil() {
			continue
		}

		// Dereference pointer if applicable
		if fv.Kind() == reflect.Pointer && !fv.IsNil() {
			fv = fv.Elem()
		}

		fieldName := field.Tag.Get("json")
		if fieldName == "" {
			fieldName = field.Name

		}

		if snakeCase {
			fieldName = CamelToSnake(fieldName)
		}
		result[fieldName] = fv.Interface()
	}

	return result
}

// CamelToSnake converts CamelCase to snake_case
func CamelToSnake(s string) string {
	re := regexp.MustCompile("([a-z0-9])([A-Z])")
	snake := re.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}

func StructToMap(input interface{}, ignoreNilField bool) map[string]interface{} {
	return structToMap(input, ignoreNilField, false)
}

func StructPointerToMap(input interface{}, ignoreNilField bool) map[string]interface{} {
	return structToMap(input, ignoreNilField, false)
}

func StructPointerToMapSnakeCase(input interface{}, ignoreNilField bool) map[string]interface{} {
	return structToMap(input, ignoreNilField, true)
}

// GetJsonStringify converts a struct to a JSON string, excluding specified fields.
func GetJsonStringify(src interface{}) string {
	byteData, err := sonic.Marshal(src)
	if err != nil {
		return ""
	}
	return string(byteData)
}

func ParseByteToStruct(src []byte, desc interface{}) error {
	return sonic.Unmarshal(src, &desc)
}
