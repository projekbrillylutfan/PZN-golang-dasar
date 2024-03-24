package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	// address1 := Address{
	// 	"bogor",
	// 	"jawa barat",
	// 	"indonesia",
	// }
	// address2 := &address1

	// address2.city = "depok"

	// address2 = &Address{
	// 	"jakarta",
	// 	"DKI Jakarta",
	// 	"indonesia",
	// }

	// fmt.Println(address1)
	// fmt.Println(address2)

	// menggunakan asterisk

	address1 := Address{
		"bogor",
		"jawa barat",
		"indonesia",
	}
	address2 := &address1

	address2.city = "depok"

	// simbol asterisk * mengubah pointer yang baru
	*address2 = Address{
		"jakarta",
		"DKI Jakarta",
		"indonesia",
	}

	fmt.Println(address1)
	fmt.Println(address2)
}