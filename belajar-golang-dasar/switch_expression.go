package main

import "fmt"

func main() {
	name := "Ekooooo"

	// switch name {
	// case "Eko":
	// 	fmt.Println("hello eko")
	// case "Budi":
	// 	fmt.Println("hello budi")
	// default:
	// 	fmt.Println("hiya hiya hiya")

	// switch short statment
	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	//switch tanpa expression
	switch length := len(name); {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
