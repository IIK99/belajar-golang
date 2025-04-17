// struct
package main

import "fmt"

type Person struct {
	Name, Address string // langsung jika satu type
	// Address       string
	Number        int
	
}

// ini struct method
func (costumer Person) SayHello(Name string) {
	fmt.Println("Hello", Name, "My name is", costumer.Name)
}

func main() {
	var user Person
	fmt.Println(user) // {  0} kosong karena belum diisi
	user.Name = "Iik"
	user.Address = "Jakarta"
	user.Number = 123
	fmt.Println(user) // {Iik Jakarta 123}

	// bisa juga secara langsung
	user2 := Person{
		Name:    "Muhammad Iik",
		Address: "Nagasaki",
		Number:  789,
	}
	fmt.Println(user2)

	// atau seperti ini juga bisa
	user3 := Person{"Rin", "Soul", 456}
	fmt.Println(user3)
	
	// memanggil method struct
	user.SayHello("Hana")
	user2.SayHello("Rin")
	user3.SayHello("Iik")
}