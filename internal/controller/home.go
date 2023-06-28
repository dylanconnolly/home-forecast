package controller

import (
	"fmt"
	"os"
	"time"

	"github.com/dylanconnolly/home-forecast/internal/device"
)

type HomeController struct {
	Thermostat   *device.Nest
	NestService  NestService
	OauthService GoogleOauthService
}

func (c *HomeController) Start() error {
	device, err := c.NestService.GetDevice(c.OauthService.AccessToken(), os.Getenv("DEVICE_ID"))
	if err != nil {
		return fmt.Errorf("error retrieving device on home controller startup: %s", err)
	}

	c.Thermostat = device
	return nil
}

func (c *HomeController) Run() {
	ticker := time.NewTicker(10 * time.Minute)
	// initialCall := time.NewTicker(500 * time.Millisecond)
	minuteInterval := time.NewTicker(60 * time.Second)
	updates := make(chan *device.Nest)
	// defer ticker.Stop()
	// defer minuteInterval.Stop()

	minuteCounter := 0

	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println("Polling at ", t)
				if err := c.PollAndUpdateThermostat(); err != nil {
					fmt.Println(err)
				}
				updates <- c.Thermostat
				// case <-initialCall.C:
				// 	fmt.Println("initial poll")
				// 	if err := c.PollAndUpdateThermostat(); err != nil {
				// 		fmt.Println(err)
				// 	}
				// 	initialCall.Stop()
				// 	updates <- c.Thermostat
				// 	return
				// case <-minuteInterval.C:
				// 	minuteCounter += 1
				// 	fmt.Println(minuteCounter)
			}
		}
	}()

	go func() {
		fmt.Println("doing initial poll")
		if err := c.PollAndUpdateThermostat(); err != nil {
			fmt.Println(err)
		}
		updates <- c.Thermostat
	}()

	go func() {
		for {
			select {
			case <-minuteInterval.C:
				minuteCounter += 1
				fmt.Println(minuteCounter)
			}
		}
	}()

	go func() {
		for {
			select {
			case device := <-updates:
				fmt.Println("Got a device update")
				fmt.Printf("new device data:\n\n %+v", device)
			}
		}
	}()
}

func (c *HomeController) PollAndUpdateThermostat() error {
	device, err := c.NestService.GetDevice(c.OauthService.AccessToken(), os.Getenv("DEVICE_ID"))
	if err != nil {
		return fmt.Errorf("error polling thermostat: %s", err)
	}

	c.Thermostat = device
	return nil
}
