package models

import "fmt"

type Sensors struct {
	Id           int           `gorm:"id"`
	SensorID     int           `gorm:"column:sensor_id"`
	SensorName   string        `json:"sensor_name" gorm:"column:sensor_name"`
	SensorType   string        `gorm:"column:sensor_type"`
	RoomID       int           `json:"room_id" gorm:"column:room_id"`
	SensorEvents []SensorEvent `json:"event_data" gorm:"foreignKey:SensorID" gorm:"column:sensor_event"`
}

func (ug *DbGorm) GetAllSensors() ([]Sensors, error) {
	var sensor []Sensors

	db := ug.Db.Table("sensors").Order("id").Find(&sensor)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil, db.Error
	}
	return sensor, nil
}
