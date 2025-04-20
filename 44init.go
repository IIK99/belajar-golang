package main

import (
	"belajar-golang/database"
	// "belajar-golang/internal" ini error karena tidak digunakan dan ketika di save auto dihapus
	// tetapi jika ingin di panggil dan menjalankan init func bisa dengan cara ini
	_ "belajar-golang/internal"
	"fmt"
)

func main() {
	fmt.Println(database.GetConnection())

}