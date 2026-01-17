package helper

import "testing"

// wajib nama function nya Test
// harus memiliki parameter (t *testing.T) dan tidak boleh ada return value


func TestHelloWord(t *testing.T) {
	result := HelloWord("Iik")

	if result != "Hello Iik" {
		// error
		t.Error("Result must be 'Hello Iik'")
	}
}

func TestHelloWordIkmal(t *testing.T) {
	result := HelloWord("Ikmal")

	if result != "Hello Ikmal" {
		// error
		t.Error("Result must be 'Hello Ikmal'")
	}
}

// cara run di terminal helper
// go test = run semua test di folder helper
// go test -v = run semua test function
// go test -run TestHelloWordIkmal = run test function TestHelloWordIkmal