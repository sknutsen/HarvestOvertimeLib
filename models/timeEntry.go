package models

import "time"

type TimeEntry struct {
	ID                int            `json:"id"`
	SpentDate         string         `json:"spent_date"`
	User              User           `json:"user"`
	Client            Client         `json:"client"`
	Project           Project        `json:"project"`
	Task              Task           `json:"task"`
	UserAssignment    UserAssignment `json:"user_assignment"`
	TaskAssignment    TaskAssignment `json:"task_assignment"`
	Hours             float64        `json:"hours"`
	HoursWithoutTimer float64        `json:"hours_without_timer"`
	RoundedHours      float64        `json:"rounded_hours"`
	Notes             string         `json:"notes"`
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	IsLocked          bool           `json:"is_locked"`
	LockedReason      string         `json:"locked_reason"`
	IsClosed          bool           `json:"is_closed"`
	IsBilled          bool           `json:"is_billed"`
	TimerStartedAt    interface{}    `json:"timer_started_at"`
	StartedTime       string         `json:"started_time"`
	EndedTime         string         `json:"ended_time"`
	IsRunning         bool           `json:"is_running"`
	Invoice           interface{}    `json:"invoice"`
	ExternalReference interface{}    `json:"external_reference"`
	Billable          bool           `json:"billable"`
	Budgeted          bool           `json:"budgeted"`
	BillableRate      float64        `json:"billable_rate"`
	CostRate          float64        `json:"cost_rate"`
}
