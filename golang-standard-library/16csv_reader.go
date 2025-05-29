// encoding, csv readers adalah paket yang digunakan untuk membaca dan menulis file CSV (Comma-Separated Values) dalam Go.
// Ini sangat berguna untuk mengimpor dan mengekspor data dalam format CSV, yang sering digunakan dalam aplikasi spreadsheet dan database.

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	// encoding/csv adalah paket yang digunakan untuk membaca dan menulis file CSV dalam Go.
	csvString := "FirstName, LastName, Age\n" +
		"Iik, Muhammad, 25\n" +
		"Rina, wati, 25"

		reader := csv.NewReader(strings.NewReader(csvString))

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break // End of file, exit the loop
			}

			fmt.Println("Record:", record)
		}
}