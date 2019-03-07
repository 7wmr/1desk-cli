package cmd

import (
	"github.com/spf13/cobra"
)

// autoCmd represents the auto command
var autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "1Desk IPautomata automation commands",
	Long:  `TBD`,
}

func init() {
	rootCmd.AddCommand(autoCmd)
}
