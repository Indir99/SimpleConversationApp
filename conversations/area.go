package conversations

import (
	"fmt"
)

// Conversion prompt functions
func SquareMetersToSquareKilometers() {
	fmt.Printf("Enter area in square meters (m²): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := ConvertSquareMetersToSquareKilometers(input)
	fmt.Printf("%.3f m² is %.6f km².\n", input, result)
}

func SquareFeetToSquareInches() {
	fmt.Printf("Enter area in square feet (ft²): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := ConvertSquareFeetToSquareInches(input)
	fmt.Printf("%.3f ft² is %.3f in².\n", input, result)
}

func AcresToSquareMeters() {
	fmt.Printf("Enter area in acres: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := ConvertAcresToSquareMeters(input)
	fmt.Printf("%.3f acres is %.3f m².\n", input, result)
}

func HectaresToSquareKilometers() {
	fmt.Printf("Enter area in hectares: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := ConvertHectaresToSquareKilometers(input)
	fmt.Printf("%.3f hectares is %.6f km².\n", input, result)
}

// Conversion functions
func ConvertSquareMetersToSquareKilometers(squareMeters float64) float64 {
	const squareMetersInSquareKilometer = 1e6
	return squareMeters / squareMetersInSquareKilometer
}

func ConvertSquareFeetToSquareInches(squareFeet float64) float64 {
	const squareInchesInSquareFoot = 144.0
	return squareFeet * squareInchesInSquareFoot
}

func ConvertAcresToSquareMeters(acres float64) float64 {
	const squareMetersInAcre = 4046.86
	return acres * squareMetersInAcre
}

func ConvertHectaresToSquareKilometers(hectares float64) float64 {
	const squareKilometersInHectare = 0.01
	return hectares * squareKilometersInHectare
}
