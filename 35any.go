// interface kosong atau any
package main

import "fmt"

func Ups() any {
	return "Hello"
	// return 123
	// return true
}

func main() {
	var kosong any = Ups()
	fmt.Println(kosong)
}