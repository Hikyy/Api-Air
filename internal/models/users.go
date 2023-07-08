package models

import (
	"App/internal/modules/hash"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewUserService(connectionInfo string) (UserService, error) {
	ug, err := newUserGorm(connectionInfo)
	if err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(HmacSecret)
	uv := &userValidator{
		UserDB: ug,
		hmac:   hmac,
	}
	return &userService{
		UserDB: uv,
	}, nil
}

// ByEmail pour get user by email
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Table("Users").Where("user_email = ?", email).First(&user)
	err := first(db, &user)
	return &user, err
}

// Implementer et retourner le userGorm
// permet co db
func newUserGorm(connectionInfo string) (*userGorm, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &userGorm{
		db: db,
	}, nil
}

// Methode Create pour add user to db
func (ug *userGorm) Create(user *User) error {
	fmt.Println("Composant Create !!!!")
	// on peut choisir le nom de la table !!!!!!!!!
	return ug.db.Table("Users").Create(user).Error
}

// will query the gorm.DB and get the first item from db and place it into
// dst , if nothing is found , it will return error.
func first(db *gorm.DB, dst interface{}) error {
	err := db.Table("Users").First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// ByID method to get a user by ID
func (ug *userGorm) ByID(id uint) (*User, error) {
	var user User
	db := ug.db.Where("id = ?", id).First(&user)
	err := first(db, &user)
	return &user, err
}

// Update method to update a user in database
func (ug *userGorm) Update(user *User, role string) error {
	return ug.db.Model(&user).Table("Users").Where("user_firstname = ?", &user.User_firstname).Update("group_name", role).Error
}

// Authenticate Method is used for Authenticate and Validate login
func (us *userService) Authenticate(email, password string) (*User, error) {
	// Vlidate if the user is existed in the database or no
	foundUser, err := us.ByEmail(email)
	fmt.Println("foundUser => ", foundUser)
	if err != nil {
		return nil, err
	}
	// Compare the login based in the Hash value
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.User_password), []byte(password))

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
func (ug *userGorm) GetAllUsers() ([]byte, error) {
	var users []User
	db := ug.db.Table("Users").Order("user_firstname").Find(&users)
	if db.Error != nil {
		return nil, db.Error
	}
	fmt.Println(users)
	jsonData, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

// CloseDB to be used as `defer us.db.Close()`
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

func (ug *userGorm) Ping() error {
	if err := ug.db.DB().Ping(); err != nil {
		ug.db.DB().Close()
		return errors.New("Connection to DB is not available")
	}
	fmt.Println("Connection ok (ping)", ug)
	return nil
}
