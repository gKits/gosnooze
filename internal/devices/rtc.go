package devices

import (
	"errors"
	"fmt"
	"time"

	"tinygo.org/x/drivers"
	"tinygo.org/x/drivers/ds3231"
)

func SetupRTC(i2cbus drivers.I2C, startTime time.Time) (ds3231.Device, error) {
	rtc := ds3231.New(i2cbus)
	if ok := rtc.Configure(); !ok {
		return ds3231.Device{}, errors.New("failed to configure rtc")
	}
	if err := rtc.SetRunning(true); err != nil {
		return ds3231.Device{}, fmt.Errorf("rtc: failed to configure: %w", err)
	}
	if err := rtc.SetTime(startTime); err != nil {
		return ds3231.Device{}, fmt.Errorf("rtc: failed to set start time: %w", err)
	}
	return rtc, nil
}
