// converse type data dari satu type ke type lain
// contoh : int8 ke int32

package main

import "fmt"

func main() {
	var a int32 = 137856
	var b int64 = int64(a)
	var c int16 = int16(b)
	fmt.Println(b)
	fmt.Println(c)
	
	var name = "Iik"
	var e uint8 = name[0]
	var eString = string(e)
	fmt.Println(name, e, eString)
}