package conv

// Konverterer Fahrenheit til Celsius
func FahrenheitToCelsius(value float64) float64 {

	return (value - 32.0) * (5.0 / 9.0)
}

// Konverterer Celsius til Fahrenheit
func CelsiusToFahrenheit(value float64) float64 {

	return (value * 9 / 5) + 32
}

// Konverterer Kelvin til Fahrenheit
func KelvinToFahrenheit(value float64) float64 {

	return (value-273.15)*9/5 + 32
}

// Konverterer Fahrenheit  til Kelvin
func FahrenheitToKelvin(value float64) float64 {

	return (value-32)*5/9 + 273.15
}

// Konverterer Celsius  til Kelvin
func CelsiusToKelvin(value float64) float64 {

	return (value + 273.15)
}

// Konverterer Kelvin  til Celsius
func KelvinToCelsius(value float64) float64 {

	return (value - 273.15)
}
