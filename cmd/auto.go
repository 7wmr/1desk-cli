package cmd

import (
	"github.com/spf13/cobra"
)

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "A brief description of your command",
	Long:  `TBD`,
}

func init() {
	rootCmd.AddCommand(autoCmd)
}
