package conversations

import "fmt"

// Conversion prompt functions
func CelsiusToFahrenheit() {
	fmt.Printf("Enter temperature in Celsius (°C): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertCelsiusToFahrenheit(input)
	fmt.Printf("%.2f °C is %.2f °F.\n", input, result)
}

func FahrenheitToCelsius() {
	fmt.Printf("Enter temperature in Fahrenheit (°F): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertFahrenheitToCelsius(input)
	fmt.Printf("%.2f °F is %.2f °C.\n", input, result)
}

func CelsiusToKelvin() {
	fmt.Printf("Enter temperature in Celsius (°C): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertCelsiusToKelvin(input)
	fmt.Printf("%.2f °C is %.2f K.\n", input, result)
}

func KelvinToCelsius() {
	fmt.Printf("Enter temperature in Kelvin (K): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertKelvinToCelsius(input)
	fmt.Printf("%.2f K is %.2f °C.\n", input, result)
}

// Conversion functions
func convertCelsiusToFahrenheit(celsius float64) float64 {
	return celsius*9/5 + 32
}

func convertFahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func convertCelsiusToKelvin(celsius float64) float64 {
	return celsius + 273.15
}

func convertKelvinToCelsius(kelvin float64) float64 {
	return kelvin - 273.15
}
