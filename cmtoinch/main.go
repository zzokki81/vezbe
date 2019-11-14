package main

import (
	"fmt"
	"math"
)

type imperialLength struct {
	feet   int
	inches float64
}

const (
	cmInInches   = 2.54
	inchesInFoot = 12
)

type centimeter float64

func (cm centimeter) convert() imperialLength {

	inches := float64(cm / cmInInches)
	feet := inches / inchesInFoot
	return imperialLength{int(feet), math.Mod(inches, inchesInFoot)}

}
func main() {
	var cm float64
	fmt.Println("Type your desired length in centimeters : ")
	result := centimeter(cm)
	fmt.Scan(&result)
	fmt.Println(result, "Centimeters in Ft and Inch is : ", result.convert())

}
