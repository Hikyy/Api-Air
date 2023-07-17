package requests

type Condition struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			AutomationName string              `json:"automation_name" gorm:"automation_name"`
			SensorId       []ConditionSensorId `gorm:"foreignKey:SensorID" gorm:"column:sensor_event"`
			DataKey        string              `json:"data_key" gorm:"data_key"`
			Operator       string              `json:"operator" gorm:"operator"`
			Value          float64             `json:"value" gorm:"value"`
		} `json:"attributes"`
	} `json:"data"`
}

type ConditionSensorId struct {
	SensorId struct{} `json:"sensor_id" gorm:"sensor_id"`
}
