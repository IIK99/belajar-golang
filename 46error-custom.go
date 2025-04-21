package main

import "fmt"

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{Message: "ID tidak boleh kosong"}
	}

	if id != "Iik" {
		return &NotFoundError{Message: "Data tidak ditemukan"}
	}

	// Jika oke
	return nil
}

func main() {
	err := SaveData("Iik", nil) // ID kosong, random, Iik = data tidak ditemukan
	if err != nil {
		// // terjadi error
		// if ValidationError, ok := err.(*ValidationError); ok {
		// 	fmt.Println("Validation Error:", ValidationError.Error())
		// } else if NotFoundError, ok := err.(*NotFoundError); ok {
		// 	fmt.Println("Not Found Error:", NotFoundError.Error())
		// } else {
		// 	fmt.Println("Unknown Error:", err.Error())
		// }

		// menggunakan switch case
		switch finalErr := err.(type) {
			case *ValidationError:
				fmt.Println("Validation Error:", finalErr.Error())
			case *NotFoundError:
				fmt.Println("Not Found Error:", finalErr.Error())
			default:
				fmt.Println("Unknown Error:", finalErr.Error())
		}
	} else {
		// sukses
		fmt.Println("Data berhasil disimpan")
	}
}