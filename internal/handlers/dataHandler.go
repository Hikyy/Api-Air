package handlers

import (
	"App/internal/models"
	"fmt"
	"net/http"
)

func (dt *Datas) Create(w http.ResponseWriter, r *http.Request) {
	var datas models.SensorData
	if err := dt.dts.Create(&datas); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("success => ", &datas)
}
