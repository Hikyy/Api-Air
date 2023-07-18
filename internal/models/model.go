package models

import (
	"App/internal/modules/hash"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	// "gorm.io/datatypes"
)

var (
	// ErrNotFound retourné si ressource absente database
	ErrNotFound = errors.New("Models: Resource Not Found")
	// ErrInvalidID utilisée quand on passe un ID à la méthode Delete pour delete un user de la DB
	ErrInvalidID = errors.New("Models: ID must be Valid ID")
	// UserPwPepper ajouté pepper value
	UserPwPepper = "secret-random-string"
	// ErrInvalidPassword pour retourne invalide password
	ErrInvalidPassword = errors.New("Models: Invalid Password")
	// HmacSecret for creating the HMAC
	HmacSecret          = "secret-hmac-key"
	_          EntityDB = &DbGorm{}
)

// UserDB interface handle toute les opérations User dans la DB
// Couche database pour les queries single user

type EntityDB interface {
	// Alter
	Create(entity interface{}, w http.ResponseWriter) error
	Update(entity interface{}, attribute string, value string, w http.ResponseWriter) error
	// SendData(user *interface{}) error
	// Query single user
	ByID(id string, entity interface{}) error
	ByEmail(email string) (*User, error)
	// Ferme Co DB
	Close() error

	// Gère Database Communication
	Ping() error

	GetAllUsers() ([]User, error)
	AddDataToDb(entity *SensorDataToDb, room_key string) error
	GetDataFromDate(start string, end string, id int) ([]SensorDatas, error)
	GetRooms() ([]Rooms, error)
	GetAllDatasByRoom(room int) ([]SensorEvent, error)
	GetAllDatasbyRoomByDate(room int, start string, end string) ([]SensorEvent, error)
	GetAllDatasbyRoomBetweenTwoDays(room int, start string, end string) ([]SensorEvent, error)
	AddCondition(entity interface{}) error
	GetAllConditions() ([]Conditions, error)
	GetAllActuators() ([]Actuators, error)
	GetDatasByIdByRoomByDate(sensor int, room int, start string, end string) ([]SensorEvent, error)
}

// EntityImplementService interface qui set les methodes utilisée pour le user model

// Database Auth Layer
type EntityImplementService interface {
	// Authenticate verifie email et password donné
	// Si correspondance retourne user email
	// Sinon retourne :
	// ErrNotFound , ErrInvalidPassword ou error
	Authenticate(email, password string) (*User, error)
	EntityDB
	// DbGorm
}

type EntityService struct {
	EntityDB
	db *gorm.DB
}

// Validation pour chaque requête DB

type dbConnectionValidator struct {
	EntityDB
	hmac hash.HMAC
}

type DbGorm struct {
	Db    *gorm.DB
	Dbase *sql.DB
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
	db := ug.Db.Where("id = ?", id).First(entity)
	err := first(db, entity)
	return err
}

// Update method to update a user in database
func (ug *DbGorm) Update(entity interface{}, attribute string, value string, w http.ResponseWriter) error {
	entityValue := reflect.Indirect(reflect.ValueOf(entity)) // Dereference the pointer if entity is a pointer

	idField := entityValue.FieldByName("Id")

	if !idField.IsValid() || idField.Kind() != reflect.Int || idField.Int() == 0 {
		return ErrInvalidID
	}

	id := strconv.Itoa(idField.Interface().(int))
	fmt.Println(ug.Db.Model(&entity).Where("id = ?", id).Update(attribute, value).Error)

	db := ug.Db.Model(&entity).Clauses(clause.Returning{Columns: []clause.Column{{Name: "group_name"}}}).Where("id = ?", id).Update(attribute, value)

	if db.Error != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return db.Error
	}

	w.WriteHeader(http.StatusCreated)
	return nil
}
