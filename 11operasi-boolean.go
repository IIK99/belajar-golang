
// operasi boolean
// operasi && (dan)
// nilai 1  	operator  	nilai 2  	hasil
// true 		&& 			true 		true  
// true 		&& 			false 		false
// false 		&& 			true 		false
// false 		&& 			false 		false

// operasi || (or)
// nilai 1  	operator  	nilai 2  	hasil
// true 		|| 			true 		true
// true 		|| 			false 		true
// false 		|| 			true 		true
// false 		|| 			false 		false

// operasi ! (not)
// nilai 1  	operator  	hasil
// true 		! 			false
// false 		! 			true

package main

import "fmt"

func main() {
	// operasi &&
	a := true
	b := true
	c := false
	d := false
	fmt.Println("operasi &&")
	fmt.Println(a && b)
	fmt.Println(a && c)
	fmt.Println(b && c)
	fmt.Println(a && d)
	fmt.Println(b && d)
	fmt.Println(c && d)

	// operasi ||
	e := true
	f := true
	g := false
	h := false
	fmt.Println("operasi ||")
	fmt.Println(e || f)
	fmt.Println(e || g)
	fmt.Println(f || g)
	fmt.Println(e || h)
	fmt.Println(f || h)
	fmt.Println(g || h)

	// operasi !
	i := true
	j := false
	fmt.Println("operasi !")
	fmt.Println(!i)
	fmt.Println(!j)

	// contoh lain
	nilaiAkhir := 80
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir > 90
	lulusAbsensi := absensi > 70

	lulus := lulusNilaiAkhir && lulusAbsensi
	fmt.Println("hasil", lulus)
}