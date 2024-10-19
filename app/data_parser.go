package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
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

	start_date := queryParams.Get("start_date")
	end_date := queryParams.Get("end_date")
	location := queryParams.Get("location")
	notes := queryParams.Get("notes")
	eventType := queryParams.Get("event_type")

	if location != "" {
		nonEmptyParams = append(nonEmptyParams, location)
	}
	if notes != "" {
		nonEmptyParams = append(nonEmptyParams, notes)
	}
	if eventType != "" {
		nonEmptyParams = append(nonEmptyParams, eventType)
	}

	var parsedStartDate time.Time
	var err error
	if start_date != "" {
		parsedStartDate, err = time.Parse("01-02-2006", start_date)
		if err != nil {
			log.Fatalf("could not parse date: %v", err)
		}
	} else {
		parsedStartDate, err = time.Parse("01-02-2006", "01-01-1970")
		if err != nil {
			log.Fatalf("could not parse date: %v", err)
		}
	}

	var parsedEndDate time.Time
	if end_date != "" {
		parsedEndDate, err = time.Parse("01-02-2006", end_date)
		if err != nil {
			log.Fatalf("could not parse date: %v", err)
		}
	} else {
		parsedEndDate, err = time.Parse("01-02-2006", "01-01-2070")
		if err != nil {
			log.Fatalf("could not parse date: %v", err)
		}
	}

	fmt.Println(parsedStartDate, parsedEndDate)

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dataEntries)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	var matchingEntries []DataEntry

	for _, entry := range dataEntries {
		for _, p := range nonEmptyParams {
			if entryContains(parsedStartDate, parsedEndDate, entry, p) {
				matchingEntries = append(matchingEntries, entry)
				break
			}
		}
	}

	if len(nonEmptyParams) == 0 {
		matchingEntries = dataEntries
	}

	fmt.Println("MATCHING ENTRIES SORTED")
	fmt.Println(matchingEntries)

	return matchingEntries
}

func entryContains(startDate time.Time, endDate time.Time, entry DataEntry, value string) bool {
	entryDate, err := time.Parse("01-02-2006", entry.Date)
	if err != nil {
		log.Fatalf("could not parse entry date: %v", entry.Date)
	}

	// if !(entryDate.Equal(startDate) || entryDate.Equal(endDate) || (entryDate.After(startDate) && entryDate.Before(endDate))) {
	// 	return false
	// }

	if !entryDate.Equal(startDate) && !entryDate.Equal(endDate) && (entryDate.Before(startDate) || entryDate.After(endDate)) {
		return false
	}

	if strings.EqualFold(entry.Location, value) ||
		strings.EqualFold(entry.EventType, value) {
		return true
	}

	if strings.Contains(strings.ToLower(entry.Notes), strings.ToLower(value)) {
		return true
	}

	return false
}
