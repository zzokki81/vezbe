package main

import (
	"fmt"
	fun "github.com/zzokki81/vezbe/mean_median_iris/functions"

)

var SetosaList = [][]float64{}
var VersiColorList = [][]float64{}

func main() {

	list1 := fun.CsvReader("files/iris.csv")
	for _,v := range list1 {
		if v[4] == "Setosa" {
			fun.ConvertSlice(v, &SetosaList)
		} else if v[4] == "Versicolor" {
			fun.ConvertSlice(v, &VersiColorList)
		}
	}

	avgSetosa := fun.AverageValue(&SetosaList)
	avgVersiColor := fun.AverageValue(&VersiColorList)

	fmt.Printf("%12s %8s %7s %9s %7s\n","","s.length","s.width","p.length","p.width")
	fmt.Println("=====================================================")
	fun.PrintList("Setosa",avgSetosa)
	fun.PrintList("VersiColor",avgVersiColor)
}
