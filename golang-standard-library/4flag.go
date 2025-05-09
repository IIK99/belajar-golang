package main

import (
	"flag"
	"fmt"
)

func main() {
	// variabel nama, default, deskripsi
	username := flag.String("username", "root", "Your username")
	password := flag.String("password", "root", "Your password")
	hos := flag.String("host", "localhost", "Database host")
	port := flag.Int("port", 0, "Database port")

	// parsing flag
	flag.Parse()

	fmt.Println("Username:", *username)
	fmt.Println("Password:", *password)
	fmt.Println("Host:", *hos)
	fmt.Println("Port:", *port)
}

// cara memanggilnya wajib menggunakan - contoh:
// go run 4flag.go -username=Iik -password="Rin" -host="195.067.99.0" -port=0507