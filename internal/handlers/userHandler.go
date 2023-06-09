package handlers

import (
	"App/internal/requests"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"fmt"
)

func Login(writer http.ResponseWriter, request *http.Request) {
	var user requests.StoreUserRequest
	body, err := io.ReadAll(request.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &user); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("It's %s\n", body)

	value := reflect.ValueOf(user)
	printFieldKeysAndValues(value)
}


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
