// nil atau null

package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"Name": name,
		}
	}
}

func main() {
	data := NewMap("Rin")

	if data == nil {
		fmt.Println("Tidak ada data")
	} else {
		fmt.Print("Data: ", data["Name"])
	}
}