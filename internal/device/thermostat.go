package device

type Thermostat interface {
	SetTempF(n int8) error
	SetTempC(n float64) error
	SetRangeF(min, max int8) error
	SetRangeC(min, max float64) error
}
