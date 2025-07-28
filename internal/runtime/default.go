package runtime

import (
	"errors"
	"fmt"
	"time"

	"github.com/gkits/gosnooze/internal/devices"
)

var modeFinished = errors.New("mode is finished")

type Runtime struct {
	disp    devices.Display
	clock   devices.Clock
	buttons [3]devices.Button

	mode Mode

	setTimeBuf time.Time
	setTimeCur TimePosition

	backlightTicksOn uint
}

func New(disp devices.Display, clock devices.Clock, buttons [3]devices.Button) *Runtime {
	return &Runtime{
		disp:    disp,
		clock:   clock,
		buttons: buttons,
	}
}

func (run *Runtime) Tick() error {
	e := run.getCurrentEvent()
	switch e {
	case EventNone:
		if run.backlightTicksOn > 0 {
			run.backlightTicksOn--
		} else {
			run.disp.BacklightOn(false)
		}
	case EventAlarmFiring:
		// TODO: Play alarm sound.
	default:
		run.disp.BacklightOn(true)
		run.backlightTicksOn = 40
	}

	switch run.mode {
	case ModeSetTime:
		if err := run.setTime(e); err != nil {
			run.mode = ModeShowTime
			return err
		}
	case ModeSetAlarm:
	case ModeShowTime:
		fallthrough
	default:
		if err := run.showTime(e); err != nil {
			return err
		}
	}
	return nil
}

func (run *Runtime) showTime(e Event) error {
	now, err := run.clock.ReadTime()
	if err != nil {
		return fmt.Errorf("failed to read time: %w", err)
	}
	run.disp.PrintTime(now)

	switch e {
	case EventButtonAPressed:
		if err := run.clock.EnableAlarm1(); err != nil {
			return err
		}
	case EventButtonBPressed:
	case EventButtonCPressed:
		if err := run.clock.EnableAlarm2(); err != nil {
			return err
		}
	}
	return nil
}

func (run *Runtime) setTime(e Event) (err error) {
	if run.setTimeBuf.IsZero() {
		run.setTimeBuf, err = run.clock.ReadTime()
		if err != nil {
			return fmt.Errorf("failed to read time: %w", err)
		}
	}

	switch e {
	case EventButtonAPressed:
		run.setTimeBuf = modifyTimePosition(run.setTimeBuf, run.setTimeCur, -1)
	case EventButtonCPressed:
		run.setTimeBuf = modifyTimePosition(run.setTimeBuf, run.setTimeCur, 1)
	case EventButtonBPressed:
		run.setTimeCur++
		if run.setTimeCur == TimeCursorOutOfBounds {
			run.setTimeCur = 0
			run.clock.SetTime(run.setTimeBuf)
			run.setTimeBuf = time.Time{}
			return modeFinished
		}
	}

	return nil
}

func (run *Runtime) setAlarm() error {
	// TODO: Implement setAlarm logic.
	return nil
}

func (run *Runtime) getCurrentEvent() Event {
	switch {
	case run.clock.IsAlarm1Fired(), run.clock.IsAlarm2Fired():
		return EventAlarmFiring
	case run.buttons[0].IsPressed():
		return EventButtonAPressed
	case run.buttons[1].IsPressed():
		return EventButtonBPressed
	case run.buttons[2].IsPressed():
		return EventButtonCPressed
	default:
		return EventNone
	}
}

func (run *Runtime) reset() {
	run.setTimeCur = 0
	run.setTimeBuf = time.Time{}
}

func modifyTimePosition(dt time.Time, cur TimePosition, mod int) time.Time {
	year, month, day, hour, minute, second, loc := dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), dt.Second(), dt.Location()
	switch cur {
	case TimeCursorHours:
		hour += mod
	case TimeCursorMinutes:
		minute += mod
	case TimeCursorSeconds:
		second += mod
	case TimeCursorDay:
		day += mod
	case TimeCursorMonth:
		month += time.Month(mod)
	case TimeCursorYear:
		year += mod
	case TimeCursorZone:
		loc = loc
	default:
		return dt
	}
	return time.Date(year, month, day, hour, minute, second, 0, loc)
}
