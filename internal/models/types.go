package models

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Success struct {
	Success bool `json:"success"`
}

type TokenClaim struct {
	Authorized bool `json:"authorized"`
	jwt.StandardClaims
}

type TokenValidity struct {
	Jwt bool
}

type TokenValidityToken struct {
	Jwt    bool
	Cookie *http.Cookie
}

func (c *TokenClaim) Valid() error {
	// Ajoutez ici la validation supplémentaire des revendications si nécessaire
	return c.StandardClaims.Valid()
}

type Datas struct {
	EventTimestamp time.Time              `json:"tx_time_ms_epoch"`
	EventData      map[string]interface{} `json:"data" gorm:"json"`
	SensorID       int                    `json:"sensor_id"`
	RoomId         int                    `json:"room_id"`
}

type userValFunc func(*User) error
