// named return value

package main

import "fmt"

func getFullName() (firstName, middleName, lastName string) {
	firstName = "Muhammad"
	middleName = "Iik"
	lastName = "Rin"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getFullName()
	fmt.Println("Nama saya adalah", firstName, middleName, lastName)
}