package csvparser

import (
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	data := ReadCSV("../testdata/easy.csv")
	expected_data := [][]string{
		{"Easy", "Test", "Data"},
		{"123", "456", "789"},
		{"101", "121", "131"},
		{"415", "161", "171"},
		{"181", "920", "212"},
	}

	if !reflect.DeepEqual(data, expected_data) {
		t.Errorf("Data read from test_data/easy_data.csv does not match what was expected")
	}
}

func TestCsvDataByColumn(t *testing.T) {
	data := [][]string{
		{"Easy", "Test", "Data"},
		{"123", "456", "789"},
		{"101", "121", "131"},
		{"415", "161", "171"},
		{"181", "920", "212"},
	}

	expectedDataByColumn := map[string][]float64{
		"Easy": {123.0, 101.0, 415.0, 181.0,},
		"Test": {456.0, 121.0, 161.0, 920.0},
		"Data": {789.0, 131.0, 171.0, 212.0},
	}

	dataByColumn := CsvDataByColumn(data)

	if !reflect.DeepEqual(dataByColumn, expectedDataByColumn) {
		t.Errorf("Parsed data does not match what was expected")
	}
}
