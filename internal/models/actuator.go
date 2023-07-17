package models

type Actuators struct {
	Id              int    `json:"id"`
	ActuatorName    string `json:"actuator_name"`
	ActuatorCommand string `json:"actuator_command"`
	DataKey         string `json:"data_key"`
	RoomId          int    `json:"room_id"`
}

func (ug *DbGorm) GetAllActuators() ([]Actuators, error) {
	var actuators []Actuators

	db := ug.Db.Table("actuators").Order("id").Find(&actuators)
	if db.Error != nil {
		return nil, db.Error
	}

	return actuators, nil
}