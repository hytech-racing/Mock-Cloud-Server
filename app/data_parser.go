package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
)

type DataEntry struct {
	Date      string `json:"date"`
	Location  string `json:"location"`
	Notes     string `json:"notes"`
	EventType string `json:"event type"`
}

func ParseJSON(file *os.File, queryParams url.Values) {
	var nonEmptyParams []string
	var dataEntries []DataEntry

	date := queryParams.Get("date")
	location := queryParams.Get("location")
	notes := queryParams.Get("notes")
	eventType := queryParams.Get("eventType")

	if date != "" {
		nonEmptyParams = append(nonEmptyParams, date)
	}
	if location != "" {
		nonEmptyParams = append(nonEmptyParams, location)
	}
	if notes != "" {
		nonEmptyParams = append(nonEmptyParams, notes)
	}
	if eventType != "" {
		nonEmptyParams = append(nonEmptyParams, eventType)
	}

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&dataEntries)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, entry := range dataEntries {
		for _, p := range nonEmptyParams {
			// fmt.Printf("Checking param %s\n", p)
			if entryContains(entry, p) {
				fmt.Println(entry)
			}
		}
	}
}

func entryContains(entry DataEntry, value string) bool {
	if entry.Date == value || entry.Location == value || entry.Notes == value || entry.EventType == value {
		return true
	}
	return false
}
