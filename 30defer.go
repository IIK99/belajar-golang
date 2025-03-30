// function defer

package main

import "fmt"

func logging () {
	fmt.Println("Success login")
}

func runApplication() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApplication()
}