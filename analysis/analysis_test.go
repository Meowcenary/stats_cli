package analysis

import (
	"strings"
	"testing"

	"github.com/montanaflynn/stats"
)

// func TestFormatSummary(t *testing.T) {
// 	// fields/columns to summarize
// 	fields := []string{}
// 	// headers ordered as they will be output
// 	order := []string{}
// }

func TestSummarizeFields(t *testing.T) {
	header := "TestHeader"
	data := map[string][]float64{header: {1, 2, 3,4, 5, 6, 7, 8, 9, 10}}

	// Three simple test cases
	headerOrder := []string{header}
	calculation := stats.Min
	summaryString := SummarizeFields(headerOrder, data, calculation)
  expectedMin := "1.000000"

	if strings.TrimSpace(summaryString) != expectedMin {
		t.Errorf(summaryString + "does not equal " + expectedMin)
	}

	headerOrder = []string{header}
	calculation = stats.Max
	summaryString = SummarizeFields(headerOrder, data, calculation)
	expectedMax := "10.000000"

	if strings.TrimSpace(summaryString) != expectedMax {
		t.Errorf(summaryString + "does not equal " + expectedMax)
	}

	headerOrder = []string{header}
	calculation = stats.Mean
	summaryString = SummarizeFields(headerOrder, data, calculation)
	expectedMean := "5.500000"

	if strings.TrimSpace(summaryString) != expectedMean {
		t.Errorf(summaryString + "does not equal " + expectedMean)
	}
}

func TestQuartiles(t *testing.T) {
	data := []float64{1, 2, 3,4, 5, 6, 7, 8, 9, 10}

	q1, _ := Q1(data)
	q2, _ := Q2(data)
	q3, _ := Q3(data)

	if q1 != 3 || q2 != 5.5 || q3 != 8 {
		t.Errorf("Quartiles are incorrect")
	}
}

func TestCount(t *testing.T) {
	data := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	count, _ := Count(data)

	if count != 5 {
		t.Errorf("Count does not match what was expected")
	}
}
