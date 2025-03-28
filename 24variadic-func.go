
// variadic function

package main

import (
	"fmt"

)

func sumAll (num ...int) int {
	total := 0
	for _, value := range num {
		total += value
	}
	return total
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5)) 

	number := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(number...))
}