package controllers

import (
	"App/internal/models"
)

type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

type LoginForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// Users Struct for holding Users variables
type Users struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	us       models.UserService
}
