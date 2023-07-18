package handlers

import (
	"App/internal/helpers"
	"App/internal/resources"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (handler *HandlerService) IndexRoomSensorEvents(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	roomInt, err := helpers.TransformStringToInt(id)

	if err != nil {
		return
	}
	data, err := handler.use.GetAllDatasByRoom(roomInt)

	if err != nil {
		fmt.Println(err)
		return
	}

	var sensorEventResource []resources.SensorEventResource

	resources.GenerateResource(&sensorEventResource, data, w)
}

func (handler *HandlerService) IndexRoomSensorEventsByDate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	date := chi.URLParam(r, "date")

	start, err := helpers.ConvertStringToStartOfDay(date)
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := helpers.ConvertStringToEndOfDay(date)
	if err != nil {
		fmt.Println(err)
		return
	}

	//roomInt, err := strconv.Atoi(id)
	roomInt, err := helpers.TransformStringToInt(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := handler.use.GetAllDatasbyRoomByDate(roomInt, start, end)

	var sensorEventResource []resources.SensorEventResource

	resources.GenerateResource(&sensorEventResource, data, w)
}

func (handler *HandlerService) IndexRoomSensorEventsBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	startDay := chi.URLParam(r, "date-debut")
	endDay := chi.URLParam(r, "date-fin")

	id := chi.URLParam(r, "id")

	start, err := helpers.ConvertStringToStartOfDay(startDay)
	if err != nil {
		fmt.Println(err)
	}

	end, err := helpers.ConvertStringToEndOfDay(endDay)
	if err != nil {
		fmt.Println(err)
	}

	roomInt, err := helpers.TransformStringToInt(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := handler.use.GetAllDatasbyRoomBetweenTwoDays(roomInt, start, end)
	if err != nil {
		return
	}

	var sensorEventResource []resources.SensorEventResource

	resources.GenerateResource(&sensorEventResource, data, w)
}

func (handler *HandlerService) IndexSensorEvents(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	id := r.URL.Query().Get("id")

	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	start, err := helpers.ConvertStringToStartOfDay(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := helpers.ConvertStringToEndOfDay(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := handler.use.GetDataFromDate(start, end, idToInt)
	if err != nil {
		fmt.Println("problèmeeeeeee => ", err)
		return
	}

	var sensorDataFromDate []resources.SensorDataFromDate

	resources.GenerateResource(&sensorDataFromDate, data, w)
}

func (handler *HandlerService) IndexSensorEventsByIdByRoomByDate(w http.ResponseWriter, r *http.Request) {
	room := chi.URLParam(r, "room_id")
	sensor := chi.URLParam(r, "sensor_id")
	date := chi.URLParam(r, "date")

	roomString, err := helpers.TransformStringToInt(room)
	if err != nil {
		return
	}

	sensorString, err := helpers.TransformStringToInt(sensor)
	if err != nil {
		return
	}

	dateStart, err := helpers.ConvertStringToStartOfDay(date)
	if err != nil {
		return
	}

	data, err := handler.IndexSensorEventsByIdByRoomByDate()
	if err != nil {
		return
	}
	w.Write(data)
}
