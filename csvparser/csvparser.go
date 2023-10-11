package csvparser

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadCSV() [][]string {
		file, err := os.Open("housesInput.csv")
		if err != nil {
			log.Fatal("Error raised while reading the file", err)
    }
		// Defer keyword allows close call to be declared next to open call, but delays execution to end of function
		defer file.Close()
		// Read the file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
    if err != nil {
        log.Fatal("Error raised while reading records", err)
    }
		return records
}

// create mapping of CSV headers to all values in column
// data is rows of CSV data read from a file with headers
func CsvDataByColumn(data [][]string) map[string][]float64 {
	headerIndex := make(map[int]string)
	dataByColumn := make(map[string][]float64)

	// Pop header data off of data and create mapping to csv data
	headers, data := data[0], data[1:]
	for i, header := range headers {
		headerIndex[i] = header
	}

	for _, row := range data {
		for i, v := range row {
			header := headerIndex[i]

			if value, err := strconv.ParseFloat(v, 64); err == nil {
				dataByColumn[header] = append(dataByColumn[header], value)
			}
		}
	}

	return dataByColumn
}
