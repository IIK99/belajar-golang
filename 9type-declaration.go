// declaration type data adalah sebuah statement yang digunakan untuk mendeklarasikan sebuah variabel
// contoh : int, string, float64, bool

package main

import "fmt"

func main() {
	// declaration type data
	type Name struct {
		firstName string
		middleName string
		lastName string
		age int
		height float64
		isStudent bool
	}

	// inisialisasi object
	obj := Name{
		firstName: "Iik",
		middleName: "Muhammad",
		lastName: "Ikmal",
		age: 25,
		height: 1.78,
		isStudent: true,
	}

	fmt.Println(obj)
}