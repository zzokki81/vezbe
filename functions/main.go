package main

import (
	"fmt"
	"math"
)

func main() {

	var x float64 = 4
	m := MySqrt(x)

	fmt.Println("Squeare root is", m)

}

//MySqrt calculate squere root
func MySqrt(x float64) float64 {
	return math.Sqrt(x)
}
