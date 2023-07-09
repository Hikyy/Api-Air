package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
)

// NewUsers for Parsing new user view/template in signup page
func NewUsers(us models.EntityImplementService) *Users {
	return &Users{
		us: us,
	}
}

// Methode create pour ajotuer new user "POST / signup"
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form requests.StoreUserRequest
	// tx := models.GetTransaction(r.Context())

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	hashedPassword := helpers.HashPassword(form.Data.Attributes.Password)

	user := models.User{
		Lastname:  form.Data.Attributes.Lastname,
		Firstname: form.Data.Attributes.Firstname,
		Email:     form.Data.Attributes.Email,
		Password:  hashedPassword,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := u.signIn(w, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect to other page after the login
	http.Redirect(w, r, "/login", http.StatusFound)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	var form requests.UserLoginRequest

	ProcessRequest(&form, r, w)

	user, err := u.us.Authenticate(form.Data.Attributes.Email, form.Data.Attributes.Password)

	fmt.Println("user => ", user, err)

	if err != nil {
		jsonData, _ := json.Marshal(err)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
		return
	}

	err = u.signIn(w, user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Login has been succeeded")
}

// signIn helps in setting the cookie "email" to the end user
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {

	cookie := http.Cookie{
		Name: "remember_token",
		// Value:    user.Remember,
		HttpOnly: true, // means that it is not accessible to scripts "to protect against XSS"
	}
	http.SetCookie(w, &cookie)
	return nil
}
