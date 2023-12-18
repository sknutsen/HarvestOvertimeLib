package models

import "time"

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type UserInfo struct {
	ID                           int       `json:"id"`
	FirstName                    string    `json:"first_name"`
	LastName                     string    `json:"last_name"`
	Email                        string    `json:"email"`
	Telephone                    string    `json:"telephone"`
	Timezone                     string    `json:"timezone"`
	HasAccessToAllFutureProjects bool      `json:"has_access_to_all_future_projects"`
	IsContractor                 bool      `json:"is_contractor"`
	IsActive                     bool      `json:"is_active"`
	CreatedAt                    time.Time `json:"created_at"`
	UpdatedAt                    time.Time `json:"updated_at"`
	WeeklyCapacity               int       `json:"weekly_capacity"`
	DefaultHourlyRate            float64   `json:"default_hourly_rate"`
	CostRate                     float64   `json:"cost_rate"`
	Roles                        []string  `json:"roles"`
	AccessRoles                  []string  `json:"access_roles"`
	AvatarURL                    string    `json:"avatar_url"`
}
