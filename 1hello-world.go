package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("Hello, Golang!")
}

// karena ini duplikasi main cara run nya go run .\hello-world2.go
// jika nge build jangan ada main yang sama dalam 1 package harus di folder terpisah 
// go mod init hello-world
// go build
// go run . (cara run di terminal)
// penamaan file go yang benar adalah huruf kecil semua, jika ada spasi diganti dengan - (strip) atau _ (underscore), tidak boleh ada nomor di depan nama file go
// nama func/type/interface wajib berawal dengan huruf besar (pascal case) 
// karena supaya bisa di import di package lain
// jika huruf kecil tidak bisa di import di package lain