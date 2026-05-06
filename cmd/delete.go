/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"log/slog"
	"os"
	"slices"

	"github.com/matthieukhl/task-tracker/internal/model"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if 'id' flag is == 0
		taskID, err := cmd.Flags().GetInt("id")
		if err != nil {
			slog.Error("error parsing task ID", "err", err)
			os.Exit(1)
		}
		rfile, err := os.OpenFile("./data/tasks.json", os.O_RDONLY, 0644)
		if err != nil {
			slog.Error("failed to read tasks.json file", "err", err)
			os.Exit(1)
		}
		defer rfile.Close()

		tasks := []model.Task{}
		err = json.NewDecoder(rfile).Decode(&tasks)
		if err != nil {
			slog.Error("failed to decode JSON data from file", "err", err)
			os.Exit(1)
		}

		for i, task := range tasks {
			if taskID == task.ID {
				tasks = slices.Delete(tasks, i, i+1)
				break
			}
		}

		wfile, err := os.OpenFile(DATA_FILE, os.O_RDWR|os.O_TRUNC, 0644)
		if err != nil {
			slog.Error("error opening file", "file", DATA_FILE, "err", err)
			os.Exit(1)
		}
		defer wfile.Close()

		err = json.NewEncoder(wfile).Encode(tasks)
		if err != nil {
			slog.Error("failed to write tasks to file", "file", DATA_FILE, "err", err)
			os.Exit(1)
		}

		slog.Info("deleted task successfully", "ID", taskID)

	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	deleteCmd.PersistentFlags().Int("id", 0, "Task ID")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
