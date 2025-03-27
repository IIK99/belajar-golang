// type data array

package main

import "fmt"

func main() {
	var caraCepat = [5]int{
		1, 2, 3, 4, 5,
	} 
	fmt.Println("ini cara cepat", caraCepat)
	fmt.Println("ambil index ke 2 adalah", caraCepat[2])

	// array kosong
	var a [5]int // [5] kapasitas array hanya 5 nilai tidak bisa ditambah
	fmt.Println(a)

	// array dengan tipe data string kosong
	var b [5]string
	fmt.Println(b)

	// array dengan tipe data dan tipe data string
	var c [5]string
	c[0] = "a"
	c[1] = "b"
	c[2] = "c"
	c[3] = "d"
	c[4] = "e"
	fmt.Println(c)

	// array dengan tipe data dan tipe data int
	var d [5]int
	d[0] = 1
	d[1] = 2
	d[2] = 3
	d[3] = 4
	d[4] = 5
	fmt.Println(d)

	// array yang tidak tentu jumlah data yang di input
	var e = []int{
		90, 55, 99,
	}
	fmt.Println(e)
	fmt.Println("panjang array", len(e))
	fmt.Println("data sebelum di ubah", e[0])
	e[0] = 30
	fmt.Println("mengubah value index ke 0", e)

}