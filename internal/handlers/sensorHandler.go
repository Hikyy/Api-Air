package handlers

import (
	"App/internal/resources"
	"net/http"
)

func (handler *HandlerService) IndexSensors(w http.ResponseWriter, r *http.Request) {
	actuator, _ := handler.use.GetAllSensors()

	var sensorsResource []resources.SensorResource

	resources.GenerateResource(&sensorsResource, actuator, w)
}
