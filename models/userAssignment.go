package models

import "time"

type UserAssignment struct {
	ID               int         `json:"id"`
	IsProjectManager bool        `json:"is_project_manager"`
	IsActive         bool        `json:"is_active"`
	Budget           interface{} `json:"budget"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	HourlyRate       float64     `json:"hourly_rate"`
}
