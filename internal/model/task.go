package model

import (
	"errors"
	"math/rand"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func New(title, description, status string) (Task, error) {
	if err := checkStatus(status); err != nil {
		return Task{}, err
	}

	if err := isTitleEmpty(title); err != nil {
		return Task{}, err
	}

	return Task{
		ID:          rand.Int(),
		Title:       title,
		Description: description,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}, nil
}

func (t *Task) UpdateStatus(status string) error {
	if err := checkStatus(status); err != nil {
		return err
	}

	t.Status = status
	t.UpdatedAt = time.Now()

	return nil
}

// Helper function that checks 'status' input
func checkStatus(status string) error {
	if status != "DONE" && status != "IN_PROGRESS" && status != "TODO" {
		return errors.New("Invalid task status")
	}

	return nil
}

// Helper function that verifies task title is not empty
func isTitleEmpty(title string) error {
	if title == "" {
		return errors.New("task title is empty - please provide a title to the task")
	}

	return nil
}

// Helper function that sanitizes user 'title' input
func sanitizeString(title string) string {
	return strings.TrimSpace(title)
}
