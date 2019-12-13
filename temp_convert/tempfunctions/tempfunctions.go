package tempfunctions

type (
	Celsius    float64
	Kelvin     float64
	Fahrenheit float64
)

const (
	absoluteZeroC = 273.15
	absoluteZeroF = 459.67
)

// ToKelvin method for convert Celsius to Kelvin
func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + absoluteZeroC)

}

// ToFahrenheit method for converting Celsius to Fahrenheit
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)

}

//ToCelsius method for converting Kelvin to Celsius
func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - absoluteZeroC)
}

// ToFahrenheit method for converting Kelvin to Fahrenheit
func (k Kelvin) ToFahrenheit() Fahrenheit {
	return Fahrenheit((k * 9 / 5) - absoluteZeroF)
}

//ToCelsius method for converting Fahrenheit to Celsius
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

//ToKelvin method for converting Fahrenheit to Kelvin
func (f Fahrenheit) ToKelvin() Kelvin {
	return Kelvin((f + absoluteZeroF) * 5 / 9)
}
