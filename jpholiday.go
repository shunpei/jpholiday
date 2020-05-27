package jpholiday

import (
	"time"
)

func (h *Holidays) IsHolidayToday() bool {
	jst, _ := TimeLocation()
	today := time.Now().In(jst)
	for _, holiday := range h.Holidays {
		if holiday.Date.Equal(today) {
			return true
		}
	}
	return false
}
