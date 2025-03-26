// operasi matematika
// + untuk tambah
// - untuk kurangi
// * untuk kali
// / untuk bagi
// % untuk modulus
// augmented assignment adalah sebuah operator yang digunakan untuk melakukan operasi matematika
// contoh : += -= *= /= %= 
// unary operator adalah sebuah operator yang digunakan untuk melakukan operasi matematika
// contoh : ++ ditambah 1, -- kurangi 1, - negatif, + positif, ! not
// operator perbandingan adalah sebuah operator yang digunakan untuk melakukan perbandingan
// contoh : == sama dengan, != tidak sama, > lebih besar, >= lebih besar atau sama, 
// < lebih kecil, <= lebih kecil atau sama
// operator ternary adalah sebuah operator yang digunakan untuk menanyakan kondisi
// contoh : ? jika benar : jika salah

package main

import "fmt"

func main() {
	// operasi matematika
	a := 10
	b := 20
	c := 30
	d := 40

	fmt.Println(a + b + c + d)
	fmt.Println(a - b - c - d)
	fmt.Println(a * b * c * d)
	fmt.Println(a / b / c / d)
	fmt.Println(a % b % c % d)

	// augmented assignment
	e := 10
	e += 20
	fmt.Println("ini augmented assignment +", e)
	f := 20
	f -= 30
	fmt.Println("ini augmented assignment -", f)
	g := 2
	g *= 40
	fmt.Println("ini augmented assignment *", g)
	h := 5
	h /= 50
	fmt.Println("ini augmented assignment /", h)
	i := 6
	i %= 60
	fmt.Println("ini augmented assignment %", i)

	// unary operator
	j := 10
	j++
	fmt.Println("ini unary operator ++", j)
	k := 20
	k--
	fmt.Println("ini unary operator --", k)
	l := 30
	l = -l
	fmt.Println("ini unary operator -", l)
	m := 40
	m = +m
	fmt.Println("ini unary operator +", m)
	n := false
	n = !n
	fmt.Println("ini unary operator !", n)

	// operator perbandingan
	o := 10
	p := 20
	q := 30
	r := 40
	s := 50
	t := 60

	fmt.Println(o == p)
	fmt.Println(o != p)
	fmt.Println(o > p)
	fmt.Println(o >= p)
	fmt.Println(o < p)
	fmt.Println(o <= p)
	fmt.Println(o == q)
	fmt.Println(o != q)
	fmt.Println(o > q)
	fmt.Println(o >= q)
	fmt.Println(o < q)
	fmt.Println(o <= q)
	fmt.Println(o == r)
	fmt.Println(o != r)
	fmt.Println(o > r)
	fmt.Println(o >= r)
	fmt.Println(o < r)
	fmt.Println(o <= r)
	fmt.Println(o == s)
	fmt.Println(o != s)
	fmt.Println(o > s)
	fmt.Println(o >= s)
	fmt.Println(o < s)
	fmt.Println(o <= s)
	fmt.Println(o == t)
	fmt.Println(o != t)
	fmt.Println(o > t)
	fmt.Println(o >= t)
	fmt.Println(o < t)	
	fmt.Println(o <= t)	

}