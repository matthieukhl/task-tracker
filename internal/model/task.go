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
	if status != "done" && status != "in_progress" && status != "todo" {
		return errors.New("Invalid status")
	}

	return nil
}

// Helper function that sanitizes user 'title' input
func sanitizeString(title string) string {
	return strings.TrimSpace(title)
}
