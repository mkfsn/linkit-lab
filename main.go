package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

type Config struct {
	SerialPort string
	Pin        string
}

func (c Config) Validate() error {
	if c.SerialPort == "" {
		return errors.New("no series port")
	}
	if c.Pin == "" {
		return errors.New("no pin")
	}
	return nil
}

func (c Config) String() string {
	return fmt.Sprintf("{SerialPort: %q, Port: %q}", c.SerialPort, c.Pin)
}

func main() {
	var config Config
	flag.StringVar(&config.SerialPort, "serial-port", "", "Serial Port to connect")
	flag.StringVar(&config.Pin, "pin", "", "Pin")
	flag.Parse()

	if err := config.Validate(); err != nil {
		log.Printf("Error: %s\n", err)
		flag.Usage()
		return
	}
	log.Printf("Config: %+v\n", config)

	firmataAdaptor := firmata.NewAdaptor(config.SerialPort)
	led := gpio.NewGroveLedDriver(firmataAdaptor, "13")
	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}
	robot := gobot.NewRobot("bot",
		[]gobot.Connection{firmataAdaptor},
		[]gobot.Device{led},
		work,
	)

	if err := robot.Start(); err != nil {
		log.Printf("Error: %v\n", err)
	}
}
