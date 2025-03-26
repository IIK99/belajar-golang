package main

import "fmt"

func main() {
	var a = 10
	// opsional menggunakan type data golang langsung tau bahwa itu int
	var b int = 20
	var c int = 30
	fmt.Println(a, b, c)
	a = a + b + c
	fmt.Println(a)
}
