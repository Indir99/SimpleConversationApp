package conversations

import "fmt"

// Conversion prompt functions
func LitersToMilliliters() {
	fmt.Printf("Enter volume in liters (L): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertLitersToMilliliters(input)
	fmt.Printf("%.3f L is %.3f mL.\n", input, result)
}

func CubicMetersToLiters() {
	fmt.Printf("Enter volume in cubic meters (m³): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertCubicMetersToLiters(input)
	fmt.Printf("%.3f m³ is %.3f L.\n", input, result)
}

func GallonsToLiters() {
	fmt.Printf("Enter volume in gallons (gal): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertGallonsToLiters(input)
	fmt.Printf("%.3f gal is %.3f L.\n", input, result)
}

func CubicInchesToCubicCentimeters() {
	fmt.Printf("Enter volume in cubic inches (in³): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertCubicInchesToCubicCentimeters(input)
	fmt.Printf("%.3f in³ is %.3f cm³.\n", input, result)
}

// Conversion functions
func convertLitersToMilliliters(liters float64) float64 {
	const millilitersInLiter = 1000.0
	return liters * millilitersInLiter
}

func convertCubicMetersToLiters(cubicMeters float64) float64 {
	const litersInCubicMeter = 1000.0
	return cubicMeters * litersInCubicMeter
}

func convertGallonsToLiters(gallons float64) float64 {
	const litersInGallon = 3.78541
	return gallons * litersInGallon
}

func convertCubicInchesToCubicCentimeters(cubicInches float64) float64 {
	const cubicCentimetersInCubicInch = 16.3871
	return cubicInches * cubicCentimetersInCubicInch
}
