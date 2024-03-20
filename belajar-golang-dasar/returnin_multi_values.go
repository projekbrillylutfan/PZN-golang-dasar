package main

import "fmt"

func getFullName() (string, string) {
	return "brilly", "lutfan"
}

func main() {
	// returning func with multi values
	// firstName, lastName := getFullName()
	// fmt.Println(firstName, lastName)

	// menghiraukan returning values
	_, lastName := getFullName()
	fmt.Println(lastName)

}