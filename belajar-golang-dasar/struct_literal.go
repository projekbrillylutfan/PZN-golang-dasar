package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func main() {
	joko := Customer{
		Name:    "joko",
		Address: "jalan raya",
		age:     22,
	}

	fmt.Println(joko)

	budi := Customer{
		"budi",
		"jalan raya 2",
		30,
	}
	fmt.Println(budi)
}
