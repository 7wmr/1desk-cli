package cmd

import (
	"fmt"
	"os"

	"github.com/ipsoft-tools/1desk-cli/conf"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var cfgFile string
var cfgPath string
var debugFlag bool
var versionFlag bool

var config conf.Config

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "1desk",
	Short: "A command line tool for administering 1Desk instances.",
	Long: `
	This command line tool should be used for 
	the administration of 1Desk instances. 
	`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.1desk.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&debugFlag, "debug", "D", false, "Enable debug logging")
	rootCmd.PersistentFlags().BoolVarP(&versionFlag, "version", "v", false, "Display version.")
}

// initConfig loads configuration from file.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		cfgPath = cfgFile
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		cfgPath = home + "/.1desk.yaml"
	}

	// Load config from file
	err := config.LoadConfig(cfgPath)
	if err != nil {
		os.Exit(1)
	}
}
