package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "brilly"
	lastName = "lutfan"
	middleName = "mulyadi"

	return firstName, middleName, lastName
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}