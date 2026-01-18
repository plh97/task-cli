package cmd

import (
	"fmt"

	"github.com/liom-source/my-cli/helper"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list messages",
	// Args:  cobra.MinimumNArgs(1),
	Long: `A command that prints a custom message provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		data, err := helper.LoadFromFile()
		if err != nil || data == nil {
			data = &helper.Config{}
		}
		if len(args) > 0 {
			for i := 0; i < len(data.Data); i++ {
				if args[0] == "done" && data.Data[i].Status == helper.StatusDone {
					fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
						data.Data[i].ID,
						data.Data[i].Description,
						data.Data[i].Status,
						data.Data[i].CreatedAt,
						data.Data[i].UpdatedAt,
					)
				}
				if args[0] == "in-progress" && data.Data[i].Status == helper.StatusInProgress {
					fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
						data.Data[i].ID,
						data.Data[i].Description,
						data.Data[i].Status,
						data.Data[i].CreatedAt,
						data.Data[i].UpdatedAt,
					)
				}
				if args[0] == "todo" && data.Data[i].Status == helper.StatusTodo {
					fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
						data.Data[i].ID,
						data.Data[i].Description,
						data.Data[i].Status,
						data.Data[i].CreatedAt,
						data.Data[i].UpdatedAt,
					)
				}
			}
		} else {
			for i := 0; i < len(data.Data); i++ {
				fmt.Printf("ID: %s\nDescription: %s\nStatus: %s\nCreatedAt: %s\nUpdatedAt: %s\n\n",
					data.Data[i].ID,
					data.Data[i].Description,
					data.Data[i].Status,
					data.Data[i].CreatedAt,
					data.Data[i].UpdatedAt,
				)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
