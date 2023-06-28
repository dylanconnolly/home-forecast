package controller

import "github.com/dylanconnolly/home-forecast/internal/device"

type NestService interface {
	GetDevices(token string) error
	GetDevice(token, id string) (*device.Nest, error)
}
