package main

import "fmt"

// interface kosong bisa berupa interface {} atau any
func Ups() interface{} {
	return "ups"
}

func main() {
	kossong := Ups()
	fmt.Println(kossong)
}