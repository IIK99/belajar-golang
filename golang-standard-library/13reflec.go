package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"100"` // tag untuk validasi
	// Tag ini bisa digunakan untuk validasi, misalnya dengan library seperti go-playground/validator
}

type Person struct {
	Name 	string `required:"true" max:"100"`
	Address string `required:"true" max:"100"`
	Email 	string `required:"true" max:"100"` 
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		structField := valueType.Field(i)
		fmt.Println(structField.Name, "With type", structField.Type)
		fmt.Println(structField.Tag.Get("required"), "and max", structField.Tag.Get("max"))
	}
}

func main() {
	readField(Sample{Name: "Iik"})
	readField(Person{Name: "Rin", Address: "Tokyo", Email: "rin@example.com"})
}