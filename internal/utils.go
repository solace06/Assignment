package internal

import (
	"errors"

	"github.com/google/uuid"
)

const (
	StatusTodo       = "todo"
	StatusInProgress = "in_progress"
	StatusDone       = "done"
)

func ValidateTask(task CreateTaskRequest) error {
	if task.Title == "" {
		return errors.New("title is required")
	}

	if len(task.Title) > 200 {
		return errors.New("title cannot be more than 200 characters")
	}

	if task.Status != "" &&
		task.Status != StatusTodo &&
		task.Status != StatusInProgress &&
		task.Status != StatusDone {
		return errors.New("invalid status")
	}

	return nil
}

func generateUUID() string {
	return uuid.New().String()
}
