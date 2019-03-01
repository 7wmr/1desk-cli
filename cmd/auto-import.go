package cmd

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/spf13/cobra"
)

var fileName string
var clientName string

// ImportResponse automation import response message.
type ImportResponse struct {
	Status   string
	Messages []string
}

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import automation XML file to instance.",
	Long:  `TBC`,
	Run: func(cmd *cobra.Command, args []string) {
		fileData, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

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

		var url string = fmt.Sprintf("https://%s/mapi/IPautomata/v1/import/automata?client=%s", context.Domain, clientName)
		if debugFlag {
			fmt.Println("Debug request url:", url)
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(fileData)))
		if err != nil {
			fmt.Println("Error creating import request:", err)
			return
		}

		req.Header.Add("Authorization", "Basic "+context.Auth)
		req.Header.Set("Content-Type", "text/xml")
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error importing automation:", err)
			return
		}

		if resp.StatusCode != 200 {
			fmt.Println("Error performing API call:", resp.StatusCode)
			return
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading response:", err)
			return
		}

		var importResponse ImportResponse

		json.Unmarshal(body, &importResponse)

		if importResponse.Status != "ok" {
			fmt.Println("Error importing automation:\n\t" + strings.Join(importResponse.Messages, "\n\t"))
			return
		}
		fmt.Println("Automation imported:", importResponse.Status)
	},
}

func init() {
	autoCmd.AddCommand(importCmd)

	importCmd.Flags().StringVarP(&fileName, "name", "n", "", "Automation filename")
	importCmd.MarkFlagRequired("name")

	importCmd.Flags().StringVarP(&clientName, "client", "c", "test", "AiT instance client name")

}
