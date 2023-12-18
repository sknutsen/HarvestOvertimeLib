package harvestovertimelib

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/sknutsen/harvestovertimelib/v2/lib"
	"github.com/sknutsen/harvestovertimelib/v2/models"
)

var client *http.Client = &http.Client{Timeout: 10 * time.Second}

var settings models.Settings = models.Settings{
	AccessToken:              "",
	AccountId:                "",
	CarryOverTime:            0,
	FromDate:                 "2023-01-01",
	ToDate:                   "2023-12-31",
	DaysInWeek:               5,
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
			ID:   10882012,
			Name: "",
		},
	},
}

func setup() {
	err := godotenv.Load(".env")

	if err == nil {
		settings.AccessToken = os.Getenv("ACCESS_TOKEN")
		settings.AccountId = os.Getenv("ACCOUNT_ID")
	}
}

func TestHarvestOvertime(t *testing.T) {
	setup()

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

func TestGetUserInfo(t *testing.T) {
	setup()

	if settings.AccessToken != "" {
		userInfo, err := GetUserInfo(client, settings)

		if err != nil {
			t.Error(err)
		}

		t.Logf("\nUserID: %d\nEmail: %s\n", userInfo.ID, userInfo.Email)
	}
}

func TestAddHoursToWeek(t *testing.T) {
	entries := []models.TimeEntry{
		{
			ID:        1,
			SpentDate: "2020-01-01",
			Hours:     7.5,
		},
		{
			ID:        2,
			SpentDate: "2020-01-01",
			Hours:     7.5,
		},
	}

	date := lib.ParseDateString("2020-01-01")

	hbw := models.HoursByWeek{
		Week: "2020-01",
		HoursByDay: map[time.Time]models.HoursByDay{
			date: {
				Date: date,
				TimeEntries: map[int]models.TimeEntry{
					0: {
						ID:        0,
						SpentDate: "2020-01-01",
						Hours:     7.5,
					},
				},
				Hours: 7.5,
			},
		},
	}

	for _, e := range entries {
		hbw = AddHoursToWeek(hbw, e)
	}

	if len(hbw.HoursByDay) != 1 {
		t.Errorf("\nDays: %d.\n", len(hbw.HoursByDay))
	}

	if hbw.SumOfHours(0) != 22.5 {
		t.Errorf("\nHours: %f.\n", hbw.SumOfHours(0))
	}
}
