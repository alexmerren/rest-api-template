package task

import (
	uuid "github.com/satori/go.uuid"
)

// Task struct is used to encapsulate functions
type Task struct {
	ID      uuid.UUID
	User    string
	Content string `json:"content"`
}

// NewTask function is called to return a new task with specified information
func NewTask() *Task {
	id := uuid.NewV4()
	return &Task{
		ID: id,
	}
}
