package runtime

type Mode uint8

const (
	ModeShowTime Mode = iota
	ModeSetTime
	ModeSetAlarm
)
