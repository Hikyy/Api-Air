package resources

import (
	"App/internal/helpers"
	"encoding/json"
	"net/http"
)

type BaseResource struct {
	Data struct {
		Type       string        `json:"type"`
		Id         int           `json:"id"`
		Attributes RoomAttribute `json:"attributes"`
	} `json:"data"`
}

func GenerateResource(resource interface{}, model interface{}, w http.ResponseWriter) {
	helpers.FillStruct(resource, model)

	respJson, _ := json.Marshal(resource)

	w.Write(respJson)
}