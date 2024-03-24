package main

import "fmt"

// penggunaan func type deklarasi
type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "babi" {
		return "..."
	} else {
		return name
	}
}

func main() {
	// bentuk pemanggilan pertama
	sayHelloWithFilter("brilly", spamFilter)

	// bentuk kedua
	filter := spamFilter
	sayHelloWithFilter("babi", filter)
}