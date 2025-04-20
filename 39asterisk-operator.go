// asterisk operator

package main

import "fmt"

type MyAddress struct {
	City, Province, Country string
	ZipCode                  int
}

func main() {
	address1 := MyAddress{"Jakarta", "DKI Jakarta", "Indonesia", 12345}
	address2 := &address1
	address2 = &MyAddress{"Bandung", "Jawa Barat", "Indonesia", 12345} 
	fmt.Println(address1) // address1 not changed
	fmt.Println(address2) // address2 changed
	*address2 = MyAddress{"Tangerang Selatan", "Banten", "Indonesia", 12345}
	// * is asterisk operator, it means address2 is a pointer to address1
	fmt.Println(address1) // address1 changed
	fmt.Println(address2) // address2 changed
}