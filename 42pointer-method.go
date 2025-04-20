package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	iik := Man{"Iik"}
	iik.Married()

	fmt.Println(iik.Name)
}

// method sangat di recommend menggunakan pointer