// encoding, base64 adalah salah satu encoding yang umum digunakan untuk mengubah data biner menjadi string ASCII.
// Ini sering digunakan untuk mengirim data biner melalui protokol yang hanya mendukung teks, seperti email atau JSON.
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Iik Muhammad Ikmal"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println("Error decoding:", err.Error())
	} else {
		fmt.Println("Decoded:", string(decoded))
	}
}