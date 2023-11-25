package models

type Tasks struct {
	Tasks []Task `json:"tasks"`
	Links Links  `json:"links"`
}
