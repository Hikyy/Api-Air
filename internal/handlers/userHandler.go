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

	success := models.Success{Success: true}
	successStatus, _ := json.Marshal(success)

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

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	if err := u.us.Create(&user); err != nil {
		success = models.Success{Success: false}
		successStatus, _ = json.Marshal(success)
		w.Write(successStatus)
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(successStatus)
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var form requests.UserLoginRequest

	ProcessRequest(&form, r, w)
	success := models.Success{Success: true}
	successStatus, _ := json.Marshal(success)

	user, err := u.us.Authenticate(form.Data.Attributes.Email, form.Data.Attributes.Password)
	if err != nil {
		success = models.Success{Success: false}
		successStatus, _ = json.Marshal(success)
		w.Write(successStatus)
		fmt.Println(err)
	}

	fmt.Println("user => ", user, err)

	returnFront := models.UserReturn{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
	test, _ := json.Marshal(returnFront)

	if err != nil {

		return
	}
	cookie := u.signIn(w, user)
	http.SetCookie(w, &cookie)
	w.Write(successStatus)
	w.Write(test)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *Users) signIn(w http.ResponseWriter, user *models.User) http.Cookie {
	token, err := auth.GenerateJWT(user.Email)
	fmt.Println(user.Email)
	fmt.Println("methode signIn =>", token)
	if err != nil {
		fmt.Println(err)
	}
	cookie := models.Cookie
	cookie.Value = token

	return cookie
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
