// if expression

package main

import "fmt"

func main() {
	a := 10
	if a > 5 {
		fmt.Println("a is greater than 5")
	}

	if a < 5 {
		fmt.Println("a is less than 5")
	}

	if a >= 5 {
		fmt.Println("a is greater than or equal to 5")
	}

	if a <= 5 {
		fmt.Println("a is less than or equal to 5")
	}	

	if a == 5 {
		fmt.Println("a is equal to 5")
	}

	if a != 5 {
		fmt.Println("a is not equal to 5")
	}

	// if else expression, license to drive
	age := 17
	if age >= 18 {
		fmt.Println("driver license")
	} else {
		fmt.Println("not enough age to drive")
	}
	
	// if else if expression, license to drive
	age = 17
	if age >= 18 {
		fmt.Println("driver license")
	} else if age >= 17 {
		fmt.Println("high school driver license")
	} else {
		fmt.Println("not enough age to drive")
	}

	// if short statement
	name := "exampleUser"
	if userName := len(name); userName > 8 {
		fmt.Println("username is too long")
	} else {
		fmt.Println("username is valid")
	}
}