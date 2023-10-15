package cmd

import (
	"log"

	"github.com/Meowcenary/stats_cli/analysis"
	"github.com/spf13/cobra"
	"github.com/Meowcenary/stats_cli/csvparser"
)

func NewSummaryCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "summary",
		Short: "Print summary statistics for a CSV file",
		Long: `Print summary statistics for a CSV file specified by
			the flag --file.

			Optionally specify CSV columns to include in output with --columns:
				stats_cli summary --file example.csv --columns column1,column2,...

			Optionally specify statistics to include and ordering to appear in summary
			output with --stats:
				stats_cli summary --file example.csv --stats count,mean,std,min,max
			`,
		Run: func(cmd *cobra.Command, args []string) {
			// handle flags
			file, _ := cmd.Flags().GetString("file")
			columns, _ := cmd.Flags().GetStringSlice("columns")
			statsorder, _ := cmd.Flags().GetStringSlice("stats")

			records, err := csvparser.ReadCSV(file)
			if err != nil {
				log.Fatal(err)
			}

			data, err := csvparser.CsvDataByColumn(records)
			if err != nil {
				log.Fatal(err)
			}

			// default to displaying all columns from csv in order they appear within file
			if len(columns) == 0 {
				columns = records[0]
			}
			// default to displaying all summary statistics
			if len(statsorder) == 0 {
				statsorder = []string{"count", "mean", "std", "stds", "min", "25%", "50%", "75%", "max"}
			}

			analysis.PrintSummary(
				data,
				columns,
				statsorder,
			)
		},
	}
}

// summaryCmd represents the summary command
var summaryCmd = NewSummaryCmd()

// used for flags
var (
	file string
	columns []string
	stats []string
)

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().StringVarP(&file, "file", "f", "", "CSV file to read from")
	summaryCmd.Flags().StringSliceVar(&columns, "columns", []string{}, "Columns to include from CSV for summary output")
	summaryCmd.Flags().StringSliceVar(&stats, "stats", []string{}, "Statistics to include for summary output")
}
