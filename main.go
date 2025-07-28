package main

import (
	"machine"
	"time"

	"github.com/gkits/gosnooze/internal/devices"
	"github.com/gkits/gosnooze/internal/log"
	"github.com/gkits/gosnooze/internal/runtime"
)

const (
	tickrate = 50 * time.Millisecond

	displaySCLPin = machine.GPIO21
	displaySDAPin = machine.GPIO21

	clockSCLPin = machine.GPIO19
	clockSDAPin = machine.GPIO18

	buttonAPin = machine.GPIO4
	buttonBPin = machine.GPIO5
	buttonCPin = machine.GPIO6
)

func main() {
	machine.LED.Low()
	machine.I2C0.Configure(machine.I2CConfig{SCL: displaySCLPin, SDA: displaySDAPin})
	machine.I2C1.Configure(machine.I2CConfig{SCL: clockSCLPin, SDA: clockSDAPin})

	var buttons [3]devices.Button
	buttons[0] = devices.NewButton(buttonAPin)
	buttons[1] = devices.NewButton(buttonBPin)
	buttons[2] = devices.NewButton(buttonCPin)

	lcd, err := devices.NewDisplay(machine.I2C0)
	if err != nil {
		println("failed to setup display device:", err.Error())
	}

	clock, err := devices.NewClock(machine.I2C1)
	if err != nil {
		println("failed to setup clock device:", err.Error())
	}

	rt := runtime.New(lcd, clock, buttons)

	tick := time.NewTicker(tickrate)
	for t := range tick.C {
		if err := rt.Tick(); err != nil {
			log.Error(t, err.Error())
		}
	}
}
