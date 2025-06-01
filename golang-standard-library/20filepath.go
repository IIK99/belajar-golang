// filepath di go menggunakan mac os atau linux, menggunakan tanda slash (/) sebagai pemisah direktori
// sedangkan di windows menggunakan backslash (\)

package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	fmt.Println(filepath.Dir("hello/world.go"))
	// path.Dir mengembalikan direktori dari path yang diberikan

	fmt.Println(filepath.Base("hello/world.go"))
	// world.go

	fmt.Println(filepath.Ext("hello/world.go"))
	// .go

	fmt.Println(filepath.IsAbs("hello/world.go"))
	// Mengembalikan path absolut dari file atau direktori yang diberikan, misal "C:\hello\world.go"

	fmt.Println(filepath.IsLocal("hello/world.go"))
	// Mengembalikan path lokal dari file atau direktori yang diberikan

	fmt.Println(filepath.Join("hello", "world", "main.go"))
}