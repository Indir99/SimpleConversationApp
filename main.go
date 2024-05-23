package main

import (
	"SimpleConversationApp/conversations"
	"fmt"
)

// TODO<IRK>: Test application and fix bugs (reminder: most cases are not covered)
// TODO<IRK>: Prepare Unit Tests for conversations package

func LengthDistanceConversation() {
	fmt.Println("You have selected an option: Length/Distance Conversations. Possible conversations are:")
	fmt.Println("\t 1. Meters (m) to Kilometers (km)")
	fmt.Println("\t 2. Meters (m) to Centimeters (cm)")
	fmt.Println("\t 3. Meters (m) to Millimeters (mm)")
	fmt.Println("\t 4. Feet (ft) to Inches (in)")
	fmt.Println("\t 5. Miles (mi) to Kilometers (km)")
	fmt.Println("\t 6. Yards (yd) to Meters (m)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.MetersToKilometers()
	case "2":
		conversations.MetersToCentimeters()
	case "3":
		conversations.MetersToMillimeters()
	case "4":
		conversations.FeetToInches()
	case "5":
		conversations.MilesToKilometers()
	case "6":
		conversations.YardsToMeters()
	default:
		HandleWrongInput(input)
	}
}

func AreaConversions() {
	fmt.Println("You have selected an option: Area Conversions. Possible conversations are:")
	fmt.Println("\t 1. Square Meters (m²) to Square Kilometers (km²)")
	fmt.Println("\t 2. Square Feet (ft²) to Square Inches (in²)")
	fmt.Println("\t 3. Acres to Square Meters (m²)")
	fmt.Println("\t 4. Hectares to Square Kilometers (km²)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.SquareMetersToSquareKilometers()
	case "2":
		conversations.SquareFeetToSquareInches()
	case "3":
		conversations.AcresToSquareMeters()
	case "4":
		conversations.HectaresToSquareKilometers()
	default:
		HandleWrongInput(input)
	}
}

func VolumeConversions() {
	fmt.Println("You have selected an option: Volume Conversions. Possible conversations are:")
	fmt.Println("\t 1. Liters (L) to Milliliters (mL)")
	fmt.Println("\t 2. Cubic Meters (m³) to Liters (L)")
	fmt.Println("\t 3. Gallons (gal) to Liters (L)")
	fmt.Println("\t 4. Cubic Inches (in³) to Cubic Centimeters (cm³)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.LitersToMilliliters()
	case "2":
		conversations.CubicMetersToLiters()
	case "3":
		conversations.GallonsToLiters()
	case "4":
		conversations.CubicInchesToCubicCentimeters()
	default:
		HandleWrongInput(input)
	}
}

func MassWeightConversions() {
	fmt.Println("You have selected an option: Mass/Weight Conversions. Possible conversations are:")
	fmt.Println("\t 1. Kilograms (kg) to Grams (g)")
	fmt.Println("\t 2. Pounds (lb) to Kilograms (kg)")
	fmt.Println("\t 3. Ounces (oz) to Grams (g)")
	fmt.Println("\t 4. Tons to Kilograms (kg)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.KilogramsToGrams()
	case "2":
		conversations.PoundsToKilograms()
	case "3":
		conversations.OuncesToGrams()
	case "4":
		conversations.TonsToKilograms()
	default:
		HandleWrongInput(input)
	}
}

func TemperatureConversions() {
	fmt.Println("You have selected an option: Temperature Conversations. Possible conversations are:")
	fmt.Println("\t 1. Celsius (°C) to Fahrenheit (°F)")
	fmt.Println("\t 2. Fahrenheit (°F) to Celsius (°C)")
	fmt.Println("\t 3. Celsius (°C) to Kelvin (K)")
	fmt.Println("\t 4. Kelvin (K) to Celsius (°C)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.CelsiusToFahrenheit()
	case "2":
		conversations.FahrenheitToCelsius()
	case "3":
		conversations.CelsiusToKelvin()
	case "4":
		conversations.KelvinToCelsius()
	default:
		HandleWrongInput(input)
	}
}

func SpeedConversions() {
	fmt.Println("You have selected an option: Speed Conversations. Possible conversations are:")
	fmt.Println("\t 1. Kilometers per Hour (km/h) to Miles per Hour (mph)")
	fmt.Println("\t 2. Meters per Second (m/s) to Kilometers per Hour (km/h)")
	fmt.Println("\t 3. Feet per Second (ft/s) to Meters per Second (m/s)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.KilometersPerHourToMilesPerHour()
	case "2":
		conversations.MetersPerSecondToKilometersPerHour()
	case "3":
		conversations.FeetPerSecondToMetersPerSecond()
	default:
		HandleWrongInput(input)
	}
}

func TimeConversions() {
	fmt.Println("You have selected an option: Time Conversations. Possible conversations are:")
	fmt.Println("\t 1. Seconds (s) to Minutes (min)")
	fmt.Println("\t 2. Minutes (min) to Hours (h)")
	fmt.Println("\t 3. Hours (h) to Days (d)")
	fmt.Println("\t 4. Days (d) to Weeks (wk)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.SecondsToMinutes()
	case "2":
		conversations.MinutesToHours()
	case "3":
		conversations.HoursToDays()
	case "4":
		conversations.DaysToWeeks()
	default:
		HandleWrongInput(input)
	}
}

func DataConversions() {
	fmt.Println("You have selected an option: Data Conversations. Possible conversations are:")
	fmt.Println("\t 1. Bits to Bytes")
	fmt.Println("\t 2. Kilobytes (KB) to Megabytes (MB)")
	fmt.Println("\t 3. Megabytes (MB) to Gigabytes (GB)")
	fmt.Println("\t 4. Gigabytes (GB) to Terabytes (TB)")
	fmt.Printf("Enter your choice: ")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		conversations.BitsToBytes()
	case "2":
		conversations.KilobytesToMegabytes()
	case "3":
		conversations.MegabytesToGigabytes()
	case "4":
		conversations.GigabytesToTerabytes()
	default:
		HandleWrongInput(input)
	}
}

func HandleWrongInput(s string) {
	fmt.Printf("You have entered the following choice: %q, but that choice is not valid. Please try again. \n", s)
}

func StartSimpleApp() {
	var input string
	fmt.Println("Welcome to Simple Go Application. Here you can do some interesting conversions. So let's start!")
	fmt.Println("Select which conversation you are interested in: ")
	fmt.Println("\t 1. Length/Distance Conversations")
	fmt.Println("\t 2. Area Conversions")
	fmt.Println("\t 3. Volume Conversions")
	fmt.Println("\t 4. Mass/Weight Conversions")
	fmt.Println("\t 5. Temperature Conversions")
	fmt.Println("\t 6. Speed Conversions")
	fmt.Println("\t 7. Time Conversions")
	fmt.Println("\t 8. Data Conversions")
	fmt.Printf("Enter your choice: ")
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		LengthDistanceConversation()
	case "2":
		AreaConversions()
	case "3":
		VolumeConversions()
	case "4":
		MassWeightConversions()
	case "5":
		TemperatureConversions()
	case "6":
		SpeedConversions()
	case "7":
		TimeConversions()
	case "8":
		DataConversions()
	default:
		HandleWrongInput(input)
	}
	//StartSimpleApp()
}

func main() {
	StartSimpleApp()
}
