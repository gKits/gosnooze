package main

import (
	"machine"
	"time"

	"github.com/gkits/gosnooze/internal/devices"
	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	machine.LED.Low()
	machine.I2C0.Configure(machine.I2CConfig{SCL: machine.GPIO27, SDA: machine.GPIO26})
	machine.I2C1.Configure(machine.I2CConfig{SCL: machine.GPIO21, SDA: machine.GPIO20})

	button := machine.GPIO4
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	lcd, err := devices.SetupLCD(machine.I2C0)
	if err != nil {
		println("failed to setup lcd device:", err.Error())
		return
	}

	clock, err := devices.SetupRTC(machine.I2C1, time.Date(2001, time.May, 15, 9, 0, 0, 0, time.UTC))
	if err != nil {
		println("failed to setup clock device:", err.Error())
		return
	}

	backlightOn := true

	for {
		if !button.Get() {
			backlightOn = !backlightOn
			lcd.BacklightOn(backlightOn)
		}

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
		println("now:", now.String(), "|", "temp:", temp/1000, "°C")
		lcdPrintTime(&lcd, now)
	}
}

func lcdPrintTime(lcd *hd44780i2c.Device, t time.Time) {
	lcd.SetCursor(0, 0)
	lcd.Print([]byte(t.Format(time.TimeOnly)))
	lcd.SetCursor(0, 1)
	lcd.Print([]byte(t.Format(time.DateOnly)))
}
