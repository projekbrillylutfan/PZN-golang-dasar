package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	// pass by value
	// address1 := Address{
	// 	"bogor",
	// 	"jawa barat",
	// 	"indonesia",
	// }
	// address2 := address1

	// address2.city = "depok"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// pass by reference pointer
	address1 := Address{
		"bogor",
		"jawa barat",
		"indonesia",
	}
	// pointer harus menggunakan simbol & pada saat pembuatan variabel
	address2 := &address1

	address2.city = "depok"
	fmt.Println(address1)
	fmt.Println(address2)
}