// for expression

package main

import (
	"fmt"
)

func main() {
	
	counter := 0
	for counter < 5 {
		counter++
		fmt.Println("perhitungan ke: ", counter)
	}
	fmt.Println("looping berakhir")
	
	// for dengan statement
	for i := 0; i < 5; i++ {
		fmt.Println("perhitungan ke: ", i)
	}
	fmt.Println("looping berakhir")

	// for range
	// manual
	name := []string{"Iik", "Rin", "Ina"}
	for i := 0; i < len(name); i++ {
		fmt.Println("nama: ", name[i])
	}

	// for range
	for index, value := range name {
		fmt.Println("index ke: ", index, "= nama: ", value)
	}

	// for range dengan underscore
	for _, value := range name {
		fmt.Println("nama: ", value)
	}
}