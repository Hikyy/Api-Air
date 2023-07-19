package models

import "fmt"

type Rooms struct {
	RoomId     int       `gorm:"column:room_id;primary_key" json:"id,omitempty"`
	RoomNumber int       `gorm:"column:room_number" json:"room_number"`
	RoomKey    string    `gorm:"column:room_key" json:"room_key"`
	FloorId    int       `gorm:"column:floor_id" json:"floor_id"`
	Sensors    []Sensors `gorm:"foreignKey:RoomID" json:"sensors"`
}

func (ug *DbGorm) GetRooms() ([]Rooms, error) {
	var rooms []Rooms

	db := ug.Db.Table("rooms").Order("room_number").Find(&rooms)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil, db.Error
	}
	return rooms, nil
}

func (ug *DbGorm) GetAllDatasbyRoomBetweenTwoDays(room int, start string, end string) ([]SensorEvent, error) {
	var sensorData []SensorEvent

	err := ug.Db.Model(&SensorEvent{}).
		Select("sensor_events.id, sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.id = sensor_events.sensor_id").
		Joins("LEFT JOIN rooms ON rooms.room_id = sensors.room_id").
		Where("rooms.id = ? AND sensor_events.event_timestamp >= ? AND sensor_events.event_timestamp <= ?", room, start, end).
		Find(&sensorData).Error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", sensorData)

	return sensorData, nil
}

func (ug *DbGorm) GetAllSensorRooms() ([]Rooms, error) {
	var rooms []Rooms

	ug.Db.Preload("Sensors").Find(&rooms)
	fmt.Printf("rooms : %+v\n", rooms)

	return rooms, nil
}
