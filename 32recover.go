// menangkap data panic proses panic akan terhenti dan melanjutkan program

package main

import "fmt"

func end() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Error message", message)
}

func runApp(error bool) {
	defer end()
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