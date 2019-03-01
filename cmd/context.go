package cmd

import (
	"github.com/spf13/cobra"
)

// contextCmd represents the context command
var contextCmd = &cobra.Command{
	Use:   "context",
	Short: "A brief description of your command",
	Long:  `TBC`,
}

func init() {
	rootCmd.AddCommand(contextCmd)
}
