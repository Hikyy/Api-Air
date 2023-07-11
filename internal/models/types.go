package models

import (
	"database/sql"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"App/internal/modules/hash"

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
	Group_name string    `gorm:"default:'admin'"`
	CreatedAt  time.Time `gorm:"type:timestamp"`
	UpdatedAt  time.Time `gorm:"type:timestamp;autoUpdateTime:true"`
}

//	type UserJSON struct {
//		Name     string `json:"User_firstname"`
//		LastName string `json:"User_lastname"`
//		Role     string `json:"Group_name"`
//	}
type Sensors struct {
	gorm.Model
	Name string
	Data string `gorm:"type:json"`
}

type UserJSON struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}
type TokenClaim struct {
	Authorized bool `json:"authorized"`
	jwt.StandardClaims
}

func (c *TokenClaim) Valid() error {
	// Ajoutez ici la validation supplémentaire des revendications si nécessaire
	return c.StandardClaims.Valid()
}

type userValFunc func(*User) error
