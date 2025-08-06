package devices

import (
	"machine"

	"tinygo.org/x/drivers/buzzer"
)

type Buzzer struct {
	buzzer.Device
}

func NewBuzzer(pin machine.Pin) Buzzer {
	return Buzzer{buzzer.New(pin)}
}
