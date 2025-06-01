package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Dir("hello/world.go"))
	// path.Dir mengembalikan direktori dari path yang diberikan
	
	fmt.Println(path.Base("hello/world.go"))
	// world.go
	
	fmt.Println(path.Ext("hello/world.go"))
	// .go

	fmt.Println(path.Join("hello", "world", "main.go"))
}