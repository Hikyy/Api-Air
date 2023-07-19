package handlers

import (
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"net/http"
)

func (handler *HandlerService) IndexRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := handler.use.GetRooms()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(rooms)
	var roomsResource []resources.RoomResource

	resources.GenerateResource(&roomsResource, rooms, w)
}

func (handler *HandlerService) IndexRoomSensors(w http.ResponseWriter, r *http.Request) {
	rooms, _ := handler.use.GetAllSensorRooms()

	baseResources := resources.NewBaseRoomSensorResource(rooms)

	jsonData, err := json.Marshal(baseResources)
	if err != nil {
		fmt.Println("Erreur de sérialisation JSON:", err)
		return
	}

	w.Write(jsonData)
}
