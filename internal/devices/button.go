package devices

import "machine"

type Button struct {
	machine.Pin
}

func NewButton(pin machine.Pin) Button {
	button := machine.GPIO14
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	return Button{button}
}

func (b Button) IsPressed() bool {
	return b.Pin.Get()
}
