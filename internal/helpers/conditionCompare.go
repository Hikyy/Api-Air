package helpers

import (
	"App/internal/models"
	"encoding/json"
	"fmt"
)

func CompareConditions(Conditions []models.Conditions, webSocket []byte) {

	var event models.SensorEventJson
	jsonWS := json.Unmarshal(webSocket, &event)

	datas, err := models.GetConditions()
	if err != nil {
		return
	}
	Conditions = datas

	fmt.Println(Conditions)
	fmt.Println("##############")
	fmt.Println(jsonWS)

	for _, condition := range Conditions {
		fmt.Println(condition.DataKey, condition.Operator, condition.Value)
		switch condition.Operator {
		case ">":
			// avant ca il faut verif => si condition.DataKey === JWTWS.DataKey ALORS
			// en gros ici if : JWTWS.Value > condition.Value =>
			fmt.Println("case : condition.DataKey ", condition.DataKey, " EST ", condition.Operator, " A ", condition.Value)
			break
		case "<":
			fmt.Println("case : ", condition.Operator)
			break
		case "=":
			fmt.Println("case : ", condition.Operator)
			break
		}
	}

}
