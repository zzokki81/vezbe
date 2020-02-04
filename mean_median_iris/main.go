package main

import (
	"fmt"

	cal "github.com/zzokki81/vezbe/mean_median_iris/mean_median"
	"github.com/zzokki81/vezbe/mean_median_iris/model/iris"
	pr "github.com/zzokki81/vezbe/mean_median_iris/printer"
	re "github.com/zzokki81/vezbe/mean_median_iris/reader"
)

var (
	setosaList     = [][]float64{}
	versiColorList = [][]float64{}
	virginicaList  = [][]float64{}
)

func addElementToList(list1 *iris.Iris, list2 *[][]float64) {
	var singleRow []float64
	singleRow = append(singleRow, list1.SepalLength, list1.SepalWidth, list1.PetalLength, list1.PetalWidth)
	*list2 = append(*list2, singleRow)
}

func main() {
	csvList := re.CsvReader("files/iris.csv")

	for _, v := range csvList {
		switch v.Variety {
		case "Setosa":
			addElementToList(v, &setosaList)
		case "Versicolor":
			addElementToList(v, &versiColorList)
		case "Virginica":
			addElementToList(v, &virginicaList)
		}
	}

	meanSetosa, meanVersiColor, meanVirginica := cal.MeanValue(&setosaList), cal.MeanValue(&versiColorList), cal.MeanValue(&virginicaList)
	medianSetosa, medianVersiColor, medianVirginica := cal.MedianValue(&setosaList), cal.MedianValue(&versiColorList), cal.MeanValue(&virginicaList)
	fmt.Printf("%12s %6s\n", "", "Mean")
	fmt.Println("====================================================")
	fmt.Printf("%12s %8s %7s %9s %7s\n", "", "s.length", "s.width", "p.length", "p.width")
	fmt.Println("=====================================================")

	pr.PrintListMean("Setosa", meanSetosa)
	pr.PrintListMean("VersiColor", meanVersiColor)
	pr.PrintListMean("Virginica", meanVirginica)

	fmt.Println("*****************************************************")
	fmt.Printf("%12s %6s\n", "", "Median")
	fmt.Println("=====================================================")

	pr.PrintListMedian("Setosa", medianSetosa)
	pr.PrintListMedian("VersiColor", medianVersiColor)
	pr.PrintListMedian("Virginica", medianVirginica)

}
