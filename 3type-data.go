// gunakan sesuai kebutuhan agar memori komputer aman

// type data integer
// type data	nilai minimum			nilai maksimal	
// int8			-128					127	
// int16		-32768					32767	
// int32		-2147483648				2147483647	
// int64		-9223372036854775808	9223372036854775807	

// type data integer 2
// type data	nilai minimum			nilai maksimal	
// uint8		0						255	
// uint16		0						65535	
// uint32		0						4294967295	
// uint64		0						18446744073709551615	

// type data float
// type data	nilai minimum				nilai maksimal	
// float32		-3.4028235e+38				3.4028235e+38	
// float64		-1.7976931348623157e+308	1.7976931348623157e+308	

// type data alias
// type data	alias untuk type data lain	
// string		[]byte						
// bool			bool						
// byte 		uint8						
// rune			int32						
// int			minimal int32				
// uint			minimal uint32			

// type data boolean	
// bool			false					true	

package main

import "fmt"

func main() {
	fmt.Println("ini adalah boolean", true)
	fmt.Println("ini adalah ASCII", 'a')
	fmt.Println("ini adalah int8", int8(-10))
	fmt.Println("ini adalah uint8", uint8(7))
	fmt.Println("ini adalah float32", float32(-1.20))
	fmt.Println("ini adalah float64", float64(1.64))
	fmt.Println("ini adalah string", string("hello"))
	fmt.Println(len("hello"))
	fmt.Println("ini adalah string"[0])
	fmt.Println("ini adalah string"[0:3])
	fmt.Println("ini adalah alias byte", byte(99))
	fmt.Println("ini adalah alias rune", rune(1234567))
	fmt.Println("ini adalah alias int", int(78965431))	
	fmt.Println("ini adalah alias uint", uint(7765498))
}

// 🔍 Bedanya 'a' dan "a" dalam Go
// Tipe	Deklarasi	Hasil	 Output
// Rune (int32)		'a'		 97
// String			"a"	 	 a 
// Jadi, 'a' di Go lebih mirip dengan angka (kode ASCII/Unicode) dibandingkan string "a".