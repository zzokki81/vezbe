package printer

import (
	"fmt"
)

func PrintListMedian(name string, list []float64) {

	fmt.Printf("%-11s:|%-8.3f|\n", name, list[0])
}

func PrintListMean(name string, list []float64) {

	fmt.Printf("%-11s:|%-8.3f|%-8.3f|%-8.3f|%-8.3f|\n", name, list[0], list[1], list[2], list[3])
}
