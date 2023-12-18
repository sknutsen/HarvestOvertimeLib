package harvestovertimelib

import (
	"fmt"
	"sort"
	"time"

	"github.com/sknutsen/harvestovertimelib/v2/lib"
	"github.com/sknutsen/harvestovertimelib/v2/models"
)

func GetTotalOvertime(entries models.TimeEntries, settings models.Settings) float64 {
	var totalHoursForPeriod float64 = 0.0
	var overtime float64 = 0.0
	var hoursByWeek map[string]models.HoursByWeek = map[string]models.HoursByWeek{}

	filteredList := FilterTimeOffTasks(entries, settings)

	var spentDate time.Time
	var week string

	for _, v := range filteredList.WorkHours {
		spentDate = lib.ParseDateString(v.SpentDate)
		week = lib.GetWeekInYearAsString(spentDate)

		hbw, exists := hoursByWeek[week]

		if !exists {
			hoursByWeek[week] = models.HoursByWeek{
				Week: week,
				HoursByDay: map[time.Time]models.HoursByDay{
					spentDate: {
						Date: spentDate,
						TimeEntries: map[int]models.TimeEntry{
							v.ID: v,
						},
						Hours: v.Hours,
					},
				},
				Hours: v.Hours,
			}
		} else {
			hoursByWeek[week] = AddHoursToWeek(hbw, v)
		}
	}

	hoursInWeek := float32(settings.DaysInWeek) * settings.WorkDayHours

	sortedWeeks := make([]string, 0)
	for k, _ := range hoursByWeek {
		sortedWeeks = append(sortedWeeks, k)
	}
	sort.Strings(sortedWeeks)

	for _, w := range sortedWeeks {
		val := hoursByWeek[w]
		if settings.SimulateFullWeekAtToDate {
			addRemainingWorkHoursForWeek(&val, settings)
		}
		fmt.Printf("Week %s\t\thours %f\t\thours in week %f\t\tdiff %f\n", w, val.Hours, hoursInWeek, val.Hours-float64(hoursInWeek))
		totalHoursForPeriod += val.Hours
	}

	totalNumberOfWeeks := len(hoursByWeek)

	overtime = totalHoursForPeriod - (float64(hoursInWeek) * float64(totalNumberOfWeeks))

	overtime = addCarryOver(overtime, settings)

	return overtime
}

func addCarryOver(overtime float64, settings models.Settings) float64 {
	return overtime + settings.CarryOverTime
}

func addRemainingWorkHoursForWeek(hbw *models.HoursByWeek, settings models.Settings) {
	if len(hbw.HoursByDay) < settings.DaysInWeek {
		fmt.Printf("days worked: %d. days in week: %d.\n", len(hbw.HoursByDay), settings.DaysInWeek)
		hbw.Hours = hbw.SumOfHours(float64(settings.DaysInWeek-len(hbw.HoursByDay)) * float64(settings.WorkDayHours))
	}
}

func AddHoursToWeek(hbw models.HoursByWeek, entry models.TimeEntry) models.HoursByWeek {
	date := lib.ParseDateString(entry.SpentDate)

	hbd, exists := hbw.HoursByDay[date]

	if exists {
		hbd.Hours += entry.Hours
	} else {
		hbw.HoursByDay[date] = models.HoursByDay{
			Date:        date,
			TimeEntries: make(map[int]models.TimeEntry),
			Hours:       entry.Hours,
		}
	}

	hbw.Hours += entry.Hours
	hbw.HoursByDay[date].TimeEntries[entry.ID] = entry

	return hbw
}
