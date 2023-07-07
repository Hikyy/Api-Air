package models

type SensorData struct {
	Current          int
	Voltage          int
	ActivePower      int
	FundamentalPower int
	ReactivePower    int
	ApparentPower    int
	Phase            int
}
