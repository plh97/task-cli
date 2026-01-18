package cmd

import (
	"github.com/liom-source/my-cli/helper"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Prints a custom message",
	Args:  cobra.MinimumNArgs(1),
	Long:  `A command that prints a custom message provided by the user.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			data, err := helper.LoadFromFile()
			if err != nil || data == nil {
				data = &helper.Config{}
			}
			var newData []helper.Task
			for i := 0; i < len(data.Data); i++ {
				if (data.Data[i].ID) != args[0] {
					newData = append(newData, data.Data[i])
				}
			}
			helper.SyncToFile(&helper.Config{Data: newData})
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
