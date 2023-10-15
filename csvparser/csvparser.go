package csvparser

import (
	"encoding/csv"
	"os"
	"path"
	"strconv"
)

type InvalidCsvPathError struct{
	Filepath string
}

func (i *InvalidCsvPathError) Error() string {
	return "Invalid CSV path: " + i.Filepath
}

func ReadCSV(filepath string) ([][]string, error) {
		// Ensure path is a csv file
		if path.Ext(filepath) != ".csv" {
			return nil, &InvalidCsvPathError{Filepath: filepath}
		}
		// Open file
		file, err := os.Open(filepath)
		if err != nil {
			return nil, err
    }
		// Defer keyword allows close call to be declared next to open call, but delays execution to end of function
		defer file.Close()

		// Read records from file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
    if err != nil {
			return nil, err
    }

		return records, err
}

// create mapping of CSV headers to all values in column
// data is rows of CSV data read from a file with headers
func CsvDataByColumn(data [][]string) (map[string][]float64, error) {
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
			value, err := strconv.ParseFloat(v, 64)

			if err == nil {
				dataByColumn[header] = append(dataByColumn[header], value)
			} else {
				return nil, err
			}
		}
	}

	return dataByColumn, nil
}
