package cmd

import (
	"encoding/json"
	"os"

	"github.com/matthieukhl/task-tracker/internal/model"
)

const DATA_FILE = "./data/tasks.json"

func getTasks() ([]model.Task, error) {
	rfile, err := os.OpenFile(DATA_FILE, os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}

	tasks := []model.Task{}

	err = json.NewDecoder(rfile).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
