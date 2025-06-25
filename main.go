package main

import (
	"fmt"
	"machine"
	"time"

	"tinygo.org/x/drivers/ds3231"
	"tinygo.org/x/drivers/hd44780i2c"
)

func main() {
	machine.LED.Low()
	machine.I2C0.Configure(machine.I2CConfig{})
	machine.I2C1.Configure(machine.I2CConfig{})

	rtc := ds3231.New(machine.I2C0)
	if ok := rtc.Configure(); !ok {
		println("failed to configure rtc")
		return
	}
	if err := rtc.SetRunning(true); err != nil {
		println("failed to start rtc:", err.Error())
		return
	}
	if err := rtc.SetTime(time.Date(2025, time.June, 25, 21, 32, 0, 0, time.UTC)); err != nil {
		println("failed to set rtc time:", err.Error())
		return
	}

	// TODO: Create custom characters using https://maxpromer.github.io/LCD-Character-Creator
	lcd := hd44780i2c.New(machine.I2C1, 0x27)
	if err := lcd.Configure(hd44780i2c.Config{Width: 16, Height: 2}); err != nil {
		println("failed to configure lcd: ", err.Error())
		return
	}
	lcd.ClearDisplay()
	lcd.SetCursor(0, 0)

	tick := time.NewTicker(time.Second)
	defer tick.Stop()
	for range tick.C {
		now, err := rtc.ReadTime()
		if err != nil {
			fmt.Println(err)
			println("failed to read time:", err.Error())
			continue
		}

		temp, err := rtc.ReadTemperature()
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
	lcd.Print([]byte(t.Format("03:04:05")))
	lcd.SetCursor(1, 0)
	lcd.Print([]byte(t.Format("02 Jan 2006")))
	lcd.SetCursor(0, 0)
}
