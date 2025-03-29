package reflection_utils

import (
	"fmt"
	"reflect"
)

// ConvertToType attempts to convert a map value to the desired type
func ConvertToType(value interface{}, targetType reflect.Type) (reflect.Value, error) {
	v := reflect.ValueOf(value)

	// Direct type match
	if v.Type() == targetType {
		return v, nil
	}

	// Handle common type conversions
	switch targetType.Kind() {
	case reflect.Int:
		if s, ok := value.(string); ok {
			var i int
			_, err := fmt.Sscanf(s, "%d", &i)
			if err == nil {
				return reflect.ValueOf(i), nil
			}
		}
	case reflect.Float64:
		if s, ok := value.(string); ok {
			var f float64
			_, err := fmt.Sscanf(s, "%f", &f)
			if err == nil {
				return reflect.ValueOf(f), nil
			}
		}
	case reflect.Ptr:
		// Handle pointers, like *int
		if targetType.Elem().Kind() == reflect.Int {
			if s, ok := value.(string); ok {
				var i int
				_, err := fmt.Sscanf(s, "%d", &i)
				if err == nil {
					ptr := reflect.New(targetType.Elem())
					ptr.Elem().SetInt(int64(i))
					return ptr, nil
				}
			}
		}

	default:
		return reflect.ValueOf(value).Convert(targetType), nil
	}
	return reflect.Value{}, fmt.Errorf("unable to convert %v to type %v", value, targetType)
}
