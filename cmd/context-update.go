package cmd

import (
	"fmt"

	"github.com/ipsoft-tools/1desk-cli/conf"

	"github.com/spf13/cobra"
)

var updateName string
var updateUsername string
var updateDomain string

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		if config.Validate(updateName) {

			auth := conf.Auth{Username: updateUsername}
			auth.PromptPassword()
			config.UpdateContext(updateName, auth, updateDomain)

			err := config.WriteConfig(cfgPath)
			if err != nil {
				return
			}
			fmt.Println("Context updated:", updateName)
		} else {
			fmt.Println("Context name not valid:", updateName)
			return
		}
	},
}

func init() {
	contextCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringVarP(&updateName, "name", "n", "", "Context name")
	updateCmd.MarkFlagRequired("name")

	updateCmd.Flags().StringVarP(&updateUsername, "username", "u", "", "Context username")
	updateCmd.MarkFlagRequired("username")

	updateCmd.Flags().StringVarP(&updateDomain, "domain", "d", "", "Context domain")

}
