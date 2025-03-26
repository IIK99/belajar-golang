package main

import "fmt"

func main() {
	fmt.Println("Hello, World2!")
	fmt.Println("Hello, Golang2!")
}

// karena ini duplikasi main cara run nya go run .\hello-world2.go
// jika nge build jangan ada main yang sama dalam 1 package harus di folder terpisah 
// go mod init hello-world
// go build
// go run . (cara run di terminal)