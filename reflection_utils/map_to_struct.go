package reflection_utils

import (
	// "fmt"
	"fmt"
	"reflect"
	"strings"
)

func MapToStruct(data map[string]string, target interface{}) error {
	targetValue := reflect.ValueOf(target).Elem()
	targetType := targetValue.Type()

	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)
		fieldName := field.Name

		// Check for json tag to map keys to field names
		if tag := field.Tag.Get("json"); tag != "" {
			fieldName = strings.Split(tag, ",")[0]
		}

		mapValue, ok := data[fieldName]
		if ok {
			fieldValue := targetValue.FieldByName(field.Name)
			if fieldValue.IsValid() && fieldValue.CanSet() {

				// Field setting logic
				convertedValue, err := ConvertToType(mapValue, fieldValue.Type())
				if err == nil {
					fieldValue.Set(convertedValue)
					// Remove the key from the map after processing
					delete(data, fieldName)
				} else {
					fmt.Printf("Warning: Failed to convert field %s. Map value type: %T, Struct field type: %v. Error: %v\n", fieldName, mapValue, fieldValue.Type(), err)
				}
			}

		}

	}

	// Handle remaining keys in data
	// Any keys left in the 'data' map after processing the struct fields
	// will be added to the 'Args' field of the target struct, if it exists
	argsField := targetValue.FieldByName("Args")
	if argsField.IsValid() && argsField.CanSet() && argsField.Type() == reflect.TypeOf(map[string]string{}) {
		argsField.Set(reflect.ValueOf(data))
	}
	return nil
}
