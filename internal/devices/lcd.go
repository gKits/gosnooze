package devices

import (
	"fmt"
	"machine"

	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/hd44780i2c"
)

type LCD struct {
	hd44780i2c.Device
}

func SetupLCD(i2cbus drivers.I2C) (LCD, error) {
	// TODO: Create custom characters using https://maxpromer.github.io/LCD-Character-Creator
	lcd := hd44780i2c.New(machine.I2C0, 0x27)
	if err := lcd.Configure(hd44780i2c.Config{Width: 16, Height: 2}); err != nil {
		return LCD{}, fmt.Errorf("failed to configure lcd: %w", err)
	}
	lcd.ClearDisplay()
	lcd.SetCursor(0, 0)
	return LCD{lcd}, nil
}

func (lcd LCD) PrintTime(t time.Time) {
	lcd.SetCursor(0, 0)
	lcd.Print([]byte(t.Format(time.TimeOnly)))
	lcd.SetCursor(0, 1)
	lcd.Print([]byte(t.Format(time.DateOnly)))
}
