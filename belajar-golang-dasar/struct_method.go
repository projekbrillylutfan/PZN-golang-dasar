package main

import "fmt"

type Customer struct {
	Name string
}

func (customer Customer) sayHello(name string) {
	fmt.Println("hello", name, "my name is", customer.Name)
}

func main() {
	rully := Customer{Name: "rully"}
	rully.sayHello("brilly")
}