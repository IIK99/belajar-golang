package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	// before test
	// contoh bisa digunakan untuk setup database
	fmt.Println("Before Test Main")
	m.Run()
	
	// after test
	// contoh bisa digunakan untuk cleanup database
	fmt.Println("Test Main Done")
}

// assertion = pengecekan apakah hasil sesuai dengan yang diharapkan
// install library/kunjungi assertion = go get github.com/stretchr/testify

// disarankan menggunakan kedua ini, sesuai kebutuhan
func TestHelloWorldAssert(t *testing.T) {
	result := HelloWord("Iik")
	assert.Equal(t, "Hello Iik", result, "Result must be 'Hello Iik'")
	// t, expected, actual, message
	fmt.Print("Test Hello with assert is success")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWord("Iik")
	require.Equal(t, "Hello Iik", result, "Result must be 'Hello Iik'")
	// t, expected, actual, message
	fmt.Print("Test Hello with require is success")
}

// go test -v -run TestHelloWorldRequire
// go test -v -run TestHelloWorldAssert

// perbedaan require dan assert
// require = error dan test function langsung berhenti
// assert = error tapi test function tetap lanjut

// skip test
func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Can't run on windows")
	}

	result := HelloWord("Iik")
	require.Equal(t, "Hello Iik", result, "Result must be 'Hello Iik'")
	fmt.Print("Test Hello with require is success")
}

// go test -v -run TestSkip

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

// perbedaan t.Error dan t.Fatal
// t.Error = error tapi test function tetap lanjut
// t.Fatal = error dan test function langsung berhenti
