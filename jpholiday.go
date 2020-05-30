package jpholiday

import (
	"context"
	"fmt"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"log"
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

func GCalendarHolidays() []Holiday {
	jst, _ := TimeLocation()
	ctx := context.Background()

	service, err := calendar.NewService(ctx,
		option.WithScopes(calendar.CalendarReadonlyScope))
	if err != nil {
		fmt.Println(err)
	}

	today, _ := time.ParseInLocation("2006-01-02", time.Now().Format("2006-01-02"), jst)
	from := today.Format(time.RFC3339)
	events, err := service.Events.List("ja.japanese#holiday@group.v.calendar.google.com").ShowDeleted(false).
		SingleEvents(true).TimeMin(from).MaxResults(250).OrderBy("startTime").Do()

	var h []Holiday
	for _, item := range events.Items {
		date := item.Start.DateTime
		if date == "" {
			date = item.Start.Date
		}
		t, e := time.ParseInLocation("2006-01-02", date, jst)
		if e != nil {
			log.Fatalf("Unable to parse date string: %v", err)
		}
		h = append(h, Holiday{
			Date:    t,
			Summary: item.Summary})
	}
	return h
}
