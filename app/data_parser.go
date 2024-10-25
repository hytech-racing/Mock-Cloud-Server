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
	var nonEmptyParams []string
	var dataEntries []DataEntry

	start_date := queryParams.Get("afterDate")
	end_date := queryParams.Get("beforeDate")
	file_name := queryParams.Get("fileName")
	location := queryParams.Get("location")
	notes := queryParams.Get("notes")
	eventType := queryParams.Get("eventType")

	if location != "" {
		nonEmptyParams = append(nonEmptyParams, location)
	}
	if notes != "" {
		nonEmptyParams = append(nonEmptyParams, notes)
	}
	if eventType != "" {
		nonEmptyParams = append(nonEmptyParams, eventType)
	}
	if file_name != "" {
		nonEmptyParams = append(nonEmptyParams, file_name)
	}


	var parsedStartDate, parsedEndDate time.Time
	var err error
	// Parse start_date
	if start_date != "" {
		parsedStartDate, err = time.Parse("01-02-2006", start_date) // MM-DD-YYYY
		if err != nil {
			log.Fatalf("could not parse start date: %v", err)
		}
		nonEmptyParams = append(nonEmptyParams, start_date)
	}
	// Parse end_date
	if end_date != "" {
		parsedEndDate, err = time.Parse("01-02-2006", end_date) // MM-DD-YYYY
		if err != nil {
			log.Fatalf("could not parse end date: %v", err)
		}
		nonEmptyParams = append(nonEmptyParams, end_date)
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

	// fmt.Println("Parsed Start Date:", parsedStartDate)
	// fmt.Println("Parsed End Date:", parsedEndDate)
	
	for _, entry := range dataEntries {
		for _, p := range nonEmptyParams {
			// fmt.Println("Entry", entry)
			if entryContains(parsedStartDate, parsedEndDate, entry, p) {
				matchingEntries = append(matchingEntries, entry)
				break
			}
		}
	}

	if len(nonEmptyParams) == 0 {
		matchingEntries = dataEntries
	}

	// sort.Slice(matchingEntries, func(i, j int) bool {
	// 	date1, err1 := time.Parse("01-02-2006", matchingEntries[i].Date)
	// 	date2, err2 := time.Parse("01-02-2006", matchingEntries[j].Date)
	// 	if err1 != nil || err2 != nil {
	// 		return false
	// 	}
	// 	return date1.Before(date2)
	// })

	fmt.Println("MATCHING ENTRIES SORTED")
	// fmt.Println(matchingEntries)

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

	// fmt.Println("Start Date:", startDate)
	// fmt.Println("End Date:", endDate)
	// fmt.Println("Entry Name:", entry.MCAPFileName)
	// fmt.Println("Value:", value)

	if !(startDate == time.Time{}) || !(endDate.Equal(time.Date(2070, 1, 1, 0, 0, 0, 0, time.UTC))) {
		if entryDate.After(startDate) && entryDate.Before(endDate) {
			return true
		}
	}

	if strings.EqualFold(entry.Location, value) || strings.EqualFold(entry.EventType, value) {
		return true
	}
	
	if strings.Contains(strings.ToLower(entry.MCAPFileName), strings.ToLower(value)) {
		return true
	}

	if strings.Contains(strings.ToLower(entry.Notes), strings.ToLower(value)) {
		return true
	}

	return false
}
