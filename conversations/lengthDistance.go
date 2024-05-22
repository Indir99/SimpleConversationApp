package conversations

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// TODO <IRK>: Create some kind of utils folder for similar things
// ReadFloatInput reads a float input from the user and returns an error if the input is invalid.
func ReadFloatInput() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	inputStr, err := reader.ReadString('\n')
	if err != nil {
		return 0.0, fmt.Errorf("error reading input: %v", err)
	}

	inputStr = strings.TrimSpace(inputStr)
	var input float64
	n, err := fmt.Sscanf(inputStr, "%f", &input)
	if err != nil {
		return 0.0, fmt.Errorf("error parsing input: %v", err)
	}
	if n != 1 {
		return 0.0, errors.New("invalid input. Please enter a valid number")
	}
	return input, nil
}

// MetersToKilometers prompts the user to enter a distance in meters and converts it to kilometers.
func MetersToKilometers() {
	fmt.Printf("Enter distance in meters: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertMetersToKilometers(input)
	fmt.Printf("%.3f m is %.6f km.\n", input, result)
}

// MetersToCentimeters prompts the user to enter a distance in meters and converts it to centimeters.
func MetersToCentimeters() {
	fmt.Printf("Eneter distance in meters: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		result := convertMetersToCentimeters(input)
		fmt.Printf("%.3f m is %.6f cm. \n", input, result)
	}
}

// MetersToMillimeters prompts the user to enter a distance in meters and converts it to millimeters.
func MetersToMillimeters() {
	fmt.Printf("Eneter distance in meters: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		result := convertMetersToMillimeters(input)
		fmt.Printf("%.3f m is %.6f mm. \n", input, result)
	}
}

// FeetToInches prompts the user to enter a distance in feet and converts it to inches.
func FeetToInches() {
	fmt.Printf("Enter distance in feet: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertFeetToInches(input)
	fmt.Printf("%.3f ft is %.3f in.\n", input, result)
}

// MilesToKilometers prompts the user to enter a distance in miles and converts it to kilometers.
func MilesToKilometers() {
	fmt.Printf("Enter distance in miles: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertMilesToKilometers(input)
	fmt.Printf("%.3f mi is %.3f km.\n", input, result)
}

// YardsToMeters prompts the user to enter a distance in yards and converts it to meters.
func YardsToMeters() {
	fmt.Printf("Enter distance in yards: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertYardsToMeters(input)
	fmt.Printf("%.3f yd is %.3f m.\n", input, result)
}

// convertMetersToKilometers converts a distance in meters to kilometers.
func convertMetersToKilometers(meters float64) float64 {
	return meters / 1000.0
}

// convertMetersToCentimeters converts distance in meters to centimeters
func convertMetersToCentimeters(meters float64) float64 {
	return meters * 100.0
}

// convertMetersToMillimeters converts distance in meters to millimeters
func convertMetersToMillimeters(meters float64) float64 {
	return meters * 1000.0
}

// convertFeetToInches converts a distance in feet to inches.
func convertFeetToInches(feet float64) float64 {
	const inchesInFoot = 12.0
	return feet * inchesInFoot
}

// convertMilesToKilometers converts a distance in miles to kilometers.
func convertMilesToKilometers(miles float64) float64 {
	const kilometersInMile = 1.60934
	return miles * kilometersInMile
}

// convertYardsToMeters converts a distance in yards to meters.
func convertYardsToMeters(yards float64) float64 {
	const metersInYard = 0.9144
	return yards * metersInYard
}
