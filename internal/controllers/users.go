package controllers 

import (
	"fmt"
	"net/http"

	"App/internal/models"
	"App/internal/helpers"
	"encoding/json"
)

// Methode create pour ajotuer new user "POST / signup"
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm

	if err := helpers.ParseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}

	r.ParseForm()
	// decode le json envoyé par le front et l'envoie dans la structure models.User
	decoder := json.NewDecoder(r.Body)
	test := decoder.Decode(&user)
	fmt.Println(test)


	fmt.Println(r)
	fmt.Println("user object", user)

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

// signIn helps in setting the cookie "email" to the end user
func (u *Users) signIn(w http.ResponseWriter, user *models.User) error {

	cookie := http.Cookie{
		Name:     "remember_token",
		// Value:    user.Remember,
		HttpOnly: true, // means that it is not accessible to scripts "to protect against XSS"
	}
	http.SetCookie(w, &cookie)
	return nil
}

// NewUsers for Parsing new user view/template in signup page
func NewUsers(us models.UserService) *Users {
	return &Users{
		us:        us,
	}
}