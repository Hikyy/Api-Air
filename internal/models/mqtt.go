package models

type SensorData struct {
	SensorAddress string                 `json:"source_address"`
	SensorID      int                    `json:"sensor_id"`
	TimeEpoch     uint                   `json:"tx_time_ms_epoch"`
	Data          map[string]interface{} `json:"data"`
}
