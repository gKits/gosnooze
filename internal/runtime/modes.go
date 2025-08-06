package runtime

type Mode uint8

const (
	ModeShowTime Mode = iota
	ModeSetTime
	ModeSetAlarm1
	ModeSetAlarm2
)

func (m Mode) String() string {
	switch m {
	case ModeShowTime:
		return "mode_show-time"
	case ModeSetTime:
		return "mode_set-time"
	case ModeSetAlarm1:
		return "mode_set-alarm1"
	case ModeSetAlarm2:
		return "mode_set-alarm2"
	}
	return ""
}

type Event uint8

const (
	EventNone Event = iota
	EventButtonAPress
	EventButtonBPress
	EventButtonCPress
	EventButtonAHold
	EventButtonBHold
	EventButtonCHold
	EventAlarmFire
)

func (e Event) String() string {
	switch e {
	case EventNone:
		return "event_none"
	case EventButtonAPress:
		return "event_button-a-press"
	case EventButtonBPress:
		return "event_button-b-press"
	case EventButtonCPress:
		return "event_button-c-press"
	case EventButtonAHold:
		return "event_button-a-hold"
	case EventButtonBHold:
		return "event_button-b-hold"
	case EventButtonCHold:
		return "event_button-c-hold"
	case EventAlarmFire:
		return "event_alarm-fire"
	}
	return ""
}

type TimePosition uint8

const (
	TimeCursorHours TimePosition = iota
	TimeCursorMinutes
	TimeCursorSeconds
	TimeCursorDay
	TimeCursorMonth
	TimeCursorYear
	TimeCursorZone
	TimeCursorOutOfBounds
)
