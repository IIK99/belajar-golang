package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string `required:"true" max:"10"` // tag untuk validasi
	// Tag ini bisa digunakan untuk validasi, misalnya dengan library seperti go-playground/validator
}

type Person struct {
	Name 	string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email 	string `required:"true" max:"10"` 
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

func IsValid(value any) (result bool) {
	result = true
	// t = target, f = field
	t := reflect.TypeOf(value)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return result
}

func main() {
	readField(Sample{Name: "Iik"})
	readField(Person{"Rin", "", ""})

	person := Person {
		Name:    "Rin",
		Address: "Jakarta",
		Email:   "rin@example.com",
	}
	fmt.Println("Is Sample valid?", IsValid(person))	
	// pengecekan validasi ini bisa digunakan untuk validasi data yang masuk ke database
	// tanpa harus mengecek satu persatu fieldnya secara manual
	// tidak boleh ada field yang kosong
}