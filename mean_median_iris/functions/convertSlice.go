package functions

import (
	"fmt"
	"strconv"
)

func ConvertSlice(list1 []string, list2 *[][]float64) {
	var list0 []float64
	for index := range list1 {
		if index > 3 {
			break
		}
		value, err := strconv.ParseFloat(list1[index], 64)
		if err != nil {
			fmt.Print(err)
		}
		list0 = append(list0, value)
	}
	*list2 = append(*list2, list0)
}