package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/laouji/hugotacata/internal/hugotacata/googleapi"
)

func main() {
	var credentialsFile string
	flag.StringVar(&credentialsFile, "c", "credentials.json", "path to the credentials.json file")
	flag.Parse()
	spreadsheetId := flag.Arg(0)

	if spreadsheetId == "" {
		fmt.Fprintln(os.Stderr, "SpreadsheetId must be provided as first argument")
		os.Exit(1)
	}

	client, err := googleapi.NewClient(credentialsFile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to create google api client %s\n", err)
		os.Exit(1)
	}

	readRange := "i18n!A2:D"
	resp, err := client.ReadSpreadsheet(spreadsheetId, readRange)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to retrieve data from sheet: %s\n", err)
		os.Exit(1)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		for _, row := range resp.Values {
			fmt.Printf("%v\n", row)
		}
	}
}
