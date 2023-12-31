package http

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/dylanconnolly/home-forecast/internal/device"
)

const nestURL = "https://smartdevicemanagement.googleapis.com/v1/enterprises/bbd8babc-56b8-4591-b429-020f5cf9d4bc"

type NestClient struct {
	URL        string
	HttpClient *http.Client
}

func NewNestClient() *NestClient {
	return &NestClient{
		URL:        nestURL,
		HttpClient: NewHttpClient(10),
	}
}

func (nc *NestClient) newRequest(token, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, nc.URL+url, body)
	if err != nil {
		return nil, fmt.Errorf("error creating new request for nest client: %s", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	return req, nil
}

func (nc *NestClient) GetDevices(token string) error {
	req, err := nc.newRequest(token, http.MethodGet, "/devices", nil)
	if err != nil {
		return err
	}

	resp, err := nc.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("error requesting devices: %s", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading nest response: %s", err)
	}

	fmt.Println(string(body))
	return nil
}

func (nc *NestClient) GetDevice(token, id string) (*device.Nest, error) {
	var device device.Nest
	req, err := nc.newRequest(token, http.MethodGet, "/devices/"+id, nil)
	if err != nil {
		return nil, err
	}

	resp, err := nc.HttpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error requesting device ID=%s: %s", id, err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&device); err != nil {
		return nil, fmt.Errorf("error scanning into nest device: %s", err)
	}

	return &device, nil
}
