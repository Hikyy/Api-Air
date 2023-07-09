package handlers

import (
	"App/internal/models"
	"App/internal/requests"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/go-chi/chi/v5"
)

func SetupRouter() http.Handler {
	router := chi.NewRouter()

	// router.Use(middlewares.TransactionMiddleware(models.InitGorm))

	// router.Use(middlewares.TransactionMiddleware)

	userHandler := NewUsers(models.Db)

	route(router, userHandler)

	return router
}

func recursiveExploreStruct(formRequest interface{}, result *[][]string) {
	value := reflect.ValueOf(formRequest)
	typ := reflect.TypeOf(formRequest)

	value = value.Elem()
	typ = typ.Elem()

	for i := 0; i < value.NumField(); i++ {
		fieldValue := value.Field(i)
		fieldType := typ.Field(i)

		if fieldValue.Kind() == reflect.Struct {
			recursiveExploreStruct(fieldValue.Addr().Interface(), result)
		} else if fieldType.Type.Kind() == reflect.Slice {
			processSliceField(fieldValue, result)
		} else {
			res := requests.ApplyRule(fieldType.Tag.Get("validate"), fieldType.Tag.Get("json"), fieldValue, value)
			if res != nil {
				*result = append(*result, res)
			}
		}
	}
}

func processSliceField(field reflect.Value, result *[][]string) {
	for i := 0; i < field.Len(); i++ {
		recursiveExploreStruct(field.Index(i).Addr().Interface(), result)
	}
}

func ProcessRequest(structRequest interface{}, request *http.Request, writer http.ResponseWriter) [][]string {
	body, err := io.ReadAll(request.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &structRequest); err != nil {
		log.Fatal(err)
	}

	var errFormRequest [][]string

	recursiveExploreStruct(structRequest, &errFormRequest)

	if len(errFormRequest) != 0 {
		jsonData, _ := json.Marshal(errFormRequest)

		writer.Header().Set("Content-Type", "application/json")
		writer.Write(jsonData)
		return errFormRequest
	}
	return nil
}
