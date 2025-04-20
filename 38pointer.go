// pointer

package main

import "fmt"

type Address struct {
	City, Province, Country string
	ZipCode                  int
}

func main() {
	address1 := Address{"Jakarta", "DKI Jakarta", "Indonesia", 12345}
	address2 := address1 //copy by value

	address2.City = "Bandung" //change address2
	address2.Province = "Jawa Barat" 
	fmt.Println(address1) //address1 not changed
	fmt.Println(address2) //address2 changed

	// use pointer for copy by reference
	address3 := &address1 //& is asterisk operator, it means address3 is a pointer to address1
	*address3 = Address{"Tangerang Selatan", "Banten", "Indonesia", 12345} // don't forget use * for new value 
	fmt.Println(address1) //address1 changed
	fmt.Println(address3) //address3 changed
}

