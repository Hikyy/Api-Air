package common

import "reflect"

type Parameter struct {
	Data          reflect.Value
	ValueRule     string
	AttributeName string
	StructRequest reflect.Value
	FieldType     reflect.StructField
}
