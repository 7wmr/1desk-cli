package cmd

import (
	"github.com/ipsoft-tools/1desk-cli/conf"

	"github.com/spf13/cobra"
)

var addName string
var addUsername string
var addPassword string
var addDomain string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a AiT instance context to local config",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		auth := conf.Auth{Username: addUsername, Password: addPassword}
		context := conf.Context{Auth: auth.Encode(), Domain: addDomain, Name: addName}
		config.AddContext(context)
		config.WriteConfig(cfgPath)
	},
}

func init() {
	contextCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addName, "name", "n", "", "Instance name")
	addCmd.MarkFlagRequired("name")

	addCmd.Flags().StringVarP(&addUsername, "username", "u", "", "Instance username")
	addCmd.MarkFlagRequired("username")

	addCmd.Flags().StringVarP(&addPassword, "password", "p", "", "Instance password")
	addCmd.MarkFlagRequired("password")

	addCmd.Flags().StringVarP(&addDomain, "domain", "d", "", "Instance domain")
	addCmd.MarkFlagRequired("domain")
}
