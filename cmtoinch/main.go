package main

import (
	"fmt"
	"math"
)

const cmInInches = 2.54
const inchesInFoot = 12

type imperialLength struct {
	feet   int
	inches float64
}

func convert(cm float64) imperialLength {

	inches := cm / cmInInches
	feet := inches / inchesInFoot
	mouduoInches := math.Mod(inches, 12)
	return imperialLength{int(feet), mouduoInches}

}

func main() {
	var cm float64
	fmt.Println("Please enter your number : ")
	fmt.Scan(&cm)
	i := convert(cm)
	fmt.Println(i)

}
