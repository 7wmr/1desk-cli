package cmd

import (
	"github.com/spf13/cobra"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "Commands to administrer 1Desk instance contexts.",
	Long:  `TBC`,
}

func init() {
	rootCmd.AddCommand(contextCmd)
}
