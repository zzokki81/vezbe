package calculate

func MeanValue(list *[][]float64) []float64 {
	as1, as2, as3, as4 := 0.0, 0.0, 0.0, 0.0

	for _, v := range *list {
		as1 += v[0]
		as2 += v[1]
		as3 += v[2]
		as4 += v[3]
	}
	sl1 := len(*list)

	av1, av2, av3, av4 := as1/float64(sl1), as2/float64(sl1), as3/float64(sl1), as4/float64(sl1)
	var avgList []float64
	avgList = append(avgList, av1, av2, av3, av4)
	return avgList
}
