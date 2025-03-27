// type data map 

package main

import "fmt"

func main() {
	// map person
	person := map[string]string{
		"name": "Iik",
		"address": "Jakarta",
	}
	fmt.Println("nama", person["name"])
	fmt.Println("alamat", person["address"])
	fmt.Println(person)

	// map person dengan tipe data int
	personInt := map[string]int{
		"no phone": 123,
		"code area": 2339,
	}
	fmt.Println("no phone", personInt["no phone"])
	fmt.Println("code area", personInt["code area"])
	fmt.Println(personInt)

	// function map 
	profile := make(map[string]int)
	profile["weight"] = 55
	profile["height"] = 170
	profile["age"] = 25

	fmt.Println(profile)
	delete(profile, "age")
	fmt.Println(profile)

}