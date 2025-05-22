package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println("Current Time: ", now)
	fmt.Println("Current Time: ", now.Local())

	UTC := time.Date(1999, time.July, 8, 10, 27, 33, 0, time.UTC)
	fmt.Println("UTC Time: ", UTC)
	fmt.Println("UTC Time: ", UTC.Local())

	formattedTime := "2006-01-02 15:04:05"

	value := "1999-07-08 10:27:33"
	// value := "terserah"
	valueTime, err := time.Parse(formattedTime, value)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Parsed Time: ", valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
}