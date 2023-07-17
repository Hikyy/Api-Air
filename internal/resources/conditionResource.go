package resources

type ConditionResource struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			AutomationName string  `json:"automation_name" gorm:"automation_name"`
			SensorId       int     `json:"sensor_id"`
			DataKey        string  `json:"data_key"`
			Operator       string  `json:"operator"`
			Value          float64 `json:"value"`
			ActuatorId     int     `json:"actuator_id"`
		} `json:"attributes"`
	} `json:"data"`
}
