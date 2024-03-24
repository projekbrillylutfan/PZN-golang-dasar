package main

import "fmt"

func random() any {
	return "ok"
}

func main() {
	result := random()
	// // value bisa di konfersi
	// resultString := result.(string)
	// fmt.Println(resultString)

	// // block kode ini pasti kena panic 
	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// pakai switch lebih aman
	switch value := result.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("unknown", value)
	}
}