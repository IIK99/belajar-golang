// function

package main

import "fmt"

// ini adalah function
func sayHello() {
	fmt.Println("Halo!, saat ini saya sedang belajar golang")
}

// dan ini memanggil function
func main() {
	sayHello() // memanggil 1x
	sayHello() // memanggil 2x
}

// func main tidak boleh lebih dari 1

// function dengan parameter
// func greetPerson(name string, program string) {
// 	fmt.Println("Halo, saat ini", name, " saya sedang belajar ", program)
// }

// func main() {
// 	greetPerson("Iik", "golang") // memanggil 1x
// 	greetPerson("Rin", "golang") // memanggil 2x
// }

// function dengan return
// func sayHelloTo(name string, program string) string {
// 	return "Halo, saat ini " + name + " sedang belajar " + program
// }

// func main() {
// 	fmt.Println(sayHelloTo("Iik", "golang")) // memanggil 1x
// 	fmt.Println(sayHelloTo("Rin", "golang")) // memanggil 2x
// }