package main

import "fmt"

func main() {
	// variable bisa langsung di deklarasikan tanpa tipe data jika variable sudah ada data nya
	// var name = "brilly"
	// fmt.Println(name)

	// name = "lutfan"
	// fmt.Println(name)

	// //tanpa var
	// nama := "brilly2"
	// fmt.Println(nama)

	// name = "lutfan2"
	// fmt.Print(nama)

	// multiple variable
	// var (
	// 	firstName = "brilly"
	// 	lastName = "lutfan"
	// )

	// fmt.Println(firstName)
	// fmt.Println(lastName)

	// constant
	// variabel kan harus dipakai, kalau const kagak perlu, yang penting sudah ada nilai

	// const name = "brilly"
	// const age = 22

	// // bakalan error nih
	// name = "lutfan"
	// age = 23

	// multiple constant
	const (
		firstName = "brilly"
		lastName = "lutfan"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}