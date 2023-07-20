package resources

import (
	"App/internal/helpers"
	"App/internal/models"
	"net/http"
	"time"
)

type SensorDataFromDate struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			EventTimestamp time.Time              `json:"tx_time_ms_epoch"`
			EventData      map[string]interface{} `json:"data"`
			SensorID       int                    `json:"sensor_id"`
		} `json:"attributes"`
	} `json:"data"`
}

type Results struct {
	Data struct {
		Type       string              `json:"type"`
		ID         uint                `gorm:"column:id"`
		Attributes models.SensorEvents `json:"attributes"`
	} `json:"data"`
}

type SensorDataAttributes struct {
	EventTimestamp time.Time              `json:"tx_time_ms_epoch"`
	EventData      map[string]interface{} `json:"data"`
	SensorID       int                    `json:"sensor_id"`
}

type SensorData struct {
	Type       string               `json:"type"`
	Id         string               `json:"id"`
	Attributes SensorDataAttributes `json:"attributes"`
}

type SensorDataFromDateResource struct {
	Data SensorData `json:"data"`
}

func GenerateResourceSensorRoom(data []models.SensorEvents, w http.ResponseWriter) []Results {
	var resource []Results
	for _, event := range data {

		event.ID = uint(convertIdToIdOptmisus(int(event.ID)))
		event.SensorID = helpers.DecodeId(event.SensorID)

		resource = append(resource, Results{
			Data: struct {
				Type       string              `json:"type"`
				ID         uint                `gorm:"column:id"`
				Attributes models.SensorEvents `json:"attributes"`
			}{
				Type: "sensor_event",
				ID:   event.ID,
				Attributes: models.SensorEvents{
					EventTimestamp: event.EventTimestamp,
					Type:           event.Type,
					Value:          event.Value,
					SensorID:       event.SensorID,
				},
			},
		})
	}
	return resource
}
