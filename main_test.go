package main

import (
	"testing"

	"github.com/Meowcenary/stats_cli/analysis"
	"github.com/Meowcenary/stats_cli/csvparser"
)


// The intention was to put the benchmark in cmd/summary_test.go, but
// there were problems that arose so in the interest of time the benchmark
// was moved to this file
func BenchmarkSummary(b *testing.B) {
	inputfile := "testdata/housesInput.csv"
	outputfile := "testdata/houseOutputGo.csv"
	statsorder := []string{"count", "mean", "std", "stds", "min", "25%", "50%", "75%", "max"}
	records, _ := csvparser.ReadCSV(inputfile)
	columns := records[0]
	data, _ := csvparser.CsvDataByColumn(records)

	fieldsSummary := analysis.FormatSummary(
		data,
		columns,
		statsorder,
	)

	csvparser.WriteCSV(outputfile, fieldsSummary)
}
