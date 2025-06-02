// permission : https://chmod-calculator.com/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	// | jika lebih dari satu opsi, gunakan bitwise OR (|) untuk menggabungkan opsi tersebut.
	// 0666 adalah permission untuk file yang dibuat, ini berarti file tersebut dapat dibaca dan ditulis oleh pemilik, grup, dan pengguna lain.
	// O_CREATE digunakan untuk membuat file baru jika file tersebut belum ada.
	// O_WRONLY digunakan untuk membuka file hanya untuk penulisan.
	if err != nil {
		return err
	}
	defer file.Close() // pastikan file ditutup setelah selesai digunakan agar tidak terjadi memory leak
	file.WriteString(message) // menulis pesan ke file
	return nil
}

func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close() 
	file.WriteString(message) 
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close() // pastikan file ditutup setelah selesai digunakan

	reader := bufio.NewReader(file)
	var message string 
	for {
		line, _, err := reader.ReadLine() // membaca baris demi baris
		if err == io.EOF {
			break // jika sudah mencapai akhir file, keluar dari loop
		}
		message += string(line) + "\n" // menambahkan baris yang dibaca ke message
	}
	return message, nil // mengembalikan pesan yang dibaca dari file
	// fungsi ini akan membaca file dengan nama yang diberikan dan mengembalikan isi file sebagai string.
}

func main() {
	// Contoh penggunaan fungsi createNewFile
	// createNewFile("example.txt", "Hello, World!\nThis is a new file created using os.OpenFile.\n")
	// Fungsi createNewFile akan membuat file baru dengan nama "example.txt" dan menulis pesan ke dalamnya.
	// bisa juga penamaan file nya "example.log" atau "example.csv" atau "example.json" sesuai kebutuhan.
	
	// Contoh penggunaan fungsi addToFile
	addToFile("example.txt", "This is an additional line added to the file.\n")
	
	// Contoh penggunaan fungsi readFile
	result, _ := readFile("example.txt")
	// Fungsi readFile akan membaca file "example.txt" dan mengembalikan isi file tersebut sebagai string.
	fmt.Println("Isi file:", result)
	// Jika file tidak ditemukan, akan mengembalikan error. Tapi error nya saya ignore dengan menggunakan underscore (_).
}