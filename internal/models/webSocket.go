package models

type WebSocket struct {
	Table  string `json:"table"`
	Action string `json:"action"`
}

type SensorEventJson struct {
	Table  string     `json:"table"`
	Action string     `json:"action"`
	Data   SensorData `json:"data"`
}

type SensorDataJson struct {
	ID             int         `json:"id"`
	EventTimestamp string      `json:"event_timestamp"`
	EventData      EventValues `json:"event_data"`
	SensorID       int         `json:"sensor_id"`
}

type EventValues struct {
	Vent int `json:"vent"`
}
