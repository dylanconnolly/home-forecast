package http

import (
	"net/http"
	"time"
)

func NewHttpClient(timeout uint8) *http.Client {
	return &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}
}
