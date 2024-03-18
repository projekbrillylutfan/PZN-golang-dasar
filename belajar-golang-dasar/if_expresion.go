package main

import "fmt"

func main() {
	// name := "brilly"

	// if name == "" {
	// 	fmt.Println("halo brilly")
	// } else if name == "lutfan" {
	// 	fmt.Println("halo lutfan")
	// } else {
	// 	fmt.Println("halo siapa?")
	// }

	// if short statment 
	name := "brilly"

	if length := len(name); length > 5 {
		fmt.Println("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}