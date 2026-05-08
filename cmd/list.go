/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		tasks, err := getTasks()
		if err != nil {
			slog.Error("failed to get tasks", "err", err)
			os.Exit(1)
		}

		data := [][]string{
			{"ID", "Title", "Description", "Status", "Created at", "Updated at"},
		}

		for _, task := range tasks {
			taskInfo := []string{
				strconv.Itoa(task.ID),
				task.Title,
				task.Description,
				task.Status,
				task.CreatedAt.Format(time.RFC822),
				task.UpdatedAt.Format(time.RFC822),
			}

			data = append(data, taskInfo)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.Header(data[0])
		table.Bulk(data[1:])
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
