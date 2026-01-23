package helper

import (
	"fmt"
	"testing"
)

// wajib nama function nya Test
// harus memiliki parameter (t *testing.T) dan tidak boleh ada return value


func TestHelloWord(t *testing.T) {
	result := HelloWord("Iik")

	if result != "Hello Iik" {
		// error
		t.Error("Result must be 'Hello Iik'")
	}

	fmt.Print("Ini menggunakan t.Error")
}

func TestHelloWordIkmal(t *testing.T) {
	result := HelloWord("Ikmal")

	if result != "Hello Ikmal" {
		// error
		t.Fatal("Result must be 'Hello Ikmal'")
	}

	fmt.Print("Ini menggunakan t.Fatal")
}

// cara run di terminal helper
// go test = run semua test di folder helper
// go test -v = run semua test function
// go test -run TestHelloWordIkmal = run test function TestHelloWordIkmal

// perbedaan t.Error dan t.Fatal
// t.Error = error tapi test function tetap lanjut
// t.Fatal = error dan test function langsung berhenti
