package analysis

import (
	"fmt"
	"strings"

	"github.com/montanaflynn/stats"
)

func AvailableStats() map[string]func(stats.Float64Data) (float64, error) {
	return map[string]func(stats.Float64Data) (float64, error){
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
}

// internal function to find the longest name of the stats calculations for formatting
func longestStatName(statnames []string) int {
	max := 0

	for _, name := range statnames {
		length := len(name)
		if length  > max {
			max = length
		}
	}

	return max
}

// print summary statistics for data set
// data , map of column/field to decimals
// columns , columns to include in summary
// statsorder , ordered statsitical calculations to include in summary
func PrintSummary(data map[string][]float64, columns []string, statsorder []string) {
	fmt.Printf("%s", FormatSummary(data, columns, statsorder))
}

func FormatSummary(data map[string][]float64, columns []string, statsorder []string) string {
	var summary string

	for _, field := range columns {
		summary += fmt.Sprintf("%20s", field)
	}
	summary += fmt.Sprintf("\n")

	calculations := AvailableStats()
	maxNameLen := longestStatName(statsorder)
	for _, calculationName := range statsorder {
		summary += fmt.Sprintf("%s: %s%s\n", calculationName, strings.Repeat(" ", maxNameLen-len(calculationName)), SummarizeFields(columns, data, calculations[calculationName]))
	}

	return summary
}

// columns array of strings that correspond to columns in a csv, e.g ["value", "income", "age"]
// the order columns appear in the array is the order they will appear on the summary left to right
func SummarizeFields(columns []string, data map[string][]float64, calculation func(stats.Float64Data) (float64, error)) string {
	var summaryString string

	for _, column := range columns {
		fieldData := data[column]
		result, _ := calculation(fieldData)
		summaryString += fmt.Sprintf("%20f", result)
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
