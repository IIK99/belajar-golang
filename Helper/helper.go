// sebetulnya di dalam go 1 project hanya 1 main dan 1 func main, tidak boleh lebih dari 1
// maka bisa kita akali harus sesuai nama package dengan file nya

package helper

// nama func/type/interface wajib berawal dengan huruf besar (pascal case)
// karena supaya bisa di import di package lain
// jika huruf kecil tidak bisa di import di package lain
func SayHello(name string) string {
	return "Hello " + name

}

// ini memanggilnya di file import