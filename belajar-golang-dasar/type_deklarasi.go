package main

import "fmt"

func main() {
	type noKtp string

	var ktpBilly noKtp = "1234567890"
	fmt.Println(ktpBilly)
	fmt.Print(noKtp("9999999999"))
}