// function dengan return

package main

import "fmt"

func sayHelloTo(name string, program string) string {
	return "Halo, saat ini " + name + " sedang belajar " + program
}

func main() {
	fmt.Println(sayHelloTo("Iik", "golang")) // memanggil 1x
	fmt.Println(sayHelloTo("Rin", "golang")) // memanggil 2x
}