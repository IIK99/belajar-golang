package main

import (
	"fmt"
	"os"
)

func main() {
	arg := os.Args
	// Args adalah agrumen yang diambil dari command line
	for _, arg := range arg {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error:", err.Error())
	}
}

// go run 3os.go Iik Muhammad Ikmal
// atau bisa juga 3os.go "Iik Muhammad Ikmal"