package handlers

import (
	"App/internal/helpers"
	"App/internal/requests"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
)

type SensorRequestToBroker struct {
	CmdId              int64  `json:"cmd_id"`
	DestinationAddress string `json:"destination_address"`
	AckFlags           int    `json:"ack_flags"`
	CmdType            int    `json:"cmd_type"`
}

func (handlers *HandlerService) YourHTTPHandler(w http.ResponseWriter, r *http.Request) {
	// Créez un canal pour recevoir les signaux OS
	id := chi.URLParam(r, "id")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	var form requests.SensorRequest

	errPayload := ProcessRequest(&form, r, w)

	if errPayload != nil {
		return
	}

	var sensorRequestToBroker SensorRequestToBroker

	sensorRequestToBroker.CmdId = helpers.GenerateUniqueID()
	// sensorRequestToBroker.AckFlags = helpers.GenerateUniqueID()

	helpers.FillStruct(&sensorRequestToBroker, form)

	// Appelez la fonction SendRequest avec les arguments requis
	SendRequest(c, sensorRequestToBroker, id)
}

func HandleSendRequestToSensorForm(w http.ResponseWriter, r *http.Request) map[string]string {
	data := make(map[string]string)

	data["cmd_id"] = r.URL.Query().Get("cmd_id")
	data["destination_address"] = r.URL.Query().Get("destination_address")
	data["ack_flags"] = r.URL.Query().Get("ack_flags")
	data["cmd_type"] = r.URL.Query().Get("cmd_type")
	return data
}
