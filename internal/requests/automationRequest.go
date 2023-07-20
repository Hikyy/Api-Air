package requests

type ConditionRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			AutomationName string  `json:"automation_name" gorm:"automation_name"`
			SensorId       int     `json:"sensor_id" gorm:"sensor_id"`
			DataKey        string  `json:"data_key" gorm:"data_key"`
			Operator       string  `json:"operator" gorm:"operator"`
			Value          float64 `json:"value" gorm:"value"`
			ActuatorId     int     `json:"actuator_id" gorm:"actuator_id"`
		} `json:"attributes"`
	} `json:"data"`
}

type ConditionSensorId struct {
	SensorId struct{} `json:"sensor_id" gorm:"sensor_id"`
}

type Conditions struct {
	AutomationName string  `json:"automation_name" gorm:"automation_name"`
	SensorId       int     `json:"sensor_id" gorm:"sensor_id"`
	DataKey        string  `json:"data_key" gorm:"data_key"`
	Operator       string  `json:"operator" gorm:"operator"`
	Value          float64 `json:"value" gorm:"value"`
	ActuatorId     int     `json:"actuator_id" gorm:"actuator_id"`
}
