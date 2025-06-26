package devices

import "machine"

func SetupButton(pin machine.Pin) machine.Pin {
	button := machine.GPIO14
	button.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	return button
}
