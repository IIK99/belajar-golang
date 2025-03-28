// return multiple value

package main

import "fmt"

func sayHelloToMultiple(name string, program string) (string, string) {
	return "Halo, saat ini " + name + " sedang belajar " + program, ""
}

func main() {
	name, program := sayHelloToMultiple("Iik", "Golang") // memanggil 1x
	fmt.Println(name, program)
	name, program = sayHelloToMultiple("Rin", "Golang") // memanggil 2x
	fmt.Println(name, program)

	// return dengan _ yang digunakan param pertama
	name, _ = sayHelloToMultiple("Iik", "Golang") // memanggil 1x
	fmt.Println(name)
	name, _ = sayHelloToMultiple("Rin", "Golang") // memanggil 2x
	fmt.Println(name)
}