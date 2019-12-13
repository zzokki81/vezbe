package tempfunctions

type (
	celsius    float64
	kelvin     float64
	fahrenheit float64
)

const (
	absoluteZeroC = 273.15
	absoluteZeroF = 459.67
)

func (c celsius) ToKelvinF() kelvin {
	return kelvin(c + absoluteZeroC)

}

func (c celsius) ToFahrenheit() fahrenheit {
	return fahrenheit((c * 9 / 5) + 32)

}

func (k kelvin) ToCelsius() celsius {
	return celsius(k - absoluteZeroC)
}
func (k kelvin) ToFahrenheit() fahrenheit {
	return fahrenheit((k * 9 / 5) - absoluteZeroF)

}
func (f fahrenheit) ToCelsius() celsius {
	return celsius((f - 32) * 5 / 9)

}
func (f fahrenheit) ToKelvin() kelvin {
	return kelvin((f + absoluteZeroF) * 5 / 9)

}
