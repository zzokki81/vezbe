package main

import (
	"fmt"
	"math"
)

const (
	cmInInches   = 2.54
	inchesInFoot = 12
)

type centimeter float64

func (cm centimeter) convert() (int, float64) {

	inches := cm / cmInInches
	feet := inches / inchesInFoot
	floatinches := float64(inches)
	return int(feet), math.Mod(floatinches, 12)
}

func main() {

	cm := centimeter(10) //Type your desired length in centimeters
	fmt.Println(cm.convert())

}
