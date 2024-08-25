package repository

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type CSVRepository struct {
	csvPath string
}

func (r *CSVRepository) Add(caffeineInMg int) error {
	// Open the CSV file
	file, err := os.OpenFile(r.csvPath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Append a new line to the CSV file
	writer := csv.NewWriter(file)
	defer writer.Flush()

	timestamp := time.Now().String()
	value := strconv.Itoa(caffeineInMg)

	err = writer.Write([]string{timestamp, value}) // Open the CSV file

	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func New(path string) *CSVRepository {
	return &CSVRepository{csvPath: path}
}

func (r *CSVRepository) Fetch() ([][]string, error) {

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
	return data, nil
}
