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

type MCAPFileType struct {
	FileName 		 string `json:"file_name"`
	SignedURL 		 string `json:"signed_url"`
}
type MATFileType struct {
	FileName 		 string `json:"file_name"`
	SignedURL 		 string `json:"signed_url"`
}

type ContentFileType struct {
	FileName 		 string `json:"file_name"`
	SignedURL 		 string `json:"signed_url"`
}

type SchemaType struct {
	FileName 		 string `json:"schema"`
}

type DataEntryNew struct {
	ID               string `json:"id"`
	MCAPFiles     	 []MCAPFileType `json:"mcap_files"`
	MATFiles   		 []MATFileType `json:"mat_files"`
	ContentFiles   	 []ContentFileType `json:"content_files"`
	Date             string `json:"date"`
	Location         string `json:"location"`
	Notes            string `json:"notes,omitempty"`
	EventType        string `json:"event_type,omitempty"`
	Schema			 map[string]string `json:"schema"`
}

func ParseJSONNew(file *os.File, queryParams url.Values) []DataEntryNew {
	var dataEntries []DataEntryNew

	startDate := queryParams.Get("afterDate")
	endDate := queryParams.Get("beforeDate")
	fileName := queryParams.Get("filename")
	location := queryParams.Get("location")
	notes := queryParams.Get("notes")
	eventType := queryParams.Get("eventType")

	var parsedStartDate, parsedEndDate time.Time
	var err error
	if startDate != "" {
		parsedStartDate, err = time.Parse("01-02-2006", startDate)
		if err != nil {
			log.Fatalf("could not parse start date: %v", err)
		}
	}
	if endDate != "" {
		parsedEndDate, err = time.Parse("01-02-2006", endDate)
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

	var matchingEntries []DataEntryNew

	for _, entry := range dataEntries {
		if matchesFiltersNew(entry, parsedStartDate, parsedEndDate, fileName, location, notes, eventType) {
			matchingEntries = append(matchingEntries, entry)
		}
	}

	return matchingEntries;
}

func matchesFiltersNew(entry DataEntryNew, startDate, endDate time.Time, fileName, location, notes, eventType string) bool {
	entryDate, err := time.Parse("01-02-2006", entry.Date)
	if err != nil {
		fmt.Printf("Error parsing entry date: %v\n", err)
		return false
	}
	if (startDate.IsZero() || entryDate.After(startDate) || entryDate.Equal(startDate)) &&
		(endDate.IsZero() || entryDate.Before(endDate) || entryDate.Equal(endDate)) {
		if fileName != "" {
			fileMatched := false
			for _, mcapFile := range entry.MCAPFiles {
				fmt.Println(mcapFile.FileName)
				if strings.Contains(mcapFile.FileName, fileName) {
					fileMatched = true
					break
				}
			}
			for _, matFile := range entry.MATFiles {
				fmt.Println(matFile.FileName)
				if strings.Contains(matFile.FileName, fileName) {
					fileMatched = true
					break
				}
			}
			if !fileMatched {
				return false
			}
		}

		if location != "" && !strings.Contains(entry.Location, location) {
			return false
		}

		if notes != "" && !strings.Contains(entry.Notes, notes) {
			return false
		}

		if eventType != "" && !strings.Contains(entry.EventType, eventType) {
			return false
		}

		return true
	}
	return false
}