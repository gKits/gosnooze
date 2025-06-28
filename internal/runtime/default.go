package runtime

import (
	"fmt"
	"machine"
	"time"

	"context"

	"tinygo.org/x/drivers/ds3231"
	"tinygo.org/x/drivers/hd44780i2c"
)

type Runtime struct {
	disp    hd44780i2c.Device
	clock   ds3231.Device
	buttons [3]machine.Pin

	tickrate time.Duration

	backlightOffAt time.Time
}

func New(disp hd44780i2c.Device, clock ds3231.Device, buttons [3]machine.Pin) *Runtime {
	return &Runtime{
		disp:     disp,
		clock:    clock,
		buttons:  buttons,
		tickrate: 10 * time.Millisecond,
	}
}

func (run *Runtime) ListenButtonEvents(ctx context.Context) <-chan int {
	event := make(chan int)
	defer close(event)

	go func(buttons [3]machine.Pin) {
		for i, but := range buttons {
			if !but.Get() {
				event <- i
			}
		}
	}(run.buttons)
	return event
}

func (run *Runtime) ConsumeButtonEvents() {
	// TODO: Implement consumer logic.
}

func (run *Runtime) DefaultMode() error {
	now, err := run.clock.ReadTime()
	if err != nil {
		return fmt.Errorf("failed to read time: %w", err)
	}

	if now.After(run.backlightOffAt) {
		run.disp.BacklightOn(false)
	}

	return nil
}

func (run *Runtime) EditTimeMode() error {
	// TODO: Implement time editting mode.
	return nil
}

func (run *Runtime) SetAlarmMode() error {
	// TODO: Implement alarm setting mode.
	return nil
}

func (run *Runtime) printTime(t time.Time) {
	run.disp.SetCursor(0, 0)
	run.disp.Print([]byte(t.Format(time.TimeOnly)))
	run.disp.SetCursor(0, 1)
	run.disp.Print([]byte(t.Format(time.DateOnly)))
}
