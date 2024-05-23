package conversations

import "fmt"

// Conversion prompt functions
func KilometersPerHourToMilesPerHour() {
	fmt.Printf("Enter speed in kilometers per hour (km/h): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertKilometersPerHourToMilesPerHour(input)
	fmt.Printf("%.2f km/h is %.2f mph.\n", input, result)
}

func MetersPerSecondToKilometersPerHour() {
	fmt.Printf("Enter speed in meters per second (m/s): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertMetersPerSecondToKilometersPerHour(input)
	fmt.Printf("%.2f m/s is %.2f km/h.\n", input, result)
}

func FeetPerSecondToMetersPerSecond() {
	fmt.Printf("Enter speed in feet per second (ft/s): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertFeetPerSecondToMetersPerSecond(input)
	fmt.Printf("%.2f ft/s is %.2f m/s.\n", input, result)
}

// Conversion functions
func convertKilometersPerHourToMilesPerHour(kmph float64) float64 {
	const milesInKilometer = 0.621371
	return kmph * milesInKilometer
}

func convertMetersPerSecondToKilometersPerHour(mps float64) float64 {
	const kilometersPerHourInMeterPerSecond = 3.6
	return mps * kilometersPerHourInMeterPerSecond
}

func convertFeetPerSecondToMetersPerSecond(fps float64) float64 {
	const metersInFoot = 0.3048
	return fps * metersInFoot
}
