package handlers

import (
	"App/internal/middlewares"
	"App/internal/requests"
	"github.com/go-chi/chi/v5"
	"net/http"
	"reflect"
)

type Handler struct {
	*chi.Mux
}

func Handlers(mux *chi.Mux) {
	handler := &Handler{
		mux,
	}

	handler.Use(middlewares.FormRequestCall)
	//FormRequestCall()

	handler.Get("/", Login)
	handler.Route("/api/login", func(r chi.Router) {

	})
}

func FormRequestCall() {
	fnValue := reflect.ValueOf(Login)
	fnType := fnValue.Type()
	//var argTypes []reflect.Type

	for i := 0; i < fnType.NumIn(); i++ {
		argType := fnType.In(i)
		if reflect.Type(argType).PkgPath() == "App/internal/requests" {

		}
		//argTypes = append(argTypes, fnType.In(i))
		//fmt.Println("Argument", i, "type:", argType)
		//fmt.Println(reflect.TypeOf(argType))
	}

	//fmt.Println(argTypes)

}

func Safe(formRequest map[string]interface{}, writer http.ResponseWriter) {
	rule := requests.Rule{Result: formRequest, Request: writer}
	rule.Validated()
}
