package main

import "fmt"

// next steps: 1. prepare package "conversations"
// 			   2. prepare lengthDistance.go with corresponding functions
//			   3. prepare area.go with corresponding functions
//			   4. prepare volume.go with corresponding functions
//			   5. prepare masswight.go with corresponding functions
//			   6. prepare temperature.go with corresponding functions
//			   7. prepare speed.go with corresponding functions
//			   8. prepare time.go with corresponding functions
//			   9. prepare data.go with corresponding functions
//			   10. prepare error handling

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
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	case "5":
		// call function
	case "6":
		// call function
	default:
		// call error handling function
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
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	default:
		// call error handling function
	}
}

func VolumeConversions() {
	fmt.Println("You have selected an option: Volume Conversions. Possible conversations are:")
	fmt.Println("\t 1. Liters (L) to Milliliters (mL)")
	fmt.Println("\t 2. Cubic Meters (m³) to Liters (L)")
	fmt.Println("\t 3. Gallons (gal) to Liters (L)")
	fmt.Println("\t 4. Cubic Inches (in³) to Cubic Centimeters (cm³)")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	default:
		// call error handling function
	}
}

func MassWeightConversions() {
	fmt.Println("You have selected an option: Mass/Weight Conversions. Possible conversations are:")
	fmt.Println("\t 1. Kilograms (kg) to Grams (g)")
	fmt.Println("\t 2. Pounds (lb) to Kilograms (kg)")
	fmt.Println("\t 3. Ounces (oz) to Grams (g)")
	fmt.Println("\t 4. Tons to Kilograms (kg)")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	default:
		// call error handling function
	}
}

func TemperatureConversions() {
	fmt.Println("You have selected an option: Temperature Conversations. Possible conversations are:")
	fmt.Println("\t 1. Celsius (°C) to Fahrenheit (°F)")
	fmt.Println("\t 2. Fahrenheit (°F) to Celsius (°C)")
	fmt.Println("\t 3. Celsius (°C) to Kelvin (K)")
	fmt.Println("\t 4. Kelvin (K) to Celsius (°C)")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	default:
		// call error handling function
	}
}

func SpeedConversions() {
	fmt.Println("You have selected an option: Speed Conversations. Possible conversations are:")
	fmt.Println("\t 1. Kilometers per Hour (km/h) to Miles per Hour (mph)")
	fmt.Println("\t 2. Meters per Second (m/s) to Kilometers per Hour (km/h)")
	fmt.Println("\t 3. Feet per Second (ft/s) to Meters per Second (m/s)")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// call function
	case "2":
		// call function
	case "3":
		// call function
	default:
		// call error handling function
	}
}

func TimeConversions() {
	fmt.Println("You have selected an option: Time Conversations. Possible conversations are:")
	fmt.Println("\t 1. Seconds (s) to Minutes (min)")
	fmt.Println("\t 2. Minutes (min) to Hours (h)")
	fmt.Println("\t 3. Hours (h) to Days (d)")
	fmt.Println("\t 4. Days (d) to Weeks (wk)")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	default:
		// call error handling function
	}
}

func DataConversions() {
	fmt.Println("You have selected an option: Data Conversations.P ossible conversations are:")
	fmt.Println("\t 1. Bits to Bytes")
	fmt.Println("\t 2. Kilobytes (KB) to Megabytes (MB)")
	fmt.Println("\t 3. Megabytes (MB) to Gigabytes (GB)")
	fmt.Println("\t 4. Gigabytes (GB) to Terabytes (TB)")
	var input string
	fmt.Scanf("%s", &input)
	switch input {
	case "1":
		// call function
	case "2":
		// call function
	case "3":
		// call function
	case "4":
		// call function
	default:
		// call error handling function
	}
}

func HandleWrongInput(s string) {
	fmt.Printf("You have entered the following choice: %q, but the required input is the range 1 to 8. Please try again. \n", s)
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
