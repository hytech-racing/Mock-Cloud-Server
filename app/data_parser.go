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
	ID               string `json:"id"`
	MCAPFileName     string `json:"mcap_file_name"`
	MatlabFileName   string `json:"matlab_file_name"`
	AWSBucket        string `json:"aws_bucket"`
	MCAPPath         string `json:"mcap_path"`
	MatPath          string `json:"mat_path"`
	VNLatLonPath     string `json:"vn_lat_lon_path"`
	VelocityPlotPath string `json:"velocity_plot_path"`
	Date             string `json:"date"`
	Location         string `json:"location"`
	Notes            string `json:"notes,omitempty"`
	EventType        string `json:"event_type,omitempty"`
	SignedURL   	 string `json:"signed_url"`
}

func ParseJSON(file *os.File, queryParams url.Values) []DataEntry {
	var dataEntries []DataEntry

	start_date := queryParams.Get("afterDate")
	end_date := queryParams.Get("beforeDate")
	file_name := queryParams.Get("filename")
	location := queryParams.Get("location")
	notes := queryParams.Get("notes")
	eventType := queryParams.Get("eventType")

	var parsedStartDate, parsedEndDate time.Time
	var err error
	// Parse start_date
	if start_date != "" {
		parsedStartDate, err = time.Parse("01-02-2006", start_date) // MM-DD-YYYY
		if err != nil {
			log.Fatalf("could not parse start date: %v", err)
		}
	}
	// Parse end_date
	if end_date != "" {
		parsedEndDate, err = time.Parse("01-02-2006", end_date) // MM-DD-YYYY
		if err != nil {
			log.Fatalf("could not parse end date: %v", err)
		}
	} else {
		parsedEndDate, err = time.Parse("01-02-2006", "01-01-2070")
		if err != nil {
			log.Fatalf("could not parse date: %v", err)
		}
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&dataEntries)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
	}

	var matchingEntries []DataEntry

	for _, entry := range dataEntries {
		if matchesFilters(parsedStartDate, parsedEndDate, entry, file_name, location, notes, eventType) {
			matchingEntries = append(matchingEntries, entry)
		}
	}

	fmt.Println("MATCHING ENTRIES SORTED")

	return matchingEntries
}

func matchesFilters(startDate time.Time, endDate time.Time, entry DataEntry, fileName, location, notes, eventType string) bool {
	entryDate, err := time.Parse("01-02-2006", entry.Date)
	if err != nil {
		log.Fatalf("could not parse entry date: %v", entry.Date)
	}

	// Date filter logic
	if !startDate.IsZero() && !entryDate.After(startDate) {
		return false
	}
	if !endDate.Equal(time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC)) && !entryDate.Before(endDate) {
		return false
	}

	// File name filter logic
	if fileName != "" && !strings.Contains(strings.ToLower(entry.MCAPFileName), strings.ToLower(fileName)) {
		return false
	}

	// Location filter logic
	if location != "" && !strings.EqualFold(entry.Location, location) {
		return false
	}

	// Notes filter logic
	if notes != "" && !strings.Contains(strings.ToLower(entry.Notes), strings.ToLower(notes)) {
		return false
	}

	// Event type filter logic
	if eventType != "" && !strings.EqualFold(entry.EventType, eventType) {
		return false
	}

	// If all checks pass, return true
	return true
}