package harvestovertimelib

import (
	"fmt"
	"net/http"

	"github.com/sknutsen/harvestovertimelib/v2/models"
)

func ListTasks(client *http.Client, settings models.Settings) ([]models.TaskDetails, error) {
	var tasks []models.TaskDetails
	var counter uint64 = 0

	var m = make(map[uint64]models.TaskDetails)
	newEntries, err := ListEntries(client, settings)
	if err == nil {
		for _, e := range newEntries.TimeEntries {
			_, exists := m[e.Task.ID]

			if !exists {
				details := models.TaskDetails{
					Task:    e.Task,
					Project: e.Project,
					Client:  e.Client,
				}
				m[e.Task.ID] = details
				tasks = append(tasks, details)
			}
		}
	}

	counter++

	println(counter)

	fmt.Printf("Tasks: %d", len(tasks))

	return tasks, nil
}
