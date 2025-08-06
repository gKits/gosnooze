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
	buzzer  devices.Buzzer
	buttons [3]devices.Button

	mode Mode

	timeBuf time.Time
	timeCur TimePosition

	backlightTicksOn uint

	ticker *time.Ticker
}

func New(
	interval time.Duration,
	disp devices.Display,
	clock devices.Clock,
	buzzer devices.Buzzer,
	buttons [3]devices.Button,
) *Runtime {
	return &Runtime{
		disp:    disp,
		clock:   clock,
		buttons: buttons,
		buzzer:  buzzer,
		ticker:  time.NewTicker(interval * time.Second),
	}
}

func (run *Runtime) Run() {
	for t := range run.ticker.C {
		if err := run.tick(t); err != nil {
			println(t.Unix(), "[ERR]", "tick encountered error", "err", err.Error())
		}
	}
}

func (run *Runtime) tick(t time.Time) (err error) {
	e := run.getCurrentEvent()
	switch e {
	case EventNone:
		if run.backlightTicksOn > 0 {
			run.backlightTicksOn--
		} else {
			run.disp.BacklightOn(false)
		}
	case EventAlarmFire:
		run.mode = ModeShowTime
		run.resetBuffer()
		// TODO: Play alarm sound.
	default:
		run.disp.BacklightOn(true)
		run.backlightTicksOn = 40
	}

	println(t.Unix(), "[INF]", "mode:", run.mode.String(), "event:", e.String())

	switch run.mode {
	case ModeSetTime:
		err = run.setTime(e)
	case ModeSetAlarm1:
	case ModeSetAlarm2:
	case ModeShowTime:
		fallthrough
	default:
		err = run.showTime(e)
	}

	if errors.Is(err, modeFinished) {
		run.mode = ModeShowTime
		return nil
	}

	if err != nil {
		return err
	}

	return nil
}

func (run *Runtime) showTime(e Event) error {
	now, err := run.clock.ReadTime()
	if err != nil {
		return fmt.Errorf("failed to read time: %w", err)
	}
	run.disp.PrintTime(now, false)

	switch e {
	case EventButtonAPress:
		if err := run.clock.EnableAlarm1(); err != nil {
			return err
		}
	case EventButtonBPress:
	case EventButtonCPress:
		if err := run.clock.EnableAlarm2(); err != nil {
			return err
		}
	case EventButtonAHold:
		run.mode = ModeSetAlarm1
	case EventButtonBHold:
		run.mode = ModeSetTime
	case EventButtonCHold:
		run.mode = ModeSetAlarm2
	}
	return nil
}

func (run *Runtime) setTime(e Event) (err error) {
	if run.timeBuf.IsZero() {
		run.timeBuf, err = run.clock.ReadTime()
		if err != nil {
			return fmt.Errorf("failed to read time: %w", err)
		}
	}

	switch e {
	case EventButtonAPress:
		run.timeBuf = modifyTimePosition(run.timeBuf, run.timeCur, -1)
	case EventButtonCPress:
		run.timeBuf = modifyTimePosition(run.timeBuf, run.timeCur, 1)
	case EventButtonBPress:
		run.timeCur++
		if run.timeCur == TimeCursorOutOfBounds {
			run.timeCur = 0
			run.clock.SetTime(run.timeBuf)
			run.timeBuf = time.Time{}
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
	// TODO: Add logic to differentiate button presses from holds.
	switch {
	case run.clock.IsAlarm1Fired(), run.clock.IsAlarm2Fired():
		return EventAlarmFire
	case run.buttons[0].IsPressed():
		return EventButtonAPress
	case run.buttons[1].IsPressed():
		return EventButtonBPress
	case run.buttons[2].IsPressed():
		return EventButtonCPress
	default:
		return EventNone
	}
}

func (run *Runtime) resetBuffer() {
	run.timeCur = 0
	run.timeBuf = time.Time{}
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
