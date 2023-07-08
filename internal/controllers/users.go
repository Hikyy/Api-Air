package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"App/internal/helpers"
	"App/internal/models"
	"encoding/json"
)

// Methode create pour ajotuer new user "POST / signup"
func (u *Users) Create(w http.ResponseWriter, r *http.Request, table string) {
	var form SignupForm

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &form); err != nil {
		log.Fatal(err)
	}

	hashedPassword := helpers.HashPassword(form.User_password)

	user := models.User{
		User_firstname: form.User_firstname,
		User_lastname:  form.User_lastname,
		User_email:     form.User_email,
		User_password:  hashedPassword,
	}

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = u.signIn(w, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect to other page after the login
	http.Redirect(w, r, "/login", http.StatusFound)
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

// NewUsers for Parsing new user view/template in signup page
func NewUsers(us models.UserService) *Users {
	return &Users{
		us: us,
	}
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	var form LoginForm
	if err := helpers.ParseForm(r, &form); err != nil {
		panic(err)
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	if err = json.Unmarshal(body, &form); err != nil {
		log.Fatal(err)
	}

	fmt.Println("form", form)
	user, err := u.us.Authenticate(form.User_email, form.User_password)
	fmt.Println("user => ", user, err)

	if err != nil {
		switch err {
		case models.ErrNotFound:
			fmt.Fprintln(w, "Invalid Email Address")
		case models.ErrInvalidPassword:
			fmt.Fprintln(w, "Invalid Password")
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Redirect to other page after the login
	fmt.Fprintf(w, "Login has been succeeded")
	// http.Redirect(w, r, "/cookietest", http.StatusFound)
}
