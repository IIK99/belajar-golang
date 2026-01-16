package main

import (
	"fmt"

	gosayhello "github.com/IIK99/go-say-hello/v2"
)

func main() {
	message := gosayhello.SayHello("Ikmal")
	fmt.Println(message)
}
