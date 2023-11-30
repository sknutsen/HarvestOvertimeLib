package harvestovertimelib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sknutsen/harvestovertimelib/lib"
	"github.com/sknutsen/harvestovertimelib/models"
)

func GetTotalOvertime(entries models.TimeEntries, settings models.Settings) float64 {
	var hoursForCurrentWeek float64
	var totalHoursForPeriod float64 = 0.0
	var overtime float64 = 0.0
	var hoursByWeek map[string]float64 = map[string]float64{}

	filteredList := filterTimeOff(entries, settings)

	var spentDate time.Time
	var week string

	for _, v := range filteredList {
		spentDate = lib.ParseDateString(v.SpentDate)
		week = lib.GetWeekInYearAsString(spentDate)
		fmt.Printf("Adding week %s\n", week)
		lib.AddHoursForWeek(hoursByWeek, week, v.Hours)
	}

	if settings.SimulateFullWeekAtToDate {
		println("Simulating final week in period")

		date := lib.ParseDateString(settings.ToDate)
		weekId := lib.GetWeekInYearAsString(date)
		hoursInFinalWeek := hoursByWeek[weekId]

		hoursByWeek[weekId] = addRemainingWorkHoursForWeek(hoursInFinalWeek, settings)
	}

	hoursInWeek := float32(settings.DaysInWeek) * settings.WorkDayHours

	for key, val := range hoursByWeek {
		fmt.Printf("Week %s\t\thours %f\t\thours in week %f\t\tdiff %f\n", key, val, hoursInWeek, val-float64(hoursInWeek))
		hoursForCurrentWeek = val

		totalHoursForPeriod += hoursForCurrentWeek
	}

	totalNumberOfWeeks := len(hoursByWeek)

	overtime = totalHoursForPeriod - (float64(hoursInWeek) * float64(totalNumberOfWeeks))

	overtime = addCarryOver(overtime, settings)

	return overtime
}

func addCarryOver(overtime float64, settings models.Settings) float64 {
	return overtime + settings.CarryOverTime
}

// TODO: Complete logic
func addRemainingWorkHoursForWeek(hoursInFinalWeek float64, settings models.Settings) float64 {
	// var hours float64

	if hoursInFinalWeek < float64(settings.DaysInWeek)*float64(settings.WorkDayHours) {
		return float64(settings.DaysInWeek) * float64(settings.WorkDayHours)
	} else {
		return hoursInFinalWeek
	}

	// if lib.Contains[time.Weekday](settings.WorkDays, date.Weekday()) {
	// 	switch date.Weekday() {
	// 	case time.Monday:
	// 	}
	// }

	// return hours
}

func filterTimeOff(entries models.TimeEntries, settings models.Settings) []models.TimeEntry {
	var filteredList []models.TimeEntry = []models.TimeEntry{}

	for i := 0; i < len(entries.TimeEntries); i++ {
		exists := false

		for j := 0; j < len(settings.TimeOffTasks); j++ {
			if entries.TimeEntries[i].Task.ID == settings.TimeOffTasks[j].ID {
				exists = true
				break
			}
		}

		if !exists {
			filteredList = append(filteredList, entries.TimeEntries[i])
		}
	}

	return filteredList
}

func ListEntries(client *http.Client, settings models.Settings) (models.TimeEntries, error) {
	var entries models.TimeEntries
	var url string = "https://api.harvestapp.com/api/v2/time_entries"
	var counter int32 = 0
	var fromDate string

	if settings.FromDate != "" {
		fromDate = settings.FromDate
	} else {
		fromDate = fmt.Sprintf("%d-01-01", time.Now().Year())
	}

	url = fmt.Sprintf("%s?from=%s", url, fromDate)

	if settings.ToDate != "" {
		url = fmt.Sprintf("%s&to=%s", url, settings.ToDate)
	}

	for url != "" {
		newEntries, err := listEntries(client, url, settings)
		if err == nil {
			url = newEntries.Links.Next
			println("New url: " + newEntries.Links.Next)

			entries.TimeEntries = append(entries.TimeEntries, newEntries.TimeEntries...)
		}

		counter++

		println(counter)
	}

	return entries, nil
}

func listEntries(client *http.Client, url string, settings models.Settings) (models.TimeEntries, error) {
	var entries models.TimeEntries

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		println("Error creating request: " + err.Error())
		return models.TimeEntries{}, err
	}

	req.Header.Add("Harvest-Account-ID", settings.AccountId)
	req.Header.Add("Authorization", "Bearer "+settings.AccessToken)
	req.Header.Add("User-Agent", "Harvest API Example")

	resp, err := client.Do(req)
	if err != nil {
		println("Error sending request: " + err.Error())
		return models.TimeEntries{}, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&entries)
	if err != nil {
		println("Error decoding response: " + err.Error())
		return models.TimeEntries{}, err
	}

	return entries, nil
}

func ListTasks(client *http.Client, settings models.Settings) ([]models.Task, error) {
	var tasks []models.Task
	var counter uint64 = 0

	var m = make(map[uint64]models.Task)
	newEntries, err := ListEntries(client, settings)
	if err == nil {
		for _, e := range newEntries.TimeEntries {
			_, exists := m[e.Task.ID]

			if !exists {
				m[e.Task.ID] = e.Task
				tasks = append(tasks, e.Task)
			}
		}
	}

	counter++

	println(counter)

	fmt.Printf("Tasks: %d", len(tasks))

	return tasks, nil
}
