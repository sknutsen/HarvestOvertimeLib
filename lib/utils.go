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

func AppendDate(dates []time.Time, date time.Time) []time.Time {
	var exists bool = false

	for i := 0; i < len(dates); i++ {
		if dates[i] == date {
			exists = true
			break
		}
	}

	if !exists {
		dates = append(dates, date)
	}

	return dates
}

func GetWeekInYearAsString(date time.Time) string {
	year, week := date.ISOWeek()

	return fmt.Sprintf("%d-%d", year, week)
}

func GetDateOfWeekday(lastDate time.Time, weekDay time.Weekday) string {
	date := lastDate.AddDate(0, 0, int(lastDate.Day())-int(lastDate.Weekday())+int(weekDay))

	// switch weekDay {
	// case time.Monday:

	// case time.Tuesday:
	// case time.Wednesday:
	// case time.Thursday:
	// case time.Friday:
	// case time.Saturday:
	// case time.Sunday:
	// }

	return fmt.Sprintf("%d-%d-%d", date.Year(), date.Month(), date.Day())
}
