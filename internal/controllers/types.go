package controllers

import (
	"App/internal/models"
)

type SignupForm struct {
	User_firstname string `gorm:"not null"`
	User_lastname  string `gorm:"not null"`
	User_email     string `gorm:"not null;unique_index"`
	User_password  string `gorm:"not null;"`
}

type LoginForm struct {
	User_email    string `gorm:"not null;unique_index"`
	User_password string `gorm:"not null;"`
}

// Users Struct for holding Users variables
type Users struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	us       models.UserService
}
