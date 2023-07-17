package models

import (
	"encoding/json"
	"gorm.io/gorm/clause"

	// "errors"
	"fmt"
	"reflect"
	"strconv"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// ByEmail pour get user by email
func (ug *DbGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Table("users").Where("email = ?", email).First(&user)

	err := first(db, &user)
	return &user, err
}

// Methode Create pour add user to db
func (ug *DbGorm) Create(entity interface{}) error {
	return ug.db.Table("users").Create(entity).Error
}

// will query the gorm.DB and get the first item from db and place it into
// dst , if nothing is found , it will return error.
func first(db *gorm.DB, entity interface{}) error {
	err := db.First(entity).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// ByID method to get a user by ID
func (ug *DbGorm) ByID(id string, entity interface{}) error {
	db := ug.db.Where("id = ?", id).First(entity)
	err := first(db, entity)
	return err
}

// Update method to update a user in database
func (ug *DbGorm) Update(entity interface{}, attribute string, value string) error {
	entityValue := reflect.Indirect(reflect.ValueOf(entity)) // Dereference the pointer if entity is a pointer

	idField := entityValue.FieldByName("Id")

	if !idField.IsValid() || idField.Kind() != reflect.Int || idField.Int() == 0 {
		return ErrInvalidID
	}
	id := strconv.Itoa(idField.Interface().(int))
	return ug.db.Model(&entity).Clauses(clause.Returning{Columns: []clause.Column{{Name: "group_name"}}}).Where("id = ?", id).Update(attribute, value).Error

}

// Authenticate Method is used for Authenticate and Validate login
func (us *DatabaseProvider) Authenticate(email, password string) (*User, error) {
	// Vlidate if the user is existed in the database or no
	foundUser, err := us.ByEmail(email)
	fmt.Println("foundUser => ", foundUser)
	if err != nil {
		return nil, err
	}
	// Compare the login based in the Hash value
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))

	if err != nil {
		switch err {
		case bcrypt.ErrMismatchedHashAndPassword:
			return nil, ErrInvalidPassword
		case nil:
			return nil, err
		default:
			return nil, err
		}
	}
	return foundUser, nil
}

func (ug *DbGorm) GetAllUsers() ([]byte, error) {
	var users []UserToFront

	db := ug.db.Table("users").Order("firstname").Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (ug *DbGorm) AddDataToDb(entity interface{}) error {
	return ug.db.Table("sensor_events").Create(entity).Error
}

func (ug *DbGorm) GetDataFromDate(start string, end string, sensorId int) ([]byte, error) {

	var datas []SensorDatas

	db := ug.db.Table("sensor_events").Where("event_timestamp >= ? AND event_timestamp <= ?", start, end).Where("sensor_id = ?", sensorId).Find(&datas)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil, db.Error
	}
	fmt.Println(datas)
	jsonData, _ := json.Marshal(datas)
	return jsonData, nil
}

func (ug *DbGorm) GetRooms() ([]byte, error) {
	var rooms []Rooms

	db := ug.db.Table("rooms").Order("room_number").Find(&rooms)
	if db.Error != nil {
		fmt.Println(db.Error)
		return nil, db.Error
	}

	jsonData, _ := json.Marshal(rooms)
	fmt.Println(jsonData)
	return jsonData, nil
}

func (ug *DbGorm) GetAllDatasByRoom(room int) ([]byte, error) {

	var sensorData []SensorEvent

	err := ug.db.Model(&SensorEvent{}).
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ?", room).
		Find(&sensorData).Error

	if err != nil {
		fmt.Println(err)
	}

	jsonData, _ := json.Marshal(sensorData)
	fmt.Printf("%+v\n", sensorData)

	fmt.Println(sensorData)

	return jsonData, nil
}

func (ug *DbGorm) GetAllDatasbyRoomBydate(room int, start string, end string) ([]byte, error) {
	var sensorData []SensorEvent

	err := ug.db.Model(&SensorEvent{}).
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ? AND sensor_events.event_timestamp >= ? AND sensor_events.event_timestamp <= ?", room, start, end).
		Find(&sensorData).Error

	if err != nil {
		fmt.Println(err)
	}
	jsonData, _ := json.Marshal(sensorData)
	fmt.Printf("%+v\n", sensorData)

	fmt.Println(sensorData)

	return jsonData, nil

}

func (ug *DbGorm) GetAllDatasbyRoomBetweenTwoDays(room int, start string, end string) ([]byte, error) {
	var sensorData []SensorEvent

	err := ug.db.Model(&SensorEvent{}).
		Select("sensor_events.event_timestamp, sensor_events.sensor_id, sensor_events.event_data, sensors.sensor_name, sensors.sensor_type, sensors.room_id").
		Joins("LEFT JOIN sensors ON sensors.sensor_id = sensor_events.sensor_id").
		Where("sensors.room_id = ? AND sensor_events.event_timestamp >= ? AND sensor_events.event_timestamp <= ?", room, start, end).
		Find(&sensorData).Error

	if err != nil {
		fmt.Println(err)
	}
	jsonData, _ := json.Marshal(sensorData)
	fmt.Printf("%+v\n", sensorData)

	fmt.Println(sensorData)

	return jsonData, nil
}

func (ug *DbGorm) AddCondition(entity interface{}) error {
	return ug.db.Table("conditions").Create(entity).Error
}

func (ug *DbGorm) GetAllConditions() ([]byte, error) {
	var automations []Condition

	db := ug.db.Table("conditions").Order("automation_name").Find(&automations)
	if db.Error != nil {
		return nil, db.Error
	}

	jsonData, _ := json.Marshal(automations)
	return jsonData, nil
}

func (ug *DbGorm) GetAllActuators() ([]byte, error) {
	var actuators []Actuators

	db := ug.db.Table("actuators").Order("id").Find(&actuators)
	if db.Error != nil {
		er, _ := json.Marshal(actuators)
		return er, nil
	}
	jsonData, _ := json.Marshal(actuators)

	return jsonData, nil
}
