package runtime

import (
	"fmt"
	"time"

	"github.com/gkits/gosnooze/internal/devices"
)

type Runtime struct {
	disp    devices.Display
	clock   devices.Clock
	buttons [3]devices.Button

	mode Mode

	setTimeBuf time.Time
	setTimeCur uint8

	backlightOffAt time.Time
}

func New(disp devices.Display, clock devices.Clock, buttons [3]devices.Button) *Runtime {
	return &Runtime{
		disp:    disp,
		clock:   clock,
		buttons: buttons,
	}
}

func (run *Runtime) Tick() error {
	switch run.mode {
	case ModeSetTime:
		if err := run.showTime(); err != nil {
			run.mode = ModeShowTime
			return err
		}
	case ModeSetAlarm:
	case ModeShowTime:
		fallthrough
	default:
		if err := run.showTime(); err != nil {
			return err
		}
	}
	return nil
}

func (run *Runtime) ConsumeButtonEvents() {
	// TODO: Implement consumer logic.
}

func (run *Runtime) showTime() error {
	now, err := run.clock.ReadTime()
	if err != nil {
		return fmt.Errorf("failed to read time: %w", err)
	}
	run.disp.PrintTime(now)
	return nil
}

func (run *Runtime) setTime() (err error) {
	// TODO: Implement setTime logic.
	if run.setTimeBuf.IsZero() {
		run.setTimeBuf, err = run.clock.ReadTime()
		if err != nil {
			return fmt.Errorf("failed to read time: %w", err)
		}
	}

	switch run.setTimeCur {
	case 0:
	case 5:
	default:
		run.mode = ModeShowTime
		run.setTimeBuf = time.Time{}
		run.setTimeCur = 0
	}

	return nil
}

func (run *Runtime) setAlarm() error {
	// TODO: Implement setAlarm logic.
	return nil
}
