package device

type ThermostatService interface {
	GetCurrentTempF() (float64, error)
	GetCurrentTempC() (float64, error)
	SetTempF(n int8) error
	SetTempC(n float64) error
	SetRangeF(min, max int8) error
	SetRangeC(min, max float64) error
}

type Thermostat struct {
	CurrentTempF int8
	CurrentTempC float64
	HeatMode     bool
	CoolMode     bool
	EcoMode      bool
}
