// code panic, digunakan saat ada error yang tidak diinginkan dan menghentikan program

package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	// recover menangkap data panic proses panic akan terhenti dan melanjutkan program
	message := recover()
	fmt.Println("Error message", message)
}

func runPanicApplication(error bool) {
	defer endApp()
	if error {
		panic("Application Error")
	}

}

func main() {
	runPanicApplication(false)
	// End App
	runPanicApplication(true)
	// End App
	// Application Error
}