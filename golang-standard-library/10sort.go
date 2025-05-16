package main

import (
	"fmt"
	"sort"
)

// membuat struct untuk menampung data
type User struct {
	Name string
	Age  int
}

// membuat slice untuk menampung data
type UserSlice []User

// membuat fungsi untuk mengurutkan data berdasarkan nama
func (u UserSlice) Len() int {
	return len(u)
}

// membuat fungsi untuk mengurutkan data berdasarkan nama
func (u UserSlice) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

// membuat fungsi untuk mengurutkan data berdasarkan age 
func (u UserSlice) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}


func main() {
	users := []User{
		{"Iik", 25},
		{"Muhammad", 30},
		{"Ikmal", 20},
		{"Rin", 19},
	}

	// mengurutkan data berdasarkan nama
	sort.Sort(UserSlice(users))

	// menampilkan data
	fmt.Println(users)
}