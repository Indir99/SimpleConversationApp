package conversations

import "fmt"

// Conversion prompt functions
func SecondsToMinutes() {
	fmt.Printf("Enter time in seconds (s): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertSecondsToMinutes(input)
	fmt.Printf("%.2f s is %.2f min.\n", input, result)
}

func MinutesToHours() {
	fmt.Printf("Enter time in minutes (min): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertMinutesToHours(input)
	fmt.Printf("%.2f min is %.2f h.\n", input, result)
}

func HoursToDays() {
	fmt.Printf("Enter time in hours (h): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertHoursToDays(input)
	fmt.Printf("%.2f h is %.2f d.\n", input, result)
}

func DaysToWeeks() {
	fmt.Printf("Enter time in days (d): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertDaysToWeeks(input)
	fmt.Printf("%.2f d is %.2f wk.\n", input, result)
}

// Conversion functions
func convertSecondsToMinutes(seconds float64) float64 {
	return seconds / 60
}

func convertMinutesToHours(minutes float64) float64 {
	return minutes / 60
}

func convertHoursToDays(hours float64) float64 {
	return hours / 24
}

func convertDaysToWeeks(days float64) float64 {
	return days / 7
}
