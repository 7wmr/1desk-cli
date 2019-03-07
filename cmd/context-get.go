package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the current 1Desk context.",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current context:", config.CurrentContext)
	},
}

func init() {
	contextCmd.AddCommand(getCmd)
}
