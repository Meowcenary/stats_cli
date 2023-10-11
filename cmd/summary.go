package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
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
		formatSummary([]string{"value", "income", "age", "rooms", "bedrooms", "pop", "hh"})
	},
}

func readCSV() [][]string {
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
func csvDataByColumn(data [][]string) map[string][]float64 {
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

func formatSummary(fields []string) {
	fmt.Println("formatSummary called")

	records := readCSV()
	data := csvDataByColumn(records)

	min_income, err := stats.Min(data["income"])
	if err == nil {
		fmt.Println("Min income: %f", min_income)
	}
	// for key, value := range data {
	// 	fmt.Printf("%s value is %v\n", key, value)
	// }

	// print requested stats for each field
	// TODO need to collect data into arrays that can be used by stats library
	// for _, value := range fields {
	// }

	// TODO add print headers function that prints out headers with spacing / delimiters
	// fmt.Println(strings.Trim(fmt.Sprint(fields), "[]"))
	// fmt.Println("Count: ", len(records)-1)
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
