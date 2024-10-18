package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
	"time"
)

type DataEntry struct {
	Date      string `json:"date"`
	Location  string `json:"location"`
	Notes     string `json:"notes"`
	EventType string `json:"event_type"`
	Bucket    string `json:"aws_bucket"`
	Path      string `json:"mcap_path"`
	FileName  string `json:"mcap_file_name"`
	SignedURL string `json:"signed_url"`
}

func ParseJSON(file *os.File, queryParams url.Values) []DataEntry {
	var nonEmptyParams []string
	var dataEntries []DataEntry

	date := queryParams.Get("date")
	location := queryParams.Get("location")
	notes := queryParams.Get("notes")
	eventType := queryParams.Get("event_type")

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
	}

	var matchingEntries []DataEntry

	for _, entry := range dataEntries {
		for _, p := range nonEmptyParams {
			if entryContains(entry, p) {
				matchingEntries = append(matchingEntries, entry)
				break
			}
		}
	}

	if len(matchingEntries) == 0 {
		matchingEntries = dataEntries
	}

	// fmt.Println("DATA ENTRIES")
	// fmt.Println(dataEntries)

	// Sort matching entries by date
	sort.Slice(matchingEntries, func(i, j int) bool {
		date1, err1 := time.Parse("01-02-2006", matchingEntries[i].Date)
		date2, err2 := time.Parse("01-02-2006", matchingEntries[j].Date)
		if err1 != nil || err2 != nil {
			// If there's an error in parsing, don't sort by date
			return false
		}
		return date1.Before(date2)
	})

	return matchingEntries
}

func entryContains(entry DataEntry, value string) bool {
	if entry.Date == value || entry.Location == value || entry.Notes == value || entry.EventType == value {
		return true
	}
	return false
}
