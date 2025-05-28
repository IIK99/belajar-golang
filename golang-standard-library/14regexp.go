// package regular expression
// untuk mencari pola tertentu di dalam string

package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`r([a-z])n`)

	fmt.Println(regex.MatchString("Rin"))
	fmt.Println(regex.MatchString("ikmal"))
	fmt.Println(regex.MatchString("Ikmal"))
	fmt.Println(regex.MatchString("rin"))

	fmt.Println(regex.FindAllString("iik ikmal Ikmal rin Rin", 5))
}