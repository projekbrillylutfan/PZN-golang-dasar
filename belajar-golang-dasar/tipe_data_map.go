package main

import "fmt"

func main() {
	// person := map[string]string{
	// 	"name":    "brilly",
	// 	"address": "jalan raya",
	// }

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "brilly"
	book["ups"] = "salah"

	delete(book, "ups")

	fmt.Println(book)
}