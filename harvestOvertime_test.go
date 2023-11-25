package harvestovertimelib

import (
	"net/http"
	"testing"
	"time"

	"github.com/sknutsen/harvestovertimelib/models"
)

func TestHarvestOvertime(t *testing.T) {
	client := &http.Client{Timeout: 10 * time.Second}

	settings := models.Settings{
		AccessToken:     "",
		AccountId:       "",
		CarryOverTime:   0,
		OnlyCurrentYear: true,
		TimeOffTasks: []models.Task{
			{
				Id:   0,
				Name: "",
			},
		},
	}

	entries, err := ListEntries(client, settings)

	if err != nil {
		t.Error(err)
	}

	hours := GetTotalOvertime(entries, settings)
	t.Logf("hours: %f", hours)
}
