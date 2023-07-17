package models

type Conditions struct {
	Id             int     `gorm:"id"`
	AutomationName string  `json:"automation_name" gorm:"automation_name"`
	SensorId       int     `json:"sensor_id" gorm:"sensor_id"`
	DataKey        string  `json:"data_key" gorm:"data_key"`
	Operator       string  `json:"operator" gorm:"operator"`
	Value          float64 `json:"value" gorm:"value"`
	ActuatorId     int     `json:"actuator_id" gorm:"actuator_id"`
}

func (ug *DbGorm) AddCondition(entity interface{}) error {
	db := ug.Db.Table("conditions").Create(entity)
	if err := db.Error; err != nil {
		return err
	}
	return nil
}

func (ug *DbGorm) GetAllConditions() ([]Conditions, error) {
	var automations []Conditions

	db := ug.Db.Table("conditions").Order("automation_name").Find(&automations)

	if err := db.Error; err != nil {
		return nil, err
	}

	return automations, nil
}
