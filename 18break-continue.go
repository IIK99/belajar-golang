// break and continue

package main

import (
	"fmt"
)

func main() {

	// break
	counter := 0
	for counter < 5 {
		counter++
		if counter == 3 {
			break
		}
		fmt.Println("perhitungan ke: ", counter)
	}
	fmt.Println("looping berakhir")

	// continue
	counter = 0
	for counter < 5 {
		counter++
		if counter == 3 {
			continue
		}
		fmt.Println("perhitungan ke: ", counter)
	}
	fmt.Println("looping berakhir")

	// continue with %
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("perhitungan ke: ", i)
	}
	fmt.Println("hanya bilangan ganjil")

	// break dan continue
	counter = 0
	for counter < 10 {
		counter++
		if counter == 3 {
			continue
		}
		if counter == 5 {
			break
		}
		fmt.Println("perhitungan ke: ", counter)
	}
	fmt.Println("looping berakhir")
}