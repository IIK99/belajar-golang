package main

import "errors"

var (
	ValidationError = errors.New("validation error")
	NotFoundError = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	} 

	if id != "Iik" {
		return NotFoundError
	}

	// if success
	return nil
}

// Men check jenis error
func main() {
	err := GetById("2")
	// Isi dengan Iik, dan tidak ada error di terminal
	if err != nil {
		if errors.Is(err, ValidationError) {
			println("Validation Error")
			// string kosong 
		} else if errors.Is(err, NotFoundError) {
			println("Not Found Error")
			// id tidak ada
		} else {
			println("Unknown Error")
			// error tidak diketahui
		}
	}

}