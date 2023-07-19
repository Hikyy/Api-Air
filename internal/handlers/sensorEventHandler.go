package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (handler *HandlerService) IndexRoomSensorEvents(w http.ResponseWriter, r *http.Request) {
	id := helpers.DecodeId(chi.URLParam(r, "id"))

	roomInt, err := helpers.TransformStringToInt(id)
	if err != nil {
		return
	}

	data, err := handler.use.GetAllDatasByRoom(roomInt)

	if err != nil {
		fmt.Println(err)
		return
	}

	var sensorEventResource []resources.SensorDataFromDate

	resources.GenerateResource(&sensorEventResource, data, w)
}

func (handler *HandlerService) IndexRoomSensorEventsByDate(w http.ResponseWriter, r *http.Request) {
	id := helpers.DecodeId(chi.URLParam(r, "id"))

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

	roomInt, err := helpers.TransformStringToInt(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	data, _ := handler.use.GetAllDatasbyRoomByDate(roomInt, start, end)

	var sensorEventResource []resources.SensorDataFromDate

	resources.GenerateResource(&sensorEventResource, data, w)
}

func (handler *HandlerService) IndexRoomSensorEventsBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	startDay := chi.URLParam(r, "date-debut")
	endDay := chi.URLParam(r, "date-fin")

	id := helpers.DecodeId(chi.URLParam(r, "id"))

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

	var sensorEventResource []resources.SensorDataFromDate

	resources.GenerateResource(&sensorEventResource, data, w)
}

func (handler *HandlerService) IndexSensorEvents(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	id := helpers.DecodeId(r.URL.Query().Get("id"))

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

	room := helpers.DecodeId(r.URL.Query().Get("room_id"))
	sensor := helpers.DecodeId(r.URL.Query().Get("sensor_id"))

	date := r.URL.Query().Get("date")

	fmt.Println(room, sensor, date)

	roomInt, err := helpers.TransformStringToInt(room)
	if err != nil {
		return
	}

	sensorInt, err := helpers.TransformStringToInt(sensor)
	if err != nil {
		return
	}

	dateStart, err := helpers.ConvertStringToStartOfDay(date)
	if err != nil {
		return
	}

	dateEnd, err := helpers.ConvertStringToEndOfDay(date)
	if err != nil {
		return
	}

	data, err := handler.use.GetDatasByIdByRoomByDate(roomInt, sensorInt, dateStart, dateEnd)
	if err != nil {
		return
	}

	var sensorDataFromDate []resources.SensorDataFromDate
	resources.GenerateResource(&sensorDataFromDate, data, w)

}

func (handler *HandlerService) IndexSensorEventsByIdByRoomBetweenTwoDate(w http.ResponseWriter, r *http.Request) {

	room := r.URL.Query().Get("room_id")
	sensor := r.URL.Query().Get("sensor_id")
	dateStart := r.URL.Query().Get("start_date")
	dateEnd := r.URL.Query().Get("end_date")

	roomInt, err := helpers.TransformStringToInt(room)
	if err != nil {
		return
	}

	sensorInt, err := helpers.TransformStringToInt(sensor)
	if err != nil {
		return
	}

	dateStartToTime, err := helpers.ConvertStringToStartOfDay(dateStart)
	if err != nil {
		return
	}

	dateEndtoTime, err := helpers.ConvertStringToEndOfDay(dateEnd)
	if err != nil {
		return
	}

	data, err := handler.use.GetDatasByIdByRoomByDate(roomInt, sensorInt, dateStartToTime, dateEndtoTime)
	if err != nil {
		return
	}

	//var sensorDataFromDate []resources.SensorDataFromDate
	//test, _ := json.Marshal(sensorDataFromDate)
	//w.Write(test)

	test := convertToResource(data)
	lol, _ := json.Marshal(test)

	w.Write(lol)

	//resources.GenerateResource(&sensorDataFromDate, data, w)

}

func convertToResource(sensorEvents []models.SensorEventNew) []resources.SensorDataFromDateResource {
	var sensorDataList []resources.SensorDataFromDateResource

	for _, sensorEvent := range sensorEvents {
		resource := resources.SensorDataFromDateResource{
			Data: resources.SensorData{
				Type: "sensor_data",
				Id:   fmt.Sprintf("%d", sensorEvent.SensorID),
				Attributes: resources.SensorDataAttributes{
					EventTimestamp: sensorEvent.EventTimestamp,
					EventData:      sensorEvent.EventData.(map[string]interface{}),
					SensorID:       sensorEvent.SensorID,
				},
			},
		}
		sensorDataList = append(sensorDataList, resource)
	}

	return sensorDataList
}
