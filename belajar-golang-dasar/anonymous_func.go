package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}
}

func main() {
	// bentuk pertama
	blaclist := func(name string) bool {
		return name == "babi"
	}
	registerUser("brilly", blaclist)
	// bentuk kedua
	registerUser("babi", func(name string) bool {
		return name == "babi"
	})
}