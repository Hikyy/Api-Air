package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	// "net/http"
)

// Methode create pour ajotuer new user "POST / signup"
func (handler *HandlerService) StoreUser(w http.ResponseWriter, r *http.Request) {
	var form requests.StoreUserRequest

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	var user models.User

	form.Data.Attributes.Password = helpers.HashPassword(form.Data.Attributes.Password)

	helpers.FillStruct(&user, form.Data.Attributes)

	fmt.Printf("user: %+v\n", user)

	if err := handler.use.Create(&user, w); err != nil {
		success := models.Success{Success: false}

		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
	}
	var userResource resources.UserResource

	resources.GenerateResource(&userResource, user, w)
}

func (handler *HandlerService) Login(w http.ResponseWriter, r *http.Request) {
	var form requests.UserLoginRequest

	errPayload := ProcessRequest(&form, r, w)
	if errPayload != nil {
		return
	}

	user, err := handler.use.Authenticate(form.Data.Attributes.Email, form.Data.Attributes.Password)

	if err != nil {
		success := models.Success{Success: false, Message: "Email or password is incorrect"}
		successStatus, _ := json.Marshal(success)

		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	fmt.Printf("user: %+v\n", *user)

	handler.setCookieFromJWT(w, user.Email)

	var userResource resources.UserResource

	resources.GenerateResource(&userResource, user, w)
}

func (handler *HandlerService) IndexProfils(w http.ResponseWriter, r *http.Request) {
	users, err := handler.use.GetAllUsers()

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		err, _ := json.Marshal(err)
		w.Write(err)
		return
	}

	var userResource []resources.UserResource

	resources.GenerateResource(&userResource, users, w)
}

func (handler *HandlerService) Update(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	role := r.URL.Query().Get("role")

	id = helpers.DecodeId(id)

	fmt.Println("id:", id)

	user := models.User{}
	err := handler.use.ByID(id, &user)

	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}

	fmt.Println(user)

	err = handler.use.Update(&user, "group_name", role, w)
	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}

	var userResource resources.UserResource

	resources.GenerateResource(&userResource, user, w)
}
