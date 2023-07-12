package handlers

import (
	"App/internal/auth"
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
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

	helpers.FillStruct(requests.StoreUserRequest{}, models.User{})

	token, erro := auth.GenerateJWT(user.Email, user.Id)
	if erro != nil {
		fmt.Println(erro)
	}

	w.Header().Set("Token", token)
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := u.signIn(w, &user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	type Test struct {
		Validite bool
	}
	token := r.Header.Get("Token")
	decoded, erro := auth.DecodeJWT(token)
	//
	//unvalid := Test{Validite: decoded.Authorized}
	//novalid, _ := json.Marshal(unvalid)

	if erro != nil {

		return
	}

	var form requests.UserLoginRequest

	ProcessRequest(&form, r, w)

	user, err := u.us.Authenticate(form.Data.Attributes.Email, form.Data.Attributes.Password)
	if err != nil {

	}
	fmt.Println("user => ", user, err)
	if err != nil {

		return
	}

	validity := Test{Validite: decoded.Authorized}
	valid, _ := json.Marshal(validity)

	w.Write(valid)

	err = u.signIn(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

func (u *Users) GetAll(w http.ResponseWriter, r *http.Request) {
	users, _ := u.us.GetAllUsers()
	w.Write(users)
}

func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	user := models.User{}
	err := u.us.ByID(id, &user)
	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}
	jsonUser, _ := json.Marshal(user)

	fmt.Println(user)
	test := u.us.Update(&user, "group_name", "admin")
	testJSON, _ := json.Marshal(test)
	w.Write(testJSON)
	w.Write(jsonUser)
}
