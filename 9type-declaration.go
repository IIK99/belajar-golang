// declaration type data adalah sebuah statement yang digunakan untuk mendeklarasikan sebuah variabel
// contoh : int, string, float64, bool

package main

import "fmt"

func main() {
	// declaration type data
	type Name struct {
		nim int
		firstName string
		middleName string
		address string
		semester int
		gpa float64
		isStudent bool
	}

	// inisialisasi object
	obj := Name{
		nim: 123456789,
		firstName: "Iik",
		middleName: "Muhammad Ikmal",
		address: "Jl. Raya No. 123, Jakarta",
		semester: 6,
		gpa: 3.98,
		isStudent: true, 
	}

	fmt.Println(obj)
}