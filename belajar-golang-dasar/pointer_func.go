package main

import "fmt"

type Address struct {
	city, province, country string
}

// penanda pointer di function yaitu ada tanda *
func ChangeAddressToIndonesia(address *Address) {
	address.country = "indonesia"
}

func main() {
	address := Address{
		"bogor",
		"jawa barat",
		"",
	}
	ChangeAddressToIndonesia(&address)

	fmt.Println(address)
}