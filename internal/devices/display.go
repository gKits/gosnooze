package devices

import (
	"fmt"
	"machine"

	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/hd44780i2c"
)

type Display struct {
	hd44780i2c.Device
}

func NewDisplay(i2cbus drivers.I2C) (Display, error) {
	// TODO: Create custom characters using https://maxpromer.github.io/LCD-Character-Creator
	lcd := hd44780i2c.New(machine.I2C0, 0x27)
	if err := lcd.Configure(hd44780i2c.Config{Width: 16, Height: 2}); err != nil {
		return Display{}, fmt.Errorf("failed to configure lcd: %w", err)
	}
	lcd.ClearDisplay()
	lcd.SetCursor(0, 0)
	return Display{lcd}, nil
}

/*
Prints dt to the LCD screen like this:

	+----------------+
	|    03:04:05    |
	| Mo Jan 02 2006 |
	+----------------+
*/
func (lcd Display) PrintTime(dt time.Time) {
	t := dt.Format("03:04:05")
	d := dt.Format("Mon")[:2] + dt.Format(" Jan 02 2006")
	lcd.SetCursor(4, 0)
	lcd.Print([]byte(t))
	lcd.SetCursor(1, 1)
	lcd.Print([]byte(d))
}
