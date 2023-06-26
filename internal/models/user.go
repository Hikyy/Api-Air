package models

import (
	"errors"
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

// Methode Create pour add user to db
func (ug *userGorm) Create(user *User) error {
	return ug.db.Create(user).Error
}