package repository

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"
)

// TODO the database needs to be abstracted away

type CSVRepository struct {
	csvPath string
}

// Add appends a new caffeine intake into the csv file
func (r *CSVRepository) Add(timestamp time.Time, caffeineInMg int) error {
	// Open the CSV file
	file, err := os.OpenFile(r.csvPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Append a new line to the CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	timestampStr := timestamp.Format("2006-01-02T15:04:05 MST")
	value := strconv.Itoa(caffeineInMg)
	err = writer.Write([]string{timestampStr, value}) // Open the CSV file

	if err != nil {
		return err
	}
	return nil
}

func New(path string) *CSVRepository {
	return &CSVRepository{csvPath: path}
}

type Entry struct {
	Timestamp    time.Time
	CaffeineInMg int
}

func (r *CSVRepository) Fetch() ([]Entry, error) {
	// Open the CSV file
	file, err := os.Open(r.csvPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 // Allow variable number of fields
	data, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	//parse data
	data = data[1:] //ignore the header row
	layout := "2006-01-02T15:04:05 MST"
	var entries []Entry
	for _, row := range data {
		timestamp, err := time.Parse(layout, row[0])
		if err != nil {
			return nil, err
		}
		value := parseInt(row[1])

		entry := Entry{
			Timestamp:    timestamp,
			CaffeineInMg: value,
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func parseInt(s string) int {
	// Implement error handling and conversion logic as needed
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error converting to int:", err)
		return 0 // Or handle the error differently
	}
	return i
}
