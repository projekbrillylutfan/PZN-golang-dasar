package main

import "fmt"

type Customer struct {
	Name, Address string
	age           int
}

func main() {
	var customer Customer
	customer.Name = "brilly"
	customer.Address = "jalan raya"
	customer.age = 22

	fmt.Println(customer)
	fmt.Println(customer.age)
	fmt.Println(customer.Address)
}