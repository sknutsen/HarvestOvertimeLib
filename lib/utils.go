package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ParseDateString(date string) time.Time {
	dateParts := strings.Split(date, "-")

	year, _ := strconv.Atoi(dateParts[0])
	month, _ := strconv.Atoi(dateParts[1])
	day, _ := strconv.Atoi(dateParts[2])

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
}

func Contains[T string | time.Weekday](arr []T, val T) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}

	return false
}

func AddHoursForWeek(weeks map[string]float64, week string, hours float64) {
	weeks[week] += hours
}

func AddHoursOnDate(dates map[time.Time]float64, date time.Time, hours float64) {
	dates[date] += hours
}

func AppendDate(dates []time.Time, date time.Time) []time.Time {
	var exists bool = false

	for i := 0; i < len(dates); i++ {
		if dates[i] == date {
			exists = true
			break
		}
	}

	if !exists {
		fmt.Printf("Adding date %s\n", date)
		dates = append(dates, date)
	}

	return dates
}

func GetWeekInYearAsString(date time.Time) string {
	year, week := date.ISOWeek()

	return fmt.Sprintf("%d-%d", year, week)
}
