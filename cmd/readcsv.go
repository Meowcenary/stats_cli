package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// readcsvCmd represents the readcsv command
var readcsvCmd = &cobra.Command{
	Use:   "readcsv",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Open file for reading
		// TODO: Generalize to a file path
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
		for _, eachrecord := range records {
        fmt.Println(eachrecord)
    }
	},
}

func init() {
	rootCmd.AddCommand(readcsvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// readcsvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// readcsvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
