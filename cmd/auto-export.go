package cmd

import (
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var automationID int

// Automation xml data.
type Automation struct {
	XMLName xml.Name `xml:"automaton"`
	Name    string   `xml:"name"`
	ID      string   `xml:"automatonID"`
}

// AutomationExport xml data.
type AutomationExport struct {
	XMLName    xml.Name   `xml:"automatonExport"`
	Automation Automation `xml:"automaton"`
}

// GetFileName get the XML file name for automation export.
func (e *AutomationExport) GetFileName() string {
	return fmt.Sprintf(
		"%s_%s.xml",
		e.Automation.Name,
		e.Automation.ID)
}

// exportCmd represents the export command
var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export automation XML file from instance.",
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

		var url string = fmt.Sprintf("https://%s/mapi/IPautomata/v1/export/automata/%d", context.Domain, automationID)
		if debugFlag {
			fmt.Println("Debug request url:", url)
		}

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating export request:", err)
			return
		}

		req.Header.Add("Authorization", "Basic "+context.Auth)
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
		automationXML, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error getting automation xml:", err)
			return
		}

		var automationExport AutomationExport
		xml.Unmarshal(automationXML, &automationExport)

		var fileName string = automationExport.GetFileName()

		file, err := os.Create(fileName)
		if err != nil {
			fmt.Println("Error issue creating file:", err)
			return
		}

		fileLength, err := file.WriteString(string(automationXML))
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
