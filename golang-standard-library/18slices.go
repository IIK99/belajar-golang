package main

import (
	"fmt"
	"slices"
)

func main() {
	name := []string{"Iik", "Ikmal", "Rina", "wati"}
	values := []int{25, 30, 25, 20}

	fmt.Println(slices.Min(name))
	fmt.Println(slices.Max(name))
	fmt.Println(slices.Contains(name, "Iik"))	
	fmt.Println(slices.Contains(name, "Rina"))
	fmt.Println(slices.Contains(name, "someone"))
	fmt.Println(slices.Index(name, "someone"))
	fmt.Println(slices.Index(name, "Rina"))
	fmt.Println(slices.Index(name, "wati"))
	fmt.Println(slices.Contains(values, 50))


}