/*
Copyright © 2026 MATTHIEU KHAIRALLAH
*/
package cmd

import (
	"encoding/json"
	"log/slog"
	"os"

	"github.com/matthieukhl/task-tracker/internal/model"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		title, err := cmd.Flags().GetString("title")
		if err != nil {
			slog.Error("error parsing title", "err", err)
			os.Exit(1)
		}

		description, err := cmd.Flags().GetString("description")
		if err != nil {
			slog.Error("error parsing description", "err", err)
			os.Exit(1)
		}

		status, err := cmd.Flags().GetString("status")
		if err != nil {
			slog.Error("error parsing status", "err", err)
			os.Exit(1)
		}

		newTask, err := model.New(title, description, status)
		if err != nil {
			slog.Error("error creating new task", "err", err)
			os.Exit(1)
		}

		file, err := os.OpenFile("./data/tasks.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			slog.Error("error reading or creating tasks.json file", "err", err)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		tasks := []model.Task{}
		decoder.Decode(tasks)

		tasks = append(tasks, newTask)

		encoder := json.NewEncoder(file)
		err = encoder.Encode(tasks)
		if err != nil {
			slog.Error("error writing to tasks.json", "err", err)
			os.Exit(1)
		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	addCmd.PersistentFlags().String("title", "", "The task title")
	addCmd.PersistentFlags().String("description", "", "The task description")
	addCmd.PersistentFlags().String("status", "TODO", "The task status")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
