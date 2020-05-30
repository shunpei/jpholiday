package jpholiday

import (
	"time"
)

func TimeLocation() (*time.Location, error) {
	return time.LoadLocation("Asia/Tokyo")
}

type Holiday struct {
	Date    time.Time
	Summary string
}

type Holidays struct {
	Holidays []Holiday
}

func Get() *Holidays {
	h := GCalendarHolidays()
	return &Holidays{h}
}

func IsHolidayToday() bool {
	return Get().IsHolidayToday()
}
