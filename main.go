package main

import (
	"fmt"
	"machine"
	"time"

	"github.com/gkits/gosnooze/internal/devices"
)

func main() {
	machine.LED.Low()
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.GPIO21, SDA: machine.GPIO20})
	machine.I2C1.Configure(machine.I2CConfig{SCL: machine.GPIO19, SDA: machine.GPIO18})

	button := machine.GPIO4
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	lcd, err := devices.SetupLCD(machine.I2C0)
	if err != nil {
		println("failed to setup lcd device:", err.Error())
		return
	}
	interrupt := machine.GPIO5
	interrupt.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	clock, err := devices.SetupRTC(machine.I2C1, time.Date(2001, time.May, 15, 9, 0, 0, 0, time.UTC))
	if err != nil {
		println("failed to setup clock device:", err.Error())
		return
	}

	for {
		now, err := clock.ReadTime()
		if err != nil {
			println("failed to read time:", err.Error())
			continue
		}

		temp, err := clock.ReadTemperature()
		if err != nil {
			println("failed to read temperature:", err.Error())
			continue
		}

		fmt.Printf("now: %s | temp: %.2f °C | control: %b%b\n", now.String(), float32(temp)/1000.)
		lcd.PrintTime(now)
	}
}
