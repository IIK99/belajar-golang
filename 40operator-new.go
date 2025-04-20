// operator new tidak memiliki data awal

package main

import "fmt"

type NewAddress struct {
	City, Province, Country string
	ZipCode                  int
}

func main() {
	address1 := new(NewAddress)
	address2 := address1

	address2.City = "Jakarta"
	address2.Province = "DKI Jakarta"
	address2.Country = "Indonesia"
	address2.ZipCode = 12345

	fmt.Println(address1) // address1 changed
	fmt.Println(address2) // address2 changed
}