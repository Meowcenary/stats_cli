package cmd

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"
	"github.com/Meowcenary/stats_cli/csvparser"
	"github.com/montanaflynn/stats"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		FormatSummary([]string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"}, []string{"count", "mean", "std", "stds", "min", "25%", "50%", "75%", "max"})
	},
}

// headerOrder array of strings that correspond to headers in the csv, e.g ["value", "income", "age"]
// the order in the array is the order they will appear on the summary left to right
func Summarize(headerOrder []string, data map[string][]float64, calculation func(stats.Float64Data) (float64, error)) string {
	var summaryString string

	for _, header := range headerOrder {
		fieldData := data[header]
		result, _ := calculation(fieldData)
		summaryString += fmt.Sprintf("%-20f", result)
	}

	return summaryString
}

// because these functions are used by Summarize they must accept a single
// argument stats.Float64Data and return a float64 and erorr
func Count(data stats.Float64Data) (float64, error) {
	return float64(len(data)), nil
}

// func GetQuartiles() {
// 	return quartiles
// }
//
// func CalcQuartiles(input stats.Float64Data) {
// 	quartiles := stats.Quartile(input)
// }
//
// the stats library calculates all three quartiles and returns them as a struct
// to improve this create a global variable "quartiles" and a setter and getter
// "setQuartiles" and "getQuartiles". setQuartiles will calculate the quartiles and
// set the variable quartiles to the returned struct. Get quartiles will return the
// current variable for quartiles
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

func FormatSummary(fields []string, order []string) {
	records := csvparser.ReadCSV()
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
		fmt.Println(calculationName + ": " + Summarize(fields, data, calculations[calculationName]))
	}
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
