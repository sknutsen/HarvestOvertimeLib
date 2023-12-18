package models

type Task struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

type TaskDetails struct {
	Task    Task
	Project Project
	Client  Client
}
