package conversations

import "fmt"

// Conversion prompt functions
func KilogramsToGrams() {
	fmt.Printf("Enter weight in kilograms (kg): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertKilogramsToGrams(input)
	fmt.Printf("%.3f kg is %.3f g.\n", input, result)
}

func PoundsToKilograms() {
	fmt.Printf("Enter weight in pounds (lb): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertPoundsToKilograms(input)
	fmt.Printf("%.3f lb is %.3f kg.\n", input, result)
}

func OuncesToGrams() {
	fmt.Printf("Enter weight in ounces (oz): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertOuncesToGrams(input)
	fmt.Printf("%.3f oz is %.3f g.\n", input, result)
}

func TonsToKilograms() {
	fmt.Printf("Enter weight in tons: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertTonsToKilograms(input)
	fmt.Printf("%.3f tons is %.3f kg.\n", input, result)
}

// Conversion functions
func convertKilogramsToGrams(kilograms float64) float64 {
	const gramsInKilogram = 1000.0
	return kilograms * gramsInKilogram
}

func convertPoundsToKilograms(pounds float64) float64 {
	const kilogramsInPound = 0.453592
	return pounds * kilogramsInPound
}

func convertOuncesToGrams(ounces float64) float64 {
	const gramsInOunce = 28.3495
	return ounces * gramsInOunce
}

func convertTonsToKilograms(tons float64) float64 {
	const kilogramsInTon = 1000.0
	return tons * kilogramsInTon
}
