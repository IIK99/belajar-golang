// encoding csv writer adalah paket yang digunakan untuk menulis data ke dalam format CSV (Comma-Separated Values) dalam Go.
// Ini sangat berguna untuk mengekspor data ke dalam format yang mudah dibaca oleh spreadsheet atau aplikasi lain yang mendukung CSV.

package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	// hasilnya akan muncul di terminal bukan ke file csv, nanti di I/O
	_ = writer.Write([]string{"FirstName", "LastName", "Age"})
	_ = writer.Write([]string{"Iik", "Muhammad", "25"})
	_ = writer.Write([]string{"Rina", "wati", "25"})
	
	writer.Flush()
	// Flush digunakan untuk memastikan semua data yang ditulis ke writer dikirim ke output.
	// Ini penting karena writer mungkin menyimpan data dalam buffer sebelum menulisnya ke output.
	if err := writer.Error(); err != nil {
		panic(err)
	}
}