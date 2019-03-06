package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var contextName string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the current AiT instance",
	Long:  `XXXXX`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.Validate(contextName) {
			config.CurrentContext = contextName
			err := config.WriteConfig(cfgPath)
			if err != nil {
				return
			}
			fmt.Println("Context set:", contextName)
		} else {
			fmt.Println("Context invalid:", contextName)
		}
	},
}

func init() {
	contextCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&contextName, "name", "n", "", "Context name")
	setCmd.MarkFlagRequired("name")

}
