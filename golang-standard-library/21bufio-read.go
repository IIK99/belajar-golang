package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	input := strings.NewReader("Hello, World!\nThis is a test.\nLet's read some lines.\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine() // membaca baris demi baris
		if err == io.EOF {
			break // jika sudah mencapai akhir file, keluar dari loop
		}

		fmt.Println(string(line)) // mencetak baris yang dibaca
	}
}