// function as param

package main

import "fmt"

type Filter func(string) string // alias

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello ", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "****"
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Iik", spamFilter) 
	sayHelloWithFilter("Anjing", spamFilter)
}