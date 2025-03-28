// anonymous function

package main

import "fmt"

type BlackList func(string) bool

func registerUser(name string, blackList BlackList) bool {
	if blackList(name) {
		fmt.Println("You are blocked", name)
		return false
	} else {
		fmt.Println("Welcome", name)
		return true
	}
}

func main() {
	blackList := func(name string) bool {
		return name == "Anjing"
	}

	registerUser("Iik", blackList)
	registerUser("Anjing", blackList)
}