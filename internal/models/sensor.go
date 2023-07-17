package models

type Sensors struct {
	Id           int           `gorm:"id"`
	SensorID     int           `gorm:"column:sensor_id"`
	SensorName   string        `json:"sensor_name" gorm:"column:sensor_name"`
	SensorType   string        `gorm:"column:sensor_type"`
	RoomID       int           `json:"room_id" gorm:"column:room_id"`
	SensorEvents []SensorEvent `json:"event_data" gorm:"foreignKey:SensorID" gorm:"column:sensor_event"`
}