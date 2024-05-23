package conversations

import "fmt"

// Conversion prompt functions
func BitsToBytes() {
	fmt.Printf("Enter data in bits: ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertBitsToBytes(input)
	fmt.Printf("%.2f bits is %.2f bytes.\n", input, result)
}

func KilobytesToMegabytes() {
	fmt.Printf("Enter data in kilobytes (KB): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertKilobytesToMegabytes(input)
	fmt.Printf("%.2f KB is %.2f MB.\n", input, result)
}

func MegabytesToGigabytes() {
	fmt.Printf("Enter data in megabytes (MB): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertMegabytesToGigabytes(input)
	fmt.Printf("%.2f MB is %.2f GB.\n", input, result)
}

func GigabytesToTerabytes() {
	fmt.Printf("Enter data in gigabytes (GB): ")
	input, err := ReadFloatInput()
	if err != nil {
		fmt.Println(err)
		return
	}
	result := convertGigabytesToTerabytes(input)
	fmt.Printf("%.2f GB is %.2f TB.\n", input, result)
}

// Conversion functions
func convertBitsToBytes(bits float64) float64 {
	return bits / 8
}

func convertKilobytesToMegabytes(kilobytes float64) float64 {
	return kilobytes / 1024
}

func convertMegabytesToGigabytes(megabytes float64) float64 {
	return megabytes / 1024
}

func convertGigabytesToTerabytes(gigabytes float64) float64 {
	return gigabytes / 1024
}
