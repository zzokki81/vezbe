package tempfunctions

type (
	Celsius    float64
	Kelvin     float64
	Fahrenheit float64
)

const (
	ABSOLUTEZEROC = 273.15
	ABSOLUTEZEROF = 459.67
	CONVERTCF = 0.556
	CONVERTFC = 1.8
	ZEROCINFAHR = 32
)

// ToKelvin method for convert Celsius to Kelvin
func (c Celsius) ToKelvin() Kelvin {
	return Kelvin(c + ABSOLUTEZEROC)


}

// ToFahrenheit method for converting Celsius to Fahrenheit
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * CONVERTFC) + ZEROCINFAHR)

}

//ToCelsius method for converting Kelvin to Celsius
func (k Kelvin) ToCelsius() Celsius {
	return Celsius(k - ABSOLUTEZEROC)
}

// ToFahrenheit method for converting Kelvin to Fahrenheit
func (k Kelvin) ToFahrenheit() Fahrenheit {
	return Fahrenheit((k * CONVERTFC) - ABSOLUTEZEROF)
}

//ToCelsius method for converting Fahrenheit to Celsius
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - ZEROCINFAHR) *CONVERTCF )
}

//ToKelvin method for converting Fahrenheit to Kelvin
func (f Fahrenheit) ToKelvin() Kelvin {
	return Kelvin((f + ABSOLUTEZEROF) * CONVERTCF )
}
