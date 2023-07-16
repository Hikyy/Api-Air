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
	"strconv"
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

	//success := models.Success{Success: true}
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

type response struct {
	Success bool              `json:"success"`
	Data    models.UserReturn `json:"data"`
}

func (u *Users) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var form requests.UserLoginRequest

	ProcessRequest(&form, r, w)

	user, err := u.us.Authenticate(form.Data.Attributes.Email, form.Data.Attributes.Password)
	if err != nil {
		fmt.Println("user => ", user, "err =>>>>", err)

		success := models.Success{Success: false}
		successStatus, _ := json.Marshal(success)
		w.Write(successStatus)
		fmt.Println(err)
		return
	}

	returnFront := models.UserReturn{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}

	// Create a new response object and fill it with data
	resp := response{
		Success: true,
		Data:    returnFront,
	}

	// Marshal the response object into JSON
	respJson, _ := json.Marshal(resp)

	cookie := u.signIn(w, user)
	http.SetCookie(w, &cookie)
	w.Write(respJson) // Send the response

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
	role := r.URL.Query().Get("role")

	user := models.User{}
	err := u.us.ByID(id, &user)
	if err != nil {
		// Gérer l'erreur
		fmt.Println(err)
	}
	jsonUser, _ := json.Marshal(user)

	fmt.Println(user)
	sendToDb := u.us.Update(&user, "group_name", role)
	sendToDbJson, _ := json.Marshal(sendToDb)
	w.Write(sendToDbJson)
	w.Write(jsonUser)
}

func (u *Users) IndexSensorEvents(w http.ResponseWriter, r *http.Request) {
	day := r.URL.Query().Get("day")
	id := r.URL.Query().Get("id")

	idToInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}

	start, err := helpers.ConvertStringToStartOfDay(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := helpers.ConvertStringToEndOfDay(day)
	if err != nil {
		fmt.Println(err)
		return
	}

	datas, err := u.us.GetDataFromDate(start, end, idToInt)
	if err != nil {
		fmt.Println("problèmeeeeeee => ", err)
		return
	}
	w.Write(datas)
}

func (u *Users) IndexRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := u.us.GetRooms()
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Write(rooms)
}

func (u *Users) IndexRoomSensorEvents(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	roomInt, err := helpers.TransformStringToInt(id)

	if err != nil {
		return
	}
	datas, err := u.us.GetAllDatasByRoom(roomInt)

	if err != nil {
		fmt.Println(err)
		return
	}

	w.Write(datas)
}

func (u *Users) IndexRoomSensorEventsByDate(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	date := chi.URLParam(r, "date")

	start, err := helpers.ConvertStringToStartOfDay(date)
	if err != nil {
		fmt.Println(err)
		return
	}

	end, err := helpers.ConvertStringToEndOfDay(date)
	if err != nil {
		fmt.Println(err)
		return
	}

	//roomInt, err := strconv.Atoi(id)
	roomInt, err := helpers.TransformStringToInt(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	datas, err := u.us.GetAllDatasbyRoomBydate(roomInt, start, end)
	w.Write(datas)
}

func (u *Users) IndexRoomSensorEventsBetweenTwoDates(w http.ResponseWriter, r *http.Request) {
	startDay := chi.URLParam(r, "date-debut")
	endDay := chi.URLParam(r, "date-fin")

	id := chi.URLParam(r, "id")

	start, err := helpers.ConvertStringToStartOfDay(startDay)
	if err != nil {
		fmt.Println(err)
	}

	end, err := helpers.ConvertStringToEndOfDay(endDay)
	if err != nil {
		fmt.Println(err)
	}

	roomInt, err := helpers.TransformStringToInt(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	datas, err := u.us.GetAllDatasbyRoomBetweenTwoDays(roomInt, start, end)
	if err != nil {
		return
	}
	w.Write(datas)
}

func (u *Users) StoreActuators(w http.ResponseWriter, r *http.Request) {

}
