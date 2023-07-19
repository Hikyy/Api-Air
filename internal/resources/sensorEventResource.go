package resources

import "time"

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

type SensorEventResource struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			EventTimestamp time.Time              `json:"event_timestamp"`
			SensorID       uint                   `json:"sensor_id"`
			EventData      map[string]interface{} `json:"event_data"`
			SensorName     string                 `json:"sensor_name"`
			SensorType     string                 `json:"sensor_type"`
			RoomID         int                    `json:"room_id"`
		} `json:"attributes"`
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
