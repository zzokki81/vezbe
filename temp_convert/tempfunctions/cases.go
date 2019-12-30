package tempfunctions

import (
	"errors"
	"fmt"
)

//func Cases() convert user temperature to other two formats.
//It can be used to convert between three temperature format C,K,F
func Cases() error {

	temperature, unit, err := Parse()
	if err != nil {
		return err
	}
	var defaultCase string

	switch unit {
	case "c":
		celsius := Celsius(temperature)
		fmt.Printf("> %v Celsius is %v Fahrenheit\n", temperature, celsius.ToFahrenheit())
		fmt.Printf("> %v Celsius is %v Kelvin\n", temperature, celsius.ToKelvin())
	case "f":
		fahrenheit := Fahrenheit(temperature)
		fmt.Printf("> %v Fahrenheit is %.2f Celsius\n", temperature, fahrenheit.ToCelsius())
		fmt.Printf("> %v Fahrenheit is %.2f Kelvin\n", temperature, fahrenheit.ToKelvin())
	case "k":
		kelvin := Kelvin(temperature)
		fmt.Printf("> %v Kelvin is %.2f Celsius\n", temperature, kelvin.ToCelsius())
		fmt.Printf("> %v Kelvin is %.2f Fahrenheit\n", temperature, kelvin.ToFahrenheit())
	default:
		defaultCase = "Temperature unit is not recognized!!"

	}
	if defaultCase != "" {
		return errors.New(defaultCase)
	}
	return nil
}
