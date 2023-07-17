package resources

import (
	"App/internal/helpers"
	"encoding/json"
	"net/http"
)

func GenerateResource(resource interface{}, model interface{}, w http.ResponseWriter) {
	helpers.FillStruct(resource, model)

	respJson, _ := json.Marshal(resource)

	w.Write(respJson)
}