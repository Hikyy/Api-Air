package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

type SensorRequestToBroker struct {
	CmdId              int    `json:"cmd_id"`
	DestinationAddress string `json:"destination_address"`
	AckFlags           int    `json:"ack_flags"`
	CmdType            int    `json:"cmd_type"`
}

func HandleDestinationAdress(w http.ResponseWriter, r *http.Request, sensorRequestToBroker *SensorRequestToBroker, roomID int, Cmdtype int) {

	models.InitGorm.Db.Table("actuators").
		Where("room_id = ? AND actuator_command = ?", roomID, Cmdtype).
		Select("destination_address").Find(&sensorRequestToBroker.DestinationAddress)

}

func (handlers *HandlerService) YourHTTPHandler(w http.ResponseWriter, r *http.Request) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var form requests.SensorRequest

	body, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var sensorRequestToBroker SensorRequestToBroker

	if err = json.Unmarshal(body, &form); err != nil {
		log.Fatal(err)
	}

	form.Data.Attributes.RoomID, _ = strconv.Atoi(helpers.DecodeId(strconv.Itoa(form.Data.Attributes.RoomID)))
	
	HandleDestinationAdress(w, r, &sensorRequestToBroker, form.Data.Attributes.RoomID, form.Data.Attributes.CmdType)
	
	sensorRequestToBroker.CmdId = int(helpers.GenerateUniqueID())
	sensorRequestToBroker.AckFlags = 0
	sensorRequestToBroker.CmdType = form.Data.Attributes.CmdType

	jsonData, err := json.Marshal(sensorRequestToBroker)
	if err != nil {
		fmt.Println("Erreur lors de la conversion en JSON :", err)
		return
	}

	SendRequest(c, jsonData, form.Data.Attributes.RoomID)
}
