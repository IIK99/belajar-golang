// format atau parsing

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// convert string ke bool
	result, err := strconv.ParseBool("true")
	// jika diisi dengan "salah" maka akan error
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println("Result:", result)
	}

	// convert bool ke string
	resultString := strconv.FormatBool(result)
	fmt.Println("Result String:", resultString)
}