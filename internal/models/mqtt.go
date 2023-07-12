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

type LambdaData struct {
	SourceAdresse string `json:"source_address"`
	SensorId      int    `json:"sensor_id"`
	Time          int    `json:"tx_time_ms_epoch"`
}

type Data struct {
	LambdaData
	Data struct {
		Data any `json:"lux" json:"persons" json:"co2" json:"heat" json:"temperature" json:"ac" json:"level" json:"humidity" json:"adc" json:"motion" json:"vent" json:"light" json:"atmospheric_pressure" json:"kwh"`
	} `json:"data"`
}
