// closure, jangan terlalu banyak memakai closure, itu akan membuat bingung

package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	increment()
	fmt.Println("counter", counter)
}