package cmd

import (
	"time"

	"github.com/liom-source/my-cli/helper"
	"github.com/spf13/cobra"
)

var updateContentCmd = &cobra.Command{
	Use:   "update",
	Short: "Prints a custom message",
	Args:  cobra.MinimumNArgs(1),
	Long:  `A command that prints a custom message provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			data, err := helper.LoadFromFile()
			if err != nil || data == nil {
				data = &helper.Config{}
			}
			for i := 0; i < len(data.Data); i++ {
				if (data.Data[i].ID) == args[0] {
					data.Data[i].Description = args[1]
					data.Data[i].UpdatedAt = time.Now().Format(time.RFC3339)
				}
			}

			helper.SyncToFile(data)
		} else {
			println("No message provided.")
		}
	},
}

var updateInProgressCmd = &cobra.Command{
	Use:   "mark-in-progress",
	Short: "Marks a task as in progress",
	Args:  cobra.MinimumNArgs(1),
	Long:  `A command that marks a task as in progress based on the provided ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			data, err := helper.LoadFromFile()
			if err != nil || data == nil {
				data = &helper.Config{}
			}
			for i := 0; i < len(data.Data); i++ {
				if (data.Data[i].ID) == args[0] {
					data.Data[i].Status = helper.StatusInProgress
					data.Data[i].UpdatedAt = time.Now().Format(time.RFC3339)
				}
			}

			helper.SyncToFile(data)
		} else {
			println("No message provided.")
		}
	},
}

var updateDoneCmd = &cobra.Command{
	Use:   "mark-done",
	Short: "Marks a task as done",
	Args:  cobra.MinimumNArgs(1),
	Long:  `A command that marks a task as done based on the provided ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			data, err := helper.LoadFromFile()
			if err != nil || data == nil {
				data = &helper.Config{}
			}
			for i := 0; i < len(data.Data); i++ {
				if (data.Data[i].ID) == args[0] {
					data.Data[i].Status = helper.StatusDone
					data.Data[i].UpdatedAt = time.Now().Format(time.RFC3339)
				}
			}

			helper.SyncToFile(data)
		} else {
			println("No message provided.")
		}
	},
}

func init() {
	rootCmd.AddCommand(updateContentCmd)
	rootCmd.AddCommand(updateInProgressCmd)
	rootCmd.AddCommand(updateDoneCmd)
}