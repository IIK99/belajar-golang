// switch expression

package main

import "fmt"

func main() {
	// switch expression
	day := "friday"
	switch day {
	case "monday":
		fmt.Println("it is monday")
	case "tuesday":
		fmt.Println("it is tuesday")
	case "wednesday":
		fmt.Println("it is wednesday")
	case "thursday":
		fmt.Println("it is thursday")
	case "friday":
		fmt.Println("it is friday")
	case "saturday":
		fmt.Println("it is saturday")
	case "sunday":
		fmt.Println("it is sunday")
	default:
		fmt.Println("day is not valid")
	}

	// switch expression short statement
	day = "friday"
	switch day {
	case "monday", "tuesday", "wednesday", "thursday", "friday":
		fmt.Println("it is weekday")
	case "saturday", "sunday":
		fmt.Println("it is weekend")
	default:
		fmt.Println("day is not valid")
	}
}