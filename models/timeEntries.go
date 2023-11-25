package models

type TimeEntries struct {
	TimeEntries []TimeEntry `json:"time_entries"`
	Links       Links       `json:"links"`
}

type Links struct {
	Next string `json:"next"`
}
