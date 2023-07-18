package resources

type SensorResource struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			SensorName string `json:"sensor_id"`
			SensorType string `json:"sensor_type"`
			RoomID     int    `json:"room_id"`
			SensorID   int    `json:"sensor_id"`
		} `json:"attributes"`
	} `json:"data"`
}
