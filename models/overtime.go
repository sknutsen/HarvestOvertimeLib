package models

import (
	"time"
)

type HoursByWeek struct {
	Week       string
	HoursByDay map[time.Time]HoursByDay
	Hours      float64
}

type HoursByDay struct {
	Date        time.Time
	TimeEntries map[int]TimeEntry
	Hours       float64
}

func (hbw *HoursByWeek) SumOfHours(hoursToAdd float64) float64 {
	hbw.Hours = 0

	for _, hbd := range hbw.HoursByDay {
		for _, e := range hbd.TimeEntries {
			hbw.Hours = hbw.Hours + e.Hours
		}
	}

	hbw.Hours += hoursToAdd

	return hbw.Hours
}
