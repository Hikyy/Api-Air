package models

import (
	"App/internal/modules/hash"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
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
	Create(entity interface{}) error
	Update(entity interface{}, attribute string, value string) error
	// SendData(user *interface{}) error
	// Query single user
	ByID(id string, entity interface{}) error
	ByEmail(email string) (*User, error)

	// Ferme Co DB
	Close() error

	// Gère Database Communication
	Ping() error

	GetAllUsers() ([]byte, error)
	AddDataToDb(entity interface{}) error
	GetDataFromDate(start string, end string) ([]byte, error)
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
}

type EntityService struct {
	EntityDB
}

// Validation pour chaque requête DB

type dbConnectionValidator struct {
	EntityDB
	hmac hash.HMAC
}

// En encapsulant un objet "*gorm.DB" dans cette structure, il devient possible de regrouper des fonctionnalités spécifiques
// à la gestion des utilisateurs ou d'ajouter des méthodes personnalisées pour manipuler les données d'utilisateurs dans la base de données.
// Ce qui permet les u *userGorm CreateUser / UpdateUser etc
type DbGorm struct {
	db    *gorm.DB
	dbase *sql.DB
}

type User struct {
	Id         int
	Firstname  string
	Lastname   string
	Email      string    `gorm:"not null;unique_index"`
	Password   string    `gorm:"no null"` // Ne pas store dans la database
	Group_name string    `gorm:"default:'administrator'"`
	CreatedAt  time.Time `gorm:"type:timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp;autoUpdateTime:true"`
}

type Success struct {
	Success bool `json:"success"`
}

type TokenClaim struct {
	Authorized bool `json:"authorized"`
	jwt.StandardClaims
}

type TokenValidity struct {
	Jwt bool
}

type TokenValidityToken struct {
	Jwt    bool
	Cookie *http.Cookie
}

func (c *TokenClaim) Valid() error {
	// Ajoutez ici la validation supplémentaire des revendications si nécessaire
	return c.StandardClaims.Valid()
}

var Cookie = http.Cookie{
	Name:     "TokenBearer",
	Value:    "",
	Path:     "/",
	Expires:  time.Now().Add(time.Minute * 2500),
	HttpOnly: true,
	Secure:   true,
	SameSite: http.SameSiteLaxMode,
}

type UserReturn struct {
	Firstname string
	Lastname  string
	Email     string
}
type SensorDatas struct {
	EventTimestamp time.Time              `json:"tx_time_ms_epoch"`
	EventData      map[string]interface{} `json:"data" gorm:"json"`
	SensorID       int                    `json:"sensor_id"`
}

type userValFunc func(*User) error
type Message struct {
	CmdID              string `json:"cmd_id"`
	DestinationAddress string `json:"destination_address"`
	Ack_Flags          string `json:"ack_flags"`
	Cmd_Type           string `json:"cmd_type"`
}
