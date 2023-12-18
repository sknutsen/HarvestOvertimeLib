package harvestovertimelib

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sknutsen/harvestovertimelib/v2/models"
)

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

	if settings.UserId != 0 {
		url = fmt.Sprintf("%s&user_id=%d", url, settings.UserId)
	}

	if settings.ProjectId != 0 {
		url = fmt.Sprintf("%s&project_id=%d", url, settings.ProjectId)
	}

	if settings.ClientId != 0 {
		url = fmt.Sprintf("%s&client_id=%d", url, settings.ClientId)
	}

	if settings.TaskId != 0 {
		url = fmt.Sprintf("%s&task_id=%d", url, settings.TaskId)
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
	req.Header.Add("User-Agent", "Harvest overtime")

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
