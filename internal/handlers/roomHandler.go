package handlers

import (
	"App/internal/resources"
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
