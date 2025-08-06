package devices

import (
	"machine"
	"time"

	"tinygo.org/x/drivers/buzzer"
)

type Buzzer struct {
	buzzer.Device
}

type BuzzerNote struct {
	Tone, Duration float64
}

func NewBuzzer(pin machine.Pin) Buzzer {
	pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	pin.Set(true)
	return Buzzer{buzzer.New(pin)}
}

func (buz *Buzzer) Play(song []BuzzerNote) {
	for _, note := range song {
		buz.Device.Tone(note.Tone, note.Duration)
		time.Sleep(10 * time.Millisecond)
	}
}
