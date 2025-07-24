package log

import "time"

func Info(t time.Time, msg string) {
	println("[INF]", t.Format(time.DateTime)+":", msg)
}

func Error(t time.Time, msg string) {
	println("[ERR]", t.Format(time.DateTime)+":", msg)
}
