// interface
package main

import "fmt"

// interface adalah kontrak yang harus diikuti oleh struct
type HasName interface {
	GetName() string
	SetHeight() int

}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName(), "My height is", value.SetHeight())
}

type Individual struct {
	Name   string
	Height int
}

func (user Individual) GetName() string {
	return user.Name
}

func (user Individual) SetHeight() int {
	return user.Height
}

func main() {
	person := Individual{
		Name:   "Iik",
		Height: 170,
	}
	SayHello(person)
}