package models

import "reflect"

func FillStruct(destination interface{}, source interface{}) {
	destinationValue := reflect.ValueOf(destination).Elem()
	sourceValue := reflect.ValueOf(source).Elem()

	for i := 0; i < destinationValue.NumField(); i++ {
		destinationField := destinationValue.Field(i)
		sourceField := sourceValue.Field(i)

		if destinationField.CanSet() && sourceField.IsValid() {
			destinationField.Set(sourceField)
		}
	}
}
