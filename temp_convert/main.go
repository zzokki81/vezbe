package main

import (
	"fmt"
	"vezbe/temp_convert/tempfunctions"
)

func main() {

	celsius := tempfunctions.Celsius(52)
	kelvin := tempfunctions.Kelvin(5)
	fahrenheit := tempfunctions.Fahrenheit(80)

	fmt.Printf("%v Celsius is %v Kelvin or %v Fahrenheit\n ", celsius, celsius.ToKelvin(), celsius.ToFahrenheit())
	fmt.Printf("%v Kelvin is %v Celsius or %v Fahrenheit\n", kelvin, kelvin.ToCelsius(), kelvin.ToFahrenheit())
	fmt.Printf("%v Fahrenheit is %.2f Celsius or  %.2f Kelvin\n ", fahrenheit, fahrenheit.ToCelsius(), fahrenheit.ToKelvin())
}
