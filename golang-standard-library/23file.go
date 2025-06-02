// permission : https://chmod-calculator.com/

package main

import "os"

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

func main() {
	createNewFile("example.txt", "Hello, World!\nThis is a new file created using os.OpenFile.\n")
	// Fungsi createNewFile akan membuat file baru dengan nama "example.txt" dan menulis pesan ke dalamnya.
	// bisa juga penamaan file nya "example.log" atau "example.csv" atau "example.json" sesuai kebutuhan.
}