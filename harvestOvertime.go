package harvestovertimelib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sknutsen/harvestovertimelib/models"
)

func GetTotalOvertime(entries models.TimeEntries, settings models.Settings) float64 {
	var sum float64 = 0.0
	var overtime float64 = 0.0
	var dates []string = []string{}

	filteredList := filterTimeOff(entries, settings)

	for i := 0; i < len(filteredList); i++ {
		sum = filteredList[i].Hours + sum

		dates = appendDate(dates, filteredList[i].SpentDate)
	}

	fmt.Printf("Number of dates: %d\n", len(dates))

	overtime = sum - (float64(len(dates)) * 7.5)

	return addCarryOver(overtime, settings)
}

func appendDate(dates []string, date string) []string {
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

func addCarryOver(overtime float64, settings models.Settings) float64 {
	return overtime + settings.CarryOverTime
}

func filterTimeOff(entries models.TimeEntries, settings models.Settings) []models.TimeEntry {
	var filteredList []models.TimeEntry = []models.TimeEntry{}

	for i := 0; i < len(entries.TimeEntries); i++ {
		exists := false

		for j := 0; j < len(settings.TimeOffTasks); j++ {
			if entries.TimeEntries[i].Task.Id == settings.TimeOffTasks[j].Id {
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
	var url string = fmt.Sprintf("https://api.harvestapp.com/api/v2/time_entries?from=%d-01-01", time.Now().Year())
	var counter int32 = 0

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
			_, exists := m[e.Task.Id]

			if !exists {
				m[e.Task.Id] = e.Task
				tasks = append(tasks, e.Task)
			}
		}
	}

	counter++

	println(counter)

	fmt.Printf("Tasks: %d", len(tasks))

	return tasks, nil
}
