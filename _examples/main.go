package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fanchann/isly"
)

func main() {
	start := time.Now()

	readRandom1Data()
	// readRandom2Data()
	// readDateFormats()
	// readSimpleData()

	fmt.Printf("Execution time: %v\n", time.Since(start))
}

// =======================================
// Struct Definitions
// =======================================

type Random struct {
	ID           int                    `isly:"id"`
	Name         string                 `isly:"name"`
	Age          int                    `isly:"age"`
	Salary       float64                `isly:"salary"`
	IsActive     bool                   `isly:"is_active"`
	Scores       []int                  `isly:"scores, list"`
	Tags         []string               `isly:"tags, list"`
	Metadata     map[string]interface{} `isly:"metadata, json"`
	CreatedAt    time.Time              `isly:"created_at, 2006-01-02"`
	NullValue    string                 `isly:"null_value"`
	SpecialChars string                 `isly:"special_chars"`
	HexValue     []byte                 `isly:"hex_value, hex"`
	BinaryData   []byte                 `isly:"binary_data, binary"`
}

type Random2 struct {
	Name      string    `isly:"Name"`
	Latitude  string    `isly:"Latitude"`
	Longitude string    `isly:"Longitude"`
	Age       int       `isly:"Age"`
	Date      time.Time `isly:"Date, 2/1/2006"`
	Height    float64   `isly:"Height"`
	Salary    int       `isly:"Salary"`
}

type DateFormat struct {
	Name         string    `isly:"Name"`
	Date         time.Time `isly:"Date, 2006/01/02"`
	DateFormat1  time.Time `isly:"DateFormat_1, 15-05-2026"`
	DateFormat2  time.Time `isly:"DateFormat_2, 02/01/2006"`
	DateFormat3  time.Time `isly:"DateFormat_3, 01-02-2006"`
	DateFormat4  time.Time `isly:"DateFormat_4, 01/02/2006"`
	DateFormat5  time.Time `isly:"DateFormat_5, 2006-01-02 15:04:05"`
	DateFormat6  time.Time `isly:"DateFormat_6, 2006/01/02 15:04:05"`
	DateFormat7  time.Time `isly:"DateFormat_7, 2006-01-02T15:04:05Z"`
	DateFormat8  time.Time `isly:"DateFormat_8, 2006-01-02T15:04:05Z07:00"`
	DateFormat9  time.Time `isly:"DateFormat_9, 2006-01-02T15:04:05"`
	DateFormat10 time.Time `isly:"DateFormat_10, 2006-01-02 15:04"`
	DateFormat11 time.Time `isly:"DateFormat_11, 2006/01/02 15:04"`
	DateFormat12 time.Time `isly:"DateFormat_12, 2006/01/02 15:04"`
}

type Raw struct {
	ID        int       `isly:"id"`
	FirstName string    `isly:"first_name"`
	LastName  string    `isly:"last_name"`
	Gender    string    `isly:"gender"`
	BirthDay  time.Time `isly:"birth_day, 2006-01-02"`
}

// =======================================
// CSV Reading & Display Functions
// =======================================

func readSimpleData() {
	var records []Raw
	islyComp := isly.NewIsly()

	islyComp.ReadFile("./examples/data/simple.csv")
	if err := islyComp.UnmarshalCSV(&records); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, r := range records {
		fmt.Printf("ID: %d | Name: %s %s | Gender: %s | BirthDay: %s\n",
			r.ID, r.FirstName, r.LastName, r.Gender, r.BirthDay.Format("2006-01-02"))
	}
}

func readRandom1Data() {
	var list []Random
	islyComp := isly.NewIsly()

	islyComp.ReadFile("./examples/data/random1.csv")
	if err := islyComp.UnmarshalCSV(&list); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, e := range list {
		fmt.Println("------ Random Data ------")
		fmt.Printf("ID: %d | Name: %s | Age: %d | Salary: %.2f | Active: %v\n",
			e.ID, e.Name, e.Age, e.Salary, e.IsActive)
		fmt.Printf("Scores: %v | Tags: %v\n", e.Scores, e.Tags)
		fmt.Printf("Metadata: %v\n", e.Metadata)
		fmt.Printf("CreatedAt: %s\n", e.CreatedAt)
		fmt.Printf("NullValue: %q | SpecialChars: %q\n", e.NullValue, e.SpecialChars)
		fmt.Printf("HexValue: %v\n", e.HexValue)
		fmt.Printf("BinaryData: %s\n", formatBinary(e.BinaryData))
	}
}

func readRandom2Data() {
	var list []Random2
	islyComp := isly.NewIsly()

	islyComp.ReadFile("./examples/data/random2.csv")
	if err := islyComp.UnmarshalCSV(&list); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, r := range list {
		fmt.Println("------ Random2 Data ------")
		fmt.Printf("Name: %s | Age: %d | Date: %s\n",
			r.Name, r.Age, r.Date.Format("02-Jan-2006"))
		fmt.Printf("Lat/Lng: %s, %s | Height: %.2f | Salary: %d\n",
			r.Latitude, r.Longitude, r.Height, r.Salary)
	}
}

func readDateFormats() {
	var list []DateFormat
	islyComp := isly.NewIsly()

	islyComp.ReadFile("./examples/data/date_formats.csv")
	if err := islyComp.UnmarshalCSV(&list); err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, d := range list {
		fmt.Println("------ Date Formats ------")
		fmt.Printf("Name: %s\n", d.Name)
		fmt.Printf("Default: %v | Format1: %v | Format2: %v\n",
			d.Date, d.DateFormat1, d.DateFormat2)
		fmt.Printf("Format3-12: %v | %v | %v | %v | %v | %v | %v | %v | %v | %v\n",
			d.DateFormat3, d.DateFormat4, d.DateFormat5, d.DateFormat6,
			d.DateFormat7, d.DateFormat8, d.DateFormat9, d.DateFormat10,
			d.DateFormat11, d.DateFormat12)
	}
}

// =======================================
// Helper
// =======================================

func formatBinary(data []byte) string {
	var result []string
	for _, b := range data {
		result = append(result, fmt.Sprintf("%08b", b))
	}
	return strings.Join(result, " ")
}
