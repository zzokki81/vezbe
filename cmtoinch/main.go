package main

import (
	"fmt"
	"math"
)

type measures struct {
	feet int
	inch float64
}

func convert(cm float64) measures {

	inch := cm / Inches
	feet := inch / Foot

	a := math.Mod(inch, 12) //function calculate  %12

	calculate := measures{int(feet), a}
	return calculate

}

//Inches is 2.54 cm
const Inches = 2.54

//Foot is 12 inch
const Foot = 12

func main() {
	var cm float64
	fmt.Println("Please enter your number : ")
	fmt.Scan(&cm)
	i := convert(cm)
	fmt.Println(i)

}
