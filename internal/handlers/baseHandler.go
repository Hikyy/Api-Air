package handlers

import (
	"App/internal/auth"
	"App/internal/models"
	"App/internal/requests"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
	"time"
)

// NewUsers for Parsing new user view/template in signup page
func NewHandler(provider models.EntityImplementService) *HandlerService {
	return &HandlerService{
		use: provider,
	}
}

func (handler *HandlerService) setCookieFromJWT(w http.ResponseWriter, email string) {
	token, err := auth.GenerateJWT(email)

	fmt.Println(email)

	if err != nil {
		fmt.Println(err)
	}

	cookie := http.Cookie{
		Name:     "TokenBearer",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(time.Minute * 2500),
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}

	http.SetCookie(w, &cookie)
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
		success := models.Success{Success: false}
		successStatus, _ := json.Marshal(success)
		writer.WriteHeader(http.StatusUnprocessableEntity)
		writer.Write(successStatus)
		return errFormRequest
	}
	return nil
}
