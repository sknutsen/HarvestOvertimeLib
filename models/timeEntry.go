package models

type TimeEntry struct {
	Id           uint64  `json:"id"`
	SpentDate    string  `json:"spent_date"`
	Project      Project `json:"project"`
	Task         Task    `json:"task"`
	Hours        float64 `json:"hours"`
	RoundedHours float64 `json:"rounded_hours"`
}
