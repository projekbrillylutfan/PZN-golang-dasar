package main

import "fmt"

type Person struct {
	Name string
}

type HasName interface {
	getName() string
}

type Animal struct {
	Name string
}

func (person Person) getName() string {
	return person.Name
}

func SayHello(hasName HasName) {
	fmt.Println("hello", hasName.getName())
}

func (animal Animal) getName() string {
	return animal.Name
}

func main() {
	person := Person{
		Name: "brilly",
	}
	SayHello(person)

	animal := Animal{
		Name: "kucing",
	}
	SayHello(animal)
}