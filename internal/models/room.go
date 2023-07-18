package models

import "fmt"

type Rooms struct {
	Id         int    `json:"id" gorm:"id"`
	RoomNumber int    `json:"room_number" gorm:"room_number"`
	RoomKey    string `json:"room_key" gorm:"room_key"`
	FloorId    int    `json:"floor_id" gorm:"floor_id"`
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
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ? AND sensor_events.event_timestamp >= ? AND sensor_events.event_timestamp <= ?", room, start, end).
		Find(&sensorData).Error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", sensorData)

	return sensorData, nil
}
