package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
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

	// Sort matching entries by date
	sort.Slice(matchingEntries, func(i, j int) bool {
		date1, err1 := time.Parse("01-02-2006", matchingEntries[i].Date)
		date2, err2 := time.Parse("01-02-2006", matchingEntries[j].Date)
		if err1 != nil || err2 != nil {
			return false
		}
		return date1.Before(date2)
	})

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

	lowerValue := strings.ToLower(value)
	if entry.Date == value || strings.ToLower(entry.Location) == lowerValue || strings.ToLower(entry.Notes) == lowerValue || strings.ToLower(entry.EventType) == lowerValue {
		return true
	}

	return false
}
