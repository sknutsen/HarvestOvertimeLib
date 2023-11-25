package models

type Settings struct {
	AccessToken     string  `json:"accessToken"`
	AccountId       string  `json:"accountId"`
	CarryOverTime   float64 `json:"carryOverTime"`
	OnlyCurrentYear bool    `json:"OnlyCurrentYear"`
	TimeOffTasks    []Task  `json:"timeOffTasks"`
}
