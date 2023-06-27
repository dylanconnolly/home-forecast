package device

type Nest struct {
	Name   string     `json:"name"`
	Type   string     `json:"type"`
	Traits NestTraits `json:"traits"`
}

type NestTraits struct {
	Info         NestInfo         `json:"sdm.devices.traits.Info"`
	Humidity     NestHumidity     `json:"sdm.devices.traits.Humidity"`
	Connectivity NestConnectivity `json:"sdm.devices.traits.Connectivity"`
	Fan          NestFan          `json:"sdm.devices.traits.Fan"`
	Eco          EcoTraits        `json:"sdm.devices.traits.ThermostatEco"`
	CurrentTemp  NestCurrentTemp  `json:"sdm.devices.traits.Temperature"`
}

type NestInfo struct {
	CustomName string `json:"customName"`
}

type NestHumidity struct {
	AmbientHumidityPercent float32 `json:"ambientHumidityPercent"`
}

type NestConnectivity struct {
	Status string `json:"status"`
}

type NestFan struct {
	TimerMode string `json:"timerMode"`
}

type EcoTraits struct {
	AvailableModes []string `json:"availableModes"`
	Mode           string   `json:"mode"`
	HeatCelsius    float64  `json:"heatCelsius"`
	CoolCelsius    float64  `json:"coolCelsius"`
}

type NestCurrentTemp struct {
	AmbientTemperatureCelsius float64 `json:"ambientTemperatureCelsius"`
}
