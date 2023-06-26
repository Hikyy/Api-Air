package models

import (
	"errors"

	"github.com/AlyRagab/golang-user-registration/modules/hash"
	"github.com/jinzhu/gorm"
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
	HmacSecret        = "secret-hmac-key"
	_          UserDB = &userGorm{}
)

// UserDB interface handle toute les opérations User dans la DB
// Couche database pour les queries single user

type UserDB interface {
	// Alter
	Create(user *User) error
	Update(user *User) error
	Delete(id uint) error

	// Query single user
	ByID(id uint) (*User, error)
	ByEmail(email string) (*User, error)
	ByRemember(token string) (*User, error)

	// Ferme Co DB
	Close() error

	// Migration Helpers
	// DBDestructiveReset()

	// Gère Database Communication
	Ping() error
}

// UserService interface qui set les methodes utilisée pour le user model

// Database Auth Layer
type UserService interface {
	// Authenticate verifie email et password donné
	// Si correspondance retourne user email
	// Sinon retourne : 
	// ErrNotFound , ErrInvalidPassword ou error 
	Authenticate(email, password string) (*User, error)
	UserDB
}

type userService struct {
	UserDB
}

// Validation pour chaque requête DB

type userValidator struct {
	UserDB
	hmac hash.HMAC
}

type userGorm struct {
	db *gorm.DB
}

type User struct {
	gorm.Model
	Name         string
	Email        string `gorm:"not null;unique_index"`
	Password     string `gorm:"-"` // Ne pas store dans la database
	PasswordHash string `gorm:"not null"`
	Remember     string `gorm:"-"`
	RememberHash string `gorm:"not null;unique_index"`
}

