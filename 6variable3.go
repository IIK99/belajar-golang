// bisa juga seperti ini := lebih simple langsung tanpa var tapi hanya di awal saja

package main

import "fmt"

func main() {
	name := "Belajar Go"
	fmt.Println(name)
	name = "Belajar Golang"
	fmt.Println(name)

	var (
		firstName = "Iik"
		middleName = "Muhammad"
		lastName = "Ikmal"
	)

	fmt.Println(firstName, middleName,)
	fmt.Println(lastName)
}