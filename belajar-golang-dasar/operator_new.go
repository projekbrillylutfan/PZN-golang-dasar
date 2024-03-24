package main

import "fmt"

type Address struct {
	city, province, country string
}

func main() {
	var alamat1 *Address = new(Address)
	var alamat2 *Address = alamat1

	alamat2.country = "indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}