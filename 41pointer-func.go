package main

import "fmt"

type NewMyAddress struct {
	City, Province, Country string
	ZipCode                  int
}

func ChangeAddress(address *NewMyAddress) { // add * address is a pointer to NewMyAddress
	address.Country = "Indonesia"
}

func main() {
	address := NewMyAddress {"Jakarta", "DKI Jakarta", "", 12345}
	ChangeAddress(&address) // pass address to ChangeAddress function
	// & is asterisk operator, it means address is a pointer to NewMyAddress
	fmt.Println(address) 
}