package main

import "fmt"

func main() {
	var cm float64 = 15
	f, i := convert(cm)
	fmt.Println(cm, " cm  is: ", f, "ft")
	fmt.Println(cm, " cm  is: ", i, "inch")

}

func convert(cm float64) (float64, float64) {
	inch := cm / 2.54
	feet := cm / 30.48

	return feet, inch
}
