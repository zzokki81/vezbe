package calculate

func MedianValue(list *[][]float64) []float64 {
	as2, as3 := 0.0, 0.0
	for _, v := range *list {
		as2 += v[1]
		as3 += v[2]
	}
	var medList []float64
	aMed := (as2 + as3) / 2
	medList = append(medList, aMed)
	return medList
}
