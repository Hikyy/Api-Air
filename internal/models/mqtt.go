package models

import "time"

type SensorData struct {
	SensorID       int                    `json:"sensor_id"`
	EventTimestamp uint                   `json:"tx_time_ms_epoch"`
	Data           map[string]interface{} `json:"data"`
}

type SensorDataToDb struct {
	EventTimestamp time.Time              `json:"tx_time_ms_epoch"`
	EventData      map[string]interface{} `json:"data" gorm:"json"`
	SensorID       int                    `json:"sensor_id"`
}
