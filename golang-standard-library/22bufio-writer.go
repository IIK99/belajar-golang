package main

import (
	"bufio"
	"os"
)

func main() {
	writer := bufio.NewWriter(os.Stdout) // os.Stdout adalah output standar, biasanya terminal atau target nya bisa berupa file
	_, _ = writer.WriteString("Hello, World!\n") // menulis string ke writer
	_, _ = writer.WriteString("This is a buffered writer example.\n") // menulis string lain ke writer
	writer.Flush() // memastikan semua data yang ditulis ke writer dikirim ke output
}