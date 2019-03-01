package cmd

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/ipsoft-tools/1desk-cli/conf"
	"github.com/spf13/cobra"
)

var automationID int

// Automation json data.
type Automation struct {
	Name          string `json:"name"`
	ID            int    `json:"id"`
	VersionNumber int    `json:versionNumber`
}

// GetFileName get the XML file name for automation export.
func (e *Automation) GetFileName() string {
	return fmt.Sprintf(
		"%s_v%d_%d.json",
		e.Name,
		e.VersionNumber,
		e.ID)
}

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export automation JSON file from instance.",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		context, err := config.GetCurrentContext()
		if err != nil {
			fmt.Println("Error getting context:", err)
			return
		}

		fmt.Println("Current context:", context.Name)

		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		client := &http.Client{Transport: tr}

		var url string = fmt.Sprintf("https://%s/api/automaton-import-export/export/%d", context.Domain, automationID)
		if debugFlag {
			fmt.Println("Debug request url:", url)
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating export request:", err)
			return
		}

		auth := conf.Auth{Encoded: context.Auth}
		token, err := auth.GetToken(context)
		if err != nil {
			fmt.Println("Error getting authentication token:", err)
			return
		}

		if debugFlag {
			fmt.Println("Debug token response", token)
		}

		req.Header.Add(token.GetHeader())
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error exporting automation:", err)
			return
		}

		if resp.StatusCode != 200 {
			fmt.Println("Error performing API call:", resp.StatusCode)
			return
		}

		defer resp.Body.Close()
		automationJSON, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error getting automation xml:", err)
			return
		}

		var automation Automation
		json.Unmarshal(automationJSON, &automation)

		var fileName string = automation.GetFileName()

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error issue creating file:", err)
			return
		}

		fileLength, err := file.WriteString(string(automationJSON))
		if err != nil {
			fmt.Println("Error issue writing to file:", err)
			file.Close()
			return
		}

		fmt.Println("Automation export file created:", fileName)
		if debugFlag {
			fmt.Println("Debug automation export file length:", fileLength)
		}

		err = file.Close()
		if err != nil {
			fmt.Println("Error issue closing file write:", err)
			return
		}
	},
}

func init() {
	autoCmd.AddCommand(exportCmd)

	exportCmd.Flags().IntVarP(&automationID, "id", "i", 0, "Automation current ID")
	exportCmd.MarkFlagRequired("id")
}
