package harvestovertimelib

import (
	"net/http"
	"testing"
	"time"

	"github.com/sknutsen/harvestovertimelib/v2/models"
)

func TestHarvestOvertime(t *testing.T) {
	client := &http.Client{Timeout: 10 * time.Second}

	settings := models.Settings{
		AccessToken:   "",
		AccountId:     "",
		CarryOverTime: 0,
		FromDate:      "",
		ToDate:        "",
		DaysInWeek:    5,
		SimulateFullWeekAtToDate: true,
		WorkDays: []time.Weekday{
			time.Monday,
			time.Tuesday,
			time.Wednesday,
			time.Thursday,
			time.Friday,
		},
		WorkDayHours: 7.5,
		TimeOffTasks: []models.Task{
			{
				ID:   0,
				Name: "",
			},
		},
	}

	if settings.AccessToken != "" {
		entries, err := ListEntries(client, settings)

		if err != nil {
			t.Error(err)
		}

		var totalHours float64
		for _, v := range entries.TimeEntries {
			totalHours += v.Hours
		}

		hours := GetTotalOvertime(entries, settings)
		t.Logf("\novertime: %f\ntotal hours worked: %f", hours, totalHours)
	}
}
