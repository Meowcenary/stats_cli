package analysis

import (
	"fmt"

	"github.com/Meowcenary/stats_cli/csvparser"
	"github.com/montanaflynn/stats"
)

func FormatSummary(fields []string, order []string) {
	records := csvparser.ReadCSV("housesInput.csv")
	data := csvparser.CsvDataByColumn(records)

	// all calculations available for summary
	calculations :=  map[string]func(stats.Float64Data) (float64, error){
		"count": Count,
		"mean": stats.Mean,
		"std": stats.StandardDeviation,
		"stds": stats.StandardDeviationSample,
		"min": stats.Min,
		"25%": Q1,
		"50%": Q2,
		"75%": Q3,
		"max": stats.Max,
	}

	for _, field := range fields {
		fmt.Printf("%20s", field)
	}
	fmt.Printf("\n")
	for _, calculationName := range order {
		fmt.Println(calculationName + ": " + SummarizeFields(fields, data, calculations[calculationName]))
	}
}

// headerOrder array of strings that correspond to headers in the csv, e.g ["value", "income", "age"]
// the order in the array is the order they will appear on the summary left to right
func SummarizeFields(headerOrder []string, data map[string][]float64, calculation func(stats.Float64Data) (float64, error)) string {
	var summaryString string

	for _, header := range headerOrder {
		fieldData := data[header]
		result, _ := calculation(fieldData)
		summaryString += fmt.Sprintf("%-20f", result)
	}

	return summaryString
}

// because these functions are used by Summarize they must accept a single
// argument stats.Float64Data and return a float64 and error i.e they need
// to be of type func(stats.Float64Data) (float64, error)
// It would be nice to not repeat stats.Quartile function call, but that would
// require a lot of refactoring that there isn't time for
func Q1(input stats.Float64Data) (float64, error) {
	quartiles, _ := stats.Quartile(input)
	return quartiles.Q1, nil
}

func Q2(input stats.Float64Data) (float64, error) {
	quartiles, _ := stats.Quartile(input)
	return quartiles.Q2, nil
}

func Q3(input stats.Float64Data) (float64, error) {
	quartiles, _ := stats.Quartile(input)
	return quartiles.Q3, nil
}

func Count(data stats.Float64Data) (float64, error) {
	return float64(len(data)), nil
}
