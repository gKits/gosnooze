package runtime

type Mode uint8

const (
	ModeShowTime Mode = iota
	ModeSetTime
	ModeSetAlarm1
	ModeSetAlarm2
)

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
