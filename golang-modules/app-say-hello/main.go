package main

import (
	"fmt"

	gosayhello "github.com/IIK99/go-say-hello"
)

func main() {
	message := gosayhello.SayHello()
	fmt.Println(message)
}
