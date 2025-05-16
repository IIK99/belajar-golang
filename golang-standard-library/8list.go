package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()

	// menambahkan data ke dalam list
	data.PushBack("Iik")
	data.PushBack("Muhammad")
	data.PushBack("Ikmal")

	// menampilkan data
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// menghapus data dari list
	data.Remove(data.Front())

	// menampilkan data
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	// menambahkan data ke dalam list dari depan
	data.PushFront("Iik")

	// menampilkan data
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}