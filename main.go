package jpholiday

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type Config struct {
	TimeLocation *time.Location
}

var DefaultConfig *Config

func TimeLocation() (*time.Location, error) {
	return time.LoadLocation("Asia/Tokyo")
}

func SetDefaultConfig() *Config {
	config := DefaultConfig
	jst, _ := TimeLocation()
	if config == nil {
		config = &Config{
			TimeLocation: jst}
	}
	return config
}

type Holiday struct {
	Date    time.Time
	Summary string
}

type Holidays struct {
	Holidays []Holiday
}

func Get() *Holidays {
	raw, err := ioutil.ReadFile("./holidays.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var items []map[string]string
	json.Unmarshal(raw, &items)

	var h []Holiday
	for _, item := range items {
		t, _ := time.Parse("2006-01-02 15:04:05 -0700 MST", item["date"])
		h = append(h, Holiday{
			Date:    t,
			Summary: item["summary"],
		})
	}
	return &Holidays{h}
}

func IsHolidayToday() bool {
	return Get().IsHolidayToday()
}
