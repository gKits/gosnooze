package devices

import (
	"errors"
	"fmt"
	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ds3231"
)

type Clock struct {
	ds3231.Device
}

func NewClock(i2cbus drivers.I2C) (Clock, error) {
	rtc := ds3231.New(i2cbus)
	if ok := rtc.Configure(); !ok {
		return Clock{}, errors.New("failed to configure rtc")
	}
	if err := rtc.SetRunning(true); err != nil {
		return Clock{}, fmt.Errorf("rtc: failed to configure: %w", err)
	}
	if !rtc.IsTimeValid() {
		if err := rtc.SetTime(time.Date(2001, time.May, 15, 9, 0, 0, 0, time.UTC)); err != nil {
			return Clock{}, fmt.Errorf("rtc: failed to set start time: %w", err)
		}
	}
	return Clock{rtc}, nil
}
