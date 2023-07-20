package handlers

import (
	"App/internal/helpers"
	"App/internal/models"
	"App/internal/requests"
	"App/internal/resources"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (handler *HandlerService) IndexCondition(w http.ResponseWriter, r *http.Request) {
	condition, _ := handler.use.GetAllConditions()

	var conditionResource []resources.ConditionResource

	resources.GenerateResource(&conditionResource, condition, w)
}

func (handler *HandlerService) StoreCondition(w http.ResponseWriter, r *http.Request) {
	var form requests.ConditionRequest

	errPayload := ProcessRequest(&form, r, w)
	if errPayload != nil {
		return
	}

	var condition models.Conditions

	helpers.FillStruct(&condition, form.Data.Attributes)

	strToInt := strconv.FormatInt(int64(condition.SensorId), 10)

	decoded := helpers.DecodeId(strToInt)

	condition.SensorId, _ = strconv.Atoi(decoded)

	id := helpers.DecodeId(strconv.FormatInt(int64(condition.ActuatorId), 10))

	condition.ActuatorId, _ = strconv.Atoi(id)

	if err := handler.use.AddCondition(&condition); err != nil {
		fmt.Println("err:", err)
		successStatus, _ := json.Marshal(err)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(successStatus)
		return
	}

	var actuatorsResource resources.ConditionResource

	resources.GenerateResource(&actuatorsResource, condition, w)
}
