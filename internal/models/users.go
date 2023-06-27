package models

import (
	"fmt"
	"errors"
	"App/internal/modules/hash"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	
)

func NewUserService(connectionInfo string) (UserService, error){
	ug, err := newUserGorm(connectionInfo)
	if	err != nil {
		return nil, err
	}
	hmac := hash.NewHMAC(HmacSecret)
	uv := &userValidator{
		UserDB:ug,
		hmac: hmac,
	}
	return &userService{
		UserDB: uv,
	}, nil
}
// ByEmail pour get user by email
func (ug *userGorm) ByEmail(email string) (*User, error) {
	var user User
	db := ug.db.Where("email = ?", email).First(&user)
	err := first(db, &user)
	return &user, err
}

// Implementer et retourner le userGorm
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
	return ug.db.Create(user).Error
}

// will query the gorm.DB and get the first item from db and place it into
// dst , if nothing is found , it will return error.
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
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
func (ug *userGorm) Update(user *User) error {
	return ug.db.Model(&user).Where("name = ?", &user.Name).Update(&user).Error
}

// Authenticate Method is used for Authenticate and Validate login
func (us *userService) Authenticate(email, password string) (*User, error) {
	// Vlidate if the user is existed in the database or no
	foundUser, err := us.ByEmail(email)
	if err != nil {
		return nil, err
	}
	// Compare the login based in the Hash value
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.PasswordHash), []byte(password+UserPwPepper))
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

// CloseDB to be used as `defer us.db.Close()`
func (ug *userGorm) Close() error {
	return ug.db.Close()
}

func (ug *userGorm) Ping() error {
	if err := ug.db.DB().Ping(); err != nil {
		ug.db.DB().Close()
		return errors.New("Connection to DB is not available")
	}
	fmt.Println("Connection ok", ug)
	return nil
}
