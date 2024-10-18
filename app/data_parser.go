package app

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"sort"
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

	sortByDate(matchingEntries)

	fmt.Println("MATCHING ENTRIES SORTED")
	fmt.Println(matchingEntries)

	return matchingEntries
}

func entryContains(entry DataEntry, value string) bool {
	if strings.EqualFold(entry.Date, value) || strings.EqualFold(entry.Location, value) ||
		strings.EqualFold(entry.EventType, value) {
		return true
	}

	if strings.Contains(strings.ToLower(entry.Notes), strings.ToLower(value)) {
		return true
	}

	return false
}

type ByDate []DataEntry

func (a ByDate) Len() int      { return len(a) }
func (a ByDate) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool {
	layout := "01-02-2006" // MM-DD-YYYY format

	// Try to parse dates
	fmt.Println(a[i].Date)
	fmt.Println(a[j].Date)
	dateI, errI := time.Parse(layout, a[i].Date)
	dateJ, errJ := time.Parse(layout, a[j].Date)

	if errI != nil {
		fmt.Printf("Error parsing date for entry %v: %v\n", a[i], errI)
	}
	if errJ != nil {
		fmt.Printf("Error parsing date for entry %v: %v\n", a[j], errJ)
	}

	// Handle parsing errors (keep original order if there's an error)
	if errI != nil || errJ != nil {
		return false
	}

	return dateI.Before(dateJ)
}

func sortByDate(entries []DataEntry) {
	sort.Sort(ByDate(entries))
}
