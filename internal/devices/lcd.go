package devices

import (
	"fmt"
	"machine"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/hd44780i2c"
)

func SetupLCD(i2cbus drivers.I2C) (hd44780i2c.Device, error) {
	// TODO: Create custom characters using https://maxpromer.github.io/LCD-Character-Creator
	lcd := hd44780i2c.New(machine.I2C0, 0x27)
	if err := lcd.Configure(hd44780i2c.Config{Width: 16, Height: 2}); err != nil {
		return hd44780i2c.Device{}, fmt.Errorf("failed to configure lcd: %w", err)
	}
	lcd.ClearDisplay()
	lcd.SetCursor(0, 0)
	return lcd, nil
}
