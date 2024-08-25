package repository

import (
	"encoding/csv"
	"log"
	"os"
)

type CSVRepository struct {
	csvPath string
}

func (r *CSVRepository) Add(caffeineInMg int) {
	//TODO implement me
	panic("implement me")
}

func New(path string) *CSVRepository {
	return &CSVRepository{csvPath: path}

}

func (r *CSVRepository) Fetch() int {

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
		panic(err)
	}
	log.Println(data)

	return 0
}
