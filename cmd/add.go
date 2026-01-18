package cmd

import (
	"crypto/rand"
	"time"

	"github.com/plh97/task-cli/helper"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Prints a custom message",
	Args:  cobra.MinimumNArgs(1),
	Long:  `A command that prints a custom message provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			data, err := helper.LoadFromFile()
			if err != nil || data == nil {
				data = &helper.Config{}
			}
			data.Data = append(data.Data, helper.Task{
				// generate random string id
				ID:          rand.Text(),
				Description: args[0],
				Status:      helper.StatusTodo,
				CreatedAt:   time.Now().Format(time.RFC3339),
				UpdatedAt:   time.Now().Format(time.RFC3339),
			})
			helper.SyncToFile(data)
		} else {
			println("No message provided.")
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
