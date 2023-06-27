package handlers

import (
	"App/internal/requests"
	// "encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"fmt"
)

func printFieldKeysAndValues(value reflect.Value) {
	request := value.Type()
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := request.Field(i)

		if fieldType.Type.Kind() == reflect.Struct {
			if fieldType.Tag.Get("validate") == "" {
			}
			printFieldKeysAndValues(field)
		} else {
			if fieldType.Tag.Get("validate") != "" {
				requests.Rule(fieldType.Tag.Get("validate"), field)
			}
		}
	}
}
