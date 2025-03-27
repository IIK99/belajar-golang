// type data slice
// contoh slice dari nama bulan

package main

import (
	"fmt"
)

func main() {
	mount := []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"}
	
	slice1 := mount[4:7]
	fmt.Println(slice1)

	slice2 := mount[:5]
	fmt.Println(slice2)

	slice3 := mount[6:]
	fmt.Println(slice3)

	slice4 := mount[:]
	fmt.Println(slice4)

	mount[0] = "januari"
	mount[1] = "februari"
	fmt.Println(mount)

	mountSlice := append(mount, "Happy New Year")
	fmt.Println(mountSlice)
	mountSlice[0] = "12 month jan"
	fmt.Println(mountSlice)
	fmt.Println(mount)
}