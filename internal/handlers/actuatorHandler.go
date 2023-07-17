package handlers

import (
	"App/internal/resources"
	"net/http"
)

func (handlers *HandlerService) IndexActuators(w http.ResponseWriter, r *http.Request) {
	actuator, _ := handlers.use.GetAllActuators()

	var roomsResource []resources.ActuatorResource

	resources.GenerateResource(&roomsResource, actuator, w)
}
