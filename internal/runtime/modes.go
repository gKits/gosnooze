package runtime

type Mode uint8

const (
	ModeShowTime Mode = iota
	ModeSetTime
	ModeSetAlarm
)

type Event uint8

const (
	EventNone Event = iota
	EventButtonAPressed
	EventButtonBPressed
	EventButtonCPressed
	EventAlarmFiring
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
