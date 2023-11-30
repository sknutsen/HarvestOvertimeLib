package models

import "time"

type TaskAssignment struct {
	ID         int         `json:"id"`
	Billable   bool        `json:"billable"`
	IsActive   bool        `json:"is_active"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	HourlyRate float64     `json:"hourly_rate"`
	Budget     interface{} `json:"budget"`
}
