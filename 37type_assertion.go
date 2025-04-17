package main

import "fmt"

func random() any {
	// return "Oke"
	return 123
	// return true
}

func main() {
	var result any = random()
	// var resultString string = result.(string)
	// fmt.Print("Result: ", resultString)
	// fmt.Print("Result: ", result.(int)) // panic: interface conversion: interface {} is int, not string
	// var resultInt int = result.(int)
	// fmt.Print("Result: ", resultInt) 
	// panic: interface conversion: interface {} is int, not int

	switch value := result.(type) {
	case string:
		fmt.Print("Result: ", value)
	case int:
		fmt.Print("Result: ", value)
	default:
		fmt.Print("Unknown type", value)
	}
}