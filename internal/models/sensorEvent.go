package models

import (
	"fmt"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type SensorEventTable struct {
	Table  string          `json:"table"`
	Action string          `json:"action"`
	Data   SensorEventData `json:"data"`
}

type SensorEventData struct {
	ID             int                    `json:"id"`
	EventTimestamp string                 `json:"event_timestamp"`
	EventData      map[string]interface{} `json:"event_data"`
	SensorID       int                    `json:"sensor_id"`
}

type SensorEvent struct {
	Id             int                    `json:"id"`
	EventTimestamp time.Time              `json:"event_timestamp"`
	SensorID       uint                   `json:"sensor_id"`
	EventData      map[string]interface{} `json:"event_data" gorm:"json"`
	SensorName     string                 `json:"sensor_name"`
	SensorType     string                 `json:"sensor_type"`
	RoomID         int                    `json:"room_id"`
}

type SensorDatas struct {
	Id             int                    `json:"id"`
	EventTimestamp time.Time              `json:"tx_time_ms_epoch"`
	EventData      map[string]interface{} `json:"data" gorm:"json"`
	SensorID       int                    `json:"sensor_id"`
}

func (ug *DbGorm) AddDataToDb(entity *SensorDataToDb, room_key string) error {

	var roomId string
	ug.Db.Table("rooms").Where("room_key = ?", room_key).Select("room_id").Find(&roomId)

	roomIdInt, err := strconv.Atoi(roomId)
	if err != nil {
		fmt.Println("Error during conversion")
		return err
	}

	var idSensor string
	ug.Db.Table("sensors").Where("sensor_id = ?", entity.SensorID).Where("room_id = ?", roomIdInt).Select("id").Find(&idSensor)

	entity.SensorID, err = strconv.Atoi(idSensor)
	if err != nil {
		fmt.Println("Error during conversion")
		return err
	}

	db := ug.Db.Table("sensor_events").Create(&entity)
	return db.Error
}

func (ug *DbGorm) GetAllDatasByRoom(room int) ([]SensorEvent, error) {

	var sensorEvent []SensorEvent

	err := ug.Db.Model(&SensorEvent{}).
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ?", room).
		Find(&sensorEvent).Error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", sensorEvent)

	return sensorEvent, nil
}

func (s *Sensors) AfterFind(tx *gorm.DB) error {
	sensorEvent := SensorEvent{
		SensorID: uint(s.SensorID),
		// SensorName: s.SensorName,
		// SensorType: s.SensorType,
		// RoomID:     s.RoomID,
	}

	s.SensorEvents = append(s.SensorEvents, sensorEvent)

	return nil
}

func (ug *DbGorm) GetDataFromDate(start string, end string, sensorId int) ([]SensorDatas, error) {

	var datas []SensorDatas

	db := ug.Db.Table("sensor_events").Where("event_timestamp >= ? AND event_timestamp <= ?", start, end).Where("sensor_id = ?", sensorId).Find(&datas)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil, db.Error
	}

	return datas, nil
}

func (ug *DbGorm) GetAllDatasbyRoomByDate(room int, start string, end string) ([]SensorEvent, error) {
	var sensorEvent []SensorEvent

	err := ug.Db.Model(&SensorEvent{}).
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ? AND sensor_events.event_timestamp >= ? AND sensor_events.event_timestamp <= ?", room, start, end).
		Find(&sensorEvent).Error

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", sensorEvent)

	return sensorEvent, nil

}

func (ug *DbGorm) GetDatasByIdByRoomByDate(room int, sensors int, start string, end string) ([]SensorEvent, error) {
	var sensorEvent []SensorEvent

	err := ug.Db.Model(&SensorEvent{}).
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ? AND sensor_id = ? AND event_timestamp >= ? AND event_timestamp <= ?", room, sensors, start, end).
		Find(&sensorEvent)
	if err != nil {
		return nil, nil
	}
	return sensorEvent, nil
}
