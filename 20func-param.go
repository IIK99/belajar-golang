// function parameter

package main

import "fmt"

func greetPerson(name string, program string) {
		fmt.Println("Halo, saat ini", name, " saya sedang belajar ", program)
	}
	
	func main() {
		greetPerson("Iik", "golang") // memanggil 1x
		greetPerson("Rin", "golang") // memanggil 2x
	}