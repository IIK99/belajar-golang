package main

import (
	"fmt"
	"time"
)

func main() {
	duration1 := time.Duration(100 * time.Second)
	duration2 := time.Duration(100 * time.Millisecond)
	duration3 := time.Duration(duration1 - duration2)

	fmt.Println("Duration 1:", duration1)
	fmt.Println("Duration 2:", duration2)
	fmt.Println("Duration 3:", duration3)
}

// berguna saat koneksi ke database, mau berapa lama durasi nya