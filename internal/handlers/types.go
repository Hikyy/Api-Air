package handlers

import (
	"App/internal/models"
)

// Users Struct for holding Users variables
type Users struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
	us       models.EntityImplementService
}

type Datas struct {
	dts models.EntityDB
}
