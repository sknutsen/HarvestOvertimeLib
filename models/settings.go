package models

import "time"

type Settings struct {
	AccessToken              string         `json:"accessToken"`
	AccountId                string         `json:"accountId"`
	CarryOverTime            float64        `json:"carryOverTime"`
	TimeOffTasks             []Task         `json:"timeOffTasks"`
	FromDate                 string         `json:"fromDate"`
	ToDate                   string         `json:"toDate"`
	WorkDays                 []time.Weekday `json:"workDays"`
	DaysInWeek               int            `json:"daysInWeek"`
	WorkDayHours             float32        `json:"workDayHours"`
	SimulateFullWeekAtToDate bool           `json:"simulateFullWeekAtToDate"`
}
