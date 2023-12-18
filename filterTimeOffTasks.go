package harvestovertimelib

import (
	"github.com/sknutsen/harvestovertimelib/v2/models"
)

type FilteredTimeEntries struct {
	TimeOff   []models.TimeEntry
	WorkHours []models.TimeEntry
}

func FilterTimeOffTasks(entries models.TimeEntries, settings models.Settings) FilteredTimeEntries {
	var filteredList FilteredTimeEntries = FilteredTimeEntries{
		TimeOff:   []models.TimeEntry{},
		WorkHours: []models.TimeEntry{},
	}

	for _, e := range entries.TimeEntries {
		exists := false

		for _, t := range settings.TimeOffTasks {
			if e.Task.ID == t.ID {
				exists = true
				break
			}
		}

		if !exists {
			filteredList.WorkHours = append(filteredList.WorkHours, e)
		} else {
			filteredList.TimeOff = append(filteredList.TimeOff, e)

			fillerEntry := e

			fillerEntry.Hours = 0

			filteredList.WorkHours = append(filteredList.WorkHours, fillerEntry)
		}
	}

	return filteredList
}
