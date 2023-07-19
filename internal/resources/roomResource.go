package resources

import (
	"App/internal/helpers"
	"App/internal/models"
)

type RoomResource struct {
	Data struct {
		Type       string `json:"type"`
		RoomId     string `json:"id"`
		Attributes struct {
			RoomNumber int `json:"room_number"`
			FloorId    int `json:"floor_id"`
		} `json:"attributes"`
	} `json:"data"`
}

type RoomAttribute struct {
	RoomNumber int              `json:"room_number"`
	RoomKey    string           `json:"room_key"`
	FloorId    int              `json:"floor_id"`
	Sensors    []models.Sensors `json:"sensors"`
}

func NewBaseRoomSensorResource(rooms []models.Rooms) []BaseResource {

	var baseResources []BaseResource

	for _, room := range rooms {
		var sensors []models.Sensors

		for _, sensor := range room.Sensors {
			newSensor := models.Sensors{
				Id:         convertIdToIdOptmisus(sensor.Id),
				SensorID:   convertIdToIdOptmisus(sensor.SensorID),
				SensorName: sensor.SensorName,
				SensorType: sensor.SensorType,
				RoomID:     convertIdToIdOptmisus(sensor.RoomID),
			}

			sensors = append(sensors, newSensor)
		}

		toto := BaseResource{
			Data: struct {
				Type       string        `json:"type"`
				Id         int           `json:"id"`
				Attributes RoomAttribute `json:"attributes"`
			}{
				Type: "Room",
				Id:   convertIdToIdOptmisus(room.RoomId),
				Attributes: RoomAttribute{
					RoomNumber: room.RoomNumber,
					RoomKey:    room.RoomKey,
					FloorId:    convertIdToIdOptmisus(room.FloorId),
					Sensors:    sensors,
				},
			},
		}
		baseResources = append(baseResources, toto)
	}

	return baseResources
}

func convertIdToIdOptmisus(id int) int {
	id, _ = helpers.TransformStringToInt(helpers.EncodeId(id))

	return id
}
