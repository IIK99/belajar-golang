// constant data yang tidak dapat diubah

package main

import "fmt"

func main() {
	const (
		firstName = "Iik"
		middleName = "Muhammad"
		lastName = "Ikmal"
		age1 = "Age :"
		age = 25
	)
	fmt.Println(firstName, middleName) 
	fmt.Println(lastName)
	fmt.Println(age1, age)
}
