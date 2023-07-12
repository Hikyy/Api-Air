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

type LightData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Lux float32 `json:"lux"`
	} `json:"data"`
}

type PresenceData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Person int `json:"persons"`
	} `json:"data"`
}

type CO2Data struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		CO2 int `json:"co2"`
	} `json:"data"`
}

type HeatData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Heat int `json:"data"`
	} `json:"data"`
}

type TemperatureData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Temperature float32 `json:"temperature"`
	} `json:"data"`
}

type ClimData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		AC int `json:"ac"`
	} `json:"data"`
}

type NiveauData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Level int `json:"data"`
	} `json:"data"`
}

type HumidityData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Humidity int `json:"humidity"`
	} `json:"data"`
}

type ADCData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		ADC int `json:"adc"`
	} `json:"data"`
}

type MotionData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Motion int `json:"motion"`
	} `json:"data"`
}

type VentData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Vent int `json:"vent"`
	} `json:"data"`
}

type LightActionData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		Light bool `json:"light"`
	} `json:"data"`
}

type AtmosphericPressureData struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		AtmosphericPressure int `json:"atmospheric_pressure"`
	} `json:"data"`
}

type Consumption struct {
	SourceAdresse int `json:"source_adresse"`
	SensorId      int `json:"sensor_id"`
	Time          int `json:"tx_time_ms_epoch"`
	Data          struct {
		KWH int `json:"kwh"`
	} `json:"data"`
}
